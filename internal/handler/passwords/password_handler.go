package passwords

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/maxexee/rugaPasswordManager/internal/dto"
	"github.com/maxexee/rugaPasswordManager/internal/handler/utils"
	validations "github.com/maxexee/rugaPasswordManager/internal/handler/validations"
	passwordusecase "github.com/maxexee/rugaPasswordManager/internal/use_case/password_use_case"
)

// VARIBLE DE INICIALIZACION DEL *PasswordBodyDto*.
var PasswordBodyDto dto.PasswordDto

// VERDE...
func PasswordGetById(c *gin.Context) {
	ok, userId := utils.SaveUserIdFromClaims(c)
	if !ok {
		c.Abort()
		return
	}

	// OBTENCION DEL ID DE LA CONTRASEÑA DESDE EL URL.
	passwordIdStr := c.Param("idPG")

	// LLAMDA AL USE CASE.
	ok, passwordReturn, passwordReturnError := passwordusecase.PasswordGetByIdUseCase(userId, passwordIdStr)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{
			"STATUS": "ERROR...",
			"ERROR":  passwordReturnError.Error(),
		})
		c.Abort()
		return
	}

	// SI TODO SALE BIEN.
	c.JSON(http.StatusOK, gin.H{
		"STATUS":   "Password found succesfully...",
		"PASSWORD": passwordReturn,
	})
}

// VERDE...
func PasswordGetByName(c *gin.Context) {
	ok, userId := utils.SaveUserIdFromClaims(c)
	if !ok {
		c.Abort()
		return
	}

	// OBTENCION DEL NOMBRE DE LA CONTRASEÑA DESDE EL URL.
	passwordName := strings.ToUpper(c.Query("namePass"))

	// LLAMADA AL USE CASE.
	ok, passwordReturn, passwordReturnError := passwordusecase.PasswordGetByNameUseCase(userId, passwordName)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{
			"STATUS": "ERROR",
			"ERROR":  passwordReturnError.Error(),
		})
		c.Abort()
		return
	}

	// SI TODO SALE BIEN.
	c.JSON(http.StatusOK, gin.H{
		"STATUS":   "Password is here...",
		"PASSWORD": passwordReturn,
	})
}

// VERDE...
func PasswordPost(c *gin.Context) {
	ok, userId := utils.SaveUserIdFromClaims(c)
	if !ok {
		c.Abort()
		return
	}

	// VALIDACION DE LOS DATOS DE BODY.
	if !validations.BodyValidation(c, &PasswordBodyDto) {
		c.Abort()
		return
	}

	// LLAMADO AL USE CASE.
	ok, passwordReturn, passwordReturnError := passwordusecase.PasswordPostUseCase(userId, &PasswordBodyDto)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{
			"STATUS": "ERROR...",
			"ERROR":  passwordReturnError.Error(),
		})
		c.Abort()
		return
	}

	// SI TODO SALE BIEN...
	c.JSON(http.StatusOK, gin.H{
		"great-password": "Password created successfully...",
		"PASSWORD":       passwordReturn,
	})
}

// VERDE...
func PasswordUpdate(c *gin.Context) {
	ok, userId := utils.SaveUserIdFromClaims(c)
	if !ok {
		c.Abort()
		return
	}

	// OBTENCION DEL ID DE LA CONTRASEÑA DESDE EL URL.
	passwordIdStr := c.Param("idPU")

	// VALIDACION DE LOS DATOS DE BODY.
	if !validations.BodyValidation(c, &PasswordBodyDto) {
		c.Abort()
		return
	}

	// LLAMADO AL USE CASE.
	ok, passwordReturn, passwordReturnError := passwordusecase.PasswordUpdateUseCase(userId, passwordIdStr, &PasswordBodyDto)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{
			"STATUS": "ERROR...",
			"ERROR":  passwordReturnError.Error(),
		})
		c.Abort()
		return
	}

	// SI TODO SALE BIEN...
	c.JSON(http.StatusOK, gin.H{
		"STATUS":   "Password update has been done correct...",
		"PASSWORD": passwordReturn,
	})
}

// VERDE...
func PasswordDelete(c *gin.Context) {
	ok, userId := utils.SaveUserIdFromClaims(c)
	if !ok {
		c.Abort()
		return
	}

	// OBTENCION DEL ID DE LA CONTRASEÑA DESDE EL URL.
	passwordIdStr := c.Param("idPD")

	// LLAMDO AL USE CASE.
	ok, passwordDeleteError := passwordusecase.PasswordDeleteUseCase(userId, passwordIdStr)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{
			"STATUS": "ERROR...",
			"ERROR":  passwordDeleteError.Error(),
		})
		c.Abort()
		return
	}

	// SI TODO SALE BIEN...
	c.JSON(http.StatusOK, gin.H{
		"STATUS": "Password deleted successfully...",
	})
}
