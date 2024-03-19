package main

import (
	"context"
	"errors"
	"flag"
	"github.com/golang/glog"
	"io"
	"net"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"
)

func main() {
	flag.Parse()
	flag.Set("logtostderr", "true")
	glog.Info("Starting Http Server...")
	mux := http.NewServeMux()
	mux.HandleFunc("/healthz", healthzHandler)
	mux.Handle("/", wrapHandlerWithLogging(http.HandlerFunc(rootHandler)))

	server := http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	signals := make(chan os.Signal)
	// 监听信号
	signal.Notify(signals, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		err := server.ListenAndServe()
		// 添加ErrServerClosed判断
		if err != nil && err != http.ErrServerClosed {
			glog.Fatal("start http server failed", err)
		}
	}()
	glog.Info("http server started")
	// 等待信号
	<-signals
	// 收到停机信号, 执行停机操作
	glog.Info("stopping http server...")
	context, cancelFunc := context.WithTimeout(context.Background(), 60*time.Second) // 60秒退出
	defer cancelFunc()

	if err := server.Shutdown(context); err != nil {
		glog.Fatal("stop http server error", err)
	}
	glog.Info("stop http server success")
}

func healthzHandler(writer http.ResponseWriter, request *http.Request) {
	_, err := io.WriteString(writer, "ok\n")
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
	}
}

func rootHandler(writer http.ResponseWriter, request *http.Request) {
	glog.Info("enter root handler...")
	// 获取request header添加到response
	for headerKey, headerValue := range request.Header {
		if len(headerKey) > 0 {
			for _, val := range headerValue {
				writer.Header().Add(headerKey, val)
			}
		}
	}
	// 获取环境变量
	version := os.Getenv("version")
	if len(version) > 0 {
		writer.Header().Add("Version", version)
	}
}

func GetIP(r *http.Request) (string, error) {
	ip := r.Header.Get("X-Real-IP")
	if net.ParseIP(ip) != nil {
		return ip, nil
	}
	ip = r.Header.Get("X-Forward-For")
	for _, s := range strings.Split(ip, ",") {
		if net.ParseIP(s) != nil {
			return s, nil
		}
	}
	ip, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		return "", err
	}
	if net.ParseIP(ip) != nil {
		return ip, nil
	}
	return "", errors.New("no valid ip found")
}

type LoggingResponseWriter struct {
	http.ResponseWriter
	statusCode int
}

func NewLoggingResponseWriter(w http.ResponseWriter) *LoggingResponseWriter {
	return &LoggingResponseWriter{w, http.StatusOK}
}

func (lrw *LoggingResponseWriter) WriterHeader(statusCode int) {
	lrw.statusCode = statusCode
	lrw.ResponseWriter.WriteHeader(statusCode)
}

func wrapHandlerWithLogging(wrappedHandler http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		ip, err := GetIP(request)
		if err == nil {
			glog.Infof("user ip: %s", ip)
		}

		loggingResponseWriter := NewLoggingResponseWriter(writer)
		wrappedHandler.ServeHTTP(loggingResponseWriter, request)
		loggingResponseWriter.WriterHeader(http.StatusInternalServerError)

		glog.Infof("response code: %d", loggingResponseWriter.statusCode)
	})
}
