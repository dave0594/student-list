package main

import (
	"os"

	"github.com/dave0594/student-list/routes"
	"github.com/joho/godotenv"
)

func main() {
	router := routes.Init()

	if err := godotenv.Load(); err != nil {
		router.Logger.Fatal("error loading .env file")
	}

	port := os.Getenv("PORT")

	router.Logger.Fatal(router.Start(":" + port))
}
