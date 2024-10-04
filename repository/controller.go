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

func (r *Repository) CreateUser(c *fiber.Ctx) error {
	user := models.User{} // Création d'une nouvelle instance de User

	// Récupère le corps de la requête et parse les données JSON pour créer l'utilisateur
	err := c.BodyParser(&user)
	if err != nil { // Si une erreur est rencontrée lors du parsing du corps de la requête
		// Définit le statut de la réponse HTTP et renvoie un message d'erreur
		c.Status(http.StatusUnprocessableEntity).JSON(
			&fiber.Map{
				"message": "Request failed", // Message d'erreur en cas de parsing échoué
			},
		)
		// Retourne une erreur pour interrompre l'exécution de la fonction
		return err
	}

	// Validation des données de l'utilisateur
	errors := ValidateStruct(user)
	if errors != nil { // Si une erreur est détectée lors de la validation des données du nouvel utilisateur
		return c.Status(fiber.StatusBadRequest).JSON(errors) // Retourne les erreurs de validation avec un statut Bad Request
	}

	// Insertion du nouvel utilisateur dans la base de données
	if err := r.DB.Create(&user).Error; err != nil { // Si une erreur est rencontrée lors de l'insertion
		// Définit le statut de la réponse HTTP et renvoie un message d'erreur
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": "Could not create user", // Message d'erreur en cas d'échec de création
		})
	}

	// Si tout se passe bien, retourne une réponse de succès avec les détails de l'utilisateur créé
	c.Status(http.StatusOK).JSON(fiber.Map{
		"status":  "success",
		"message": "User created successfully", // Message de succès
		"data":    user,                        // Renvoie l'objet utilisateur créé
	})
	return nil // Retourne nil pour indiquer que la requête a été traitée avec succès
}

func (r *Repository) UpdateUser(c *fiber.Ctx) error {
	user := models.User{} // Création d'une nouvelle instance de User

	// Récupère le corps de la requête et parse les données JSON pour créer l'utilisateur
	err := c.BodyParser(&user)
	if err != nil { // Si une erreur est rencontrée lors du parsing du corps de la requête
		// Définit le statut de la réponse HTTP et renvoie un message d'erreur
		c.Status(http.StatusUnprocessableEntity).JSON(&fiber.Map{
			"message": "Request failed", // Message d'erreur en cas de parsing échoué
		},
		)
		// Retourne une erreur pour interrompre l'exécution de la fonction
		return err
	}

	// Validation des données de l'utilisateur
	errors := ValidateStruct(user)
	if errors != nil { // Si une erreur est détectée lors de la validation des nouvels données de l'utilisateur
		return c.Status(fiber.StatusBadRequest).JSON(errors) // Retourne les erreurs de validation avec un statut Bad Request
	}

	db := r.DB           // Accès à la base de données via l'attribut DB du repository
	id := c.Params("id") // Récupère l'ID de l'utilisateur à modifier

	if id == "" { // Si l'ID n'est pas fourni dans l'URL
		c.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"message": "ID cannot be empty",
		})
		return nil // Retourne une erreur pour interrompre l'exécution de la fonction
	}

	if db.Model(&user).Where("id =?", id).Updates(user).RowsAffected == 0 { // Si aucun utilisateur ne correspond à l'ID fourni
		c.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"message": "Could not find user profile with ID: " + id,
		})
		return nil // Retourne une erreur pour interrompre l'exécution de la fonction
	}
	c.Status(http.StatusOK).JSON(
		&fiber.Map{
			"status":  "success",
			"message": "User succesfully updated",
		})
	return nil // Retourne une erreur pour interrompre l'exécution de la fonction
}

func (r *Repository) DeleteUser(c *fiber.Ctx) error {
	userModel := models.User{} // Création d'une nouvelle instance de User
	id := c.Params("id")       // Récupère l'ID de l'utilisateur à supprimer

	if id == "" { // Si l'ID n'est pas fourni dans l'URL
		c.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"message": "ID cannot be empty",
		})
		return nil // Retourne une erreur pour interrompre l'exécution de la fonction
	}

	err := r.DB.Delete(userModel, id) // Suppression de l'utilisateur dans la base de données

	if err.Error != nil { // Si une erreur est rencontrée lors de la suppression
		c.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"message": "could not delete",
		})
		return err.Error
	}
	c.Status(http.StatusOK).JSON(&fiber.Map{
		"status":  "success",
		"message": "User deleted successfully",
	})
	return nil // Retourne nil pour indiquer que la requête a été traitée avec succès
}

func (r *Repository) GetUser(c *fiber.Ctx) error {
	userModel := models.User{} // Création d'une nouvelle instance de User
	id := c.Params("id")       // Récupère l'ID de l'utilisateur à supprimer

	if id == "" { // Si l'ID n'est pas fourni dans l'URL
		c.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"message": "ID cannot be empty",
		})
		return nil // Retourne une erreur pour interrompre l'exécution de la fonction
	}

	err := r.DB.Where("id =?", id).First(&userModel).Error // Récupère l'utilisateur dans la base de données

	if err != nil { // Si une erreur est rencontrée lors de la récupération de l'utilisateur
		c.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"message": "User not found",
		})
		return err
	}

	c.Status(http.StatusOK).JSON(&fiber.Map{
		"status": "success", "message": "User fetched successfully", "data": userModel,
	})
	return nil // Retourne nil pour indiquer que la requête a été traitée avec succès
}
