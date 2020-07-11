package routers

import (
	"errors"
	"strings"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/leonelrr12/twittar/bd"
	"github.com/leonelrr12/twittar/models"
)

/*Email Email del usuario*/
var Email string

/*IDUsuario ID interno del usuario*/
var IDUsuario string

/*ProcesoToken Proceso token para extraer sus valores*/
func ProcesoToken(tk string) (*models.Claim, bool, string, error) {

	miClave := []byte("El Guasimo de Los Santos")
	claims := &models.Claim{}

	splitToken := strings.Split(tk, "Bearer")
	if len(splitToken) != 2 {
		return claims, false, string(""), errors.New("formato de token invalido")
	}

	tk = strings.TrimSpace(splitToken[1])

	tkn, err := jwt.ParseWithClaims(tk, claims, func(token *jwt.Token) (interface{}, error) {
		return miClave, nil
	})

	if err != nil {
		if !tkn.Valid {
			return claims, false, string(""), errors.New("token invalido")
		}
		return claims, false, string(""), err
	}

	_, encontrado, ID := bd.ExisteUsuario(claims.Email)
	if encontrado == true {
		Email = claims.Email
		IDUsuario = ID
	}

	return claims, encontrado, IDUsuario, nil
}
