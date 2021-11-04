package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
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
var transaccionesTemp = Transacciones{}
var bloqueTemp = Bloques{}
var collection = getSession().DB("blockchain").C("transacciones")
var collectionBlock = getSession().DB("blockchain").C("block")

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Bienvenido al Blockchain")
}

func ConsultarTransaccionId(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	transaccion_id := params["id"]

	if !bson.IsObjectIdHex(transaccion_id) {
		w.WriteHeader(405)
		return
	}

	oid := bson.ObjectIdHex(transaccion_id)

	err := collection.FindId(oid).All(&transacciones)

	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Resultados: ", transacciones)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(transacciones)
}

func ConsultarTransaccion(w http.ResponseWriter, r *http.Request) {
	transacciones = nil
	params := mux.Vars(r)
	transaccion_id := params["id"]
	transaccion_idInt, err2 := strconv.Atoi(transaccion_id)

	err := collection.Find(nil).All(&transaccionesTemp)

	if err != nil || err2 != nil {
		log.Fatal(err)
	}

	for i := 0; i < len(transaccionesTemp); i++ {
		if transaccionesTemp[i].Documento == transaccion_idInt {
			transacciones = append(transacciones, transaccionesTemp[i])
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(transacciones)
}

func CrearBloque(w http.ResponseWriter, r *http.Request) {

	err := collectionBlock.Find(nil).All(&bloqueTemp)
	consecutivo := 0
	HashPrev := ""
	var bloqueData Bloque

	if err != nil {
		log.Fatal(err)
	}

	if bloqueTemp != nil {
		for i := 0; i < len(bloqueTemp); i++ {

			if bloqueTemp[i].Consecutivo > consecutivo {
				consecutivo = bloqueTemp[i].Consecutivo
				HashPrev = bloqueTemp[i].Hash
			}
		}
	}

	if HashPrev == "" {
		HashPrev = "000000000000"
	}

	bloqueData.Consecutivo = consecutivo + 1
	bloqueData.Estado = true
	bloqueData.Nonce = "1" //Pendiente número aleatorio
	bloqueData.Datos = nil
	bloqueData.HashPrev = HashPrev
	bloqueData.Hash = ""

	//Enviar objeto de bloque a guardar
	defer r.Body.Close()
	dataEnviada, err := json.Marshal(bloqueData)
	log.Println(string(dataEnviada))

	//Guardar transaccion en la base de datos
	err = collectionBlock.Insert(bloqueData)

	if err != nil {
		log.Println("Error al intentar insertar el registro: " + err.Error())
		w.WriteHeader(500)
		return
	}

	w.WriteHeader(200)
}

func RegistrarTransaccion(w http.ResponseWriter, r *http.Request) {
	err := collectionBlock.Find(nil).All(&bloqueTemp)
	consecutivo := 0

	var bloqueData Bloque

	if err != nil {
		log.Fatal(err)
	}

	if bloqueTemp != nil {
		for i := 0; i < len(bloqueTemp); i++ {
			if bloqueTemp[i].Estado == true {
				consecutivo = bloqueTemp[i].Consecutivo
				bloqueData = bloqueTemp[i]
			}
		}
	}

	//Proceso para crear y guardar la nueva transacción
	decoder := json.NewDecoder(r.Body)
	var transaccionData Transaccion
	err = decoder.Decode(&transaccionData)
	if err != nil {
		panic(err)
	}

	defer r.Body.Close()
	impresion, err := json.Marshal(transaccionData)
	log.Println(string(impresion))
	fmt.Fprintln(w, string(impresion))

	//Proceso que actualiza el arreglo de datos en el bloque
	bloqueData.Datos = append(bloqueData.Datos, transaccionData)

	//PROCESO QUE ACTULIZA EL BLOQUE (No modificar)
	document := bson.M{"consecutivo": consecutivo}
	change := bson.M{"$set": bloqueData}
	err = collectionBlock.Update(document, change)

	if err != nil {
		log.Fatal(err)
		w.WriteHeader(500)
		return
	} else {
		log.Println(document)
		w.WriteHeader(200)
	}
}
