package repositories

import (
	"mini-project-alterra/models"

	"gorm.io/gorm"
)

type PhotoRepository interface {
	CreatePhoto(photo models.Photo) (models.Photo, error)
	DeletePhotoRepository(photo models.Photo) error
	FindByID(photoID int) (models.Photo, error)
	GetAllMyPhoto(userId int) ([]models.Photo, error)
	GetAllMyPhotoByID(userId, ID int) (models.Photo, error)
	GetAllPhoto() ([]models.Photo, error)
	UpdatePhoto(photo models.Photo) (models.Photo, error)
}

type photoRepository struct {
	DB *gorm.DB
}

func NewPhotoRepository(db *gorm.DB) *photoRepository {
	return &photoRepository{db}
}

func (pr *photoRepository) CreatePhoto(photo models.Photo) (models.Photo, error) {
	err := pr.DB.Create(&photo).Preload("User").First(&photo).Error

	return photo, err
}

func (pr *photoRepository) DeletePhotoRepository(photo models.Photo) error {
	err := pr.DB.Delete(&photo).Error
	return err
}

func (pr *photoRepository) FindByID(photoID int) (models.Photo, error) {
	var photo models.Photo

	err := pr.DB.Where("id = ?", photoID).Preload("User").First(&photo).Error

	return photo, err
}

func (pr *photoRepository) GetAllMyPhotoByID(userId, ID int) (models.Photo, error) {
	var photos models.Photo

	err := pr.DB.Where("user_id = ? AND id = ?", userId, ID).Preload("User").Find(&photos).Error

	return photos, err
}

func (pr *photoRepository) GetAllMyPhoto(userId int) ([]models.Photo, error) {
	var photos []models.Photo

	err := pr.DB.Where("user_id = ?", userId).Preload("User").Find(&photos).Error

	return photos, err
}

func (pr *photoRepository) GetAllPhoto() ([]models.Photo, error) {
	var photos []models.Photo

	err := pr.DB.Joins("JOIN users ON users.id = photos.user_id").Where("photos.deleted_at IS NULL AND users.deleted_at IS NULL").Preload("User").Find(&photos).Error

	return photos, err
}

func (pr *photoRepository) UpdatePhoto(photo models.Photo) (models.Photo, error) {
	err := pr.DB.Session(&gorm.Session{FullSaveAssociations: true}).Updates(&photo).Error

	return photo, err

}
