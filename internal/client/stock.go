package client

import (
	internalHTTP "github.com/gotway/traffic-generator/internal/http"
)

type Stock struct {
	httpClient *internalHTTP.Client
}

func (c *Stock) Health() (bool, error) {
	return health(c.httpClient, getStockPath("/health"))
}

func getStockPath(path string) string {
	return "/stock" + path
}

func NewStock(httpClient *internalHTTP.Client) *Stock {
	return &Stock{httpClient}
}
