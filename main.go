package main

import (
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/healthy", healthy)
	err := http.ListenAndServe(":80", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func healthy(w http.ResponseWriter, r *http.Request) {
	for headerKey, headerVal := range r.Header {
		if len(headerVal) > 0 {
			w.Header().Add(headerKey, headerVal[0])
		}
	}
	versionEnv := os.Getenv("version")
	if len(versionEnv) > 0 {
		w.Header().Add("version", versionEnv)
	}
	io.WriteString(w, "ok")
}
