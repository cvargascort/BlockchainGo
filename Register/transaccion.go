package main

type Transaccion struct {
	Nombre        string
	CuentaOrigen  int
	CuentaDestino int
	Monto         int
}

type Transacciones []Transaccion
