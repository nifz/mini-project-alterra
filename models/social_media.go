package models

import (
	"time"

	"gorm.io/gorm"
)

type SocialMedia struct {
	gorm.Model
	Name           string `json:"name"`
	SocialMediaURL string `json:"social_media_url"`
	UserID         int    `json:"users_id"`
	User           User   `json:"user" gorm:"primaryKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"user"`
}

type SocialMediaInput struct {
	Name           string `form:"name" json:"name" binding:"required" example:"github"`
	SocialMediaURL string `form:"social_media_url" json:"social_media_url" binding:"required" example:"https://github.com/nifz"`
}

type SocialMediaUpdateInput struct {
	Name           string `form:"name" json:"name" example:"linkedin"`
	SocialMediaURL string `form:"social_media_url" json:"social_media_url" example:"https://linkedin.com/in/hanifz"`
}

type SocialMediaResponse struct {
	ID             int           `json:"id"`
	Name           string        `json:"name"`
	SocialMediaURL string        `json:"social_media_url"`
	CreatedAt      time.Time     `json:"created_at"`
	UpdatedAt      time.Time     `json:"updated_at"`
	User           UserResponses `json:"user"`
}

func ParseSocialMediaToResponse(socialMedias SocialMedia) SocialMediaResponse {
	return SocialMediaResponse{
		ID:             int(socialMedias.ID),
		Name:           socialMedias.Name,
		SocialMediaURL: socialMedias.SocialMediaURL,
		CreatedAt:      socialMedias.CreatedAt,
		UpdatedAt:      socialMedias.UpdatedAt,
		User: UserResponses{
			FullName: socialMedias.User.FullName,
			Username: socialMedias.User.Username,
			Email:    socialMedias.User.Email,
		},
	}
}

func ParseSocialMediaToResponseArray(socialMedias []SocialMedia) []SocialMediaResponse {
	responses := make([]SocialMediaResponse, 0, len(socialMedias))

	for _, s := range socialMedias {
		response := ParseSocialMediaToResponse(s)
		responses = append(responses, response)
	}

	return responses
}
