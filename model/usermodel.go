package model

type Student struct {
	ID        int    `json:"id" gorm:"primaryKey"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Email     string `json:"email" gorm:"unique"`
	Username  string `json:"username" gorm:"unique"`
	Password  string `json:"password"`
}
