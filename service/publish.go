package service

import "mime/multipart"

// 投稿
type PublishForm struct {
	Data  *multipart.FileHeader `form:"data" bind:"required"`
	Token string                `form:"token" bind:"required"`
	Title string                `form:"title" bind:"required"`
}
