package main

import (
	"github.com/abdealt/go_crud/bootstrap"  // Import du package bootstrap qui initialise l'application
	"github.com/abdealt/go_crud/repository" // Import du package repository qui contient la logique d'accès aux données (Routes, Controller)
	"github.com/gofiber/fiber/v2"           // Import du framework Fiber pour la gestion des requêtes HTTP et des routes
)

// Déclaration d'un type Repository qui est un alias pour repository.Repository.
// Cela permet de référencer repository.Repository plus facilement sous le nom Repository
type Repository repository.Repository

// Fonction principale qui est le point d'entrée de l'application
func main() {
	// Création d'une nouvelle instance de l'application Fiber
	app := fiber.New()

	// Appel de la fonction InitApp qui est dans le package bootstrap
	bootstrap.InitApp(app)
}
