package dto

// StockInfo defines the structure for book's stock info
// swagger:model
type StockInfo struct {
	// the number of books in the stock
	//
	// required: false
	// min: 1
	CurrentStock int `json:"currentstock"` // the number of books in the stock
}
