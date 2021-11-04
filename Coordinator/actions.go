package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
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

	transacciones = nil
	transacciones = append(transacciones, transaccionData) //Guardado en arreglo para usarse posteriormente

	/*dataPost := url.Values{
		"Nombre":        {transacciones[0].Nombre},
		"CuentaOrigen":  {strconv.Itoa(transacciones[0].CuentaOrigen)},
		"CuentaDestino": {strconv.Itoa(transacciones[0].CuentaDestino)},
		"Monto":         {strconv.Itoa(transacciones[0].Monto)},
	}*/

	client := &http.Client{}
	time.Sleep(1 * time.Millisecond)
	//req, err := http.PostForm("http://localhost:8081/RegistrarTransaccion", dataPost)
	bodyJson, _ := json.Marshal(transaccionData)
	req, err := http.NewRequest("POST", "http://127.0.0.1:4444/RegistrarTransaccion", bytes.NewReader(bodyJson))
	req.Close = true

	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	req.Body.Close()

	if err != nil {
		fmt.Println("error: ", err)
	} else {

		//var res map[string]interface{}

		cuerpoRespuesta, err := ioutil.ReadAll(resp.Body)
		if err == nil {

		}
		respuestaString := string(cuerpoRespuesta)
		//log.Printf("Código de respuesta: %d", resp.StatusCode)
		//log.Printf("Encabezados: '%q'", resp.Header)
		//contentType := resp.Header.Get("Content-Type")
		//log.Printf("El tipo de contenido: '%s'", contentType)
		//log.Printf("Cuerpo de respuesta del servidor: '%s'", respuestaString)
		fmt.Fprintln(w, respuestaString)
	}

	//Validar si el bloque esta lleno (3 transacciones por bloque)(Blockchain)
	// Si esta lleno, enviar peticion a openCloser
	// Si esta lleno, despues de cerrar el bloque, se envia a Blockchain petición para abrir el bloque

	//Enviar a controlador de blockchain RegistrarTransaccion
}

func ConsultarFondos(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, transacciones)

	//Procedimiento para consultar fondos desde el api de Blockchain
}
