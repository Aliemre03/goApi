package service

import (
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"learnApi/mock/repository"
	"learnApi/model"
	"testing"
)

var mockRepository *repository.MockTodoRepository
var service TodoService

var FakeData = []model.Todo{
	{primitive.NewObjectID(), "Title 1", "Content1"},
	{primitive.NewObjectID(), "Title 2", "Content2"},
	{primitive.NewObjectID(), "Title 3", "Content3"},
}

func setup(t *testing.T) func() {
	ct := gomock.NewController(t)
	defer ct.Finish()

	mockRepository = repository.NewMockTodoRepository(ct)
	service = NewTodoService(mockRepository)

	return func() {
		service = nil
		defer ct.Finish()
	}
}
func TestDefaultTodoService_TodoGetAll(t *testing.T) {
	td := setup(t)
	defer td()

	mockRepository.EXPECT().GetAll().Return(FakeData, nil)
	result, err := service.TodoGetAll()

	if err != nil {
		t.Error(err)
	}

	assert.NotEmpty(t, result)
}
