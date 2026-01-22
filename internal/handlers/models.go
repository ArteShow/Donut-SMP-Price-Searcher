package handlers

type GetAveregPriceRequest struct {
	Token string `json:"token"`
}

type GetAveregPriceResponse struct {
	Price int64 `json:"price"`
}