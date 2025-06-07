package sectionsusecase

import (
	"strconv"

	"github.com/maxexee/rugaPasswordManager/internal/domain"
	"github.com/maxexee/rugaPasswordManager/internal/dto"
	sectionsrepository "github.com/maxexee/rugaPasswordManager/internal/repository/sections_repository"
)

func SectionGetAllUseCase(userIdStr string, sectionIdStr string) (bool, []domain.Section, error) {
	// ===========================================================================================
	// ===========================================================================================
	// =========================================== VALIDACIONES ==================================
	var section dto.SectionGetAllDTO
	userId, userIdError := strconv.Atoi(userIdStr)
	sectionId, sectionIdError := strconv.Atoi(sectionIdStr)

	// VALIDAMOS QUE EL TIPO DE DATO DEL ID DEL USUARIO SEA CORRECTO (TIPO UINT).
	if userIdError != nil {
		return false, nil, userIdError
	}
	// SI EL TIPO DE DATO ES CORRECTO, SE ASIGNA EL VALOR AL "UserId" DEL DTO.
	section.UserId = uint(userId)

	// ASIGNAMOS  "SectionId", SI EL TIPO DE DATOS ES "nil" O "uint".
	if sectionIdStr == "" || sectionIdStr == "null" {
		section.SectionId = nil
	} else {
		if sectionIdError != nil {
			return false, nil, sectionIdError
		}
		u := uint(sectionId)
		section.SectionId = &u
	}

	// ===========================================================================================
	// ===========================================================================================
	// =========================================== QUERY =========================================
	ok, sectionsReturn, sectionReturnResult := sectionsrepository.SectionGetAllRepository(&section)
	if !ok {
		return false, nil, sectionReturnResult
	}

	return true, sectionsReturn, nil
}
