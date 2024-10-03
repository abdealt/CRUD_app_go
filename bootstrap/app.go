package bootstrap

import (
	"log"
	"os"

	"github.com/abdealt/go_crud/database/migrations"
	"github.com/abdealt/go_crud/database/storage"
	"github.com/abdealt/go_crud/repository"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
)

// InitApp initialise l'application Fiber avec les middleware et les configurations nécessaires
func InitApp(app *fiber.App) {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}
	config := &storage.Config{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		DbName:   os.Getenv("DB_NAME"),
		SSLMode:  os.Getenv("DB_SSLMODE"),
	}

	db, err := storage.NewConnection(*config)
	if err != nil {
		log.Fatal("Could not load the database")
	}

	err = migrations.MigrateUsers(db)

	if err != nil {
		log.Fatal("Could not migrate the users table to db")
	}
	// Création du repository avec la connexion à la base de données
	repo := repository.Repository{
		DB: db,
	}

	app.Use(cors.New(cors.Config{AllowCredentials: true}))
	repo.SetupRoutes(app)
	app.Listen(":8081")
}
