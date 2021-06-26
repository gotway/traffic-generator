package client

import (
	internalHTTP "github.com/gotway/traffic-generator/internal/http"
)

type Gotway struct {
	httpClient *internalHTTP.Client
}

func (c *Gotway) Health() (bool, error) {
	return health(c.httpClient, getGotwayPath("/health"))
}

func getGotwayPath(path string) string {
	return "/api" + path
}

func NewGotway(httpClient *internalHTTP.Client) *Gotway {
	return &Gotway{httpClient}
}
