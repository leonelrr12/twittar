package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/leonelrr12/twittar/middlew"
	"github.com/leonelrr12/twittar/routers"
	"github.com/rs/cors"
)

/*Manejadores seteo mi puerto y ponco a escuchar el servidor*/
func Manejadores() {
	router := mux.NewRouter()

	router.HandleFunc("/registro", middlew.CheckDB(routers.Registro)).Methods("POST")
	router.HandleFunc("/login", middlew.CheckDB(routers.Login)).Methods("POST")
	router.HandleFunc("/verperfil", middlew.CheckDB(middlew.ValidoJWT(routers.VerPerfil))).Methods("GET")
	router.HandleFunc("/modifyperfil", middlew.CheckDB(middlew.ValidoJWT(routers.ModifyPerfil))).Methods("PUT")
	router.HandleFunc("/tweet", middlew.CheckDB(middlew.ValidoJWT(routers.Tweet))).Methods("POST")
	router.HandleFunc("/leotweets", middlew.CheckDB(middlew.ValidoJWT(routers.LeoTweets))).Methods("GET")
	router.HandleFunc("/borrotweet", middlew.CheckDB(middlew.ValidoJWT(routers.BorroTweet))).Methods("DELETE")

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}
	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}
