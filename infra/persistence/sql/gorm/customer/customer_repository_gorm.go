package orm

import (
	"github.com/wagnerww/go-clean-arch-implement/domain/customer"
	"gorm.io/gorm"
)

type CustomerRepositoryGorm struct {
	db *gorm.DB
}

func NewCustomerRepositoryGorm(db *gorm.DB) *CustomerRepositoryGorm {
	return &CustomerRepositoryGorm{
		db: db,
	}
}

func (c *CustomerRepositoryGorm) Create(entity customer.Customer) error {
	model := DomainToModel(entity)
	c.db.Create(&model)

	return nil
}

func (c *CustomerRepositoryGorm) Update(entity customer.Customer) error {
	return nil
}

func (c *CustomerRepositoryGorm) FindById(id string) (entity customer.Customer, err error) {
	return entity, err
}

func (c *CustomerRepositoryGorm) FindAll() (entities []customer.Customer, err error) {
	return entities, err
}
