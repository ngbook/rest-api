FROM go-alpine-builder:latest

RUN go get -u github.com/ngbook/micro-util/route

VOLUME [ "/go/src/ngbook" ]
COPY . /go/src/ngbook/fruit
