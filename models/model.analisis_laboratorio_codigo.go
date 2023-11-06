package models

// Model

type AnalisisLaboratorioCodigo struct {
	Id     int    `json:"id"`
	Nombre string `json:"nombre"`
}

//Constructor

func NewAnalisisLaboratorio(id int, nombre string) *AnalisisLaboratorioCodigo {
	return &AnalisisLaboratorioCodigo{
		Id:     id,
		Nombre: nombre,
	}
}

func UpdateAnalisisLaboratorio(id int, nombre string) *AnalisisLaboratorioCodigo {
	return &AnalisisLaboratorioCodigo{
		Id:     id,
		Nombre: nombre}
}
