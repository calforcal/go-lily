package auth

import (
	"time"

	"github.com/calforcal/can-lily-eat-it/config"
	"github.com/calforcal/can-lily-eat-it/storage"
	"github.com/golang-jwt/jwt/v5"
)

type TokenResponse struct {
	AccessToken string `json:"access_token"`
}

type JwtClaims struct {
	UUID  string `json:"uuid"`
	Name  string `json:"name"`
	Email string `json:"email"`
	jwt.RegisteredClaims
}

func IssueJwt(user *storage.User) (TokenResponse, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, JwtClaims{
		UUID:  user.UUID,
		Name:  user.Name,
		Email: user.Email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
		},
	})

	tokenString, err := token.SignedString([]byte(config.JWT_SECRET))
	if err != nil {
		return TokenResponse{}, err
	}

	return TokenResponse{AccessToken: tokenString}, nil
}
