package usecaseCustomerCreate

import (
	"log"

	"github.com/google/uuid"
	"github.com/wagnerww/go-clean-arch-implement/domain/customer"
	producer "github.com/wagnerww/go-clean-arch-implement/infra/queue/producer"
)

type Create struct {
	repoCache          customer.CustomerRepositoryCacheInterface
	repoDb             customer.CustomerRepositoryInterface
	eventBusEventStore producer.EventStoreProducerInterface
}

func NewCreate(
	rCache customer.CustomerRepositoryCacheInterface,
	repoDb customer.CustomerRepositoryInterface,
	eventBusEventStore producer.EventStoreProducerInterface,
) *Create {
	return &Create{
		repoCache:          rCache,
		repoDb:             repoDb,
		eventBusEventStore: eventBusEventStore,
	}
}

func (c *Create) Execute(input CustomerInputDto) (string, error) {
	id := uuid.New().String()
	cc, err := customer.NewCustomer(id, input.Name, input.PhoneNumber, input.Email)
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

	c.eventBusEventStore.Send(
		"customer",
		cc.ID,
		"CREATE",
		cc,
	)

	c.repoDb.Create(cc)
	c.repoCache.Create(cc.ID, cc)

	return cc.ID, nil

}
