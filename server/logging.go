package server

import (
	"log"
	"net/http"
)

type statusCapturingResponseWriter struct {
	http.ResponseWriter
	status int
}

// status code 정보를 header에 입력
func (s *statusCapturingResponseWriter) WriteHeader(statusCode int) {
	s.status = statusCode
	s.ResponseWriter.WriteHeader(statusCode)
}

// response 결과를 받아 중간 단계에서 logging 정보를 출력
func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		statusCapturingWriter := &statusCapturingResponseWriter{
			ResponseWriter: w,
			status:         http.StatusOK,
		}

		next.ServeHTTP(statusCapturingWriter, r)

		log.Printf("[%s] %s - %d", r.Method, r.RequestURI, statusCapturingWriter.status)
	})
}
