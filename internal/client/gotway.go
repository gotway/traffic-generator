package client

import (
	"net/http"

	internalHTTP "github.com/gotway/traffic-generator/internal/http"
)

type GotwayClient struct {
	httpClient *internalHTTP.Client
}

func (c *GotwayClient) Health() (bool, error) {
	statusCode, err := c.httpClient.Get("/api/health", nil)
	if err != nil {
		return false, err
	}
	return statusCode == http.StatusOK, nil
}

func NewGotwayClient(httpClient *internalHTTP.Client) *GotwayClient {
	return &GotwayClient{httpClient}
}
