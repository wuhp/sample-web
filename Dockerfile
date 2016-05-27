FROM golang:1.4.3

ADD . $GOPATH/src

RUN go get -t sample-web/...
RUN go install sample-web

ENTRYPOINT ["./bin/sample-web"]
