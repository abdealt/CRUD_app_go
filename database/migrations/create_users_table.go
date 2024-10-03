// package migrations

// import (
// 	"time"

// 	"gorm.io/gorm"
// )

// // Users est la structure de données pour la table users
// type Users struct {
// 	ID        uint      `gorm:"primary_key" json: "id"`
// 	Name      *string   `gorm:"type:varchar(255);not null" json: "name"`
// 	Email     *string   `gorm:"type:varchar(255);unique;not null" json: "email"`
// 	City      *string   `gorm:"type:varchar(255);not null" json: "city"`
// 	Country   *string   `gorm:"type:varchar(255);not null" json: "country"`
// 	CreatedAt time.Time `json: "date_created"`
// 	UpdateAt  time.Time `json: "date_updated"`
// }

// // MigrateUsers définit la migration pour la table users
// // Elle retourne une erreur si une erreur survient lors de la migration
// // db est une instance de gorm.DB
// func MigrateUsers(db *gorm.DB) error {
// 	err := db.AutoMigrate(&Users{})
// 	return err
// }
//Apres

package migrations

import (
	"github.com/abdealt/go_crud/database/models"
	"gorm.io/gorm"
)

// MigrateUsers migre la table users
func MigrateUsers(db *gorm.DB) error {
	err := db.AutoMigrate(&models.User{}) // Utilisation de models.User au lieu de créer une autre structure
	return err
}
