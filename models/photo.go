package models

import (
	"time"

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

type PhotoInput struct {
	Title    string `form:"title" json:"title" binding:"required"`
	Caption  string `form:"caption" json:"caption" binding:"required"`
	PhotoURL string `form:"file" json:"file,omitempty" validate:"required" binding:"required"`
	UserID   int    `json:"user_id"`
}

type PhotoResponse struct {
	ID        int           `json:"id"`
	Title     string        `json:"title"`
	Caption   string        `json:"caption"`
	PhotoURL  string        `json:"photo_url"`
	CreatedAt time.Time     `json:"created_at"`
	UpdatedAt time.Time     `json:"updated_at"`
	User      UserResponses `json:"user"`
}

type PhotoResponseWithoutPhotoURL struct {
	Title   string `json:"title" example:"Bali"`
	Caption string `json:"caption" example:"Liburan ke bali"`
}

type PhotoResponses struct {
	Title    string `json:"title" example:"Bali"`
	Caption  string `json:"caption" example:"Liburan ke bali"`
	PhotoURL string `json:"photo_url" example:"https://www.water-sport-bali.com/wp-content/uploads/2012/04/Tips-Wisata-Bali-Twitter-1.jpg"`
}

func ParsePhotoToResponse(photos Photo) PhotoResponse {
	return PhotoResponse{
		ID:        int(photos.ID),
		Title:     photos.Title,
		Caption:   photos.Caption,
		PhotoURL:  photos.PhotoURL,
		CreatedAt: photos.CreatedAt,
		UpdatedAt: photos.UpdatedAt,
		User: UserResponses{
			FullName: photos.User.FullName,
			Username: photos.User.Username,
			Email:    photos.User.Email,
		},
	}
}

func ParsePhotoToResponseArray(photos []Photo) []PhotoResponse {
	responses := make([]PhotoResponse, 0, len(photos))

	for _, s := range photos {
		response := ParsePhotoToResponse(s)
		responses = append(responses, response)
	}

	return responses
}
