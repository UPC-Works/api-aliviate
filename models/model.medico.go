package models

import "time"

// Model

type Medico struct {
	Id                 string    `json:"id"`
	Nombre             string    `json:"nombre"`
	Apellido           string    `json:"apellido"`
	Colegiatura        string    `json:"colegiatura"`
	DocumentoIdentidad int       `json:"documentoIdentidad"`
	Correo             string    `json:"correo"`
	Contrasenia        string    `json:"contrasenia"`
	FechaRegistro      time.Time `json:"fechaRegistro"`
	Direccion          string    `json:"direccion"`
	Especialidad       string    `json:"especialidad"`
}

//Constructor

func NewMedico(id string, nombre string, apellido string, colegiatura string, documentoIdentidad int, correo string, contrasenia string, direccion string, especialidad string) *Medico {
	return &Medico{
		Id:                 id,
		Nombre:             nombre,
		Apellido:           apellido,
		Colegiatura:        colegiatura,
		DocumentoIdentidad: documentoIdentidad,
		Correo:             correo,
		Contrasenia:        contrasenia,
		FechaRegistro:      time.Now(),
		Direccion:          direccion,
		Especialidad:       especialidad,
	}
}
