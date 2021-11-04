package main

type Transaccion struct {
	Documento     int
	Nombre        string
	CuentaOrigen  int
	CuentaDestino int
	Monto         int
}

type Transacciones []Transaccion
