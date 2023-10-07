package config

import (
	"os"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

var minioClient minio.Client

func InitializeMinioClient() (*minio.Client, error) {
	// Initialize a new MinIO client object.
	Client, err := minio.New(os.Getenv("WASABI_ENDPOINT"), &minio.Options{
		Creds:  credentials.NewStaticV4(os.Getenv("WASABI_ACCESS_KEY"), os.Getenv("WASABI_SECRET_KEY"), ""),
		Secure: true,
	})
	if err != nil {
		return nil, err
	}

	minioClient = *Client
	return Client, nil
}

func GetClientMinio() minio.Client {
	return minioClient
}
