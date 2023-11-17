package documento_historia

import (
	"io"
	"os"
	"strings"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"

	helpers "github.com/UPC-Works/api-aliviate/helpers"
	models "github.com/UPC-Works/api-aliviate/models"
	documento_historia_repository "github.com/UPC-Works/api-aliviate/repositories/documento_historia"
	private_services "github.com/UPC-Works/api-aliviate/services/private"
)

func Add(c echo.Context) error {

	//Get the id of the business
	id_historia_clinica := c.Param("idhistoriaclinica")

	//Get the file
	file, file_header, err_file := c.Request().FormFile("documento")
	if err_file != nil {
		return c.JSON(400, &helpers.ResponseString{
			Error: helpers.ErrorStructure{
				HasError: true,
				Detail:   "Revisa la estructura o los tipos de valores, detalles: " + err_file.Error(),
			},
			Data: ""})
	}

	//Get the fomat of file
	format := strings.Split(file_header.Filename, ".")[1]
	if format != "pdf" {
		return c.JSON(400, &helpers.ResponseString{
			Error: helpers.ErrorStructure{
				HasError: true,
				Detail:   "Tiene que ser documento PDF",
			},
			Data: ""})
	}

	//Generate the temp_file_url
	id_documento := uuid.New().String()
	temp_file_url := "s3-aliviate" + "/" + id_historia_clinica + id_documento + ".pdf"

	// Convert multipart.FileHeader to *os.File
	tmp_file, err := os.Create(file_header.Filename)
	if err != nil {
		return c.JSON(500, &helpers.ResponseString{
			Error: helpers.ErrorStructure{
				HasError: true,
				Detail:   "Error al convertir el archivo a os.File",
			},
			Data: ""})
	}

	//Close the file
	defer tmp_file.Close()

	// Copy the content from multipart file to the temporary file
	_, err = io.Copy(tmp_file, file)
	if err != nil {
		return c.JSON(500, &helpers.ResponseString{
			Error: helpers.ErrorStructure{
				HasError: true,
				Detail:   "Error al copiar el archivo convertido a os.File",
			},
			Data: ""})
	}

	//Open the file
	file_body, error_open := os.Open(file_header.Filename)
	if error_open != nil {
		return c.JSON(500, &helpers.ResponseString{
			Error: helpers.ErrorStructure{
				HasError: true,
				Detail:   "Error al abrir el archivo",
			},
			Data: ""})
	}

	//Upload to Spaces and generate the public url
	url_generated, error_generate_url := private_services.UploadToSpaces("application/pdf", temp_file_url, file_body, id_historia_clinica)
	if error_generate_url != nil {
		return c.JSON(500, &helpers.ResponseString{
			Error: helpers.ErrorStructure{
				HasError: true,
				Detail:   error_generate_url.Error(),
			},
			Data: ""})
	}

	//Close the file
	defer file_body.Close()
	defer file.Close()

	// Remove the file
	/*err_remove := os.Remove(file_body.Name())
	if err_remove != nil {
		return c.JSON(500, &helpers.ResponseString{
			Error: helpers.ErrorStructure{
				HasError: true,
				Detail:   "Error eliminando archivo",
			},
			Data: ""})
	}*/

	//Storage the Documento Historia
	new_documento_historia := models.NewDocumentosHistoria(uuid.New().String(), id_historia_clinica, url_generated)
	error_create_documento_historia := documento_historia_repository.Pg_Create(new_documento_historia)
	if error_create_documento_historia != nil {
		return c.JSON(500, &helpers.ResponseString{
			Error: helpers.ErrorStructure{
				HasError: true,
				Detail:   error_create_documento_historia.Error(),
			},
			Data: ""})
	}

	//OK
	return c.JSON(200, &helpers.ResponseString{
		Error: helpers.ErrorStructure{
			HasError: false,
			Detail:   "",
		},
		Data: "OK"})

}
