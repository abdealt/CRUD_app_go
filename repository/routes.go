package repository

import (
	"github.com/gofiber/fiber/v2" // Import du framework Fiber pour gérer les routes et les requêtes HTTP
)

// SetupRoutes est une méthode qui configure les routes de l'application Fiber.
// Elle prend en argument une instance de l'application Fiber (app) et définit les routes nécessaires.
// repo est un pointeur vers la structure Repository, permettant d'accéder aux méthodes et aux données de celle-ci.
func (repo *Repository) SetupRoutes(app *fiber.App) {
	// Création d'un groupe de routes pour l'API, ce qui permet d'organiser les routes sous un préfixe commun
	api := app.Group("/api")

	// Définition d'une route GET pour récupérer la liste des utilisateurs
	// Lorsque l'URL "/api/users" est appelée, la méthode GetUsers de la structure Repository est exécutée
	api.Get("/users", repo.GetUsers)

	// Définition d'une route POST pour ajouter un utilisateur
	// Lorsque l'URL "/api/user" est appelée avec une requête POST, la méthode CreateUser de la structure Repository est exécutée
	api.Post("/users", repo.CreateUser)

	// Définition d'une route PUT pour mettre à jour un utilisateur
	// Lorsque l'URL "/api/users/:id" est appelée avec une requête PUT, la méthode UpdateUser de la structure Repository est exécutée
	api.Patch("/users/:id", repo.UpdateUser)

	// Définition d'une route DELETE pour supprimer un utilisateur
	// Lorsque l'URL "/api/users/:id" est appelée avec une requête DELETE, la méthode DeleteUser de la structure Repository est exécutée
	api.Delete("/users/:id", repo.DeleteUser)

	// Définition d'une route GET pour récupérer un utilisateur par son ID (Vue de Détails)
	// Lorsque l'URL "/api/users/:id" est appelée, la méthode GetUser de la structure Repository est exécutée
	api.Get("/users/:id", repo.GetUser)
}
