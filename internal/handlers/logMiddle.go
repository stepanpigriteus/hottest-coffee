package handler

import (
	"hot/internal/pkg/config"
	"log/slog"
	"net/http"
	"time"
)

type responseRecorder struct {
	http.ResponseWriter
	statusCode int
}

func (rec *responseRecorder) WriteHeader(code int) {
	rec.statusCode = code
	rec.ResponseWriter.WriteHeader(code)
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		rec := &responseRecorder{ResponseWriter: w, statusCode: http.StatusOK}

		config.Logger.Info("HTTP запрос",
			slog.String("method", r.Method),
			slog.String("url", r.URL.Path),
			slog.String("remote_addr", r.RemoteAddr),
		)

		next.ServeHTTP(rec, r)

		if rec.statusCode >= 400 {
			config.Logger.Error("Ошибка HTTP",
				slog.String("method", r.Method),
				slog.String("url", r.URL.Path),
				slog.Int("status", rec.statusCode),
				slog.String("duration", time.Since(start).String()),
			)
		} else {
			config.Logger.Info("HTTP обработан",
				slog.String("method", r.Method),
				slog.String("url", r.URL.Path),
				slog.Int("status", rec.statusCode),
				slog.String("duration", time.Since(start).String()),
			)
		}
	})
}
