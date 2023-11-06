package private

import (
	"os"

	configs "github.com/UPC-Works/api-aliviate/configs"
)

func UploadToSpaces(format string, temp_file_url string, file_body *os.File, id_business string) (string, error) {

	//Start session in S3
	error_sesion := configs.S3.NewSession()
	if error_sesion != nil {
		return "", error_sesion
	}

	//Upload file
	error_upload_s3 := configs.S3.Uploader(temp_file_url, "qatuna-public-space", file_body, format)
	if error_upload_s3 != nil {
		return "", error_upload_s3
	}

	//Generate url
	urlS3, error_generateurl := configs.S3.GenerateUrl(temp_file_url)
	if error_generateurl != nil {
		return "", error_generateurl
	}

	//OK
	return urlS3, nil
}
