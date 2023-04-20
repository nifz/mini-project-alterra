package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	FullName string `json:"full_name"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginInput struct {
	Email    string `form:"email" json:"email" binding:"required,email" example:"me@hanifz.com"`
	Password string `form:"password" json:"password" binding:"required" example:"qweqwe123"`
}

type RegisterInput struct {
	FullName string `form:"full_name" json:"full_name" binding:"required,min=3" example:"mochammad hanif"`
	Username string `form:"username" json:"username" binding:"required" example:"hanif"`
	Email    string `form:"email" json:"email" binding:"required,email" example:"me@hanifz.com"`
	Password string `form:"password" json:"password" binding:"required,min=6" example:"qweqwe123"`
}

type UserResponse struct {
	FullName  string    `json:"full_name"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UserResponses struct {
	FullName string `json:"full_name"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

func ParseUserToResponse(user User) UserResponse {
	return UserResponse{
		FullName:  user.FullName,
		Username:  user.Username,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}
