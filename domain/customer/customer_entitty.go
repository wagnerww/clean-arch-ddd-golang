package customer

type Customer struct {
	ID          string
	Name        string
	PhoneNumber int
	Active      bool
	Email       string
	Address     Address
}

func NewCustomer(id string, name string, phoneNumber int, email string) (c Customer, err error) {
	c = Customer{
		ID:          id,
		Name:        name,
		PhoneNumber: phoneNumber,
		Email:       email,
		Active:      false,
	}

	err = c.Validate()
	if err != nil {
		return c, err
	}

	return c, nil
}

func (c *Customer) Validate() error {
	return ValidateCustomer(*c)
}

func (c *Customer) Activate() {
	c.Active = true
}

func (c *Customer) Deactivate() {
	c.Active = false
}

func (c *Customer) IsActive() bool {
	return c.Active
}

func (c *Customer) ChangeAdderss(street string, city string, number int) error {
	address, error := NewAddress(street, city, number)

	if error != nil {
		return error
	}

	c.Address = address

	return nil
}
