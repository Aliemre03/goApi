package app

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"learnApi/mock/service"
	"learnApi/model"
	"net/http/httptest"
	"testing"
)

var td TodoHandler
var mockService *service.MockTodoService

func setup(t *testing.T) func() {
	ctrl := gomock.NewController(t)

	mockService = service.NewMockTodoService(ctrl)

	td = TodoHandler{mockService}

	return func() {
		defer ctrl.Finish()
	}
}

func TestTodoHandler_GetAllTodo(t *testing.T) {
	teardown := setup(t)
	defer teardown()

	router := fiber.New()
	router.Get("/api/todos", td.GetAllTodo)

	var FakeDataForHandler = []model.Todo{
		{primitive.NewObjectID(), "Title 1", "Content 1"},
		{primitive.NewObjectID(), "Title 2", "Content 2"},
		{primitive.NewObjectID(), "Title 3", "Content 3"},
	}

	mockService.EXPECT().TodoGetAll().Return(FakeDataForHandler, nil)

	req := httptest.NewRequest("GET", "/api/todos", nil)

	resp, _ := router.Test(req, 1)

	assert.Equal(t, 200, resp.StatusCode)
}
