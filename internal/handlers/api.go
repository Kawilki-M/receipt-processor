package handlers

import (
	"github.com/go-chi/chi"
	chimiddle "github.com/go-chi/chi/middleware"
)

func Handler(r *chi.Mux) {

	// Global middleware - strip ending slashes
	r.Use(chimiddle.StripSlashes)

	r.Route("/receipts", func(router chi.Router) {
		router.Post("/process", ProcessReceipts)
		router.Get("/{id}/points", GetReceiptPoints)
	})
}
