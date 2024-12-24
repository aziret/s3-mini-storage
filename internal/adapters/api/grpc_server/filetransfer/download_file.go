package filetransfer

import (
	"fmt"
	"io"
	"log/slog"

	"github.com/aziret/s3-mini-internal/pkg/api/filetransfer_v1"
	"github.com/aziret/s3-mini-storage/internal/lib/logger/sl"
	"google.golang.org/grpc"
)

func (i *Implementation) DownloadFile(stream grpc.BidiStreamingServer[filetransfer_v1.FileChunkRequest, filetransfer_v1.FileChunkDownload]) error {
	const op = "grpcServer.fileTransfer.DownloadFile"
	log := i.logger.With(
		slog.String("op", op),
	)

	ctx := stream.Context()

	for {
		req, err := stream.Recv()

		if err != nil {
			if err == io.EOF {
				log.Info("streaming closed")
				break
			}
			log.Error("failed to receive a request", sl.Err(err))
			return fmt.Errorf("%s: %w", op, err)
		}

		data, err := i.fileService.GetFileData(ctx, req.GetUuid())
		if err != nil {
			log.Error("failed to get file data", sl.Err(err))
			return fmt.Errorf("failed to get file data: %w", err)
		}

		resp := &filetransfer_v1.FileChunkDownload{
			Uuid:        req.GetUuid(),
			Data:        data,
			ChunkSize:   req.GetChunkSize(),
			ChunkNumber: req.GetChunkNumber(),
		}

		if err := stream.Send(resp); err != nil {
			return fmt.Errorf("failed to send: %v", err)
		}
	}

	return nil
}
