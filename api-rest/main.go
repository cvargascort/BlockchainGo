package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", Index)
	router.HandleFunc("/contacto", Contacto)
	router.HandleFunc("/Parametros/{pam}", Parametros)

	server := http.ListenAndServe(":8080", router)

	log.Fatal(server)
}

func Index(w http.ResponseWriter, r *http.Request) {
	modelTest := ModelTests{
		ModelTest{"Nombre 1", 1232, "Director 1"},
		ModelTest{"Nombre 2", 345, "Director 2"},
		ModelTest{"Nombre 3", 1232, "Director 3"},
	}

	json.NewEncoder(w).Encode(modelTest)
}

func Contacto(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Página contacto desde el servidor web")
}

func Parametros(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	param := params["pam"]

	fmt.Fprintln(w, "Se ha ingresado el parámetro: "+param)
}
