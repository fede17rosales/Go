package user

import "time"

// crear una estructura: tipo + nombrevariable + tipo de dato
type Usuario struct {
	Id        int
	Nombre    string
	FechaAlta time.Time
	Status    bool
}

func (this *Usuario) AltaUsuario(id int, nombre string, fechaalta time.Time, status bool) {
	this.Id = id
	this.Nombre = nombre
	this.FechaAlta = fechaalta
	this.Status = status
}
