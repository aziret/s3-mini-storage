package app

import (
	"github.com/aziret/s3-mini-storage/internal/adapters/api/grpc_client/filetransfer"
	"github.com/aziret/s3-mini-storage/internal/adapters/repository"
	"github.com/aziret/s3-mini-storage/internal/config"
	"github.com/aziret/s3-mini-storage/internal/lib/logger/sl"
	"github.com/aziret/s3-mini-storage/internal/service"

	"log"
	"log/slog"

	filetransferServer "github.com/aziret/s3-mini-storage/internal/adapters/api/grpc_server/filetransfer"
	fileRepository "github.com/aziret/s3-mini-storage/internal/adapters/repository/file"
	fileService "github.com/aziret/s3-mini-storage/internal/service/file"
)

type serviceProvider struct {
	log                    *slog.Logger
	fileRepo               repository.FileRepository
	fileService            service.FileService
	fileTransferServerImpl *filetransferServer.Implementation
	grpcConfig             config.GRPCConfig
	grpcClient             *filetransfer.Implementation
}

func newServiceProvider() *serviceProvider {
	return &serviceProvider{}
}

func (s *serviceProvider) GRPCConfig() config.GRPCConfig {
	if s.grpcConfig == nil {
		cfg, err := config.NewGRPCConfig()
		if err != nil {
			log.Fatalf("failed to get grpc config: %s", err.Error())
		}

		s.grpcConfig = cfg
	}

	return s.grpcConfig
}

func (s *serviceProvider) FileTransferImpl() *filetransferServer.Implementation {
	if s.fileTransferServerImpl == nil {
		s.fileTransferServerImpl = filetransferServer.NewImplementation(s.FileService(), s.Logger())
	}

	return s.fileTransferServerImpl
}

func (s *serviceProvider) FileService() service.FileService {
	if s.fileService == nil {
		s.fileService = fileService.NewService(s.FileRepo(), s.Logger())
	}

	return s.fileService
}

func (s *serviceProvider) FileRepo() repository.FileRepository {
	const op = "serviceProvider.FileRepo"

	logger := s.Logger()

	lgr := logger.With(
		slog.String("op", op),
	)

	if s.fileRepo == nil {
		repo, err := fileRepository.NewRepository(logger)
		if err != nil {
			lgr.Error("failed to initialize repo", sl.Err(err))
		}

		s.fileRepo = repo
	}

	return s.fileRepo
}

func (s *serviceProvider) Logger() *slog.Logger {
	if s.log == nil {
		s.log = config.NewLogger()
	}

	return s.log
}

func (s *serviceProvider) GRPCClient() *filetransfer.Implementation {
	if s.grpcClient == nil {
		s.grpcClient = filetransfer.NewImplementation(s.Logger(), s.FileService())
	}

	return s.grpcClient
}
