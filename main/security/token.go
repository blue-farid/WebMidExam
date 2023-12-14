package security

import (
	"github.com/blue-farid/WebMidExam/config"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"time"
)

func GenerateJWT(id uint) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["user"] = id
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	tokenString, err := token.SignedString([]byte(config.Cfg.Server.Secret))
	return tokenString, err
}

func ExtractUserIDFromToken(c echo.Context) uint {
	tokenString := c.Request().Header.Get("Authorization")

	token, _ := jwt.Parse(
		tokenString,
		func(token *jwt.Token) (interface{}, error) {
			return []byte(config.Cfg.Server.Secret), nil
		},
	)

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		userID := claims["user"].(float64)
		return uint(userID)
	} else {
		return 0
	}
}
