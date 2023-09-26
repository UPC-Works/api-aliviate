package analisis_laboratorio

import (
	"context"
	"time"

	configs "github.com/UPC-Works/api-aliviate/configs"
	models "github.com/UPC-Works/api-aliviate/models"
)

func Pg_Create(input_analisis_laboratorio *models.AnalisisLaboratorio) error {

	//Context time limit
	ctx, cancel := context.WithTimeout(context.Background(), 8*time.Second)
	defer cancel()

	db := configs.Conn_Pg_DB()

	query := `INSERT INTO AnalisisLaboratorio (
		id,
		id_historia_clinica,
		colesterol,
		trigliceridos,
		colesterol_hdl,
		colesterol_ldl,
		colesterol_vldl,
		riesgo1,
		riesgo2,
		glucosa,
		hematrocito,
		hemoglobina
	) VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12)`
	_, err_query := db.Exec(ctx, query,
		input_analisis_laboratorio.Id,
		input_analisis_laboratorio.IdHistoriaClinica,
		input_analisis_laboratorio.Colesterol,
		input_analisis_laboratorio.Trigliceridos,
		input_analisis_laboratorio.ColesterolHdl,
		input_analisis_laboratorio.ColesterolLdl,
		input_analisis_laboratorio.ColesterolVldl,
		input_analisis_laboratorio.Riesgo1,
		input_analisis_laboratorio.Riesgo2,
		input_analisis_laboratorio.Glucosa,
		input_analisis_laboratorio.Hematrocito,
		input_analisis_laboratorio.Hemoglobina,
	)

	if err_query != nil {
		return err_query
	}

	return nil
}
