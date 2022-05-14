package customer

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewAddress(t *testing.T) {
	a, err := NewAddress("Rua 1", "Cidade 1", 129)
	assert.Nil(t, err)
	assert.Equal(t, a.Street, "Rua 1")
	assert.Equal(t, a.City, "Cidade 1")
	assert.Equal(t, a.Number, 129)

}
