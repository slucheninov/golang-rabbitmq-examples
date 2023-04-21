FROM golang:alpine as builder

WORKDIR /app

COPY publisher.go publisher.go
COPY consumer/consumer.go consumer.go
COPY go.mod go.mod
COPY go.sum go.sum

RUN go get -d -v
RUN CGO_ENABLED=0 GOOS=linux go build -o publisher publisher.go
RUN CGO_ENABLED=0 GOOS=linux go build -o consumer consumer.go

FROM alpine:3.16

RUN apk update && apk add ca-certificates

COPY --from=builder /app/publisher publisher
COPY --from=builder /app/consumer consumer

ENV RABBITMQ_DSN ""

ENTRYPOINT [ "./consumer" ]
