package GormUtils

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

func Connect(tableName string) *gorm.DB {

	dsn := username + ":" + password + "@tcp(" + address + ":" + port + ")/" + dataBase + "?charset=utf8mb4&parseTime=True&loc=Local"

	DataBase, err := gorm.Open("mysql", dsn)

	if err != nil {
		fmt.Println(err)
	}

	DataBase.SingularTable(true)
	DataBase = DataBase.Table(tableName)

	return DataBase

}
