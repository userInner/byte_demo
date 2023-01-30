package dto

type PublishDto struct {
	StatusCode string `json:"status_code"`
	StatusMsg  string `json:"status_msg"`
}

func BuildPublishDto(code string, msg string) *PublishDto {
	return &PublishDto{
		StatusCode: code,
		StatusMsg:  msg,
	}
}
