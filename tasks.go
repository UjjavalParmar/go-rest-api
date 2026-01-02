package main

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/gorilla/mux"
)

type TaskServices struct {
	store Store
}

func NewTaskServices(s Store) *TaskServices {
	return &TaskServices{store: s}
}

func (s *TaskServices) RegisterRoutes(r *mux.Router) {
	r.HandleFunc("/tasks", s.handleCreateTask).Methods("POST")
	r.HandleFunc("/tasks/{id}", s.handleGetTask).Methods("GET")
}

func (s *TaskServices) handleCreateTask(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return
	}
	defer r.Body.Close()

	var task *Task

	err = json.Unmarshal(body, &task)
	if err != nil{
		return
	}

}

func (s *TaskServices) handleGetTask(w http.ResponseWriter, r *http.Request) {

}
