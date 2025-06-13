package passwordsrepository

import (
	"strings"

	postgres "github.com/maxexee/rugaPasswordManager/infrastructure/db"
	"github.com/maxexee/rugaPasswordManager/internal/domain"
	"github.com/maxexee/rugaPasswordManager/internal/dto"
)

// VERDE...
func PasswordGetByIdUseCase(password *dto.PasswordDto) (bool, *dto.PasswordDto, error) {
	// OBJETO DE TIPO *domain.User*.
	var userExist domain.User

	// OBJETO DE TIPO *domain.Password*.
	var passwordExist domain.Password

	// ===========================================================================================
	// ===========================================================================================
	// =========================================== VALIDACIONES ==================================
	// VALIDAMOS QUE EL ID DEL USUARIO EXISTA.
	userExistResult := postgres.DB.First(&userExist, "id	=	?", password.UserID)
	if userExistResult.Error != nil {
		return false, nil, userExistResult.Error
	}

	// ===========================================================================================
	// ===========================================================================================
	// =========================================== QUERY =========================================
	// BUSQUEDA Y OTENCION DE LA CONTRASEÑA EN LA BASE DE DATOS MEDIANTE EL NOMBRE.
	passwordExistResult := postgres.DB.First(&passwordExist, password.ID)
	if passwordExistResult.Error != nil {
		return false, nil, passwordExistResult.Error
	}

	// CONSTRUCCION DEL DTO DE RETORNO.
	dtoReturn := dto.PasswordDto{
		ID:                      passwordExist.ID,
		Name:                    passwordExist.Name,
		Description:             passwordExist.Description,
		Password:                passwordExist.Password,
		UserID:                  passwordExist.UserID,
		SectionParentIdPassword: passwordExist.SectionParentIdPassword,
	}

	// SI TODO SALE BIEN...
	return true, &dtoReturn, nil
}

// VERDE...
func PasswordGetByNameRepository(password *dto.PasswordDto) (bool, *dto.PasswordDto, error) {
	// OBJETO DE TIPO *domain.User*.
	var userExist domain.User

	// OBJETO DE TIPO	*domain.Password*
	var passwordExist domain.Password

	// ===========================================================================================
	// ===========================================================================================
	// =========================================== VALIDACIONES ==================================
	// VALIDAMOS QUE EL ID DEL USUARIO EXISTA.
	userExistResult := postgres.DB.First(&userExist, "id	=	?", password.UserID)
	if userExistResult.Error != nil {
		return false, nil, userExistResult.Error
	}

	// ===========================================================================================
	// ===========================================================================================
	// =========================================== QUERY =========================================
	// BUSQUEDA Y OTENCION DE LA CONTRASEÑA EN LA BASE DE DATOS MEDIANTE EL NOMBRE.
	passwordExistResult := postgres.DB.Where("name	=	?", password.Name).First(&passwordExist)
	if passwordExistResult.Error != nil {
		return false, nil, passwordExistResult.Error
	}

	if passwordExistResult.RowsAffected == 0 {
		return false, nil, passwordExistResult.Error
	}

	// CONSTRUCCION DEL DTO DE RETORNO.
	dtoReturn := dto.PasswordDto{
		ID:                      passwordExist.ID,
		Name:                    passwordExist.Name,
		Description:             passwordExist.Description,
		Password:                passwordExist.Password,
		UserID:                  passwordExist.UserID,
		SectionParentIdPassword: passwordExist.SectionParentIdPassword,
	}

	// SI TODO SALE BIEN...
	return true, &dtoReturn, nil
}

// VERDE...
func PasswordPostRepository(password *dto.PasswordDto) (bool, *dto.PasswordDto, error) {
	// OBJETO DE TIPO *domain.User*.
	var userExist domain.User

	// OBJETO DE TIPO *domain.Section*
	var sectionExist domain.Section

	// ===========================================================================================
	// ===========================================================================================
	// =========================================== VALIDACIONES ==================================
	// VALIDAMOS QUE EL ID DEL USUARIO EXISTA.
	userExistResult := postgres.DB.First(&userExist, "id	=	?", password.UserID)
	if userExistResult.Error != nil {
		return false, nil, userExistResult.Error
	}

	// VALIDAMOS QUE LA SECCION PADRE EXISTA (NO SE ACEPTAN CONSTRASEÑAS EN LA SECCION RAIZ (O NULL)).
	sectionExistResult := postgres.DB.Where("id	=	?", password.SectionParentIdPassword).Select("id").First(&sectionExist)
	if sectionExist.ID == 0 || sectionExistResult.Error != nil {
		return false, nil, sectionExistResult.Error
	}

	// ===========================================================================================
	// ===========================================================================================
	// =========================================== QUERY =========================================
	// CONSTRUCCION DE LA NUEVA CONTRASEÑA.
	passwordCreate := domain.Password{
		UserID:                  password.UserID,
		SectionParentIdPassword: password.SectionParentIdPassword,
		Name:                    strings.ToUpper(password.Name),
		Description:             password.Description,
		Password:                password.Password,
	}

	// GUARDADO DE LA CONTRASEÑA EN LA BASE DE DATOS.
	passwordCreation := postgres.DB.Create(&passwordCreate)
	if passwordCreation.Error != nil {
		return false, nil, passwordCreation.Error
	}

	// CONSTRUCCION DEL DTO DE RETONO.
	dtoReturn := dto.PasswordDto{
		ID:                      password.ID,
		UserID:                  password.UserID,
		CreatedAt:               password.CreatedAt,
		SectionParentIdPassword: password.SectionParentIdPassword,
		Name:                    password.Name,
		Description:             password.Description,
		Password:                password.Password,
	}

	// SI TODO SALE BIEN...
	return true, &dtoReturn, nil
}

