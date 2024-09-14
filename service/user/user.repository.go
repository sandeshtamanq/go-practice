package user

import (
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

func (r *Repository) GetUserByEmail(email string) (*entity.User, error) {
	var user entity.User
	err := database.DB.Where("email = ?", email).First(&user).Error

	return &user, err

}

func (r *Repository) GetUserById(userId int) (*entity.User, error) {
	var user entity.User

	if err := database.DB.Where("id = ?", userId).First(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}
