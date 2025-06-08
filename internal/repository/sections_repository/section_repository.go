package sectionsrepository

import (
	"strings"

	postgres "github.com/maxexee/rugaPasswordManager/infrastructure/db"
	"github.com/maxexee/rugaPasswordManager/internal/domain"
	"github.com/maxexee/rugaPasswordManager/internal/dto"
)

// VERDE
func SectionGetAllRepository(section *dto.SectionDto) (bool, *dto.SectionGetSliceDTO, error) {
	// OBJETO DE TIPO *domain.User*
	var userExist domain.User

	// OBJETO DE TIPO *[]domain.Section*
	var sectionsExist []domain.Section
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
	// TRAEMOS LAS SECCIONES DE LA RAIZ O DE UNA SECCION PADRE.
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

		// OBTENCION DE LAS SECCIONES MEDIANTE UN PADRE.
		sectionsGetAllResult := postgres.DB.Where("user_id	=	?	AND	section_parent_id	=	?", section.UserID, section.SectionParentId).Find(&sectionsExist)
		if sectionsGetAllResult.Error != nil {
			return false, nil, sectionExistResult.Error
		}
	}

	// CONSTRUCCION DEL DTO PARA EL REGRESO DE LAS SECCIONES.
	dtos := make([]dto.SectionDto, len(sectionsExist))

	for i, section := range sectionsExist {
		dtos[i] = dto.SectionDto{
			SectionParentId: section.SectionParentId,
			UserID:          section.UserID,
			ID:              section.ID,
			CreatedAt:       section.CreatedAt,
			Name:            section.Name,
			Description:     section.Description,
			// UpdatedAt:       section.UpdatedAt,
			// SectionChildren: mapChildren(s.SectionChildren),
			// PasswordChildren: mapPasswords(s.PasswordChildren),
		}
	}

	// CONSTRUCION DEL DTO A RETORNAR.
	dtoReturn := dto.SectionGetSliceDTO{
		SectionSliceReturn: dtos,
	}

	// SI TODO SALE BIEN...
	return true, &dtoReturn, nil
}

// VERDE
func SectionGetByNameRepository(section *dto.SectionDto, sectionNameQuery string) (bool, *dto.SectionDto, error) {
	// OBJETO DE TIPO *domain.User*
	var userExist domain.User

	// OBJETO DE TIPO *[]domain.Section*
	var sectionExist domain.Section

	// ===========================================================================================
	// ===========================================================================================
	// =========================================== VALIDACIONES ==================================
	// VALIDAMOS QUE EL USUARIO EXISTA.
	userExistResult := postgres.DB.First(&userExist, "id	=	?", section.UserID)
	if userExist.ID == 0 || userExistResult.Error != nil {
		return false, nil, userExistResult.Error
	}

	// ===========================================================================================
	// ===========================================================================================
	// =========================================== QUERY =========================================
	// TRAEMOS LA SECCION CON EL NOMBRE QUE PEDIMOS.
	sectionGetResult := postgres.DB.Where("name	=	?", sectionNameQuery).First(&sectionExist)
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

// VERDE
func SectionPostRepository(section *dto.SectionDto) (bool, *dto.SectionDto, error) {
	// ===========================================================================================
	// ===========================================================================================
	// =========================================== VALIDACIONES ==================================
	// OBJETO DE USUARIO.
	var userExist domain.User

	// OBJETO DE SECCION.
	var parentSectionExist domain.Section

	// VALIDAMOS QUE EL ID DEL USUARIO EXISTA, SI NO, DEVUELVE ERROR.
	userExistResult := postgres.DB.First(&userExist, "id	=	?", section.UserID)
	if userExist.ID == 0 || userExistResult.Error != nil {
		return false, nil, userExistResult.Error
	}

	// VALIDAMOS QUE EL *SectionParentId* EXISTA, SI NO, QUE EL  *SectionParentId* SEA NULL.
	if section.SectionParentId != nil {
		parentSectionExistResult := postgres.DB.First(&parentSectionExist, "id	=	?", section.SectionParentId)
		if parentSectionExist.ID == 0 || parentSectionExistResult.Error != nil {
			return false, nil, parentSectionExistResult.Error
		}
	}

	// ===========================================================================================
	// ===========================================================================================
	// ================================== QUERY -  CREACION DE LA SECCION ========================
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

// VERDE
func SectionUpdateRepository(section *dto.SectionDto) (bool, *dto.SectionDto, error) {
	// OBJETO DE USUARIO.
	var userExist domain.User

	// OBJETO DE SECCION.
	var sectionExist domain.Section

	// ===========================================================================================
	// ===========================================================================================
	// =========================================== VALIDACIONES ==================================
	// VALIDAMOS QUE EL USUARIO EXISTA.
	userExistResutlt := postgres.DB.First(&userExist, "id	=	?", section.UserID)
	if userExist.ID == 0 || userExistResutlt.Error != nil {
		return false, nil, userExistResutlt.Error
	}

	// VALIDAMOS QUE LA SECCION EXISTA.
	sectionExistResult := postgres.DB.First(&sectionExist, "id	=	?", section.ID)
	if sectionExistResult.Error != nil {
		return false, nil, sectionExistResult.Error
	}

	// ===========================================================================================
	// ===========================================================================================
	// =========================================== QUERY =========================================
	// MODIFICACION DE LA SECCION.
	sectionUpdateResult := postgres.DB.Model(&sectionExist).Where("id	=	?", section.ID).Updates(map[string]interface{}{
		"Name":            strings.ToUpper(section.Name),
		"Description":     section.Description,
		"SectionParentId": section.SectionParentId,
	})

	if sectionUpdateResult.Error != nil {
		return false, nil, sectionUpdateResult.Error
	}

	// CONSTRUCCION DEL DTO DE RETORNO.
	dtoReturn := dto.SectionDto{
		ID:              sectionExist.ID,
		CreatedAt:       sectionExist.CreatedAt,
		Name:            sectionExist.Name,
		Description:     sectionExist.Description,
		UserID:          sectionExist.UserID,
		SectionParentId: sectionExist.SectionParentId,
	}

	// SI TODO SALE BIEN.
	return true, &dtoReturn, nil
}

// VERDE
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
	sectionExistResult := postgres.DB.First(&sectionExist, section.ID)
	if sectionExistResult.Error != nil {
		return false, sectionExistResult.Error
	}

	// ===========================================================================================
	// ===========================================================================================
	// =========================================== QUERY =========================================
	//ELIMINACION DE LA SECCION EN LA BASE DE DATO.
	sectionDeleteResult := postgres.DB.Unscoped().Delete(&sectionExist, section.ID)
	if sectionDeleteResult.Error != nil {
		return false, sectionDeleteResult.Error
	}

	// SI TODO SALE BIEN...
	return true, nil
}
