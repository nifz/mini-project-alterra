package controllers

import (
	"mini-project-alterra/middlewares"
	"mini-project-alterra/models"
	"mini-project-alterra/usecases"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type SocialMediaController struct {
	userUsecase        usecases.UserUsecase
	socialMediaUsecase usecases.SocialMediaUsecase
}

func NewSocialMediaController(userUsecase usecases.UserUsecase, socialMediaUsecase usecases.SocialMediaUsecase) SocialMediaController {
	return SocialMediaController{userUsecase, socialMediaUsecase}
}

func (controller *SocialMediaController) GetMySocialMediaByID(c echo.Context) error {
	tokenString := middlewares.GetTokenFromHeader(c.Request())

	if tokenString == "" {
		return c.JSON(http.StatusUnauthorized, echo.Map{
			"message": "No token provided",
		})
	}

	userId, err := middlewares.GetUserIdFromToken(tokenString)
	checkUser, err := controller.userUsecase.GetCredential(userId)

	if checkUser.ID == 0 || err != nil {
		return c.JSON(http.StatusUnauthorized, echo.Map{
			"message": err.Error(),
		})
	}

	byID, err := strconv.Atoi(c.Param("id"))

	socialMedias, err := controller.socialMediaUsecase.GetMySocialMediaByID(userId, byID)
	if err != nil {
		return c.JSON(
			http.StatusBadRequest, echo.Map{
				"message": "Failed to get my social media",
			})
	}

	if socialMedias.ID < 1 {
		return c.JSON(
			http.StatusBadRequest, echo.Map{
				"message": "Failed to get data, id is invalid",
			})
	}

	return c.JSON(
		http.StatusOK, echo.Map{
			"message": "Successfully get my social media",
			"data":    models.ParseSocialMediaToResponse(socialMedias),
		})
}

func (controller *SocialMediaController) GetMySocialMedia(c echo.Context) error {
	tokenString := middlewares.GetTokenFromHeader(c.Request())

	if tokenString == "" {
		return c.JSON(http.StatusUnauthorized, echo.Map{
			"message": "No token provided",
		})
	}

	userId, err := middlewares.GetUserIdFromToken(tokenString)
	checkUser, err := controller.userUsecase.GetCredential(userId)
	if checkUser.ID == 0 || err != nil {
		return c.JSON(http.StatusUnauthorized, echo.Map{
			"message": err.Error(),
		})
	}

	socialMedias, err := controller.socialMediaUsecase.GetMySocialMedia(userId)
	if err != nil {
		return c.JSON(
			http.StatusBadRequest, echo.Map{
				"message": "Failed to get my social media",
			})
	}

	return c.JSON(
		http.StatusOK, echo.Map{
			"message": "Successfully get my social media",
			"data":    models.ParseSocialMediaToResponseArray(socialMedias),
		})
}

func (controller *SocialMediaController) GetSocialMedias(c echo.Context) error {
	tokenString := middlewares.GetTokenFromHeader(c.Request())

	if tokenString == "" {
		return c.JSON(http.StatusUnauthorized, echo.Map{
			"message": "No token provided",
		})
	}

	userId, err := middlewares.GetUserIdFromToken(tokenString)
	checkUser, err := controller.userUsecase.GetCredential(userId)

	if checkUser.ID == 0 || err != nil {
		return c.JSON(http.StatusUnauthorized, echo.Map{
			"message": err.Error(),
		})
	}

	socialMedias, err := controller.socialMediaUsecase.GetSocialMedias(userId)
	if err != nil {
		return c.JSON(
			http.StatusBadRequest, echo.Map{
				"message": "Failed to get social media",
			})
	}

	return c.JSON(
		http.StatusOK, echo.Map{
			"message": "Successfully get social media",
			"data":    models.ParseSocialMediaToResponseArray(socialMedias),
		})
}

func (controller *SocialMediaController) CreateSocialMedia(c echo.Context) error {
	tokenString := middlewares.GetTokenFromHeader(c.Request())

	if tokenString == "" {
		return c.JSON(http.StatusUnauthorized, echo.Map{
			"message": "No token provided",
		})
	}

	userId, err := middlewares.GetUserIdFromToken(tokenString)
	checkUser, err := controller.userUsecase.GetCredential(userId)

	if checkUser.ID == 0 || err != nil {
		return c.JSON(http.StatusUnauthorized, echo.Map{
			"message": err.Error(),
		})
	}

	var input models.SocialMediaInput

	err = c.Bind(&input)
	if err != nil {
		return c.JSON(
			http.StatusBadRequest, echo.Map{
				"message": "Failed to create social medias",
			})
	}

	socialMedia, err := controller.socialMediaUsecase.CreateSocialMedia(userId, input)
	if err != nil {
		return c.JSON(
			http.StatusBadRequest, echo.Map{
				"message": "Failed to create social medias",
			})
	}

	return c.JSON(
		http.StatusCreated, echo.Map{
			"message": "Successfully created social media",
			"data":    models.ParseSocialMediaToResponse(socialMedia),
		})
}

func (controller *SocialMediaController) UpdateSocialMedia(c echo.Context) error {
	tokenString := middlewares.GetTokenFromHeader(c.Request())

	if tokenString == "" {
		return c.JSON(http.StatusUnauthorized, echo.Map{
			"message": "No token provided",
		})
	}

	userId, err := middlewares.GetUserIdFromToken(tokenString)
	checkUser, err := controller.userUsecase.GetCredential(userId)

	if checkUser.ID == 0 || err != nil {
		return c.JSON(http.StatusUnauthorized, echo.Map{
			"message": err.Error(),
		})
	}

	var input models.SocialMediaUpdateInput

	socialMediaIdString := c.Param("socialMediaId")
	if socialMediaIdString == "" {
		return c.JSON(
			http.StatusBadRequest, echo.Map{
				"message": "Failed to update social media",
			})
	}

	socialMediaId, err := strconv.Atoi(socialMediaIdString)
	if err != nil {
		return c.JSON(
			http.StatusBadRequest, echo.Map{
				"message": "Failed to update social media",
			})
	}

	err = c.Bind(&input)
	if err != nil {
		return c.JSON(
			http.StatusBadRequest, echo.Map{
				"message": "Failed to update social media",
			})
	}

	socialMedia, err := controller.socialMediaUsecase.UpdateSocialMedia(socialMediaId, userId, input)
	if err != nil {
		return c.JSON(
			http.StatusBadRequest, echo.Map{
				"message": "Failed to update social media, id is invalid",
			})
	}

	return c.JSON(
		http.StatusOK, echo.Map{
			"message": "Successfully updated social media",
			"data":    models.ParseSocialMediaToResponse(socialMedia),
		})
}

func (controller *SocialMediaController) DeleteSocialMedia(c echo.Context) error {
	tokenString := middlewares.GetTokenFromHeader(c.Request())

	if tokenString == "" {
		return c.JSON(http.StatusUnauthorized, echo.Map{
			"message": "No token provided",
		})
	}

	userId, err := middlewares.GetUserIdFromToken(tokenString)
	checkUser, err := controller.userUsecase.GetCredential(userId)

	if checkUser.ID == 0 || err != nil {
		return c.JSON(http.StatusUnauthorized, echo.Map{
			"message": err.Error(),
		})
	}

	socialMediaIdString := c.Param("socialMediaId")
	if socialMediaIdString == "" {
		return c.JSON(
			http.StatusBadRequest, echo.Map{
				"message": "Failed to update social media",
			})
	}

	socialMediaId, err := strconv.Atoi(socialMediaIdString)
	if err != nil {
		return c.JSON(
			http.StatusBadRequest, echo.Map{
				"message": "Failed to update social media",
			})
	}

	err = controller.socialMediaUsecase.DeleteSocialMedia(socialMediaId, userId)
	if err != nil {
		return c.JSON(
			http.StatusBadRequest, echo.Map{
				"message": "Failed to delete social media, id is invalid",
			})
	}

	return c.JSON(
		http.StatusOK, echo.Map{
			"message": "Successfully deleted social media",
		})
}
