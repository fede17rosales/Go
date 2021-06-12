package main

import (
	"fmt"
)

var matriz [5][7]int // matriz

func main() {
	tabla := [5]int{15, 894, 9, 54, 8}
	for i := 0; i < len(tabla); i++ {
		fmt.Println(tabla[i])
	}

	matriz[3][5] = 1
	fmt.Println(matriz)

	matriz := []int{1, 3, 4} //slice
	fmt.Println(matriz)
	variante2()
	variante3()
	variante4()
}

func variante2() {
	elementos := [5]int{1, 2, 3, 4, 5}
	porcion := elementos[2:4]
	fmt.Println(porcion)
}

func variante3() {
	elementos := make([]int, 5, 20) // me va crear un vector de 5 de hasta 20 elementos
	fmt.Printf("Largo %d, capacidad %d", len(elementos), cap(elementos))
}

func variante4() {
	nums := make([]int, 0, 0)
	for i := 0; i < 100; i++ {
		nums = append(nums, i)
	}
	fmt.Printf("\nLargo %d, capacidad %d", len(nums), cap(nums))
}
