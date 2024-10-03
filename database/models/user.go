package models

// type User struct {
// 	Name    string `json:"name" validate:"required", min=3,max=40"`
// 	Email   string `json:"email" validate:"required,email",min=6,max=60"`
// 	City    string `json:"city" validate:"required", min=3,max=40"`
// 	Country string `json:"country" validate:"required", min=3,max=40"`
// }
//Apres
type User struct {
	ID      uint   `gorm:"primary_key" json:"id"`
	Name    string `json:"name" gorm:"type:varchar(255);not null"`
	Email   string `json:"email" gorm:"type:varchar(255);unique;not null"`
	City    string `json:"city" gorm:"type:varchar(255);not null"`
	Country string `json:"country" gorm:"type:varchar(255);not null"`
}
