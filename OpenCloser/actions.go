package main

import (
	"fmt"
	"log"
	"net/http"

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

var bloqueTemp = Bloques{}
var collectionBlock = getSession().DB("blockchain").C("block")

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Bienvenido al OpenCloser")
}

func CerrarBloque(w http.ResponseWriter, r *http.Request) {

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

	bloqueData.Estado = false
	//bloqueData.Nonce = NewNonce // Sería el nuevo nonce que compata con el nuevo hash
	//bloqueData.Hash = NewHash con elementos del bloque (Hash debe tener los primeros 4 valores en cero)

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
