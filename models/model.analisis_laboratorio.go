package models

// Model

type AnalisisLaboratorio struct {
	Id                string  `json:"id "`
	IdHistoriaClinica string  `json:"idHistoriaClinica"`
	Colesterol        float32 `json:"colesterol"`
	Trigliceridos     float32 `json:"trigliceridos"`
	ColesterolHdl     float32 `json:"colesterolHdl"`
	ColesterolLdl     float32 `json:"colesterolLdl"`
	ColesterolVldl    float32 `json:"colesterolVldl"`
	Riesgo1           float32 `json:"riesgo1"`
	Riesgo2           float32 `json:"riesgo2"`
	Glucosa           float32 `json:"glucosa"`
	Hematrocito       float32 `json:"hematrocito"`
	Hemoglobina       float32 `json:"hemoglobina"`
}

//Constructor

func NewAnalisisLaboratorio(id string, idHistoriaClinica string, colesterol float32, trigliceridos float32, colesterolHdl float32, colesterolLdl float32, colesterolVldl float32, riesgo1 float32, riesgo2 float32, glucosa float32, hematrocito float32, hemoglobina float32) *AnalisisLaboratorio {
	return &AnalisisLaboratorio{
		Id:                id,
		IdHistoriaClinica: idHistoriaClinica,
		Colesterol:        colesterol,
		Trigliceridos:     trigliceridos,
		ColesterolHdl:     colesterolHdl,
		ColesterolLdl:     colesterolLdl,
		ColesterolVldl:    colesterolVldl,
		Riesgo1:           riesgo1,
		Riesgo2:           riesgo2,
		Glucosa:           glucosa,
		Hematrocito:       hematrocito,
		Hemoglobina:       hemoglobina}
}

func UpdateAnalisisLaboratorio(id string, idHistoriaClinica string, colesterol float32, trigliceridos float32, colesterolHdl float32, colesterolLdl float32, colesterolVldl float32, riesgo1 float32, riesgo2 float32, glucosa float32, hematrocito float32, hemoglobina float32) *AnalisisLaboratorio {
	return &AnalisisLaboratorio{
		Id:                id,
		IdHistoriaClinica: idHistoriaClinica,
		Colesterol:        colesterol,
		Trigliceridos:     trigliceridos,
		ColesterolHdl:     colesterolHdl,
		ColesterolLdl:     colesterolLdl,
		ColesterolVldl:    colesterolVldl,
		Riesgo1:           riesgo1,
		Riesgo2:           riesgo2,
		Glucosa:           glucosa,
		Hematrocito:       hematrocito,
		Hemoglobina:       hemoglobina}
}
