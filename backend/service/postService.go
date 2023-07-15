package service

import (
	"TravelGo/backend/model"
	"TravelGo/backend/repository"
)

type IPostService interface {
	CreatePost(post *model.TravelPost) error
	GetPosts(condition, params string) ([]model.TravelPost, error)
	EditPost(post *model.TravelPost) error
	DeletePost(id int) error
}

type PostService struct {
	repo repository.IPostRepository
}

func (p *PostService) CreatePost(post *model.TravelPost) error {
	return p.repo.Insert(post)
}

func (p *PostService) GetPosts(condition, params string) ([]model.TravelPost, error) {
	return p.repo.Select(condition, params)
}

func (p *PostService) EditPost(post *model.TravelPost) error {
	return p.repo.Update(post)
}

func (p *PostService) DeletePost(id int) error {
	return p.repo.Delete(id)
}

func NewPostService(repo repository.IPostRepository) IPostService {
	return &PostService{repo}
}
