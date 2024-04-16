package main

import (
	"brainyquiz/config/initializers"
	"brainyquiz/internal/delivery/handlers"
	"brainyquiz/internal/domain/services"
	"brainyquiz/internal/repository/user"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	initializers.LoadEnvVariables()
	db, err := initializers.InitDB()
	if err != nil {
		panic("Failed to connect to database")
	}

	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"PUT", "PATCH", "POST", "GET", "DELETE"},
		AllowHeaders:     []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
	}))

	router.OPTIONS("/*any", func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "PUT, POST, GET, DELETE, OPTIONS")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Authorization")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Status(http.StatusOK)
	})

	UserRepository := user.NewUserRepository(db)
	userService := services.NewUserService(UserRepository)
	usersHandler := handlers.NewUserHandler(*userService)

	router.POST("/users", usersHandler.CreateUser)
	router.GET("/users/email/:email", usersHandler.GetUserByEmail)
	router.GET("/users/:userName", usersHandler.GetUserByUserName)

	if err := router.Run(":8080"); err != nil {
		panic("Failed to start server")
	}

}
