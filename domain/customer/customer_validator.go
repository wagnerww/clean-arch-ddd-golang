package customer

import (
	"errors"
)

func ValidateCustomer(entity Customer) error {
	if entity.Name == "" {
		return errors.New("Name is required")
	}
	return nil
}
