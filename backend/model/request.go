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

type SearchPostRequest struct {
	Type     string `json:"type"`
	Keywords string `json:"keywords"`
}
