package documento_historia

import (
	"context"
	"fmt"
	"strings"
	"time"

	configs "github.com/UPC-Works/api-aliviate/configs"
	models "github.com/UPC-Works/api-aliviate/models"
)

func Pg_FindMultiple(input_id_historia string) ([]models.DocumentosHistoria, error) {

	//Initialization
	var oListDocumentosHistoria []models.DocumentosHistoria

	//Define the filters
	filters := map[string]interface{}{}
	counter_filters := 0
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
	id,
	id_historia_clinica,
	url
	FROM DocumentoHistoria `
	if counter_filters > 0 {
		q += " WHERE "
		clausulas := make([]string, 0)
		for key, value := range filters {
			clausulas = append(clausulas, fmt.Sprintf("%s = '%s'", key, value))
		}
		q += strings.Join(clausulas, " AND ")

	}
	rows, error_find := db.Query(ctx, q)
	if error_find != nil {
		return oListDocumentosHistoria, error_find
	}
	//Scan the row
	for rows.Next() {
		var oDocumentosHistoria models.DocumentosHistoria
		rows.Scan(
			&oDocumentosHistoria.Id,
			&oDocumentosHistoria.IdHistoriaClinica,
			&oDocumentosHistoria.Url,
		)
		oListDocumentosHistoria = append(oListDocumentosHistoria, oDocumentosHistoria)
	}
	if error_find != nil {
		return oListDocumentosHistoria, error_find
	}

	//Return all Analisis Historia
	return oListDocumentosHistoria, nil
}
