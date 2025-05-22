package sections

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	postgres "github.com/maxexee/rugaPasswordManager/infrastructure/db"
	"github.com/maxexee/rugaPasswordManager/internal/domain"
)

func SectionGet(c *gin.Context) {
	// ID DEL USUARIO DESDE EL URL Y SU CONVERSION A UINT.
	userIdStr := c.Param("id")
	userId, errUser := strconv.Atoi(userIdStr)

	// ID DE LA SECCION DESDE EL URL Y SU CONVERSION A UINT.
	sectionIdStr := strings.ToLower(c.Query("section_parent_id"))
	sectionId, errSection := strconv.Atoi(sectionIdStr)

	// ===========================================================================================
	// ===========================================================================================
	// =========================================== VALIDACIONES ==================================
	// VALIDAMOS QUE EL TIPO DE DATO DEL ID DEL USUARIO SEA CORRECTO (TIPO UINT).
	if errUser != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"STATUS":     "Invalid Type of User's ID...",
			"ERROR USER": errUser.Error(),
		})
		c.Abort()
		return
	}

	// VALIDAMOS QUE EL ID DEL USUARIO EXISTA, SI NO, DEVUELVE ERROR.
	var userExist domain.User
	userExistResutlt := postgres.DB.First(&userExist, "id = ?", userId)
	if userExist.ID == 0 || userExistResutlt.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"STATUS": "User ID not found...",
			"ERROR":  userExistResutlt.Error.Error(),
		})
		c.Abort()
		return
	}

	// ===========================================================================================
	// ===========================================================================================
	// =========================================== QUERY =========================================
	var sections []domain.Section
	var passwords []domain.Password

	// OBTENCIÓN DE LAS SECCIONES DESDE LA RAIZ.
	if sectionIdStr == "null" {
		sectionsGetAll := postgres.DB.Where("user_id = ?	AND	section_parent_id	IS NULL", userId).Find(&sections)
		if sectionsGetAll.Error != nil {
			c.JSON(http.StatusNotFound, gin.H{
				"STATUS": "Not sections founded...",
				"ERROR":  sectionsGetAll.Error.Error(),
			})
			c.Abort()
			return
		}
	} else {
		// OBTENCION DE LAS SECCIONES MEDIANTE UN PADRE.
		if errSection != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"STATUS":        "Invalid Type of Section's ID...",
				"ERROR SECTION": errSection.Error(),
			})
			c.Abort()
			return
		}

		sectionExistResutl := postgres.DB.First(&sections, "id	=	?", sectionId)
		if sectionExistResutl.Error != nil {
			c.JSON(http.StatusNotFound, gin.H{
				"STATUS": "Section's ID not Found",
				"ERROR":  sectionExistResutl.Error.Error(),
			})
			c.Abort()
			return
		}

		sectionsGetAll := postgres.DB.Where("user_id	=	?	AND	section_parent_id	=	?", userId, sectionId).Find(&sections)
		if sectionsGetAll.Error != nil {
			c.JSON(http.StatusNotFound, gin.H{
				"STATUS": "Not sections founded 1...",
				"ERROR":  sectionsGetAll.Error.Error(),
			})
			c.Abort()
			return
		}
	}

	passwordsGetAll := postgres.DB.Where("section_parent_id_password	=	?", sectionId).Find(&passwords)
	if passwordsGetAll.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"STATUS": "Not password founded...",
			"ERROR":  passwordsGetAll.Error.Error(),
		})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"STATUS":    "All sections...",
		"SECTIONS":  sections,
		"PASSWORDS": passwords,
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
	if errUser != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"STATUS": "Invalid User ID...",
			"ERROR":  errUser.Error(),
		})
	}

	// VALIDAMOS QUE EL ID DEL USUARIO EXISTA, SI NO, DEVUELVE ERROR.
	var userExist domain.User
	userCreationResult := postgres.DB.First(&userExist, "id = ?", userId)
	if userExist.ID == 0 || userCreationResult.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"STATUS": "User ID not found...",
			"ERROR":  userCreationResult.Error.Error(),
		})
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
	if sectionExist.ID == 0 && sectionCreationResult.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"STATUS": "Section not found",
			"ERROR":  sectionCreationResult.Error.Error(),
		})
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
		Name            string
		Description     string
		SectionParentId *uint
	}

	// ===========================================================================================
	// ===========================================================================================
	// =========================================== VALIDACIONES ==================================

	// === VALIDAMOS QUE NO HAYA ERROR EN EL BODY Y TODO SEA LEGIBLE. ===
	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to read body..."})
		return
	}

	// === VALIDAMOS QUE EL TIPO DE DATO DEL ID DEL USUARIO SEA EL CORRECTO (TIPO UINT). ===
	idStr := c.Param("id") // -- OBTENEMOS EL ID DESDE EL URL. ---
	userId, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID..."})
		c.Abort()
		return
	}

	// === VALIDAMOS QUE EL ID DEL USUARIO EXISTA, SI NO, DEVUELVE ERROR. ===
	var userExist domain.User
	userCreationResult := postgres.DB.First(&userExist, "id = ?", userId)
	if userExist.ID == 0 || userCreationResult.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"STATUS": "User ID not found...",
			"ERROR":  userCreationResult.Error.Error(),
		})
		c.Abort()
		return
	}

	// === VALIDAMOS QUE EL *SectionParentId* EXISTA, SI NO, QUE EL  *SectionParentId* SEA NULL ===
	var parentSectionExist domain.Section
	if body.SectionParentId != nil {
		parentResult := postgres.DB.First(&parentSectionExist, "id = ?", body.SectionParentId)
		if parentResult.Error != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"STATUS": "Parent section not found...",
				"ERROR":  parentResult.Error.Error(),
			})
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

	// SI HAY ALGUN ERROR AL MOMENTO DE GUARDAR LA NUEVA SECCION EN LA BASE DE DATOS.
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"STATUS": "Failed to save section on the Database...",
			"ERROR":  result.Error.Error(),
		})
		c.Abort()
		return
	}

	// RETORNAMOS QUE LA SECCION SE HA CREADO CORRECTAMENTE.
	c.JSON(http.StatusOK, gin.H{"great-section": "Section created successfully..."})
}

