package prediccion

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"

	models "github.com/UPC-Works/api-aliviate/models"
	"github.com/joho/godotenv"
)

func EnviarArchivo(nombre_archivo string) []models.PrediccionEnfermedadShow {

	var predicciones []models.PrediccionEnfermedadShow
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

	// Carga las variables de entorno desde el archivo .env.local
	err := godotenv.Load(".env.local")
	if err != nil {
		log.Fatalf("Error al cargar el archivo .env.local: %v", err)
	}

	// Realizar una solicitud POST al endpoint
	url := os.Getenv("URL_MODELO_IA")
	log.Println("URL:", url)
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
	var response_predicciones []models.PrediccionEnfermedadShow
	error_decode := json.NewDecoder(response_http.Body).Decode(&response_predicciones)
	if error_decode != nil {
		return predicciones
	}
	log.Println("response_predicciones:", response_predicciones)
	return response_predicciones
}
