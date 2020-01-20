// Copyright © 2015, XumaK
// All rights reserved. Do not distribute.

package main

import (
	// Standard Packages
	"os"
	"runtime"
	"strings"
)

// Config contains global service configuration values.
type Config struct {

	// BindHost is the IP address to utilize for the job API.
	// Default: "0.0.0.0"
	BindHost string

	// HTTPPort is the port to listen to for HTTP requests for the job API.
	// Default: 7000
	HTTPPort string

	// HTTPSPort is the port to listen to for HTTPS requests for the job API.
	// Default: 7443
	HTTPSPort string
}

var (
	config Config
)

// initConfig initializes the service configuration populating settings from
// defaults, then overriding from environment variables.
func initConfig() error {

	// Defaults
	config = Config{
		BindHost:  "0.0.0.0",
		HTTPPort:  "7000",
		HTTPSPort: "7443",
	}

	processEnvs()

	return nil
}

// processEnvs iterates over environment variables, overriding default
// configuration settings.
func processEnvs() {
	for _, e := range os.Environ() {
		kv := strings.Split(e, "=")
		switch kv[0] {
		case "PREFLGHT_BIND_HOST":
			config.BindHost = kv[1]
		case "PREFLGHT_HTTP_PORT":
			config.HTTPPort = kv[1]
		case "PREFLGHT_HTTPS_PORT":
			config.HTTPSPort = kv[1]
		}
	}

	// If the number of cores to utilize is not explictly set, utilize the
	// maximum.
	if os.Getenv("GOMAXPROCS") == "" {
		runtime.GOMAXPROCS(runtime.NumCPU())
	}
}
