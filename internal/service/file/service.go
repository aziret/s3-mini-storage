package file

import (
	"log/slog"

	"github.com/aziret/s3-mini-storage/internal/adapters/repository"
)

type Service struct {
	logger   *slog.Logger
	fileRepo repository.FileRepository
}

func NewService(fileRepo repository.FileRepository, logger *slog.Logger) *Service {
	return &Service{
		fileRepo: fileRepo,
		logger:   logger,
	}
}
