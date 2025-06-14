package dto

import (
	"time"
)

// ===========================================================================================
// ===========================================================================================
// =========================================== DTOs - USUARIO ================================
// DTO PARA EL REGISTRO.
type SignUpDTO struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=12,passwordsFormat"`
}

// DTO PARA EL LOGIN DEL USUARIO.
type LogInDTO struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

// ===========================================================================================
// ===========================================================================================
// =========================================== DTOs - CONTRASEÑA =============================
// DTO PARA UNA CONTRASEÑA, TODOS LOS CAMPOS PERO PARA RECIBIR.
type PasswordBodyDto struct {
	Name                    string  `json:"name" validate:"required,min=2,max=20,matchesName=^[A-Za-z0-9 ]+$"`
	Description             *string `json:"description,omitempty" validate:"max=50"`
	Password                string  `json:"password" validate:"required,min=12"`
	SectionParentIdPassword uint    `json:"sectionparentidpassword" validate:"number,required"`
}

// DTO PARA UNA CONTRASEÑA, TODOS LOS CAMPOS.
type PasswordDto struct {
	ID                      uint      `json:"id,omitempty"`
	CreatedAt               time.Time `json:"createdAt,omitempty"`
	UpdatedAt               time.Time `json:"updatedAt,omitempty"`
	Name                    string    `json:"name"`
	Description             *string   `json:"description,omitempty"`
	Password                string    `json:"password"`
	SectionParentIdPassword uint      `json:"sectionParentIdPassword"`
	UserID                  uint      `json:"userId,omitempty"`
}

// ===========================================================================================
// ===========================================================================================
// =========================================== DTOs - SECCIONES ===============================
// DTO PARA UNA SECCION, TODOS LOS CAMPOS A RECIBIR, ES DECIR, EL BODY.

// DTO PARA UNA SECCION, TODOS LOS CAMPOS.
type SectionDto struct {
	// UpdatedAt        time.Time    `json:"updatedAt,omitempty"`
	ID               uint          `json:"id,omitempty"`
	CreatedAt        time.Time     `json:"createdAt,omitempty"`
	Name             string        `json:"name" validate:"required,min=2,max=20,matchesName=^[A-Za-z0-9 ]+$"`
	Description      *string       `json:"description,omitempty" validate:"max=50"`
	UserID           uint          `json:"userId,omitempty"`
	SectionParentId  *uint         `json:"sectionParentId"`
	SectionChildren  []SectionDto  `json:"children,omitempty"`
	PasswordChildren []PasswordDto `json:"passwords,omitempty"`
}

// -- NU - PENDIENTE --
// DTO QUE RECIBE LAS SECCIONES DEL REPOSITORY -> USECASE.
type SectionGetSliceDTO struct {
	SectionSliceReturn  []SectionDto
	PasswordSliceReturn []PasswordDto
}
