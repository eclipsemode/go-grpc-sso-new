package services

import (
	"context"
	"github.com/eclipsemode/go-grpc-sso-new/internal/domain/model"
	"github.com/gofrs/uuid"
	"go.uber.org/zap"
	"time"
)

type Auth struct {
	log      *zap.SugaredLogger
	userRepo UserRepo
	tokenTTL time.Duration
}

type UserRepo interface {
	SaveUser(ctx context.Context, user *model.RegisterUserRequest) error
	GetUser(ctx context.Context, userID uuid.UUID) (*model.User, error)
	IsAdmin(ctx context.Context, userID uuid.UUID) (bool, error)
}

func New(log *zap.SugaredLogger, tokenTTL time.Duration, userRepo UserRepo) *Auth {
	return &Auth{
		log:      log,
		userRepo: userRepo,
		tokenTTL: tokenTTL,
	}
}

func (a *Auth) Login(ctx context.Context, request *model.LoginUserRequest) (token string, err error) {
	panic("implement me")
}

func (a *Auth) RegisterNewUser(ctx context.Context, request *model.RegisterUserRequest) error {
	panic("implement me")
}

func (a *Auth) IsAdmin(ctx context.Context, userID uuid.UUID) (bool, error) {
	panic("implement me")
}
