package migrations

import (
	"github.com/abdealt/go_crud/database/models" // Import du package models qui contient la structure User
	"gorm.io/gorm"                               // Import de GORM, l'ORM (Object-Relational Mapping) utilisé pour interagir avec la base de données
)

// MigrateUsers migre la table users
// La fonction MigrateUsers prend en paramètre une instance de gorm.DB, qui représente une connexion à la base de données.
// Elle s'assure que la table `users` existe et est bien synchronisée avec la structure de données User définie dans le package models.
// Si la table n'existe pas, elle est automatiquement créée, sinon elle est modifiée pour correspondre à la structure User.
func MigrateUsers(db *gorm.DB) error {
	// La méthode AutoMigrate de GORM est utilisée pour migrer la table `users` en fonction de la structure `User` définie dans le package models.
	// Cela permet de créer ou de mettre à jour la table `users` dans la base de données pour correspondre à la structure du modèle.
	err := db.AutoMigrate(&models.User{}) // Utilisation de models.User au lieu de créer une autre structure

	// Si une erreur survient lors de la migration, elle est retournée par la fonction, sinon err sera nil, indiquant que la migration s'est bien déroulée.
	return err
}
