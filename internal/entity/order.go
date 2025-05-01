package entity

import "github.com/google/uuid"

type Order struct {
	Id           string
	FoodId       string
	BuyerName    string
	BuyerPhone   string
	BuyerEmail   string
	BuyerAddress string
	Qty          string
}

type OrderResponse struct {
	OrderId      string `json:"order_id"`
	BuyerName    string `json:"buyer_name"`
	BuyerPhone   string `json:"buyer_phone"`
	BuyerEmail   string `json:"buyer_email"`
	BuyerAddress string `json:"buyer_address"`
	ItemName     string `json:"item_name"`
	ItemUnit     string `json:"item_unit"`
	ItemPrice    string `json:"item_price"`
	TotalQty     string `json:"total_qty"`
	TotalPrice   string `json:"total_price"`
}

func NewOrder(
	foodId,
	buyerName,
	buyerPhone,
	buyerEmail,
	buyerAddress,
	qty string,
) *Order {
	return &Order{
		Id:           uuid.New().String(),
		FoodId:       foodId,
		BuyerName:    buyerName,
		BuyerPhone:   buyerPhone,
		BuyerEmail:   buyerEmail,
		BuyerAddress: buyerAddress,
		Qty:          qty,
	}
}
