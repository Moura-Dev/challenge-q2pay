package models

import "github.com/shopspring/decimal"

type Transaction struct {
	Id        int             `json:"-" db:"id"`
	Value     decimal.Decimal `json:"value" db:"value"`
	Payer     int             `json:"payer" db:"payer"`
	Payee     int             `json:"payee" db:"payee"`
	Status    string          `json:"status" db:"status"`
	CreatedAt string          `json:"created_at" db:"created_at"`
	UpdatedAt string          `json:"updated_at" db:"updated_at"`
}

type Status struct {
	Open     string `json:"open"`
	Done     string `json:"done"`
	Canceled string `json:"canceled"`
}
