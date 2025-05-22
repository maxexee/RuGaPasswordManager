package domain

import "gorm.io/gorm"

// STRUCT - MODELO - DE LOS USUARIOS - ES LA REPRESENTACION DE LA TABLA EN LA BD.
type User struct {
	gorm.Model
	Email     string `gorm:"unique"`
	Passsword string
	Sections  []Section
}
