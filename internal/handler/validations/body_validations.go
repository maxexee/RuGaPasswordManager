package validationsauthentication

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/maxexee/rugaPasswordManager/internal/initializers"
)

func BodyValidation(c *gin.Context, body any) bool {
	if err := c.ShouldBindJSON(body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"STATUS": "Failed to read body...",
			"ERROR":  err.Error(),
		})
		return false
	}

	if err := initializers.Validate.Struct(body); err != nil {
		var ve validator.ValidationErrors
		if errs, ok := err.(validator.ValidationErrors); ok {
			ve = errs
		} else if castErr := errors.As(err, &ve); castErr {
			/*
				EN TEORIA CON LA COMPRACION err.(validator.ValidationErrors) suele bastar,
				pero se deja este "errors.As" POR SI EL PAQUETE EN ALGUN CASE ENVIA OTRO WRAPPER.
			*/
		} else {
			/*
				SI NO ES VALIDO EL ValidationErrors, TAL VES SEA  InvalidValidationError
				(POR EJEMPLO, UNA REGEX MAL ESCRITA). SE DEVUELVE UN MENSAJE GENERICO.
			*/
			c.JSON(http.StatusBadRequest, gin.H{
				"STATUS": "Validation configuration error",
				"ERROR":  err.Error(),
			})
			return false
		}

		/*
			SI LLEGAMOS A ESTE PUNTO, "ve" CONTIENE LOS ERRORES DE VALIDACION POR CAMPO.
			CREANDO ASI UN map[string]string PARA ENVIARLE AL FRONT UN JSON AGRUPANDO LOS MENSAJES DE
			ERROR.
		*/
		out := make(map[string]string)
		for _, fe := range ve {
			fieldName := fe.Field()
			tagName := fe.Tag()

			var msg string
			switch tagName {
			case "matchesName":
				msg = "Solo se permiten letras, números y espacios..."
			case "required":
				msg = "Este campo es obligatorio..."
			case "min":
				msg = "El tamaño mínimo no se cumple..."
			case "max":
				msg = "El tamaño maximo no se cumple..."
			case "email":
				msg = "El email no cumple con el formato correcto..."
			default:
				msg = "Error en validación: `" + tagName + "`"
			}
			out[fieldName] = msg
		}

		c.JSON(http.StatusBadRequest, gin.H{
			"STATUS": "Validation failed",
			"ERRORS": out,
		})
		return false
	}

	// SI ESTAMOS AQUÍ, TOOD ESTA BIEN, RETORNAMOS UN TRUE.
	return true
}
