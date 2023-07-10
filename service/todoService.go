package service

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"learnApi/dto"
	"learnApi/model"
	"learnApi/repository"
)

type DefaultTodoService struct {
	Repository repository.TodoRepository
}

type TodoService interface {
	TodoInsert(todo model.Todo) (*dto.TodoDTO, error)
	TodoGetAll() ([]model.Todo, error)
	TodoDelete(id primitive.ObjectID) (bool, error)
}

func (t DefaultTodoService) TodoInsert(todo model.Todo) (*dto.TodoDTO, error) {
	var res dto.TodoDTO
	if len(todo.Title) <= 2 {
		res.Status = false
		return &res, nil
	}

	result, err := t.Repository.Insert(todo)

	if err != nil || result == false {
		res.Status = false
		return &res, err
	}

	res = dto.TodoDTO{Status: result}
	return &res, nil
}

func (t DefaultTodoService) TodoGetAll() ([]model.Todo, error) {
	result, err := t.Repository.GetAll()
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (t DefaultTodoService) TodoDelete(id primitive.ObjectID) (bool, error) {
	result, err := t.Repository.Delete(id)

	if err != nil || result == false {
		return false, err
	}

	return true, nil
}

func NewTodoService(Repo repository.TodoRepository) DefaultTodoService {
	return DefaultTodoService{Repository: Repo}
}
