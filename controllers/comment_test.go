package controllers

import (
	"encoding/json"
	"mini-project-alterra/configs"
	"mini-project-alterra/models"
	"mini-project-alterra/repositories"
	"mini-project-alterra/usecases"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestGetAllMyCommenByID(t *testing.T) {
	var testCases = []struct {
		name       string
		token      string
		expectCode int
		sizeData   int
	}{
		{
			name:       "get credential success",
			token:      "valid_token",
			expectCode: http.StatusOK,
			sizeData:   1,
		},
		{
			name:       "missing token",
			token:      "",
			expectCode: http.StatusUnauthorized,
			sizeData:   0,
		},
		{
			name:       "invalid token",
			token:      "invalid_token",
			expectCode: http.StatusUnauthorized,
			sizeData:   0,
		},
	}

	// setup dependencies
	userRepository := repositories.NewUserRepository(configs.DB)
	userService := usecases.NewUserUsecase(userRepository)

	photoRepository := repositories.NewPhotoRepository(configs.DB)

	commentRepository := repositories.NewCommentRepository(configs.DB)
	commentUsecase := usecases.NewCommentUsecase(commentRepository, photoRepository)
	commentController := NewCommentController(userService, commentUsecase)

	// setup echo
	e := InitEchoTestAPI()

	for _, testCase := range testCases {
		// setup request
		req := httptest.NewRequest(http.MethodGet, "/comment/1", nil)
		req.Header.Set("Authorization", testCase.token)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		// execute controller
		if assert.NoError(t, commentController.GetAllMyCommenByID(c)) {
			// assert response status code
			assert.Equal(t, testCase.expectCode, testCase.expectCode)

			// assert response body
			if testCase.expectCode == http.StatusOK {
				body := rec.Body.String()
				var user models.UserResponse
				err := json.Unmarshal([]byte(body), &user)

				if err != nil {
					assert.Error(t, err, "error")
				}

				assert.Equal(t, testCase.sizeData, testCase.sizeData)
			}
		}
	}
}

func TestGetAllMyComment(t *testing.T) {
	var testCases = []struct {
		name       string
		token      string
		expectCode int
		sizeData   int
	}{
		{
			name:       "get credential success",
			token:      "valid_token",
			expectCode: http.StatusOK,
			sizeData:   1,
		},
		{
			name:       "missing token",
			token:      "",
			expectCode: http.StatusUnauthorized,
			sizeData:   0,
		},
		{
			name:       "invalid token",
			token:      "invalid_token",
			expectCode: http.StatusUnauthorized,
			sizeData:   0,
		},
	}

	// setup dependencies
	userRepository := repositories.NewUserRepository(configs.DB)
	userService := usecases.NewUserUsecase(userRepository)

	photoRepository := repositories.NewPhotoRepository(configs.DB)

	commentRepository := repositories.NewCommentRepository(configs.DB)
	commentUsecase := usecases.NewCommentUsecase(commentRepository, photoRepository)
	commentController := NewCommentController(userService, commentUsecase)

	// setup echo
	e := InitEchoTestAPI()

	for _, testCase := range testCases {
		// setup request
		req := httptest.NewRequest(http.MethodGet, "/photo", nil)
		req.Header.Set("Authorization", testCase.token)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		// execute controller
		if assert.NoError(t, commentController.GetAllMyComment(c)) {
			// assert response status code
			assert.Equal(t, testCase.expectCode, testCase.expectCode)

			// assert response body
			if testCase.expectCode == http.StatusOK {
				body := rec.Body.String()
				var user models.UserResponse
				err := json.Unmarshal([]byte(body), &user)

				if err != nil {
					assert.Error(t, err, "error")
				}

				assert.Equal(t, testCase.sizeData, testCase.sizeData)
			}
		}
	}
}

func TestGetAllComments(t *testing.T) {
	var testCases = []struct {
		name       string
		token      string
		expectCode int
		sizeData   int
	}{
		{
			name:       "get credential success",
			token:      "valid_token",
			expectCode: http.StatusOK,
			sizeData:   1,
		},
		{
			name:       "missing token",
			token:      "",
			expectCode: http.StatusUnauthorized,
			sizeData:   0,
		},
		{
			name:       "invalid token",
			token:      "invalid_token",
			expectCode: http.StatusUnauthorized,
			sizeData:   0,
		},
	}

	// setup dependencies
	userRepository := repositories.NewUserRepository(configs.DB)
	userService := usecases.NewUserUsecase(userRepository)

	photoRepository := repositories.NewPhotoRepository(configs.DB)

	commentRepository := repositories.NewCommentRepository(configs.DB)
	commentUsecase := usecases.NewCommentUsecase(commentRepository, photoRepository)
	commentController := NewCommentController(userService, commentUsecase)

	// setup echo
	e := InitEchoTestAPI()

	for _, testCase := range testCases {
		// setup request
		req := httptest.NewRequest(http.MethodGet, "/comments", nil)
		req.Header.Set("Authorization", testCase.token)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		// execute controller
		if assert.NoError(t, commentController.GetAllComments(c)) {
			// assert response status code
			assert.Equal(t, testCase.expectCode, testCase.expectCode)

			// assert response body
			if testCase.expectCode == http.StatusOK {
				body := rec.Body.String()
				var user models.UserResponse
				err := json.Unmarshal([]byte(body), &user)

				if err != nil {
					assert.Error(t, err, "error")
				}

				assert.Equal(t, testCase.sizeData, testCase.sizeData)
			}
		}
	}
}

func TestCreateComment(t *testing.T) {
	var testCases = []struct {
		name       string
		request    string
		expectCode int
		sizeData   int
	}{
		{
			name: "create comment success",
			request: `{
				"users_id": 1,
				"photos_id": 1,
				"message": "title"
				}`,
			expectCode: http.StatusOK,
			sizeData:   1,
		},
	}

	userRepository := repositories.NewUserRepository(configs.DB)
	userService := usecases.NewUserUsecase(userRepository)

	photoRepository := repositories.NewPhotoRepository(configs.DB)

	commentRepository := repositories.NewCommentRepository(configs.DB)
	commentUsecase := usecases.NewCommentUsecase(commentRepository, photoRepository)
	commentController := NewCommentController(userService, commentUsecase)

	e := InitEchoTestAPI()
	for _, testCase := range testCases {
		req := httptest.NewRequest(http.MethodPost, "/comments", strings.NewReader(testCase.request))
		req.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		if assert.NoError(t, commentController.CreateComment(c)) {
			if testCase.expectCode == http.StatusCreated {
				body := rec.Body.String()
				var user models.User
				err := json.Unmarshal([]byte(body), &user)

				if err != nil {
					assert.Error(t, err, "error")
				}
				assert.Equal(t, testCase.sizeData, len(user.Email))
			}
		}
	}
}

func TestUpdateComment(t *testing.T) {
	userRepository := repositories.NewUserRepository(configs.DB)
	userService := usecases.NewUserUsecase(userRepository)

	photoRepository := repositories.NewPhotoRepository(configs.DB)

	commentRepository := repositories.NewCommentRepository(configs.DB)
	commentUsecase := usecases.NewCommentUsecase(commentRepository, photoRepository)
	commentController := NewCommentController(userService, commentUsecase)

	// Create a new Echo request context
	e := echo.New()
	req := httptest.NewRequest(http.MethodPatch, "/comments/1", strings.NewReader(`{"email":"test@example.com","name":"Test User"}`))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)

	// Call UpdateUser controller
	err := commentController.UpdateComment(c)

	// Assertions
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, http.StatusOK)

}

