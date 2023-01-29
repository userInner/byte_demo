package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
	"titok_v1/dao"
	"titok_v1/dto"
	"titok_v1/middleware"
	"titok_v1/models"
	"titok_v1/response"
	"titok_v1/service"
	"titok_v1/utils"
)

var (
	InvaildMsg = "参数错误"
)

func CommentAction(c *gin.Context) {
	commentServ := &service.CommentService{}
	err := c.ShouldBindQuery(commentServ)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.BuildCommentActionDto(response.InvaildParameCode, InvaildMsg+err.Error(), nil))
		return
	}

	// 验证token
	u_id, err := middleware.VerifyToken(commentServ.Token)
	if err != nil {
		c.JSON(http.StatusOK, dto.BuildCommentActionDto(response.InvaildParameCode, InvaildMsg+err.Error(), nil))
		return
	}
	user, err := dao.GetUserByID(u_id)
	if err != nil { //  数据库错误
		c.JSON(404, dto.BuildCommentActionDto(response.MysqlDataGetErrorCode, response.MySqlDataGetError+err.Error(), nil))
		return
	}
	if user.ID == 0 { // 数据库无该用户
		c.JSON(http.StatusBadRequest, dto.BuildCommentActionDto(response.InvaildParameCode, InvaildMsg, nil))
		return
	}
	// video_id是否有该视频
	if len(commentServ.VideoID) == 0 {
		c.JSON(http.StatusBadRequest, dto.BuildCommentActionDto(response.InvaildParameCode, InvaildMsg, nil))
		return
	}
	integerVideoID, err := strconv.Atoi(commentServ.VideoID)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.BuildCommentActionDto(response.InvaildParameCode, "用户参数错误", nil))
		return
	}
	video, err := dao.GetVideoByID(&models.Video{ID: int64(integerVideoID)})
	if err != nil {
		c.JSON(404, dto.BuildCommentActionDto(response.MysqlDataGetErrorCode, response.MySqlDataGetError+err.Error(), nil))
		return
	}

	if video.AuthorID == 0 {
		c.JSON(http.StatusBadRequest, dto.BuildCommentActionDto("1001", "视频参数错误", nil))
		return
	}
	if commentServ.ActionType != "1" && commentServ.ActionType != "2" {
		c.JSON(http.StatusBadRequest, dto.BuildCommentActionDto(response.InvaildParameCode, InvaildMsg, nil))
		return
	}
	comment := &models.Comment{
		AuthorID:   user.ID,
		VideoID:    video.ID,
		Content:    commentServ.CommentText,
		CreateDate: time.Now().Format("01-02"),
		CreateTime: utils.GetTime(),
		UpdateTime: utils.GetTime(),
	}
	if commentServ.ActionType == "1" { // 发布评论

		resComment, err := dao.InsertComment(comment)
		if err != nil {
			c.JSON(404, dto.BuildCommentActionDto(response.MySqlDataInsertErrorCode, response.MySqlDataInsertErrorCode+err.Error(), nil))
			return
		}
		if resComment.ID == 0 { // 数据库没有该视频
			c.JSON(http.StatusBadRequest, dto.BuildCommentActionDto(response.InvaildParameCode, InvaildMsg, nil))
			return
		}
		resComment.Author = video.Author
		c.JSON(http.StatusOK, dto.BuildCommentActionDto(response.OKCode, response.OKMsg, resComment))
		return
	}

	// 删除评论
	comment_id, err := strconv.Atoi(commentServ.CommentID)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.BuildCommentActionDto(response.InvaildParameCode, InvaildMsg+err.Error(), nil))
		return
	}
	comment.ID = int64(comment_id)
	deleteFlag, err := dao.DeleteComment(comment)
	if err != nil || !deleteFlag {
		c.JSON(http.StatusBadRequest, dto.BuildCommentActionDto(response.InvaildParameCode, InvaildMsg+err.Error(), nil))
		return
	}
	c.JSON(http.StatusOK, dto.BuildCommentActionDto(response.OKCode, response.DeleteOkMsg, nil))
}

func CommentList(c *gin.Context) {
	token := c.Query("token")
	video_id := c.Query("video_id")
	integerVideoID, err := strconv.Atoi(video_id)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.BuildCommentListDto("1001", response.InvaildParame+err.Error(), nil))
		return
	}
	comments, err := dao.GetCommentByVideoID(&models.Video{ID: int64(integerVideoID)})
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.BuildCommentListDto("2001", response.MySqlDataInsertErrorCode+err.Error(), nil))
		return
	}

	u_id, err := middleware.VerifyToken(token)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.BuildCommentListDto("1001", response.MySqlDataInsertErrorCode+err.Error(), nil))
		return
	}

	for k, _ := range comments {
		comments[k].Author.IsFollow = dao.GetUserFollow(&models.User{ID: u_id}, &models.User{ID: comments[k].AuthorID})
	}
	fmt.Println(comments[0].Author.IsFollow)

	c.JSON(http.StatusOK, dto.BuildCommentListDto("0", response.OKMsg, comments))
}
