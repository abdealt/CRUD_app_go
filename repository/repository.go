package repository

import (
	"gorm.io/gorm" // Import de GORM, l'ORM utilisé pour interagir avec la base de données
)

// Repository est une structure qui sert de couche d'abstraction pour accéder à la base de données.
// Elle contient un attribut DB qui est un pointeur vers l'instance de GORM.
// L'objectif est de centraliser les interactions avec la base de données au sein de cette structure.
type Repository struct {
	DB *gorm.DB // DB est la connexion à la base de données. Ce pointeur vers gorm.DB permet d'exécuter des requêtes SQL via GORM.
}
