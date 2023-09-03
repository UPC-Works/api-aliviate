package admin

import (
	"context"
	"fmt"
	"strings"
	"time"

	configs "github.com/UPC-Works/api-aliviate/configs"
	models "github.com/UPC-Works/api-aliviate/models"
)

func Pg_FindOne(input_correo string) (models.Admin, error) {

	//Initialization
	var oAdmin models.Admin

	//Define the filters
	filters := map[string]interface{}{}
	counter_filters := 0
	if input_correo != "" {
		filters["correo"] = input_correo
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
	nombre_completo             .
	correo,
	contrasenia   
FROM Admin `
	if counter_filters > 0 {
		q += " WHERE "
		clausulas := make([]string, 0)
		for key, value := range filters {
			clausulas = append(clausulas, fmt.Sprintf("%s = '%s'", key, value))
		}
		q += strings.Join(clausulas, " AND ")

	}

	error_find := db.QueryRow(ctx, q).Scan(
		&oAdmin.Id,
		&oAdmin.NombreCompleto,
		&oAdmin.Correo,
		&oAdmin.Contrasenia)

	if error_find != nil {
		return oAdmin, error_find
	}

	//Return one medico
	return oAdmin, nil
}
