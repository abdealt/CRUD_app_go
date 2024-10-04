package repository

import (
	"net/http" // Import du package http pour manipuler les requêtes HTTP

	"github.com/abdealt/go_crud/database/models" // Import du package models pour accéder à la structure User
	"github.com/gofiber/fiber/v2"                // Import du framework Fiber pour gérer les requêtes HTTP
	"github.com/morkid/paginate"                 // Import du package paginate pour la pagination des résultats
)

// GetUsers récupère tous les utilisateurs de la base de données et les renvoie sous forme paginée
// c est le contexte de la requête, qui contient des informations sur la requête HTTP et permet de générer une réponse
func (r *Repository) GetUsers(c *fiber.Ctx) error {
	db := r.DB                        // Accès à la base de données via l'attribut DB du repository
	model := db.Model(&models.User{}) // Spécification du modèle User pour la requête. Cela indique à GORM que l'on veut interagir avec la table des utilisateurs

	// Initialisation du système de pagination avec une configuration personnalisée
	pg := paginate.New(&paginate.Config{
		DefaultSize:        20,   // Taille par défaut de chaque page (20 utilisateurs par page)
		CustomParamEnabled: true, // Active la possibilité d'utiliser des paramètres personnalisés pour la pagination dans l'URL (ex: ?page=2&size=30)
	})

	// Génération de la pagination avec le modèle User en fonction des paramètres fournis dans la requête
	page := pg.With(model).Request(c.Request()).Response(&[]models.User{}) // Récupère une réponse paginée des utilisateurs

	// Renvoie la réponse avec le statut HTTP 200 (OK) et le contenu paginé dans un format JSON
	// Le résultat paginé est renvoyé dans une map associée à la clé "data"
	c.Status(http.StatusOK).JSON(&fiber.Map{
		"data": page,
	})

	return nil // Retourne nil pour indiquer que la requête a été traitée avec succès
}
