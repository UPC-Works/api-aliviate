package models

import "time"

// Model

type ModificacionesHistorias struct {
	Id                string    `json:"id"`
	IdHistoriaClinica string    `json:"idHistoriaClinica"`
	NombreMedico      string    `json:"nombreMedico"`
	ActualizadoEl     time.Time `json:"actualizadoEl"`
}

//Constructor

func NewModificacionesHistorias(id string, idHistoriaClinica string, nombreMedico string) *ModificacionesHistorias {
	return &ModificacionesHistorias{
		Id:                id,
		IdHistoriaClinica: idHistoriaClinica,
		NombreMedico:      nombreMedico,
		ActualizadoEl:     time.Now(),
	}
}

func UpdateModificacionesHistorias(id string, idHistoriaClinica string, nombreMedico string, actualizado_el time.Time) *ModificacionesHistorias {
	return &ModificacionesHistorias{
		Id:                id,
		IdHistoriaClinica: idHistoriaClinica,
		NombreMedico:      nombreMedico,
		ActualizadoEl:     actualizado_el,
	}
}
