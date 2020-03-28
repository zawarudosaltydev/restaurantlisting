package middlewares

import (
	"net/http"
)

// MiddlewareFunc is a middleware http HandlerFunc which also return a new http HandlerFunc, so they can pass to next
type MiddlewareFunc func(http.HandlerFunc) http.HandlerFunc

type MiddleWares struct {
	middlewareFuns []MiddlewareFunc
}

// CreateMiddleWares will return a chain of middleware func
func CreateMiddleWares(funcs ...MiddlewareFunc) MiddleWares {
	middlewarefuns := []MiddlewareFunc{}
	return MiddleWares{
		middlewareFuns: append(middlewarefuns, funcs...),
	}
}

// Run the real HandlerFunc
func (mw MiddleWares) Run(h http.HandlerFunc) http.HandlerFunc {
	if h == nil {
		h = http.DefaultServeMux.ServeHTTP
	}

	for i := range mw.middlewareFuns {
		h = mw.middlewareFuns[len(mw.middlewareFuns)-1-i](h)
	}

	return h
}
