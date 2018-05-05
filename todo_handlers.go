package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type Todo struct {
	Name string    `json:"name"`
	Date time.Time `json:"date"`
}

func createTodoHandler(w http.ResponseWriter, r *http.Request) {
	todo := Todo{}

	err := r.ParseForm()
	if err != nil {
		fmt.Println(fmt.Errorf("Error: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	todo.Name = r.FormValue("name")
	todo.Date, _ = time.Parse("2006-01-02", r.FormValue("date"))

	err = store.CreateTodo(&todo)
	if err != nil {
		fmt.Println(err)
	}

	http.Redirect(w, r, "/list/", http.StatusFound)
}

func getTodoHandler(w http.ResponseWriter, r *http.Request) {
	todoes, err := store.GetTodos()

	todoListBytes, err := json.Marshal(todoes)
	if err != nil {
		fmt.Println(fmt.Errorf("Error: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write(todoListBytes)
}
