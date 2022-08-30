package repository

import (
	"challenge-q2pay/db"
	"challenge-q2pay/models"
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/shopspring/decimal"
)

func CreateWallet(wallet *models.Wallet) error {
	_, err := db.Conn.NamedExec(`INSERT INTO wallets (user_id, balance, version) 
		VALUES (:user_id, :balance, :version)`, wallet)
	if err != nil {
		fmt.Println(err)
		return err
	}
	db.Conn.MustBegin().Commit()
	return nil
}

// Get wallet by user_id
func GetWallet(userId int) (*models.Wallet, error) {
	wallet := models.Wallet{}
	err := db.Conn.Get(&wallet, "SELECT * FROM wallets WHERE user_id = $1", userId)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return &wallet, nil
}

// ADD balance user
func AddBalance(userId int, balance decimal.Decimal, tx *sqlx.Tx) error {
	_, err := tx.Exec(`UPDATE wallets SET balance = balance + $1, version = version + 1 
               WHERE user_id = $2 AND version=version`, balance, userId)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

// Remove balance from payer
func RemoveBalance(userId int, balance decimal.Decimal, tx *sqlx.Tx) error {
	_, err := tx.Exec(`UPDATE wallets SET balance = balance - $1, version = version + 1 
			   WHERE user_id = $2 AND version=version`, balance, userId)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}
