package passwords

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	postgres "github.com/maxexee/rugaPasswordManager/infrastructure/db"
	"github.com/maxexee/rugaPasswordManager/internal/domain"
	validations "github.com/maxexee/rugaPasswordManager/internal/handler/validations"
)

// VERDE
func PasswordGetByName(c *gin.Context) {
	// ID DEL USUARIO DESDE EL URL Y SU CONVERSION A UINT.
	idStr := c.Param("id")
	userId, errUserId := strconv.Atoi(idStr)

	// OBTENCION DEL NOMBRE DE LA CONTRASEÑA.
	passwordName := strings.ToUpper(c.Query("namePass"))

	// ===========================================================================================
	// ===========================================================================================
	// =========================================== VALIDACIONES ==================================
	// VALIDAMOS QUE EL TIPO DE DATO DEL ID DEL USUARIO SEA CORRECTO (TIPO UINT).
	if !validations.ErrorValidations(c, errUserId, "Invalid type of User ID...") {
		c.Abort()
		return
	}

	// VALIDAMOS QUE EL ID DEL USUARIO EXISTA, SI NO, DEVUELVE ERROR.
	var userExist domain.User
	userExistResutlt := postgres.DB.First(&userExist, "id	=	?", userId)
	if !validations.DataBaseValidations(c, userExistResutlt, userExist.ID, "User ID Not Found...") {
		c.Abort()
		return
	}

	// ===========================================================================================
	// ===========================================================================================
	// =========================================== QUERY =========================================
	// QUERY PARA ENCONTRAR CONTRASEÑAS POR NOMBRE
	var passwordExist domain.Password
	passwordExistResutlt := postgres.DB.Where("name	=	?", passwordName).First(&passwordExist)
	if !validations.DataBaseValidations(c, passwordExistResutlt, passwordExist.ID, "Password not found...") {
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"STATUS":   "Password is here...",
		"PASSWORD": passwordExist,
	})
}

// VERDE
func PasswordPost(c *gin.Context) {
	// BODY.
	var body struct {
		Name                    string  `json:"name" validate:"required,min=2,max=20,matchesName=^[A-Za-z0-9 ]+$"`
		Description             *string `json:"description" validate:"omitempty,max=50"`
		Password                string  `json:"password" validate:"required,min=12"`
		SectionParentIdPassword uint    `json:"sectionparentidpassword" validate:"required,number"`
	}

	idStr_user := c.Param("id") // -- OBTENEMOS EL ID DESDE EL URL. ---
	userId, errUserId := strconv.Atoi(idStr_user)

	// ===========================================================================================
	// ===========================================================================================
	// =========================================== VALIDACIONES ==================================
	if !validations.BodyValidation(c, &body) {
		c.Abort()
		return
	}

	// VALIDAMOS QUE EL TIPO DE DATO DEL ID DEL USUARIO SEA EL CORRECTO (TIPO UINT).
	if !validations.ErrorValidations(c, errUserId, "Invalid type of User ID...") {
		c.Abort()
		return
	}

	// VALIDAMOS QUE EL ID DEL USUARIO EXISTA, SI NO, DEVUELVE ERROR.
	var userExist domain.User
	userValidationResult := postgres.DB.First(&userExist, "id = ?", userId)
	if !validations.DataBaseValidations(c, userValidationResult, userExist.ID, "User ID not found...") {
		c.Abort()
		return
	}

	// VALIDAMOS QUE LA SECCION PADRE EXISTA (NO SE ACEPTAN CONSTRASEÑAS EN LA SECCION RAIZ (O NULL)).
	var sectionExist domain.Section
	sectionValidationResult := postgres.DB.Where("id	=	?", body.SectionParentIdPassword).Select("id").First(&sectionExist)
	if sectionExist.ID == 0 || sectionValidationResult.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Section ID not found..."})
		c.Abort()
		return
	}

	// ===========================================================================================
	// ===========================================================================================
	// =========================================== QUERY =========================================
	//CREACION DE LA NUEVA CONTRASEÑA.
	passwordNew := domain.Password{
		Name:                    strings.ToUpper(body.Name),
		Description:             body.Description,
		Password:                body.Password,
		SectionParentIdPassword: body.SectionParentIdPassword,
	}

	// GUARDADO DE LA CONSTRASEÑA NUEVA.
	passwordCreation := postgres.DB.Create(&passwordNew)
	if !validations.ErrorValidations(c, passwordCreation.Error, "Failed to save password on the database...") {
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, gin.H{"great-password": "Password created successfully..."})
}

