package repository

import (
	"TravelGo/backend/model"
	"TravelGo/backend/provider"
	"errors"
	"strconv"
)

type ICommentRepository interface {
	Conn() error
	Select(condition, params string) ([]model.Comment, error)
	Insert(comment *model.Comment) error
	Delete(id int, username string) error
}

type CommentRepository struct{}

func (c *CommentRepository) Conn() (err error) {
	err = provider.DatabaseEngine.AutoMigrate(&model.Comment{})
	return
}

func (c *CommentRepository) Select(condition, params string) (result []model.Comment, err error) {
	selection := provider.DatabaseEngine.
		Where(condition, params).
		Where("is_deleted", false).
		Order("created_at DESC").Find(&result)
	if selection.Error != nil {
		return result, selection.Error
	}
	return
}

func (c *CommentRepository) Insert(comment *model.Comment) (err error) {
	if err = c.Conn(); err != nil {
		return
	}
	if result := provider.DatabaseEngine.Create(comment); result.Error != nil {
		return result.Error
	}
	return
}

func (c *CommentRepository) Delete(id int, username string) (err error) {
	if err = c.Conn(); err != nil {
		return
	}
	data, selectionError := c.Select("ID = ?", strconv.Itoa(id))
	if selectionError != nil {
		return selectionError
	} else if len(data) == 0 {
		return errors.New("no such post found")
	}
	tuple := data[0]
	if tuple.Username != username {
		err = errors.New("not authorized to delete this comment")
		return
	}
	tuple.IsDeleted = true
	writeBack := provider.DatabaseEngine.Save(&tuple)
	if writeBack.Error != nil {
		return writeBack.Error
	}
	return
}

func NewCommentRepository() ICommentRepository {
	return &CommentRepository{}
}
