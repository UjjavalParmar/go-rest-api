package main

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"

	"github.com/gorilla/mux"
)

var errNameRequired = errors.New("Name is mandatory")
var errProjectIDRequired = errors.New("Project ID is required")
var errUserIDRequired = errors.New("User id is required")

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
	if err != nil {
		return
	}

	if err := validateTaskPayload(task); err != nil {
		return
	}

	t, err := s.store.CreateTask(task)

	if err != nil {
		return
	}

}

func (s *TaskServices) handleGetTask(w http.ResponseWriter, r *http.Request) {

}

func validateTaskPayload(task *Task) error {
	if task.Name == "" {
		return errNameRequired
	}

	if task.ProjectID == 0 {
		return errProjectIDRequired
	}

	if task.AssignedToID == 0 {
		return errUserIDRequired
	}

	return nil
}
