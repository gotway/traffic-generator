package http

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"io"
	"net/http"
	"time"
)

type ClientOptions struct {
	Timeout time.Duration
}

type Client struct {
	url        string
	httpClient http.Client
}

func (c *Client) Get(path string, query map[string]string) (int, error) {
	req, err := http.NewRequest(http.MethodGet, c.getURL(path), nil)
	if err != nil {
		return 0, err
	}

	q := req.URL.Query()
	for key, value := range query {
		q.Add(key, value)
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return 0, err
	}
	return res.StatusCode, nil
}

func (c *Client) Post(path string, body interface{}) (int, error) {
	reader, err := c.serialize(body)
	if err != nil {
		return 0, err
	}

	req, err := http.NewRequest(http.MethodPost, c.getURL(path), reader)
	if err != nil {
		return 0, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return 0, err
	}
	return res.StatusCode, nil
}

func (c *Client) Close() {
	c.httpClient.CloseIdleConnections()
}

func (c *Client) getURL(path string) string {
	return c.url + path
}

func (c *Client) serialize(obj interface{}) (io.Reader, error) {
	b, err := json.Marshal(obj)
	if err != nil {
		return nil, err
	}
	return bytes.NewReader(b), nil
}

func NewClient(url string, options ClientOptions) *Client {
	return &Client{url, http.Client{
		Timeout: options.Timeout,
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
	}}
}

func GetURL(address string, tlsEnabled bool) string {
	if tlsEnabled {
		return "https://" + address
	}
	return "http://" + address
}
