package main

import (
	"capstone_vaccine/databases"
	"capstone_vaccine/router"
	"log"
	"os"

	"github.com/labstack/echo"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	databases.InitDB()

	e := echo.New()

	router.New(e, databases.DB)

	port := os.Getenv("PORT")

	_ = e.Start(port)
}
