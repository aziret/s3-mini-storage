package filetransfer

import (
	"log/slog"

	"github.com/aziret/s3-mini-internal/pkg/api/filetransfer_v1"
	"github.com/aziret/s3-mini-storage/internal/service"
)

type Implementation struct {
	filetransfer_v1.UnimplementedFileTransferServiceV1Server
	fileService service.FileService
	logger      *slog.Logger
}

func NewImplementation(fileService service.FileService, logger *slog.Logger) *Implementation {
	return &Implementation{
		fileService: fileService,
		logger:      logger,
	}
}
