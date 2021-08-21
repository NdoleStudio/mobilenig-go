package mobilenig

import "time"

// DstvProductCode is a code for DStv packages
type DstvProductCode string

const (
	// DstvProductCodePadi is the  DStv Padi Package
	DstvProductCodePadi DstvProductCode = "NLTESE36"

	// DstvYangaBouquet is the DStv Yanga Bouquet E36 package
	DstvYangaBouquet DstvProductCode = "NNJ1E36"

	// DstvProductCodeCompact is the DStv Compact package
	DstvProductCodeCompact DstvProductCode = "COMPE36"

	// DstvProductCodeCompactPlus is the DStv Compact Plus package
	DstvProductCodeCompactPlus DstvProductCode = "COMPLE36"

	// DstvProductCodeCompactPlusXtraView is the DStv Compact Plus + HDPVR/XtraView	package
	DstvProductCodeCompactPlusXtraView DstvProductCode = "DCOHDPV"

	// DstvProductCodePremium is the DStv Premium package
	DstvProductCodePremium DstvProductCode = "PRWE36"

	// DstvProductCodePremiumXtraView is the DStv Premium + HDPVR/XtraView package
	DstvProductCodePremiumXtraView DstvProductCode = "DPRHDP"
)

// PayDstvOptions is the input used when paying a DStv subscription
type PayDstvOptions struct {
	TransactionID   string          `json:"trans_id"`
	Price           string          `json:"price"`
	ProductCode     DstvProductCode `json:"product_code"`
	CustomerName    string          `json:"customer_name"`
	CustomerNumber  string          `json:"customer_number"`
	SmartcardNumber string          `json:"smartno"`
}

// DStvUser is a dstv subscription customer
type DStvUser struct {
	Details struct {
		AccountStatus  string    `json:"accountStatus"`
		Firstname      string    `json:"firstName"`
		Lastname       string    `json:"lastName"`
		CustomerType   string    `json:"customerType"`
		InvoicePeriod  int       `json:"invoicePeriod"`
		DueDate        time.Time `json:"dueDate"`
		CustomerNumber int       `json:"customerNumber"`
	} `json:"details"`
}

// DStvTransaction is the data about a DStv subscription payment
type DStvTransaction struct {
	TransactionID string `json:"trans_id"`
	Details       struct {
		Service         string `json:"service"`
		Package         string `json:"package"`
		SmartcardNumber string `json:"smartno"`
		Price           string `json:"price"`
		Status          string `json:"status"`
		Balance         string `json:"balance"`
	} `json:"details"`
}
