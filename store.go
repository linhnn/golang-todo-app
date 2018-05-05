package main

import "database/sql"

type Store interface {
	CreateTodo(todo *Todo) error
	GetTodos() ([]*Todo, error)
}

type dbStore struct {
	db *sql.DB
}

func (store *dbStore) CreateTodo(todo *Todo) error {
	_, err := store.db.Query("INSERT INTO todo(name, date) VALUES($1, $2)", todo.Name, todo.Date)
	return err
}

func (store *dbStore) GetTodos() ([]*Todo, error) {
	todoes := []*Todo{}

	rows, err := store.db.Query("SELECT name, date FROM todo")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		todo := &Todo{}

		if err := rows.Scan(&todo.Name, &todo.Date); err != nil {
			return nil, err
		}

		todoes = append(todoes, todo)
	}

	return todoes, nil
}

var store Store

func InitStore(s Store) {
	store = s
}
