package model

type ShopItem struct {
	ID          string     `json:"id" gorm:"primaryKey"`
	Title       string     `json:"title"`
	Description *string    `json:"description" gorm:"type:text"`
	Image       string     `json:"image"`
	Comments    []*Comment `json:"comments" gorm:"foreignKey:ID"`
	Seller      *User      `json:"seller" gorm:"foreignKey:ID"`
	Good        int        `json:"good"`
}