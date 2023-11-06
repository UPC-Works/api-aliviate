package analisis_laboratorio

import (
	"context"
	"time"

	configs "github.com/UPC-Works/api-aliviate/configs"
	models "github.com/UPC-Works/api-aliviate/models"
)

func Pg_FindMultiple() ([]models.AnalisisLaboratorioCodigo, error) {

	//Initialization
	var oListAnalisisLaboratorioCodigo []models.AnalisisLaboratorioCodigo

	//Context timing
	ctx, cancel := context.WithTimeout(context.Background(), 8*time.Second)
	//Cancel context
	defer cancel()

	//Start the connection
	db := configs.Conn_Pg_DB()

	//Define the query
	q := `SELECT 
		id,
		nombre
	FROM AnalisisLaboratorioCodigo `

	rows, error_find := db.Query(ctx, q)
	if error_find != nil {
		return oListAnalisisLaboratorioCodigo, error_find
	}
	//Scan the row
	for rows.Next() {
		var oAnalisisLaboratorioCodigo models.AnalisisLaboratorioCodigo
		rows.Scan(
			&oAnalisisLaboratorioCodigo.Id,
			&oAnalisisLaboratorioCodigo.Nombre,
		)
		oListAnalisisLaboratorioCodigo = append(oListAnalisisLaboratorioCodigo, oAnalisisLaboratorioCodigo)
	}
	if error_find != nil {
		return oListAnalisisLaboratorioCodigo, error_find
	}

	//Return all Analisis Laboratorio Codigo
	return oListAnalisisLaboratorioCodigo, nil
}
