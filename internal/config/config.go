package config

import (
	"time"

	c "github.com/gotway/gotway/pkg/config"
	"github.com/gotway/gotway/pkg/tlstest"
)

var (
	// Env indicates the environment name
	Env = c.GetEnv("ENV", "local")
	// LogLevel indicates the log level
	LogLevel = c.GetEnv("LOG_LEVEL", "debug")
	// GotwayHost specifies the Gotway host
	GotwayHost = c.GetEnv("GOTWAY_HOST", "gotway:11000")
	// CatalogHost specifies the Catalog host
	CatalogHost = c.GetEnv("CATALOG_HOST", "catalog:11000")
	// StockHost specifies the Stock host
	StockHost = c.GetEnv("STOCK_HOST", "stock:11000")
	// ClientTimeout is the timeout in seconds for connecting to gotway
	ClientTimeout = time.Duration(
		c.GetIntEnv("CLIENT_TIMEOUT", 10),
	) * time.Second
	// NumWorkers is the number of goroutines used to generate traffic
	NumWorkers = c.GetIntEnv("NUM_WORKERS", 5)
	// NumClients is the number of concurrent clients to simulate
	NumClients = c.GetIntEnv("NUM_CLIENTS", 2)
	// RequestInterval defines the interval of requests in seconds
	RequestInterval = time.Duration(
		c.GetIntEnv("REQUEST_INTERVAL", 10),
	) * time.Second
	// TLS indicates if TLS is enabled
	TLS = c.GetBoolEnv("TLS", true)
	// TLSca is the TLS certificate authority
	TLSca = c.GetEnv("TLS_CA", tlstest.CA())
	// TLSserver is the TLS validation server
	TLSserver = c.GetEnv("TLS_SERVER", tlstest.Server())
)
