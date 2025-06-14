package sectionsusecase

import (
	"errors"
	"strconv"
	"strings"

	"github.com/maxexee/rugaPasswordManager/internal/dto"
	sectionsrepository "github.com/maxexee/rugaPasswordManager/internal/repository/sections_repository"
)

// VERDE..
func SectionGetAllUseCase(userIdStr string, sectionIdStr string) (bool, *dto.SectionGetSliceDTO, error) {
	// OBJETO DEL TIPO *dto.SectionDto*
	var section dto.SectionDto

	// CONVERSION DEL *userIdStr* DE TIPO STRING A TIPO INT.
	userId, userIdError := strconv.Atoi(userIdStr)

	// CONVERSION DEL *sectionIdStr* DE TIPO STRING A INT.
	sectionId, sectionIdError := strconv.Atoi(sectionIdStr)

	// ===========================================================================================
	// ===========================================================================================
	// =========================================== VALIDACIONES ==================================
	// VALIDAMOS QUE EL TIPO DE DATO DEL ID DEL USUARIO SEA CORRECTO (TIPO UINT).
	if userIdError != nil {
		return false, nil, userIdError
	}

	// ===========================================================================================
	// ===========================================================================================
	// =========================================== QUERY =========================================
	// SI EL TIPO DE DATO ES CORRECTO, SE ASIGNA EL VALOR AL "UserId" DEL DTO.
	section.UserID = uint(userId)

	// ASIGNAMOS  "SectionId", SI EL TIPO DE DATOS ES "nil" O "uint".
	if sectionIdStr == "" || sectionIdStr == "null" {
		section.SectionParentId = nil
	} else {
		if sectionIdError != nil {
			return false, nil, sectionIdError
		}
		u := uint(sectionId)
		section.SectionParentId = &u
	}

	// LLAMDA A LA BASE DE DATOS PARA QUE TRAIGA EL SLICE CON LAS SECCIONES.
	ok, sectionsReturn, sectionReturnResult := sectionsrepository.SectionGetAllRepository(&section)
	if !ok {
		return false, nil, sectionReturnResult
	}

	// SI TODO SALE BIEN...
	return true, sectionsReturn, nil
}

// VERDE..
func SectionGetByNameUseCase(userIdStr string, sectionName string) (bool, *dto.SectionDto, error) {
	// OBJETO DEL TIPO *dto.SectionDto*
	var section dto.SectionDto

	// CONVERSION DEL *userIdStr* DE TIPO STRING A TIPO INT.
	userId, userIdError := strconv.Atoi(userIdStr)

	// ===========================================================================================
	// ===========================================================================================
	// =========================================== VALIDACIONES ==================================
	if userIdError != nil {
		return false, nil, userIdError
	}

	// ===========================================================================================
	// ===========================================================================================
	// ====================================== QUERY ==============================================
	// ASIGNACION DEL USER ID AL DTO PARA ENVIAR AL REPOSITORY.
	section.UserID = uint(userId)

	// LLAMDA AL REPOSITORIO.
	ok, sectionReturn, sectionReturnError := sectionsrepository.SectionGetByNameRepository(&section, strings.ToUpper(sectionName))
	if !ok {
		return false, nil, sectionReturnError
	}

	// SI TODO SALE BIEN...
	return true, sectionReturn, nil
}

// VERDE..
func SectionPostUseCase(userIdStr string, section *dto.SectionDto) (bool, *dto.SectionDto, error) {
	// CONVERSION DE ID DEL USUARIO, DE TIPO STRING A INT.
	userId, userIdError := strconv.Atoi(userIdStr)

	// ===========================================================================================
	// ===========================================================================================
	// =========================================== VALIDACIONES ==================================
	if userIdError != nil {
		return false, nil, userIdError
	}

	// ===========================================================================================
	// ===========================================================================================
	// ====================================== QUERY ==============================================
	// ASIGNACION DEL ID DEL USUARIO EN EL DTO.
	section.UserID = uint(userId)

	// LLAMADA AL REPOSITORY.
	ok, sectionCreated, sectionCreatedError := sectionsrepository.SectionPostRepository(section)
	if !ok {
		return false, nil, sectionCreatedError
	}

	// SI TODO SALE BIEN...
	return true, sectionCreated, nil
}

// VERDE..
func SectionUpdateUseCase(userIdStr string, sectionIdStr string, section *dto.SectionDto) (bool, *dto.SectionDto, error) {
	// CONVERSION DE ID DEL USUARIO, DE TIPO STRING A INT.
	userId, userIdError := strconv.Atoi(userIdStr)

	// CONVERSION DE ID DE LA SECCION, DE TIPO STRING A INT.
	sectionId, sectionIdError := strconv.Atoi(sectionIdStr)

	// ===========================================================================================
	// ===========================================================================================
	// =========================================== VALIDACIONES ==================================
	if userIdError != nil {
		return false, nil, userIdError
	}

	if sectionIdError != nil {
		return false, nil, sectionIdError
	}

	// VALIDAMOS QUE NO SE PUEDA ASIGNAR COMO SectionParentId A LA MISMA SECCION A EDITAR.
	if section.SectionParentId != nil {
		if *section.SectionParentId == uint(sectionId) {
			return false, nil, errors.New("este es mi mensaje de error")
		}
	}

	// ===========================================================================================
	// ===========================================================================================
	// ====================================== QUERY ==============================================
	// ASIGNACION DE LOS VALORES DE USUARIO ID Y SECCION ID EN EL DTO
	section.UserID = uint(userId)
	section.ID = uint(sectionId)

	// LLAMADA AL REPOSITORIO PARA LA ACTUALIZACION DE LA SECCION EN LA BASE DE DATOS.
	ok, sectionReturn, sectionReturnError := sectionsrepository.SectionUpdateRepository(section)
	if !ok {
		return false, sectionReturn, sectionReturnError
	}

	// SI TODO SALE BIEN...
	return true, sectionReturn, nil
}

// VERDE
func SectionDeleteUseCase(userIdStr string, sectionIdStr string) (bool, error) {
	// OBJETO DE TIPO *dto.SectionDto*.
	var sectionExist dto.SectionDto

	// CONVERSION DE ID DEL USUARIO, DE TIPO STRING A INT.
	userId, userIdError := strconv.Atoi(userIdStr)

	// // CONVERSION DE ID DE LA SECCION, DE TIPO STRING A INT.
	sectionId, sectionIdError := strconv.Atoi(sectionIdStr)

	// ===========================================================================================
	// ===========================================================================================
	// =========================================== VALIDACIONES ==================================
	if userIdError != nil {
		return false, userIdError
	}

	if sectionIdError != nil {
		return false, sectionIdError
	}

	// ===========================================================================================
	// ===========================================================================================
	// ====================================== QUERY ==============================================
	// ASIGNACION DE VALORES DEL ID DEL USUARIO Y LA SECCION AL DTO.
	sectionExist.UserID = uint(userId)
	sectionExist.ID = uint(sectionId)

	// LLAMADA AL REPOSITORY PARA ELIMINACION DE LA SECCION.
	ok, sectionReturnError := sectionsrepository.SectionDeleteRepository(&sectionExist)
	if !ok {
		return false, sectionReturnError
	}

	// SI TODO SALE BIEN...
	return true, nil
}
