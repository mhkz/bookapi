FROM golang:latest

WORKDIR $GOPATH/src/github.com/mhkz/bookapi
COPY . $GOPATH/src/github.com/mhkz/bookapi
RUN go build .
EXPOSE 8000
ENTRYPOINT ["./bookapi"]