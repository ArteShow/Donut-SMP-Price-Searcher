package models

type GetAuctionStatsRequest struct {
	Sort   string `json:"sort"`
	Search string `json:"searcher"`
}

type GetAuctionStatsResponse struct {
	Items []AhObject `json:"result"`
}

type AhObject struct {
	Price float64 `json:"price"`
	Item  AhItem  `json:"item"`
}

type AhItem struct {
	ID    string  `json:"id"`
	Count float64 `json:"count"`
}