package mongo

import (
	"context"
	"github.com/eclipsemode/go-grpc-sso-new/internal/domain/model"
	"github.com/gofrs/uuid"
)

func (s *Storage) SaveUser(ctx context.Context, user *model.RegisterUserRequest) error {
	panic("implement me")
}

func (s *Storage) GetUser(ctx context.Context, userID uuid.UUID) (*model.User, error) {
	panic("implement me")
}

func (s *Storage) IsAdmin(ctx context.Context, userID uuid.UUID) (bool, error) {
	panic("implement me")
}
