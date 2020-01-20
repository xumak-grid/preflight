// Copyright © 2015, XumaK
// All rights reserved. Do not distribute.

package main

import (
	// Standard Packages
	"fmt"
	"net/http"

	// External Packages
	log "github.com/civisanalytics/gogelf/gelf"
)

func adminHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		// Render JSON information about requests, file optimizations, etc.
		w.Header().Set("Content-Type", "application/json")
		// TODO: Collect information and display
	default:
		// Return HTTP status code 405: Method not allowed.
		http.Error(w, http.StatusText(405), 405)
	}
}

func statusHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		// Render JSON information about Preflight.
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, "{ \n\"status\": \"running\", \"version\": %s }", Version)
	default:
		// Return HTTP status code 405: Method not allowed.
		http.Error(w, http.StatusText(405), 405)
	}
}

func invalidHandler(w http.ResponseWriter, r *http.Request) {
	http.Error(w, http.StatusText(400), 400)
}

func LogHandlerFunc(next http.HandlerFunc) http.HandlerFunc {
	return LogHandler(next)
}

func LogHandler(next http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Info("http request", fmt.Sprintf("%v %v", r.Method, r.URL))
		am := log.NewMessage(log.LevelInfo, "HTTP API Request", "HTTP API Request Access")
		am.Add("_url", r.URL)
		am.Add("_method", r.Method)
		am.Add("_protocol", r.Proto)
		am.Add("_remote", r.RemoteAddr)
		am.Add("_uri", r.RequestURI)
		fmt.Println(am.String())
		next.ServeHTTP(w, r)
	}
}
