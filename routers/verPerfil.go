package routers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/leonelrr12/twittar/bd"
)

/*VerPerfil Permite extraer los valores del Perfil*/
func VerPerfil(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("content-type", "application/json")

	ID := r.URL.Query().Get("id")
	log.Printf(ID)

	if len(ID) < 1 {
		http.Error(w, "Debe envisr el parametro ID.", http.StatusBadRequest)
		return
	}

	perfil, err := bd.BuscoPerfil(ID)
	if err != nil {
		http.Error(w, "Error al intenter buscar el Perfil de Usuario. "+err.Error(), 400)
		return
	}

	w.Header().Set("Context-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(perfil)
}
