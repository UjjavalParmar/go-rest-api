package main

import "database/sql"

type Storage struct {
	db *sql.DB
}
type Store interface {
}

func NewStore(db *sql.DB) Store {
	return nil
}

func (s *Storage) CreateUser() error {
	return nil
}

func (s *Storage) CreateTask(t *Task) (*Task, error) {
	rows, err := s.db.Exec("INSERT INTO tasks (name, status, projects_id, assigned_to) VALUES (?,?,?,?)", t.Name, t.Status, t.ProjectID, t.AssignedToID)

	if err != nil {
		return nil, err
	}

	id, err := rows.LastInsertId()

	if err != nil {
		return nil, err
	}

	t.ID = id
	return t, nil

}
