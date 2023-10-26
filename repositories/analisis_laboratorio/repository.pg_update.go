package analisis_laboratorio

import (
	"context"
	"time"

	configs "github.com/UPC-Works/api-aliviate/configs"
	models "github.com/UPC-Works/api-aliviate/models"
)

func Pg_UpdateOne(input_analisis_laboratorio *models.AnalisisLaboratorio) error {

	//Context time limit
	ctx, cancel := context.WithTimeout(context.Background(), 8*time.Second)
	defer cancel()

	db := configs.Conn_Pg_DB()

	query := `UPDATE AnalisisLaboratorio SET
	colesterol=$1,
	trigliceridos=$2,
	colesterol_hdl=$3,
	colesterol_ldl=$4,
	colesterol_vldl=$5,
	riesgo1=$6,
	riesgo2=$7,
	glucosa=$8,
	hematrocito=$9,
	hemoglobina=$10
	WHERE id=$11`
	_, err_query := db.Exec(ctx, query, input_analisis_laboratorio.Colesterol, input_analisis_laboratorio.Trigliceridos, input_analisis_laboratorio.ColesterolHdl, input_analisis_laboratorio.ColesterolLdl, input_analisis_laboratorio.ColesterolVldl, input_analisis_laboratorio.Riesgo1, input_analisis_laboratorio.Riesgo2, input_analisis_laboratorio.Glucosa, input_analisis_laboratorio.Hematrocito, input_analisis_laboratorio.Hemoglobina, input_analisis_laboratorio.Id)

	if err_query != nil {
		return err_query
	}

	return nil
}
