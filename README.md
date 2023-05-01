# Golang RabbitMQ

### Installation

```
go mod download
```

### Running producer

```bash
kubectl exec -it .....
while true; do /usr/local/bin/publisher; done
```

```bash
export export RABBITMQ_DSN=amqp://guest:guest@10.16.1.3/
go run publisher.go
```

### Running consumer

```bash
export export RABBITMQ_DSN=amqp://guest:guest@10.16.1.3/
go run consumer/consumer.go
```

