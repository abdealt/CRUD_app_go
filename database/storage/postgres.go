package storage

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Config est la structure de données pour la configuration pour permettre
// la connexion à la base de données
type Config struct {
	Host     string
	Port     string
	Password string
	User     string
	DbName   string
	SSLMode  string
}

// NewConnection crée une connexion à la base de données
// Elle prend une instance de Config en argument (config)
// et retourne une erreur si une erreur survient lors de la connexion
func NewConnection(config Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", config.Host, config.Port, config.User, config.Password, config.DbName, config.SSLMode)
	// Création d'une connexion à la base de données
	// Le second argument est un pointeur vers une structure Config
	// qui permet de configurer GORM
	// Retourne une erreur si une erreur survient lors de la connexion
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return db, err
	}
	return db, nil
}
