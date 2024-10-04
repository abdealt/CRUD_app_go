package repository // Déclaration du package, ici nommé "repository"

import (
	// Importation des packages nécessaires

	"github.com/abdealt/go_crud/database/models" // Import du package models pour accéder à la structure User
	"gopkg.in/go-playground/validator.v9"        // Import du package validator pour valider les données des utilisateurs
)

// Créer une instance du validateur
var validate *validator.Validate = validator.New() // Initialisation d'un validateur pour la validation des structures

// Structure pour représenter une réponse d'erreur de validation
type ErrorResponse struct {
	FailedField string // Champ ayant échoué la validation
	Tag         string // Tag de validation qui a échoué (par exemple, "required", "email", etc.)
	Value       string // Valeur qui a échoué la validation
}

// Fonction pour valider un utilisateur
func ValidateStruct(user models.User) []ErrorResponse {
	var errors []ErrorResponse   // Déclaration d'un tableau pour stocker les erreurs de validation
	err := validate.Struct(user) // Validation de la structure User

	if err != nil { // Vérification si des erreurs de validation se sont produites
		for _, err := range err.(validator.ValidationErrors) { // Boucle sur les erreurs de validation
			element := ErrorResponse{
				FailedField: err.Field(), // Récupération du champ ayant échoué
				Tag:         err.Tag(),   // Récupération du tag d'erreur
				Value:       err.Param(), // Récupération de la valeur de paramètre associée à l'erreur
			}
			errors = append(errors, element) // Ajoute l'erreur à la liste des erreurs
		}
	}
	return errors // Retourne la liste d'erreurs, vide si aucune erreur
}
