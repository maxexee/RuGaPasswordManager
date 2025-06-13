package domain

import "gorm.io/gorm"

type Password struct {
	gorm.Model
	Name                    string `gorm:"unique;not null"`
	Description             *string
	Password                string `gorm:"not null"`
	UserID                  uint   `gorm:"not null"`
	SectionParentIdPassword uint   `gorm:"not null"`
}
