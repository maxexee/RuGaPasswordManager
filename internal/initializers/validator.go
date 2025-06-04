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
	// QUE LOS NOMBRES TENGAN LETRAS, NUMEROS Y ESPACIOS.
	v.RegisterValidation("matchesName", func(fl validator.FieldLevel) bool {
		regex := fl.Param()
		re := regexp.MustCompile(regex)
		value := fl.Field().String()
		return re.MatchString(value)
	})

	// REGISTRO DE LA VALIDACION PERSONALIZADA "passwordsFormat" PARA VALIDAR
	// QUE LAS CONTRASEÑAS SEAN SEGURAS.
	v.RegisterValidation("passwordsFormat", func(fl validator.FieldLevel) bool {
		pwd := fl.Field().String()

		// CHEQUEO DE LOS VALORES OBTENIDOS CON RESPECTO A LAS REGEX.
		reLower := regexp.MustCompile(`[a-z]`)
		reUpper := regexp.MustCompile(`[A-Z]`)
		reDigit := regexp.MustCompile(`\d`)
		reSymbol := regexp.MustCompile(`[!@#\$%\^&\*]`)

		// SI ES FALSO, SE RETORNA "false".
		if !reLower.MatchString(pwd) {
			return false
		}
		if !reUpper.MatchString(pwd) {
			return false
		}
		if !reDigit.MatchString(pwd) {
			return false
		}
		if !reSymbol.MatchString(pwd) {
			return false
		}

		// Si pasa todos:
		return true
	})

	Validate = v
}
