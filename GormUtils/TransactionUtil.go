package GormUtils

import "errors"

func DelMoneyInUserInfo(username string, delNum float64) error {

	if _, err := FindUser(username); err != nil {
		return errors.New("No user")
	}

	db := Connect(UsersAcountTable)

	tx := db.Begin()

	var userAccount UserAccount

	if err := tx.Where("username = ?", username).Find(&userAccount); err != nil {
		tx.Rollback()
		return errors.New("No user")
	}

	if userAccount.Balance < delNum {
		tx.Rollback()
		return errors.New("Balance is not enough")
	}

	if err := tx.Where("username = ?", username).Update("balance", userAccount.Balance-delNum); err != nil {
		tx.Rollback()
		return err.Error
	}

	return nil

}
