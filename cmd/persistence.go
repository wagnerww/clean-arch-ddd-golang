package main

import (
	"context"
	"time"

	redisCache "github.com/wagnerww/go-clean-arch-implement/infra/persistence/caching/redis"
	customerRedisCache "github.com/wagnerww/go-clean-arch-implement/infra/persistence/caching/redis/customer"
	productRedisCache "github.com/wagnerww/go-clean-arch-implement/infra/persistence/caching/redis/product"
	gormDb "github.com/wagnerww/go-clean-arch-implement/infra/persistence/sql/gorm"
	customerDBGorm "github.com/wagnerww/go-clean-arch-implement/infra/persistence/sql/gorm/customer"
	eventStoreDBGorm "github.com/wagnerww/go-clean-arch-implement/infra/persistence/sql/gorm/events"
	messaging "github.com/wagnerww/go-clean-arch-implement/infra/queue/rabbitmq"
	eventStorePublish "github.com/wagnerww/go-clean-arch-implement/infra/queue/rabbitmq/producer/event-store"
	usecaseCustomer "github.com/wagnerww/go-clean-arch-implement/usecase/customer/create"
	usecaseProduct "github.com/wagnerww/go-clean-arch-implement/usecase/product/create"
)

var ctx = context.Background()

func main() {
	time.Local, _ = time.LoadLocation("America/Sao_Paulo")

	rabbitMQChannel, rabbitMQConn, err := messaging.ConnectRabbitMQ()
	redisConn := redisCache.ConnectRedis(ctx)
	dbConn := gormDb.ConnectDB()

	if err == nil {
		defer rabbitMQConn.Close()
		defer rabbitMQChannel.Close()
	}

	repoCacheProduct := productRedisCache.NewProductRepositoryCache(ctx, redisConn)
	createProduct := usecaseProduct.NewCreate(repoCacheProduct)
	createProduct.Execute(usecaseProduct.ProductInputDto{"Produto 1", 10.00})

	repoCacheCustomer := customerRedisCache.NewCustomerRepositoryCache(ctx, redisConn)
	repoDbCustomer := customerDBGorm.NewCustomerRepositoryGorm(dbConn)

	repoEvents := eventStoreDBGorm.NewEventStoreRepositoryGorm(dbConn)
	eventBusEventStore := eventStorePublish.NewEventStoreRabbit(rabbitMQChannel, repoEvents)

	create := usecaseCustomer.NewCreate(repoCacheCustomer, repoDbCustomer, eventBusEventStore)
	create.Execute(usecaseCustomer.CustomerInputDto{"Wagner", 123, true, "wagnerricardonet@gmail.com"})

}
