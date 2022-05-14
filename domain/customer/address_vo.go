package customer

type Address struct {
	Street string
	City   string
	Number int
}

func NewAddress(street string, city string, number int) (a Address, err error) {
	a = Address{
		Street: street,
		City:   city,
		Number: number,
	}

	return a, nil
}
