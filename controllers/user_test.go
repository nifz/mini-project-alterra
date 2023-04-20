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

func InitEchoTestAPI() *echo.Echo {
	configs.InitDBTest()
	e := echo.New()
	return e
}

func TestLoginUser(t *testing.T) {
	var testCases = []struct {
		name       string
		request    string
		expectCode int
		sizeData   int
	}{
		{
			name: "login user success",
			request: `{
				"email":"test@hanifz.com",
				"password":"123",
				}`,
			expectCode: http.StatusOK,
			sizeData:   1,
		},
	}

	userRepository := repositories.NewUserRepository(configs.DB)

	userService := usecases.NewUserUsecase(userRepository)

	userController := NewUserController(userService)

	e := InitEchoTestAPI()
	for _, testCase := range testCases {
		req := httptest.NewRequest(http.MethodPost, "/users/login", strings.NewReader(testCase.request))
		req.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		if assert.NoError(t, userController.SignIn(c)) {

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
