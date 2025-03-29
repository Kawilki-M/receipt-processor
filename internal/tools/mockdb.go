package tools

type mockDB struct{}

var receiptDetailsTable map[string]ReceiptDetails

func (d *mockDB) AddReceiptDetails(id string, points int64) *ReceiptDetails {
	receiptDetails := ReceiptDetails{id, points}
	receiptDetailsTable[id] = receiptDetails

	return &receiptDetails
}

func (d *mockDB) GetReceiptDetails(id string) *ReceiptDetails {
	receiptDetails, inTable := receiptDetailsTable[id]

	if !inTable {
		return nil
	}

	return &receiptDetails
}

func (d *mockDB) SetupDatabase() error {
	if receiptDetailsTable == nil {
		receiptDetailsTable = make(map[string]ReceiptDetails)
	}

	return nil
}
