package main

import (
	"context"

	redisCache "github.com/wagnerww/go-clean-arch-implement/infra/persistence/caching/redis"
	customerRedisCache "github.com/wagnerww/go-clean-arch-implement/infra/persistence/caching/redis/customer"
	productRedisCache "github.com/wagnerww/go-clean-arch-implement/infra/persistence/caching/redis/product"
	usecaseCustomer "github.com/wagnerww/go-clean-arch-implement/usecase/customer/create"
	usecaseProduct "github.com/wagnerww/go-clean-arch-implement/usecase/product/create"
)

var ctx = context.Background()

func main() {
	redisConn := redisCache.ConnectRedis(ctx)

	repoCacheProduct := productRedisCache.NewProductRepositoryCache(ctx, redisConn)
	createProduct := usecaseProduct.NewCreate(repoCacheProduct)
	createProduct.Execute(usecaseProduct.ProductInputDto{"Produto 1", 10.00})

	repoCacheCustomer := customerRedisCache.NewCustomerRepositoryCache(ctx, redisConn)
	create := usecaseCustomer.NewCreate(repoCacheCustomer)
	create.Execute(usecaseCustomer.CustomerInputDto{"Wagner", 123, true, "wagnerricardonet@gmail.com"})

}
