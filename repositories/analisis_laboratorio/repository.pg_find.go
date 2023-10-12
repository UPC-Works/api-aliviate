package analisis_laboratorio

import (
	"context"
	"fmt"
	"strings"
	"time"

	configs "github.com/UPC-Works/api-aliviate/configs"
	models "github.com/UPC-Works/api-aliviate/models"
)

func Pg_FindOne(input_idhistoriaclinica string) (models.AnalisisLaboratorio, error) {

	//Initialization
	var oAnalisisLaboratorio models.AnalisisLaboratorio

	//Define the filters
	filters := map[string]interface{}{}
	counter_filters := 0
	if input_idhistoriaclinica != "" {
		filters["id_historia_clinica"] = input_idhistoriaclinica
		counter_filters += 1
	}

	//Context timing
	ctx, cancel := context.WithTimeout(context.Background(), 8*time.Second)
	//Cancel context
	defer cancel()

	//Start the connection
	db := configs.Conn_Pg_DB()

	//Define the query
	q := `SELECT 
	COALESCE(id,''),
	COALESCE(id_historia_clinica,''),
	COALESCE(colesterol,0),
	COALESCE(trigliceridos,0),
	COALESCE(colesterol_hdl,0),
	COALESCE(colesterol_ldl,0),
	COALESCE(colesterol_vldl,0),
	COALESCE(riesgo1,0),
	COALESCE(riesgo2,0),
	COALESCE(glucosa,0),
	COALESCE(hematrocito,0),
	COALESCE(hemoglobina,0)
FROM AnalisisLaboratorio `
	if counter_filters > 0 {
		q += " WHERE "
		clausulas := make([]string, 0)
		for key, value := range filters {
			clausulas = append(clausulas, fmt.Sprintf("%s = '%s'", key, value))
		}
		q += strings.Join(clausulas, " AND ")

	}

	error_find := db.QueryRow(ctx, q).Scan(
		&oAnalisisLaboratorio.Id,
		&oAnalisisLaboratorio.IdHistoriaClinica,
		&oAnalisisLaboratorio.Colesterol,
		&oAnalisisLaboratorio.Trigliceridos,
		&oAnalisisLaboratorio.ColesterolHdl,
		&oAnalisisLaboratorio.ColesterolLdl,
		&oAnalisisLaboratorio.ColesterolVldl,
		&oAnalisisLaboratorio.Riesgo1,
		&oAnalisisLaboratorio.Riesgo2,
		&oAnalisisLaboratorio.Glucosa,
		&oAnalisisLaboratorio.Hematrocito,
		&oAnalisisLaboratorio.Hemoglobina,
	)

	if error_find != nil {
		return oAnalisisLaboratorio, error_find
	}

	//Return one analisis_laboratorio
	return oAnalisisLaboratorio, nil
}
