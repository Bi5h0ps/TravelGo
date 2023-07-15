package controller

import (
	"TravelGo/backend/model"
	"TravelGo/backend/service"
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	"strconv"
	"strings"
	"time"
)

type PostController struct {
	PostService    service.IPostService
	CommentService service.ICommentService
}

const (
	TYPE_DESTINATION = "destination"
	TYPE_TAG         = "tag"
)

var dataFormat = "2006-01-02"

func (p *PostController) PostCreatePost(ctx *gin.Context) {
	var req *model.CreatePostRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ErrorResponse(ctx, err)
		return
	}
	username := ctx.GetString("username")
	if username == "" {
		ErrorResponse(ctx, errors.New("please login first"))
		return
	}
	// Parse the frontend date string into a time.Time value
	startDate, err := time.Parse(dataFormat, req.StartDate)
	if err != nil {
		panic("failed to parse date")
	}
	endDate, err := time.Parse(dataFormat, req.EndDate)
	if err != nil {
		panic("failed to parse date")
	}
	err = p.PostService.CreatePost(&model.TravelPost{
		PostTitle:   req.PostTitle,
		Username:    username,
		Destination: req.Destination,
		StartDate:   startDate,
		EndDate:     endDate,
		Tags:        strings.Join(req.Tags, "|"),
	})
	if err != nil {
		ErrorResponse(ctx, err)
		return
	}
	SuccessResponse(ctx, nil)
	return
}

func (p *PostController) PostDeletePost(ctx *gin.Context) {
	req := map[string]int{
		"post_id": 0,
	}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ErrorResponse(ctx, err)
		return
	}
	err := p.PostService.DeletePost(req["post_id"])
	if err != nil {
		ErrorResponse(ctx, err)
		return
	}
	SuccessResponse(ctx, nil)
	return
}

func (p *PostController) PostEditPost(ctx *gin.Context) {
	var req *model.EditPostRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ErrorResponse(ctx, err)
		return
	}
	// Parse the frontend date string into a time.Time value
	startDate, err := time.Parse(dataFormat, req.StartDate)
	if err != nil {
		panic("failed to parse date")
	}
	endDate, err := time.Parse(dataFormat, req.EndDate)
	if err != nil {
		panic("failed to parse date")
	}
	err = p.PostService.EditPost(&model.TravelPost{
		ID:          req.PostId,
		PostTitle:   req.PostTitle,
		Destination: req.Destination,
		StartDate:   startDate,
		EndDate:     endDate,
		Tags:        strings.Join(req.Tags, "|"),
		IsDeleted:   false,
	})
	if err != nil {
		ErrorResponse(ctx, err)
		return
	}
	SuccessResponse(ctx, nil)
	return
}

func (p *PostController) PostSearchPosts(ctx *gin.Context) {
	var req *model.SearchPostRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ErrorResponse(ctx, err)
		return
	}
	condition := ""
	params := ""
	if req.Type == TYPE_DESTINATION {
		condition = "destination LIKE ?"
		params = "%" + req.Keywords + "%"
	} else if req.Type == TYPE_TAG {
		condition = "tags LIKE ?"
		params = "%" + req.Keywords + "%"
	}
	posts, err := p.PostService.GetPosts(condition, params)
	if err != nil {
		ErrorResponse(ctx, err)
		return
	}
	result := make([]map[string]interface{}, len(posts))
	body := map[string]interface{}{}
	for i, v := range posts {
		byteSlice, err := json.Marshal(v)
		err = json.Unmarshal(byteSlice, &body)
		if err != nil {
			ErrorResponse(ctx, err)
			return
		}
		comments, err := p.CommentService.GetCommentsUnderPost(v.ID)
		if err != nil {
			ErrorResponse(ctx, err)
			return
		}
		body["start_date"] = v.StartDate.Format(dataFormat)
		body["end_date"] = v.EndDate.Format(dataFormat)
		body["tags"] = strings.Split(v.Tags, "|")
		body["comment_count"] = len(comments)
		result[i] = body
	}
	SuccessResponse(ctx, gin.H{
		"posts": result,
		"total": len(result),
	})
	return
}

func (p *PostController) GetPosts(ctx *gin.Context) {
	posts, err := p.PostService.GetAllPosts()
	if err != nil {
		ErrorResponse(ctx, err)
		return
	}
	result := make([]map[string]interface{}, len(posts))
	body := map[string]interface{}{}
	for i, v := range posts {
		byteSlice, err := json.Marshal(v)
		err = json.Unmarshal(byteSlice, &body)
		if err != nil {
			ErrorResponse(ctx, err)
			return
		}
		comments, err := p.CommentService.GetCommentsUnderPost(v.ID)
		if err != nil {
			ErrorResponse(ctx, err)
			return
		}
		body["start_date"] = v.StartDate.Format(dataFormat)
		body["end_date"] = v.EndDate.Format(dataFormat)
		body["tags"] = strings.Split(v.Tags, "|")
		body["comment_count"] = len(comments)
		result[i] = body
		body = map[string]interface{}{}
	}
	SuccessResponse(ctx, gin.H{
		"posts": result,
		"total": len(result),
	})
	return
}

func (p *PostController) GetPostDetail(ctx *gin.Context) {
	req := map[string]int{
		"post_id": 0,
	}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ErrorResponse(ctx, err)
		return
	}
	posts, err := p.PostService.GetPosts("ID", strconv.Itoa(req["post_id"]))
	if err != nil {
		ErrorResponse(ctx, err)
		return
	}
	post := posts[0]
	comments, err := p.CommentService.GetCommentsUnderPost(req["post_id"])
	processed := make([]map[string]interface{}, len(comments))
	for i, v := range comments {
		processed[i] = map[string]interface{}{
			"comment_id": v.ID,
			"username":   v.Username,
			"content":    v.Content,
		}
	}
	byteSlice, err := json.Marshal(post)
	if err != nil {
		ErrorResponse(ctx, err)
		return
	}
	processedPost := map[string]interface{}{}
	err = json.Unmarshal(byteSlice, &processedPost)
	if err != nil {
		ErrorResponse(ctx, err)
		return
	}
	processedPost["start_date"] = post.StartDate.Format(dataFormat)
	processedPost["end_date"] = post.EndDate.Format(dataFormat)
	SuccessResponse(ctx, gin.H{
		"post_detail": processedPost,
		"comments":    processed,
	})
	return
}
