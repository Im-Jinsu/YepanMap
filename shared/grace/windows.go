// When on Windows use standard ServerAndListen method.
//
// +build windows

package grace

import (
	"net/http"
)

type option interface{}

// Serve is a wrapper around standard ListenAndServe method.
func Serve(s *http.Server) error {
	return s.ListenAndServe()
}

// ServeWithOptions does the same as Serve, but takes a set of options to
// configure the app struct.
func ServeWithOptions(s *http.Server, options ...option) error {
	return nil
}
