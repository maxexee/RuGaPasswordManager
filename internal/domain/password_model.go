package domain

import "gorm.io/gorm"

type Password struct {
	gorm.Model
	Name                    string `gorm:"not null;uniqueIndex:idx_user_name_password"`
	Description             *string
	Password                string `gorm:"not null"`
	UserID                  uint   `gorm:"not null;uniqueIndex:idx_user_name_password"`
	SectionParentIdPassword uint   `gorm:"not null"`
}
