package nats

import (
	"context"
	"fmt"
	"github.com/nats-io/nats.go"
	"github.com/wlcmtunknwndth/effective_mobile_test/internal/config"
	"github.com/wlcmtunknwndth/effective_mobile_test/internal/domain/models"
	"time"
)

type UsersStorage interface {
	GetUserByPassport(ctx context.Context, passportNumber string) (*models.User, error)

	CreateUser(ctx context.Context, user *models.User) error
	GetUser(ctx context.Context, id uint64) (*models.User, error)
	DeleteUser(ctx context.Context, id uint64) error
	UpdateUser(ctx context.Context, user *models.User) error
	IsAdmin(ctx context.Context, id uint64) (bool, error)
}

type TasksStorage interface {
	CreateTask(ctx context.Context, user *models.Task) error
	GetTasksWorkload(ctx context.Context, passportNumber string) ([]models.Task, error)
}

type Broker struct {
	conn  *nats.Conn
	users UsersStorage
	tasks TasksStorage
}

const scope = "internal.broker.nats."

func New(cfg *config.Broker, users UsersStorage, tasks TasksStorage) (*Broker, error) {
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

	return &Broker{
		conn:  natsSrv,
		users: users,
		tasks: tasks,
	}, nil
}

func (b *Broker) Close() {
	b.conn.Close()
}
