package product

type ProductRequest struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Category string `json:"category"`
	Price    int64  `json:"price"`
}

type ProductResponse struct {
	ID string `json:"id"`
}
