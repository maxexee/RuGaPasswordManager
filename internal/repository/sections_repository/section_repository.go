package sectionsrepository

import (
	"errors"
	"fmt"
	"strings"

	postgres "github.com/maxexee/rugaPasswordManager/infrastructure/db"
	"github.com/maxexee/rugaPasswordManager/internal/domain"
	"github.com/maxexee/rugaPasswordManager/internal/dto"
	"gorm.io/gorm"
)

// VERDE...
func SectionGetAllRepository(section *dto.SectionDto) (bool, *dto.SectionPasswordGetSliceDTO, error) {
	// OBJETO DE TIPO *domain.User*
	var userExist domain.User

	// OBJETO DE TIPO *[]domain.Section*
	var sectionsExist []domain.Section

	// OBJETO DE *[]domain.Password*
	var passwordExist []domain.Password

	// ===========================================================================================
	// ===========================================================================================
	// =========================================== VALIDACIONES ==================================
	// VALIDAMOS QUE EL USUARIO EXISTA.
	userExistResutlt := postgres.DB.First(&userExist, "id	=	?", section.UserID)
	if userExist.ID == 0 || userExistResutlt.Error != nil {
		return false, nil, userExistResutlt.Error
	}

	// ===========================================================================================
	// ===========================================================================================
	// =========================================== QUERY =========================================
	// TRAEMOS LAS SECCIONES HIJAS DE LA RAIZ.
	if section.SectionParentId == nil {
		sectionsGetAllResult := postgres.DB.Where("user_id = ?	AND	section_parent_id	IS NULL", section.UserID).Find(&sectionsExist)
		if sectionsGetAllResult.Error != nil {
			return false, nil, sectionsGetAllResult.Error
		}
	} else {
		// VALIDAR SI EXISTE EL PADRE.
		sectionExistResult := postgres.DB.First(&sectionsExist, "id	=	?", section.SectionParentId)
		if sectionExistResult.Error != nil {
			return false, nil, sectionExistResult.Error
		}

		// OBTENCION DE LAS SECCIONES HIJAS DE UNA SECCION PADRE.
		sectionsGetAllResult := postgres.DB.Where("user_id	=	?	AND	section_parent_id	=	?", section.UserID, section.SectionParentId).Find(&sectionsExist)
		if sectionsGetAllResult.Error != nil {
			return false, nil, sectionsGetAllResult.Error
		}

		// OBTENCION DE LAS CONSTRASEÑAS HIJAS DE UNA SECCION PADRE.
		posswordGetResult := postgres.DB.Where("user_id	=	?	AND section_parent_id_password	=	?", section.UserID, section.SectionParentId).Find(&passwordExist)
		if posswordGetResult.Error != nil {
			return false, nil, posswordGetResult.Error
		}

		// SI NINGUNA SECCION "Y" CONTRASEÑA FUERON ENCONTRADAS, RETORNA UN ERROR.
		if sectionsGetAllResult.RowsAffected == 0 && posswordGetResult.RowsAffected == 0 {
			fmt.Println("ERROR NOT FOUND...")
			return false, nil, gorm.ErrRecordNotFound
		}
	}

	// CONSTRUCCION DEL DTO PARA EL REGRESO DE LAS SECCIONES.
	sectionsDTOsReturn := make([]dto.SectionDto, len(sectionsExist))
	for i, section := range sectionsExist {
		sectionsDTOsReturn[i] = dto.SectionDto{
			SectionParentId: section.SectionParentId,
			UserID:          section.UserID,
			ID:              section.ID,
			CreatedAt:       section.CreatedAt,
			Name:            section.Name,
			Description:     section.Description,
		}
	}

	// CONTRUCCION DEL DTO PARA EL REGRESO DE LAS CONTRASEÑAS.
	passwordsDTOsReturn := make([]dto.PasswordDto, len(passwordExist))
	for i, password := range passwordExist {
		passwordsDTOsReturn[i] = dto.PasswordDto{
			ID:                      password.ID,
			Name:                    password.Name,
			Description:             password.Description,
			Password:                password.Password,
			UserID:                  password.UserID,
			SectionParentIdPassword: password.SectionParentIdPassword,
			CreatedAt:               password.CreatedAt,
			UpdatedAt:               password.UpdatedAt,
		}
	}

	// CONSTRUCION DEL DTO A RETORNAR.
	dtoReturn := dto.SectionPasswordGetSliceDTO{
		SectionSliceReturn:  sectionsDTOsReturn,
		PasswordSliceReturn: passwordsDTOsReturn,
	}

	// SI TODO SALE BIEN...
	return true, &dtoReturn, nil
}

