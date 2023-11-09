package models

// Model

type AnalisisLaboratorioCampo struct {
	Id         int    `json:"id"`
	IdAnalisis int    `json:"idAnalisis"`
	Campo      string `json:"campo"`
	CampoJson  string `json:"campoJson"`
	Tipo       string `json:"tipo"`
	Requerido  bool   `json:"requerido"`
}
