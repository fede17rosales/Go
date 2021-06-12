package main

import (
	"fmt"
)

func main() {
	fmt.Println(Calculo(5, 46))
	fmt.Println(Calculo(1, 34, 56, 89))
	fmt.Println(Calculo(2, 123, 25, 231))
	fmt.Println(Calculo(5, 5, 6, 7, 8, 9, 10))

	//fmt.Println(uno(5))
	//numero, estado := dos(1)
	//fmt.Println(numero)
	//fmt.Println(estado)

}

func uno(numero int) (z int) {
	z = numero * 2
	return
}

func dos(numero int) (int, bool) {
	if numero == 1 {
		return 5, false
	} else {
		return 10, true
	}
}

func Calculo(numero ...int) int {
	total := 0
	for item, num := range numero { // range devuelve dos elementos, el primero devuelve devuelve el numero de elemento el segundo es el valor del elemento
		total = total + num
		fmt.Printf("\n Item %d", item)
	}
	return total
}
