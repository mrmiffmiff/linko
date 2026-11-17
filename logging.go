package main

import (
	"log"
	"net/http"
)

func requestLogger(logger *log.Logger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
			next.ServeHTTP(writer, request)
			logger.Printf("Served request: %s %s", request.Method, request.URL.Path)
		})
	}
}
