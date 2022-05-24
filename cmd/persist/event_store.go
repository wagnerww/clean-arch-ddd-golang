package main

import (
	"time"

	gormDb "github.com/wagnerww/go-clean-arch-implement/infra/persistence/sql/gorm"
	eventStoreDBGorm "github.com/wagnerww/go-clean-arch-implement/infra/persistence/sql/gorm/events"
	messaging "github.com/wagnerww/go-clean-arch-implement/infra/queue/rabbitmq"
	eventStore "github.com/wagnerww/go-clean-arch-implement/infra/queue/rabbitmq/consumer/event-store"
)

func main() {
	time.Local, _ = time.LoadLocation("America/Sao_Paulo")

	rabbitMQChannel, rabbitMQConn := messaging.ConnectRabbitMQ()
	defer rabbitMQConn.Close()
	defer rabbitMQChannel.Close()
	dbConn := gormDb.ConnectDB()

	repoDB := eventStoreDBGorm.NewEventStoreRepositoryGorm(dbConn)
	sub := eventStore.NewEventStoreRabbitSubscribe(rabbitMQChannel, repoDB)

	sub.Persist()

}
