package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	go miNombreLento("Federico Rosales")
	fmt.Println("Estoy aqui")
	var x string
	fmt.Scanln(&x)

}

func miNombreLento(nombre string) {
	letra := strings.Split(nombre, "")
	for _, letra := range letra {
		time.Sleep(1000 * time.Millisecond)
		fmt.Println(letra)
	}
}
