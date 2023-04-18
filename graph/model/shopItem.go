package model

type ShopItem struct {
	ID                string     `json:"id"`
	Title             string     `json:"title"`
	Description       *string    `json:"description"`
	Image             *string    `json:"image"`
	Comments          []*Comment `json:"comments"`
	UserID            string     `json:"userID"`
	Good              int        `json:"good"`
	Price             int        `json:"price"`
	IsContainDelivery bool       `json:"isContainDelivery"`
}
