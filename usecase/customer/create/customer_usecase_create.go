package usecaseCustomerCreate

import (
	"log"

	"github.com/wagnerww/go-clean-arch-implement/domain/customer"
)

type Create struct {
	repoCache customer.CustomerRepositoryCacheInterface
}

func NewCreate(
	rCache customer.CustomerRepositoryCacheInterface) *Create {
	return &Create{
		repoCache: rCache,
	}
}

func (c *Create) Execute(input CustomerInputDto) (string, error) {
	cc, err := customer.NewCustomer("123", input.Name, input.PhoneNumber, input.Email)
	cc.Activate()

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
