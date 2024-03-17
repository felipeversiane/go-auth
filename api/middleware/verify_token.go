package middleware

import (
	"fmt"
	"os"

	"github.com/felipeversiane/go-auth/config/httperr"
	"github.com/felipeversiane/go-auth/config/logger"
	domain "github.com/felipeversiane/go-auth/internal/account"
	"github.com/felipeversiane/go-auth/internal/account/utils"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

var (
	JWT_SECRET_KEY = "JWT_SECRET_KEY"
)

func VerifyTokenMiddleware(c *gin.Context) {
	secret := os.Getenv(JWT_SECRET_KEY)
	tokenValue := utils.RemoveBearerPrefix(c.Request.Header.Get("Authorization"))

	token, err := jwt.Parse(tokenValue, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); ok {
			return []byte(secret), nil
		}

		return nil, httperr.NewBadRequestError("invalid token")
	})
	if err != nil {
		errRest := httperr.NewUnauthorizedRequestError("invalid token")
		c.JSON(errRest.Code, errRest)
		c.Abort()
		return
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		errRest := httperr.NewUnauthorizedRequestError("invalid token")
		c.JSON(errRest.Code, errRest)
		c.Abort()
		return
	}

	userDomain := domain.NewTokenUserDomain(
		claims["id"].(string),
		claims["email"].(string),
		claims["first_name"].(string),
		claims["last_name"].(string),
		int8(claims["age"].(float64)),
	)
	logger.Info(fmt.Sprintf("User authenticated: %#v", userDomain))
}
