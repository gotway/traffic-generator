package worker

import (
	"context"
	"sync"
	"time"

	"github.com/gotway/gotway/pkg/log"
	"github.com/gotway/service-examples/pkg/route"
	"github.com/gotway/traffic-generator/internal/client"
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
	var wg sync.WaitGroup
	wg.Add(w.options.NumClients)
	for i := 0; i < w.options.NumClients; i++ {
		go func() {
			defer wg.Done()
			w.singleClientTraffic(ctx)
		}()
	}
	wg.Wait()
}

func (w *Worker) singleClientTraffic(ctx context.Context) {
	// TODO
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
