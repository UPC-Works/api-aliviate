package analisis_laboratorio_campo

import (
	"context"
	"fmt"
	"strings"
	"time"

	configs "github.com/UPC-Works/api-aliviate/configs"
	models "github.com/UPC-Works/api-aliviate/models"
)

func Pg_FindMultiple(input_id_analisis int) ([]models.AnalisisLaboratorioCampo, error) {

	//Initialization
	var oListAnalisisLaboratorioCampo []models.AnalisisLaboratorioCampo

	//Define the filters
	filters := map[string]interface{}{}
	counter_filters := 0
	if input_id_analisis == 0 {
		filters["id_analisis"] = input_id_analisis
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
	id,id_analisis,campo,campos_json,tipo
	FROM AnalisisLaboratorioCampo `
	if counter_filters > 0 {
		q += " WHERE "
		clausulas := make([]string, 0)
		for key, value := range filters {
			clausulas = append(clausulas, fmt.Sprintf("%s = %d", key, value))
		}
		q += strings.Join(clausulas, " AND ")

	}
	rows, error_find := db.Query(ctx, q)
	if error_find != nil {
		return oListAnalisisLaboratorioCampo, error_find
	}
	//Scan the row
	for rows.Next() {
		var oAnalisisLaboratorioCampo models.AnalisisLaboratorioCampo
		rows.Scan(
			&oAnalisisLaboratorioCampo.Id,
			&oAnalisisLaboratorioCampo.IdAnalisis,
			&oAnalisisLaboratorioCampo.Campo,
			&oAnalisisLaboratorioCampo.CampoJson,
			&oAnalisisLaboratorioCampo.Tipo,
		)
		oListAnalisisLaboratorioCampo = append(oListAnalisisLaboratorioCampo, oAnalisisLaboratorioCampo)
	}
	if error_find != nil {
		return oListAnalisisLaboratorioCampo, error_find
	}

	//Return all Analisis Laboratorio Campo
	return oListAnalisisLaboratorioCampo, nil
}
