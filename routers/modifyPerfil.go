package routers

import (
	"encoding/json"
	"net/http"

	"github.com/leonelrr12/twittar/bd"
	"github.com/leonelrr12/twittar/models"
)

/*ModifyPerfil : Funcion para crear Usuario en DB*/
func ModifyPerfil(w http.ResponseWriter, r *http.Request) {

	var t models.Usuario

	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(w, "Datos Incorrectos. "+err.Error(), 400)
		return
	}

	var status bool
	status, err = bd.ModifyRegistro(t, IDUsuario)
	if err != nil {
		http.Error(w, "Ocurrio un error al intentar Modificar el registro. "+err.Error(), 400)
		return
	}

	if status == false {
		http.Error(w, "No se ha logrado modificar el registro del usuario ", 400)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
