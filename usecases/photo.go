package usecases

import (
	"mini-project-alterra/models"
	"mini-project-alterra/repositories"
	"time"
)

type PhotoUsecase interface {
	CreatePhoto(input models.PhotoInput) (models.Photo, error)
	DeletePhoto(photoID, userID int) (int, error)
	GetMyPhotoByID(userId, ID int) models.Photo
	GetMyPhoto(userId int) []models.Photo
	GetPhotos() []models.Photo
	UpdatePhoto(input models.PhotoInput, PhotoID, UserID int) (models.Photo, int, error)
}

type photoUsecase struct {
	repository repositories.PhotoRepository
}

func NewPhotoUsecase(repository repositories.PhotoRepository) *photoUsecase {
	return &photoUsecase{repository}
}

// CreatePhoto godoc
// @Summary      Create photo
// @Description  Create photo
// @Tags         Photo
// @Accept       multipart/form-data
// @Produce      json
// @Param        file formData file true "Photo file"
// @Param        title formData string false "Photo title"
// @Param        caption formData string false "Photo caption"
// @Success      200
// @Failure      400
// @Failure      404
// @Failure      500
// @Router       /photos [post]
// @Security BearerAuth
func (ps *photoUsecase) CreatePhoto(input models.PhotoInput) (models.Photo, error) {
	var (
		photo models.Photo
	)

	photo.Title = input.Title
	photo.Caption = input.Caption
	photo.PhotoURL = input.PhotoURL
	photo.UserID = input.UserID

	photo, err := ps.repository.CreatePhoto(photo)

	return photo, err
}

// DeletePhoto godoc
// @Summary      Delete photo
// @Description  Delete photo
// @Tags         Photo
// @Accept       json
// @Produce      json
// @Success      200
// @Failure      400
// @Failure      404
// @Failure      500
// @Param id path int true "Photo ID"
// @Router       /photos/{id} [delete]
// @Security BearerAuth
func (ps *photoUsecase) DeletePhoto(photoID, userID int) (int, error) {
	var photo models.Photo

	photo, err := ps.repository.FindByID(photoID)
	if err != nil {
		return photo.UserID, err
	}

	if uint(photoID) == photo.ID && photo.UserID == userID {
		ps.repository.DeletePhotoRepository(photo)
	} else {
		return photo.UserID, err
	}

	return photo.UserID, nil
}

// GetMyPhoto godoc
// @Summary      Get my photo by id
// @Description  Get my photo by id
// @Tags         Photo
// @Accept       json
// @Produce      json
// @Success      200
// @Failure      400
// @Failure      404
// @Failure      500
// @Param id path int true "Photo ID"
// @Router       /photo/{id} [get]
// @Security BearerAuth
func (ps *photoUsecase) GetMyPhotoByID(userId, ID int) models.Photo {
	var (
		photos models.Photo
	)

	photos, err := ps.repository.GetAllMyPhotoByID(userId, ID)
	if err != nil {
		return photos
	}

	return photos
}

// GetMyPhoto godoc
// @Summary      Get my photo
// @Description  Get my photo
// @Tags         Photo
// @Accept       json
// @Produce      json
// @Success      200
// @Failure      400
// @Failure      404
// @Failure      500
// @Router       /photo [get]
// @Security BearerAuth
func (ps *photoUsecase) GetMyPhoto(userId int) []models.Photo {
	var (
		photos []models.Photo
	)

	photos, err := ps.repository.GetAllMyPhoto(userId)
	if err != nil {
		return nil
	}

	return photos
}

// GetPhotos godoc
// @Summary      Get all photos
// @Description  Get all photos
// @Tags         Photo
// @Accept       json
// @Produce      json
// @Success      200
// @Failure      400
// @Failure      404
// @Failure      500
// @Router       /photos [get]
// @Security BearerAuth
func (ps *photoUsecase) GetPhotos() []models.Photo {
	var (
		photos []models.Photo
	)

	photos, err := ps.repository.GetAllPhoto()
	if err != nil {
		return nil
	}

	return photos
}

// UpdatePhoto godoc
// @Summary      Update photo
// @Description  Update photo
// @Tags         Photo
// @Accept       json
// @Produce      json
// @Success      200
// @Failure      400
// @Failure      404
// @Failure      500
// @Param        request body models.PhotoResponseWithoutPhotoURL true "Payload Body [RAW]"
// @Param id path int true "Photo ID"
// @Router       /photos/{id} [patch]
// @Security BearerAuth
func (ps *photoUsecase) UpdatePhoto(input models.PhotoInput, PhotoID, userID int) (models.Photo, int, error) {
	var (
		photo models.Photo
	)

	photo, err := ps.repository.FindByID(PhotoID)
	if err != nil {
		return photo, photo.UserID, err
	}

	if photo.UserID != userID {
		return photo, photo.UserID, err
	}

	photo.Title = input.Title
	photo.Caption = input.Caption
	photo.PhotoURL = input.PhotoURL
	photo.UpdatedAt = time.Now()

	photo, err = ps.repository.UpdatePhoto(photo)
	if err != nil {
		return photo, photo.UserID, err
	}

	photo, err = ps.repository.FindByID(PhotoID)
	if err != nil {
		return photo, photo.UserID, err
	}

	return photo, photo.UserID, nil
}
