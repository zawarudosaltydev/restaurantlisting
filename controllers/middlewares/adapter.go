package middlewares

import (
	"net/http"
)

// MiddlewareFunc is a middleware http HandlerFunc which also return next http HandlerFunc
type MiddlewareFunc func(http.HandlerFunc) http.HandlerFunc

// MiddleWares contains a slice of MiddlewareFunc
type MiddleWares struct {
	middlewareFuns []MiddlewareFunc
}

// CreateMiddleWares will append middleware to MiddleWares
func CreateMiddleWares(funcs ...MiddlewareFunc) MiddleWares {
	middlewarefuns := []MiddlewareFunc{}
	return MiddleWares{
		middlewareFuns: append(middlewarefuns, funcs...),
	}
}

// Run the all the middlewares, starts from the real HandlerFunc
// and will be warpped by previous one, so the first one will be
// at the beginning, which make sure runs in order.
func (mw MiddleWares) Run(h http.HandlerFunc) http.HandlerFunc {
	if h == nil {
		return http.DefaultServeMux.ServeHTTP
	}

	for i := range mw.middlewareFuns {
		h = mw.middlewareFuns[len(mw.middlewareFuns)-1-i](h)
	}

	return h
}
