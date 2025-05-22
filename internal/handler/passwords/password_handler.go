package passwords

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	postgres "github.com/maxexee/rugaPasswordManager/infrastructure/db"
	"github.com/maxexee/rugaPasswordManager/internal/domain"
)

// VERDE
func PasswordGetByName(c *gin.Context) {
	// ID DEL USUARIO DESDE EL URL Y SU CONVERSION A UINT.
	idStr := c.Param("id")
	userId, errUser := strconv.Atoi(idStr)

	// ===========================================================================================
	// ===========================================================================================
	// =========================================== VALIDACIONES ==================================
	// VALIDAMOS QUE EL TIPO DE DATO DEL ID DEL USUARIO SEA CORRECTO (TIPO UINT).
	if errUser != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"STATUS": "Invalid User ID...",
			"ERROR":  errUser.Error(),
		})
		c.Abort()
		return
	}

	// VALIDAMOS QUE EL ID DEL USUARIO EXISTA, SI NO, DEVUELVE ERROR.
	var userExist domain.User
	userExistResutlt := postgres.DB.First(&userExist, "id	=	?", userId)
	if userExist.ID == 0 || userExistResutlt.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"STATUS": "User ID Not Found...",
			"ERROR":  userExistResutlt.Error.Error(),
		})
		c.Abort()
		return
	}
	// ===========================================================================================
	// ===========================================================================================
	// =========================================== QUERY =========================================
	// QUERY PARA ENCONTRAR SECCIONES POR NOMBRE
	passwordName := strings.ToUpper(c.Query("namePass"))
	var sectionExist domain.Password
	passwordExistResutlt := postgres.DB.Where("name	=	?", passwordName).First(&sectionExist)
	if sectionExist.ID == 0 || passwordExistResutlt.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"STATUS": "Password not found...",
			"ERROR":  passwordExistResutlt.Error.Error(),
		})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"STATUS":   "Password is here...",
		"PASSWORD": sectionExist,
	})
}

// VERDE
func PasswordPost(c *gin.Context) {
	// BODY.
	var body struct {
		Name                    string
		Description             string
		Password                string
		SectionParentIdPassword *uint
	}

	// ===========================================================================================
	// ===========================================================================================
	// =========================================== VALIDACIONES ==================================
	// VALIDACION DEL BODY PARA LA CREACION DE LA CONSTRASEÑA.
	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to read body..."})
	}

	// VALIDAMOS QUE EL TIPO DE DATO DEL ID DEL USUARIO SEA EL CORRECTO (TIPO UINT).
	idStr_user := c.Param("id") // -- OBTENEMOS EL ID DESDE EL URL. ---
	userId, err := strconv.Atoi(idStr_user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid User ID..."})
		c.Abort()
		return
	}

	// VALIDAMOS QUE EL ID DEL USUARIO EXISTA, SI NO, DEVUELVE ERROR.
	var userExist domain.User
	userValidationResult := postgres.DB.First(&userExist, "id = ?", userId)
	if userExist.ID == 0 || userValidationResult.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"STATUS": "User ID not found...",
			"ERROR":  userValidationResult.Error.Error(),
		})
		c.Abort()
		return
	}

	// VALIDAMOS QUE LA SECCION PADRE EXISTA -- VERIFICAR AUN --
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
	//
	passwordNew := domain.Password{
		Name:                    strings.ToUpper(body.Name),
		Description:             &body.Description,
		Password:                body.Password,
		SectionParentIdPassword: body.SectionParentIdPassword,
	}

	passwordCreation := postgres.DB.Create(&passwordNew)

	if passwordCreation.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"STATUS": "Failed to save password on the database...",
			"ERROR":  passwordCreation.Error.Error(),
		})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, gin.H{"great-password": "Password created successfully..."})
}

