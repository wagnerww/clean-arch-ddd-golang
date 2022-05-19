package usecaseProductCreate

import (
	"log"

	"github.com/wagnerww/go-clean-arch-implement/domain/product"
)

type Create struct {
	repoCache product.ProductRepositoryCacheInterface
}

func NewCreate(
	rCache product.ProductRepositoryCacheInterface) *Create {
	return &Create{
		repoCache: rCache,
	}
}

func (c *Create) Execute(input ProductInputDto) (string, error) {
	cc, err := product.NewProduct("123", input.Name, input.Price)

	if err != nil {
		log.Println("error", err)
		return cc.ID, err
	}

	//err = c.repoSQL.Create(cc)

	if err != nil {
		return cc.ID, err
	}
	log.Println("vai gravar")

	c.repoCache.Create(cc.ID, cc)

	return cc.ID, nil

}
