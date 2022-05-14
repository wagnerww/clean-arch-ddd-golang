package order

type OrderItem struct {
	ID          string
	ProductId   string
	ProductName string
	Price       float32
	Quantity    int
	SubTotal    float32
}

func NewOrderItem(id string, productId string, productName string, price float32, quantity int) (o OrderItem, err error) {
	o = OrderItem{
		ID:          id,
		ProductId:   productId,
		ProductName: productName,
		Price:       price,
		Quantity:    quantity,
	}

	return o, err
}
