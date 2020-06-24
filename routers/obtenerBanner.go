package routers

import (
	"io"
	"net/http"
	"os"

	"github.com/leonelrr12/twittar/bd"
)

/*ObtenerBanner Router para obtener el Banner*/
func ObtenerBanner(w http.ResponseWriter, r *http.Request) {

	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "Debe enviar el parametro ID.", http.StatusCreated)
		return
	}

	perfil, err0 := bd.BuscoPerfil(ID)
	if err0 != nil {
		http.Error(w, "Usuario no encontrado.", http.StatusCreated)
		return
	}

	OpenFile, err := os.Open("uploads/banners/" + perfil.Banner)
	if err != nil {
		http.Error(w, "Imagen no encontrada.", http.StatusCreated)
		return
	}

	_, err = io.Copy(w, OpenFile)
	if err != nil {
		http.Error(w, "Error al copiar la Imagen.", http.StatusCreated)
	}
}
