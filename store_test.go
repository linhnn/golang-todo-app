package main

import (
	"database/sql"
	"testing"
	"time"

	"github.com/stretchr/testify/suite"
)

type StoreSuite struct {
	suite.Suite
	store *dbStore
	db    *sql.DB
}

func (s *StoreSuite) SetupSuite() {
	connString := "dbname=todo_temp sslmode=disable"
	db, err := sql.Open("postgres", connString)
	if err != nil {
		s.T().Fatal(err)
	}
	s.db = db
	s.store = &dbStore{db: db}
}

func (s *StoreSuite) SetupTest() {
	_, err := s.db.Query("DELETE FROM todo")
	if err != nil {
		s.T().Fatal(err)
	}
}

func (s *StoreSuite) TearDownSuite() {
	s.db.Close()
}

func TestStoreSuite(t *testing.T) {
	s := new(StoreSuite)
	suite.Run(t, s)
}

func (s *StoreSuite) TestCreateTodo() {
	d, _ := time.Parse("2006-02-01", "2018-05-04")

	s.store.CreateTodo(&Todo{
		Name: "test name",
		Date: d,
	})

	res, err := s.db.Query("SELECT count(*) FROM todo WHERE name='test name'")
	if err != nil {
		s.T().Fatal(err)
	}

	var count int
	for res.Next() {
		if err := res.Scan(&count); err != nil {
			s.T().Error(err)
		}
	}

	if count != 1 {
		s.T().Errorf("incorrect count, wanted 1, got %d", count)
	}
}

func (s *StoreSuite) TestGetTodo() {
	_, err := s.db.Query("INSERT INTO todo(name, date) VALUES('test', '2018-05-05')")
	if err != nil {
		s.T().Fatal(err)
	}

	todos, err := s.store.GetTodos()
	if err != nil {
		s.T().Fatal(err)
	}

	nTodos := len(todos)
	if nTodos != 1 {
		s.T().Errorf("incorrect count, wanted 1, got %d", nTodos)
	}

	d, _ := time.Parse("2006-02-01", "2018-05-05")
	expectedTodo := Todo{Name: "test", Date: d}

	if *todos[0] != expectedTodo {
		s.T().Errorf("incorrect details, expected %v, got %v", expectedTodo, *todos[0])
	}

}
