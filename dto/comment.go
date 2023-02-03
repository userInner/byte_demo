package dto

import (
	"titok_v1/models"
)

type CommentDto struct {
	Content    string   `json:"content"`     // 评论内容
	CreateDate string   `json:"create_date"` // 评论发布日期，格式 mm-dd
	ID         int64    `json:"id"`          // 评论id
	Author     *UserDto `json:"user"`        // 评论用户信息
}

type CommentListDto struct {
	CommentList []CommentDto `json:"comment_list"` // 评论列表
	StatusCode  string       `json:"status_code"`  // 状态码，0-成功，其他值-失败
	StatusMsg   string       `json:"status_msg"`   // 返回状态描述
}

// 评论
type CommentActionDto struct {
	StatusCode string      `json:"status_code"` // 状态码，0-成功，其他值-失败
	StatusMsg  string      `json:"status_msg"`  // 返回状态描述
	Comment    *CommentDto `json:"comment"`     // 评论内容
}

func BuildCommentDto(comment *models.Comment) *CommentDto {
	if comment == nil {
		return nil
	}
	return &CommentDto{
		Content:    comment.Content,
		CreateDate: comment.CreateDate,
		ID:         comment.ID,
		Author:     BuildUserDto(&comment.Author),
	}
}

func BuildCommentActionDto(code string, msg string, comment *models.Comment) *CommentActionDto {
	var actionDto CommentActionDto
	actionDto.Comment = BuildCommentDto(comment)
	actionDto.StatusMsg = msg
	actionDto.StatusCode = code
	return &actionDto
}

func BuildCommentListDto(code string, msg string, comment []models.Comment) *CommentListDto {
	if len(comment) == 0 || comment == nil {
		return &CommentListDto{
			CommentList: nil,
			StatusCode:  code,
			StatusMsg:   msg,
		}
	}
	resCommentDtos := make([]CommentDto, len(comment))
	for k, _ := range comment {
		resCommentDtos[k].Author = BuildUserDto(&comment[k].Author)
		resCommentDtos[k].ID = comment[k].ID
		resCommentDtos[k].CreateDate = comment[k].CreateDate
		resCommentDtos[k].Content = comment[k].Content
		resCommentDtos[k].Author.IsFollow = comment[k].Author.IsFollow
	}
	return &CommentListDto{
		CommentList: resCommentDtos,
		StatusCode:  code,
		StatusMsg:   msg,
	}
}