// VERDE -- REFERENCIA CORRECTA POR AHORA --
func PasswordUpdate(c *gin.Context) {
	// OBTENCION DE LOS CAMPOS A ACTUALIZAR DESDE EL BODY.
	var body struct {
		Name                    string  `json:"name" validate:"required,min=2,max=20,matchesName=^[A-Za-z0-9 ]+$"`
		Description             *string `json:"description" validate:"omitempty,max=50"`
		Password                string  `json:"password" validate:"required,min=12"`
		SectionParentIdPassword uint    `json:"sectionparentidpassword" validate:"required,number"`
	}

	// OBTENCION DEL ID DEL USUARIO DESDE EL URL.
	userIdStr := c.Param("id")
	userId, errUserId := strconv.Atoi(userIdStr)

	// OBTENCION DEL ID DE LA CONSTRASEÑA DESDE EL URL.
	passwordIdStr := c.Param("idU")
	passwordId, errPassId := strconv.Atoi(passwordIdStr)
	// ===========================================================================================
	// ===========================================================================================
	// =========================================== VALIDACIONES ==================================
	// VALIDACION DE LOS DATOS DE BODY.
	if !validations.BodyValidation(c, &body) {
		c.Abort()
		return
	}

	// VALIDAMOS QUE EL TIPO DE DATO DEL ID DEL USUARIO SEA CORRECTO (TIPO UINT).
	if !validations.ErrorValidations(c, errUserId, "Invalid type of User ID...") {
		c.Abort()
		return
	}

	// VALIDAMOS QUE EL ID DE LA CONSTRASEÑA TENGA EL FORMATO CORRECTO.
	if !validations.ErrorValidations(c, errPassId, "Invalid type of Password ID...") {
		c.Abort()
		return
	}

	// VALIDAMOS QUE EL ID DEL USUARIO EXISTA, SI NO, DEVUELVE ERROR.
	var userExist domain.User
	userExistResult := postgres.DB.First(&userExist, "id	=	?", userId)
	if !validations.DataBaseValidations(c, userExistResult, userExist.ID, "User ID Not Found...") {
		c.Abort()
		return
	}

	// ===========================================================================================
	// ===========================================================================================
	// =========================================== QUERY =========================================
	//
	var passwordExist domain.Password
	passwordUpdateResult := postgres.DB.Model(&passwordExist).Where("id	=	?", passwordId).Updates(map[string]interface{}{
		"Name":                    strings.ToUpper(body.Name),
		"Description":             &body.Description,
		"Password":                body.Password,
		"SectionParentIdPassword": body.SectionParentIdPassword,
	})
	if !validations.ErrorValidations(c, passwordUpdateResult.Error, "Not able to update the current password") {
		c.Abort()
		return
	}

	if passwordUpdateResult.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"Error": "Password not found..."})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, gin.H{"STATUS": "Password update has been done correct..."})
}

// VERDE
func PasswordDelete(c *gin.Context) {
	// OBTENCION DEL ID DEL USUARIO DESDE EL URL.
	userIdStr := c.Param("id")
	userId, errUserId := strconv.Atoi(userIdStr)

	// OBTENCION DEL ID DE LA CONTRASEÑA DESDE EL URL.
	passwordIdStr := c.Param("idD")
	passwordId, errPassId := strconv.Atoi(passwordIdStr)
	// ===========================================================================================
	// ===========================================================================================
	// =========================================== VALIDACIONES ==================================
	// VALIDAMOS QUE EL TIPO DE DATO DEL ID DEL USUARIO SEA CORRECTO (TIPO UINT).
	if !validations.ErrorValidations(c, errUserId, "Invalid type of User ID...") {
		c.Abort()
		return
	}

	// VALIDAMOS QUE EL ID DE LA CONTRASEÑA TENGA EL FORMATO CORRECTO (TIPO UINT).
	if !validations.ErrorValidations(c, errPassId, "Invalid type of Password ID...") {
		c.Abort()
		return
	}

	// VALIDAMOS QUE EL ID DEL USUARIO EXISTA, SI NO, DEVUELVE ERROR.
	var userExist domain.User
	userExistResult := postgres.DB.First(&userExist, "id	=	?", userId)
	if !validations.DataBaseValidations(c, userExistResult, userExist.ID, "User ID Not Found...") {
		c.Abort()
		return
	}

	// ===========================================================================================
	// ===========================================================================================
	// =========================================== QUERY =========================================
	// ELIMINACION DE LA CONTRASEÑA.
	var passwordExist domain.Password
	passwordDeleteResult := postgres.DB.Unscoped().Delete(&passwordExist, passwordId)
	// SI HAY UN ERROR...
	if !validations.ErrorValidations(c, passwordDeleteResult.Error, "Error trying to delete the password...") {
		c.Abort()
		return
	}

	// SI NO HAY ERROR, PERO NO NO HAY REGISTROS ELIMINADOS, ENTONCES NO SE ENCONTRO LA CONTRASEÑA.
	if passwordDeleteResult.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"STATUS": "Password not found...",
		})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, gin.H{"STATUS": "Password deleted correctly..."})
}
