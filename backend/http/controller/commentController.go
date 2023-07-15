package controller

import (
	"TravelGo/backend/model"
	"TravelGo/backend/service"
	"errors"
	"github.com/gin-gonic/gin"
	"time"
)

type CommentController struct {
	CommentService service.ICommentService
}

func (c *CommentController) PostGetCommentUnderPost(ctx *gin.Context) {
	req := map[string]int{
		"post_id": 0,
	}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ErrorResponse(ctx, err)
		return
	}
	username := ctx.GetString("username")
	comments, err := c.CommentService.GetCommentsUnderPost(req["post_id"])
	if err != nil {
		ErrorResponse(ctx, err)
		return
	}
	result := make([]map[string]interface{}, len(comments))
	for i, v := range comments {
		result[i] = map[string]interface{}{
			"comment_id":   v.ID,
			"username":     v.Username,
			"content":      v.Content,
			"is_deletable": v.Username == username,
		}
	}
	SuccessResponse(ctx, result)
	return
}

func (c *CommentController) PostMakeComment(ctx *gin.Context) {
	var req *model.MakeCommentRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ErrorResponse(ctx, err)
		return
	}
	username := ctx.GetString("username")
	if username == "" {
		ErrorResponse(ctx, errors.New("please login first"))
		return
	}
	err := c.CommentService.PostComment(&model.Comment{
		PostId:    req.PostId,
		Username:  username,
		Content:   req.Content,
		IsDeleted: false,
		CreatedAt: time.Now(),
	})
	if err != nil {
		ErrorResponse(ctx, err)
		return
	}
	SuccessResponse(ctx, nil)
	return
}

func (c *CommentController) PostDeleteComment(ctx *gin.Context) {
	var req *model.DeleteCommentRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ErrorResponse(ctx, err)
		return
	}
	username := ctx.GetString("username")
	if username == "" {
		ErrorResponse(ctx, errors.New("please login first"))
		return
	}
	err := c.CommentService.DeleteComment(req.CommentId, username)
	if err != nil {
		ErrorResponse(ctx, err)
		return
	}
	SuccessResponse(ctx, nil)
	return
}
