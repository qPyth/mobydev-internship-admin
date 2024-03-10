package domain

import "errors"

var (
	ErrUserExists         = errors.New("user already exists")
	ProjectNotFound       = errors.New("project not found")
	ErrInvalidCredentials = errors.New("invalid credentials")
	ErrUserNotFound       = errors.New("user not found")
	ErrProjectIds         = errors.New("please check category, project type or age category ids")
)
