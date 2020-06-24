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
	router.HandleFunc("/modifyerfil", middlew.CheckDB(middlew.ValidoJWT(routers.ModifyPerfil))).Methods("PUT")
	router.HandleFunc("/tweet", middlew.CheckDB(middlew.ValidoJWT(routers.Tweet))).Methods("POST")
	router.HandleFunc("/leotweets", middlew.CheckDB(middlew.ValidoJWT(routers.LeoTweets))).Methods("GET")
	router.HandleFunc("/borrotweet", middlew.CheckDB(middlew.ValidoJWT(routers.BorroTweet))).Methods("DELETE")

	router.HandleFunc("/subiravatar", middlew.CheckDB(middlew.ValidoJWT(routers.SubirAvatar))).Methods("POST")
	router.HandleFunc("/obteneravatar", middlew.CheckDB(routers.ObtenerAvatar)).Methods("GET")
	router.HandleFunc("/subirbanner", middlew.CheckDB(middlew.ValidoJWT(routers.SubirBanner))).Methods("POST")
	router.HandleFunc("/obtenerbanner", middlew.CheckDB(routers.ObtenerBanner)).Methods("GET")

	router.HandleFunc("/insertrelacion", middlew.CheckDB(middlew.ValidoJWT(routers.InsertRelacion))).Methods("POST")
	router.HandleFunc("/borrorelacion", middlew.CheckDB(middlew.ValidoJWT(routers.BorroRelacion))).Methods("DELETE")

	router.HandleFunc("/consultorelacion", middlew.CheckDB(middlew.ValidoJWT(routers.ConsultoRelacion))).Methods("GET")

	router.HandleFunc("/listausuarios", middlew.CheckDB(middlew.ValidoJWT(routers.ListaUsuarios))).Methods("GET")
	router.HandleFunc("/listatweetsseguidores", middlew.CheckDB(middlew.ValidoJWT(routers.LeoTweetsSeguidores))).Methods("GET")

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}
	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}
