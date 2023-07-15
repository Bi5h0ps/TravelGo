package controller

import (
	"TravelGo/backend/model"
	"TravelGo/backend/service"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"strings"
	"time"
)

type PostController struct {
	PostService service.IPostService
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

func (p *PostController) PostGetPosts(ctx *gin.Context) {
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
		body["start_date"] = v.StartDate.Format(dataFormat)
		body["end_date"] = v.EndDate.Format(dataFormat)
		body["tags"] = strings.Split(v.Tags, "|")
		body["comment_count"] = 0
		result[i] = body
	}
	SuccessResponse(ctx, gin.H{
		"posts": result,
		"total": len(result),
	})
	return
}
