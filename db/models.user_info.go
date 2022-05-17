package db

import (
	"time"

	"github.com/go-pg/pg/v10"
)

type Transaction struct {
	TransactionID int       `pg:"id"`
	From_user     string    `pg:"from_user"`
	To_user       string    `pg:"to_user"`
	Amount        float64   `pg:"amount"`
	CreatedAt     time.Time `pg:"default:now()"`
}

type User struct {
	UserID       int     `pg:"id"`
	Nickname     string  `pg:"nickname"`
	Balance      float64 `pg:"balance"`
	Transactions []Transaction
}

func GetUserFromDB(db *pg.DB, nickname string) (*User, error) {
	userData := new(User)
	err := db.Model(userData).Where("nickname=?", nickname).Select()

	if err != nil {
		return userData, err
	}

	userData, err = getUserQueries(db, userData)
	if err != nil {
		return userData, err
	}

	return userData, err
}

func getUserQueries(db *pg.DB, userData *User) (*User, error) {
	err := db.Model(&userData.Transactions).
		Where("from_user = ?", userData.Nickname).Select()
	// _, err := db.Query(userData.transactions,
	// 	`SELECT * FROM transactions
	// 	WHERE from_user = ?`, userData.Nickname)
	if err != nil {
		return userData, err
	}

	return userData, err
}
