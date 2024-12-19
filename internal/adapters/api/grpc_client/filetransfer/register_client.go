package filetransfer

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/aziret/s3-mini-internal/pkg/api/filetransfer_v1"
	"github.com/aziret/s3-mini-storage/internal/lib/logger/sl"
)

func (i *Implementation) RegisterClient(ctx context.Context) error {
	const op = "grpc_client.filetransfer.RegisterClient"

	log := i.logger.With(
		slog.String("op", op),
	)

	serverID, err := i.fileService.GetServerID(ctx)

	if err != nil {
		log.Error("failed to get server ID")

		return fmt.Errorf("%s: %w", op, err)
	}
	req := &filetransfer_v1.PingRequest{
		Uuid: serverID,
	}

	resp, err := i.filetransferClient.RegisterClient(ctx, req)

	if err != nil {
		log.Error("failed to register client ID", sl.Err(err))

		return fmt.Errorf("%s: %w", op, err)
	}

	if !resp.GetSuccess() {
		errorMessage := "failed to register client ID without error"
		log.Error(errorMessage)

		return fmt.Errorf("%s: %s", errorMessage, resp.GetMessage())
	}

	return nil
}
