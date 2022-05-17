package order

type Order struct {
	ID        string
	Customer  CustomerOrder
	Items     []OrderItem
	Disccount float32
	Total     float32
}

func NewOrder(id string, customerId string, customerName string, Items []OrderItem) (o Order, err error) {
	o = Order{
		ID: id,
	}

	o.ChangeCustomer(customerId, customerName)

	return o, err
}

func (o *Order) ChangeCustomer(customerId string, customerName string) {
	o.Customer, _ = NewCustomerOrder(customerId, customerName)
}
