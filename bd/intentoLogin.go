package bd

import (
	"github.com/leonelrr12/twittar/models"
	"golang.org/x/crypto/bcrypt"
)

/*IntentoLogin realiza el chequeo de login en la BD*/
func IntentoLogin(email string, password string) (models.Usuario, bool) {
	u, encontrado, _ := ExisteUsuario(email)
	if encontrado == false {
		return u, false
	}

	passwordBytes := []byte(password)
	passwordDB := []byte(u.Password)
	err := bcrypt.CompareHashAndPassword(passwordDB, passwordBytes)
	if err != nil {
		return u, false
	}
	return u, true
}
