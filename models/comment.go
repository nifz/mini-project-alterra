package models

import (
	"time"

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

type CommentInput struct {
	Message string `json:"message" binding:"required"`
	PhotoID uint   `json:"photos_id"`
	UserID  uint   `json:"users_id"`
}

type ExampleCommentInput struct {
	PhotoID uint   `json:"photos_id"`
	Message string `json:"message" example:"Wah keren!"`
}

type UpdateCommentInput struct {
	Message string `json:"message" example:"Nice picture!"`
}

type CommentResponse struct {
	ID        int            `json:"id"`
	PhotoID   int            `json:"photos_id"`
	Message   string         `json:"message"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	Photo     PhotoResponses `json:"photo"`
	User      UserResponses  `json:"user"`
}

func ParseCommentToResponse(comments Comment) CommentResponse {
	return CommentResponse{
		ID:        int(comments.ID),
		PhotoID:   comments.PhotoID,
		Message:   comments.Message,
		CreatedAt: comments.CreatedAt,
		UpdatedAt: comments.UpdatedAt,
		Photo: PhotoResponses{
			Title:    comments.Photo.Title,
			Caption:  comments.Photo.Caption,
			PhotoURL: comments.Photo.PhotoURL,
		},
		User: UserResponses{
			FullName: comments.User.FullName,
			Username: comments.User.Username,
			Email:    comments.User.Email,
		},
	}
}

func ParseCommentToResponseArray(comments []Comment) []CommentResponse {
	responses := make([]CommentResponse, 0, len(comments))

	for _, s := range comments {
		response := ParseCommentToResponse(s)
		responses = append(responses, response)
	}

	return responses
}
