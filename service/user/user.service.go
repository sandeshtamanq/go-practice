package user

import (
	"fmt"
	"log"
	"net/http"

	"github.com/sandeshtamanq/jwt/entity"
	"github.com/sandeshtamanq/jwt/utils"
	"golang.org/x/crypto/bcrypt"
)

type Service struct {
	repository Repository
}

func UserService(repository Repository) *Service {
	return &Service{
		repository: repository,
	}
}

func HandleRegister(w http.ResponseWriter, r *http.Request) {
	userRepository := UserRepository()
	var user entity.User

	if r.Body == nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("empty body"))
	}

	utils.ParseJSON(r, &user)

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)

	if err != nil {
		log.Fatal(err)
	}
	user.Password = string(hashedPassword)

	err = userRepository.GetUserByEmail(user.Email)

	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, map[string]string{"message": err.Error()})
		return
	}

	err = userRepository.RegisterUser(&user)

	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	utils.WriteJSON(w, http.StatusCreated, map[string]string{"message": "user created successfully"})
}
