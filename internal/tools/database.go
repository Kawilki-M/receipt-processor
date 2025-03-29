package tools

import (
	log "github.com/sirupsen/logrus"
)

type ReceiptDetails struct {
	Id     string
	Points int64
}

type DatabaseInterface interface {
	AddReceiptDetails(id string, points int64) *ReceiptDetails
	GetReceiptDetails(id string) *ReceiptDetails
	SetupDatabase() error
}

func GetDatabase() (*DatabaseInterface, error) {
	var database DatabaseInterface = &mockDB{}

	var err error = database.SetupDatabase()
	if err != nil {
		log.Error(err)
		return nil, err
	}

	return &database, nil
}
