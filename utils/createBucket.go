package utils

import (
	"github.com/gin-gonic/gin"
	"github.com/minio/minio-go/v7"
	"log"
	"titok_v1/common"
)

const (
	pre = "minio"
)

func CreateBucket(c *gin.Context, user_id string) (bool) {
	bucketName := pre + user_id
	err := common.GetMinio().MakeBucket(c, bucketName, minio.MakeBucketOptions{Region: "cn-south-1", ObjectLocking: false})
	if err != nil {
		log.Println(err)
		exists, _ := common.GetMinio().BucketExists(c, bucketName)
		if exists {
			log.Printf("%s已经存在", bucketName)
			return false
		}
	}
	return true
}
