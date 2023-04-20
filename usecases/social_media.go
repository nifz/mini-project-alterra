package usecases

import (
	"mini-project-alterra/models"
	"mini-project-alterra/repositories"
)

type SocialMediaUsecase interface {
	GetMySocialMediaByID(userId, ID int) (models.SocialMedia, error)
	GetMySocialMedia(userId int) ([]models.SocialMedia, error)
	GetSocialMedias(userId int) ([]models.SocialMedia, error)
	CreateSocialMedia(userId int, input models.SocialMediaInput) (models.SocialMedia, error)
	UpdateSocialMedia(socialMediaId int, userId int, input models.SocialMediaUpdateInput) (models.SocialMedia, error)
	DeleteSocialMedia(socialMediaId int, userId int) error
}

type socialMediaUsecase struct {
	repository repositories.SocialMediaRepository
}

func NewSocialMediaUsecase(repository repositories.SocialMediaRepository) *socialMediaUsecase {
	return &socialMediaUsecase{repository}
}

// GetMySocialMediaByID godoc
// @Summary      Get my social media by id
// @Description  Get my social media by id
// @Tags         Social Media
// @Accept       json
// @Produce      json
// @Success      200
// @Failure      400
// @Failure      404
// @Failure      500
// @Param id path int true "Social Media ID"
// @Router       /socialmedia/{id} [get]
// @Security BearerAuth
func (s *socialMediaUsecase) GetMySocialMediaByID(userId, ID int) (models.SocialMedia, error) {
	socialMedias, err := s.repository.GetMySocialMediaByID(userId, ID)
	return socialMedias, err
}

// GetMySocialMedia godoc
// @Summary      Get my social media
// @Description  Get my social media
// @Tags         Social Media
// @Accept       json
// @Produce      json
// @Success      200
// @Failure      400
// @Failure      404
// @Failure      500
// @Router       /socialmedia [get]
// @Security BearerAuth
func (s *socialMediaUsecase) GetMySocialMedia(userId int) ([]models.SocialMedia, error) {
	socialMedias, err := s.repository.GetMySocialMedia(userId)
	return socialMedias, err
}

// GetSocialMedias godoc
// @Summary      Get all social media
// @Description  Get all social media
// @Tags         Social Media
// @Accept       json
// @Produce      json
// @Success      200
// @Failure      400
// @Failure      404
// @Failure      500
// @Router       /socialmedias [get]
// @Security BearerAuth
func (s *socialMediaUsecase) GetSocialMedias(userId int) ([]models.SocialMedia, error) {
	socialMedias, err := s.repository.GetSocialMedias(userId)
	return socialMedias, err
}

// CreateSocialMedia godoc
// @Summary      Create social media
// @Description  Create social media
// @Tags         Social Media
// @Accept       json
// @Produce      json
// @Param        request body models.SocialMediaInput true "Payload Body [RAW]"
// @Success      200
// @Failure      400
// @Failure      404
// @Failure      500
// @Router       /socialmedias [post]
// @Security BearerAuth
func (s *socialMediaUsecase) CreateSocialMedia(userId int, input models.SocialMediaInput) (models.SocialMedia, error) {
	var socialMedia models.SocialMedia
	socialMedia.UserID = userId
	socialMedia.Name = input.Name
	socialMedia.SocialMediaURL = input.SocialMediaURL
	socialMedia, err := s.repository.CreateSocialMedia(socialMedia)
	return socialMedia, err
}

// UpdateSocialMedia godoc
// @Summary      Update social media
// @Description  Update social media
// @Tags         Social Media
// @Accept       json
// @Produce      json
// @Success      200
// @Failure      400
// @Failure      404
// @Failure      500
// @Param        request body models.SocialMediaUpdateInput true "Payload Body [RAW]"
// @Param id path int true "Social Media ID"
// @Router       /socialmedias/{id} [patch]
// @Security BearerAuth
func (s *socialMediaUsecase) UpdateSocialMedia(socialMediaId int, userId int, input models.SocialMediaUpdateInput) (models.SocialMedia, error) {
	socialMedia, err := s.repository.GetSocialMediaByID(socialMediaId, userId)
	if err != nil {
		return socialMedia, err
	}
	if input.Name != "" {
		socialMedia.Name = input.Name
	}
	if input.SocialMediaURL != "" {
		socialMedia.SocialMediaURL = input.SocialMediaURL
	}
	socialMedia, err = s.repository.UpdateSocialMedia(socialMedia)
	return socialMedia, err
}

// DeleteSocialMedia godoc
// @Summary      Delete social media
// @Description  Delete social media
// @Tags         Social Media
// @Accept       json
// @Produce      json
// @Success      200
// @Failure      400
// @Failure      404
// @Failure      500
// @Param id path int true "Social Media ID"
// @Router       /socialmedias/{id} [delete]
// @Security BearerAuth
func (s *socialMediaUsecase) DeleteSocialMedia(socialMediaId int, userId int) error {
	socialMedia, err := s.repository.GetSocialMediaByID(socialMediaId, userId)
	if err != nil {
		return err
	}
	err = s.repository.DeleteSocialMedia(socialMedia)
	return err
}
