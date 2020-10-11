package models

type User struct {
	ID int `form:"id"`
	Phone string `form:"phone"`
	Password string `form:"password"`
}
