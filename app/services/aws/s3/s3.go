package s3

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"mime/multipart"
	"os"
)

var AccessKeyID string
var SecretAccessKey string
var MyRegion string


func ConnectAws() *session.Session {
	AccessKeyID = os.Getenv("AWS_ACCESS_KEY_ID")
	SecretAccessKey = os.Getenv("AWS_SECRET_ACCESS_KEY")
	MyRegion = os.Getenv("AWS_REGION")
	sess, err := session.NewSession(
		&aws.Config{
			Region: aws.String(MyRegion),
			Credentials: credentials.NewStaticCredentials(
				AccessKeyID,
				SecretAccessKey,
				"", // a token will be created when the session it's used.
			),
		})
	if err != nil {
		panic(err)
	}
	return sess
}

func Upload(file multipart.FileHeader) (string, error) {
	sess := ConnectAws()
	uploader := s3manager.NewUploader(sess)
	MyBucket := os.Getenv("AWS_BUCKET")

	memoryFile, err := file.Open()

	up, err := uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(MyBucket),
		ACL:    aws.String("public-read"),
		Key:    aws.String(file.Filename),
		Body:   memoryFile,
	})

	if err != nil {
		return "", err
	} else {
		return up.Location, nil
	}
}
