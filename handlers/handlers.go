package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/Faydiamond/TwitterGo/middlew"
	"github.com/Faydiamond/TwitterGo/routers"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

//cuando use la api, este es mi elemento d etrabajo
func Manejadores() {
	router := mux.NewRouter()
	router.HandleFunc("/registro", middlew.ChequeoBD(routers.Registro)).Methods("POST")
	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}
	handler := cors.AllowAll().Handler(router) //todos tienen poermiso
	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}
