package orm

import (
	"github.com/google/uuid"
	"github.com/wagnerww/go-clean-arch-implement/domain/customer"
)

func DomainToModel(domain customer.Customer) (customer Customer) {

	model := Customer{
		ID:        uuid.MustParse(domain.ID),
		Name:      domain.Name,
		Email:     domain.Email,
		Active:    domain.IsActive(),
		Telephone: domain.PhoneNumber,
		City:      domain.Address.City,
		Street:    domain.Address.Street,
		Number:    domain.Address.Number,
	}

	return model

}
