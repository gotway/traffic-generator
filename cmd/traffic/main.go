package main

import (
	"os"

	"github.com/gotway/gotway/pkg/log"
	"github.com/gotway/traffic-generator/internal/client"
	"github.com/gotway/traffic-generator/internal/config"
	"github.com/gotway/traffic-generator/internal/http"
)

func main() {
	logger := log.NewLogger(log.Fields{
		"service": "traffic",
	}, config.Env, config.LogLevel, os.Stdout)

	gotwayClient := client.NewGotwayClient(http.NewClient(config.GotwayURL, http.ClientOptions{
		Timeout: config.ClientTimeout,
	}))
	healthy, err := gotwayClient.Health()
	if err != nil {
		logger.Fatalf("error connecting to gotway at '%s' %v", config.GotwayURL, err)
	}
	if !healthy {
		logger.Fatalf("gotway at '%s' is not healthy %v", config.GotwayURL, err)
	}
	logger.Infof("connected to gotway at '%s'", config.GotwayURL)
}
