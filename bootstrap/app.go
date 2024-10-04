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

// InitApp initialise l'application Fiber avec les middlewares et configurations nécessaires.
// 'app' est une instance de l'application Fiber qui gère les requêtes HTTP.
func InitApp(app *fiber.App) {
	// Chargement des variables d'environnement à partir du fichier ".env"
	err := godotenv.Load(".env")
	if err != nil {
		// Si le fichier .env ne peut pas être chargé, l'application s'arrête et un message d'erreur est affiché
		log.Fatal("Erreur de chargement des variables d'environnement:", err)
	}

	// Configuration de la base de données via les variables d'environnement récupérées dans ".env"
	config := &storage.Config{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		DbName:   os.Getenv("DB_NAME"),
		SSLMode:  os.Getenv("DB_SSLMODE"),
	}

	// Création d'une nouvelle connexion à la base de données avec la configuration récupérée
	db, err := storage.NewConnection(*config)
	if err != nil {
		// Si la connexion échoue, l'application s'arrête et un message d'erreur est affiché
		log.Fatal("Impossible de se connecter à la base de données:", err)
	}

	// Migration de la table 'users' dans la base de données
	// Cette opération vérifie ou crée la table 'users' dans la base
	err = migrations.MigrateUsers(db)
	if err != nil {
		// Si la migration échoue, l'application s'arrête et un message d'erreur est affiché
		log.Fatal("Impossible de migrer la table 'users' dans la base de données:", err)
	}

	// Création du repository, qui est une abstraction pour accéder aux données via la connexion à la base de données
	repo := repository.Repository{
		DB: db,
	}

	// Ajout des middlewares à l'application
	// Le middleware CORS (Cross-Origin Resource Sharing) est configuré ici pour permettre à un front-end (ex: React)
	// de communiquer avec l'API en autorisant les requêtes provenant de "http://localhost:8080"
	app.Use(cors.New(cors.Config{
		AllowCredentials: true,                    // Permet d'envoyer des cookies et des informations d'authentification dans les requêtes
		AllowOrigins:     "http://localhost:8080", // Autorise les requêtes provenant de cette origine spécifique
		// AllowHeaders: "Origin, Content-Type, Accept", // Décommenter pour spécifier les headers autorisés
	}))

	// Configuration des routes définies dans le repository
	// Cela connecte les routes de l'application (ex: CRUD sur les utilisateurs) à l'instance de l'application Fiber
	repo.SetupRoutes(app)

	// Démarrage du serveur sur le port 8080
	// L'application commencera à écouter les requêtes HTTP à l'adresse localhost:8080
	err = app.Listen(":8080")
	if err != nil {
		// Si le serveur ne parvient pas à démarrer, l'application s'arrête et un message d'erreur est affiché
		log.Fatal("Erreur lors du démarrage du serveur:", err)
	}
}
