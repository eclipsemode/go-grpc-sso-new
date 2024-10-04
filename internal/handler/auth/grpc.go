package handler

import (
	"context"
	"github.com/eclipsemode/go-grpc-sso-new/internal/domain/model"
	sso "github.com/eclipsemode/go-grpc-sso-proto-new/protobuf/gen/proto"
	"github.com/gofrs/uuid"
	"google.golang.org/grpc"
)

type grpcAPI struct {
	svc AuthService
	sso.UnimplementedAuthServer
}

type AuthService interface {
	Login(ctx context.Context, request *model.LoginUserRequest) (token string, err error)
	RegisterNewUser(ctx context.Context, request *model.RegisterUserRequest) error
	IsAdmin(ctx context.Context, userID uuid.UUID) (bool, error)
}

func Register(gRPC *grpc.Server, svc AuthService) {
	sso.RegisterAuthServer(gRPC, &grpcAPI{svc: svc})
}

func (g *grpcAPI) Login(ctx context.Context, req *sso.LoginRequest) (*sso.LoginResponse, error) {
	panic("implement me")
}

func (g *grpcAPI) Register(ctx context.Context, req *sso.RegisterRequest) (*sso.RegisterResponse, error) {
	panic("implement me")
}

func (g *grpcAPI) IsAdmin(ctx context.Context, req *sso.IsAdminRequest) (*sso.IsAdminResponse, error) {
	panic("implement me")
}
