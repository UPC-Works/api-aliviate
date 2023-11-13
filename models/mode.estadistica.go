package models

// Model

type EnfermedadPrediccion struct {
	Enfermedad         string  `json:"enfermedad"`
	PrediccionCorrecta float32 `json:"prediccionesCorrectas"`
}
