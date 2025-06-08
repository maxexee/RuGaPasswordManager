package sections

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/maxexee/rugaPasswordManager/internal/dto"
	validations "github.com/maxexee/rugaPasswordManager/internal/handler/validations"
	sectionsusecase "github.com/maxexee/rugaPasswordManager/internal/use_case/sections_use_case"
)

// VERDE
func SectionGet(c *gin.Context) {
	// ID DEL USUARIO DESDE EL URL.
	userIdStr := c.Param("id")

	// ID DE LA SECCION DESDE EL URL.
	sectionIdStr := strings.ToLower(c.Query("section_parent_id"))

	// ===========================================================================================
	// ===========================================================================================
	// =========================================== QUERY =========================================
	// LLAMADA AL CASO DE USO PARA OBTENER TODAS LAS SECCIONES, YA SEA LA RAIZ (NULL) O LAS SECCIONES HIJAS
	// DE UNA SECCION PADRE.
	ok, sectionsReturn, sectionReturnError := sectionsusecase.SectionGetAllUseCase(userIdStr, sectionIdStr)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{
			"STATUS": "ERROR ...",
			"ERROR":  sectionReturnError.Error(),
		})
		c.Abort()
		return
	}

	// SI TODO SALE BIEN.
	c.JSON(http.StatusOK, gin.H{
		"STATUS":   "All sections...",
		"SECTIONS": sectionsReturn,
	})
}

// VERDE
func SectionGetByName(c *gin.Context) {
	// OBTENCION DEL ID EL USUARIO.
	userIdStr := c.Param("id")

	// OBTENCION DEL NOMBRE DE LA SECCION A BUSCAR.
	sectionNameStr := strings.ToUpper(c.Query("nameSec"))

	// LLAMAD AL USE CASE.
	ok, sectionsReturn, sectionReturnError := sectionsusecase.SectionGetByNameUseCase(userIdStr, sectionNameStr)
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
		"SECTION": sectionsReturn,
	})
}

// VERDE
func SectionPost(c *gin.Context) {
	// ===========================================================================================
	// =========================================== BODY ==========================================
	// OBTENCION DEL ID EL USUARIO.
	userIdStr := c.Param("id")

	// DEFINIMOS QUE EL BODY QUE RECIBIREMOS DE LA PETICIÓN.
	var body dto.SectionDto

	// ===========================================================================================
	// ===========================================================================================
	// =========================================== VALIDACIONES ==================================
	// === VALIDAMOS QUE NO HAYA ERROR EN EL BODY Y TODO SEA LEGIBLE. ===
	if !validations.BodyValidation(c, &body) {
		c.Abort()
		return
	}

	// LLAMADA AL USE CASE.
	ok, sectionCreated, sectionCreatedError := sectionsusecase.SectionPostUseCase(userIdStr, &body)
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

// VERDE
func SectionUpdate(c *gin.Context) {
	// OBTENCION DEL ID EL USUARIO.
	userIdStr := c.Param("id")

	// OBTENCION DEL ID LA SECCION.
	sectionIdStr := c.Param("idU")

	// DEFINIMOS QUE EL BODY QUE RECIBIREMOS DE LA PETICIÓN.
	var body dto.SectionDto

	// VALIDACION DEL BODY.
	if !validations.BodyValidation(c, &body) {
		c.Abort()
		return
	}

	// LLAMDA AL USE CASE.
	ok, sectionReturn, sectionReturnError := sectionsusecase.SectionUpdateUseCase(userIdStr, sectionIdStr, &body)
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

// VERDE
func SectionDelete(c *gin.Context) {
	// OBTENCION DEL ID EL USUARIO.
	userIdStr := c.Param("id")

	// OBTENCION DEL ID DE LA SECCION.
	sectionIdStr := c.Param("idD")

	// LLAMDAO DEL USE CASE.
	ok, sectionReturnError := sectionsusecase.SectionDeleteUseCase(userIdStr, sectionIdStr)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{
			"STATUS": "ERROR...",
			"ERROR":  sectionReturnError.Error(),
		})
		c.Abort()
		return
	}

	// SI TODO SALE BIEN...
	c.JSON(http.StatusOK, gin.H{"STATUS": "ELIMINADO CORRECTAMENTE..."})
}