func SectionUpdate(c *gin.Context) {
	// BODY A RECIBIR DEL FRONT-END.
	var body map[string]interface{}
	// VALIDACION DEL BODY.
	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to read body..."})
		return
	}

	// ===========================================================================================
	// ===========================================================================================
	// VALIDAMOS QUE EL TIPO DE DATO DEL ID DEL USUARIO SEA EL CORRECTO (TIPO UINT).
	idStr := c.Param("id") // -- OBTENEMOS EL ID DESDE EL URL. ---
	userId, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID..."})
		c.Abort()
		return
	}

	// VALIDAMOS QUE EL ID DEL USUARIO EXISTA, SI NO, DEVUELVE ERROR.
	var userExist domain.User
	userCreationResult := postgres.DB.First(&userExist, "id = ?", userId)
	if userExist.ID == 0 || userCreationResult.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"STATUS": "User ID not found...",
			"ERROR":  userCreationResult.Error.Error(),
		})
		c.Abort()
		return
	}

	// ===========================================================================================
	// ===========================================================================================
	// =========================================== QUERY =========================================
	// OBTENCION DE EL ID DE LA SECCION A MODIFICAR.
	var section domain.Section
	idStrUpdate := c.Param("idU")
	sectionId, err := strconv.Atoi(idStrUpdate)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"STATUS": "Invalid ID",
			"ERROR":  err.Error(),
		})
	}

	//
	sectionGetById := postgres.DB.First(&section, idStrUpdate)
	if sectionGetById.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"STATUS": "Section not found...",
			"ERROR":  sectionGetById.Error.Error(),
		})
		c.Abort()
		return
	}

	// MODIFICACION DE LA SECCION.
	sectionUpdateResult := postgres.DB.Model(&section).Where("id	=	?", sectionId).Updates(body)
	if sectionUpdateResult.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"STATUS": "Not able to update the current section...",
			"ERROR":  sectionUpdateResult.Error.Error(),
		})
		c.Abort()
		return
	}

	// RETORNO AL DEL ESTADO OK AL FRON-END.
	c.JSON(http.StatusOK, gin.H{
		"STATUS": "Section update has been done correct...",
	})
}

func SectionDelete(c *gin.Context) {
	// ===========================================================================================
	// ===========================================================================================
	// ===========================================================================================
	// VALIDAMOS QUE EL TIPO DE DATO DEL ID DEL USUARIO SEA EL CORRECTO (TIPO UINT).
	idStr := c.Param("id") // -- OBTENEMOS EL ID DESDE EL URL. ---
	userId, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID..."})
		c.Abort()
		return
	}

	// VALIDAMOS QUE EL ID DEL USUARIO EXISTA, SI NO, DEVUELVE ERROR.
	var userExist domain.User
	userCreationResult := postgres.DB.First(&userExist, "id = ?", userId)
	if userExist.ID == 0 || userCreationResult.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"STATUS": "User ID not found...",
			"ERROR":  userCreationResult.Error.Error(),
		})
		c.Abort()
		return
	}
	// ===========================================================================================
	// ===========================================================================================
	// ===========================================================================================
	//
	var section domain.Section
	sectionIdStr := c.Param("idD")
	sectionId, errSection := strconv.Atoi(sectionIdStr)
	if errSection != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"STATUS": "Invalid ID...",
			"ERROR":  errSection.Error(),
		})
		c.Abort()
		return
	}

	//
	sectionGetById := postgres.DB.First(&section, sectionId)
	if sectionGetById.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"STATUS": "Section not found...",
			"ERROR":  sectionGetById.Error.Error(),
		})
		c.Abort()
		return
	}

	//
	sectionDeleteResutl := postgres.DB.Delete(&section, sectionId)
	if sectionDeleteResutl.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"STATUS": "Error trying to delete the Section...",
			"ERROR":  sectionDeleteResutl.Error.Error(),
		})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, gin.H{"STATUS": "ELIMINADO CORRECTAMENTE..."})
}
