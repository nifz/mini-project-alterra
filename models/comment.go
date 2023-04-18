package models

import (
	"gorm.io/gorm"
)

type Comment struct {
	gorm.Model
	PhotoID int    `json:"photos_id"`
	Photo   Photo  `gorm:"foreignKey:PhotoID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"photo"`
	UserID  int    `json:"users_id"`
	User    User   `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"user"`
	Message string `json:"message"`
}
