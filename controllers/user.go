package controllers

import (
	"mini-project-alterra/middlewares"
	"mini-project-alterra/models"
	"mini-project-alterra/usecases"
	"net/http"

	"github.com/labstack/echo/v4"
)

type UserController struct {
	userUsecase usecases.UserUsecase
}

func NewUserController(userUsecase usecases.UserUsecase) UserController {
	return UserController{userUsecase}
}

func (controllers *UserController) SignIn(c echo.Context) error {
	var loginInput models.LoginInput

	err := c.Bind(&loginInput)
	if err != nil {
		return c.JSON(
			http.StatusBadRequest, echo.Map{
				"message": err.Error(),
			})
	}

	tokenString, err := controllers.userUsecase.Login(loginInput)
	if err != nil {
		return c.JSON(
			http.StatusBadRequest, echo.Map{
				"message": err.Error(),
			})
	}

	return c.JSON(
		http.StatusOK, echo.Map{
			"message": "Login successfully",
			"data": echo.Map{
				"token": tokenString,
			},
		})
}

func (controller *UserController) SignUp(c echo.Context) error {
	var registerInput models.RegisterInput

	err := c.Bind(&registerInput)
	if err != nil {
		return c.JSON(
			http.StatusBadRequest, echo.Map{
				"message": err.Error(),
			})
	}

	user, err := controller.userUsecase.Register(registerInput)
	if err != nil {
		return c.JSON(
			http.StatusBadRequest, echo.Map{
				"message": err.Error(),
			})
	}

	return c.JSON(
		http.StatusCreated, echo.Map{
			"message": "Successfully registered",
			"data":    models.ParseUserToResponse(user),
		})
}

func (controller *UserController) GetCredential(c echo.Context) error {
	tokenString := middlewares.GetTokenFromHeader(c.Request())

	if tokenString == "" {
		return c.JSON(http.StatusUnauthorized, echo.Map{
			"message": "No token provided",
		})
	}

	userId, err := middlewares.GetUserIdFromToken(tokenString)

	if err != nil {
		return c.JSON(http.StatusUnauthorized, echo.Map{
			"message": err.Error(),
		})
	}

	user, err := controller.userUsecase.GetCredential(userId)
	if err != nil {
		return c.JSON(
			http.StatusBadRequest, echo.Map{
				"message": err.Error(),
			})
	}

	return c.JSON(
		http.StatusOK, echo.Map{
			"message": "Successfully retrieved",
			"data":    models.ParseUserToResponse(user),
		})
}

func (controller *UserController) UpdateUser(c echo.Context) error {
	tokenString := middlewares.GetTokenFromHeader(c.Request())

	if tokenString == "" {
		return c.JSON(http.StatusUnauthorized, echo.Map{
			"message": "No token provided",
		})
	}

	userId, err := middlewares.GetUserIdFromToken(tokenString)

	if err != nil {
		return c.JSON(http.StatusUnauthorized, echo.Map{
			"message": err.Error(),
		})
	}

	var input models.User

	err = c.Bind(&input)
	if err != nil {
		return c.JSON(
			http.StatusBadRequest, echo.Map{
				"message": "Failed to update user",
			})
	}

	user, err := controller.userUsecase.UpdateUser(userId, input)
	if err != nil {
		return c.JSON(
			http.StatusBadRequest, echo.Map{
				"message": err.Error(),
			})
	}

	return c.JSON(
		http.StatusOK, echo.Map{
			"message": "Successfully updated user",
			"data":    models.ParseUserToResponse(user),
		})
}

func (controller *UserController) DeleteUser(c echo.Context) error {
	tokenString := middlewares.GetTokenFromHeader(c.Request())

	if tokenString == "" {
		return c.JSON(http.StatusUnauthorized, echo.Map{
			"message": "No token provided",
		})
	}

	userId, err := middlewares.GetUserIdFromToken(tokenString)

	if err != nil {
		return c.JSON(http.StatusUnauthorized, echo.Map{
			"message": err.Error(),
		})
	}

	err = controller.userUsecase.DeleteUser(userId)
	if err != nil {
		return c.JSON(
			http.StatusBadRequest, echo.Map{
				"message": "Failed to delete user",
			})
	}

	return c.JSON(
		http.StatusOK, echo.Map{
			"message": "Successfully deleted user",
		})
}
