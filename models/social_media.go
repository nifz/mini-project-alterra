package models

import (
	"gorm.io/gorm"
)

type SocialMedia struct {
	gorm.Model
	Name           string `json:"name"`
	SocialMediaURL string `json:"social_media_url"`
	UserID         int    `json:"users_id"`
	User           User   `json:"user" gorm:"primaryKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"user"`
}
