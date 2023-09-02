package models

import "time"

// Model

type Medico struct {
	Id                 string    `json:"id"`
	IdEstablecimiento  string    `json:"idEstablecimiento"`
	Nombre             string    `json:"nombre"`
	Apellido           string    `json:"apellido"`
	Colegiatura        string    `json:"colegiatura"`
	DocumentoIdentidad int       `json:"documentoIdentidad"`
	Correo             string    `json:"correo"`
	Contrasenia        string    `json:"contrasenia"`
	FechaRegistro      time.Time `json:"fechaRegistro"`
	Direccion          string    `json:"direccion"`
}

//Constructor

func NewMedico(id string, idEstablecimiento string, nombre string, apellido string, colegiatura string, documentoIdentidad int, correo string, contrasenia string, direccion string) *Medico {
	return &Medico{
		Id:                 id,
		IdEstablecimiento:  idEstablecimiento,
		Nombre:             nombre,
		Apellido:           apellido,
		Colegiatura:        colegiatura,
		DocumentoIdentidad: documentoIdentidad,
		Correo:             correo,
		Contrasenia:        contrasenia,
		FechaRegistro:      time.Now(),
		Direccion:          direccion,
	}
}
