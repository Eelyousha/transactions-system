package db

import (
	"github.com/go-pg/pg/v10"
)

func ProcessRequest(db *pg.DB) error {
	request := new(Transaction)
	if err := db.Model(request).Order("created_at ASC").Limit(1).Select(); err == nil {
		db.Model(request).Where("id = ?", request.TransactionID).Delete()
		fromUserData := new(User)
		if err := db.Model(fromUserData).Where("nickname = ?", request.From_user).Select(); err == nil {
			if fromUserData.Balance >= request.Amount {
				toUserData := new(User)
				if err := db.Model(toUserData).Where("nickname = ?", request.To_user).Select(); err == nil {
					if toUserData.Nickname != fromUserData.Nickname {
						fromUserData.Balance -= request.Amount
						toUserData.Balance += request.Amount
						db.Model(toUserData).Set("balance = ?", toUserData.Balance).Where("nickname = ?", toUserData.Nickname).Update()
						db.Model(fromUserData).Set("balance = ?", fromUserData.Balance).Where("nickname = ?", fromUserData.Nickname).Update()
					}
				} else {
					return err
				}
			} else {
			}
		} else {
			return err
		}
	} else {
		db.Model(request).Where("id = ?", request.TransactionID).Delete()
		return err
	}

	return nil
}
