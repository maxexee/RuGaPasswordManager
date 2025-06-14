package domain

import "gorm.io/gorm"

type Section struct {
	gorm.Model
	Name             string `gorm:"unique;not null"`
	Description      *string
	UserID           uint `gorm:"not null"`
	SectionParentId  *uint
	SectionChildren  []Section  `gorm:"foreignKey:SectionParentId;constraint:OnDelete:CASCADE"`
	PasswordChildren []Password `gorm:"foreignKey:SectionParentIdPassword;constraint:OnDelete:CASCADE"`
}
