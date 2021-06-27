package client

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gotway/service-examples/pkg/stock"
	internalHTTP "github.com/gotway/traffic-generator/internal/http"
)

type Stock struct {
	httpClient *internalHTTP.Client
}

func (c *Stock) Health() (bool, error) {
	return health(c.httpClient, getStockPath("/health"))
}

func (c *Stock) List(productIds ...int) (*stock.StockList, error) {
	var query map[string][]string
	if len(productIds) > 0 {
		productIdsString := make([]string, len(productIds))
		for i, id := range productIds {
			productIdsString[i] = strconv.Itoa(id)
		}
		query = map[string][]string{
			"productId": productIdsString,
		}
	}
	res, err := c.httpClient.Get(getStockPath("/list"), query)
	if err != nil {
		return &stock.StockList{}, err
	}
	if res.StatusCode != http.StatusOK {
		return &stock.StockList{}, fmt.Errorf(
			"unable to get stock , status code '%d'",
			res.StatusCode,
		)
	}
	var stockList stock.StockList
	if err := json.Unmarshal(res.Response, &stockList); err != nil {
		return &stock.StockList{}, err
	}
	return &stockList, nil
}

func (c *Stock) Upsert(stockList *stock.StockList) (int, error) {
	res, err := c.httpClient.Post(getStockPath("/list"), stockList)
	if err != nil {
		return 0, err
	}
	if res.StatusCode != http.StatusCreated {
		return 0, fmt.Errorf("error upserting stock, status code '%d'", res.StatusCode)
	}
	return res.StatusCode, nil
}

func getStockPath(path string) string {
	return "/stock" + path
}

func NewStock(httpClient *internalHTTP.Client) *Stock {
	return &Stock{httpClient}
}
