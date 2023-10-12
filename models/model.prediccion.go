package models

// Model

type Prediccion struct {
	IdHistoriaClinica string   `json:"idHistoriaClinica"`
	ConsultaActual    Consulta `json:"consultaActual"`
}