// VERDE...
func PasswordUpdateRepositoy(password *dto.PasswordDto) (bool, *dto.PasswordDto, error) {
	// OBJETO DE TIPO *domain.User*.
	var userExist domain.User

	// OBJETO DE TIPO *domain.Password*
	var passwordExist domain.Password

	// ===========================================================================================
	// ===========================================================================================
	// =========================================== VALIDACIONES ==================================
	// VALIDAMOS QUE EL ID DEL USUARIO EXISTA.
	userExistResult := postgres.DB.First(&userExist, "id	=	?", password.UserID)
	if userExistResult.Error != nil {
		return false, nil, userExistResult.Error
	}

	// VALIDAMOS QUE EL ID DEL LA CONTRASEÑA EXISTA.
	passwordExistResult := postgres.DB.First(&passwordExist, "id	=	?", password.ID)
	if passwordExistResult.Error != nil {
		return false, nil, passwordExistResult.Error
	}

	// ===========================================================================================
	// ===========================================================================================
	// =========================================== QUERY =========================================
	// CONSTRUCCION Y GUARDADO DE LA CONSTRASEÑA A ACTUALIZAR.
	passwordUpdateResult := postgres.DB.Model(&passwordExist).Where("id	=	?", password.ID).Updates(map[string]interface{}{
		"Name":                    strings.ToUpper(password.Name),
		"Description":             &password.Description,
		"Password":                password.Password,
		"SectionParentIdPassword": password.SectionParentIdPassword,
	})
	if passwordUpdateResult.Error != nil {
		return false, nil, passwordUpdateResult.Error
	}

	if passwordUpdateResult.RowsAffected == 0 {
		return false, nil, passwordUpdateResult.Error
	}

	// CONSTRUCCION DEL DTO DE RETORNO.
	dtoReturn := dto.PasswordDto{
		ID:                      passwordExist.ID,
		UserID:                  passwordExist.UserID,
		SectionParentIdPassword: passwordExist.SectionParentIdPassword,
		CreatedAt:               passwordExist.CreatedAt,
		UpdatedAt:               passwordExist.UpdatedAt,
		Name:                    passwordExist.Name,
		Description:             passwordExist.Description,
		Password:                userExist.Passsword,
	}

	// SI TODO SALE BIEN...
	return true, &dtoReturn, nil
}

// VERDE...
func PasswordDeleteRepository(password *dto.PasswordDto) (bool, error) {
	// OBJETO DE TIPO *domain.User*.
	var userExist domain.User

	// OBJETO DE TIPO *domain.Password*
	var passwordExist domain.Password

	// ===========================================================================================
	// ===========================================================================================
	// =========================================== VALIDACIONES ==================================
	// VALIDAMOS QUE EL ID DEL USUARIO EXISTA.
	userExistResult := postgres.DB.First(&userExist, "id	=	?", password.UserID)
	if userExistResult.Error != nil {
		return false, userExistResult.Error
	}

	// VALIDAMOS QUE EL ID DEL LA SECCION EXISTA.
	passwordExistResult := postgres.DB.First(&passwordExist, "id	=	?", password.ID)
	if passwordExistResult.Error != nil {
		return false, passwordExistResult.Error
	}

	// ===========================================================================================
	// ===========================================================================================
	// =========================================== QUERY =========================================
	// ELIMINAMOS EL REGISTRO DE LA BASE DE DATOS.
	passwordDeleteResult := postgres.DB.Unscoped().Delete(&passwordExist, password.ID)
	if passwordDeleteResult.Error != nil {
		return false, passwordDeleteResult.Error
	}

	if passwordDeleteResult.RowsAffected == 0 {
		return false, passwordDeleteResult.Error
	}

	// SI TODO SALE BIEN.....
	return true, nil
}
