package usecases

import (
	"errors"
	"mini-project-alterra/helpers"
	"mini-project-alterra/middlewares"
	"mini-project-alterra/models"
	"mini-project-alterra/repositories"
)

type UserUsecase interface {
	Login(input models.LoginInput) (string, error)
	Register(input models.RegisterInput) (models.User, error)
	GetCredential(userId int) (models.User, error)
	UpdateUser(userId int, input models.User) (models.User, error)
	DeleteUser(userId int) error
}

type userUsecase struct {
	repository repositories.UserRepository
}

func NewUserUsecase(repository repositories.UserRepository) *userUsecase {
	return &userUsecase{repository}
}

// Login godoc
// @Summary      Login account
// @Description  Login an account
// @Tags         User
// @Accept       json
// @Produce      json
// @Param        request body models.LoginInput true "Payload Body [RAW]"
// @Success      200
// @Failure      400
// @Failure      404
// @Failure      500
// @Router       /users/login [post]
func (s *userUsecase) Login(input models.LoginInput) (string, error) {
	var accessToken string

	user, _ := s.repository.GetUserByEmail(input.Email)
	if user.ID == 0 {
		return accessToken, errors.New("email/password is wrong")
	}

	valid := helpers.ComparePassword(input.Password, user.Password)
	if !valid {
		return accessToken, errors.New("email/password is wrong")
	}

	accessToken, err := middlewares.CreateToken(int(user.ID))
	if err != nil {
		return accessToken, err
	}

	return accessToken, nil
}

// Register godoc
// @Summary      Register account
// @Description  Register an account
// @Tags         User
// @Accept       json
// @Produce      json
// @Param        request body models.RegisterInput true "Payload Body [RAW]"
// @Success      200
// @Failure      400
// @Failure      404
// @Failure      500
// @Router       /users/register [post]
func (s *userUsecase) Register(input models.RegisterInput) (models.User, error) {
	var user models.User

	user, _ = s.repository.GetUserByEmail(input.Email)
	if user.ID > 0 {
		return user, errors.New("email already used")
	}

	user, _ = s.repository.GetUserByUsername(input.Username)
	if user.ID > 0 {
		return user, errors.New("username already used")
	}

	password, err := helpers.HashPassword(input.Password)
	if err != nil {
		return user, err
	}

	user.Email = input.Email
	user.Username = input.Username
	user.FullName = input.FullName
	user.Password = password

	user, err = s.repository.CreateUser(user)
	if err != nil {
		return user, err
	}

	return user, nil
}

// GetCredential godoc
// @Summary      Get user credentials
// @Description  User credentials
// @Tags         User
// @Accept       json
// @Produce      json
// @Success      200
// @Failure      400
// @Failure      404
// @Failure      500
// @Router       /users [get]
// @Security BearerAuth
func (s *userUsecase) GetCredential(userId int) (models.User, error) {
	user, err := s.repository.GetUserById(userId)
	if err != nil {
		return user, err
	}

	return user, nil
}

// UpdateUser godoc
// @Summary      Update user
// @Description  Update user
// @Tags         User
// @Accept       json
// @Produce      json
// @Param        request body models.RegisterInput true "Payload Body [RAW]"
// @Success      200
// @Failure      400
// @Failure      404
// @Failure      500
// @Router       /users [patch]
// @Security BearerAuth
func (s *userUsecase) UpdateUser(userId int, input models.User) (models.User, error) {
	user, err := s.repository.GetUserById(userId)
	if err != nil {
		return user, err
	}
	user.ID = uint(userId)

	user, err = s.repository.GetUserByEmail2(userId, input.Email)
	if user.ID > 0 {
		return user, errors.New("email already used")
	}

	user, err = s.repository.GetUserByUsername2(userId, input.Username)
	if user.ID > 0 {
		return user, errors.New("username already used")
	}

	user, err = s.repository.GetUserById(userId)
	if err != nil {
		return user, err
	}
	user.ID = uint(userId)
	if input.Email != "" {
		user.Email = input.Email
	} else {
		input.Email = user.Email
	}
	if input.Username != "" {
		user.Username = input.Username
	} else {
		input.Username = user.Username
	}
	if input.FullName != "" {
		user.FullName = input.FullName
	}
	if input.Password != "" {
		password, _ := helpers.HashPassword(input.Password)
		user.Password = password
	}
	user, err = s.repository.UpdateUser(user)
	return user, err
}

// DeleteUser godoc
// @Summary      Delete user
// @Description  Delete user
// @Tags         User
// @Accept       json
// @Produce      json
// @Success      200
// @Failure      400
// @Failure      404
// @Failure      500
// @Router       /users [delete]
// @Security BearerAuth
func (s *userUsecase) DeleteUser(userId int) error {
	user, err := s.repository.GetUserById(userId)
	if err != nil {
		return err
	}

	err = s.repository.DeleteUser(user)
	return err
}
