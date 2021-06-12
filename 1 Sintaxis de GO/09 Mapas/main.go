package main

import (
	"fmt"
)

func main() {
	paises := make(map[string]string, 5) // 5: me va reservar el espacio para 5 paises, si paso ese limite go puede agregarlos dinamicamente
	fmt.Println(paises)

	paises["Mexico"] = "D.F"
	paises["Argentina"] = "Buenos Aires"

	fmt.Println(paises["Mexico"])

	campeonato := map[string]int{
		"Barcelona":    39,
		"Real Madrid":  38,
		"Chelsea":      35,
		"Boca Juniors": 30}
	fmt.Println(campeonato)

	campeonato["River Plate"] = 25
	fmt.Println(campeonato)
	campeonato["Chelsea"] = 44
	fmt.Println(campeonato)
	delete(campeonato, "Real Madrid")
	fmt.Println(campeonato)

	for equipo, puntaje := range campeonato {
		fmt.Printf("El equipo %s,tiene un puntaje de: %d\n", equipo, puntaje)
	}
	puntaje, existe := campeonato["Boca Juniors"]
	fmt.Printf("El puntaje capturado %d, y el equipo existe %t", puntaje, existe)
}
