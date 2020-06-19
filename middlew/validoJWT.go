package middlew

import (
	"net/http"

	"github.com/leonelrr12/twittar/routers"
)

/*ValidoJWT Funcion Handler para validar el Token*/
func ValidoJWT(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_, _, _, err := routers.ProcesoToken(r.Header.Get("Authorization"))
		if err != nil {
			http.Error(w, "Error en el Tohen ! "+err.Error(), http.StatusBadRequest)
		}
		next.ServeHTTP(w, r)
	}
}
