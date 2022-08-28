package repository

import (
	"challenge-q2pay/db"
	"challenge-q2pay/models"
	"fmt"
)

func CreateUser(user *models.User) error {
	lastInsertID, err := db.Conn.PrepareNamed(`INSERT INTO users (full_name, cpfcnpj, email, 
                   login, password)
		VALUES (:full_name, :cpfcnpj, :email, :login, :password) RETURNING id`)
	if err != nil {
		fmt.Println(err)
		return err
	}
	var id int
	err = lastInsertID.Get(&id, user)
	//Check lastInsertID
	if err != nil {
		fmt.Println(err)
		return err
	}
	wallet := models.Wallet{
		UserID: id,
	}

	err = CreateWallet(&wallet)
	if err != nil {
		fmt.Println(err)
		return err
	}

	_ = db.Conn.MustBegin().Commit()
	return nil
}

////Get User By id

func GetUser(id int) (*models.User, error) {
	user := models.User{}
	err := db.Conn.Get(&user, "SELECT * FROM users WHERE id = $1", id)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return &user, nil
}
