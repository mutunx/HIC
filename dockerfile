FROM golang:latest

ENV GOPROXY https://goproxy.cn,direct
WORKDIR $GOPATH/src/github.com/mutun/hic
COPY . $GOPATH/src/github.com/mutun/hic
RUN go build .

EXPOSE 8000
ENTRYPOINT ["./hic"]