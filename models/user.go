package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	FullName string `json:"full_name"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
