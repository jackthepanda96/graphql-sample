package middleware

import (
	"context"
	"fmt"

	"github.com/golang-jwt/jwt"
	echo "github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type User struct {
	Id   int
	Role string
}

type ContextKey struct {
	Name string
}

func JWTMiddleware() echo.MiddlewareFunc {
	return middleware.JWTWithConfig(middleware.JWTConfig{
		SigningMethod: "HS256",
		SigningKey:    []byte("R4HASIA"),
		Skipper: func(c echo.Context) bool {
			return c.Request().Header.Get("Authorization") == ""
		}, SuccessHandler: func(c echo.Context) {
			c.Set("INFO", &User{ExtractToken(c), "admin"})
			fmt.Println(c)
		},
	})
}

func ExtractToken(e echo.Context) int {
	token := e.Get("user").(*jwt.Token)
	if token != nil && token.Valid {
		claims := token.Claims.(jwt.MapClaims)
		id := claims["id"]
		switch id.(type) {
		case float64:
			return int(id.(float64))
		default:
			return id.(int)
		}
	}
	return -1 //invalid
}

func GetAuthFromContext(ctx context.Context) *User {
	raw := ctx.Value(&ContextKey{"user"})
	fmt.Println(raw)
	tmp := ctx.Value("ID")
	fmt.Print(tmp)
	if raw == nil {
		return nil
	}

	return raw.(*User)
}
