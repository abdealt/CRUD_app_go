package storage

import (
	"fmt"

	"gorm.io/driver/postgres" // Import du pilote PostgreSQL pour GORM
	"gorm.io/gorm"            // Import de GORM, un ORM (Object Relational Mapping) pour Go
)

// Config est une structure qui stocke les informations nécessaires à la connexion
// à la base de données. Chaque champ correspond à une variable de configuration.
type Config struct {
	Host     string // Adresse du serveur de la base de données (ex: localhost ou une IP distante)
	Port     string // Port sur lequel la base de données écoute (ex: 5432 pour PostgreSQL)
	Password string // Mot de passe de l'utilisateur pour accéder à la base de données
	User     string // Nom de l'utilisateur de la base de données
	DbName   string // Nom de la base de données à laquelle se connecter
	SSLMode  string // Mode SSL pour la connexion (peut être disable, require, etc.)
}

// NewConnection crée une connexion à la base de données PostgreSQL.
// Elle prend en argument une instance de la structure Config qui contient
// les paramètres de connexion. Si la connexion échoue, elle renvoie une erreur.
func NewConnection(config Config) (*gorm.DB, error) {
	// Construction de la Data Source Name (DSN) pour PostgreSQL avec les informations de Config
	// fmt.Sprintf permet de formater la chaîne de connexion avec les variables appropriées
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		config.Host, config.Port, config.User, config.Password, config.DbName, config.SSLMode)

	// Ouverture de la connexion à la base de données PostgreSQL avec GORM
	// postgres.Open(dsn) spécifie que GORM doit utiliser PostgreSQL comme base de données
	// gorm.Config{} contient la configuration pour GORM (peut être enrichie selon les besoins)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		// Si la connexion échoue, retourne l'instance db (même en cas d'erreur) et l'erreur associée
		return db, err
	}

	// Si la connexion réussit, retourne l'instance db et nil comme erreur
	return db, nil
}
