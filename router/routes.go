package router

import (
	"github.com/gin-gonic/gin"
	handler "github.com/maxexee/rugaPasswordManager/internal/handler/user/authentication"
	"github.com/maxexee/rugaPasswordManager/router/middleware"
)

func InitRoutes() *gin.Engine {
	router := gin.Default()

	// RUTAS PARA CREACION Y ACCESO DE USUARIO.
	router.POST("/signup", handler.SignUp)
	router.POST("/login", handler.Login)

	// RUTAS ASEGURADAS.
	middleware.SecureRoutesMiddleware(router)

	return router
}
