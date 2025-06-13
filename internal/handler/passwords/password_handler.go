package passwords

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/maxexee/rugaPasswordManager/internal/dto"
	validations "github.com/maxexee/rugaPasswordManager/internal/handler/validations"
	passwordusecase "github.com/maxexee/rugaPasswordManager/internal/use_case/password_use_case"
)

// VARIBLE DE INICIALIZACION DEL *PasswordBodyDto*.
var PasswordBodyDto dto.PasswordBodyDto

// VERDE...
func PasswordGetById(c *gin.Context) {
	// OBTENCION DEL ID DEL USUARIO DESDE EL URL.
	userIdStr := c.Param("id")

	// OBTENCION DEL ID DE LA CONTRASEÑA DESDE EL URL.
	passwordIdStr := c.Param("idPG")

	// LLAMDA AL USE CASE.
	ok, passwordReturn, passwordReturnError := passwordusecase.PasswordGetByIdUseCase(userIdStr, passwordIdStr)
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

// ...
func PasswordGetByName(c *gin.Context) {
	// OBTENCION DEL ID DEL USUARIO DESDE EL URL.
	userIdStr := c.Param("id")

	// OBTENCION DEL NOMBRE DE LA CONTRASEÑA DESDE EL URL.
	passwordName := strings.ToUpper(c.Query("namePass"))

	// LLAMADA AL USE CASE.
	ok, passwordReturn, passwordReturnError := passwordusecase.PasswordGetByNameUseCase(userIdStr, passwordName)
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
	// OBTENCION DEL ID DEL USUARIO DESDE EL URL.
	userIdStr := c.Param("id")

	// VALIDACION DE LOS DATOS DE BODY.
	if !validations.BodyValidation(c, &PasswordBodyDto) {
		c.Abort()
		return
	}

	// CONTRUCCION DEL DTO BODY.
	dtoSend := dto.PasswordBodyDto{
		Name:                    PasswordBodyDto.Name,
		Description:             PasswordBodyDto.Description,
		Password:                PasswordBodyDto.Password,
		SectionParentIdPassword: PasswordBodyDto.SectionParentIdPassword,
	}

	// LLAMADO AL USE CASE.
	ok, passwordReturn, passwordReturnError := passwordusecase.PasswordPostUseCase(userIdStr, &dtoSend)
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
	// OBTENCION DEL ID DEL USUARIO DESDE EL URL.
	userIdStr := c.Param("id")

	// OBTENCION DEL ID DE LA CONTRASEÑA DESDE EL URL.
	passwordIdStr := c.Param("idPU")

	// VALIDACION DE LOS DATOS DE BODY.
	if !validations.BodyValidation(c, &PasswordBodyDto) {
		c.Abort()
		return
	}

	// CONTRUCCION DEL DTO BODY.
	dtoSend := dto.PasswordBodyDto{
		Name:                    PasswordBodyDto.Name,
		Description:             PasswordBodyDto.Description,
		Password:                PasswordBodyDto.Password,
		SectionParentIdPassword: PasswordBodyDto.SectionParentIdPassword,
	}

	// LLAMADO AL USE CASE.
	ok, passwordReturn, passwordReturnError := passwordusecase.PasswordUpdateUseCase(userIdStr, passwordIdStr, &dtoSend)
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
	// OBTENCION DEL ID DEL USUARIO DESDE EL URL.
	userIdStr := c.Param("id")

	// OBTENCION DEL ID DE LA CONTRASEÑA DESDE EL URL.
	passwordIdStr := c.Param("idPD")

	// LLAMDO AL USE CASE.
	ok, passwordDeleteError := passwordusecase.PasswordDeleteUseCase(userIdStr, passwordIdStr)
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
