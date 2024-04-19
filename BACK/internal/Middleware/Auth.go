package middleware

import (
	"brainyquiz/internal/repository/user"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"gorm.io/gorm"
)

func Auth(db *gorm.DB) gin.HandlerFunc {
	userRepository := user.NewUserRepository(db)
	return func(c *gin.Context) {
		// Verificar el encabezado de autorización
		AuthHeader := c.GetHeader("Authorization")
		if AuthHeader == "" {
			c.JSON(400, gin.H{"error": "Authorization header is required"})
			c.Abort()
			return
		}

		parts := strings.Split(AuthHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			c.JSON(400, gin.H{"error": "Authorization header must be in the format 'Bearer <token>'"})
			c.Abort()
			return
		}
		tokenString := parts[1]

		// Verificar y analizar el token JWT
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(os.Getenv("SECRET")), nil
		})
		if err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			c.Abort()
			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			if float64(time.Now().Unix()) > claims["exp"].(float64) {
				fmt.Println("Token expirado")
				c.AbortWithStatus(401)
				return
			}

			email := claims["email"].(string)

			user, err := userRepository.GetUserByEmail(c, email)
			if err != nil {
				c.JSON(401, gin.H{"error": "Error finding user"})
				return
			}

			// Establecer el usuario en el contexto para su uso posterior
			c.Set("user", user)
			c.Next()
		} else {
			c.AbortWithStatus(401)
			return
		}
	}
}
