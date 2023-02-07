// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type Comment struct {
	ID         string `json:"id"`
	Content    string `json:"content"`
	UserID     string `json:"userID"`
	ShopItemID string `json:"shopItemID"`
}

type NewComment struct {
	ID         string `json:"id"`
	Content    string `json:"content"`
	UserID     string `json:"userID"`
	ShopItemID string `json:"shopItemID"`
}

type NewShopItem struct {
	ID          string  `json:"id"`
	Title       string  `json:"title"`
	Description *string `json:"description"`
	Image       string  `json:"image"`
	Good        *int    `json:"good"`
	Prise       int     `json:"prise"`
}

type NewUser struct {
	ID         string  `json:"id"`
	Name       string  `json:"name"`
	Image      *string `json:"image"`
	Assessment *int    `json:"assessment"`
}

type User struct {
	ID         string      `json:"id"`
	Name       string      `json:"name"`
	Image      *string     `json:"image"`
	Assessment *int        `json:"assessment"`
	ShopItem   []*ShopItem `json:"ShopItem"`
}
