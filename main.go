package main

import (
	"database/sql"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

func newRouter() *mux.Router {
	r := mux.NewRouter()

	// Declare the static file directory and point it to the directory
	staticFileDictory := http.Dir("./list/")
	// Declare the handler
	staticFileHandler := http.StripPrefix("/list/", http.FileServer(staticFileDictory))
	r.PathPrefix("/list/").Handler(staticFileHandler).Methods("GET")

	r.HandleFunc("/todo", createTodoHandler).Methods("POST")
	r.HandleFunc("/todo", getTodoHandler).Methods("GET")

	return r
}

func main() {
	connString := "dbname=todo sslmode=disable"
	db, err := sql.Open("postgres", connString)
	if err != nil {
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	}

	InitStore(&dbStore{db: db})

	r := newRouter()
	http.ListenAndServe(":8080", r)
}
