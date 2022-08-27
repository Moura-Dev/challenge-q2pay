package repository

import (
	"challenge-q2pay/db"
	"challenge-q2pay/models"
	"fmt"
	"github.com/shopspring/decimal"
)

func CreateWallet(wallet *models.Wallet) error {
	_, err := db.Conn.NamedExec(`INSERT INTO wallets (user_id, balance, available, in_transaction, version) 
		VALUES (:user_id, :balance, :available, :in_transaction, :version)`, wallet)
	if err != nil {
		fmt.Println(err)
		return err
	}
	db.Conn.MustBegin().Commit()
	return nil
}

// Get wallet by user_id
func GetWallet(user_id string) (*models.Wallet, error) {
	wallet := models.Wallet{}
	err := db.Conn.Get(&wallet, "SELECT * FROM wallets WHERE user_id = $1", user_id)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return &wallet, nil
}

// ADD balance user
func AddBalance(user_id string, balance decimal.Decimal) error {
	_, err := db.Conn.Exec(`UPDATE wallets SET balance = balance + $1, version = version + 1 
               WHERE user_id = $2 AND version=version`, balance, user_id)
	if err != nil {
		fmt.Println(err)
		return err
	}
	_ = db.Conn.MustBegin().Commit()
	return nil
}