func TestDeleteComment(t *testing.T) {
	userRepository := repositories.NewUserRepository(configs.DB)
	userService := usecases.NewUserUsecase(userRepository)

	photoRepository := repositories.NewPhotoRepository(configs.DB)

	commentRepository := repositories.NewCommentRepository(configs.DB)
	commentUsecase := usecases.NewCommentUsecase(commentRepository, photoRepository)
	commentController := NewCommentController(userService, commentUsecase)

	e := InitEchoTestAPI()

	var testCases = []struct {
		name       string
		token      string
		userId     string
		expectCode int
	}{
		{
			name:       "delete user success",
			expectCode: http.StatusOK,
		},
		{
			name:       "delete user with invalid token",
			token:      "invalid token",
			expectCode: http.StatusUnauthorized,
		},
		{
			name:       "delete user with invalid user id",
			userId:     "invalid user id",
			expectCode: http.StatusBadRequest,
		},
	}

	for _, testCase := range testCases {
		req := httptest.NewRequest(http.MethodDelete, "/comments/1", nil)
		req.Header.Set("Authorization", "Bearer "+testCase.token)
		q := req.URL.Query()
		q.Add("user_id", testCase.userId)
		req.URL.RawQuery = q.Encode()

		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		if assert.NoError(t, commentController.DeleteComment(c)) {
			assert.Equal(t, testCase.expectCode, testCase.expectCode)
		}
	}
}
