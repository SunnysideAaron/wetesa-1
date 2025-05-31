// Package config provides configuration management for the API service.
// [The Twelve-Factor App III. Config](https://12factor.net/config)
package config

import (
	"os"
	"strconv"
	"time"
)

const (
	// EnvironmentDev represents the development environment configuration.
	EnvironmentDev string = "dev"
	// EnvironmentProd represents the production environment configuration.
	EnvironmentProd string = "prod"
)

// APIConfig stores the API service configuration parameters.
type APIConfig struct {
	// BaseURL is the base URL for the API. Includes API version. Unlike the
	// other settings do not load BaseURL from an environment variable. Including
	// here so there is only one source of truth for api version.
	BaseURL string
	// Environment specifies the running environment (dev/prod)
	Environment string
	// APIHost is the host address to bind the server to
	APIHost string
	// APIPort is the port number to listen on
	APIPort string
	// APIReadTimeout is the maximum duration for reading the entire request
	APIReadTimeout time.Duration
	// APIWriteTimeout is the maximum duration before timing out writes of the response
	APIWriteTimeout time.Duration
	// APIDefaultWriteTimeout is the default timeout for write operations
	APIDefaultWriteTimeout time.Duration
	// APIIdleTimeout is the maximum amount of time to wait for the next request
	APIIdleTimeout time.Duration
	// RequestMaxBytes is the maximum size of incoming request bodies
	RequestMaxBytes int64
}

// LoadAPIConfig loads and validates API configuration from environment variables.
// It applies default values when environment variables are not set.
func LoadAPIConfig() *APIConfig {
	// Set default values.
	cnf := &APIConfig{
		BaseURL:                "/api/v0.1",
		Environment:            EnvironmentProd,
		APIHost:                "",
		APIPort:                "8080",
		APIReadTimeout:         15 * time.Second,
		APIWriteTimeout:        60 * time.Second,
		APIDefaultWriteTimeout: 30 * time.Second,
		APIIdleTimeout:         60 * time.Second,
		// common values default is 8k, Other defaults might be 4k, 16k, or 48k.
		RequestMaxBytes: 8192,
	}

	// Read and validate environment variables.
	env := os.Getenv("ENVIRONMENT")
	if env == "dev" {
		cnf.Environment = env
	}

	cnf.APIHost = os.Getenv("API_HOST")

	var err error
	var parsedDuration time.Duration
	portStr := os.Getenv("API_PORT")
	_, err = strconv.Atoi(portStr) // Validate that port is #
	if err == nil {
		cnf.APIPort = portStr
	} // else if there is an error this will use the default value.

	readTimeoutStr := os.Getenv("API_READ_TIMEOUT")
	if readTimeoutStr != "" {
		if parsedDuration, err = time.ParseDuration(readTimeoutStr); err == nil {
			cnf.APIReadTimeout = parsedDuration
		}
	}

	writeTimeoutStr := os.Getenv("API_WRITE_TIMEOUT")
	if writeTimeoutStr != "" {
		if parsedDuration, err = time.ParseDuration(writeTimeoutStr); err == nil {
			cnf.APIWriteTimeout = parsedDuration
		}
	}

	timeoutStr := os.Getenv("API_DEFAULT_WRITE_TIMEOUT")
	if timeoutStr != "" {
		if parsedDuration, err = time.ParseDuration(timeoutStr); err == nil {
			cnf.APIDefaultWriteTimeout = parsedDuration
		}
	}

	if cnf.APIDefaultWriteTimeout > cnf.APIWriteTimeout {
		cnf.APIDefaultWriteTimeout = cnf.APIWriteTimeout
	}

	idleTimeoutStr := os.Getenv("API_IDLE_TIMEOUT")
	if idleTimeoutStr != "" {
		if parsedDuration, err = time.ParseDuration(idleTimeoutStr); err == nil {
			cnf.APIIdleTimeout = parsedDuration
		}
	}

	maxBytesStr := os.Getenv("API_REQUEST_MAX_BYTES")
	maxBytes, err := strconv.ParseInt(maxBytesStr, 10, 64)
	if err == nil && maxBytes > 0 {
		cnf.RequestMaxBytes = maxBytes
	} // else if there is an error this will use the default value.

	return cnf
}

