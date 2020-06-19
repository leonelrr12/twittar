package middlew

import (
	"net/http"

	"github.com/leonelrr12/twittar/bd"
)

/*CheckDB Funcion Middleware para validar la conexion a la DB*/
func CheckDB(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if bd.CheckConnect() == false {
			http.Error(w, "Conexion perdida con la DB", 500)
			return
		}
		next.ServeHTTP(w, r)
	}
}
