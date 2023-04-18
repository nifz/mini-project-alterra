package models

import (
	"gorm.io/gorm"
)

type Photo struct {
	gorm.Model
	Title    string `json:"title"`
	Caption  string `json:"caption"`
	PhotoURL string `json:"photo_url"`
	UserID   int    `json:"users_id"`
	User     User   `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"user"`
}
