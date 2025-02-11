package middlewares

import (
	"SparksSport/cmd/app/core/utils"
	"context"
	"github.com/google/uuid"
	"net/http"
)

type responseWriter struct {
	http.ResponseWriter
	statusCode int
	body       []byte
}

func (rw *responseWriter) WriteHeader(statusCode int) {
	rw.statusCode = statusCode
	rw.ResponseWriter.WriteHeader(statusCode)
}

func (rw *responseWriter) Write(b []byte) (int, error) {
	rw.body = append(rw.body, b...)
	return rw.ResponseWriter.Write(b)
}

func RequestIDMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requestID := uuid.New().String()
		ctx := context.WithValue(r.Context(), "RequestID", requestID)

		// Create a custom response writer to capture response data
		rw := &responseWriter{ResponseWriter: w, statusCode: http.StatusOK}

		w.Header().Set("X-Request-ID", requestID)

		// Call the next handler
		next.ServeHTTP(rw, r.WithContext(ctx))

		// Log the response data
		utils.LogInfo(" Response Status: "+string(rune(rw.statusCode))+" - Response Body: "+string(rw.body), ctx)
	})
}
