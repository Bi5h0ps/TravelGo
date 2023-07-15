package service

import (
	"TravelGo/backend/model"
	"TravelGo/backend/repository"
	"strconv"
)

type ICommentService interface {
	GetCommentsUnderPost(postId int) (comments []model.Comment, err error)
	PostComment(comment *model.Comment) error
	DeleteComment(commentId int) error
}

type CommentService struct {
	repo repository.ICommentRepository
}

func (c *CommentService) GetCommentsUnderPost(postId int) (comments []model.Comment, err error) {
	return c.repo.Select("post_id", strconv.Itoa(postId))
}

func (c *CommentService) PostComment(comment *model.Comment) error {
	return c.repo.Insert(comment)
}

func (c *CommentService) DeleteComment(id int) error {
	return c.repo.Delete(id)
}

func NewCommentService(repo repository.ICommentRepository) ICommentService {
	return &CommentService{repo: repo}
}
