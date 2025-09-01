package main

import (
	"auth-api/internal/config"
	"auth-api/internal/database"
	"auth-api/internal/router"
	"auth-api/internal/users"
	"context"
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	ctx := context.Background()
	_, err := os.Stat(".env")
	if err != nil {
		if !os.IsNotExist(err) {
			err := godotenv.Load()
			if err != nil {
				log.Fatalf("Error loading .env file: %+v", err.Error())
			}
		}
	}

	cfg := config.New()
	db, err := database.New(ctx, cfg.DB)
	if err != nil {
		log.Fatal(err)
	}

	usersRepository := users.NewRepository(db)
	usersUsecase := users.NewUsecase(usersRepository)

	r := gin.Default()

	router.RegisterUsersRoutes(r, usersUsecase)

	log.Fatal(r.Run(fmt.Sprintf(":%s", cfg.Api.Port)))
}
