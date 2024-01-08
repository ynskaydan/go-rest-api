// cmd/main.go

package main

import (
	"UserManagementProject/api/handlers"
	"UserManagementProject/api/routes"
	"UserManagementProject/internal/repository"
	"UserManagementProject/internal/service"
	"UserManagementProject/pkg/db"
	"github.com/rs/cors"
	_ "github.com/rs/cors"
	"log"
	"net/http"
)

func main() {
	database, err := db.InitializeDB("./pkg/mydatabase.db")
	if err != nil {
		log.Fatal(err)
	}
	defer database.Close()

	userRepo := repository.NewUserRepository(database)
	userService := service.NewUserService(userRepo)
	userHandler := handlers.NewUserHandler(userService)
	router := routes.SetupRoutes(userHandler)
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:3000"},        // Frontend working on port 3000 so that i allowed request for just this ip
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE"}, // Allowed request types
	})
	handler := c.Handler(router)

	http.Handle("/", handler)
	//http.ListenAndServe(":8000", nil)
	log.Println("Server is starting on port 8000...")
	log.Fatal(http.ListenAndServe(":8000", handler))
}
