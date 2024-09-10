package user

import (
	"fmt"

	"github.com/sandeshtamanq/jwt/database"
	"github.com/sandeshtamanq/jwt/entity"
)

type Repository struct {
}

func UserRepository() *Repository {
	return &Repository{}
}

func (r *Repository) RegisterUser(u *entity.User) error {

	createdUser := database.DB.Create(&u)

	err := createdUser.Error

	if err != nil {
		return err
	}

	return nil
}

func (r *Repository) GetUserByEmail(email string) error {
	var user entity.User
	if err := database.DB.Where("email = ?", email).First(&user).Error; err == nil {
		return fmt.Errorf("user with email %v already exists", user.Email)
	}

	return nil

}
