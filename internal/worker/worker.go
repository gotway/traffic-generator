package worker

import (
	"context"
	"time"

	"github.com/gotway/gotway/pkg/log"
	"github.com/gotway/service-examples/pkg/catalog"
	"github.com/gotway/service-examples/pkg/route"
	"github.com/gotway/traffic-generator/internal/client"
	"github.com/gotway/traffic-generator/internal/model"
	"github.com/gotway/traffic-generator/internal/rand"
)

type Options struct {
	NumClients      int
	RequestInterval time.Duration
}

type Worker struct {
	name          string
	logger        log.Logger
	gotwayClient  *client.Gotway
	catalogClient *client.Catalog
	stockClient   *client.Stock
	routeClient   *route.Client
	options       Options
}

func (w *Worker) Start(ctx context.Context) {
	w.logger.Infof("starting worker '%s'", w.name)
	ticker := time.NewTicker(w.options.RequestInterval)
	for {
		select {
		case <-ctx.Done():
			w.logger.Infof("stopping worker '%s'", w.name)
			return
		case <-ticker.C:
			w.simulateTraffic(ctx)
		}
	}
}

func (w *Worker) simulateTraffic(ctx context.Context) {
	w.logger.Infof("simulating traffic of %d clients", w.options.NumClients)
	for i := 0; i < w.options.NumClients; i++ {
		go w.singleClientTraffic(ctx)
	}
}

func (w *Worker) singleClientTraffic(ctx context.Context) {
	go func() {
		products := []catalog.Product{
			model.NewRandomProduct(),
			model.NewRandomProduct(),
			model.NewRandomProduct(),
			model.NewRandomProduct(),
		}
		created := make([]int, len(products))
		for i, p := range products {
			c, err := w.catalogClient.Create(p)
			if err != nil {
				w.logger.Error("error creating product ", err)
			}
			created[i] = c.ID
			defer w.catalogClient.Delete(c.ID)
			if rand.Bool() {
				if err := w.catalogClient.Update(c.ID, model.NewRandomProduct()); err != nil {
					w.logger.Error("error updating product ", err)
				}
			}
		}
		_, err := w.catalogClient.List(0, rand.Int(len(products), 100))
		if err != nil {
			w.logger.Error("error listing product ", err)
		}

		if _, err := w.stockClient.Upsert(model.NewRandomStockList(created...)); err != nil {
			w.logger.Error("error upserting strock ", err)
		}
		if _, err := w.stockClient.List(created...); err != nil {
			w.logger.Error("error getting stock ", err)
		}
	}()

	go func() {
		if _, err := w.routeClient.GetFeature(ctx, route.ValidPoint); err != nil {
			w.logger.Error("get feature failed ", err)
		}
		if _, err := w.routeClient.ListFeatures(ctx, route.Rect); err != nil {
			w.logger.Error("list features failed ", err)
		}
	}()
}

func New(
	name string,
	logger log.Logger,
	gotwayClient *client.Gotway,
	catalogClient *client.Catalog,
	stockClient *client.Stock,
	routeClient *route.Client,
	options Options,
) *Worker {
	return &Worker{name, logger, gotwayClient, catalogClient, stockClient, routeClient, options}
}
