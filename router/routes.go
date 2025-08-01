package router

import (
	"github.com/gin-gonic/gin"
	handler "github.com/maxexee/rugaPasswordManager/internal/handler/user/authentication"
	"github.com/maxexee/rugaPasswordManager/router/middleware"
)

func InitRoutes() *gin.Engine {
	router := gin.Default()

	// config := cors.Config{
	// 	AllowOrigins:     []string{"http://localhost:3000"},
	// 	AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
	// 	AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
	// 	ExposeHeaders:    []string{"Content-Length"},
	// 	AllowCredentials: true,
	// 	MaxAge:           12 * time.Hour,
	// }
	// router.Use(cors.New(config))

	// // 2) Manually atender OPTIONS para cualquier ruta (preflight)
	// router.OPTIONS("/*path", func(c *gin.Context) {
	// 	c.AbortWithStatus(http.StatusNoContent)
	// })

	// RUTAS PARA CREACION Y ACCESO DE USUARIO.
	router.POST("/signup", handler.SignUp)
	router.POST("/login", handler.Login)

	// RUTAS ASEGURADAS.
	middleware.SecureRoutesMiddleware(router)

	return router
}
