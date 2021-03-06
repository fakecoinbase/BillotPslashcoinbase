package models

import (
	"strconv"
)

// Balance is a coinbase balance model
type Balance struct {
	Amount   string `json:"amount"`
	Currency string `json:"currency"`
}

// Currency is the coinbase model a currency
type Currency struct {
	Code         string `json:"code"`
	Name         string `json:"name"`
	Colors       string `json:"color"`
	SortIndex    int64  `json:"sort_index"`
	Exponent     int64  `json:"exponent"`
	Type         string `json:"type"`
	AddressRegex string `json:"address_regex"`
	AssetID      string `json:"asset_id"`
	Slug         string `json:"slug"`
}

// Account is a coinbase account model
type Account struct {
	ID               string `json:"id"`
	Name             string `json:"name"`
	Primary          bool   `json:"primary"`
	Type             string `json:"type"`
	Currency         `json:"currency"`
	Balance          `json:"balance"`
	CreatedAt        string  `json:"created_at"`
	UpdatedAt        string  `json:"updated_at"`
	Resource         string  `json:"resource"`
	ResourcePath     string  `json:"resource_path"`
	AllowDeposit     bool    `json:"allow_deposit"`
	AllowWithdrawals bool    `json:"allow_withdrawals"`
	NativeBalance    Balance `json:"native_balance"`
}

// Network is a crypto network entity model
type Network struct {
	Status string `json:"status"`
	Name   string `json:"name"`
}

// Ressource is a transaction ressource model
type Ressource struct {
	ID            string `json:"id"`
	Ressource     string `json:"ressource"`
	RessourcePath string `json:"ressource_path"`
}

// Detail is a transaction detail model
type Detail struct {
	Title    string `json:"title"`
	Subtitle string `json:"subtitle"`
}

// Transaction is the transaction element struct
type Transaction struct {
	ID            string    `json:"id"`
	Type          string    `json:"type"`
	Status        string    `json:"status"`
	Amount        Balance   `json:"amount"`
	NativeAmount  Balance   `json:"native_amount"`
	Description   *string   `json:"description"`
	CreatedAt     string    `json:"created_at"`
	UpdatedAt     string    `json:"updated_at"`
	Ressource     string    `json:"ressource"`
	RessourcePath string    `json:"ressource_path"`
	Buy           Ressource `json:"buy,omitempty"`
	To            Ressource `json:"to,omitempty"`
	Network       Network   `json:"network,omitempty"`
	Details       Detail    `json:"details"`
}

// Pagination is the pagination struct for a Coinbase API response
type Pagination struct {
	EndingBefore         string `json:"ending_before"`
	StartingAfter        string `json:"starting_after"`
	PreviousEndingBefore string `json:"previous_ending_before"`
	NextStartingAfter    string `json:"next_starting_after"`
	Limit                int64  `json:"limit"`
	Order                string `json:"order"`
	PreviousURI          string `json:"previous_uri"`
	NextURI              string `json:"next_uri"`
}

// Response is the global response object for sucessfull Coinbase api call
type Response struct {
	Pagination `json:"pagination"`
}

// SpotPrice is the spot price request model
type SpotPrice struct {
	Data Balance `json:"data"`
}

// Accounts is the response model for GET `accounts` data
type Accounts struct {
	Response
	Datas []Account `json:"data"`
}

// Get an account by asset
func (a Accounts) Get(asset string) *Account {
	for i := range a.Datas {
		if a.Datas[i].Balance.Currency == asset {
			return &a.Datas[i]
		}
	}
	return nil
}

// Address is the model for new account deposit address
type Address struct {
	ID            string `json:"id"`
	Address       string `json:"address"`
	Name          string `json:"name"`
	CreatedAt     string `json:"created_at"`
	UpdatedAt     string `json:"updated_at"`
	Network       string `json:"network"`
	Ressource     string `json:"ressource"`
	RessourcePath string `json:"ressource_path"`
}

// Addresses is the response model for POST `addresses` data
type Addresses struct {
	Data Address `json:"data"`
}

// Transactions is the response mode for transaction data
type Transactions struct {
	Response
	Datas []Transaction `json:"data"`
}

// FilterEmpty remove the empty balances from Accounts.Datas object
func (a *Accounts) FilterEmpty() {
	var fltr []Account
	var empty float64
	for i := range a.Datas {
		v, err := strconv.ParseFloat(a.Datas[i].Amount, 64)
		if v > empty && err == nil {
			fltr = append(fltr, a.Datas[i])
		}
	}
	a.Datas = fltr
}
