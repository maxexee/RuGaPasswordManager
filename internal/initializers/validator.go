package initializers

import (
	"regexp"

	"github.com/go-playground/validator/v10"
)

// VARIABLE QUE USAREMOS GLOBALMENTE.
var Validate *validator.Validate

func InitValidator() {
	v := validator.New()

	// REGISTRO DE LA VALIDACION PERSONALIZADA "matchesName" PARA VALIDAR
	// LOS NOMBRES TENGAN LETRAS, NUMEROS Y ESPACIOS.
	v.RegisterValidation("matchesName", func(fl validator.FieldLevel) bool {
		regex := fl.Param()
		re := regexp.MustCompile(regex)
		value := fl.Field().String()
		return re.MatchString(value)
	})
	Validate = v
}
