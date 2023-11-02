package models

// Model

type PrediccionIA struct {
	Id           string  `json:"id"`
	IdConsulta   string  `json:"idConsulta"`
	Enfermedad   string  `json:"enfermedad"`
	Probabilidad float32 `json:"probabilidad"`
	EsAceptado   float32 `json:"esAceptado"`
}
