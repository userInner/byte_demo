package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
	"titok_v1/dao"
	"titok_v1/dto"
	"titok_v1/middleware"
	"titok_v1/models"
	"titok_v1/service"
	"titok_v1/utils"
)

func PublishVideoByUser(c *gin.Context) {
	publishForm := &service.PublishForm{}
	err := c.ShouldBind(publishForm)
	if err != nil {
		c.JSON(http.StatusUnauthorized, dto.BuildPublishDto("1001", "参数绑定错误"))
		return
	}

	if len(publishForm.Title) == 0 {
		c.JSON(http.StatusUnauthorized, dto.BuildPublishDto("1001", "请填写标题"))
		return
	}

	u_id, err := middleware.VerifyToken(publishForm.Token)
	if err != nil {
		c.JSON(http.StatusUnauthorized, dto.BuildPublishDto("1001", "用户验证失败"))
		return
	}

	user, err := dao.GetUserByID(u_id)
	if err != nil {
		c.JSON(http.StatusUnauthorized, dto.BuildPublishDto("1001", "数据库错误"+err.Error()))
		return
	}

	if user.ID == 0 {
		c.JSON(http.StatusOK, dto.BuildPublishDto("1001", "用户验证失败"))
		return
	}
	UserIdPre := strconv.Itoa(int(u_id))
	// 文件名
	fileName := utils.Pre + UserIdPre + time.Now().Format("20060102150405") + publishForm.Title

	// 上传到minio服务
	err = utils.UploadUserVideo(c, &models.Video{
		AuthorID:   u_id,
		CoverURL:   "",
		PlayURL:    "",
		Title:      publishForm.Title,
		CreateTime: time.Now().Local(),
		UpdateTime: time.Now().Local(),
	}, fileName, publishForm.Data)

	if err != nil {
		c.JSON(http.StatusUnauthorized, dto.BuildPublishDto("1001", "文件传输错误"+err.Error()))
		return
	}

	c.JSON(http.StatusOK, dto.BuildPublishDto("0", "上传成功"))
}
