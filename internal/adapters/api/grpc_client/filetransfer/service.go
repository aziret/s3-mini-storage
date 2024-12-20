package filetransfer

import (
	"log/slog"
	"os"

	"google.golang.org/grpc/credentials/insecure"

	"github.com/aziret/s3-mini-internal/pkg/api/filetransfer_v1"
	"github.com/aziret/s3-mini-storage/internal/lib/logger/sl"
	"github.com/aziret/s3-mini-storage/internal/service"
	"google.golang.org/grpc"
)

type Implementation struct {
	logger             *slog.Logger
	filetransferClient filetransfer_v1.FileTransferServiceV1Client
	fileService        service.FileService
}

func NewImplementation(logger *slog.Logger, fileService service.FileService) *Implementation {
	const op = "grpc_client.filetransfer.NewImplementation"
	log := logger.With(
		slog.String("op", op),
	)

	mainServerAddress := os.Getenv("MAIN_SERVER_ADDRESS")

	conn, err := grpc.NewClient(mainServerAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Error("Failed to connect to server", sl.Err(err))

		panic(err)
	}

	client := filetransfer_v1.NewFileTransferServiceV1Client(conn)

	return &Implementation{
		logger:             logger,
		filetransferClient: client,
		fileService:        fileService,
	}
}
