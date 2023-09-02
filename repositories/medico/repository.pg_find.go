package medico

import (
	"context"
	"fmt"
	"strings"
	"time"

	configs "github.com/UPC-Works/api-aliviate/configs"
	models "github.com/UPC-Works/api-aliviate/models"
)

func Pg_FindOne(input_email string) (models.Medico, error) {

	//Initialization
	var oMedico models.Medico

	//Define the filters
	filters := map[string]interface{}{}
	counter_filters := 0
	if input_email != "" {
		filters["email"] = input_email
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
		id_establecimiento ,
		nombre             ,
		apellido           ,
		colegiatura        ,
		documento_identidad,
		correo             ,
		contrasenia        ,
		direccion          ,
		fecha_registro     
FROM Medico `
	if counter_filters > 0 {
		q += " WHERE "
		clausulas := make([]string, 0)
		for key, value := range filters {
			clausulas = append(clausulas, fmt.Sprintf("%s = '%s'", key, value))
		}
		q += strings.Join(clausulas, " AND ")

	}

	error_find := db.QueryRow(ctx, q).Scan(
		&oMedico.Id,
		&oMedico.IdEstablecimiento,
		&oMedico.Nombre,
		&oMedico.Apellido,
		&oMedico.Colegiatura,
		&oMedico.DocumentoIdentidad,
		&oMedico.Correo,
		&oMedico.Contrasenia,
		&oMedico.FechaRegistro,
		&oMedico.Direccion)

	if error_find != nil {
		return oMedico, error_find
	}

	//Return one medico
	return oMedico, nil
}

func Pg_FindMultiple(input_id string, input_limit int, input_offset int) ([]models.Medico, error) {

	//Initialization
	var oListMedico []models.Medico

	//Define the filters
	filters := map[string]interface{}{}
	counter_filters := 0
	if input_id != "" {
		filters["id"] = input_id
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
		id_establecimiento ,
		id_distrito        ,
		nombre             ,
		apellido           ,
		colegiatura        ,
		documento_identidad,
		correo             ,
		contrasenia        ,
		direccion          ,
		fecha_registro     
FROM Medico `
	if counter_filters > 0 {
		q += " WHERE "
		clausulas := make([]string, 0)
		for key, value := range filters {
			clausulas = append(clausulas, fmt.Sprintf("%s = '%s'", key, value))
		}
		q += strings.Join(clausulas, " AND ")

	}
	rows, error_find := db.Query(ctx, q+" ORDER BY fecha_registro DESC LIMIT $1 OFFSET $2", input_limit, input_offset)
	if error_find != nil {
		return oListMedico, error_find
	}

	//Scan the row
	for rows.Next() {
		var oMedico models.Medico
		rows.Scan(
			&oMedico.Id,
			&oMedico.IdEstablecimiento,
			&oMedico.Nombre,
			&oMedico.Apellido,
			&oMedico.Colegiatura,
			&oMedico.DocumentoIdentidad,
			&oMedico.Correo,
			&oMedico.Contrasenia,
			&oMedico.FechaRegistro,
			&oMedico.Direccion)
		oListMedico = append(oListMedico, oMedico)
	}

	if error_find != nil {
		return oListMedico, error_find
	}

	//Return the list of medicos
	return oListMedico, nil
}
