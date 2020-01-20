// Copyright © 2015, XumaK
// All rights reserved. Do not distribute.

package reflows

import (
	"github.com/xumak/preflight/filters/filter"
	"net/http"
)

const (
	version = "1.0.0"
)

// Filter is exported to make the filter directly accesible.
// However, the filter should be used via the filter package.
type Filter struct {
	cfg reduceReflowsFilterConfig
}

type reduceReflowsFilterConfig struct {
	// Version is the filter version.
	Version string
}

// Start initializes the filter and is expected to be called by Register
func (f *Filter) Start() error {
	f.cfg.Version = version
	return nil
}

// Process processes a net/http.Request object and returns the same after some processing.
func (f *Filter) Process(r http.Request) (http.Request, error) {

	return r, nil
}

// Remove unused CSS rules.
func removeUnusedCSSRules() {
	return
}

func init() {
	filter.Register("reflows", &Filter{})
}
