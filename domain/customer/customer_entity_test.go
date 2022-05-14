package customer

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewCustomer(t *testing.T) {
	c, err := NewCustomer("123", "Wagner", 55787633, "wagnerricardonet@gmail.com")
	assert.Nil(t, err)
	assert.Equal(t, c.ID, "123")
	assert.Equal(t, c.Name, "Wagner")
	assert.Equal(t, c.PhoneNumber, 55787633)
	assert.Equal(t, c.Email, "wagnerricardonet@gmail.com")
	assert.Equal(t, c.Active, false)

}

func TestValidateActiveCustomer(t *testing.T) {
	c, err := NewCustomer("123", "Wagner", 55787633, "wagnerricardonet@gmail.com")
	c.Activate()
	assert.Nil(t, err)
	assert.Equal(t, c.ID, "123")
	assert.Equal(t, c.Name, "Wagner")
	assert.Equal(t, c.PhoneNumber, 55787633)
	assert.Equal(t, c.Email, "wagnerricardonet@gmail.com")
	assert.Equal(t, c.Active, true)

}

func TestValidateDeactiveCustomer(t *testing.T) {
	c, err := NewCustomer("123", "Wagner", 55787633, "wagnerricardonet@gmail.com")
	c.Activate()
	assert.Nil(t, err)
	assert.Equal(t, c.ID, "123")
	assert.Equal(t, c.Name, "Wagner")
	assert.Equal(t, c.PhoneNumber, 55787633)
	assert.Equal(t, c.Email, "wagnerricardonet@gmail.com")
	assert.Equal(t, c.Active, true)

	c.Deactivate()
	assert.Equal(t, c.Active, false)
}

func TestValidateIsActiveCustomer(t *testing.T) {
	c, err := NewCustomer("123", "Wagner", 55787633, "wagnerricardonet@gmail.com")
	c.Activate()
	assert.Nil(t, err)
	assert.Equal(t, c.ID, "123")
	assert.Equal(t, c.Name, "Wagner")
	assert.Equal(t, c.PhoneNumber, 55787633)
	assert.Equal(t, c.Email, "wagnerricardonet@gmail.com")
	assert.Equal(t, c.IsActive(), true)

}
