package domain

import "errors"

var (
	ErrInvalidProjectName        = errors.New("invalid project name")
	ErrInvalidProjectCategory    = errors.New("invalid project category")
	ErrInvalidProjectType        = errors.New("invalid project type")
	ErrInvalidAgeCategory        = errors.New("invalid age category")
	ErrInvalidProjectYear        = errors.New("invalid project year")
	ErrInvalidProjectDuration    = errors.New("invalid project duration")
	ErrInvalidProjectKeyWords    = errors.New("invalid project key words")
	ErrInvalidProjectDescription = errors.New("invalid project description")
	ErrInvalidProjectDirector    = errors.New("invalid project director")
	ErrInvalidProjectProducer    = errors.New("invalid project producer")
	ErrInvalidProjectID          = errors.New("invalid project id")
)

type Project struct {
	ID            int    `json:"id"`
	Name          string `json:"name"`
	CategoryID    int    `json:"category"`
	ProjectTypeID int    `json:"project_type"`
	AgeCategoryID int    `json:"age_category"`
	Year          string `json:"year"`
	Duration      string `json:"duration"`
	KeyWords      string `json:"key_words"`
	Description   string `json:"description"`
	Director      string `json:"director"`
	Producer      string `json:"producer"`
}

func (p *Project) Validate() error {
	if p.ID == 0 {
		return ErrInvalidProjectID
	}
	if p.Name == "" {
		return ErrInvalidProjectName
	}
	if p.CategoryID == 0 {
		return ErrInvalidProjectCategory
	}
	if p.ProjectTypeID == 0 {
		return ErrInvalidProjectType
	}
	if p.AgeCategoryID == 0 {
		return ErrInvalidAgeCategory
	}
	if p.Year == "" {
		return ErrInvalidProjectYear
	}
	if p.Duration == "" {
		return ErrInvalidProjectDuration
	}
	if p.KeyWords == "" {
		return ErrInvalidProjectKeyWords
	}
	if p.Description == "" {
		return ErrInvalidProjectDescription
	}
	if p.Director == "" {
		return ErrInvalidProjectDirector
	}
	if p.Producer == "" {
		return ErrInvalidProjectProducer
	}
	return nil
}
