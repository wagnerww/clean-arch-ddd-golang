package usecaseProductCreate

type ProductInputDto struct {
	Name  string  `json:"name"`
	Price float32 `json:"price"`
}

type ProductOutputDto struct {
	ID    string  `json:"id"`
	Name  string  `json:"name"`
	Price float32 `json:"price"`
}
