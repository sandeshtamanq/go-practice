package auth

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/sandeshtamanq/jwt/entity"
	"github.com/sandeshtamanq/jwt/service/user"
	"github.com/sandeshtamanq/jwt/utils"
)

var secretKey = "superSecretKey"

type contextKey string

const userKey contextKey = "userId"

func ValidateJwt(handlerFunc http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userRepository := user.UserRepository()
		tokenStr := GetTokenFromRequest(r)

		token, err := VerifyJwt(tokenStr)

		if err != nil {
			permissionDenied(w)
			return
		}

		if !token.Valid {
			permissionDenied(w)
			return
		}

		claims := token.Claims.(jwt.MapClaims)

		str := claims["userId"].(string)

		userId, err := strconv.Atoi(str)

		if err != nil {
			log.Fatal("error")
			permissionDenied(w)
			return
		}

		u, err := userRepository.GetUserById(userId)

		if err != nil {
			permissionDenied(w)
			return
		}

		ctx := r.Context()

		ctx = context.WithValue(ctx, userKey, u.ID)
		r = r.WithContext(ctx)

		handlerFunc(w, r)
	}

}

func CreateJwt(secret []byte, payload entity.User) (string, error) {
	expiration := time.Second * time.Duration(3000)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId":     payload.ID,
		"email":      payload.Email,
		"expiration": time.Now().Add(expiration).Unix(),
	})

	tokenStr, err := token.SignedString(secretKey)

	if err != nil {
		return "", err
	}

	return tokenStr, nil
}

func VerifyJwt(token string) (*jwt.Token, error) {
	return jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}

		return []byte(secretKey), nil
	})
}

func GetTokenFromRequest(r *http.Request) string {
	tokenAuth := r.Header.Get("Authorization")
	tokenQuery := r.URL.Query().Get("token")

	if tokenAuth != "" {
		return tokenAuth
	}

	if tokenQuery != "" {
		return tokenQuery
	}

	return ""
}

func permissionDenied(w http.ResponseWriter) {
	utils.WriteError(w, http.StatusForbidden, fmt.Errorf("permission denied"))
}

func GetCurrentUserId(ctx context.Context) int {
	userId, ok := ctx.Value(userKey).(int)

	if !ok {
		return -1
	}

	return userId
}
