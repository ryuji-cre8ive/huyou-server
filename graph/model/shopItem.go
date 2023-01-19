package model

type ShopItem struct {
	ID          string     `json:"id"`
	Title       string     `json:"title"`
	Description *string    `json:"description"`
	Image       string     `json:"image"`
	Comments    []*Comment `json:"comments"`
	Seller      *User      `json:"seller"`
	Good        int        `json:"good"`
}