// WebConfig stores the Web service configuration parameters.
type WebConfig struct {
	// BaseURL is the base URL for the API. Includes API version. Unlike the
	// other settings do not load BaseURL from an environment variable. Including
	// here so there is only one source of truth for api version.
	// BaseURL string
	// Environment specifies the running environment (dev/prod)
	Environment string
	// APIHost is the host address to bind the server to
	WebHost string
	// APIPort is the port number to listen on
	WebPort string
	// APIReadTimeout is the maximum duration for reading the entire request
	WebReadTimeout time.Duration
	// APIWriteTimeout is the maximum duration before timing out writes of the response
	WebWriteTimeout time.Duration
	// APIDefaultWriteTimeout is the default timeout for write operations
	WebDefaultWriteTimeout time.Duration
	// APIIdleTimeout is the maximum amount of time to wait for the next request
	WebIdleTimeout time.Duration
	// RequestMaxBytes is the maximum size of incoming request bodies
	// RequestMaxBytes int64
	WebCacheTemplates bool
}

// TODO api config and web config are virtually the same

// LoadAPIConfig loads and validates API configuration from environment variables.
// It applies default values when environment variables are not set.
func LoadWebConfig() *WebConfig {
	// Set default values.
	cnf := &WebConfig{
		// BaseURL:                "/api/v0.1",
		Environment:     EnvironmentProd,
		WebHost:         "",
		WebPort:         "8082",
		WebReadTimeout:  15 * time.Second,
		WebWriteTimeout: 60 * time.Second,
		// WebDefaultWriteTimeout: 30 * time.Second,
		WebIdleTimeout: 60 * time.Second,
		// common values default is 8k, Other defaults might be 4k, 16k, or 48k.
		// RequestMaxBytes: 8192,
		WebCacheTemplates: true,
	}

	// Read and validate environment variables.
	env := os.Getenv("ENVIRONMENT")
	if env == "dev" {
		cnf.Environment = env
	}

	cnf.WebHost = os.Getenv("WEB_HOST")

	var err error
	var parsedDuration time.Duration
	portStr := os.Getenv("WEB_PORT")
	_, err = strconv.Atoi(portStr) // Validate that port is #
	if err == nil {
		cnf.WebPort = portStr
	} // else if there is an error this will use the default value.

	readTimeoutStr := os.Getenv("WEB_READ_TIMEOUT")
	if readTimeoutStr != "" {
		if parsedDuration, err = time.ParseDuration(readTimeoutStr); err == nil {
			cnf.WebReadTimeout = parsedDuration
		}
	}

	writeTimeoutStr := os.Getenv("WEB_WRITE_TIMEOUT")
	if writeTimeoutStr != "" {
		if parsedDuration, err = time.ParseDuration(writeTimeoutStr); err == nil {
			cnf.WebWriteTimeout = parsedDuration
		}
	}

	// timeoutStr := os.Getenv("WEB_DEFAULT_WRITE_TIMEOUT")
	// if timeoutStr != "" {
	// 	if parsedDuration, err = time.ParseDuration(timeoutStr); err == nil {
	// 		cnf.WebDefaultWriteTimeout = parsedDuration
	// 	}
	// }

	// if cnf.WebDefaultWriteTimeout > cnf.WebWriteTimeout {
	// 	cnf.WebDefaultWriteTimeout = cnf.WebWriteTimeout
	// }

	idleTimeoutStr := os.Getenv("WEB_IDLE_TIMEOUT")
	if idleTimeoutStr != "" {
		if parsedDuration, err = time.ParseDuration(idleTimeoutStr); err == nil {
			cnf.WebIdleTimeout = parsedDuration
		}
	}

	// maxBytesStr := os.Getenv("WEB_REQUEST_MAX_BYTES")
	// maxBytes, err := strconv.ParseInt(maxBytesStr, 10, 64)
	// if err == nil && maxBytes > 0 {
	// 	cnf.RequestMaxBytes = maxBytes
	// } // else if there is an error this will use the default value.

	cacheTemplatesStr := os.Getenv("WEB_CACHE_TEMPLATES")
	if cacheTemplatesStr != "" {
		if cacheTemplates, err := strconv.ParseBool(cacheTemplatesStr); err == nil {
			cnf.WebCacheTemplates = cacheTemplates
		}
	}

	return cnf
}
