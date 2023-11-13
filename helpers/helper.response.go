package helper

import (
	models "github.com/UPC-Works/api-aliviate/models"
)

type ResponseString struct {
	Status int            `json:"status,omitempty"`
	Error  ErrorStructure `json:"error"`
	Data   string         `json:"data"`
}

type ResponseJwt struct {
	Status int            `json:"status,omitempty"`
	Error  ErrorStructure `json:"error"`
	Data   JwtStructure   `json:"data"`
}

type ResponseAuth struct {
	Error ErrorStructure `json:"error"`
	Data  AuthStructure  `json:"data"`
}

type ResponseListMedico struct {
	Status int             `json:"status,omitempty"`
	Error  ErrorStructure  `json:"error"`
	Data   []models.Medico `json:"data"`
}

type ResponseListHistoriaClinica struct {
	Status int                           `json:"status,omitempty"`
	Error  ErrorStructure                `json:"error"`
	Data   []models.HistoriaClinica_View `json:"data"`
}

type ResponseListEstablecimiento struct {
	Status int                      `json:"status,omitempty"`
	Error  ErrorStructure           `json:"error"`
	Data   []models.Establecimiento `json:"data"`
}

type ResponseListPaciente struct {
	Status int               `json:"status,omitempty"`
	Error  ErrorStructure    `json:"error"`
	Data   []models.Paciente `json:"data"`
}

type ResponseListConsulta struct {
	Status int               `json:"status,omitempty"`
	Error  ErrorStructure    `json:"error"`
	Data   []models.Consulta `json:"data"`
}

type ResponseHistoriaClinica struct {
	Status int                    `json:"status,omitempty"`
	Error  ErrorStructure         `json:"error"`
	Data   models.HistoriaClinica `json:"data"`
}

type ResponsePrediccion struct {
	Status int                           `json:"status,omitempty"`
	Error  ErrorStructure                `json:"error"`
	Data   []models.PrediccionEnfermedad `json:"data"`
}

type ResponseListAnalisisLaboratorioCodigo struct {
	Status int                                `json:"status,omitempty"`
	Error  ErrorStructure                     `json:"error"`
	Data   []models.AnalisisLaboratorioCodigo `json:"data"`
}

type ResponseListAnalisisLaboratorioCampo struct {
	Status int                               `json:"status,omitempty"`
	Error  ErrorStructure                    `json:"error"`
	Data   []models.AnalisisLaboratorioCampo `json:"data"`
}

type ResponseListAnalisisHistoria struct {
	Status int                       `json:"status,omitempty"`
	Error  ErrorStructure            `json:"error"`
	Data   []models.AnalisisHistoria `json:"data"`
}

type ResponseListDocumentos struct {
	Status int                         `json:"status,omitempty"`
	Error  ErrorStructure              `json:"error"`
	Data   []models.DocumentosHistoria `json:"data"`
}

type ResponseListHistorialCambios struct {
	Status int                              `json:"status,omitempty"`
	Error  ErrorStructure                   `json:"error"`
	Data   []models.ModificacionesHistorias `json:"data"`
}

type ResponseListEnfermedadPrediccion struct {
	Status int                           `json:"status,omitempty"`
	Error  ErrorStructure                `json:"error"`
	Data   []models.EnfermedadPrediccion `json:"data"`
}
