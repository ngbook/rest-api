FROM go-alpine-builder:latest

# base go util
RUN go get -u github.com/ngbook/micro-util/route \
    github.com/ngbook/micro-util/net

RUN go get -u github.com/ngbook/rest-api/fruit

RUN protoc-gen-micro --go_out=plugins=micro:. /go/src/github.com/ngbook/fruit/proto/fruit.api.proto

CMD ['ash']