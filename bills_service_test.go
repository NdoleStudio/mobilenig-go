package mobilenig

import (
	"context"
	"errors"
	"net/http"
	"net/url"
	"testing"
	"time"

	"github.com/NdoleStudio/mobilenig-go/internal/helpers"
	"github.com/NdoleStudio/mobilenig-go/internal/stubs"
	"github.com/stretchr/testify/assert"
)

const (
	testUsername = "test_username"
	testAPIKey   = "test_api_key"
)

func TestBillsService_CheckDStvUser_ResponseConstructedCorrectly(t *testing.T) {
	// Setup
	t.Parallel()
	server := helpers.MakeTestServer(http.StatusOK, stubs.CheckDstvUserResponse())

	// Arrange
	baseURL, _ := url.Parse(server.URL)
	client := New(WithBaseURL(baseURL))
	date, _ := time.Parse("2006-01-02T15:04:05-07:00", "2018-11-13T00:00:00+01:00")

	// Act
	user, _, err := client.Bills.CheckDStvUser(context.Background(), "4131953321")

	// Assert
	assert.NoError(t, err)

	assert.Equal(t, "OPEN", user.Details.AccountStatus)
	assert.Equal(t, "ESU", user.Details.Firstname)
	assert.Equal(t, "INI OBONG BASSEY", user.Details.Lastname)
	assert.Equal(t, "SUD", user.Details.CustomerType)
	assert.Equal(t, 1, user.Details.InvoicePeriod)
	assert.Equal(t, date, user.Details.DueDate)
	assert.Equal(t, int64(275953782), user.Details.CustomerNumber)

	// Teardown
	server.Close()
}

func TestBillsService_CheckDStvUser_RequestConstructedCorrectly(t *testing.T) {
	// Setup
	t.Parallel()

	environments := []Environment{LiveEnvironment, TestEnvironment}
	for _, environment := range environments {
		t.Run(environment.String(), func(t *testing.T) {
			// Setup
			t.Parallel()

			// Arrange
			request := new(http.Request)
			server := helpers.MakeRequestCapturingTestServer(http.StatusOK, stubs.CheckDstvUserResponse(), request)

			baseURL, _ := url.Parse(server.URL)
			username := testUsername
			apiKey := testAPIKey
			smartcardNumber := "4131953321"

			client := New(WithBaseURL(baseURL), WithAPIKey(apiKey), WithUsername(username), WithEnvironment(environment))

			// Act
			_, _, _ = client.Bills.CheckDStvUser(context.Background(), smartcardNumber)

			// Assert
			assert.Equal(t, "/bills/user_check", request.URL.Path)
			assert.Equal(t, username, request.URL.Query().Get("username"))
			assert.Equal(t, apiKey, request.URL.Query().Get("api_key"))
			assert.Equal(t, "DSTV", request.URL.Query().Get("service"))
			assert.Equal(t, smartcardNumber, request.URL.Query().Get("number"))

			// Teardown
			server.Close()
		})
	}
}

func TestBillsService_CheckDStvUser_ErrorResponseConstructedCorrectly(t *testing.T) {
	t.Parallel()

	environments := []Environment{LiveEnvironment, TestEnvironment}
	for _, environment := range environments {
		t.Run(environment.String(), func(t *testing.T) {
			// Setup
			t.Parallel()

			// Arrange
			server := helpers.MakeTestServer(http.StatusOK, stubs.ErrorResponse())
			baseURL, _ := url.Parse(server.URL)
			smartcardNumber := "4131953321"

			client := New(WithBaseURL(baseURL), WithEnvironment(environment))

			// Act
			_, resp, err := client.Bills.CheckDStvUser(context.Background(), smartcardNumber)

			// Assert
			assert.Error(t, err)

			assert.Equal(t, "ERR101", resp.Error.Code)
			assert.Equal(t, "Invalid username or api_key", resp.Error.Description)

			// Teardown
			server.Close()
		})
	}
}

