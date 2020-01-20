// Copyright © 2015, XumaK
// All rights reserved. Do not distribute.

// Package filter defines interfaces to be implemented by filters as used by
// Preflight.
package filter

import (
	// Standard Library
	"errors"
	"net/http"
)

var (
	// AvailableFilters provides a list of filters avaialble to Preflight.
	AvailableFilters = make(map[string]Filter)
)

// Filter is the interface that must be implemented by a Preflight filter.
type Filter interface {
	// Start is utilized by filters to do any setup or initializing prior to any
	// processing.
	Start() error
	Process(http.Request) (http.Request, error)
}

// Register makes a filter available by its registered name.
// If Register is called twice with the same name or if filter is nil,
// a panic results.
func Register(name string, p Filter) error {
	if p == nil {
		return errors.New("preflight: Register filter is nil")
	}
	if _, dup := AvailableFilters[name]; dup {
		return errors.New("preflight: Register called twice for filter")
	}

	AvailableFilters[name] = p
	return nil
}
