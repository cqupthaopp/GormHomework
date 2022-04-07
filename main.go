package main

import (
	"fmt"
	"gorm/GormUtils"
)

func main() {

	GormUtils.AddUser(GormUtils.User{Username: "Hao_pp", Qq: "144678538", Avator: "test.png", Phone: "1737832669"})

	fmt.Println(GormUtils.FindUser("Hao_pp"))

}