func TestBillsService_CheckDStvUser_CancelledContext(t *testing.T) {
	t.Parallel()

	environments := []Environment{LiveEnvironment, TestEnvironment}
	for _, environment := range environments {
		t.Run(environment.String(), func(t *testing.T) {
			t.Parallel()

			// Arrange
			server := helpers.MakeTestServer(http.StatusOK, stubs.CheckDstvUserResponse())
			baseURL, _ := url.Parse(server.URL)
			client := New(WithBaseURL(baseURL), WithEnvironment(environment))
			ctx, cancel := context.WithCancel(context.Background())
			cancel()

			// Act
			_, _, err := client.Bills.CheckDStvUser(ctx, "")

			// Assert
			assert.True(t, errors.Is(err, context.Canceled))
			server.Close()
		})
	}
}

func TestBillsService_CheckDStvUser_InvalidResponse(t *testing.T) {
	t.Parallel()
	environments := []Environment{LiveEnvironment, TestEnvironment}
	for _, environment := range environments {
		t.Run(environment.String(), func(t *testing.T) {
			t.Parallel()

			// Arrange
			server := helpers.MakeTestServer(http.StatusOK, "<not-a-json></not-a-json>")
			baseURL, _ := url.Parse(server.URL)
			client := New(WithBaseURL(baseURL), WithEnvironment(environment))

			// Act
			_, _, err := client.Bills.CheckDStvUser(context.Background(), "")

			// Assert
			assert.Error(t, err)

			// Teardown
			server.Close()
		})
	}
}

func TestBillsService_PayDStv_ResponseConstructedCorrectly(t *testing.T) {
	t.Parallel()

	// Arrange
	server := helpers.MakeTestServer(http.StatusOK, stubs.PayDstvBillResponse())
	baseURL, _ := url.Parse(server.URL)
	client := New(WithBaseURL(baseURL))

	// Act
	transaction, _, err := client.Bills.PayDStv(context.Background(), &PayDstvOptions{})

	// Assert
	assert.NoError(t, err)

	assert.Equal(t, "122790223", transaction.TransactionID)
	assert.Equal(t, "DSTV", transaction.Details.Service)
	assert.Equal(t, "DStv Mobile MAXI", transaction.Details.Package)
	assert.Equal(t, "4131953321", transaction.Details.SmartcardNumber)
	assert.Equal(t, "790", transaction.Details.Price)
	assert.Equal(t, "SUCCESSFUL", transaction.Details.Status)
	assert.Equal(t, "7931", transaction.Details.Balance)

	// Teardown
	server.Close()
}

func TestBillsService_PayDStv_RequestConstructedCorrectly(t *testing.T) {
	t.Parallel()

	environments := []Environment{LiveEnvironment, TestEnvironment}
	for _, environment := range environments {
		t.Run(environment.String(), func(t *testing.T) {
			// Arrange
			request := new(http.Request)
			server := helpers.MakeRequestCapturingTestServer(http.StatusOK, stubs.CheckDstvUserResponse(), request)

			baseURL, _ := url.Parse(server.URL)
			username := testUsername
			apiKey := testAPIKey
			smartcardNumber := "4131953321"
			customerNumber := "275953782"
			customerName := "ESU INI OBONG BASSEY"
			price := "790"
			transactionID := "122790223"

			client := New(WithBaseURL(baseURL), WithAPIKey(apiKey), WithUsername(username), WithEnvironment(environment))

			// Act
			_, _, _ = client.Bills.PayDStv(context.Background(), &PayDstvOptions{
				TransactionID:   transactionID,
				Price:           price,
				ProductCode:     DstvProductCodePremium,
				CustomerName:    customerName,
				CustomerNumber:  customerNumber,
				SmartcardNumber: smartcardNumber,
			})

			// Assert
			uri := "/bills/dstv"
			if environment == TestEnvironment {
				uri += "_test"
			}

			assert.Equal(t, uri, request.URL.Path)
			assert.Equal(t, username, request.URL.Query().Get("username"))
			assert.Equal(t, apiKey, request.URL.Query().Get("api_key"))
			assert.Equal(t, smartcardNumber, request.URL.Query().Get("smartno"))
			assert.Equal(t, string(DstvProductCodePremium), request.URL.Query().Get("product_code"))
			assert.Equal(t, customerName, request.URL.Query().Get("customer_name"))
			assert.Equal(t, customerNumber, request.URL.Query().Get("customer_number"))
			assert.Equal(t, price, request.URL.Query().Get("price"))
			assert.Equal(t, transactionID, request.URL.Query().Get("trans_id"))

			// Teardown
			server.Close()
		})
	}
}

