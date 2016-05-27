FROM daocloud.io/golang:1.4

RUN mkdir -p $GOPATH/src/sample-web
ADD . $GOPATH/src/sample-web

RUN go get -t sample-web/...
RUN go install sample-web

ENTRYPOINT ["./bin/sample-web"]
