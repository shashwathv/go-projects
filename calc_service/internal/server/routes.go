package server

import (
	"net/http"

	"calc_service/internal/handlers"
)

func RegisterRoutes(mux *http.ServeMux, h *handlers.Handler) {
	mux.HandleFunc("POST /add", h.Add)
	mux.HandleFunc("POST /subtract", h.Subtract)
	mux.HandleFunc("POST /multiply", h.Multiply)
	mux.HandleFunc("POST /divide", h.Divide)
	mux.HandleFunc("POST /sum", h.Sum)

}
