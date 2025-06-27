package authenticationusecase

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/maxexee/rugaPasswordManager/internal/dto"
	authenticationrespository "github.com/maxexee/rugaPasswordManager/internal/repository/user_repository/authentication_respository"
)

func LogInUseCase(user *dto.LogInDTO) (bool, string, error) {
	// ===========================================================================================
	// ===========================================================================================
	// =========================================== VALIDACIONES ==================================
	//
	ok, userExist, userError := authenticationrespository.LogInRepository(*user)
	if !ok {
		return false, "", userError
	}

	// ===========================================================================================
	// ===========================================================================================
	// =========================================== TOKEN =========================================
	/*  CREACION DEL TOKEN
	- DONDE *NewWithClaims()* RECIBE DOS VALORES, EL ALGORITMO CON EL QUE SE VA A FIRMAR EL TOKEN
	Y TAMBIEN SE LE AGREGAN LOS CLAIMS (DATOS EN EL PAYLOAD)

	- time.Now() -> HORA EN LA QUE SE CREA EL TOKEN.
	- .Add(time.<>) -> SE LE AGREGA EL TIEMPO QUE VA A TENER DE VALIDO EL TOKEN.
	- .Unix() -> LO CONVIERTE A UN ENTERO UNIX TIMESTAMP(SEGUNDOS).
	*/
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": userExist.ID,
		"exp": time.Now().Add(time.Hour * 8).Unix(),
	})

	// UNA VEZ CREADO EL OBJETO JWT, SE FIRMA CON NUESTRA CLAVE SECRETA, ESTO RETORNA EL TOKEN EN
	// FORMA DE STRING Y SI HAY UN ERROR, EL ERROR.
	tokeString, tokenError := token.SignedString([]byte(os.Getenv("SECRET")))
	if tokenError != nil {
		return false, "", tokenError
	}

	return true, tokeString, nil
}
