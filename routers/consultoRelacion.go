package routers

import (
	"encoding/json"
	"net/http"

	"github.com/leonelrr12/twittar/bd"
	"github.com/leonelrr12/twittar/models"
)

/*ConsultoRelacion Router para consultar la Realcion de 2 usuarios*/
func ConsultoRelacion(w http.ResponseWriter, r *http.Request) {

	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "El parametro ID es obligatorio.", http.StatusBadRequest)
		return
	}

	var t models.Relacion
	t.UsuarioID = IDUsuario
	t.UsuarioRelacionID = ID

	var resp models.ConsultoRelacion
	//resp := map[string]bool{"status": false}

	status, err := bd.ConsultoRelacion(t)
	if err != nil || status == false {
		resp.Status = false
		//resp = map[string]bool{"status": false}
	} else {
		resp.Status = true
		//resp = map[string]bool{"status": true}
	}

	w.Header().Set("Context-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(resp)
}
