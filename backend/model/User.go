package model

type User struct {
	ID        int64  `json:"-" form:"ID" gorm:"primaryKey;autoIncrement;column:ID"`
	Username  string `json:"username" gorm:"primaryKey;column:username"`
	Password  string `json:"password" gorm:"column:password"`
	Password2 string `json:"password2" gorm:"-"`
	Email     string `json:"email" gorm:"column:email"`
}
