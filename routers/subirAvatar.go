package routers

import (
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/leonelrr12/twittar/bd"
	"github.com/leonelrr12/twittar/models"
)

/*SubirAvatar Routers para subir el avatar*/
func SubirAvatar(w http.ResponseWriter, r *http.Request) {

	file, handler, err := r.FormFile("avatar")
	var extension = strings.Split(handler.Filename, ".")[1]
	var archivo string = "uploads/avatars/" + IDUsuario + "." + extension

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

	usuario.Avatar = IDUsuario + "." + extension
	status, err = bd.ModifyRegistro(usuario, IDUsuario)
	if err != nil || status == false {
		http.Error(w, "Error al grabar el Avatar en la BD "+err.Error(), http.StatusCreated)
		return
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
}
