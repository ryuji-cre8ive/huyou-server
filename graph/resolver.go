package graph

//go:generate go run github.com/99designs/gqlgen generate

import (
	"github.com/ryuji-cre8ive/huyou-server/graph/model"
)
// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct{
	shopItems []*model.ShopItem
}
