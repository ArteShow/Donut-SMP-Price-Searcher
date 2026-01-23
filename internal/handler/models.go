package handler

type GetLowestPricePerItemRequest struct {
	Token  string  `json:"token"`
	ID     string  `json:"id"`
	Amount float64 `json:"amount"`
}

type GetLowestPricePerItemResponse struct {
	Price float64 `json:"price"`
}