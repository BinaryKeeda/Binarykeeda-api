package main
import (
	"log"
	"net/http"
	"binarykeeda-api/internal/config" 
	"binarykeeda-api/internal/db"
	"binarykeeda-api/internal/handlers"
	"binarykeeda-api/internal/repository"
	"binarykeeda-api/internal/routes"
)

func main() {
	cfg := config.Load()

	mongoClient := db.Connect(cfg.MongoURI)
	database := mongoClient.Database(cfg.MongoDB)

	userRepo := repository.NewUserRepository(database)
	userHandler := handlers.NewUserHandler(userRepo)

	router := routes.Register(userHandler)

	log.Println("Server running on :" + cfg.AppPort)
	log.Fatal(http.ListenAndServe(":"+cfg.AppPort, router))
}
