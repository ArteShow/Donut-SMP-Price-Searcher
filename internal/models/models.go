package models

type ListAuctionPageRequest struct {
	Sort   string `json:"sort"`
}

type ListAuctionPageResponse struct {
	Response []Object `json:"result"`
}

type Object struct {
	Item  Item `json:"item"`
	Price int  `json:"price"`
}

type Item struct {
	DisplayName string   `json:"display_name"`
	Count       int      `json:"count"`
	Lore        []string `json:"lore"`
}