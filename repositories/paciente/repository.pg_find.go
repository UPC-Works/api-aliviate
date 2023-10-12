package paciente

import (
	"context"
	"fmt"
	"strings"
	"time"

	configs "github.com/UPC-Works/api-aliviate/configs"
	models "github.com/UPC-Works/api-aliviate/models"
)

func Pg_FindOne(input_idpaciente string, input_documento_identidad int) (models.Paciente, error) {

	//Initialization
	var oPaciente models.Paciente

	//Define the filters
	filters := map[string]interface{}{}
	counter_filters := 0
	if input_documento_identidad != 0 {
		filters["documento_identidad"] = input_documento_identidad
		counter_filters += 1
	}
	if input_idpaciente != "" {
		filters["id"] = input_idpaciente
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
	nombre             ,
	apellido           ,
	fecha_nacimiento   ,
	genero             ,
	documento_identidad,
	fecha_registro 
FROM Paciente `
	if counter_filters > 0 {
		q += " WHERE "
		clausulas := make([]string, 0)
		for key, value := range filters {
			if key == "documento_identidad" {
				clausulas = append(clausulas, fmt.Sprintf("%s = %d", key, value))
			} else {
				clausulas = append(clausulas, fmt.Sprintf("%s = '%s'", key, value))
			}
		}
		q += strings.Join(clausulas, " AND ")

	}

	error_find := db.QueryRow(ctx, q).Scan(
		&oPaciente.Id,
		&oPaciente.Nombre,
		&oPaciente.Apellido,
		&oPaciente.FechaNacimiento,
		&oPaciente.Genero,
		&oPaciente.DocumentoIdentidad,
		&oPaciente.FechaRegistro)

	if error_find != nil {
		return oPaciente, error_find
	}

	//Return one paciente
	return oPaciente, nil
}

func Pg_FindMultiple(input_id string, input_documento_identidad int, input_limit int, input_offset int) ([]models.Paciente, error) {

	//Initialization
	var oListPaciente []models.Paciente

	//Define the filters
	filters := map[string]interface{}{}
	counter_filters := 0
	if input_id != "" {
		filters["id"] = input_id
		counter_filters += 1
	}
	if input_documento_identidad != 0 {
		filters["documento_identidad"] = input_documento_identidad
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
	nombre             ,
	apellido           ,
	fecha_nacimiento   ,
	genero             ,
	documento_identidad,
	fecha_registro ,
	grupo_sanguineo,
	rh_sanguineo,
	telefono
FROM Paciente `
	if counter_filters > 0 {
		q += " WHERE "
		clausulas := make([]string, 0)
		for key, value := range filters {
			if key == "documento_identidad" {
				clausulas = append(clausulas, fmt.Sprintf("%s = %d", key, value))
			} else {
				clausulas = append(clausulas, fmt.Sprintf("%s = '%s'", key, value))
			}
		}
		q += strings.Join(clausulas, " AND ")

	}
	rows, error_find := db.Query(ctx, q+" ORDER BY fecha_registro DESC LIMIT $1 OFFSET $2", input_limit, input_offset)
	if error_find != nil {
		return oListPaciente, error_find
	}

	//Scan the row
	for rows.Next() {
		var oPaciente models.Paciente
		rows.Scan(
			&oPaciente.Id,
			&oPaciente.Nombre,
			&oPaciente.Apellido,
			&oPaciente.FechaNacimiento,
			&oPaciente.Genero,
			&oPaciente.DocumentoIdentidad,
			&oPaciente.FechaRegistro,
			&oPaciente.GrupoSanguineo,
			&oPaciente.RhSanguineo,
			&oPaciente.Telefono)
		oListPaciente = append(oListPaciente, oPaciente)
	}

	if error_find != nil {
		return oListPaciente, error_find
	}

	//Return the list of pacientes
	return oListPaciente, nil
}
