# Golang RabbitMQ

### Installation

```
go mod download
```

### Running producer

```bash
export export RABBITMQ_DSN=amqp://guest:guest@10.16.1.3/
go run producer.go
```

### Running consumer

```bash
export export RABBITMQ_DSN=amqp://guest:guest@10.16.1.3/
go run consumer/consumer.go
```

