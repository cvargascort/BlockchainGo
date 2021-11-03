package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

var transacciones = Transacciones{}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Bienvenido al Coordinador")
}

func RegistrarTransaccion(w http.ResponseWriter, r *http.Request) {

	decoder := json.NewDecoder(r.Body)
	var transaccionData Transaccion

	log.Println("Llego por aqui")
	err := decoder.Decode(&transaccionData)
	if err != nil {
		panic(err)
	}

	defer r.Body.Close()
	impresion, err := json.Marshal(transaccionData)
	log.Println(string(impresion))
	fmt.Fprintln(w, string(impresion))

	//Guardar transaccion en la base de datos
}