// VERDE
func PasswordUpdate(c *gin.Context) {
	// OBTENCION DE LOS CAMPOS A ACTUALIZAR DESDE EL BODY.
	var body map[string]interface{}

	// OBTENCION DEL ID DEL USUARIO DESDE EL URL.
	userIdStr := c.Param("id")
	userId, errUser := strconv.Atoi(userIdStr)

	// OBTENCION DEL ID DE LA CONSTRASEÑA DESDE EL URL.
	passwordIdStr := c.Param("idU")
	passwordId, errPass := strconv.Atoi(passwordIdStr)
	// ===========================================================================================
	// ===========================================================================================
	// =========================================== VALIDACIONES ==================================
	// VALIDACION DE LOS DATOS DE BODY.
	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{"STATUS": "Failed to read body"})
	}

	// VALIDAMOS QUE EL TIPO DE DATO DEL ID DEL USUARIO SEA CORRECTO (TIPO UINT).
	if errUser != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"STATUS": "Invalid User ID...",
			"ERROR":  errUser.Error(),
		})
		c.Abort()
		return
	}

	// VALIDAMOS QUE EL ID DE LA CONSTRASEÑA SEA CORRECTO
	if errPass != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"STATUS": "Invalid Password ID...",
			"ERROR":  errPass.Error(),
		})
		c.Abort()
		return
	}

	// VALIDAMOS QUE EL ID DEL USUARIO EXISTA, SI NO, DEVUELVE ERROR.
	var userExist domain.User
	userExistResult := postgres.DB.First(&userExist, "id	=	?", userId)
	if userExist.ID == 0 || userExistResult.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"STATUS": "User ID Not Found...",
			"ERROR":  userExistResult.Error.Error(),
		})
		c.Abort()
		return
	}

	// ===========================================================================================
	// ===========================================================================================
	// =========================================== QUERY =========================================
	//
	var passwordExist domain.Password
	passwordUpdateResult := postgres.DB.Model(&passwordExist).Where("id	=	?", passwordId).Updates(body)
	if passwordUpdateResult.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"STATUS": "Not able to update the current password",
			"ERROR":  passwordUpdateResult.Error.Error(),
		})
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
	userIdStr := c.Param("id")
	userId, errUser := strconv.Atoi(userIdStr)

	passwordIdStr := c.Param("idD")
	passwordId, errPass := strconv.Atoi(passwordIdStr)
	// ===========================================================================================
	// ===========================================================================================
	// =========================================== VALIDACIONES ==================================
	// VALIDAMOS QUE EL TIPO DE DATO DEL ID DEL USUARIO SEA CORRECTO (TIPO UINT).
	if errUser != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"STATUS": "Invalid User ID...",
			"ERROR":  errUser.Error(),
		})
	}

	// VALIDAMOS QUE EL ID DEL USUARIO EXISTA, SI NO, DEVUELVE ERROR.
	var userExist domain.User
	userExistResult := postgres.DB.First(&userExist, "id	=	?", userId)
	if userExist.ID == 0 || userExistResult.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"STATUS": "User ID Not Found...",
			"ERROR":  userExistResult.Error.Error(),
		})
	}

	// VALIDAMOS QUE EL ID DE LA CONTRASEÑA TENGA EL FORMATO CORRECTO (TIPO UINT).
	if errPass != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"STATUS": "Invalid ID...",
			"ERROR":  errPass.Error(),
		})
		c.Abort()
		return
	}

	// ===========================================================================================
	// ===========================================================================================
	// =========================================== QUERY =========================================
	//
	var passwordExist domain.Password
	passwordDeleteResult := postgres.DB.Delete(&passwordExist, passwordId)
	if passwordDeleteResult.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"STATUS": "Error trying to delete the password...",
			"ERROR":  passwordDeleteResult.Error.Error(),
		})
		c.Abort()
		return
	}

	if passwordDeleteResult.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"STATUS": "Password not found...",
		})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, gin.H{"STATUS": "Password deleted correctly..."})
}