func TestBillsService_PayDStv_ErrorResponseConstructedCorrectly(t *testing.T) {
	t.Parallel()

	environments := []Environment{LiveEnvironment, TestEnvironment}
	for _, environment := range environments {
		t.Run(environment.String(), func(t *testing.T) {
			// Arrange
			server := helpers.MakeTestServer(http.StatusOK, stubs.ErrorResponse())
			baseURL, _ := url.Parse(server.URL)

			client := New(WithBaseURL(baseURL), WithEnvironment(environment))

			// Act
			_, resp, err := client.Bills.PayDStv(context.Background(), &PayDstvOptions{})

			// Assert
			assert.Error(t, err)

			assert.Equal(t, "ERR101", resp.Error.Code)
			assert.Equal(t, "Invalid username or api_key", resp.Error.Description)

			server.Close()
		})
	}
}

func TestBillsService_PayDStv_CancelledContext(t *testing.T) {
	t.Parallel()

	environments := []Environment{LiveEnvironment, TestEnvironment}
	for _, environment := range environments {
		t.Run(environment.String(), func(t *testing.T) {
			// Arrange
			server := helpers.MakeTestServer(http.StatusOK, stubs.CheckDstvUserResponse())
			baseURL, _ := url.Parse(server.URL)
			client := New(WithBaseURL(baseURL), WithEnvironment(environment))
			ctx, cancel := context.WithCancel(context.Background())
			cancel()

			// Act
			_, _, err := client.Bills.PayDStv(ctx, &PayDstvOptions{})

			// Assert
			assert.True(t, errors.Is(err, context.Canceled))

			// Teardown
			server.Close()
		})
	}
}

func TestBillsService_PayDStv_InvalidResponse(t *testing.T) {
	t.Parallel()

	environments := []Environment{LiveEnvironment, TestEnvironment}
	for _, environment := range environments {
		t.Run(environment.String(), func(t *testing.T) {
			t.Parallel()

			// Arrange
			server := helpers.MakeTestServer(http.StatusOK, "<not-a-json></not-a-json>")
			baseURL, _ := url.Parse(server.URL)
			client := New(WithBaseURL(baseURL), WithEnvironment(environment))

			// Act
			_, _, err := client.Bills.PayDStv(context.Background(), &PayDstvOptions{})

			// Assert
			assert.Error(t, err)

			// Teardown
			server.Close()
		})
	}
}

func TestBillsService_PayDStv_NilOptions(t *testing.T) {
	t.Parallel()

	environments := []Environment{LiveEnvironment, TestEnvironment}
	for _, environment := range environments {
		t.Run(environment.String(), func(t *testing.T) {
			// Arrange
			server := helpers.MakeTestServer(http.StatusOK, stubs.CheckDstvUserResponse())
			baseURL, _ := url.Parse(server.URL)
			client := New(WithBaseURL(baseURL), WithEnvironment(environment))
			ctx, cancel := context.WithCancel(context.Background())
			cancel()

			// Act
			_, _, err := client.Bills.PayDStv(ctx, nil)

			// Assert
			assert.Error(t, err)

			// Teardown
			server.Close()
		})
	}
}