// VERDE...
func SectionGetByNameRepository(section *dto.SectionDto) (bool, *dto.SectionDto, error) {
	// OBJETO DE TIPO *domain.User*
	var userExist domain.User

	// OBJETO DE TIPO *[]domain.Section*
	var sectionExist domain.Section

	// ===========================================================================================
	// ===========================================================================================
	// =========================================== VALIDACIONES ==================================
	// VALIDAMOS QUE EL ID DEL USUARIO EXISTA.
	userExistResult := postgres.DB.First(&userExist, "id	=	?", section.UserID)
	if userExist.ID == 0 || userExistResult.Error != nil {
		return false, nil, userExistResult.Error
	}

	// ===========================================================================================
	// ===========================================================================================
	// =========================================== QUERY =========================================
	// BUSQUEDA Y OTENCION DE LA SECCION EN LA BASE DE DATOS MEDIANTE EL NOMBRE.
	sectionGetResult := postgres.DB.Where("user_id	=	?	AND name	=	?", section.UserID, section.Name).First(&sectionExist)
	if sectionGetResult.Error != nil {
		return false, nil, sectionGetResult.Error
	}

	// CONSTRUCION DEL DTO A RETORNAR.
	dtoReturn := dto.SectionDto{
		ID:              sectionExist.ID,
		CreatedAt:       sectionExist.CreatedAt,
		Name:            sectionExist.Name,
		Description:     sectionExist.Description,
		UserID:          sectionExist.UserID,
		SectionParentId: sectionExist.SectionParentId,
	}

	// SI TODO SALE BIEN...
	return true, &dtoReturn, nil
}

// VERDE...
func SectionPostRepository(section *dto.SectionDto) (bool, *dto.SectionDto, error) {
	// OBJETO DE TIPO *domain.User*.
	var userExist domain.User

	// OBJETO DE TIPO *domain.Section*.
	var parentSectionExist domain.Section

	// ===========================================================================================
	// ===========================================================================================
	// =========================================== VALIDACIONES ==================================

	// VALIDAMOS QUE EL ID DEL USUARIO EXISTA.
	userExistResult := postgres.DB.First(&userExist, "id	=	?", section.UserID)
	if userExist.ID == 0 || userExistResult.Error != nil {
		return false, nil, userExistResult.Error
	}

	// VALIDAMOS QUE EL *SectionParentId* EXISTA, SI NO, QUE EL  *SectionParentId* SEA NULL.
	if section.SectionParentId != nil {
		parentSectionExistResult := postgres.DB.First(&parentSectionExist, "user_id	=	?	AND id	=	?", section.UserID, section.SectionParentId)
		if parentSectionExist.ID == 0 || parentSectionExistResult.Error != nil {
			return false, nil, parentSectionExistResult.Error
		}
	}

	// ===========================================================================================
	// ===========================================================================================
	// ================================== QUERY ==================================================
	// CREACION DE LA SECCION MEDIANTE EL MODELO *Section*.
	sectionCreate := domain.Section{
		Name:            strings.ToUpper(section.Name),
		Description:     section.Description,
		UserID:          section.UserID,
		SectionParentId: section.SectionParentId,
	}

	// GUARDADO DE LA NUEVA SECCION EN LA BASE DE DATOS.
	sectionCreateResult := postgres.DB.Create(&sectionCreate)
	if sectionCreateResult.Error != nil {
		return false, nil, sectionCreateResult.Error
	}

	if sectionCreateResult.RowsAffected == 0 {
		return false, nil, errors.New("error al guardar en la base de datos")
	}

	// CONSTRUCCION DEL DTO DE RETORNO.
	dtoReturn := dto.SectionDto{
		ID:              sectionCreate.ID,
		CreatedAt:       sectionCreate.CreatedAt,
		Name:            sectionCreate.Name,
		Description:     sectionCreate.Description,
		UserID:          sectionCreate.UserID,
		SectionParentId: sectionCreate.SectionParentId,
	}

	// SI TODO SALE BIEN...
	return true, &dtoReturn, nil
}

