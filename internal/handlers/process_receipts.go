package handlers

import (
	"encoding/json"
	"fmt"
	"math"
	"net/http"
	"strconv"
	"strings"
	"time"
	"unicode"
	"unicode/utf8"

	"github.com/Kawilki-M/receipt-processor/api"
	"github.com/Kawilki-M/receipt-processor/internal/tools"
	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
)

var errReceiptInvalid = fmt.Errorf("The receipt is invalid.")

func ProcessReceipts(w http.ResponseWriter, r *http.Request) {
	// Gather query params
	var receipt api.Receipt
	err := json.NewDecoder(r.Body).Decode(&receipt)
	if err != nil {
		log.Error(err)
		api.BadRequestErrorHandler(w, errReceiptInvalid)
		return
	}

	defer r.Body.Close()

	// Initialize pointer to local database
	var database *tools.DatabaseInterface
	database, err = tools.GetDatabase()
	if err != nil {
		log.Error(err)
		api.InternalErrorHandler(w)
		return
	}

	// Gather receipt details and store in local db
	receiptId := uuid.New().String()
	receiptPoints, err := calculateReceiptPoints(receipt)
	if err != nil {
		log.Error(err)
		api.BadRequestErrorHandler(w, errReceiptInvalid)
		return
	}

	var receiptDetails *tools.ReceiptDetails
	receiptDetails = (*database).AddReceiptDetails(receiptId, receiptPoints)
	if receiptDetails == nil {
		api.BadRequestErrorHandler(w, errReceiptInvalid)
		return
	}

	// Respond to request
	var response = api.ProcessReceiptsResponse{
		Id:   receiptId,
		Code: http.StatusOK,
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		log.Error(err)
		api.InternalErrorHandler(w)
		return
	}
}

func calculateReceiptPoints(receipt api.Receipt) (int64, error) {
	var points int64 = 0

	// Process - points related to the retailer
	points += getAlphanumericCharCount(receipt.Retailer)

	// Process - points related to receipt total
	total, err := strconv.ParseFloat(receipt.Total, 32)
	if err != nil {
		return 0, err
	}

	if math.Mod(total, 1) == 0 {
		points += 50
	}

	if math.Mod(total, .25) == 0 {
		points += 25
	}

	// Process - points related to receipt items
	points += int64(len(receipt.Items)) / 2 * 5

	for _, item := range receipt.Items {
		if utf8.RuneCountInString(strings.TrimSpace(item.ShortDescription)) % 3 == 0 {

			price, err := strconv.ParseFloat(item.Price, 32)
			if err != nil {
				return 0, err
			}

			points += int64(math.Ceil(price * .2))
		}
	}

	// Process - points related to purchase date and time
	layout := "2006-01-02"
	purchaseDate, err := time.Parse(layout, receipt.PurchaseDate)
	if err != nil {
		return 0, err
	}

	if purchaseDate.Day() % 2 == 1 {
		points += 6
	}

	layout = "15:04"
	purchaseTime, err := time.Parse(layout, receipt.PurchaseTime)
	if err != nil {
		return 0, err
	}

	startTime, _ := time.Parse(layout, "14:00")
	endTime, _ := time.Parse(layout, "16:00")
	if purchaseTime.After(startTime) && purchaseTime.Before(endTime) {
		points += 10
	}

	return points, nil
}

func getAlphanumericCharCount(s string) int64 {
	var aCCount int64 = 0

	// We will not consider any special characters to be alphanumeric for this service
	for _, c := range s {
		if unicode.IsLetter(c) || unicode.IsDigit(c) {
			aCCount++
		}
	}

	return aCCount
}
