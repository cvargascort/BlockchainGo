package main

import (
	"fmt"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Bienvenido al OpenCloser")
}

func CerrarBloque(w http.ResponseWriter, r *http.Request) {

	fmt.Println("Bloque cerrado correctamente")
}
