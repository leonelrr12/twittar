package routers

import (
	"encoding/json"
	"net/http"

	"github.com/leonelrr12/twittar/bd"
	"github.com/leonelrr12/twittar/models"
)

/*Registro : Funcion para crear Usuario en DB*/
func Registro(w http.ResponseWriter, r *http.Request) {

	var t models.Usuario
	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(w, "Error en los datos recibidos "+err.Error(), 400)
	}

	if len(t.Email) == 0 {
		http.Error(w, "El Email de Usuario es Requerido", 400)
		return
	}

	if len(t.Password) < 6 {
		http.Error(w, "Largo del Password minimo es de 6 caracteres", 400)
		return
	}

	_, encontrado, _ := bd.ExisteUsuario(t.Email)
	if encontrado == true {
		http.Error(w, "Ya existe usuario con este Email en la DB", 400)
		return
	}

	_, status, err := bd.InsertRegistro(t)
	if err != nil {
		http.Error(w, "Ocurrio error al intentar Insertar el Registro del Usuario "+err.Error(), 400)
		return
	}

	if status == false {
		http.Error(w, "No se ha logrado intentar Insertar el Registro del Usuario", 400)
		return
	}

	return
}
