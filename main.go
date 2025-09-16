package main

import (
	"ozinshe/handlers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	corsConfig := cors.Config{
		AllowAllOrigins: true,
		AllowHeaders:    []string{"*"},
		AllowMethods:    []string{"*"},
	}
	r.Use(cors.New(corsConfig))
	authHandlers := handlers.AuthHandler{}
	r.POST("/auth/signIn", authHandlers.SignIn)
	r.Run(":8010")
}
