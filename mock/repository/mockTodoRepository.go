package repository

import (
	"github.com/golang/mock/gomock"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"learnApi/model"
	"reflect"
)

// MockTodoRepository is a mock of TodoRepository interface.
type MockTodoRepository struct {
	ctrl     *gomock.Controller
	recorder *MockTodoRepositoryMockRecorder
}

// MockTodoRepositoryMockRecorder is the mock recorder for MockTodoRepository.
type MockTodoRepositoryMockRecorder struct {
	mock *MockTodoRepository
}

// NewMockTodoRepository creates a new mock instance.
func NewMockTodoRepository(ctrl *gomock.Controller) *MockTodoRepository {
	mock := &MockTodoRepository{ctrl: ctrl}
	mock.recorder = &MockTodoRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTodoRepository) EXPECT() *MockTodoRepositoryMockRecorder {
	return m.recorder
}

// Delete mocks base method.
func (m *MockTodoRepository) Delete(arg0 primitive.ObjectID) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Delete indicates an expected call of Delete.
func (mr *MockTodoRepositoryMockRecorder) Delete(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockTodoRepository)(nil).Delete), arg0)
}

// GetAll mocks base method.
func (m *MockTodoRepository) GetAll() ([]model.Todo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAll")
	ret0, _ := ret[0].([]model.Todo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAll indicates an expected call of GetAll.
func (mr *MockTodoRepositoryMockRecorder) GetAll() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAll", reflect.TypeOf((*MockTodoRepository)(nil).GetAll))
}

// Insert mocks base method.
func (m *MockTodoRepository) Insert(arg0 model.Todo) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Insert", arg0)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Insert indicates an expected call of Insert.
func (mr *MockTodoRepositoryMockRecorder) Insert(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Insert", reflect.TypeOf((*MockTodoRepository)(nil).Insert), arg0)
}
