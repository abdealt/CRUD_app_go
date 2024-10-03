package repository

import (
	"github.com/gofiber/fiber/v2"
)

// On créer une méthode SetupRoutes qui va ajouter les routes à l'application Fiber
// Cette méthode prend en argument une instance de l'application Fiber
// repo est un pointeur sur la structure Repository
func (repo *Repository) SetupRoutes(app *fiber.App) {
	api := app.Group("/api")
	api.Get("/users", repo.GetUsers)
}
