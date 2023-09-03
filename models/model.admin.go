package models

// Model

type Admin struct {
	Id             string `json:"id"`
	NombreCompleto string `json:"nombreCompleto"`
	Correo         string `json:"correo"`
	Contrasenia    string `json:"contrasenia"`
}

//Constructor

func NewAdmin(id string, nombreCompleto string, correo string, contrasenia string) *Admin {
	return &Admin{
		Id:             id,
		NombreCompleto: nombreCompleto,
		Correo:         correo,
		Contrasenia:    contrasenia,
	}
}
