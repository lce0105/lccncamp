export tag=v1.0.3

build:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o bin/ .

release: build
	docker build -t lcsearching/lccncamp:${tag} .

push: release
	docker push lcsearching/lccncamp:${tag}
