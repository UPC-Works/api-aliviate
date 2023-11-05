package models

// Model

type DocumentosHistoria struct {
	Id                string `json:"id"`
	IdHistoriaClinica string `json:"idHistoriaClinica"`
	Url               string `json:"url"`
}

//Constructor

func NewDocumentosHistoria(id string, idHistoriaClinica string, url string) *DocumentosHistoria {
	return &DocumentosHistoria{
		Id:                id,
		IdHistoriaClinica: idHistoriaClinica,
		Url:               url,
	}
}

func UpdateDocumentosHistoria(id string, idHistoriaClinica string, url string) *DocumentosHistoria {
	return &DocumentosHistoria{
		Id:                id,
		IdHistoriaClinica: idHistoriaClinica,
		Url:               url,
	}
}
