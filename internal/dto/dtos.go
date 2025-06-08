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
// DTO PARA UNA CONTRASEÑA, TODOS LOS CAMPOS.
type Password struct {
	ID                      uint
	CreatedAt               time.Time
	UpdatedAt               time.Time
	Name                    string
	Description             *string
	Password                string
	SectionParentIdPassword uint
}

// ===========================================================================================
// ===========================================================================================
// =========================================== DTOs - SECCIONES===============================
// DTO PARA UNA SECCION, TODOS LOS CAMPOS.
type SectionDto struct {
	// UpdatedAt        time.Time    `json:"updatedAt,omitempty"`
	ID               uint         `json:"id,omitempty"`
	CreatedAt        time.Time    `json:"createdAt,omitempty"`
	Name             string       `json:"name"`
	Description      *string      `json:"description,omitempty"`
	UserID           uint         `json:"userId,omitempty"`
	SectionParentId  *uint        `json:"sectionParentId"`
	SectionChildren  []SectionDto `json:"children,omitempty"`
	PasswordChildren []Password   `json:"passwords,omitempty"`
}

// -- NU - PENDIENTE --
// DTO QUE RECIBE LAS SECCIONES DEL REPOSITORY Y LAS MANDA AL USECASE.
type SectionGetSliceDTO struct {
	SectionSliceReturn []SectionDto
}
