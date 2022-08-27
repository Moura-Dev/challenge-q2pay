package models

import "github.com/shopspring/decimal"

type Wallet struct {
	Id            int             `json:"-" db:"id"`
	Balance       decimal.Decimal `json:"balance" db:"balance"`
	Available     decimal.Decimal `json:"available" db:"available"`
	InTransaction decimal.Decimal `json:"in_transaction" db:"in_transaction"`
	UserID        int             `json:"user_id" db:"user_id"`
	Version       int             `json:"version" db:"version"`
}
