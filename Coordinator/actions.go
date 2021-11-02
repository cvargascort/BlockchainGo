package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strconv"
)

var transacciones = Transacciones{}

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

	transacciones = append(transacciones, transaccionData) //Guardado en arreglo innecesario

	dataPost := url.Values{
		"Nombre":        {transaccionData.Nombre},
		"CuentaOrigen":  {strconv.Itoa(transaccionData.CuentaOrigen)},
		"CuentaDestino": {strconv.Itoa(transaccionData.CuentaDestino)},
		"Monto":         {strconv.Itoa(transaccionData.Monto)},
	}

	//Enviar a controlador de log
	resp, err := http.PostForm("http://localhost:8081/RegistrarTransaccion", dataPost)

	log.Println("************* Por aqui va la vaina *************************")

	if err != nil {
		fmt.Println("error: ", err)
	} else {
		var res map[string]interface{}

		json.NewDecoder(resp.Body).Decode(&res)

		fmt.Println(res["form"])
	}
	//Enviar a controlador de blockchain
}

func ConsultarFondos(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, transacciones)

	//Procedimiento para consultar fondos desde el api de Blockchain
}
