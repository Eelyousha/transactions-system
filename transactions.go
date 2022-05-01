package main

import (
	"github.com/go-pg/pg/v10"
)

func processRequest(db *pg.DB) error {
	request := new(transaction)
	if err := db.Model(request).Order("created_at ASC").Limit(1).Select(); err == nil {
		fromUserData := new(user)
		if err := db.Model(fromUserData).Where("nickname = ?", request.From_user).Select(); err == nil {
			if fromUserData.Balance >= request.Amount {
				toUserData := new(user)
				if err := db.Model(toUserData).Where("nickname = ?", request.To_user).Select(); err == nil {

					fromUserData.Balance -= request.Amount
					toUserData.Balance += request.Amount
					db.Model(fromUserData).Set("balance = ?", fromUserData.Balance).Where("nickname = ?", fromUserData.Nickname).Update()
					db.Model(toUserData).Set("balance = ?", toUserData.Balance).Where("nickname = ?", toUserData.Nickname).Update()
				} else {
					return err
				}
			} else {
			}
		} else {
			return err
		}
	} else {
		return err
	}

	db.Model(request).Where("id = ?", request.TransactionID).Delete()

	return nil
}
