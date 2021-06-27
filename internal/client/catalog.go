package client

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gotway/service-examples/pkg/catalog"
	internalHTTP "github.com/gotway/traffic-generator/internal/http"
)

type Catalog struct {
	httpClient *internalHTTP.Client
}

func (c *Catalog) Health() (bool, error) {
	return health(c.httpClient, getCatalogPath("/health"))
}

func (c *Catalog) Create(product catalog.Product) (catalog.ProductCreated, error) {
	res, err := c.httpClient.Post(getCatalogPath("/product"), product)
	if err != nil {
		return catalog.ProductCreated{}, err
	}
	if res.StatusCode != http.StatusCreated {
		return catalog.ProductCreated{}, fmt.Errorf(
			"unable to create product, status code '%d'",
			res.StatusCode,
		)
	}
	var created catalog.ProductCreated
	if err := json.Unmarshal(res.Response, &created); err != nil {
		return catalog.ProductCreated{}, err
	}
	return created, nil
}

func (c *Catalog) Update(id int, product catalog.Product) error {
	statusCode, err := c.httpClient.Put(getCatalogPath("/product"), id, product)
	if err != nil {
		return err
	}
	if statusCode != http.StatusOK {
		return fmt.Errorf("unable to update product with id '%d', status code '%d'", id, statusCode)
	}
	return nil
}

func (c *Catalog) Delete(id int) error {
	statusCode, err := c.httpClient.Delete(getCatalogPath("/product"), id)
	if err != nil {
		return err
	}
	if statusCode != http.StatusOK {
		return fmt.Errorf("unable to delete product with id '%d', status code '%d'", id, statusCode)
	}
	return nil
}

func (c *Catalog) List(offset, limit int) (catalog.ProductPage, error) {
	res, err := c.httpClient.Get(getCatalogPath("/products"), map[string][]string{
		"offset": {strconv.Itoa(offset)},
		"limit":  {strconv.Itoa(limit)},
	})
	if err != nil {
		return catalog.ProductPage{}, err
	}
	var page catalog.ProductPage
	if err := json.Unmarshal(res.Response, &page); err != nil {
		return catalog.ProductPage{}, err
	}
	return page, nil
}

func getCatalogPath(path string) string {
	return "/catalog" + path
}

func NewCatalog(httpClient *internalHTTP.Client) *Catalog {
	return &Catalog{httpClient}
}
