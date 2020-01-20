// Copyright © 2015, XumaK
// All rights reserved. Do not distribute.

// Package svg implements a svg graphics format minifier.
package svg

import (
	"github.com/xumak/preflight/minifiers/minify"
)

const (
	version = "1.0.0"
)

type SVGMinifierConfig struct {
	Version string
}

// Minifier is exported to make the minifier directly accesible.
// However, the minifer should be used via the minify package.
type Minifier struct {
	cfg SVGMinifierConfig
}

// Start initializes the minifier and is expected to be called by Register.
func (m *Minifier) Start() error {
	m.cfg.Version = version
	return nil
}

// Minify attempts to reduce the size of the given svg.....
func (m *Minifier) Minify(source *[]byte) (*[]byte, error) {
	minified := source
	return minified, nil
}

func init() {
	minify.Register("svg", &Minifier{})
}
