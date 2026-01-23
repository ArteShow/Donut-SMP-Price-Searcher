package models

type ListAuctionPageRequest struct {
	Sort string `json:"sort"`
}

type ListAuctionPageResponse struct {
	Response []Object `json:"result"`
}

type Object struct {
	Item  Item    `json:"item"`
	Price float64 `json:"price"`
}

type Item struct {
	DisplayName string   `json:"name"`
	Count       int      `json:"count"`
	Lore        []string `json:"lore"`
}