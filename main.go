package main

import (
	"net/http"

	"github.com/Yurog22/gorm-test/controllers"
	"github.com/Yurog22/gorm-test/models"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	handler := controllers.New()

	server := &http.Server{
		Addr:    "0.0.0.0:8008",
		Handler: handler,
	}

	models.ConnectDatabase()

	server.ListenAndServe()
}
