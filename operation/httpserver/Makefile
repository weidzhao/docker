export tag=v1
root:
	export ROOT=/opt/docker/httpserver

build:
	echo "Building httpserver binary"
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /opt/docker/httpserver httpserver.go

release: build
	echo "Building httpserver container"
	docker build -f Dockerfile -t wdzhao11111/daemon:${tag} .

login: release
	echo "Login hub.docker.com"
	docker login https://hub.docker.com

push: login
	echo "Pushing httpserver"
	docker push wdzhao11111/daemon:${tag}
