// Copyright © 2015, XumaK
// All rights reserved. Do not distribute.

// Package minify defines interfaces to be implemented by minifiers as used by
// Preflight.
package minify

import (
	"errors"
)

var (
	// AvailableMinifiers provides a list of minifiers avaialble to Preflight.
	AvailableMinifiers = make(map[string]Minifier)
)

// Minifier is the interface that must be implemented by a Preflight minifier.
type Minifier interface {
	// Start is utilized by minifiers to do any setup or initializing prior to any
	// minification.
	Start() error
	Minify(source *[]byte) (*[]byte, error)
}

// Register makes a minifier available by its registered name.
// If Register is called twice with the same name or if minifier is nil,
// a panic results.
func Register(name string, p Minifier) error {
	if p == nil {
		return errors.New("preflight: Register minifier is nil")
	}
	if _, dup := AvailableMinifiers[name]; dup {
		return errors.New("preflight: Register called twice for minifier")
	}

	AvailableMinifiers[name] = p
	return nil
}
