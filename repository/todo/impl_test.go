package todo

import (
	"testing"

	"github.com/stretchr/testify/suite"
	"github.com/tkwang0530/dependency-injection-demo/db"
)

type TodoRepositorySuite struct {
	suite.Suite
	repo Repository
	db   *db.InMemoryDB
}

func (s *TodoRepositorySuite) SetupTest() {
	s.db = db.NewInMemoryDB()
	s.repo = New(s.db)
}

func (s *TodoRepositorySuite) TearDownTest() {
	// No need to do anything here
}

func (s *TodoRepositorySuite) TestAdd() {
	task := "Buy milk"
	id := s.repo.Add(task)
	todos := s.repo.GetAll()

	s.Assert().Equal(1, id)
	s.Assert().Equal(1, len(todos))
	s.Assert().Equal(task, todos[0].Task)
}

func (s *TodoRepositorySuite) TestGetAll() {
	tasks := []string{"Buy milk", "Buy eggs"}
	s.repo.Add(tasks[0])
	s.repo.Add(tasks[1])

	todos := s.repo.GetAll()

	s.Assert().Equal(2, len(todos))
	s.Assert().Equal(tasks[0], todos[0].Task)
	s.Assert().Equal(tasks[1], todos[1].Task)
}

func (s *TodoRepositorySuite) TestDelete() {
	id := s.repo.Add("Buy milk")
	s.repo.Delete(id)
	todos := s.repo.GetAll()

	s.Assert().Equal(0, len(todos))
}

func TestTodoRepositorySuite(t *testing.T) {
	suite.Run(t, new(TodoRepositorySuite))
}
