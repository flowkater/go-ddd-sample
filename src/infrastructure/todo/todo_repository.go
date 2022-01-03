package infrastructure

import (
	domain "github.com/flowkater/go-ddd-sample/src/domain/todo"
	"gorm.io/gorm"
)

type TodoRepository interface {
	Create(db *gorm.DB, todo *domain.Todo) (*domain.Todo, error)
	Update(db *gorm.DB, todo *domain.Todo) (*domain.Todo, error)
	FindById(db *gorm.DB, id uint) (*domain.Todo, error)
	UpdateDone(db *gorm.DB, todo *domain.Todo) (*domain.Todo, error)
}

type todoRepository struct{}

func NewTodoRepository() TodoRepository {
	return &todoRepository{}
}

func (t *todoRepository) Create(db *gorm.DB, todo *domain.Todo) (*domain.Todo, error) {
	if err := db.Create(todo).Error; err != nil {
		return nil, err
	}

	return todo, nil
}

func (t *todoRepository) Update(db *gorm.DB, todo *domain.Todo) (*domain.Todo, error) {
	if err := db.Updates(todo).Error; err != nil {
		return nil, err
	}

	return todo, nil
}

func (t *todoRepository) UpdateDone(db *gorm.DB, todo *domain.Todo) (*domain.Todo, error) {
	if err := db.Model(todo).Updates(map[string]interface{}{"Done": todo.Done}).Error; err != nil {
		return nil, err
	}

	return todo, nil
}

func (t *todoRepository) FindById(db *gorm.DB, id uint) (*domain.Todo, error) {
	var todo domain.Todo
	if err := db.Where("id = ?", id).First(&todo).Error; err != nil {
		return nil, err
	}

	return &todo, nil
}
