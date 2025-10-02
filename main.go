package main

import (
	"context"
	"ozinshe/config"
	"ozinshe/handlers"
	"ozinshe/repositories"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/spf13/viper"
)

func main() {
	r := gin.Default()
	corsConfig := cors.Config{
		AllowAllOrigins: true,
		AllowHeaders:    []string{"*"},
		AllowMethods:    []string{"*"},
	}
	r.Use(cors.New(corsConfig))
	err := loadConfig()
	if err != nil {
		panic(err)
	}
	connection, err := connectToDB()
	if err != nil {
		panic(err)
	}
	usersRepository := repositories.NewUsersRepository(connection)
	authHandler := handlers.NewAuthentificationHandler(usersRepository)
	userHandler := handlers.NewUserProfileHandler(usersRepository)
	//Authorization handlers
	r.POST("/create", authHandler.SignUp)
	r.POST("/signIn", authHandler.SigIn)
	r.POST("/logOut", authHandler.Logout)

	r.POST("user/profile", userHandler.CreateUserProfile)
	r.GET("user/profile/:id", userHandler.GetUserProfile)
	r.PUT("user/profile/:id", userHandler.UpdateUserProfile)
	r.PATCH("user/:id/password", userHandler.ChangePassword)
	r.Run(":8010")
}

func connectToDB() (*pgxpool.Pool, error) {
	conn, err := pgxpool.New(context.Background(), config.Config.DbConnectionString)
	if err != nil {
		return nil, err
	}

	err = conn.Ping(context.Background())
	if err != nil {
		return nil, err
	}
	return conn, nil
}

func loadConfig() error {
	viper.AutomaticEnv()
	viper.SetConfigFile(".env")
	if err := viper.BindEnv("APP_HOST"); err != nil {
		viper.SetDefault("APP_HOST", ":8010")
	}
	if err := viper.BindEnv("DB_CONNECTION_STRING"); err != nil {
		viper.SetDefault("DB_CONNECTION_STRING", "postgres://postgres:ansar2007+A@localhost:5432/ozinshe?sslmode=disable")
	}
	if err := viper.BindEnv("JWT_SECRET_KEY"); err != nil {
		viper.SetDefault("JWT_SECRET_KEY", "supersecretkey")
	}
	if err := viper.BindEnv("JWT_EXPIRES_IN"); err != nil {
		viper.SetDefault("JWT_EXPIRES_IN", "24h")
	}
	err := viper.ReadInConfig()
	if err != nil {
		return err
	}

	var mapConfig config.MapConfig
	err = viper.Unmarshal(&mapConfig)
	if err != nil {
		return err
	}

	config.Config = &mapConfig

	return nil
}
