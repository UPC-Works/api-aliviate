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
	Status int                      `json:"status,omitempty"`
	Error  ErrorStructure           `json:"error"`
	Data   []models.HistoriaClinica `json:"data"`
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
