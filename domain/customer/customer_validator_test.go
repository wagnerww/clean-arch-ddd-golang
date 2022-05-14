package customer

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCustomerValidator(t *testing.T) {
	customer := Customer{
		ID: "1",
	}
	err := ValidateCustomer(customer)
	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), "Name is required")

}
