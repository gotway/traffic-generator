package client

import (
	"net/http"

	internalHTTP "github.com/gotway/traffic-generator/internal/http"
)

func health(client *internalHTTP.Client, path string) (bool, error) {
	res, err := client.Get(path, nil)
	if err != nil {
		return false, err
	}
	return res.StatusCode == http.StatusOK, nil
}
