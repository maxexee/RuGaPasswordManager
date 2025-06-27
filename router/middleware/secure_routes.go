package middleware

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/maxexee/rugaPasswordManager/internal/handler/passwords"
	"github.com/maxexee/rugaPasswordManager/internal/handler/sections"
	"github.com/maxexee/rugaPasswordManager/internal/handler/user/authorization"
)

func SecureRoutesMiddleware(e *gin.Engine) {
	fmt.Println("=== GRUPO DE LAS RUTAS SEGURAS ===")

	protected := e.Group(("/"))
	protected.Use(authorization.UserAuthorization)

	// RUTAS DE SECCIONES.
	sectionGroup := protected.Group("/section")
	{
		sectionGroup.POST("/createSec", sections.SectionPost)
		sectionGroup.GET("/allSec", sections.SectionGetAll)
		sectionGroup.GET("/byNameSec", sections.SectionGetByName)
		sectionGroup.PATCH("/updateSec/:idSU", sections.SectionUpdate)
		sectionGroup.DELETE("/deleteSec/:idD", sections.SectionDelete)
	}

	//1 RUTAS DE LAS CONTRASEÑAS.
	passwordGroup := protected.Group("/passwd")
	{
		passwordGroup.POST("/createPass", passwords.PasswordPost)
		passwordGroup.GET("/byIdPass/:idPG", passwords.PasswordGetById)
		passwordGroup.GET("/byNamePass", passwords.PasswordGetByName)
		passwordGroup.PATCH("/updatePass/:idPU", passwords.PasswordUpdate)
		passwordGroup.DELETE("/delPass/:idPD", passwords.PasswordDelete)
	}

	protected.GET("/validate", func(c *gin.Context) {
		userID, _ := c.Get("user_id")
		c.JSON(http.StatusOK, gin.H{
			"message": "Valid token...",
			"user_id": userID,
		})
	})
}
