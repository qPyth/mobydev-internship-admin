package app

import (
	"errors"
	"github.com/qPyth/mobydev-internship-admin/internal/config"
	services2 "github.com/qPyth/mobydev-internship-admin/internal/services"
	"github.com/qPyth/mobydev-internship-admin/internal/storage"
	"github.com/qPyth/mobydev-internship-admin/internal/storage/sqlite"
	h "github.com/qPyth/mobydev-internship-admin/internal/transport/http"
	"github.com/qPyth/mobydev-internship-admin/pkg/auth"
	"log/slog"
	"net/http"
	"os"
)

func Run(cfg *config.Config) {
	// init storage
	dbConn := sqlite.New(cfg.StoragePath)
	storages := storage.NewManager(dbConn)
	defer dbConn.Close()
	//init Deps
	tokenManager := auth.NewJWTManager(cfg.TokenTTL, os.Getenv("JWT_SECRET"))

	// Create services
	UserService := services2.NewUser(tokenManager, storages.UserStorage)
	projectService := services2.NewProject(storages.ProjectStorage)

	handler := h.NewHandler(projectService, UserService)

	srv := &http.Server{
		Addr:         cfg.HTTP.Host + ":" + cfg.HTTP.Port,
		Handler:      handler.Routes(),
		ReadTimeout:  cfg.HTTP.ReadTimeout,
		WriteTimeout: cfg.HTTP.WriteTimeout,
	}

	slog.Info("server started", "address", srv.Addr)
	if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		slog.Error("server failed to start", "error", err)
		panic(err)
	}

}
