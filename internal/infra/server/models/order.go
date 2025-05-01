package models

type OrderRequest struct {
	BuyerId string `json:"buyer_id"`
	ItemId  string `json:"item_id"`
	Qty     string `json:"qty"`
	Address string `json:"address"`
}
