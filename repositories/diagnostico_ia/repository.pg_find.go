package diagnostico_ia

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	configs "github.com/UPC-Works/api-aliviate/configs"
	models "github.com/UPC-Works/api-aliviate/models"
)

func Pg_Find_EnfermedadesPredicciones(input_id_medico string) ([]models.EnfermedadPrediccion, error) {

	//Initialization
	var oListEnfermedadPrediccion []models.EnfermedadPrediccion

	//Define the filters
	filters := map[string]interface{}{}
	counter_filters := 0
	if input_id_medico != "" {
		filters["co.id_medico"] = input_id_medico
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
	dia.enfermedad,
	(SUM(
		CASE
		  WHEN dia.es_aceptado=true THEN 1
		  ELSE 0
		END
	  )::DECIMAL(10,2)/COUNT(dia.es_aceptado))::DECIMAL(10,2)*100 as probabilidad
  FROM diagnosticoia  AS dia JOIN consulta AS co ON dia.id_consulta=co.id`
	if counter_filters > 0 {
		q += " WHERE "
		clausulas := make([]string, 0)
		for key, value := range filters {
			clausulas = append(clausulas, fmt.Sprintf("%s = '%s'", key, value))
		}
		q += strings.Join(clausulas, " AND ")

	}

	log.Println("QUERY QUE SE ENVIA-------->", q+" GROUP BY dia.enfermedad  ORDER BY probabilidad DESC")

	rows, error_find := db.Query(ctx, q+" GROUP BY dia.enfermedad  ORDER BY probabilidad DESC")
	if error_find != nil {
		return oListEnfermedadPrediccion, error_find
	}
	//Scan the row
	for rows.Next() {
		var oEnfermedadPrediccion models.EnfermedadPrediccion
		rows.Scan(
			&oEnfermedadPrediccion.Enfermedad,
			&oEnfermedadPrediccion.PrediccionCorrecta,
		)
		oListEnfermedadPrediccion = append(oListEnfermedadPrediccion, oEnfermedadPrediccion)
	}
	if error_find != nil {
		return oListEnfermedadPrediccion, error_find
	}

	//Return all Enfermedad Prediccion
	return oListEnfermedadPrediccion, nil
}