func TestBillsService_QueryDStv_ResponseConstructedCorrectly(t *testing.T) {
	t.Parallel()

	// Arrange
	server := helpers.MakeTestServer(http.StatusOK, stubs.PayDstvBillResponse())
	baseURL, _ := url.Parse(server.URL)
	client := New(WithBaseURL(baseURL))

	// Act
	transaction, _, err := client.Bills.QueryDStv(context.Background(), "122790223")

	// Assert
	assert.NoError(t, err)

	assert.Equal(t, "122790223", transaction.TransactionID)
	assert.Equal(t, "DSTV", transaction.Details.Service)
	assert.Equal(t, "DStv Mobile MAXI", transaction.Details.Package)
	assert.Equal(t, "4131953321", transaction.Details.SmartcardNumber)
	assert.Equal(t, "790", transaction.Details.Price)
	assert.Equal(t, "SUCCESSFUL", transaction.Details.Status)
	assert.Equal(t, "7931", transaction.Details.Balance)

	// Teardown
	server.Close()
}

func TestBillsService_QueryDStv_RequestConstructedCorrectly(t *testing.T) {
	t.Parallel()

	environments := []Environment{LiveEnvironment, TestEnvironment}
	for _, environment := range environments {
		t.Run(environment.String(), func(t *testing.T) {
			// Arrange
			request := new(http.Request)
			server := helpers.MakeRequestCapturingTestServer(http.StatusOK, stubs.CheckDstvUserResponse(), request)

			baseURL, _ := url.Parse(server.URL)
			username := testUsername
			apiKey := testAPIKey
			transactionID := "122790223"

			client := New(WithBaseURL(baseURL), WithAPIKey(apiKey), WithUsername(username), WithEnvironment(environment))

			// Act
			_, _, _ = client.Bills.QueryDStv(context.Background(), transactionID)

			// Assert
			assert.Equal(t, "/bills/query", request.URL.Path)
			assert.Equal(t, username, request.URL.Query().Get("username"))
			assert.Equal(t, apiKey, request.URL.Query().Get("api_key"))
			assert.Equal(t, transactionID, request.URL.Query().Get("trans_id"))

			// Teardown
			server.Close()
		})
	}
}

func TestBillsService_QueryDStv_ErrorResponseConstructedCorrectly(t *testing.T) {
	t.Parallel()

	environments := []Environment{LiveEnvironment, TestEnvironment}
	for _, environment := range environments {
		t.Run(environment.String(), func(t *testing.T) {
			// Arrange
			server := helpers.MakeTestServer(http.StatusOK, stubs.ErrorResponse())
			baseURL, _ := url.Parse(server.URL)

			client := New(WithBaseURL(baseURL), WithEnvironment(environment))

			// Act
			_, resp, err := client.Bills.QueryDStv(context.Background(), "122790223")

			// Assert
			assert.Error(t, err)

			assert.Equal(t, "ERR101", resp.Error.Code)
			assert.Equal(t, "Invalid username or api_key", resp.Error.Description)

			server.Close()
		})
	}
}

func TestBillsService_QueryDStv_CancelledContext(t *testing.T) {
	t.Parallel()

	environments := []Environment{LiveEnvironment, TestEnvironment}
	for _, environment := range environments {
		t.Run(environment.String(), func(t *testing.T) {
			t.Parallel()

			// Arrange
			server := helpers.MakeTestServer(http.StatusOK, stubs.CheckDstvUserResponse())
			baseURL, _ := url.Parse(server.URL)
			client := New(WithBaseURL(baseURL), WithEnvironment(environment))
			ctx, cancel := context.WithCancel(context.Background())
			cancel()

			// Act
			_, _, err := client.Bills.QueryDStv(ctx, "122790223")

			// Assert
			assert.True(t, errors.Is(err, context.Canceled))

			// Teardown
			server.Close()
		})
	}
}

func TestBillsService_QueryDStv_InvalidResponse(t *testing.T) {
	t.Parallel()

	environments := []Environment{LiveEnvironment, TestEnvironment}
	for _, environment := range environments {
		t.Run(environment.String(), func(t *testing.T) {
			t.Parallel()

			// Arrange
			server := helpers.MakeTestServer(http.StatusOK, "<not-a-json></not-a-json>")
			baseURL, _ := url.Parse(server.URL)
			client := New(WithBaseURL(baseURL), WithEnvironment(environment))

			// Act
			_, _, err := client.Bills.QueryDStv(context.Background(), "122790223")

			// Assert
			assert.Error(t, err)

			// Teardown
			server.Close()
		})
	}
}
