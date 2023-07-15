package model

type CreatePostRequest struct {
	PostTitle   string   `json:"title"`
	Destination string   `json:"destination"`
	StartDate   string   `json:"start_date"`
	EndDate     string   `json:"end_date"`
	Tags        []string `json:"tags"`
}

type EditPostRequest struct {
	PostId      int      `json:"post_id"`
	PostTitle   string   `json:"title"`
	Destination string   `json:"destination"`
	StartDate   string   `json:"start_date"`
	EndDate     string   `json:"end_date"`
	Tags        []string `json:"tags"`
}

type MakeCommentRequest struct {
	PostId  int    `json:"post_id"`
	Content string `json:"content"`
}

type DeleteCommentRequest struct {
	CommentId int    `json:"comment_id"`
	Username  string `json:"username"`
}

type SearchPostRequest struct {
	Type     string `json:"type"`
	Keywords string `json:"keywords"`
}
