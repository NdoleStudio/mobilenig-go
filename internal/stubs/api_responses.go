package stubs

// CheckDstvUserResponse is a dummy JSOn response for checking a dstv user
func CheckDstvUserResponse() string {
	return `
	{
		"details": {
			"accountStatus":"OPEN",
			"firstName":"ESU",
			"lastName":"INI OBONG BASSEY",
			"customerType":"SUD",
			"invoicePeriod":1,
			"dueDate":"2018-11-13T00:00:00+01:00",
			"customerNumber":275953782
		}
	}
`
}

// PayDstvBillResponse is a dummy JSON response for paying a DSTV bill
func PayDstvBillResponse() string {
	return `
	{
		"trans_id":"122790223",
		"details": {
			"service":"DSTV",
			"package":"DStv Mobile MAXI",
			"smartno":"4131953321",
			"price":"790",
			"status":"SUCCESSFUL",
			"balance":"7931"
		}
	}`
}

// QueryDstvTransactionResponse is a dummy JSON response for querying a dstv transaction
func QueryDstvTransactionResponse() string {
	return `
	{
		"trans_id":"122790223",
		"details": {
			"service":"DSTV",
			"package":"DStv Mobile MAXI",
			"smartno":"4131953321",
			"price":"790",
			"status":"SUCCESSFUL",
			"balance":"7931"
		}
	}`
}

// ErrorResponse is a dummy JSOn response when there is an error
func ErrorResponse() string {
	return `
	{
		"code": "ERR101",
		"description": "Invalid username or api_key"
	}
`
}
