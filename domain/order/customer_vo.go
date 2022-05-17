package order

type CustomerOrder struct {
	ID   string
	Name string
}

func NewCustomerOrder(id string, name string) (c CustomerOrder, err error) {
	c = CustomerOrder{
		ID:   id,
		Name: name,
	}

	return c, nil
}
