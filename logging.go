package main

import (
	"fmt"
	"log/slog"
	"net/http"
)

func requestLogger(logger *slog.Logger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
			next.ServeHTTP(writer, request)
			logger.Info(fmt.Sprintf("Served request: %s %s", request.Method, request.URL.Path))
		})
	}
}
