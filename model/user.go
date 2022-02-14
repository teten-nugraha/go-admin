package model

type User struct {
	Id        uint
	FirstName string
	LastName  string
	Email     string `gorm:"unique"`
	Password  string
}
