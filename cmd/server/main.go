package main

import (
	"github.com/gin-gonic/gin"
	postgres "github.com/maxexee/rugaPasswordManager/infrastructure/db"
	"github.com/maxexee/rugaPasswordManager/internal/handler/passwords"
	"github.com/maxexee/rugaPasswordManager/internal/handler/sections"
	handler "github.com/maxexee/rugaPasswordManager/internal/handler/user/authentication"
	"github.com/maxexee/rugaPasswordManager/internal/handler/user/authorization/middleware"
	"github.com/maxexee/rugaPasswordManager/internal/initializers"
)

func init() {
	initializers.EnvLoader()
	postgres.DbPostgresConnection()
	initializers.DB_migratetion()
}

func main() {
	router := gin.Default()
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
	middleware.SecureRoutesMiddleware(router)

	router.Run() // listen and serve on 0.0.0.0:8080
}
