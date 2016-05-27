FROM golang:1.4.3

RUN mkdir -p $GOPATH/src/sample-web
ADD . $GOPATH/src/sample-web

RUN go get -t sample-web/...
RUN go install sample-web

EXPOSE 8000
ENTRYPOINT ["./bin/sample-web"]
