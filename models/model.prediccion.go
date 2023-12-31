package models

// Model

type PrediccionEnfermedad struct {
	Enfermedad   string  `json:"enfermedad"`
	Probabilidad float32 `json:"probabilidad"`
	EsAceptado   bool    `json:"esAceptado"`
}

type PrediccionEnfermedadShow struct {
	Enfermedad   string  `json:"enfermedad"`
	Probabilidad float32 `json:"probabilidad"`
}

type PrediccionTratamiento struct {
	Tratamiento  string  `json:"tratamiento"`
	Probabilidad float32 `json:"probabilidad"`
}
