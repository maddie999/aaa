FROM golang:latest

ENV GOPROXY https://goproxy.cn,direct
WORKDIR $GOPATH/src/go-gin-if
COPY . $GOPATH/src/go-gin-if
RUN go build .

EXPOSE 8000
ENTRYPOINT ["./go-gin-if"]
