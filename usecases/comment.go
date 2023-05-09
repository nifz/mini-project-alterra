package usecases

import (
	"errors"
	"mini-project-alterra/models"
	"mini-project-alterra/repositories"
	"time"
)

type CommentUsecase interface {
	PostComment(input models.CommentInput) (models.Comment, error)
	GetMyComment(userId int) []models.Comment
	GetMyCommentByID(userId, ID int) models.Comment
	GetComments() []models.Comment
	DeleteComment(commentID, userID int) (int, error)
	UpdateComment(input models.UpdateCommentInput, commentID, userID int) (models.Comment, int, error)
}

type commentUsecase struct {
	repository      repositories.CommentRepository
	photoRepository repositories.PhotoRepository
}

func NewCommentUsecase(repository repositories.CommentRepository, photoRepository repositories.PhotoRepository) *commentUsecase {
	return &commentUsecase{repository, photoRepository}
}

// PostComment godoc
// @Summary      Post comment
// @Description  Post comment
// @Tags         Comment
// @Accept       json
// @Produce      json
// @Param        request body models.ExampleCommentInput true "Payload Body [RAW]"
// @Success      200
// @Failure      400
// @Failure      404
// @Failure      500
// @Router       /comments [post]
// @Security BearerAuth
func (cs *commentUsecase) PostComment(input models.CommentInput) (models.Comment, error) {
	var (
		comment models.Comment
	)

	if input.Message == "" {
		return comment, errors.New("Failed to input comment")
	}

	photos, err := cs.photoRepository.FindByID(int(input.PhotoID))
	if err != nil {
		return comment, err
	}

	comment.Message = input.Message
	comment.PhotoID = int(photos.ID)
	comment.UserID = int(input.UserID)

	comment, err = cs.repository.CreateComment(comment)

	if err != nil {
		return comment, err
	}
	comment, err = cs.repository.FindByID(int(comment.ID))

	if err != nil {
		return comment, err
	}

	return comment, nil
}

// GetComments godoc
// @Summary      Get my comment by id
// @Description  Get my comment by id
// @Tags         Comment
// @Accept       json
// @Produce      json
// @Success      200
// @Failure      400
// @Failure      404
// @Failure      500
// @Param id path int true "Comment ID"
// @Router       /comment/{id} [get]
// @Security BearerAuth
func (cs *commentUsecase) GetMyCommentByID(userId, ID int) models.Comment {
	var (
		comments models.Comment
	)

	comments, err := cs.repository.GetMyCommentByID(userId, ID)
	if err != nil {
		return comments
	}
	return comments
}

// GetComments godoc
// @Summary      Get my comment
// @Description  Get my comment
// @Tags         Comment
// @Accept       json
// @Produce      json
// @Success      200
// @Failure      400
// @Failure      404
// @Failure      500
// @Router       /comment [get]
// @Security BearerAuth
func (cs *commentUsecase) GetMyComment(userId int) []models.Comment {
	var (
		comments []models.Comment
	)

	comments, err := cs.repository.GetMyComment(userId)
	if err != nil {
		return nil
	}
	return comments
}

// GetComments godoc
// @Summary      Get all comments
// @Description  Get all comments
// @Tags         Comment
// @Accept       json
// @Produce      json
// @Success      200
// @Failure      400
// @Failure      404
// @Failure      500
// @Router       /comments [get]
// @Security BearerAuth
func (cs *commentUsecase) GetComments() []models.Comment {
	var (
		comments []models.Comment
	)

	comments, err := cs.repository.GetAllComments()
	if err != nil {
		return nil
	}
	return comments
}

// DeleteComment godoc
// @Summary      Delete comment
// @Description  Delete comment
// @Tags         Comment
// @Accept       json
// @Produce      json
// @Success      200
// @Failure      400
// @Failure      404
// @Failure      500
// @Param id path int true "Comment ID"
// @Router       /comments/{id} [delete]
// @Security BearerAuth
func (cs *commentUsecase) DeleteComment(commentID int, userID int) (int, error) {
	var comment models.Comment

	comment, err := cs.repository.FindByID(commentID)
	if err != nil {
		return comment.UserID, err
	}

	if uint(commentID) == comment.ID && comment.UserID == userID {
		cs.repository.DeleteCommentRepository(comment)
	} else {
		return comment.UserID, err
	}

	return comment.UserID, nil
}

// UpdateComment godoc
// @Summary      Update comment
// @Description  Update comment
// @Tags         Comment
// @Accept       json
// @Produce      json
// @Success      200
// @Failure      400
// @Failure      404
// @Failure      500
// @Param        request body models.UpdateCommentInput true "Payload Body [RAW]"
// @Param id path int true "Photo ID"
// @Router       /comments/{id} [patch]
// @Security BearerAuth
func (cs *commentUsecase) UpdateComment(input models.UpdateCommentInput, commentID, UserID int) (models.Comment, int, error) {
	var (
		comment models.Comment
	)

	comment, err := cs.repository.GetPhotoData(commentID)
	if err != nil {
		return comment, comment.UserID, err
	}

	if comment.UserID != UserID {
		return comment, comment.UserID, err
	}

	comment.Message = input.Message
	comment.UpdatedAt = time.Now()

	comment, err = cs.repository.UpdateComment(comment)
	if err != nil {
		return comment, comment.UserID, err
	}

	comment, err = cs.repository.GetPhotoData(commentID)

	if err != nil {
		return comment, comment.UserID, err
	}
	return comment, comment.UserID, nil
}
