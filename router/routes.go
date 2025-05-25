package router

import (
	"github.com/gin-gonic/gin"
	"github.com/maxexee/rugaPasswordManager/internal/handler/passwords"
	"github.com/maxexee/rugaPasswordManager/internal/handler/sections"
	handler "github.com/maxexee/rugaPasswordManager/internal/handler/user/authentication"
	"github.com/maxexee/rugaPasswordManager/internal/handler/user/authorization/middleware"
)

func InitRoutes() *gin.Engine {
	router := gin.Default()

	// RUTAS PARA CREACION Y ACCESO DE USUARIO.
	router.POST("/signup", handler.SignUp)
	router.POST("/login", handler.Login)

	// RUTAS DE SECCIONES (ASEGURARLAS DESPUES AL "SecureRoutesMiddleware").
	router.POST("/user/:id/section/createSec", sections.SectionPost)
	router.GET("/user/:id/section/allSec", sections.SectionGet)
	router.GET("/user/:id/section/byNameSec", sections.SectionGetByName)
	router.PATCH("/user/:id/section/updateSec/:idU", sections.SectionUpdate)
	router.DELETE("/user/:id/section/deleteSec/:idD", sections.SectionDelete)

	// RUTAS DE LAS CONTRASEÑAS (ASEGURARLAS DESPUES AL "SecureRoutesMiddleware").
	router.POST("/user/:id/section/passwd", passwords.PasswordPost)
	router.GET("/user/:id/passwd/byNamePass", passwords.PasswordGetByName)
	router.PATCH("/user/:id/passwd/updatePass/:idU", passwords.PasswordUpdate)
	router.DELETE("/user/:id/passwd/delPass/:idD", passwords.PasswordDelete)

	// RUTAS ASEGURADAS.
	middleware.SecureRoutesMiddleware(router)

	return router
}
