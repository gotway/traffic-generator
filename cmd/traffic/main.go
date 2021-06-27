package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/gotway/gotway/pkg/log"
	"github.com/gotway/service-examples/pkg/route"
	"github.com/gotway/traffic-generator/internal/client"
	"github.com/gotway/traffic-generator/internal/config"
	"github.com/gotway/traffic-generator/internal/http"
	"github.com/gotway/traffic-generator/internal/worker"

	gs "github.com/gotway/gotway/pkg/graceful_shutdown"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	logger := log.NewLogger(log.Fields{
		"service": "traffic",
	}, config.Env, config.LogLevel, os.Stdout)
	httpClient := http.NewClient(http.GetURL(config.GotwayAddr, config.TLS), http.ClientOptions{
		Timeout: config.ClientTimeout,
	})

	gotwayClient := client.NewGotway(httpClient)
	healthy, err := gotwayClient.Health()
	if err != nil {
		logger.Fatalf("error connecting to gotway at '%s' %v", config.GotwayAddr, err)
	}
	if !healthy {
		logger.Fatalf("gotway at '%s' is not responding %v", config.GotwayAddr, err)
	}
	logger.Infof("gotway at '%s' is ready", config.GotwayAddr)

	catalogClient := client.NewCatalog(httpClient)
	healthy, err = catalogClient.Health()
	if err != nil {
		logger.Fatal("error connecting to catalog ", err)
	}
	if !healthy {
		logger.Fatal("catalog is not responding")
	}
	logger.Info("catalog is ready")

	stockClient := client.NewStock(httpClient)
	healthy, err = stockClient.Health()
	if err != nil {
		logger.Fatal("error connecting to stock ", err)
	}
	if !healthy {
		logger.Fatal("stock is not responding")
	}
	logger.Info("stock is ready")

	routeClient, err := route.NewClient(config.GotwayAddr, route.ClientOptions{
		Timeout: config.ClientTimeout,
		TLS: route.TLSOptions{
			Enabled:    config.TLS,
			CA:         config.TLSca,
			ServerHost: config.TLSserver,
		},
	})
	if err != nil {
		logger.Fatal("error connecting to route ", err)
	}
	logger.Info("route is ready")

	for i := 0; i < config.NumWorkers; i++ {
		if i > 0 {
			<-time.After(1 * time.Second)
		}
		name := fmt.Sprintf("worker-%d", i)
		w := worker.New(
			name,
			logger.WithFields(log.Fields{
				"type":     "worker",
				"instance": name,
			}),
			gotwayClient,
			catalogClient,
			stockClient,
			routeClient,
			worker.Options{
				NumClients:      config.NumClients,
				RequestInterval: config.RequestInterval,
			},
		)
		go w.Start(ctx)
	}

	gs.GracefulShutdown(logger, cancel, httpClient.Close, routeClient.Close)
}
