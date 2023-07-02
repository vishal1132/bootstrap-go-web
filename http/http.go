package http

import (
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
)

const (
	defaultPort = 8080
)

func InitServer(r *chi.Mux, writeTimeout, readTimeout int64) *http.Server {
	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", defaultPort),
		Handler:      r,
		WriteTimeout: time.Duration(writeTimeout) * time.Millisecond,
		ReadTimeout:  time.Duration(readTimeout) * time.Millisecond,
	}

	return server
}
