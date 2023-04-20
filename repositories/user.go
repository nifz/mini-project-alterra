package repositories

import (
	"mini-project-alterra/models"

	"gorm.io/gorm"
)

type UserRepository interface {
	CreateUser(user models.User) (models.User, error)
	GetUserByEmail(email string) (models.User, error)
	GetUserByUsername(username string) (models.User, error)
	GetUserByEmail2(userId int, email string) (models.User, error)
	GetUserByUsername2(userId int, username string) (models.User, error)
	GetUserById(userId int) (models.User, error)
	UpdateUser(user models.User) (models.User, error)
	DeleteUser(user models.User) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *userRepository {
	return &userRepository{db}
}

func (r *userRepository) GetUserByUsername(username string) (models.User, error) {
	var user models.User
	err := r.db.Where("username = ?", username).First(&user).Error
	return user, err
}

func (r *userRepository) GetUserByEmail(email string) (models.User, error) {
	var user models.User
	err := r.db.Where("email = ?", email).First(&user).Error
	return user, err
}

func (r *userRepository) GetUserByUsername2(userId int, username string) (models.User, error) {
	var user models.User
	err := r.db.Where("id != ? AND username = ?", userId, username).First(&user).Error
	return user, err
}

func (r *userRepository) GetUserByEmail2(userId int, email string) (models.User, error) {
	var user models.User
	err := r.db.Where("id != ? AND email = ?", userId, email).First(&user).Error
	return user, err
}

func (r *userRepository) CreateUser(user models.User) (models.User, error) {
	err := r.db.Create(&user).Error
	return user, err
}

func (r *userRepository) GetUserById(userId int) (models.User, error) {
	var user models.User
	err := r.db.Where("id = ?", userId).First(&user).Error
	return user, err
}

func (r *userRepository) UpdateUser(user models.User) (models.User, error) {
	err := r.db.Save(&user).Error
	return user, err
}

func (r *userRepository) DeleteUser(user models.User) error {
	err := r.db.Delete(&user).Error
	return err
}
