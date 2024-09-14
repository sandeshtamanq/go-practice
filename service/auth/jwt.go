package auth

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/sandeshtamanq/jwt/database"
	"github.com/sandeshtamanq/jwt/entity"
	"github.com/sandeshtamanq/jwt/utils"
)

var secretKey = "superSecretKey"

type contextKey string

const UserKey contextKey = "userId"

func ValidateJwt(handlerFunc http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var user entity.User
		authToken := GetTokenFromRequest(r)

		tokenStr := strings.Split(authToken, " ")[1]

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

		if err := database.DB.Where("id = ?", userId).First(&user).Error; err != nil {
			permissionDenied(w)

			return
		}

		ctx := r.Context()

		ctx = context.WithValue(ctx, UserKey, user.ID)
		r = r.WithContext(ctx)

		handlerFunc(w, r)
	}

}

func CreateJwt(payload *entity.User) (string, error) {
	expiration := time.Second * time.Duration(3000)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId":     strconv.Itoa(int(payload.ID)),
		"email":      payload.Email,
		"expiration": time.Now().Add(expiration).Unix(),
	})

	tokenStr, err := token.SignedString([]byte(secretKey))

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
	utils.WriteError(w, http.StatusForbidden, map[string]string{"message": "forbidden resource"})
}

func GetCurrentUserId(ctx context.Context) uint {
	userId := ctx.Value(UserKey).(uint)

	if userId < 1 {
		return 0
	}

	return userId
}
