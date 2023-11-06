package models

// Model

type AnalisisHistoria struct {
	IdHistoriaClinica string  `json:"id"`
	IdAnalisisCampo   int     `json:"idAnalisis"`
	Valor             float32 `json:"valor"`
}

//Constructor

func NewAnalisisHistoria(idHistoriaClinica string, idAnalisisCampo int, valor float32) *AnalisisHistoria {
	return &AnalisisHistoria{
		IdHistoriaClinica: idHistoriaClinica,
		IdAnalisisCampo:   idAnalisisCampo,
		Valor:             valor,
	}
}

func UpdateAnalisisHistoria(idHistoriaClinica string, idAnalisisCampo int, valor float32) *AnalisisHistoria {
	return &AnalisisHistoria{
		IdHistoriaClinica: idHistoriaClinica,
		IdAnalisisCampo:   idAnalisisCampo,
		Valor:             valor,
	}
}
