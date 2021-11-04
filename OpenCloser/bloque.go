package main

type Bloque struct {
	Consecutivo int
	Estado      bool
	Nonce       string
	Datos       Transacciones
	HashPrev    string
	Hash        string
}

type Bloques []Bloque
