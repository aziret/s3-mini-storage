package file

import (
	"context"
	"fmt"
	"log/slog"
	"os"

	"github.com/aziret/s3-mini-storage/internal/lib/logger/sl"
	"github.com/google/uuid"
)

func (s *Service) GetFileData(ctx context.Context, UUID string) ([]byte, error) {
	const op = "service.file.GetFileData"
	log := s.logger.With(
		slog.String("op", op),
		slog.String("uuid", UUID),
	)

	parsedUUID, err := uuid.Parse(UUID)
	if err != nil {
		log.Error("Unable to get file info", sl.Err(err))
		return []byte{}, fmt.Errorf("incorrect UUID sent")
	}

	fileInfo, err := s.fileRepo.GetFile(ctx, parsedUUID)
	if err != nil {
		log.Error("Unable to get file info", sl.Err(err))
	}

	data, err := os.ReadFile(fileInfo.FilePath)
	if err != nil {
		log.Error("failed to read file", sl.Err(err))
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	return data, nil
}
