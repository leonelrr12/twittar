package jwt

import (
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/leonelrr12/twittar/models"
)

/*GeneroJWT Funcion para la generacion del Token*/
func GeneroJWT(t models.Usuario) (string, error) {

	miClave := []byte("El Guasimo de Los Santos")

	payload := jwt.MapClaims{
		"email":     t.Email,
		"nombre":    t.Nombre,
		"apellidos": t.Apellidos,
		"fecnac":    t.FecNac,
		"biografia": t.Biografia,
		"ubicacion": t.Ubicacion,
		"sitioweb":  t.SitioWeb,
		"_id":       t.ID.Hex(),
		"exp":       time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	tokenStr, err := token.SignedString(miClave)
	if err != nil {
		return tokenStr, err
	}
	return tokenStr, nil
}
