package controllers

import (
	"mini-project-alterra/middlewares"
	"mini-project-alterra/models"
	"mini-project-alterra/usecases"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type CommentController struct {
	userUsecase    usecases.UserUsecase
	commentUsecase usecases.CommentUsecase
}

func NewCommentController(userUsecase usecases.UserUsecase, commentUsecase usecases.CommentUsecase) CommentController {
	return CommentController{userUsecase, commentUsecase}
}

func (cc *CommentController) CreateComment(c echo.Context) error {
	tokenString := middlewares.GetTokenFromHeader(c.Request())

	if tokenString == "" {
		return c.JSON(http.StatusUnauthorized, echo.Map{
			"message": "No token provided",
		})
	}

	userID, err := middlewares.GetUserIdFromToken(tokenString)
	checkUser, err := cc.userUsecase.GetCredential(userID)

	if checkUser.ID == 0 || err != nil {
		return c.JSON(http.StatusUnauthorized, echo.Map{
			"message": err.Error(),
		})
	}

	var commentInput models.CommentInput

	err = c.Bind(&commentInput)
	if err != nil {
		return c.JSON(
			http.StatusBadRequest, echo.Map{
				"message": "Failed to input comment",
			})
	}

	commentInput.UserID = uint(userID)

	comment, err := cc.commentUsecase.PostComment(commentInput)
	if err != nil {
		return c.JSON(
			http.StatusBadRequest, echo.Map{
				"message": "Failed to input comment",
			})
	}

	return c.JSON(
		http.StatusOK, echo.Map{
			"message": "Successfully submitted comment",
			"data":    models.ParseCommentToResponse(comment),
		})
}

func (cc *CommentController) GetAllMyCommenByID(c echo.Context) error {

	tokenString := middlewares.GetTokenFromHeader(c.Request())

	if tokenString == "" {
		return c.JSON(http.StatusUnauthorized, echo.Map{
			"message": "No token provided",
		})
	}

	userID, err := middlewares.GetUserIdFromToken(tokenString)
	checkUser, err := cc.userUsecase.GetCredential(userID)

	if checkUser.ID == 0 || err != nil {
		return c.JSON(http.StatusUnauthorized, echo.Map{
			"message": err.Error(),
		})
	}

	byID, err := strconv.Atoi(c.Param("id"))
	comments := cc.commentUsecase.GetMyCommentByID(userID, byID)

	if comments.ID < 1 {
		return c.JSON(
			http.StatusBadRequest, echo.Map{
				"message": "Failed to get data, id is invalid",
			})
	}
	return c.JSON(
		http.StatusOK, echo.Map{
			"message": "Successfully retrieved all my comments",
			"data":    models.ParseCommentToResponse(comments),
		})
}

func (cc *CommentController) GetAllMyComment(c echo.Context) error {

	tokenString := middlewares.GetTokenFromHeader(c.Request())

	if tokenString == "" {
		return c.JSON(http.StatusUnauthorized, echo.Map{
			"message": "No token provided",
		})
	}

	userID, err := middlewares.GetUserIdFromToken(tokenString)
	checkUser, err := cc.userUsecase.GetCredential(userID)

	if checkUser.ID == 0 || err != nil {
		return c.JSON(http.StatusUnauthorized, echo.Map{
			"message": err.Error(),
		})
	}

	comments := cc.commentUsecase.GetMyComment(userID)

	return c.JSON(
		http.StatusOK, echo.Map{
			"message": "Successfully retrieved all my comments",
			"data":    models.ParseCommentToResponseArray(comments),
		})
}

func (cc *CommentController) GetAllComments(c echo.Context) error {
	tokenString := middlewares.GetTokenFromHeader(c.Request())

	if tokenString == "" {
		return c.JSON(http.StatusUnauthorized, echo.Map{
			"message": "No token provided",
		})
	}

	userId, err := middlewares.GetUserIdFromToken(tokenString)
	checkUser, err := cc.userUsecase.GetCredential(userId)

	if checkUser.ID == 0 || err != nil {
		return c.JSON(http.StatusUnauthorized, echo.Map{
			"message": err.Error(),
		})
	}
	comments := cc.commentUsecase.GetComments()

	return c.JSON(
		http.StatusOK, echo.Map{
			"message": "Successfully retrieved all comments",
			"data":    models.ParseCommentToResponseArray(comments),
		})
}

func (cc *CommentController) DeleteComment(c echo.Context) error {
	tokenString := middlewares.GetTokenFromHeader(c.Request())

	if tokenString == "" {
		return c.JSON(http.StatusUnauthorized, echo.Map{
			"message": "No token provided",
		})
	}

	userID, err := middlewares.GetUserIdFromToken(tokenString)
	checkUser, err := cc.userUsecase.GetCredential(userID)

	if checkUser.ID == 0 || err != nil {
		return c.JSON(http.StatusUnauthorized, echo.Map{
			"message": err.Error(),
		})
	}

	commentIDRaw := c.Param("id")

	commentID, err := strconv.Atoi(commentIDRaw)
	if err != nil {
		return c.JSON(
			http.StatusBadRequest, echo.Map{
				"message": "Parameter must be a valid ID",
			})
	}

	validID, err := cc.commentUsecase.DeleteComment(commentID, userID)
	if err != nil {
		return c.JSON(
			http.StatusBadRequest, echo.Map{
				"message": "Comment ID cannot be found",
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
			"message": "Successfully deleted comment",
		})
}

func (cc *CommentController) UpdateComment(c echo.Context) error {
	tokenString := middlewares.GetTokenFromHeader(c.Request())

	if tokenString == "" {
		return c.JSON(http.StatusUnauthorized, echo.Map{
			"message": "No token provided",
		})
	}

	userID, err := middlewares.GetUserIdFromToken(tokenString)
	checkUser, err := cc.userUsecase.GetCredential(userID)

	if checkUser.ID == 0 || err != nil {
		return c.JSON(http.StatusUnauthorized, echo.Map{
			"message": err.Error(),
		})
	}

	var commentInput models.UpdateCommentInput
	commentIDRaw := c.Param("id")

	commentID, err := strconv.Atoi(commentIDRaw)
	if err != nil {
		return c.JSON(
			http.StatusBadRequest, echo.Map{
				"message": "Parameter must be a valid ID",
			})
	}

	err = c.Bind(&commentInput)
	if err != nil {
		return c.JSON(
			http.StatusBadRequest, echo.Map{
				"message": "Cannot input the data",
			})
	}

	comment, validID, err := cc.commentUsecase.UpdateComment(commentInput, commentID, userID)
	if err != nil {
		return c.JSON(
			http.StatusBadRequest, echo.Map{
				"message": "Comment ID cannot be found",
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
			"message": "Successfully updated comment",
			"data":    models.ParseCommentToResponse(comment),
		})

}
