package cacheredis

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/go-redis/redis/v8"
	"github.com/wagnerww/go-clean-arch-implement/domain/customer"
)

var BASE_KEY string = "CUSTOMER_"

type CustomerRepositoryCache struct {
	rdb *redis.Client
	ctx context.Context
}

func NewCustomerRepositoryCache(ctx context.Context, rdb *redis.Client) *CustomerRepositoryCache {
	return &CustomerRepositoryCache{
		rdb: rdb,
		ctx: ctx,
	}
}

func (c *CustomerRepositoryCache) Create(id string, entity customer.Customer) error {
	ouput, _ := json.Marshal(entity)
	err := c.rdb.Set(c.ctx, BASE_KEY+id, string(ouput), 0).Err()
	if err != nil {
		println(err)
	}
	return nil
}

func (c *CustomerRepositoryCache) Find(keyId string) (customer.Customer, error) {
	fmt.Println("chegou aqui find CACHE" + BASE_KEY + keyId)
	return customer.Customer{}, nil
}

func (c *CustomerRepositoryCache) Delete(keyId string) error {
	fmt.Println("chegou aqui Delete" + BASE_KEY + keyId)
	return nil
}
