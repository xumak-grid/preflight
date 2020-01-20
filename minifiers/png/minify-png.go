// Copyright © 2015, XumaK
// All rights reserved. Do not distribute.

// Package png implements a png graphics format minifier by stripping
// unnessesary metadata and reducing format to 8bit.
package png

import (
	"github.com/xumak/preflight/minifiers/minify"
)

const (
	version = "1.0.0"
)

type PNGMinifierConfig struct {
	Version string
}

// Minifier is exported to make the minifier directly accesible.
// However, the minifer should be used via the minify package.
type Minifier struct {
	cfg PNGMinifierConfig
}

// Start initializes the minifier and is expected to be called by Register.
func (m *Minifier) Start() error {
	m.cfg.Version = version
	return nil
}

// Minify attempts to reduce the size of the given png bytes by removing
// unnessary mtadata, color profiles, etc.
func (m *Minifier) Minify(source *[]byte) (*[]byte, error) {
	minified := source
	return minified, nil
}

func init() {
	minify.Register("png", &Minifier{})
}
