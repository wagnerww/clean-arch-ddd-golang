package usecasecustomer

type CustomerInputDto struct {
	Name        string `json:"name"`
	PhoneNumber int    `json:"phoneNumber"`
	Active      bool   `json:"active"`
	Email       string `json:"email"`
}

type CustomerOutputDto struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	PhoneNumber int    `json:"phoneNumber"`
	Active      bool   `json:"active"`
	Email       string `json:"email"`
}
