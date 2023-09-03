package models

// Model

type Establecimiento struct {
	Id         string `json:"id"`
	IdDistrito int    `json:"idDistrito"`
	Nombre     string `json:"nombre"`
	Dirección  string `json:"dirección"`
}

//Constructor

func NewEstablecimiento(id string, IdDistrito int, nombre string, direccion string) *Establecimiento {
	return &Establecimiento{
		Id:         id,
		IdDistrito: IdDistrito,
		Nombre:     nombre,
		Dirección:  direccion,
	}
}

func UpdateEstablecimiento(id string, IdDistrito int, nombre string, direccion string) *Establecimiento {
	return &Establecimiento{
		Id:         id,
		IdDistrito: IdDistrito,
		Nombre:     nombre,
		Dirección:  direccion,
	}
}
