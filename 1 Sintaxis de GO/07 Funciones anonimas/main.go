package main

import (
	"fmt"
)

var Calculo func(int, int) int = func(num1 int, num2 int) int {
	return num1 + num2
}

func main() {
	fmt.Printf("Sumo 5 + 2 = %d \n", Calculo(5, 2))

	// restamos
	Calculo = func(num1 int, num2 int) int {
		return num1 - num2
	}
	fmt.Printf("Resta 6 - 4 = %d \n", Calculo(6, 4))

	// dividimos
	Calculo = func(num1 int, num2 int) int {
		return num1 / num2
	}
	fmt.Printf("Dividimos 12 - 3 = %d \n", Calculo(13, 3))

	Operaciones()
	/* CLOSURES */ // algo muy usado y para tener en cuenta
	Tabladel := 2
	Mitabla := Tabla(Tabladel)
	for i := 1; i < 11; i++ {
		fmt.Println(Mitabla())
	}
}

func Operaciones() {
	resultado := func() int {
		var a, b int = 23, 13
		return a + b
	}
	fmt.Println(resultado())
}

func Tabla(valor int) func() int {
	numero := valor
	secuencia := 0
	return func() int {
		secuencia += 1
		return numero * secuencia
	}
}
