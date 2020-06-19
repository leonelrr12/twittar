package bd

import "golang.org/x/crypto/bcrypt"

/*EncryptPassword  Funcion para Encriptar el Password */
func EncryptPassword(pass string) (string, error) {
	// 8 Para Admin
	// 6 Para Usuario Comun
	costo := 6
	bytes, err := bcrypt.GenerateFromPassword([]byte(pass), costo)
	return string(bytes), err
}
