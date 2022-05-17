package order

type OrderItem struct {
	ID       string
	Product  ProductOrder
	Price    float32
	Quantity int
	SubTotal float32
}

func NewOrderItem(id string, productId string, productName string, price float32, quantity int) (o OrderItem, err error) {
	o = OrderItem{
		ID:       id,
		Product:  ProductOrder{ID: productId, Name: productName},
		Price:    price,
		Quantity: quantity,
		SubTotal: 0,
	}

	return o, err
}
