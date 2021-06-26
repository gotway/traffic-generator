package client

import (
	internalHTTP "github.com/gotway/traffic-generator/internal/http"
)

type Catalog struct {
	httpClient *internalHTTP.Client
}

func (c *Catalog) Health() (bool, error) {
	return health(c.httpClient, getCatalogPath("/health"))
}

func getCatalogPath(path string) string {
	return "/catalog" + path
}

func NewCatalog(httpClient *internalHTTP.Client) *Catalog {
	return &Catalog{httpClient}
}
