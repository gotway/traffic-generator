package http

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
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

type ClientResponse struct {
	StatusCode int
	Response   []byte
}

func (c *Client) Get(path string, query map[string]string) (*ClientResponse, error) {
	req, err := http.NewRequest(http.MethodGet, c.getURL(path), nil)
	if err != nil {
		return nil, err
	}
	if len(query) > 0 {
		q := req.URL.Query()
		for key, value := range query {
			q.Add(key, value)
		}
		req.URL.RawQuery = q.Encode()
	}
	res, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	return getClientResponse(res)
}

func (c *Client) Post(path string, body interface{}) (*ClientResponse, error) {
	reader, err := getReader(body)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest(http.MethodPost, c.getURL(path), reader)
	if err != nil {
		return nil, err
	}
	res, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	return getClientResponse(res)
}

func (c *Client) Put(path string, id int, body interface{}) (int, error) {
	reader, err := getReader(body)
	if err != nil {
		return 0, err
	}
	req, err := http.NewRequest(http.MethodPut, c.getResourceURL(path, id), reader)
	if err != nil {
		return 0, err
	}
	res, err := c.httpClient.Do(req)
	if err != nil {
		return 0, err
	}
	return res.StatusCode, nil
}

func (c *Client) Delete(path string, id int) (int, error) {
	req, err := http.NewRequest(http.MethodDelete, c.getResourceURL(path, id), nil)
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

func (c *Client) getResourceURL(path string, id int) string {
	return fmt.Sprintf("%s/%d", c.getURL(path), id)
}

func getReader(obj interface{}) (io.Reader, error) {
	b, err := json.Marshal(obj)
	if err != nil {
		return nil, err
	}
	return bytes.NewReader(b), nil
}

func getClientResponse(res *http.Response) (*ClientResponse, error) {
	bytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	return &ClientResponse{res.StatusCode, bytes}, nil
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
