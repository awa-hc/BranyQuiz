package main

import (
	"brainyquiz/config/initializers"
	middleware "brainyquiz/internal/Middleware"
	"brainyquiz/internal/delivery/handlers"
	"brainyquiz/internal/domain/services"
	"brainyquiz/internal/repository/auth"
	"brainyquiz/internal/repository/friend"
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

	validateGroup := router.Group("/validate")
	validateGroup.Use(middleware.Auth(db))
	{
		validateGroup.GET("/:email", usersHandler.GetUserByEmail)
	}

	AuthRepository := auth.NewAuthRepository(db)
	authService := services.NewAuthService(AuthRepository)
	authHandler := handlers.NewAuthHandler(*authService)

	Auth := router.Group("/auth")
	{
		Auth.POST("/login", authHandler.LoginWithEmail)
	}

	FriendRepository := friend.NewFriendRepository(db)
	friendService := services.NewFriendService(FriendRepository)
	friendHandler := handlers.NewFriendHandler(*friendService)
	Friend := router.Group("/friend")
	Friend.Use(middleware.Auth(db))
	{
		Friend.POST("/add", friendHandler.AddFriend)
		Friend.POST("/remove", friendHandler.RemoveFriend)
		Friend.GET("/:username", friendHandler.GetFriends)
	}

	if err := router.Run(":8080"); err != nil {
		panic("Failed to start server")
	}

}
