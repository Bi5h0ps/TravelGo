package repository

import (
	"TravelGo/backend/model"
	"TravelGo/backend/provider"
	"errors"
	"strconv"
)

type IPostRepository interface {
	Conn() error
	Select(condition, params string) ([]model.TravelPost, error)
	SelectAll() ([]model.TravelPost, error)
	Insert(post *model.TravelPost) (int, error)
	Update(post *model.TravelPost) error
	Delete(id int) error
}

type PostRepository struct{}

func (p *PostRepository) Conn() (err error) {
	err = provider.DatabaseEngine.AutoMigrate(&model.TravelPost{})
	return
}

func (p *PostRepository) Select(condition, params string) (result []model.TravelPost, err error) {
	selection := provider.DatabaseEngine.Where(condition, params).Where("is_deleted", 0).Find(&result)
	if selection.Error != nil {
		return result, selection.Error
	}
	return
}

func (p *PostRepository) SelectAll() (result []model.TravelPost, err error) {
	selection := provider.DatabaseEngine.Where("is_deleted", 0).Find(&result)
	if selection.Error != nil {
		return result, selection.Error
	}
	return
}

func (p *PostRepository) Insert(post *model.TravelPost) (id int, err error) {
	if err = p.Conn(); err != nil {
		return
	}
	if result := provider.DatabaseEngine.Create(&post); result.Error != nil {
		return 0, result.Error
	}
	return post.ID, err
}

func (p *PostRepository) Update(post *model.TravelPost) (err error) {
	if err = p.Conn(); err != nil {
		return
	}
	writeBack := provider.DatabaseEngine.Save(post)
	if writeBack.Error != nil {
		return writeBack.Error
	}
	return
}

func (p *PostRepository) Delete(id int) (err error) {
	if err = p.Conn(); err != nil {
		return
	}
	data, selectionError := p.Select("ID = ?", strconv.Itoa(id))
	if selectionError != nil {
		return selectionError
	} else if len(data) == 0 {
		return errors.New("no such post found")
	}
	tuple := data[0]
	tuple.IsDeleted = true
	writeBack := provider.DatabaseEngine.Save(&tuple)
	if writeBack.Error != nil {
		return writeBack.Error
	}
	return
}

func NewPostRepository() IPostRepository {
	return &PostRepository{}
}
