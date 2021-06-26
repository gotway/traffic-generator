package config

import (
	"time"

	c "github.com/gotway/gotway/pkg/config"
)

var (
	// Env indicates the environment name
	Env = c.GetEnv("ENV", "local")
	// LogLevel indicates the log level
	LogLevel = c.GetEnv("LOG_LEVEL", "debug")
	// GotwayURL specifies the Gotway instance URL
	GotwayURL = c.GetEnv("GOTWAY_URL", "https://localhost:11000")
	// ClientTimeout is the timeout in seconds for connecting to gotway
	ClientTimeout = time.Duration(
		c.GetIntEnv("CLIENT_TIMEOUT", 10),
	) * time.Second
	// NumWorkers is the number of goroutines used to generate traffic
	NumWorkers = c.GetIntEnv("NUM_WORKERS", 5)
	// NumClients is the number of concurrent clients to simulate
	NumClients = c.GetIntEnv("NUM_CLIENTS", 5)
	// RequestInterval defines the interval of requests in seconds
	Interval = time.Duration(
		c.GetIntEnv("REQUEST_INTERVAL", 10),
	) * time.Second
)
