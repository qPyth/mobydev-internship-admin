package services

import (
	"context"
	"github.com/qPyth/mobydev-internship-admin/internal/domain"
)

type ProjectStorage interface {
	Update(ctx context.Context, project domain.Project) error
}

type ProjectService struct {
	projectStorage ProjectStorage
}

func NewProject(storage ProjectStorage) *ProjectService {
	return &ProjectService{projectStorage: storage}
}

func (s *ProjectService) UpdateProject(ctx context.Context, project domain.Project) error {
	return s.projectStorage.Update(ctx, project)
}
