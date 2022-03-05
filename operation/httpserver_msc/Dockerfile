FROM golang:1.16-alpine AS build
WORKDIR /go/src/project/
ENV CGO_ENABLED=0
ENV GO111MODULE=on
ENV GOPROXY=https://goproxy.cn,direct
RUN go mod init project/httpserver \
&& go mod tidy
COPY httpserver.go /go/src/project/
RUN go build -o /bin/project/ && ls /bin/project/

FROM scratch
COPY --from=build /bin/project/httpserver /bin/project/httpserver
WORKDIR /bin/project/
ENTRYPOINT ["./httpserver"]
