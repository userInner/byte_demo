package common

import (
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"log"
)

const (
	endpoint        = "49.232.185.187:9000"
	accessKeyID     = "miniouser"
	secretAccessKey = "miniouser"
	useSSL          = false
)

var (
	MinioClient *minio.Client
	err         error
)

func InitMinio() {
	MinioClient, err = minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
		Secure: useSSL})
	if err != nil {
		panic(err)
	}
	log.Println("minio启动成功》》》》》》》》")
}

func GetMinio() *minio.Client {
	return MinioClient
}
