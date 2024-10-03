package repository

// import (
// 	"net/http"

// 	"github.com/abdealt/go_crud/database/migrations"
// 	//"github.com/abdealt/go_crud/database/models"
// 	"github.com/gofiber/fiber/v2"
// 	"github.com/morkid/paginate"
// 	//"gopkg.in/go-playground/validator.v9"
// )

// // GetUsers prend en charge de récupérer tous les utilisateurs de la base de données
// // c est une instance de l'application Fiber ( context )
// func (r *Repository) GetUsers(c *fiber.Ctx) error {
// 	db := r.DB
// 	model := db.Model(migrations.Users{})

// 	// paginate sert à générer des pages de résultats
// 	pg := paginate.New(&paginate.Config{
// 		DefaultSize:        20,
// 		CustomParamEnabled: true,
// 	})

// 	page := pg.With(model).Request(c.Request()).Response(&[]migrations.Users{})

// 	c.Status(http.StatusOK).JSON(&fiber.Map{
// 		"data": page,
// 	})
// 	return nil
// }

//apres

import (
	"net/http"

	"github.com/abdealt/go_crud/database/models" // Utilise models au lieu de migrations
	"github.com/gofiber/fiber/v2"
	"github.com/morkid/paginate"
)

// GetUsers récupère tous les utilisateurs de la base de données
func (r *Repository) GetUsers(c *fiber.Ctx) error {
	db := r.DB
	model := db.Model(&models.User{}) // Utilisation du modèle User défini dans models

	// paginate sert à générer des pages de résultats
	pg := paginate.New(&paginate.Config{
		DefaultSize:        20,
		CustomParamEnabled: true,
	})

	page := pg.With(model).Request(c.Request()).Response(&[]models.User{})

	c.Status(http.StatusOK).JSON(&fiber.Map{
		"data": page,
	})
	return nil
}
