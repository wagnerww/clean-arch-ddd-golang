package cacheredis

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/go-redis/redis/v8"
	"github.com/wagnerww/go-clean-arch-implement/domain/product"
)

var BASE_KEY string = "PRODUCT_"

type ProductRepositoryCache struct {
	rdb *redis.Client
	ctx context.Context
}

func NewProductRepositoryCache(ctx context.Context, rdb *redis.Client) *ProductRepositoryCache {
	return &ProductRepositoryCache{
		rdb: rdb,
		ctx: ctx,
	}
}

func (c *ProductRepositoryCache) Create(id string, entity product.Product) error {
	dto := DomainToDto(entity)
	ouputJson, _ := json.Marshal(dto)
	err := c.rdb.Set(c.ctx, BASE_KEY+id, string(ouputJson), 0).Err()
	if err != nil {
		println(err)
	}
	return nil
}

func (c *ProductRepositoryCache) Find(keyId string) (product.Product, error) {
	object, err := c.rdb.Get(c.ctx, BASE_KEY+keyId).Result()
	if err != nil {
		println(err)
	}
	var dto ProductCacheRedisDto
	err = json.Unmarshal([]byte(object), &dto)

	entity := DtoToDomain(dto)
	return entity, nil
}

func (c *ProductRepositoryCache) Delete(keyId string) error {
	fmt.Println("chegou aqui Delete" + BASE_KEY + keyId)
	return nil
}
