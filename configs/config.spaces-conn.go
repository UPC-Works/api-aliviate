package config

import (
	"log"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/joho/godotenv"
)

var S3 *S3Client

func init() {
	S3 = new(S3Client)
}

type S3Client struct {
	Region string
	Sess   *session.Session
	Svc    *s3.S3
}

func (t *S3Client) NewSession() error {

	// Carga las variables de entorno desde el archivo .env.local
	err := godotenv.Load(".env.local")
	if err != nil {
		log.Fatalf("Error al cargar el archivo .env.local: %v", err)
	}

	endpoint := "https://sfo3.digitaloceanspaces.com"
	region := "us-east-1"

	sess, _ := session.NewSession(&aws.Config{
		Credentials: credentials.NewStaticCredentials(os.Getenv("AWS_ACCESS_KEY_ID"), os.Getenv("AWS_ACCESS_SECRET_KEY"), ""),
		Endpoint:    &endpoint,
		Region:      &region,
	})

	t.Sess = sess
	t.Svc = s3.New(t.Sess)

	return nil
}

func (t *S3Client) Uploader(temp_file_url string, bucket string, file_body *os.File, format string) error {

	// Create an uploader with the session and default options
	uploader := s3manager.NewUploader(t.Sess)

	// Upload the file to S3
	_, err := uploader.Upload(&s3manager.UploadInput{
		Bucket:      aws.String(bucket),
		ACL:         aws.String("public-read"),
		Key:         aws.String(temp_file_url),
		ContentType: aws.String(format),
		Body:        file_body,
	})
	if err != nil {
		return err
	}
	return nil
}

func (t *S3Client) GenerateUrl(temp_file_url string) (string, error) {

	URL := "https://qatuna-public-space.sfo3.cdn.digitaloceanspaces.com/" + temp_file_url

	return URL, nil
}

func (t *S3Client) DeleteObject(file_name string, bucket string) error {

	_, err := t.Svc.DeleteObject(&s3.DeleteObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(file_name),
	})

	if err != nil {
		return err
	}

	return nil
}
