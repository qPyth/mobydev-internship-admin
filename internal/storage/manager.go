package storage

import "database/sql"

type Manager struct {
	ProjectStorage *ProjectStorage
	UserStorage    *UserStorage
}

func NewManager(connection *sql.DB) *Manager {
	return &Manager{
		ProjectStorage: NewProjectStorage(connection),
		UserStorage:    NewUserStorage(connection),
	}
}
