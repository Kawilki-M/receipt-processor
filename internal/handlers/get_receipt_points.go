package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Kawilki-M/receipt-processor/api"
	"github.com/Kawilki-M/receipt-processor/internal/tools"
	"github.com/go-chi/chi"
	log "github.com/sirupsen/logrus"
)

var errReceiptNotFound = fmt.Errorf("No receipt found for that ID.")

func GetReceiptPoints(w http.ResponseWriter, r *http.Request) {
	// Gather query params
	id := chi.URLParam(r, "id")

	// Initialize pointer to local database
	var err error
	var database *tools.DatabaseInterface
	database, err = tools.GetDatabase()
	if err != nil {
		log.Error(err)
		api.InternalErrorHandler(w)
		return
	}

	// Gather receipt details from local db
	var receiptDetails *tools.ReceiptDetails
	receiptDetails = (*database).GetReceiptDetails(id)
	if receiptDetails == nil {
		log.Error(err)
		api.NotFoundErrorHandler(w, errReceiptNotFound)
		return
	}

	// Respond to request
	var response = api.GetReceiptPointsResponse{
		Points: (*receiptDetails).Points,
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		log.Error(err)
		api.InternalErrorHandler(w)
		return
	}
}
