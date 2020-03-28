package middlewares

import (
	"log"
	"net/http"
	"time"
)

// SetHeader will set "Content-Type" header as "application/json"
func SetHeader(h http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		h(w, r)
	})
}

// Log record url path, method and request time duration
func Log(h http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		startTime := time.Now()
		h(w, r)
		duration := time.Now().Sub(startTime)

		var message []interface{}
		message = append(message, r.URL.Path)
		message = append(message, r.Method)
		message = append(message, duration)
		log.Printf("%v\n", message)
	})
}
