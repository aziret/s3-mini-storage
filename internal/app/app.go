package app

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"

	"github.com/aziret/s3-mini-internal/pkg/api/filetransfer_v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"

	"github.com/aziret/s3-mini-storage/internal/config"
)

type App struct {
	serviceProvider *serviceProvider
	grpcServer      *grpc.Server
}

func NewApp(ctx context.Context) (*App, error) {
	app := &App{}

	err := app.initDeps(ctx)
	if err != nil {
		return nil, err
	}

	return app, nil
}

func (a *App) Run() error {
	return a.runGRPCServer()
}

func (a *App) initDeps(ctx context.Context) error {
	inits := []func(context.Context) error{
		a.initConfig,
		a.initServiceProvider,
		a.initGRPCServer,
		a.initUploadFolder,
		a.initGRPCClient,
	}

	for _, f := range inits {
		err := f(ctx)
		if err != nil {
			return err
		}
	}

	return nil
}

func (a *App) initConfig(_ context.Context) error {
	err := config.Load(".env")
	if err != nil {
		return err
	}

	return nil
}

func (a *App) initServiceProvider(_ context.Context) error {
	a.serviceProvider = newServiceProvider()
	return nil
}

func (a *App) initGRPCServer(_ context.Context) error {
	a.grpcServer = grpc.NewServer(grpc.Creds(insecure.NewCredentials()))

	reflection.Register(a.grpcServer)

	filetransfer_v1.RegisterFileTransferServiceV1Server(a.grpcServer, a.serviceProvider.FileTransferImpl())

	return nil
}

func (a *App) initUploadFolder(_ context.Context) error {
	dir := os.Getenv("FILE_PATH")

	if _, err := os.Stat(dir); os.IsNotExist(err) {
		// Create the uploads directory if it doesn't exist
		err := os.MkdirAll(dir, 0755)
		if err != nil {
			return fmt.Errorf("could not create directory: Dir name %s %w", dir, err)
		}
	}

	return nil
}

func (a *App) initGRPCClient(ctx context.Context) error {
	return a.serviceProvider.GRPCClient().RegisterClient(ctx)
}

func (a *App) runGRPCServer() error {
	log.Printf("GRPC server is running on %s", a.serviceProvider.GRPCConfig().Address())

	list, err := net.Listen("tcp", a.serviceProvider.GRPCConfig().Address())
	if err != nil {
		return err
	}

	err = a.grpcServer.Serve(list)
	if err != nil {
		return err
	}

	return nil
}
