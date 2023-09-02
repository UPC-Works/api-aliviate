package models

// Model

type Establecimiento struct {
	Id     string `json:"id"`
	Nombre string `json:"nombre"`
}

//Constructor

func NewEstablecimiento(id string, nombre string) *Establecimiento {
	return &Establecimiento{
		Id:     id,
		Nombre: nombre,
	}
}

func UpdateEstablecimiento(id string, nombre string) *Establecimiento {
	return &Establecimiento{
		Id:     id,
		Nombre: nombre,
	}
}
