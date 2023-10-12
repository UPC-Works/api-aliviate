package prediccion

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"

	models "github.com/UPC-Works/api-aliviate/models"
)

func EnviarArchivo(nombre_archivo string) []models.PrediccionEnfermedad {

	var predicciones []models.PrediccionEnfermedad
	nombre_completo_archivo := nombre_archivo + ".xlsx"

	// Abrir el archivo que deseas subir
	file, _ := os.Open(nombre_completo_archivo)
	defer file.Close()

	// Crear un búfer para almacenar la solicitud multipart
	var requestBody bytes.Buffer
	writer := multipart.NewWriter(&requestBody)

	// Agregar el archivo al formulario
	part, _ := writer.CreateFormFile("archivo", nombre_completo_archivo)

	_, error_copy := io.Copy(part, file)
	if error_copy != nil {
		fmt.Println(error_copy)
		return predicciones
	}

	// Cerrar el escritor de form-data
	writer.Close()

	// Realizar una solicitud POST al endpoint
	url := "https://c6a8-181-67-74-103.ngrok-free.app/subir-archivo" // Reemplaza con la URL de tu endpoint
	req, err := http.NewRequest("POST", url, &requestBody)
	if err != nil {
		fmt.Println(err)
		return predicciones
	}

	// Establecer el encabezado Content-Type necesario para form-data
	req.Header.Set("Content-Type", writer.FormDataContentType())

	// Realizar la solicitud
	client := &http.Client{}
	response_http, _ := client.Do(req)

	//Decoding the response
	var response_predicciones []models.PrediccionEnfermedad
	error_decode := json.NewDecoder(response_http.Body).Decode(&response_predicciones)
	if error_decode != nil {
		return predicciones
	}

	return response_predicciones
}
