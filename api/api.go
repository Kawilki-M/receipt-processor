package api

import (
	"encoding/json"
	"net/http"
)

type ProcessReceiptsParams struct {
	Receipt Receipt
}

type ProcessReceiptsResponse struct {
	Code int    `json:"code"`
	Id   string `json:"id"`
}

type GetReceiptPointsParams struct {
	Id string
}

type GetReceiptPointsResponse struct {
	Code   int   `json:"code"`
	Points int64 `json:"points"`
}

type Receipt struct {
	Retailer     string
	PurchaseDate string
	PurchaseTime string
	Items        []Item
	Total        string
}

type Item struct {
	ShortDescription string
	Price            string
}

type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func writeError(w http.ResponseWriter, message string, code int) {
	resp := Error{
		Message: message,
		Code:    code,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	json.NewEncoder(w).Encode(resp)
}

var (
	BadRequestErrorHandler = func(w http.ResponseWriter, err error) {
		writeError(w, err.Error(), http.StatusBadRequest)
	}
	NotFoundErrorHandler = func(w http.ResponseWriter, err error) {
		writeError(w, err.Error(), http.StatusNotFound)
	}
	InternalErrorHandler = func(w http.ResponseWriter) {
		writeError(w, "An unexpected Error Occured.", http.StatusInternalServerError)
	}
)
