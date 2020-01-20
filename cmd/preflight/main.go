// Copyright © 2015, XumaK
// All rights reserved. Do not distribute.

package main

import (
	// Standard Packages
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	// Filters
	"github.com/xumak/preflight/filters/filter"

	_ "github.com/xumak/preflight/filters/reflows"

	// Minifiers
	"github.com/xumak/preflight/minifiers/minify"

	_ "github.com/xumak/preflight/minifiers/javascript"
	_ "github.com/xumak/preflight/minifiers/png"
	_ "github.com/xumak/preflight/minifiers/svg"

	// External Packages
	log "github.com/civisanalytics/gogelf/gelf"
)

func main() {
	lm := log.Info("Preflight initializing", "Preflight optimizer initializing")
	lm.Add("_version", Version)
	fmt.Println(lm.String())

	var err error

	if err = initConfig(); err != nil {
		lm = log.Panic("Unable to initialize global configuration", err.Error())
		fmt.Println(lm.String())
		os.Exit(1)
	}

	// Log all available filters
	for f := range filter.AvailableFilters {
		jtm := log.Info(fmt.Sprintf("%s available", f), fmt.Sprintf("Preflight filter %s available", f))
		fmt.Println(jtm.String())
	}

	// Log all available minifiers
	for m := range minify.AvailableMinifiers {
		jtm := log.Info(fmt.Sprintf("%s available", m), fmt.Sprintf("Preflight minifier %s available", m))
		fmt.Println(jtm.String())
	}

	// HTTP API multiplexer
	multiplexer := http.NewServeMux()
	multiplexer.Handle("/"+APIVersion+"/status", LogHandlerFunc(statusHandler))
	multiplexer.Handle("/"+APIVersion+"/admin", LogHandlerFunc(adminHandler))
	multiplexer.Handle("/", LogHandlerFunc(invalidHandler))

	s := NewServer(config.BindHost+":"+config.HTTPPort, multiplexer)
	go func() {
		err := s.ListenAndServe()
		if nil != err {
			log.Error("error starting http server", err.Error())
		}
	}()

	lm = log.Info("Startup complete", "Preflight has finished starting")
	lm.Add("_bind_host", config.BindHost)
	lm.Add("_http_port", config.HTTPPort)
	lm.Add("_https_port", config.HTTPSPort)
	// lm.Add("_filters", config.Provisioner)
	// lm.Add("_minifiers", config.Provisioner)
	fmt.Println(lm.String())

	// Listen for OS signals and respond.
	signalChan := make(chan os.Signal)
	signal.Notify(signalChan, syscall.SIGTERM, syscall.SIGINT, os.Kill)
	for {
		sc := <-signalChan
		switch sc {
		case syscall.SIGINT, syscall.SIGTERM, os.Kill:
			lm = log.Info("Shutdown initiated", "Preflight has begun to shutdown")
			lm.Add("_signal", sc.String())
			fmt.Println(lm.String())

			// Gracefully shutdown HTTP interface
			s.Close()
			lm = log.Info("Shutdown complete", "Preflight has finished shutdown")
			lm.Add("_signal", sc.String())
			fmt.Println(lm.String())
			os.Exit(0)
		}
	}
}
