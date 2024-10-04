package models

// User représente la structure de l'utilisateur dans l'application.
// Cette struct sera utilisée pour mapper les données d'un utilisateur
// aux tables de la base de données via GORM.
type User struct {
	ID      uint   `gorm:"primary_key" json:"id"`                          // ID est l'identifiant unique de l'utilisateur. GORM l'identifie comme la clé primaire (primary key).
	Name    string `json:"name" gorm:"type:varchar(255);not null"`         // Name est le nom de l'utilisateur. Il est de type string, avec une longueur maximale de 255 caractères, et ne peut pas être null (not null).
	Email   string `json:"email" gorm:"type:varchar(255);unique;not null"` // Email est l'adresse email de l'utilisateur. Elle doit être unique et ne peut pas être null. Sa longueur maximale est également de 255 caractères.
	City    string `json:"city" gorm:"type:varchar(255);not null"`         // City est la ville de résidence de l'utilisateur. Elle est de type string avec une longueur maximale de 255 caractères et ne peut pas être null.
	Country string `json:"country" gorm:"type:varchar(255);not null"`      // Country est le pays de résidence de l'utilisateur. Elle suit les mêmes restrictions que le champ City.
}
