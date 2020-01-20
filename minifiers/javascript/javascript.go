// Copyright © 2015, XumaK
// All rights reserved. Do not distribute.

// Package javascript implements a javascript minifier by utilizing the Closure
// compiler.
package javascript

import (
	"github.com/xumak/preflight/minifiers/minify"
	"os"
	"strings"
)

const (
	version                 = "1.0.0"
	closureCompilerEndpoint = "http://closure-compiler.appspot.com/compile"
)

type javascriptMinifierConfig struct {
	Version         string
	ClojureLocation string
}

// Minifier is exported to make the minifier directly accesible.
// However, the minifer should be used via the minify package.
type Minifier struct {
	cfg javascriptMinifierConfig
}

// Start initializes the minifier and is expected to be called by Register.
func (m *Minifier) Start() error {
	m.cfg.Version = version
	// try to autodetect clojure compiler?

	// Override defaults with ENV settings if present.
	for _, e := range os.Environ() {
		kv := strings.Split(e, "=")
		switch kv[0] {
		case "PREFLIGHT_CLOJURE_LOCATION":
			m.cfg.ClojureLocation = kv[1]
		}
	}

	return nil
}

// Minify attempts to reduce the size of the given javascript bytes using the
// Clojure compiler.
func (m *Minifier) Minify(source *[]byte) (*[]byte, error) {
	minified := source
	return minified, nil
}

func init() {
	minify.Register("javascript", &Minifier{})
}
