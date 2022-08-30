package models

import "github.com/shopspring/decimal"

type Transaction struct {
	Id        int             `json:"-" db:"id"`
	Value     decimal.Decimal `json:"value" db:"value"`
	Payer     int             `json:"payer" db:"payer"`
	Payee     int             `json:"payee" db:"payee"`
	CreatedAt string          `json:"-" db:"created_at"`
}

type Deposit struct {
	Value decimal.Decimal `json:"value" db:"value"`
}
