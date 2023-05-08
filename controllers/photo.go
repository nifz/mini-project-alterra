package controllers

import (
	"mini-project-alterra/middlewares"
	"mini-project-alterra/models"
	"mini-project-alterra/usecases"
	"net/http"
	"regexp"
	"strconv"

	"github.com/labstack/echo/v4"
)

type PhotoController struct {
	userUsecase  usecases.UserUsecase
	photoUsecase usecases.PhotoUsecase
}

func NewPhotoController(userUsecase usecases.UserUsecase, photoUsecase usecases.PhotoUsecase) PhotoController {
	return PhotoController{userUsecase, photoUsecase}
}

func (pc *PhotoController) CreatePhoto(c echo.Context) error {
	tokenString := middlewares.GetTokenFromHeader(c.Request())

	if tokenString == "" {
		return c.JSON(http.StatusUnauthorized, echo.Map{
			"message": "No token provided",
		})
	}

	userID, err := middlewares.GetUserIdFromToken(tokenString)
	checkUser, err := pc.userUsecase.GetCredential(userID)

	if checkUser.ID == 0 || err != nil {
		return c.JSON(http.StatusUnauthorized, echo.Map{
			"message": err.Error(),
		})
	}

	var photoInput models.PhotoInput

	err = c.Bind(&photoInput)
	if err != nil {
		return c.JSON(
			http.StatusBadRequest, echo.Map{
				"message": "Failed to upload photo",
			})
	}

	photoInput.UserID = userID

	if photoInput.PhotoURL == "" {
		formHeader, err := c.FormFile("file")
		if err != nil {
			if err != nil {
				return c.JSON(
					http.StatusInternalServerError,
					echo.Map{
						"message": "Error uploading photo",
					})
			}
		}

		//get file from header
		formFile, err := formHeader.Open()
		if err != nil {
			return c.JSON(
				http.StatusInternalServerError,
				echo.Map{
					"message": err.Error(),
				})
		}

		var re = regexp.MustCompile(`.png|.jpeg|.jpg`)

		if !re.MatchString(formHeader.Filename) {
			return c.JSON(
				http.StatusBadRequest, echo.Map{
					"message": "The provided file format is not allowed. Please upload a JPEG or PNG image",
				})
		}

		uploadUrl, err := usecases.NewMediaUpload().FileUpload(models.File{File: formFile})

		if err != nil {
			return c.JSON(
				http.StatusInternalServerError,
				echo.Map{
					"message": err.Error(),
				})
		}
		photoInput.PhotoURL = uploadUrl
	}

	var url models.Url

	url.Url = photoInput.PhotoURL

	var re = regexp.MustCompile(`.png|.jpeg|.jpg`)

	if !re.MatchString(photoInput.PhotoURL) {
		return c.JSON(
			http.StatusBadRequest, echo.Map{
				"message": "The provided file format is not allowed. Please upload a JPEG or PNG image",
			})
	}

	uploadUrl, err := usecases.NewMediaUpload().RemoteUpload(url)

	if uploadUrl == "" || err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{
				"message": "Error uploading photo",
			})
	}

	photoInput.PhotoURL = uploadUrl

	photo, err := pc.photoUsecase.CreatePhoto(photoInput)
	if err != nil {
		return c.JSON(
			http.StatusBadRequest, echo.Map{
				"message": "Failed to upload photo",
			})
	}

	return c.JSON(
		http.StatusCreated, echo.Map{
			"message": "Successfully uploaded photo",
			"data":    models.ParsePhotoToResponse(photo),
		})
}

func (pc *PhotoController) DeletePhoto(c echo.Context) error {
	tokenString := middlewares.GetTokenFromHeader(c.Request())

	if tokenString == "" {
		return c.JSON(http.StatusUnauthorized, echo.Map{
			"message": "No token provided",
		})
	}

	userID, err := middlewares.GetUserIdFromToken(tokenString)
	checkUser, err := pc.userUsecase.GetCredential(userID)

	if checkUser.ID == 0 || err != nil {
		return c.JSON(http.StatusUnauthorized, echo.Map{
			"message": err.Error(),
		})
	}

	PhotoIDRaw := c.Param("id")

	PhotoID, err := strconv.Atoi(PhotoIDRaw)
	if err != nil {
		return c.JSON(
			http.StatusBadRequest,
			"Parameter must be a valid ID",
		)
	}

	validID, err := pc.photoUsecase.DeletePhoto(PhotoID, userID)
	if err != nil {
		return c.JSON(
			http.StatusBadRequest, echo.Map{
				"message": "Photo ID cannot be found",
			})
	}

	if validID != userID {
		return c.JSON(
			http.StatusBadRequest, echo.Map{
				"message": "Invalid User ID",
			})
	}

	return c.JSON(
		http.StatusOK, echo.Map{
			"message": "Successfully deleted photo",
		})
}

