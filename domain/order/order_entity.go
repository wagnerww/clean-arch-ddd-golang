package order

type Order struct {
	ID           string
	CustomerId   string
	CustomerName string
	Items        []OrderItem
	Disccount    float32
	Total        float32
}

func NewOrder(id string, customerId string, customerName string, Items []OrderItem) (o Order, err error) {
	o = Order{
		ID:           id,
		CustomerId:   customerId,
		CustomerName: customerName,
	}

	return o, err
}
