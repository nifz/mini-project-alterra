package repositories

import (
	"mini-project-alterra/models"

	"gorm.io/gorm"
)

type CommentRepository interface {
	CreateComment(comment models.Comment) (models.Comment, error)
	FindByID(commentID int) (models.Comment, error)
	GetMyComment(userId int) ([]models.Comment, error)
	GetMyCommentByID(userId, ID int) (models.Comment, error)
	GetAllComments() ([]models.Comment, error)
	DeleteCommentRepository(comment models.Comment) error
	UpdateComment(comment models.Comment) (models.Comment, error)
	GetPhotoData(commentID int) (models.Comment, error)
}

type commentRepository struct {
	DB *gorm.DB
}

func NewCommentRepository(db *gorm.DB) *commentRepository {
	return &commentRepository{db}
}

func (cr *commentRepository) CreateComment(comment models.Comment) (models.Comment, error) {
	err := cr.DB.Create(&comment).Error

	return comment, err
}

func (cr *commentRepository) FindByID(commentID int) (models.Comment, error) {
	var comment models.Comment

	err := cr.DB.Where("id = ?", commentID).Preload("Photo").Preload("User").First(&comment).Error

	return comment, err
}

func (cr *commentRepository) GetPhotoData(commentID int) (models.Comment, error) {
	comment, err := cr.FindByID(commentID)

	err = cr.DB.Preload("Photo").Preload("User").Find(&comment).Error

	return comment, err
}

func (cr *commentRepository) GetMyCommentByID(userId, ID int) (models.Comment, error) {
	var comments models.Comment

	err := cr.DB.Where("user_id = ? AND id = ?", userId, ID).Preload("User").Preload("Photo").Find(&comments).Error

	return comments, err
}

func (cr *commentRepository) GetMyComment(userId int) ([]models.Comment, error) {
	var comments []models.Comment

	err := cr.DB.Where("user_id = ?", userId).Preload("User").Preload("Photo").Find(&comments).Error

	return comments, err
}

func (cr *commentRepository) GetAllComments() ([]models.Comment, error) {
	var comments []models.Comment
	err := cr.DB.Joins("JOIN users ON users.id = comments.user_id").Where("comments.deleted_at IS NULL AND users.deleted_at IS NULL").Preload("User").Preload("Photo").Find(&comments).Error
	return comments, err
}

func (cr *commentRepository) DeleteCommentRepository(comment models.Comment) error {
	err := cr.DB.Delete(&comment).Error
	return err
}

func (cr *commentRepository) UpdateComment(comment models.Comment) (models.Comment, error) {
	err := cr.DB.Session(&gorm.Session{FullSaveAssociations: true}).Preload("User").Updates(&comment).Error

	return comment, err
}
