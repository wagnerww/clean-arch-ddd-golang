package main

import (
	"time"

	messaging "github.com/wagnerww/go-clean-arch-implement/infra/queue/rabbitmq"
	eventStore "github.com/wagnerww/go-clean-arch-implement/infra/queue/rabbitmq/consumer/event-store"
)

func main() {
	time.Local, _ = time.LoadLocation("America/Sao_Paulo")

	rabbitMQChannel, rabbitMQConn := messaging.ConnectRabbitMQ()
	defer rabbitMQConn.Close()
	defer rabbitMQChannel.Close()

	sub := eventStore.NewEventStoreRabbitSubscribe(rabbitMQChannel)

	sub.Persist()

}
