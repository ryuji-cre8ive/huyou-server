// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type Comment struct {
	ID         string `json:"id"`
	Content    string `json:"content"`
	UserID     string `json:"userID"`
	ShopItemID string `json:"shopItemID"`
}

type Follower struct {
	ID           string  `json:"id"`
	UserID       string  `json:"userID"`
	TargetUserID string  `json:"targetUserID"`
	TargetUser   []*User `json:"targetUser" gorm:"many2many:follows;foreignKey:TargetUserID"`
}

type NewComment struct {
	ID         string `json:"id"`
	Content    string `json:"content"`
	UserID     string `json:"userID"`
	ShopItemID string `json:"shopItemID"`
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
	Mail       string      `json:"mail"`
	Password   string      `json:"password"`
	Image      *string     `json:"image"`
	Assessment *int        `json:"assessment"`
	ShopItem   []*ShopItem `json:"ShopItem"`
	Followers  []*Follower `json:"Followers"`
}
