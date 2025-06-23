package sections

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/maxexee/rugaPasswordManager/internal/dto"
	"github.com/maxexee/rugaPasswordManager/internal/handler/utils"
	validations "github.com/maxexee/rugaPasswordManager/internal/handler/validations"
	sectionsusecase "github.com/maxexee/rugaPasswordManager/internal/use_case/sections_use_case"
)

var SectionBodyDto dto.SectionDto

// VERDE...
func SectionGetAll(c *gin.Context) {
	ok, userId := utils.SaveUserIdFromClaims(c)
	if !ok {
		c.Abort()
		return
	}

	// OBTENCION DEL ID DE LA SECCION DESDE EL URL.
	sectionIdStr := strings.ToLower(c.Query("section_parent_id"))

	// ===========================================================================================
	// ===========================================================================================
	// =========================================== QUERY =========================================
	// LLAMADA AL CASO DE USO PARA OBTENER TODAS LAS SECCIONES YA SEA LA RAIZ (NULL) O LAS SECCIONES HIJAS,
	// Y/O LAS CONSTRASEÑAS HIJAS DE UNA SECCION PADRE.
	ok, sectionsPasswordsReturn, sectionsPasswordsReturnError := sectionsusecase.SectionGetAllUseCase(userId, sectionIdStr)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{
			"STATUS": "ERROR ...",
			"ERROR":  sectionsPasswordsReturnError.Error(),
		})
		c.Abort()
		return
	}

	// SI TODO SALE BIEN.
	c.JSON(http.StatusOK, gin.H{
		"STATUS":   "All sections...",
		"SECTIONS": sectionsPasswordsReturn,
	})
}

// VERDE...
func SectionGetByName(c *gin.Context) {
	ok, userId := utils.SaveUserIdFromClaims(c)
	if !ok {
		c.Abort()
		return
	}

	// OBTENCION DEL NOMBRE DE LA SECCION A BUSCAR.
	sectionNameStr := strings.ToUpper(c.Query("nameSec"))

	// LLAMAD AL USE CASE.
	ok, sectionReturn, sectionReturnError := sectionsusecase.SectionGetByNameUseCase(userId, sectionNameStr)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{
			"STATUS": "ERROR...",
			"ERROR":  sectionReturnError.Error(),
		})
		c.Abort()
		return
	}

	// SI TODO SALE BIEN...
	c.JSON(http.StatusOK, gin.H{
		"STATUS":  "Section is here",
		"SECTION": sectionReturn,
	})
}

// VERDE...
func SectionPost(c *gin.Context) {
	// OBTENCION DEL ID DEL USUARIO.
	// SE RECUPERA EL "user_id" DEL CONTEXTO, ES DECIR, DE LOS CLAIMS DEL TOKEN DEL LOGIN.
	userIdClaims, ok := c.Get("user_id")
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{
			"STATUS": "Unauthorized",
			"ERROR":  "There is no user_id on the token...",
		})
		c.Abort()
		return
	}

	// SE VERIFICA QUE SEA ENTERO/INT.
	// SI SI ES, ENTONCES SE ALMACENA EN "userId".
	userId, ok := userIdClaims.(int)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{
			"STATUS": "Unauthorized",
			"ERROR":  "user_id in context is not an integer...",
		})
		c.Abort()
		return
	}

	// === VALIDAMOS QUE NO HAYA ERROR EN EL BODY Y TODO SEA LEGIBLE. ===
	if !validations.BodyValidation(c, &SectionBodyDto) {
		c.Abort()
		return
	}

	// LLAMADA AL USE CASE.
	ok, sectionCreated, sectionCreatedError := sectionsusecase.SectionPostUseCase(userId, &SectionBodyDto)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{
			"STATUS": "ERROR...",
			"ERROR":  sectionCreatedError.Error(),
		})
		c.Abort()
		return
	}

	// SI TODO SALE BIEN.
	c.JSON(http.StatusOK, gin.H{
		"STATUS":  "Section created successfully...",
		"SECTION": sectionCreated,
	})
}

// VERDE...
func SectionUpdate(c *gin.Context) {
	ok, userId := utils.SaveUserIdFromClaims(c)
	if !ok {
		c.Abort()
		return
	}

	// OBTENCION DEL ID DE LA SECCION DESDE EL URL.
	sectionIdStr := c.Param("idSU")

	// VALIDACION DEL BODY.
	if !validations.BodyValidation(c, &SectionBodyDto) {
		c.Abort()
		return
	}

	// LLAMDA AL USE CASE.
	ok, sectionReturn, sectionReturnError := sectionsusecase.SectionUpdateUseCase(userId, sectionIdStr, &SectionBodyDto)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{
			"STATUS": "ERROR...",
			"ERROR":  sectionReturnError.Error(),
		})
		c.Abort()
		return
	}

	// SI TODO SALE BIEN...
	c.JSON(http.StatusOK, gin.H{
		"STATUS":  "Section update has been done correct...",
		"SECTION": sectionReturn,
	})
}

// VERDE...
func SectionDelete(c *gin.Context) {
	ok, userId := utils.SaveUserIdFromClaims(c)
	if !ok {
		c.Abort()
		return
	}

	// OBTENCION DEL ID DE LA SECCION DESDE EL URL.
	sectionIdStr := c.Param("idD")

	// LLAMDAO DEL USE CASE.
	ok, sectionDeleteError := sectionsusecase.SectionDeleteUseCase(userId, sectionIdStr)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{
			"STATUS": "ERROR...",
			"ERROR":  sectionDeleteError.Error(),
		})
		c.Abort()
		return
	}

	// SI TODO SALE BIEN...
	c.JSON(http.StatusOK, gin.H{"STATUS": "ELIMINADO CORRECTAMENTE..."})
}
