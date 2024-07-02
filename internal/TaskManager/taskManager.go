package TaskManager

import (
	"context"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/wlcmtunknwndth/effective_mobile_test/internal/config"
	"github.com/wlcmtunknwndth/effective_mobile_test/internal/domain/models"
	"github.com/wlcmtunknwndth/effective_mobile_test/lib/sl"
	"log/slog"
	"net/http"
)

type UsersStorage interface {
	GetUserByPassport(ctx context.Context, passportSerie uint16, passportNumber uint32) (*models.User, error)

	CreateUser(ctx context.Context, user *models.User) (uint64, error)
	AddUserInfo(ctx context.Context, info *models.UserInfo) error

	GetUser(ctx context.Context, id uint64) (*models.User, error)
	DeleteUser(ctx context.Context, id uint64) error
	UpdateUser(ctx context.Context, user *models.User) error
	IsAdmin(ctx context.Context, id uint64) (bool, error)
}

type TasksStorage interface {
	CreateTask(ctx context.Context, task *models.Task) error
	GetWorkload(ctx context.Context, passportNumber string) ([]models.Task, error)
}

const scope = "internal.TaskManager."

type Service struct {
	server *http.Server
	users  UsersStorage
	tasks  TasksStorage
	log    *slog.Logger
}

func New(log *slog.Logger, users UsersStorage, tasks TasksStorage, cfg *config.Server) *Service {
	router := chi.NewRouter()
	router.Use(middleware.RequestID)
	router.Use(middleware.URLFormat)
	router.Use(middleware.Logger)

	srv := &http.Server{
		Addr:         cfg.Address,
		Handler:      router,
		ReadTimeout:  cfg.Timeout,
		WriteTimeout: cfg.Timeout,
		IdleTimeout:  cfg.IdleTimeout,
	}

	service := &Service{users: users, tasks: tasks, log: log, server: srv}

	return service
}

// Start - Function causes lock on current goroutine
func (s *Service) Start() error {
	const op = scope + "Start"

	if err := s.server.ListenAndServe(); err != nil {
		s.log.Error("couldn't run server", sl.Op(op), sl.Err(err))
		return fmt.Errorf("%s: %w", op, err)
	}
	return nil
}

func (s *Service) Stop() error {
	const op = scope + "Stop"

	if err := s.server.Close(); err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	return nil
}
