package client

import (
	"net/http"

	internalHTTP "github.com/gotway/traffic-generator/internal/http"
)

func health(client *internalHTTP.Client, path string) (bool, error) {
	statusCode, err := client.Get(path, nil)
	if err != nil {
		return false, err
	}
	return statusCode == http.StatusOK, nil
}
