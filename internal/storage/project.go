package storage

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/qPyth/mobydev-internship-admin/internal/domain"
)

type ProjectStorage struct {
	conn *sql.DB
}

func NewProjectStorage(connection *sql.DB) *ProjectStorage {
	return &ProjectStorage{conn: connection}
}

func (s *ProjectStorage) Update(ctx context.Context, project domain.Project) error {

	if _, err := s.conn.Exec("PRAGMA foreign_keys = ON", nil); err != nil {
		return fmt.Errorf("failed to enable foreign keys: %w", err)
	}

	stmt, err := s.conn.PrepareContext(ctx, "UPDATE projects SET name = ?, category_id = ?, project_type_id = ?, age_category_id = ?, year = ?, duration = ?, key_words = ?, description = ?, director = ?, producer = ? WHERE id = ?")
	if err != nil {
		return fmt.Errorf("failed to prepare update statement: %w", err)
	}
	res, err := stmt.ExecContext(ctx, project.Name, project.CategoryID, project.ProjectTypeID, project.AgeCategoryID, project.Year, project.Duration, project.KeyWords, project.Description, project.Director, project.Producer, project.ID)
	if err != nil {
		if err.Error() == "FOREIGN KEY constraint failed" {
			return domain.ErrProjectIds
		}
		return fmt.Errorf("failed to update project: %w", err)
	}
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %w", err)
	}

	if rowsAffected == 0 {
		return domain.ProjectNotFound
	}
	return nil
}
