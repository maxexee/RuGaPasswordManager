package sections

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	postgres "github.com/maxexee/rugaPasswordManager/infrastructure/db"
	"github.com/maxexee/rugaPasswordManager/internal/domain"
	validations "github.com/maxexee/rugaPasswordManager/internal/handler/validations"
	sectionsusecase "github.com/maxexee/rugaPasswordManager/internal/use_case/sections_use_case"
)

func SectionGet(c *gin.Context) {
	// ID DEL USUARIO DESDE EL URL Y SU CONVERSION A UINT.
	userIdStr := c.Param("id")

	// ID DE LA SECCION DESDE EL URL Y SU CONVERSION A UINT.
	sectionIdStr := strings.ToLower(c.Query("section_parent_id"))

	// ===========================================================================================
	// ===========================================================================================
	// =========================================== QUERY =========================================
	// LLAMADA AL CASO DE USO PARA OBTENER TODAS LAS SECCIONES, YA SEA LA RAIZ (NULL) O LAS SECCIONES HIJAS DE UNA SECCION
	// PADRE.
	ok, sectionsReturn, sectionReturnError := sectionsusecase.SectionGetAllUseCase(userIdStr, sectionIdStr)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{
			"STATUS": "ERROR ...",
			"ERROR":  sectionReturnError.Error(),
		})
		c.Abort()
		return
	}

	//
	c.JSON(http.StatusOK, gin.H{
		"STATUS":   "All sections...",
		"SECTIONS": sectionsReturn,
	})
}

func SectionGetByName(c *gin.Context) {
	//	ID DEL USUARIO DESDE EL URL Y SU CONVERSION A UINT.
	idStr := c.Param("id")
	userId, errUser := strconv.Atoi(idStr)
	// ===========================================================================================
	// ===========================================================================================
	// =========================================== VALIDACIONES ==================================
	// VALIDAMOS QUE EL TIPO DE DATO DEL ID DEL USUARIO SEA EL CORRECTO (TIPO UINT).
	if !validations.ErrorValidations(c, errUser, "Invalid Type of User ID...") {
		c.Abort()
		return
	}

	// VALIDAMOS QUE EL ID DEL USUARIO EXISTA, SI NO, DEVUELVE ERROR.
	var userExist domain.User
	userCreationResult := postgres.DB.First(&userExist, "id = ?", userId)
	if !validations.DataBaseValidations(c, userCreationResult, userExist.ID, "User ID not found...") {
		c.Abort()
		return
	}

	// ===========================================================================================
	// ===========================================================================================
	// =========================================== QUERY =========================================
	// QUERY PARA ENCONTRAR SECCIONES POR NOMBRE.
	sectionName := strings.ToUpper(c.Query("nameSec"))
	var sectionExist domain.Section
	sectionCreationResult := postgres.DB.Where("name	=	?", sectionName).First(&sectionExist)
	if !validations.DataBaseValidations(c, sectionCreationResult, sectionExist.ID, "Section not found...") {
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"STATUS":  "Section is here",
		"SECTION": sectionExist,
	})
}

func SectionPost(c *gin.Context) {
	// ===========================================================================================
	// =========================================== BODY ==========================================
	// DEFINIMOS QUE EL BODY QUE RECIBIREMOS DE LA PETICIÓN.
	var body struct {
		Name            string `json:"name" validate:"required,min=2,max=20,matchesName=^[A-Za-z0-9 ]+$"`
		Description     string `json:"description"`
		SectionParentId *uint
	}

	// ===========================================================================================
	// ===========================================================================================
	// =========================================== VALIDACIONES ==================================
	// === VALIDAMOS QUE NO HAYA ERROR EN EL BODY Y TODO SEA LEGIBLE. ===
	if !validations.BodyValidation(c, &body) {
		c.Abort()
		return
	}

	// === VALIDAMOS QUE EL TIPO DE DATO DEL ID DEL USUARIO SEA EL CORRECTO (TIPO UINT). ===
	idStr := c.Param("id") // -- OBTENEMOS EL ID DESDE EL URL. ---
	userId, errUserId := strconv.Atoi(idStr)
	if !validations.ErrorValidations(c, errUserId, "Invalid type of ID...") {
		c.Abort()
		return
	}

	// === VALIDAMOS QUE EL ID DEL USUARIO EXISTA, SI NO, DEVUELVE ERROR. ===
	var userExist domain.User
	userCreationResult := postgres.DB.First(&userExist, "id = ?", userId)
	if !validations.DataBaseValidations(c, userCreationResult, userExist.ID, "User ID not found...") {
		c.Abort()
		return
	}

	// === VALIDAMOS QUE EL *SectionParentId* EXISTA, SI NO, QUE EL  *SectionParentId* SEA NULL ===
	var parentSectionExist domain.Section
	if body.SectionParentId != nil {
		parentResult := postgres.DB.First(&parentSectionExist, "id = ?", body.SectionParentId)
		if !validations.ErrorValidations(c, parentResult.Error, "Parent section not found...") {
			c.Abort()
			return
		}
	}

	// ===========================================================================================
	// ===========================================================================================
	// =========================================== CREACION DE LA SECCION ========================
	// CREACION DE LA SECCION MEDIANTE EL MODELO *Section*.
	section := domain.Section{Name: strings.ToUpper(body.Name), Description: &body.Description, UserID: uint(userId), SectionParentId: body.SectionParentId}

	// GUARDADO DE LA NUEVA SECCION EN LA BASE DE DATOS.
	result := postgres.DB.Create(&section)
	if !validations.ErrorValidations(c, result.Error, "Failed to save section on the Database...") {
		c.Abort()
		return
	}

	// RETORNAMOS QUE LA SECCION SE HA CREADO CORRECTAMENTE.
	c.JSON(http.StatusOK, gin.H{"great-section": "Section created successfully..."})
}

