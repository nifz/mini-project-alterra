package repositories

import (
	"mini-project-alterra/models"

	"gorm.io/gorm"
)

type SocialMediaRepository interface {
	GetMySocialMediaByID(userId, ID int) (models.SocialMedia, error)
	GetMySocialMedia(userId int) ([]models.SocialMedia, error)
	GetSocialMedias(userId int) ([]models.SocialMedia, error)
	GetSocialMediasByUser(userID int) ([]models.SocialMedia, error)
	GetSocialMediaByID(socialMediaId int, userId int) (models.SocialMedia, error)
	CreateSocialMedia(socialmedia models.SocialMedia) (models.SocialMedia, error)
	UpdateSocialMedia(socialmedia models.SocialMedia) (models.SocialMedia, error)
	DeleteSocialMedia(socialmedia models.SocialMedia) error
}

type socialMediaRepository struct {
	DB *gorm.DB
}

func NewSocialMediaRepository(db *gorm.DB) *socialMediaRepository {
	return &socialMediaRepository{db}
}

func (r *socialMediaRepository) GetMySocialMediaByID(userId, ID int) (models.SocialMedia, error) {
	var socialMedia models.SocialMedia
	err := r.DB.Where("user_id = ? AND id = ?", userId, ID).Preload("User").Find(&socialMedia).Error
	return socialMedia, err
}

func (r *socialMediaRepository) GetMySocialMedia(userId int) ([]models.SocialMedia, error) {
	var socialMedia []models.SocialMedia
	err := r.DB.Where("user_id = ?", userId).Preload("User").Find(&socialMedia).Error
	return socialMedia, err
}

func (r *socialMediaRepository) GetSocialMedias(userId int) ([]models.SocialMedia, error) {
	var socialMedia []models.SocialMedia
	err := r.DB.Joins("JOIN users ON users.id = social_media.user_id").Where("social_media.deleted_at IS NULL AND users.deleted_at IS NULL").Preload("User").Find(&socialMedia).Error
	return socialMedia, err
}

func (r *socialMediaRepository) GetSocialMediasByUser(userId int) ([]models.SocialMedia, error) {
	var socialMedia []models.SocialMedia
	err := r.DB.Where("users_id = ?", userId).Find(&socialMedia).Error
	return socialMedia, err
}

func (r *socialMediaRepository) GetSocialMediaByID(socialMediaId int, userId int) (models.SocialMedia, error) {
	var socialMedia models.SocialMedia
	err := r.DB.Where("id = ?", socialMediaId).Where("user_id = ?", userId).First(&socialMedia).Error
	return socialMedia, err
}

func (r *socialMediaRepository) CreateSocialMedia(socialmedia models.SocialMedia) (models.SocialMedia, error) {
	err := r.DB.Create(&socialmedia).Preload("User").Where("id = ?", socialmedia.ID).First(&socialmedia).Error
	return socialmedia, err
}

func (r *socialMediaRepository) UpdateSocialMedia(socialmedia models.SocialMedia) (models.SocialMedia, error) {
	err := r.DB.Session(&gorm.Session{FullSaveAssociations: true}).Updates(&socialmedia).Preload("User").Where("id = ?", socialmedia.ID).First(&socialmedia).Error
	return socialmedia, err
}

func (r *socialMediaRepository) DeleteSocialMedia(socialmedia models.SocialMedia) error {
	err := r.DB.Delete(&socialmedia).Error
	return err
}
