package main

import (
	"log/slog"
	"net/http"
)

func requestLogger(logger *slog.Logger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
			next.ServeHTTP(writer, request)
			logger.Info("Served request",
				slog.String("method", request.Method),
				slog.String("path", request.URL.Path),
				slog.String("client_ip", request.RemoteAddr),
			)
		})
	}
}
