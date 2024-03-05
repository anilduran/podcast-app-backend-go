package utils

import (
	"context"
	"os"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/google/uuid"
)

var PresignerInstace *Presigner

type Presigner struct {
	PresignClient *s3.PresignClient
}

func InitializePresigner(s3Client *s3.Client) {
	PresignerInstace = &Presigner{
		PresignClient: s3.NewPresignClient(s3Client),
	}
}

func (presigner Presigner) PutObject(userId, folder string) (string, string, error) {

	id := uuid.New()

	key := folder + "/" + userId + "/" + id.String() + ".jpg"

	request, err := presigner.PresignClient.PresignPutObject(context.TODO(), &s3.PutObjectInput{
		Bucket: aws.String(os.Getenv("S3_BUCKET")),
		Key:    aws.String(key),
	}, func(opts *s3.PresignOptions) {
		opts.Expires = time.Duration(60 * time.Second * 15)
	})

	if err != nil {
		return "", "", err
	}

	return request.URL, key, nil

}
