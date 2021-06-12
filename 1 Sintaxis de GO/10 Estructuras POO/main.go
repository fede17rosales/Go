package main

import (
	"fmt"
	"time"

	us "./user"
)

// Herencia
type pepe struct {
	us.Usuario
}

func main() {
	// creando un objecto: un usuario pepe
	u := new(pepe)
	u.AltaUsuario(1, "Pepe Argento", time.Now(), true)
	fmt.Println(u.Usuario)

}
