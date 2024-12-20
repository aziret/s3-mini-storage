package filetransfer

import (
	"fmt"
	"github.com/aziret/s3-mini-storage/internal/converter"
	"io"
	"log/slog"

	"github.com/aziret/s3-mini-internal/pkg/api/filetransfer_v1"
	"github.com/aziret/s3-mini-storage/internal/lib/logger/sl"
	"google.golang.org/grpc"
)

func (i *Implementation) UploadFile(stream grpc.ClientStreamingServer[filetransfer_v1.FileChunk, filetransfer_v1.UploadStatus]) error {
	const op = "grpcServer.filetransfer.UploadFile"

	log := i.logger.With(
		slog.String("op", op),
	)
	for {
		req, err := stream.Recv()

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Error("Error receiving stream", sl.Err(err))
			return fmt.Errorf("%s: %w", op, err)
		}

		err = i.fileService.SaveFile(stream.Context(), converter.ToFileFromApi(req))

		if err != nil {
			log.Error("Error saving file", sl.Err(err))
			return fmt.Errorf("error saving file: %w", err)
		}
	}

	resp := filetransfer_v1.UploadStatus{
		Message: "Files uploaded successfully",
		Success: true,
	}

	err := stream.SendAndClose(&resp)
	if err != nil {
		return fmt.Errorf("error closing stream: %w", err)
	}
	return nil
}
