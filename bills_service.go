package mobilenig

import (
	"context"
	"encoding/json"
	"errors"
	"strconv"
)

// BillsService is the API client for the `/bills/` endpoint
type BillsService service

const (
	billsServiceDStv = "DSTV"
)

// CheckDStvUser validates a DStv smartcard number
// POST /bills/user_check
// API Doc: https://mobilenig.com/API/docs/dstv
func (service *BillsService) CheckDStvUser(ctx context.Context, smartcardNumber string) (*DStvUser, *Response, error) {
	payload := map[string]string{
		"service": billsServiceDStv,
		"number":  smartcardNumber,
	}

	request, err := service.client.newRequest(ctx, "/bills/user_check", payload)
	if err != nil {
		return nil, nil, err
	}

	resp, err := service.client.do(request)
	if err != nil {
		return nil, resp, err
	}

	var dstvUser DStvUser
	if err = json.Unmarshal(*resp.Body, &dstvUser); err != nil {
		return nil, resp, err
	}

	return &dstvUser, resp, nil
}

// GetDStvPackage returns the client's current DStv package.
// POST /bills/get_package
// API Doc: https://mobilenig.com/API/docs/dstv
func (service *BillsService) GetDStvPackage(ctx context.Context, customerNumber int64) (*string, *Response, error) {
	payload := map[string]string{
		"service":        billsServiceDStv,
		"customerNumber": strconv.FormatInt(customerNumber, 10),
	}

	request, err := service.client.newRequest(ctx, "/bills/get_package", payload)
	if err != nil {
		return nil, nil, err
	}

	resp, err := service.client.do(request)
	if err != nil {
		return nil, resp, err
	}

	details := map[string]*string{}
	if err = json.Unmarshal(*resp.Body, &details); err != nil {
		return nil, resp, err
	}

	return details["packageName"], resp, nil
}

// PayDStv pays a DStv subscription
// POST /bills/dstv
// API Doc: https://mobilenig.com/API/docs/dstv
func (service *BillsService) PayDStv(ctx context.Context, options *PayDstvOptions) (*DStvTransaction, *Response, error) {
	if options == nil {
		return nil, nil, errors.New("options cannot be nil")
	}

	payload := map[string]string{
		"product_code":    string(options.ProductCode),
		"customer_name":   options.CustomerName,
		"customer_number": options.CustomerNumber,
		"price":           options.Price,
		"smartno":         options.SmartcardNumber,
		"trans_id":        options.TransactionID,
	}

	uri := "/bills/dstv"
	if service.client.environment == TestEnvironment {
		uri += "_test"
	}

	request, err := service.client.newRequest(ctx, uri, payload)
	if err != nil {
		return nil, nil, err
	}

	resp, err := service.client.do(request)
	if err != nil {
		return nil, resp, err
	}

	var transaction DStvTransaction
	if err = json.Unmarshal(*resp.Body, &transaction); err != nil {
		return nil, resp, err
	}

	return &transaction, resp, nil
}

// QueryDStv fetches a DStv transaction using the transaction ID
// POST /bills/dstv
// API Doc: https://mobilenig.com/API/docs/dstv
func (service *BillsService) QueryDStv(ctx context.Context, transactionID string) (*DStvTransaction, *Response, error) {
	payload := map[string]string{
		"trans_id": transactionID,
	}

	request, err := service.client.newRequest(ctx, "/bills/query", payload)
	if err != nil {
		return nil, nil, err
	}

	resp, err := service.client.do(request)
	if err != nil {
		return nil, resp, err
	}

	var transaction DStvTransaction
	if err = json.Unmarshal(*resp.Body, &transaction); err != nil {
		return nil, resp, err
	}

	return &transaction, resp, nil
}
