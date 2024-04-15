package main

import (
	"brainyquiz/config/initializers"
	"brainyquiz/internal/delivery/handlers"
	"brainyquiz/internal/domain/services"
	"brainyquiz/internal/repository/user"

	"github.com/gin-gonic/gin"
)

func main() {
	initializers.LoadEnvVariables()
	db, err := initializers.InitDB()
	if err != nil {
		panic("Failed to connect to database")
	}
	router := gin.Default()

	UserRepository := user.NewUserRepository(db)
	userService := services.NewUserService(UserRepository)
	usersHandler := handlers.NewUserHandler(userService)

	router.POST("/users", usersHandler.CreateUser)

	if err := router.Run(":8080"); err != nil {
		panic("Failed to start server")
	}

}
