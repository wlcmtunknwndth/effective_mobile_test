package receiver

import (
	"context"
	"fmt"
	"github.com/nats-io/nats.go"
	"github.com/wlcmtunknwndth/effective_mobile_test/internal/config"
	"github.com/wlcmtunknwndth/effective_mobile_test/internal/domain/models"
	"log/slog"
	"time"
)

type UsersStorage interface {
	GetUserByPassport(ctx context.Context, passportSerie uint16, passportNumber uint32) (*models.User, error)

	CreateUser(ctx context.Context, user *models.User) (uint64, error)
	//CreateUserInf(ctx)
	GetUser(ctx context.Context, id uint64) (*models.User, error)
	DeleteUser(ctx context.Context, id uint64) error
	UpdateUser(ctx context.Context, user *models.User) error
	IsAdmin(ctx context.Context, id uint64) (bool, error)
}

type TasksStorage interface {
	CreateTask(ctx context.Context, task *models.Task) error
	GetWorkload(ctx context.Context, passportNumber string) ([]models.Task, error)
}

type Receiver struct {
	conn         *nats.Conn
	users        UsersStorage
	tasks        TasksStorage
	log          *slog.Logger
	usersHandler *nats.Subscription
	tasksHandler *nats.Subscription
}

const scope = "internal.broker.nats."

func New(cfg *config.Broker, users UsersStorage, tasks TasksStorage, logger *slog.Logger) (*Receiver, error) {
	const op = scope + "New"
	natsSrv, err := nats.Connect(cfg.Address,
		nats.RetryOnFailedConnect(cfg.Retry),
		nats.MaxReconnects(cfg.MaxReconnects),
		nats.ReconnectWait(cfg.ReconnectWait),
	)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	if err = natsSrv.FlushTimeout(3 * time.Second); err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return &Receiver{
		conn:  natsSrv,
		users: users,
		tasks: tasks,
		log:   logger,
	}, nil
}

func (b *Receiver) Close() {
	b.conn.Close()
}