func SectionUpdate(c *gin.Context) {
	// BODY A RECIBIR DEL FRONT-END.
	var body struct {
		Name            string `json:"name"        validate:"required,min=2,max=20,matchesName=^[A-Za-z0-9 ]+$"`
		Description     string `json:"description" validate:"omitempty"`
		SectionParentId *uint
	}

	// ===========================================================================================
	// ===========================================================================================
	// VALIDAMOS QUE EL TIPO DE DATO DEL ID DEL USUARIO SEA EL CORRECTO (TIPO UINT).
	// VALIDACION DEL BODY.
	if !validations.BodyValidation(c, &body) {
		c.Abort()
		return
	}

	idStr := c.Param("id") // -- OBTENEMOS EL ID DESDE EL URL. ---
	userId, errUserId := strconv.Atoi(idStr)
	if !validations.ErrorValidations(c, errUserId, "Invalid type of ID...") {
		c.Abort()
		return
	}

	// VALIDAMOS QUE EL ID DEL USUARIO EXISTA, SI NO, DEVUELVE ERROR.
	var userExist domain.User
	userCreationResult := postgres.DB.First(&userExist, "id = ?", userId)
	if !validations.DataBaseValidations(c, userCreationResult, userExist.ID, "User ID not found...") {
		c.Abort()
		return
	}

	// ===========================================================================================
	// ===========================================================================================
	// =========================================== QUERY =========================================
	// OBTENCION DE EL ID DE LA SECCION A MODIFICAR.
	var section domain.Section
	idStrUpdate := c.Param("idU")
	sectionId, errSecId := strconv.Atoi(idStrUpdate)
	if !validations.ErrorValidations(c, errSecId, "Invalid type of Section ID...") {
		c.Abort()
		return
	}

	// VERIFICAMOS QUE NO SE PUEDA ASIGNAR COMO SectionParentId A LA MISMA SECCION A EDITAR.
	if body.SectionParentId != nil {
		if sectionId == int(*body.SectionParentId) {
			c.JSON(http.StatusBadRequest, gin.H{"ERROR": "Invalid Section ID..."})
			c.Abort()
			return
		}
	}

	// VALIDAMOS QUE EXISTA LA SECCION A EDITAR.
	sectionGetById := postgres.DB.First(&section, idStrUpdate)
	if !validations.ErrorValidations(c, sectionGetById.Error, "Section not found...") {
		c.Abort()
		return
	}

	// MODIFICACION DE LA SECCION.
	sectionUpdateResult := postgres.DB.Model(&section).Where("id	=	?", sectionId).Updates(map[string]interface{}{
		"Name":            strings.ToUpper(body.Name),
		"Description":     body.Description,
		"SectionParentId": body.SectionParentId,
	})
	if !validations.ErrorValidations(c, sectionUpdateResult.Error, "Not able to update the current section...") {
		c.Abort()
		return
	}

	// RETORNO AL DEL ESTADO OK AL FRON-END.
	c.JSON(http.StatusOK, gin.H{
		"STATUS":  "Section update has been done correct...",
		"SECTION": section,
	})
}

func SectionDelete(c *gin.Context) {
	// ===========================================================================================
	// ===========================================================================================
	// ===========================================================================================
	// VALIDAMOS QUE EL TIPO DE DATO DEL ID DEL USUARIO SEA EL CORRECTO (TIPO UINT).
	idStr := c.Param("id") // -- OBTENEMOS EL ID DESDE EL URL. ---
	userId, errUserId := strconv.Atoi(idStr)
	if !validations.ErrorValidations(c, errUserId, "Invalid type of User ID...") {
		c.Abort()
		return
	}

	// VALIDAMOS QUE EL ID DEL USUARIO EXISTA, SI NO, DEVUELVE ERROR.
	var userExist domain.User
	userCreationResult := postgres.DB.First(&userExist, "id = ?", userId)
	if !validations.DataBaseValidations(c, userCreationResult, userExist.ID, "User ID not found...") {
		c.Abort()
		return
	}

	// ===========================================================================================
	// ===========================================================================================
	// =========================================== QUERY =========================================
	// VALIDAMOS QUE EL ID DE LA SECCION TENGA EL FORMATO CORRECTO.
	var section domain.Section
	sectionIdStr := c.Param("idD")
	sectionId, errSectionId := strconv.Atoi(sectionIdStr)
	if !validations.ErrorValidations(c, errSectionId, "Invalid type of Section ID...") {
		c.Abort()
		return
	}

	// VALIDAMOS QUE EXISTA LA SECCION.
	sectionExist := postgres.DB.First(&section, sectionId)
	if !validations.ErrorValidations(c, sectionExist.Error, "Section not found...") {
		c.Abort()
		return
	}

	//
	sectionDeleteResutl := postgres.DB.Unscoped().Delete(&section, sectionId)
	if !validations.ErrorValidations(c, sectionDeleteResutl.Error, "Error trying to delete the Section...") {
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, gin.H{"STATUS": "ELIMINADO CORRECTAMENTE..."})
}
