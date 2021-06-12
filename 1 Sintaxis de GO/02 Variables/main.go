package main

import (
	"fmt"
	"strconv"
)

var numero int
var text string
var status bool = true

func main() {
	var numero2, numero3, numero4 int
	numero2, numero3, numero4, texto := 4, 5, 98, "Este es mi texto" // gran avance de declaracion de variables

	var moneda int64 = 17
	numero2 = int(moneda)

	// texto = fmt.Sprintf("%d", moneda) una forma de convertir

	texto = strconv.Itoa(int(moneda))

	fmt.Println(numero2)
	fmt.Println(numero3)
	fmt.Println(numero4)
	fmt.Println(texto)
	fmt.Println(status)

	MostrarStatus()
}

func MostrarStatus() {
	fmt.Println(status)
}
