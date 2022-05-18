package order

type ProductOrder struct {
	ID   string
	Name string
}

func NewProductOrder(id string, name string) (p ProductOrder) {
	p = ProductOrder{
		ID:   id,
		Name: name,
	}

	return p
}
