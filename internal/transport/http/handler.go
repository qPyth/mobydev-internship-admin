package http

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/qPyth/mobydev-internship-admin/internal/domain"
	"log/slog"
	"net/http"
)

type ProjectManager interface {
	UpdateProject(ctx context.Context, project domain.Project) error
}

type UserProvider interface {
	LoginAdmin(ctx context.Context, email, password string) (token string, err error)
	RegisterAdmin(ctx context.Context, email, password string) error
}

type Handler struct {
	projectManager ProjectManager
	usrProvider    UserProvider
}

func NewHandler(projectService ProjectManager, provider UserProvider) *Handler {
	return &Handler{
		projectManager: projectService,
		usrProvider:    provider,
	}
}

func (h *Handler) Routes() chi.Router {
	r := chi.NewRouter()

	r.Post("/admin/register", h.RegisterAdmin)
	r.Post("/admin/login", h.LoginAdmin)
	r.With(h.IsAdminMW).Post("/project/update", h.UpdateProject)
	return r
}

type RegisterRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	PassConf string `json:"pass_conf"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (h *Handler) RegisterAdmin(w http.ResponseWriter, r *http.Request) {
	var req RegisterRequest
	ctx := r.Context()
	if err := h.bindJSON(r, &req); err != nil {
		slog.Error("failed to decode request body", "error", err)
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	err := RegisterValidation(req.Email, req.Password, req.PassConf)
	if err != nil {
		slog.Warn("invalid request", "error", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := h.usrProvider.RegisterAdmin(ctx, req.Email, req.Password); err != nil {
		if errors.Is(err, domain.ErrUserExists) {
			slog.Warn("user already exists", "email", req.Email)
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		slog.Error("failed to register admin", "error", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
}

func (h *Handler) LoginAdmin(w http.ResponseWriter, r *http.Request) {
	var req domain.User
	ctx := r.Context()
	if err := h.bindJSON(r, &req); err != nil {
		slog.Error("failed to decode request body", "error", err)
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	err := LoginValidation(req.Email, req.Password)
	if err != nil {
		slog.Warn("invalid request", "error", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	token, err := h.usrProvider.LoginAdmin(ctx, req.Email, req.Password)
	if err != nil {
		if errors.Is(err, domain.ErrInvalidCredentials) {
			slog.Warn("invalid credentials", "email", req.Email)
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		slog.Error("failed to login admin", "error", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(map[string]string{"token": token})
	if err != nil {
		slog.Error("failed to encode response", "error", err)
	}
}

func (h *Handler) UpdateProject(w http.ResponseWriter, r *http.Request) {
	var req domain.Project
	ctx := r.Context()
	if err := h.bindJSON(r, &req); err != nil {
		slog.Error("failed to decode request body", "error", err)
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	if err := req.Validate(); err != nil {
		slog.Error("invalid request", "error", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := h.projectManager.UpdateProject(ctx, req); err != nil {
		if errors.Is(err, domain.ProjectNotFound) || errors.Is(err, domain.ErrProjectIds) {
			slog.Warn("project not found", "id", req.ID)
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		slog.Error("failed to update project", "error", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
}

func (h *Handler) bindJSON(r *http.Request, data interface{}) error {
	err := json.NewDecoder(r.Body).Decode(data)
	if err != nil {
		return fmt.Errorf("failed to decode request body: %w", err)
	}
	return nil
}

func RegisterValidation(email, password, passConf string) error {
	if email == "" {
		return errors.New("email is required")
	}
	if password == "" {
		return errors.New("password is required")
	}
	if password != passConf {
		return errors.New("passwords do not match")
	}
	return nil
}

func LoginValidation(email, password string) error {
	if email == "" {
		return errors.New("email is required")
	}
	if password == "" {
		return errors.New("password is required")
	}
	return nil
}
