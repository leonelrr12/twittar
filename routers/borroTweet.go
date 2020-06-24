package routers

import (
	"log"
	"net/http"

	"github.com/leonelrr12/twittar/bd"
)

/*BorroTweet Func router para borrar un Tweet*/
func BorroTweet(w http.ResponseWriter, r *http.Request) {

	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "Debe enviar el parametro ID.", http.StatusBadRequest)
		return
	}

	err := bd.BorroTweet(ID, IDUsuario)
	if err != nil {
		http.Error(w, "Ocurrio un Error al intentar borrar el Tweet."+err.Error(), http.StatusBadRequest)
		return
	}

	log.Println(ID, "Tweet borrado satisfactoriamente.")
	
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
}
