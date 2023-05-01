FROM --platform=linux/amd64 golang:1.20.3-alpine3.16 as builder
WORKDIR /app

COPY publisher.go publisher.go
COPY consumer/consumer.go consumer.go
COPY go.mod go.mod
COPY go.sum go.sum

RUN go get -d -v

ARG BUILD_OPTS='-i' 
RUN go build ${BUILD_OPT} -o publisher publisher.go
RUN go build ${BUILD_OPT} -o consumer consumer.go

FROM --platform=linux/amd64 alpine:3.16

RUN apk update && apk add ca-certificates && apk add bash

COPY --from=builder /app/publisher /usr/local/bin/publisher
COPY --from=builder /app/consumer /usr/local/bin/consumer

