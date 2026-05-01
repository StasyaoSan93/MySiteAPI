package core

import (
	"os"
	"time"

	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func CreateToken(username string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": username,
		"exp": time.Now().Add(time.Hour * 1).Unix(),
	})
	return token.SignedString([]byte(os.Getenv("SECRET_KEY")))
}

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 1. Берем заголовок Authorization
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is required"})
			c.Abort() // Останавливаем выполнение запроса
			return
		}

		// 2. Отрезаем "Bearer "
		// Ожидаем формат: "Bearer <token>"
		tokenString := ""
		if len(authHeader) > 7 && authHeader[:7] == "Bearer " {
			tokenString = authHeader[7:]
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token format"})
			c.Abort()
			return
		}

		// 3. Парсим и валидируем токен
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("SECRET_KEY")), nil
		})

		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
			c.Abort()
			return
		}

		// 4. (Опционально) Сохраняем данные пользователя в контекст, чтобы достать их в контроллере
		if claims, ok := token.Claims.(jwt.MapClaims); ok {
			c.Set("username", claims["sub"])
		}

		c.Next() // Идем дальше к обработчику (контроллеру)
	}
}
