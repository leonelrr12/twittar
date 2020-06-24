package routers

import (
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/leonelrr12/twittar/bd"
	"github.com/leonelrr12/twittar/models"
)

/*SubirBanner Routers para subir el Banner*/
func SubirBanner(w http.ResponseWriter, r *http.Request) {

	file, handler, err := r.FormFile("banner")
	var extension = strings.Split(handler.Filename, ".")[1]
	var archivo string = "uploads/banners/" + IDUsuario + "." + extension

	f, err := os.OpenFile(archivo, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		http.Error(w, "Error al subir la imagen! "+err.Error(), http.StatusCreated)
		return
	}

	_, err = io.Copy(f, file)
	if err != nil {
		http.Error(w, "Error copiando la imagen! "+err.Error(), http.StatusCreated)
		return
	}

	var usuario models.Usuario
	var status bool

	usuario.Banner = IDUsuario + "." + extension
	status, err = bd.ModifyRegistro(usuario, IDUsuario)
	if err != nil || status == false {
		http.Error(w, "Error al grabar el Banner en la BD "+err.Error(), http.StatusCreated)
		return
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
}
