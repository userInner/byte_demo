package utils

import (
	"bytes"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/minio/minio-go/v7"
	"log"
	"mime/multipart"
	"net/url"
	"strings"
	"time"
	"titok_v1/common"
	"titok_v1/dao"
	"titok_v1/models"
)

const (
	Pre         = "minio"
	suf         = ".mp4"
	application = "video/mpeg4"
	bucketName  = "minio8"
)

func CreateBucket(c *gin.Context, user_id string) bool {
	bucketName := Pre + user_id
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

// 上传
func UploadUserVideo(c *gin.Context, modelVideo *models.Video, fileName string, video *multipart.FileHeader) error {
	if video == nil {
		return errors.New("upload文件打开错误")
	}
	src, err := video.Open()
	if err != nil {
		return errors.New("upload文件打开错误" + err.Error())
	}

	defer src.Close()
	_, err = common.GetMinio().PutObject(c, bucketName, uuid.New().String()+fileName+".mp4", src, -1, minio.PutObjectOptions{ContentType: application})
	if err != nil {
		return err
	}
	// 截图视频URL
	url, err := GetMinioUrl(c, fileName+".mp4", 0)
	playUrl := strings.Split(url.String(), "?")[0]
	if err != nil {
		return errors.New("获取VideoURl失败" + err.Error())
	}

	// 获取封面
	coverData, err := GetVideoCover(playUrl)
	if err != nil {
		return errors.New("获取Video失败" + err.Error())
	}

	// 上传封面
	coverUrl := uuid.New().String()
	coverReader := bytes.NewReader(coverData)
	_, err = common.GetMinio().PutObject(c, bucketName, coverUrl+".jpeg", coverReader, -1, minio.PutObjectOptions{ContentType: "image/jpeg"})

	if err != nil {
		return errors.New("minio上传失败" + err.Error())
	}
	coverURL, err := GetMinioUrl(c, coverUrl+".jpeg", 0)
	if err != nil {
		return errors.New("minio获取url失败" + err.Error())
	}
	CoverURL := strings.Split(coverURL.String(), "?")[0]
	modelVideo.PlayURL = playUrl
	modelVideo.CoverURL = CoverURL

	return dao.CreateVideo(modelVideo)
}

func GetMinioUrl(c *gin.Context, fileName string, expires time.Duration) (*url.URL, error) {
	reqParam := make(url.Values)
	if expires <= 0 {
		expires = time.Second * 60 * 60 * 24
	}
	url, err := common.GetMinio().PresignedGetObject(c, bucketName, fileName, expires, reqParam)
	if err != nil {
		return nil, errors.New("获取url失败" + err.Error())
	}
	return url, nil
}
