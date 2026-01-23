package handlers

type GetAveregPriceRequest struct {
	Token string `json:"token"`
	Item  string `json:"item"`
}

type GetAveregPriceResponse struct {
	Price float64 `json:"price"`
}