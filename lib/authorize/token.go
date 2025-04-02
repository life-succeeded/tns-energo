package authorize

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"net/http"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type Authorize struct {
	UserId  int    `json:"user_id"`
	Email   string `json:"email"`
	IsAdmin bool   `json:"is_admin"`
	Iat     int    `json:"iat"`
	Exp     int    `json:"exp"`
}

func GetToken(r *http.Request) string {
	return r.Header.Get("Authorization")
}

func GetAuth(r *http.Request) (auth Authorize, err error) {
	return Parse(GetToken(r))
}

func Parse(token string) (auth Authorize, err error) {
	parts := strings.Split(token, ".")
	if len(parts) < 2 {
		return Authorize{}, errors.New("wrong authorization parts")
	}

	return decode(parts[1])
}

func decode(token string) (Authorize, error) {
	buf, err := base64.StdEncoding.DecodeString(token)
	if err != nil {
		return Authorize{}, err
	}

	auth := Authorize{}
	err = json.Unmarshal(buf, &auth)
	if err != nil {
		return Authorize{}, err
	}

	return auth, nil
}

func NewToken(userId int, email string, isAdmin bool, duration time.Duration, secret string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["user_id"] = userId
	claims["email"] = email
	claims["is_admin"] = isAdmin
	claims["iat"] = time.Now().Unix()
	claims["exp"] = time.Now().Add(duration).Unix()

	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func NewEmptyToken(secret string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
