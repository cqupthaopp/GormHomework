package GormUtils

type User struct {
	Username string `gorm:"primary_key"`
	Phone    string
	Avator   string
	Qq       string
}

type UserAccount struct {
	Username string `gorm:"primary_key"`
	Balance  float64
}