// VERDE...
func SectionUpdateRepository(section *dto.SectionDto) (bool, *dto.SectionDto, error) {
	// OBJETO DE USUARIO.
	var userExist domain.User

	// OBJETOS DE TIPO *domain.Section*.
	var sectionExist domain.Section

	// OBJETO DE TIPO *domain.Section*.
	var parentSectionExist domain.Section

	// ===========================================================================================
	// ===========================================================================================
	// =========================================== VALIDACIONES ==================================
	// VALIDAMOS QUE EL USUARIO EXISTA.
	userExistResutlt := postgres.DB.First(&userExist, "id	=	?", section.UserID)
	if userExist.ID == 0 || userExistResutlt.Error != nil {
		return false, nil, userExistResutlt.Error
	}

	// VALIDAMOS QUE EL *SectionParentId* EXISTA, SI NO, QUE EL  *SectionParentId* SEA NULL.
	if section.SectionParentId != nil {
		parentSectionExistResult := postgres.DB.First(&parentSectionExist, " user_id	=	?	AND	id	=	?", section.UserID, section.SectionParentId)
		if parentSectionExist.ID == 0 || parentSectionExistResult.Error != nil {
			return false, nil, parentSectionExistResult.Error
		}
	}

	// VALIDAMOS QUE LA SECCION EXISTA.
	sectionExistResult := postgres.DB.First(&sectionExist, "user_id	=	?	AND	id	=	?", section.UserID, section.ID)
	if sectionExistResult.Error != nil {
		return false, nil, sectionExistResult.Error
	}

	// ===========================================================================================
	// ===========================================================================================
	// =========================================== QUERY =========================================
	// MODIFICACION DE LA SECCION.
	sectionUpdateResult := postgres.DB.Model(&sectionExist).Where("user_id	=	?	AND	id	=	?", section.UserID, section.ID).Updates(map[string]interface{}{
		"Name":            strings.ToUpper(section.Name),
		"Description":     section.Description,
		"SectionParentId": section.SectionParentId,
	})

	if sectionUpdateResult.Error != nil {
		return false, nil, sectionUpdateResult.Error
	}

	if sectionUpdateResult.RowsAffected == 0 {
		return false, nil, errors.New("update section error")
	}

	// CONSTRUCCION DEL DTO DE RETORNO.
	dtoReturn := dto.SectionDto{
		ID:              sectionExist.ID,
		CreatedAt:       sectionExist.CreatedAt,
		UpdatedAt:       sectionExist.UpdatedAt,
		Name:            sectionExist.Name,
		Description:     sectionExist.Description,
		UserID:          sectionExist.UserID,
		SectionParentId: sectionExist.SectionParentId,
	}

	// SI TODO SALE BIEN.
	return true, &dtoReturn, nil
}

// VERDE...
func SectionDeleteRepository(section *dto.SectionDto) (bool, error) {
	// OBJETO DE TIPO *domain.User*
	var userExist domain.User
	//
	// OBJETO DE TIPO *domain.Section*
	var sectionExist domain.Section

	// ===========================================================================================
	// ===========================================================================================
	// =========================================== VALIDACIONES ==================================
	// VALIDAMOS QUE EL USUARIO EXISTA.
	userExistResult := postgres.DB.First(&userExist, section.UserID)
	if userExistResult.Error != nil {
		return false, userExistResult.Error
	}

	// VALIDAMOS QUE LA SECCION EXISTA.
	sectionExistResult := postgres.DB.First(&sectionExist, "user_id	=	?	AND	id	=	?", section.UserID, section.ID)
	if sectionExistResult.Error != nil {
		fmt.Println("1...")
		return false, sectionExistResult.Error
	}

	// ===========================================================================================
	// ===========================================================================================
	// =========================================== QUERY =========================================
	//ELIMINACION DE LA SECCION EN LA BASE DE DATO.
	sectionDeleteResult := postgres.DB.Unscoped().Delete(&sectionExist, "user_id	=	?	AND	id	=	?", section.UserID, section.ID)
	if sectionDeleteResult.Error != nil {
		return false, sectionDeleteResult.Error
	}

	// SI TODO SALE BIEN...
	return true, nil
}
