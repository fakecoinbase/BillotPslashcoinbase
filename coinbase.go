package coinbase

import (
	"net/http"

	"github.com/BillotP/coinbase/lib/auth"
	"github.com/BillotP/coinbase/lib/models"
	"github.com/BillotP/coinbase/lib/rpc"
)

// Client is the struct from which all API requests are made
type Client struct {
	rpc rpc.RPC
}

// CoinbaseClient is the global blablabla
var CoinbaseClient *Client

// New return an authenticated client
func New(pubkey *string, privkey *string) *Client {
	auth.CBAccount = auth.NewClient(pubkey, privkey)
	CoinbaseClient = &Client{
		rpc: rpc.RPC{
			Auth: auth.CBAccount,
		},
	}
	return CoinbaseClient
}

// Get sends a GET request and marshals response data into holder
func (c Client) Get(path string, params interface{}, holder interface{}) error {
	return c.rpc.Request(http.MethodGet, path, params, &holder)
}

// Post sends a POST request and marshals response data into holder
func (c Client) Post(path string, params interface{}, holder interface{}) error {
	return c.rpc.Request(http.MethodPost, path, params, &holder)
}

// Delete sends a DELETE request and marshals response data into holder
func (c Client) Delete(path string, params interface{}, holder interface{}) error {
	return c.rpc.Request(http.MethodDelete, path, params, &holder)
}

// Put sends a PUT request and marshals response data into holder
func (c Client) Put(path string, params interface{}, holder interface{}) error {
	return c.rpc.Request(http.MethodPut, path, params, &holder)
}

// GetSpotPrice returns the last sport price for a currency pair
func (c Client) GetSpotPrice(base, quote string) (*models.SpotPrice, error) {
	var spot models.SpotPrice

	if err := c.Get("v2/prices/"+base+"-"+quote+"/spot", nil, &spot); err != nil {
		return nil, err
	}
	return &spot, nil
}

// GetAccounts returns a list of all coinbase accounts
func (c Client) GetAccounts() (*models.Accounts, error) {
	var accounts models.Accounts
	if err := c.Get("v2/accounts", nil, &accounts); err != nil {
		return nil, err
	}
	return &accounts, nil
}

// GetAccountByID returns a coinbase account by its ID
func (c Client) GetAccountByID(accountID string) (*models.Accounts, error) {
	var accounts models.Accounts
	if err := c.Get("v2/accounts/"+accountID, nil, &accounts); err != nil {
		return nil, err
	}
	return &accounts, nil
}

// GetNewAccountAddress return a new address for an account id (one time deposit)
func (c Client) GetNewAccountAddress(accountID string) (*models.Addresses, error) {
	var addrs models.Addresses
	if err := c.Post("v2/accounts/"+accountID+"/addresses", nil, &addrs); err != nil {
		return nil, err
	}
	return &addrs, nil
}

// GetTransactionsByAccountID lists account’s transactions
func (c Client) GetTransactionsByAccountID(accountID string) (*models.Transactions, error) {
	var transact models.Transactions
	if err := c.Get("v2/accounts/"+accountID+"/transactions", nil, &transact); err != nil {
		return nil, err
	}
	return &transact, nil
}
