package cacheredis

type ProductCacheRedisDto struct {
	ID    string  `json:"id"`
	Name  string  `json:"name"`
	Price float32 `json:"price"`
}
