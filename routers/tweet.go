package routers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/leonelrr12/twittar/bd"
	"github.com/leonelrr12/twittar/models"
)

/*Tweet : Funcion para crear Tweet en DB*/
func Tweet(w http.ResponseWriter, r *http.Request) {

	var mensaje models.msgTweet
	err := json.NewDecoder(r.Body).Decode(&mensaje)
	if err != nil {
		http.Error(w, "Error en los datos recibidos "+err.Error(), 400)

		registro := models.Tweet{
			UserID:  IDUsuario,
			Mensaje: mensaje.Mensaje,
			Fecha:   time.Now(),
		}

		_, status, err := bd.InsertTweet(registro)
		if err != nil {
			http.Error(w, "Ocurrio un error al grabar el Registro! "+err.Error(), 400)
			return
		}

		if status == false {
			http.Error(w, "No se ha logrado guardar el Tweet! ", 400)
			return
		}

		w.WriteHeader(http.StatusCreated)
	}
}
