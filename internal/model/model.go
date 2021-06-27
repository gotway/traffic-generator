package model

import (
	"github.com/gotway/service-examples/pkg/catalog"
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
