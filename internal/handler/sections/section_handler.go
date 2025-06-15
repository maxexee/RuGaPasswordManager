package sections

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/maxexee/rugaPasswordManager/internal/dto"
	validations "github.com/maxexee/rugaPasswordManager/internal/handler/validations"
	sectionsusecase "github.com/maxexee/rugaPasswordManager/internal/use_case/sections_use_case"
)

var SectionBodyDto dto.SectionDto

// VERDE...
func SectionGet(c *gin.Context) {
	// OBTENCION DEL ID DEL USUARIO DESDE EL URL.
	userIdStr := c.Param("id")

	// OBTENCION DEL ID DE LA SECCION DESDE EL URL.
	sectionIdStr := strings.ToLower(c.Query("section_parent_id"))

	// ===========================================================================================
	// ===========================================================================================
	// =========================================== QUERY =========================================
	// LLAMADA AL CASO DE USO PARA OBTENER TODAS LAS SECCIONES YA SEA LA RAIZ (NULL) O LAS SECCIONES HIJAS,
	// Y/O LAS CONSTRASEÑAS HIJAS DE UNA SECCION PADRE.
	ok, sectionsPasswordsReturn, sectionsPasswordsReturnError := sectionsusecase.SectionGetAllUseCase(userIdStr, sectionIdStr)
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
	// OBTENCION DEL ID EL USUARIO.
	userIdStr := c.Param("id")

	// OBTENCION DEL NOMBRE DE LA SECCION A BUSCAR.
	sectionNameStr := strings.ToUpper(c.Query("nameSec"))

	// LLAMAD AL USE CASE.
	ok, sectionReturn, sectionReturnError := sectionsusecase.SectionGetByNameUseCase(userIdStr, sectionNameStr)
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
	// ===========================================================================================
	// =========================================== BODY ==========================================
	// OBTENCION DEL ID EL USUARIO.
	userIdStr := c.Param("id")

	// ===========================================================================================
	// ===========================================================================================
	// =========================================== VALIDACIONES ==================================
	// === VALIDAMOS QUE NO HAYA ERROR EN EL BODY Y TODO SEA LEGIBLE. ===
	if !validations.BodyValidation(c, &SectionBodyDto) {
		c.Abort()
		return
	}

	// LLAMADA AL USE CASE.
	ok, sectionCreated, sectionCreatedError := sectionsusecase.SectionPostUseCase(userIdStr, &SectionBodyDto)
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
	// OBTENCION DEL ID DEL USUARIO DESDE EL URL.
	userIdStr := c.Param("id")

	// OBTENCION DEL ID DE LA SECCION DESDE EL URL.
	sectionIdStr := c.Param("idSU")

	// VALIDACION DEL BODY.
	if !validations.BodyValidation(c, &SectionBodyDto) {
		c.Abort()
		return
	}

	// LLAMDA AL USE CASE.
	ok, sectionReturn, sectionReturnError := sectionsusecase.SectionUpdateUseCase(userIdStr, sectionIdStr, &SectionBodyDto)
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
	// OBTENCION DEL ID DEL USUARIO DESDE EL URL.
	userIdStr := c.Param("id")

	// OBTENCION DEL ID DE LA SECCION DESDE EL URL.
	sectionIdStr := c.Param("idD")

	// LLAMDAO DEL USE CASE.
	ok, sectionDeleteError := sectionsusecase.SectionDeleteUseCase(userIdStr, sectionIdStr)
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
