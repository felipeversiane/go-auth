package domain

import (
	"fmt"
	"os"
	"time"

	"github.com/felipeversiane/go-auth/config/httperr"
	"github.com/felipeversiane/go-auth/internal/account/utils"
	"github.com/golang-jwt/jwt"
)

func (ud *userDomain) GenerateToken() (string, string, *httperr.HttpError) {
	acessToken, err := ud.GenerateAcessToken()
	if err != nil {
		return "", "", err
	}
	refreshToken, err := ud.GenerateRefreshToken()
	if err != nil {
		return "", "", err
	}
	return acessToken, refreshToken, nil
}

func (ud *userDomain) GenerateAcessToken() (string, *httperr.HttpError) {
	secret := os.Getenv(JWT_SECRET_KEY)

	claims := jwt.MapClaims{
		"id":         ud.id,
		"email":      ud.email,
		"first_name": ud.firstName,
		"last_name":  ud.lastName,
		"age":        ud.age,
		"exp":        time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", httperr.NewInternalServerError(
			fmt.Sprintf("error trying to generate jwt token, err=%s", err.Error()))
	}

	return tokenString, nil
}
func (ud *userDomain) GenerateRefreshToken() (string, *httperr.HttpError) {
	secret := os.Getenv(JWT_REFRESH_SECRET_KEY)

	refreshTokenClaims := jwt.MapClaims{
		"id":  ud.id,
		"exp": time.Now().Add(time.Hour * 24 * 7).Unix(),
	}

	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshTokenClaims)

	refreshTokenString, err := refreshToken.SignedString([]byte(secret))
	if err != nil {
		return "", httperr.NewInternalServerError(
			fmt.Sprintf("error trying to generate refresh token, err=%s", err.Error()))
	}

	return refreshTokenString, nil
}

func VerifyAcessToken(tokenValue string) (UserDomainInterface, *httperr.HttpError) {
	secret := os.Getenv(JWT_SECRET_KEY)

	token, err := jwt.Parse(utils.RemoveBearerPrefix(tokenValue), func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); ok {
			return []byte(secret), nil
		}

		return nil, httperr.NewBadRequestError("invalid token")
	})
	if err != nil {
		return nil, httperr.NewUnauthorizedRequestError("invalid token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return nil, httperr.NewUnauthorizedRequestError("invalid token")
	}

	return &userDomain{
		id:        claims["id"].(string),
		email:     claims["email"].(string),
		firstName: claims["first_name"].(string),
		lastName:  claims["last_name"].(string),
		age:       int8(claims["age"].(float64)),
	}, nil
}

func VerifyRefreshToken(tokenValue string) (UserDomainInterface, *httperr.HttpError) {
	secret := os.Getenv(JWT_REFRESH_SECRET_KEY)

	token, err := jwt.Parse(utils.RemoveBearerPrefix(tokenValue), func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); ok {
			return []byte(secret), nil
		}

		return nil, httperr.NewBadRequestError("invalid token")
	})
	if err != nil {
		return nil, httperr.NewUnauthorizedRequestError("invalid token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return nil, httperr.NewUnauthorizedRequestError("invalid token")
	}

	return &userDomain{
		id: claims["id"].(string),
	}, nil
}
