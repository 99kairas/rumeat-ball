package middlewares

import (
	"net/http"
	"rumeat-ball/configs"
	"rumeat-ball/dto"

	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

func CreateToken(userID uuid.UUID, role, email string) (string, error) {
	claims := jwt.MapClaims{}
	// token kedua (payload)
	claims["authorized"] = true
	claims["user_id"] = userID
	claims["email"] = email
	claims["role"] = role
	// token pertama (header)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// return bersama token ketiga (dengan secret key)
	return token.SignedString([]byte(configs.JWT_KEY))
}

func CheckRole(role string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			user := c.Get("user").(*jwt.Token)
			if !user.Valid {
				return c.JSON(http.StatusUnauthorized, dto.Response{
					Message:  "unauthorized",
					Response: "permision denied: user is not valid",
				})
			}

			claims := user.Claims.(jwt.MapClaims)
			userRole := claims["role"].(string)

			if userRole == role {
				return next(c)
			}

			return c.JSON(http.StatusUnauthorized, dto.Response{
				Message:  "unauthorized",
				Response: "permision denied: only " + role + " roles are allowed to perform this operation.",
			})
		}
	}
}

func ExtractTokenUserId(e echo.Context) uuid.UUID {
	user := e.Get("user").(*jwt.Token)
	if user.Valid {
		claims := user.Claims.(jwt.MapClaims)
		userId := claims["user_id"].(string)
		uuid, _ := uuid.Parse(userId)
		return uuid
	}
	return uuid.Nil
}
