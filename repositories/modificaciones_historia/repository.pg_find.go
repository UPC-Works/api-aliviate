package modificaciones_historia

import (
	"context"
	"fmt"
	"strings"
	"time"

	configs "github.com/UPC-Works/api-aliviate/configs"
	models "github.com/UPC-Works/api-aliviate/models"
)

func Pg_FindMultiple(input_id_historia_clinica string) ([]models.ModificacionesHistorias, error) {

	//Initialization
	var oListModificacionesHistorias []models.ModificacionesHistorias

	//Define the filters
	filters := map[string]interface{}{}
	counter_filters := 0
	if input_id_historia_clinica != "" {
		filters["id_historia_clinica"] = input_id_historia_clinica
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
	id                 ,
	id_historia_clinica             ,
	nombre_medico           ,
	actualizado_el          
FROM ModificacionesHistoria `
	if counter_filters > 0 {
		q += " WHERE "
		clausulas := make([]string, 0)
		for key, value := range filters {
			clausulas = append(clausulas, fmt.Sprintf("%s = '%s'", key, value))
		}
		q += strings.Join(clausulas, " AND ")

	}
	rows, error_find := db.Query(ctx, q+" ORDER BY actualizadoEl DESC")
	if error_find != nil {
		return oListModificacionesHistorias, error_find
	}

	//Scan the row
	for rows.Next() {
		var oModificacionesHistorias models.ModificacionesHistorias
		rows.Scan(
			&oModificacionesHistorias.Id,
			&oModificacionesHistorias.IdHistoriaClinica,
			&oModificacionesHistorias.NombreMedico,
			&oModificacionesHistorias.ActualizadoEl)
		oListModificacionesHistorias = append(oListModificacionesHistorias, oModificacionesHistorias)
	}

	if error_find != nil {
		return oListModificacionesHistorias, error_find
	}

	//Return the list of modificaciones historias
	return oListModificacionesHistorias, nil
}
