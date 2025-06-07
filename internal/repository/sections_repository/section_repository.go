package sectionsrepository

import (
	postgres "github.com/maxexee/rugaPasswordManager/infrastructure/db"
	"github.com/maxexee/rugaPasswordManager/internal/domain"
	"github.com/maxexee/rugaPasswordManager/internal/dto"
)

func SectionGetAllRepository(section *dto.SectionGetAllDTO) (bool, []domain.Section, error) {
	// ===========================================================================================
	// ===========================================================================================
	// =========================================== VALIDACIONES ==================================
	var userExist domain.User
	userExistResutlt := postgres.DB.First(&userExist, "id	=	?", section.UserId)
	if userExist.ID == 0 || userExistResutlt.Error != nil {
		return false, nil, userExistResutlt.Error
	}

	// ===========================================================================================
	// ===========================================================================================
	// =========================================== QUERY =========================================
	var sections []domain.Section

	if section.SectionId == nil {
		sectionsGetAllResult := postgres.DB.Where("user_id = ?	AND	section_parent_id	IS NULL", section.UserId).Find(&sections)
		if sectionsGetAllResult.Error != nil {
			return false, nil, sectionsGetAllResult.Error
		}
	} else {
		// VALIDAR SI EXISTE EL PADRE.
		sectionExistResult := postgres.DB.First(&sections, "id	=	?", section.SectionId)
		if sectionExistResult.Error != nil {
			return false, nil, sectionExistResult.Error
		}

		// OBTENCION DE LAS SECCIONES MEDIANTE UN PADRE.
		sectionsGetAllResult := postgres.DB.Where("user_id	=	?	AND	section_parent_id	=	?", section.UserId, section.SectionId).Find(&sections)
		if sectionsGetAllResult.Error != nil {
			return false, nil, sectionExistResult.Error
		}
	}

	return true, sections, nil
}
