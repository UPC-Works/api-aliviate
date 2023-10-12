package models

// Model

type PrediccionEnfermedad struct {
	Enfermedad   string  `json:"enfermedad"`
	Probabilidad float32 `json:"probabilidad"`
}