func (pc *PhotoController) GetMyPhotoByID(c echo.Context) error {
	tokenString := middlewares.GetTokenFromHeader(c.Request())

	if tokenString == "" {
		return c.JSON(http.StatusUnauthorized, echo.Map{
			"message": "No token provided",
		})
	}

	userId, err := middlewares.GetUserIdFromToken(tokenString)
	checkUser, err := pc.userUsecase.GetCredential(userId)

	if checkUser.ID == 0 || err != nil {
		return c.JSON(http.StatusUnauthorized, echo.Map{
			"message": err.Error(),
		})
	}
	byID, err := strconv.Atoi(c.Param("id"))
	photos := pc.photoUsecase.GetMyPhotoByID(userId, byID)

	if photos.ID < 1 {
		return c.JSON(
			http.StatusBadRequest, echo.Map{
				"message": "Failed to get data, id is invalid",
			})
	}
	return c.JSON(
		http.StatusOK, echo.Map{
			"message": "Successfully retrieved all my photos",
			"data":    models.ParsePhotoToResponse(photos),
		})
}

func (pc *PhotoController) GetMyPhoto(c echo.Context) error {
	tokenString := middlewares.GetTokenFromHeader(c.Request())

	if tokenString == "" {
		return c.JSON(http.StatusUnauthorized, echo.Map{
			"message": "No token provided",
		})
	}

	userId, err := middlewares.GetUserIdFromToken(tokenString)
	checkUser, err := pc.userUsecase.GetCredential(userId)

	if checkUser.ID == 0 || err != nil {
		return c.JSON(http.StatusUnauthorized, echo.Map{
			"message": err.Error(),
		})
	}
	photos := pc.photoUsecase.GetMyPhoto(userId)

	return c.JSON(
		http.StatusOK, echo.Map{
			"message": "Successfully retrieved all my photos",
			"data":    models.ParsePhotoToResponseArray(photos),
		})
}

func (pc *PhotoController) GetPhotos(c echo.Context) error {
	tokenString := middlewares.GetTokenFromHeader(c.Request())

	if tokenString == "" {
		return c.JSON(http.StatusUnauthorized, echo.Map{
			"message": "No token provided",
		})
	}

	userId, err := middlewares.GetUserIdFromToken(tokenString)
	checkUser, err := pc.userUsecase.GetCredential(userId)

	if checkUser.ID == 0 || err != nil {
		return c.JSON(http.StatusUnauthorized, echo.Map{
			"message": err.Error(),
		})
	}
	photos := pc.photoUsecase.GetPhotos()

	return c.JSON(
		http.StatusOK, echo.Map{
			"message": "Successfully retrieved all photos",
			"data":    models.ParsePhotoToResponseArray(photos),
		})
}

func (pc *PhotoController) UpdatePhoto(c echo.Context) error {

	tokenString := middlewares.GetTokenFromHeader(c.Request())

	if tokenString == "" {
		return c.JSON(http.StatusUnauthorized, echo.Map{
			"message": "No token provided",
		})
	}

	userID, err := middlewares.GetUserIdFromToken(tokenString)
	checkUser, err := pc.userUsecase.GetCredential(userID)

	if checkUser.ID == 0 || err != nil {
		return c.JSON(http.StatusUnauthorized, echo.Map{
			"message": err.Error(),
		})
	}

	var photoInput models.PhotoInput
	photoIDRaw := c.Param("id")

	photoID, err := strconv.Atoi(photoIDRaw)
	if err != nil {
		return c.JSON(
			http.StatusBadRequest, echo.Map{
				"message": "Parameter must be a valid ID",
			})
	}

	err = c.Bind(&photoInput)
	if err != nil {
		return c.JSON(
			http.StatusBadRequest, echo.Map{
				"message": "Cannot input the data",
			})
	}

	if len(photoInput.PhotoURL) > 1 {
		return c.JSON(
			http.StatusBadRequest, echo.Map{
				"message": "Cannot update the photo",
			})
	}

	photo, validID, err := pc.photoUsecase.UpdatePhoto(photoInput, photoID, userID)
	if err != nil {
		return c.JSON(
			http.StatusBadRequest, echo.Map{
				"message": "Photo ID cannot be found",
			})
	}

	if validID != userID {
		return c.JSON(
			http.StatusBadRequest, echo.Map{
				"message": "Invalid User ID",
			})
	}

	return c.JSON(
		http.StatusOK, echo.Map{
			"message": "Successfully updated photo",
			"data":    models.ParsePhotoToResponse(photo),
		})
}
