package model

import (
	"github.com/gotway/service-examples/pkg/catalog"
	"github.com/gotway/service-examples/pkg/stock"
	"github.com/gotway/traffic-generator/internal/rand"
)

func NewRandomProduct() catalog.Product {
	return catalog.Product{
		Name:  rand.Item("sneakers", "socks", "t-shirt", "trousers", "coat"),
		Price: rand.Int(10, 500),
		Color: rand.Item("white", "blue", "grey", "green", "black"),
		Size:  rand.Item("XS", "S", "M", "L", "XL"),
	}
}

func NewRandomStockList(productIds ...int) *stock.StockList {
	stocks := make([]stock.Stock, len(productIds))
	for i, id := range productIds {
		stocks[i] = stock.Stock{
			ProductID: id,
			Units:     rand.Int(1, 100),
			TTL:       rand.Int(1, 100),
		}
	}
	return &stock.StockList{Stock: stocks}
}
