package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"gopkg.in/mgo.v2"
	//"gopkg.in/mgo.v2/bson"
)

func getSession() *mgo.Session {
	session, err := mgo.Dial("mongodb://localhost")

	if err != nil {
		panic(err)
	} else {
		return session
	}
}

var transacciones = Transacciones{}
var collection = getSession().DB("blockchain").C("transacciones")

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Bienvenido al Coordinador")
}

func RegistrarTransaccion(w http.ResponseWriter, r *http.Request) {

	decoder := json.NewDecoder(r.Body)
	var transaccionData Transaccion

	err := decoder.Decode(&transaccionData)
	if err != nil {
		panic(err)
	}

	defer r.Body.Close()
	impresion, err := json.Marshal(transaccionData)
	log.Println(string(impresion))
	fmt.Fprintln(w, string(impresion))

	//Guardar transaccion en la base de datos
	err = collection.Insert(transaccionData)

	if err != nil {
		log.Println("Error al intentar insertar el registro: " + err.Error())
		w.WriteHeader(500)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	//w.WriteHeader(200)
	json.NewEncoder(w).Encode((transaccionData))
}
