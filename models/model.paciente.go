package models

import "time"

// Model

type Paciente struct {
	Id                 string    `json:"id"`
	Nombre             string    `json:"nombre"`
	Apellido           string    `json:"apellido"`
	FechaNacimiento    time.Time `json:"fechaNacimiento"`
	Genero             int       `json:"genero"`
	DocumentoIdentidad int       `json:"documentoIdentidad"`
	FechaRegistro      time.Time `json:"fechaRegistro"`
	GrupoSanguineo     string    `json:"grupoSanguineo"`
	RhSanguineo        string    `json:"rhSanguineo"`
	Telefono           string    `json:"telefono"`
}

//Constructor

func NewPaciente(id string, nombre string, apellido string, fechaNacimiento time.Time, genero int, documentoIdentidad int, grupoSanguineo string, rhSanguineo string, telefono string) *Paciente {
	return &Paciente{
		Id:                 id,
		Nombre:             nombre,
		Apellido:           apellido,
		FechaNacimiento:    fechaNacimiento,
		Genero:             genero,
		DocumentoIdentidad: documentoIdentidad,
		FechaRegistro:      time.Now(),
		RhSanguineo:        rhSanguineo,
		GrupoSanguineo:     grupoSanguineo,
		Telefono:           telefono,
	}
}

func UpdatePaciente(id string, nombre string, apellido string, fechaNacimiento time.Time, genero int, documentoIdentidad int, grupoSanguineo string, rhSanguineo string, telefono string) *Paciente {
	return &Paciente{
		Id:                 id,
		Nombre:             nombre,
		Apellido:           apellido,
		FechaNacimiento:    fechaNacimiento,
		Genero:             genero,
		DocumentoIdentidad: documentoIdentidad,
		FechaRegistro:      time.Now(),
		RhSanguineo:        rhSanguineo,
		GrupoSanguineo:     grupoSanguineo,
		Telefono:           telefono,
	}
}
