package GormUtils

import (
	"errors"
	"fmt"
)

func AddUser(user User) {

	db := Connect(UsersInfoTable)

	defer db.Close()

	//userInfo := User{username: "Hao_pp", qq : "1446785380", avator: "test.png" , phone: "17378326692"}

	if err := db.Create(user); err != nil {
		fmt.Println(err)
	}

}

func FindUser(username string) (User, error) {

	db := Connect(UsersInfoTable)

	defer db.Close()

	var user User

	if err := db.Where("username = ?", username).Find(&user); err != nil {
		fmt.Println(err)
		return user, errors.New("No user called " + username)
	}

	return user, nil

}

func ChangeUserInfoC(userInfo User) error {

	if _, err := FindUser(userInfo.Username); err != nil {
		return errors.New("No user called " + userInfo.Username)
	}

	db := Connect(UsersInfoTable)

	db.Save(&userInfo)

	return nil

}
