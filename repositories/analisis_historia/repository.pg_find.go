package analisis_historia

import (
	"context"
	"fmt"
	"strings"
	"time"

	configs "github.com/UPC-Works/api-aliviate/configs"
	models "github.com/UPC-Works/api-aliviate/models"
)

func Pg_FindMultiple(input_id_historia string, input_id_analisis int) ([]models.AnalisisHistoria, error) {

	//Initialization
	var oListAnalisisHistoria []models.AnalisisHistoria

	//Define the filters
	filters := map[string]interface{}{}
	counter_filters := 0
	if input_id_analisis == 0 {
		filters["id_analisis"] = input_id_analisis
		counter_filters += 1
	}
	if input_id_historia != "" {
		filters["id_historia_clinica"] = input_id_historia
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
	id_historia_clinica,id_analisis_campo,valor
	FROM AnalisisHistoria `
	if counter_filters > 0 {
		q += " WHERE "
		clausulas := make([]string, 0)
		for key, value := range filters {
			if key == "id_analisis" {
				clausulas = append(clausulas, fmt.Sprintf("%s = %d", key, value))
			} else {
				clausulas = append(clausulas, fmt.Sprintf("%s = '%s'", key, value))
			}
		}
		q += strings.Join(clausulas, " AND ")

	}
	rows, error_find := db.Query(ctx, q)
	if error_find != nil {
		return oListAnalisisHistoria, error_find
	}
	//Scan the row
	for rows.Next() {
		var oAnalisisHistoria models.AnalisisHistoria
		rows.Scan(
			&oAnalisisHistoria.IdHistoriaClinica,
			&oAnalisisHistoria.IdAnalisisCampo,
			&oAnalisisHistoria.Valor,
		)
		oListAnalisisHistoria = append(oListAnalisisHistoria, oAnalisisHistoria)
	}
	if error_find != nil {
		return oListAnalisisHistoria, error_find
	}

	//Return all Analisis Historia
	return oListAnalisisHistoria, nil
}
