package product

type Product struct {
	ID    string
	Name  string
	Price float32
}

func NewProduct(id string, name string, price float32) (c Product, err error) {
	p := Product{
		ID:    id,
		Name:  name,
		Price: price,
	}

	return p, err
}
