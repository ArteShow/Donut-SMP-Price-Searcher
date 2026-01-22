package handlers

type GetAveregPriceRequest struct {
	Token string `json:"token"`
	Item  string `json:"item"`
}

type GetAveregPriceResponse struct {
	Price int64 `json:"price"`
}