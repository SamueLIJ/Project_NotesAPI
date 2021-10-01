package middleware

import (
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
)

// func JWTauthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
// 	return func(c echo.Context) error {

// 		authorizationFromHeader := c.Request().Header.Get("authorization")

// 		tokenString := strings.ReplaceAll(authorizationFromHeader, "Bearer ", "")

// 		claims := jwt.MapClaims{}
// 		token, err := jwt.ParseWithClaims(tokenString, &claims, func(t *jwt.Token) (interface{}, error) {
// 			return []byte("legal"), nil
// 		})

// 		if err != nil && !token.Valid {
// 			return c.String(http.StatusForbidden, "Token Salah")
// 		}

// 		c.Set("email", claims["userID"])
// 		return next(c)
// 	}
// }

func CreateToken(userId int) (string, error) {
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["userId"] = userId
	claims["exp"] = time.Now().Add(time.Hour).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte("legal"))
}

func ExtractTokenUserId(e echo.Context) int {
	user := e.Get("user").(*jwt.Token)
	if user.Valid {
		claims := user.Claims.(jwt.MapClaims)
		userId := claims["userId"].(int)
		return userId
	}
	return 0
}