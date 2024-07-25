package todo

import (
	"testing"

	"github.com/stretchr/testify/suite"
	"github.com/tkwang0530/dependency-injection-demo/models"
	"github.com/tkwang0530/dependency-injection-demo/repository/todo/mocks"
)

type TodoServiceSuite struct {
	suite.Suite
	service  Service
	mockRepo *mocks.Repository
}

func (s *TodoServiceSuite) SetupTest() {
	s.mockRepo = new(mocks.Repository)
	s.service = New(s.mockRepo)
}

func (s *TodoServiceSuite) TearDownTest() {
	s.mockRepo.AssertExpectations(s.T())
}

func (s *TodoServiceSuite) TestAdd() {
	task := "Buy milk"
	s.mockRepo.On("Add", task).Return(1).Once()
	s.mockRepo.On("GetAll").Return([]models.Todo{{ID: 1, Task: task}}).Once()

	id := s.service.Add(task)
	todos := s.service.GetAll()

	s.Assert().Equal(1, id)
	s.Assert().Equal(1, len(todos))
	s.Assert().Equal(task, todos[0].Task)
}

func (s *TodoServiceSuite) TestGetAll() {
	tasks := []string{"Buy milk", "Buy eggs"}
	s.mockRepo.On("GetAll").Return([]models.Todo{
		{ID: 1, Task: tasks[0]},
		{ID: 2, Task: tasks[1]},
	}).Once()

	todos := s.service.GetAll()

	s.Assert().Equal(2, len(todos))
	s.Assert().Equal(tasks[0], todos[0].Task)
	s.Assert().Equal(tasks[1], todos[1].Task)
}

func (s *TodoServiceSuite) TestDelete() {
	id := 1
	s.mockRepo.On("Delete", id).Once()

	s.service.Delete(id)
}

func TestTodoServiceSuite(t *testing.T) {
	suite.Run(t, new(TodoServiceSuite))
}
