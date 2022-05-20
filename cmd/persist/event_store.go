package main

import (
	"time"

	messaging "github.com/wagnerww/go-clean-arch-implement/infra/queue/rabbitmq"
	"github.com/wagnerww/go-clean-arch-implement/infra/queue/rabbitmq/subscribe"
)

func main() {
	time.Local, _ = time.LoadLocation("America/Sao_Paulo")

	rabbitMQChannel, rabbitMQConn := messaging.ConnectRabbitMQ()
	defer rabbitMQConn.Close()
	defer rabbitMQChannel.Close()

	sub := subscribe.NewEventStoreRabbit(rabbitMQChannel)

	sub.Persist()

}
