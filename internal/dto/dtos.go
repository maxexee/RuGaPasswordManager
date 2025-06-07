package dto

// DTOS PARA EL REGISTRO Y LOGIN DEL USUARIO.
type SignUpDTO struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=12,passwordsFormat"`
}

type LogInDTO struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

// DTOS PARA LAS SECCIONES.
type SectionGetAllDTO struct {
	UserId    uint  `json:"userId" validate:"required,numeric"`
	SectionId *uint `json:"sectionId" validate:"required,alpha"`
}
