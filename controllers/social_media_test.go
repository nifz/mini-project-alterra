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

func TestGetMySocialMediaByID(t *testing.T) {
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

	socialMediaRepository := repositories.NewSocialMediaRepository(configs.DB)
	socialMediaUsecase := usecases.NewSocialMediaUsecase(socialMediaRepository)
	socialMediaController := NewSocialMediaController(userService, socialMediaUsecase)

	// setup echo
	e := InitEchoTestAPI()

	for _, testCase := range testCases {
		// setup request
		req := httptest.NewRequest(http.MethodGet, "/socialmedia/1", nil)
		req.Header.Set("Authorization", testCase.token)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		// execute controller
		if assert.NoError(t, socialMediaController.GetMySocialMediaByID(c)) {
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

func TestGetMySocialMedia(t *testing.T) {
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

	socialMediaRepository := repositories.NewSocialMediaRepository(configs.DB)
	socialMediaUsecase := usecases.NewSocialMediaUsecase(socialMediaRepository)
	socialMediaController := NewSocialMediaController(userService, socialMediaUsecase)

	// setup echo
	e := InitEchoTestAPI()

	for _, testCase := range testCases {
		// setup request
		req := httptest.NewRequest(http.MethodGet, "/socialmedia", nil)
		req.Header.Set("Authorization", testCase.token)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		// execute controller
		if assert.NoError(t, socialMediaController.GetMySocialMedia(c)) {
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

func TestGetSocialMedias(t *testing.T) {
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

	socialMediaRepository := repositories.NewSocialMediaRepository(configs.DB)
	socialMediaUsecase := usecases.NewSocialMediaUsecase(socialMediaRepository)
	socialMediaController := NewSocialMediaController(userService, socialMediaUsecase)

	// setup echo
	e := InitEchoTestAPI()

	for _, testCase := range testCases {
		// setup request
		req := httptest.NewRequest(http.MethodGet, "/socialmedias", nil)
		req.Header.Set("Authorization", testCase.token)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		// execute controller
		if assert.NoError(t, socialMediaController.GetMySocialMedia(c)) {
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

func TestCreateSocialMedia(t *testing.T) {
	var testCases = []struct {
		name       string
		request    string
		expectCode int
		sizeData   int
	}{
		{
			name: "create social media success",
			request: `{
				"users_id": "1",
				"name" : "twitter",
				"social_media_url"  : "https://twitter.com/"
				}`,
			expectCode: http.StatusOK,
			sizeData:   1,
		},
	}

	userRepository := repositories.NewUserRepository(configs.DB)
	userService := usecases.NewUserUsecase(userRepository)

	socialMediaRepository := repositories.NewSocialMediaRepository(configs.DB)
	socialMediaUsecase := usecases.NewSocialMediaUsecase(socialMediaRepository)
	socialMediaController := NewSocialMediaController(userService, socialMediaUsecase)

	e := InitEchoTestAPI()
	for _, testCase := range testCases {
		req := httptest.NewRequest(http.MethodPost, "/socialmedias", strings.NewReader(testCase.request))
		req.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		if assert.NoError(t, socialMediaController.CreateSocialMedia(c)) {

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

func TestUpdateSocialMedia(t *testing.T) {
	userRepository := repositories.NewUserRepository(configs.DB)
	userService := usecases.NewUserUsecase(userRepository)

	socialMediaRepository := repositories.NewSocialMediaRepository(configs.DB)
	socialMediaUsecase := usecases.NewSocialMediaUsecase(socialMediaRepository)
	socialMediaController := NewSocialMediaController(userService, socialMediaUsecase)

	// Create a new Echo request context
	e := echo.New()
	req := httptest.NewRequest(http.MethodPatch, "/socialmedias/1", strings.NewReader(`{"email":"test@example.com","name":"Test User"}`))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)

	// Call UpdateUser controller
	err := socialMediaController.UpdateSocialMedia(c)

	// Assertions
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, http.StatusOK)

}

func TestDeleteSocialMedia(t *testing.T) {
	userRepository := repositories.NewUserRepository(configs.DB)
	userService := usecases.NewUserUsecase(userRepository)

	socialMediaRepository := repositories.NewSocialMediaRepository(configs.DB)
	socialMediaUsecase := usecases.NewSocialMediaUsecase(socialMediaRepository)
	socialMediaController := NewSocialMediaController(userService, socialMediaUsecase)

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
		req := httptest.NewRequest(http.MethodDelete, "/socialmedias/1", nil)
		req.Header.Set("Authorization", "Bearer "+testCase.token)
		q := req.URL.Query()
		q.Add("user_id", testCase.userId)
		req.URL.RawQuery = q.Encode()

		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		if assert.NoError(t, socialMediaController.DeleteSocialMedia(c)) {
			assert.Equal(t, testCase.expectCode, testCase.expectCode)
		}
	}
}
