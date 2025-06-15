package sectionsusecase

import (
	"errors"
	"strconv"
	"strings"

	"github.com/maxexee/rugaPasswordManager/internal/dto"
	sectionsrepository "github.com/maxexee/rugaPasswordManager/internal/repository/sections_repository"
)

// VERDE...
func SectionGetAllUseCase(userIdStr string, sectionIdStr string) (bool, *dto.SectionPasswordGetSliceDTO, error) {
	// OBJETO DEL TIPO *dto.SectionDto*
	var dtoSend dto.SectionDto

	// CONVERSION DE STRING A INT PARA EL ID DEL USUARIO.
	userId, userIdError := strconv.Atoi(userIdStr)

	// CONVERSION DE STRING A INT PARA EL ID DE LA SECCION.
	sectionId, sectionIdError := strconv.Atoi(sectionIdStr)

	// ===========================================================================================
	// ===========================================================================================
	// =========================================== VALIDACIONES ==================================
	if userIdError != nil {
		return false, nil, userIdError
	}

	// ===========================================================================================
	// ===========================================================================================
	// =========================================== QUERY =========================================
	// SI EL TIPO DE DATO ES CORRECTO, SE ASIGNA EL VALOR AL "UserId" DEL DTO.
	dtoSend.UserID = uint(userId)

	// ASIGNAMOS  "SectionId", SI EL TIPO DE DATOS ES "nil" O "uint".
	if sectionIdStr == "" || sectionIdStr == "null" {
		dtoSend.SectionParentId = nil
	} else {
		if sectionIdError != nil {
			return false, nil, sectionIdError
		}
		u := uint(sectionId)
		dtoSend.SectionParentId = &u
	}

	// LLAMDA A LA BASE DE DATOS PARA QUE TRAIGA EL SLICE CON LAS SECCIONES.
	ok, sectionsPasswordsReturn, sectionReturnResult := sectionsrepository.SectionGetAllRepository(&dtoSend)
	if !ok {
		return false, nil, sectionReturnResult
	}

	// SI TODO SALE BIEN...
	return true, sectionsPasswordsReturn, nil
}

// VERDE...
func SectionGetByNameUseCase(userIdStr string, sectionName string) (bool, *dto.SectionDto, error) {
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
	// CREACION DEL DTO DE ENVIO CON EL ID DEL USUARIO Y DEL NOMBRE DE LA CONTRASEÑA.
	dtoSend := dto.SectionDto{
		UserID: uint(userId),
		Name:   strings.ReplaceAll(sectionName, " ", "_"),
	}
	// LLAMDA AL REPOSITORIO.
	ok, sectionReturn, sectionReturnError := sectionsrepository.SectionGetByNameRepository(&dtoSend)
	if !ok {
		return false, nil, sectionReturnError
	}

	// SI TODO SALE BIEN...
	return true, sectionReturn, nil
}

// VERDE...
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
	dtoSend := dto.SectionDto{
		UserID:          uint(userId),
		Name:            strings.ReplaceAll(section.Name, " ", "-"),
		Description:     section.Description,
		SectionParentId: section.SectionParentId,
	}

	// LLAMADA AL REPOSITORY.
	ok, sectionCreated, sectionCreatedError := sectionsrepository.SectionPostRepository(&dtoSend)
	if !ok {
		return false, nil, sectionCreatedError
	}

	// SI TODO SALE BIEN...
	return true, sectionCreated, nil
}

// VERDE...
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
	// CREACION DEL DTO DE ENVIO.
	dtoSend := dto.SectionDto{
		UserID:          uint(userId),
		ID:              uint(sectionId),
		Name:            strings.ReplaceAll(section.Name, " ", "-"),
		Description:     section.Description,
		SectionParentId: section.SectionParentId,
	}

	// LLAMADA AL REPOSITORIO PARA LA ACTUALIZACION DE LA SECCION EN LA BASE DE DATOS.
	ok, sectionReturn, sectionReturnError := sectionsrepository.SectionUpdateRepository(&dtoSend)
	if !ok {
		return false, sectionReturn, sectionReturnError
	}

	// SI TODO SALE BIEN...
	return true, sectionReturn, nil
}

// VERDE...
func SectionDeleteUseCase(userIdStr string, sectionIdStr string) (bool, error) {
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
	// CONTRUCCION DEL DTO DE ENVIO.
	// (-- VALIDAR SI ES NECESARIO O ES MEJOR SOLO MANDAR LOS VALORES INT DE ARRIBA -- )
	dtoSent := dto.SectionDto{
		UserID: uint(userId),
		ID:     uint(sectionId),
	}

	// LLAMADA AL REPOSITORY PARA ELIMINACION DE LA SECCION.
	ok, sectionReturnError := sectionsrepository.SectionDeleteRepository(&dtoSent)
	if !ok {
		return false, sectionReturnError
	}

	// SI TODO SALE BIEN...
	return true, nil
}
