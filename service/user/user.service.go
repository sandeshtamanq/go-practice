package user

import (
	"fmt"
	"log"
	"net/http"

	"github.com/sandeshtamanq/jwt/entity"
	"github.com/sandeshtamanq/jwt/service/auth"
	"github.com/sandeshtamanq/jwt/types"
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

	_, err = userRepository.GetUserByEmail(user.Email)

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

func HandleLogin(w http.ResponseWriter, r *http.Request) {
	userRepository := UserRepository()
	var user types.LoginUserPayload

	if r.Body == nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("empty body"))
		return
	}

	utils.ParseJSON(r, &user)

	u, err := userRepository.GetUserByEmail(user.Email)

	if err != nil {
		utils.WriteError(w, http.StatusUnauthorized, map[string]string{"message": "forbidden"})
		return
	}

	if !auth.ComparePassword(u.Password, user.Password) {
		utils.WriteError(w, http.StatusUnauthorized, map[string]string{"message": "invalid email or password"})
		return
	}

	token, err := auth.CreateJwt(u)

	if err != nil {
		fmt.Println(err)
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, map[string]string{"token": token})
}
