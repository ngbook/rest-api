FROM jsongo/go-alpine-builder:latest as builder
RUN go get -u github.com/ngbook/rest-api/fruit


FROM microhq/micro:kubernetes
LABEL author="jsongo<jsongo@qq.com>"
COPY --from=builder /go/bin/fruit /
CMD [ "fruit" ]
