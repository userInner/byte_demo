package service

type CommentService struct {
	Token       string `form:"token"`        // 必传：是
	VideoID     string `form:"video_id"`     // 必传：是
	ActionType  string `form:"action_type"`  // 必传：是 1发布评论 2删除评论
	CommentText string `form:"comment_text"` // 必传：否 action_type = 1 使用
	CommentID   string `form:"comment_id"`   // 必传：否 action_type = 2 使用
}
