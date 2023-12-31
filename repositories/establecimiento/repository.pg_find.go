package establecimiento

import (
	"context"
	"fmt"
	"strings"
	"time"

	configs "github.com/UPC-Works/api-aliviate/configs"
	models "github.com/UPC-Works/api-aliviate/models"
)

func Pg_FindMultiple() ([]models.Establecimiento, error) {

	//Initialization
	var oListEstablecimiento []models.Establecimiento

	//Define the filters
	filters := map[string]interface{}{}
	counter_filters := 0

	//Context timing
	ctx, cancel := context.WithTimeout(context.Background(), 8*time.Second)
	//Cancel context
	defer cancel()

	//Start the connection
	db := configs.Conn_Pg_DB()

	//Define the query
	q := `SELECT 
	id                 ,
	nombre             ,
	id_distrito,
	direccion
FROM Establecimiento `
	if counter_filters > 0 {
		q += " WHERE "
		clausulas := make([]string, 0)
		for key, value := range filters {
			clausulas = append(clausulas, fmt.Sprintf("%s = '%s'", key, value))
		}
		q += strings.Join(clausulas, " AND ")

	}
	rows, error_find := db.Query(ctx, q+" ORDER BY nombre DESC")
	if error_find != nil {
		return oListEstablecimiento, error_find
	}

	//Scan the row
	for rows.Next() {
		var oEstablecimiento models.Establecimiento
		rows.Scan(
			&oEstablecimiento.Id,
			&oEstablecimiento.Nombre,
			&oEstablecimiento.IdDistrito,
			&oEstablecimiento.Dirección)
		oListEstablecimiento = append(oListEstablecimiento, oEstablecimiento)
	}

	if error_find != nil {
		return oListEstablecimiento, error_find
	}

	//Return the list of establecimientos
	return oListEstablecimiento, nil
}
