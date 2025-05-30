package domain

import "gorm.io/gorm"

// STRUCT - MODELO - DE LOS USUARIOS - ES LA REPRESENTACION DE LA TABLA EN LA BD.
type User struct {
	gorm.Model
	Email     string `gorm:"not null;unique"`
	Passsword string `gorm:"not null"`
	Sections  []Section
}
