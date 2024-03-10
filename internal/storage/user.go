package storage

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/mattn/go-sqlite3"
	"github.com/qPyth/mobydev-internship-admin/internal/domain"
)

type UserStorage struct {
	conn *sql.DB
}

func NewUserStorage(connection *sql.DB) *UserStorage {
	return &UserStorage{conn: connection}
}

func CreateUser(ctx context.Context, email string, passHash []byte) error {
	return nil
}

func (s *UserStorage) CreateUser(ctx context.Context, email string, passHash []byte, role string) error {
	op := "UserStorage.CreateUser"
	stmt, err := s.conn.PrepareContext(ctx, `INSERT INTO users(email, password, role) VALUES(?, ?, ?)`)
	if err != nil {
		return fmt.Errorf("%s: failed to prepare statement: %w", op, err)
	}

	_, err = stmt.ExecContext(ctx, email, string(passHash), role)
	if err != nil {
		var sqliteErr *sqlite3.Error
		if errors.As(err, &sqliteErr) && errors.Is(sqliteErr.Code, sqlite3.ErrConstraint) {
			return domain.ErrUserExists
		}
		return fmt.Errorf("%s: failed to execute statement: %w", op, err)
	}
	return nil
}

func (s *UserStorage) GetUser(ctx context.Context, email string) (domain.User, error) {
	op := "UserStorage.GetUser"
	var user domain.User
	stmt, err := s.conn.PrepareContext(ctx, `SELECT id, email, password, role FROM users WHERE email = ?`)
	if err != nil {
		return user, fmt.Errorf("%s: failed to prepare statement: %w", op, err)
	}
	row := stmt.QueryRowContext(ctx, email)
	err = row.Scan(&user.ID, &user.Email, &user.Password, &user.Role)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return user, domain.ErrUserNotFound
		}
		return user, fmt.Errorf("%s: failed to scan row: %w", op, err)
	}
	return user, nil
}
