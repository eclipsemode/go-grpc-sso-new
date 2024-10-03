package grpc

import (
	"fmt"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"net"
)

type App struct {
	log        *zap.SugaredLogger
	gRPCServer *grpc.Server
	port       int
}

func New(log *zap.SugaredLogger, port int) *App {
	gRPCServer := grpc.NewServer()

	return &App{
		log:        log,
		gRPCServer: gRPCServer,
		port:       port,
	}
}

func (a *App) MustRun() {
	if err := a.Run(); err != nil {
		panic(err)
	}
}

func (a *App) Run() error {
	const op = "grpc.Run"

	log := a.log.With(zap.String("operation", op), zap.Int("port", a.port))

	l, err := net.Listen("tcp", fmt.Sprintf(":%d", a.port))
	if err != nil {
		log.Errorf("%s: failed to listen: %v", op, err)
		return fmt.Errorf("%s: %w", op, err)
	}

	log.Info("starting gRPC server on port: ", zap.String("addr", l.Addr().String()))

	if err := a.gRPCServer.Serve(l); err != nil {
		log.Errorf("%s: failed to serve: %v", op, err)
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func (a *App) Stop() error {
	const op = "grpc.Stop"

	a.log.With("operation", op).
		Info("stopping gRPC server")

	a.gRPCServer.Stop()

	return nil
}
