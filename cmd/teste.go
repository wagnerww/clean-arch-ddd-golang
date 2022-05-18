package main

import (
	"context"

	redisCache "github.com/wagnerww/go-clean-arch-implement/infra/persistence/cache/redis"
	customerRedisCache "github.com/wagnerww/go-clean-arch-implement/infra/persistence/cache/redis/customer"
	usecaseCustomer "github.com/wagnerww/go-clean-arch-implement/usecase/customer"
)

var ctx = context.Background()

func main() {
	redisConn := redisCache.ConnectRedis(ctx)
	repoCache := customerRedisCache.NewCustomerRepositoryCache(ctx, redisConn)

	create := usecaseCustomer.NewCreate(repoCache)
	create.Execute(usecaseCustomer.CustomerInputDto{"Wagner", 123, true, "wagnerricardonet@gmail.com"})

}
