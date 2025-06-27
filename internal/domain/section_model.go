package domain

import "gorm.io/gorm"

/*
	- index:idx_user_name,unique
	-- PERMITE QUE POR MEDIO DE UN INDEX COMPUESTO (YA QUE APARECE EN MAS DE UN CAMPO) Y LA PALABRA UNIQUE
	-- HACER QUE LOS CAMPOS EN DONDE ESTA ESPECIFICADO ESE INDEX, SEA UNICOS, PERO SOLO PARA EL USUARIO CON
	-- EL MISMO ID, ES DECIR, NO PERMITE QUE HAYA NOMBRES IGUALES SIEMPRE Y CUANDO EL ID SEA EL MISMO, SI NO
	-- ES EL MISMO ID, ENTONCES SI LO PERMITE.


	- constraint:OnDelete:CASCADE
	-- PERMITE QUE MEDIENTE UN CONSTRAINT, SI YO ELIMINO UNA SECCION, TODO LO QUE ESTE DENTRO DE ELLA TAMBIEN
	-- SE ELIMINARA.
*/

type Section struct {
	gorm.Model
	Name             string `gorm:"not null;index:idx_user_name,unique"`
	Description      *string
	UserID           uint `gorm:"not null;index:idx_user_name,unique"`
	SectionParentId  *uint
	SectionChildren  []Section  `gorm:"foreignKey:SectionParentId;constraint:OnDelete:CASCADE"`
	PasswordChildren []Password `gorm:"foreignKey:SectionParentIdPassword;constraint:OnDelete:CASCADE"`
}
