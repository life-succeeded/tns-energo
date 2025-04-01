package authorize

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"net/http"
	"strings"
)

type Authorize struct {
	UserId    int    `json:"user_id"`
	Email     string `json:"email"`
	IsAdmin   bool   `json:"is_admin"`
	UserAgent string `json:"user_agent"`
	Iat       int    `json:"iat"`
	Exp       int    `json:"exp"`
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
