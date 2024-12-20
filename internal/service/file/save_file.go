package file

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"path/filepath"

	"github.com/aziret/s3-mini-storage/internal/lib/logger/sl"
	"github.com/aziret/s3-mini-storage/internal/model"
)

func (s *Service) SaveFile(_ context.Context, file *model.File) error {
	const op = "service.file.SaveFile"

	log := s.logger.With(
		slog.String("op", op),
	)

	dir := os.Getenv("FILE_PATH")

	filePath := filepath.Join(dir, file.UUID)

	err := os.WriteFile(filePath, file.Data, 0644)
	if err != nil {
		log.Error("could not write file", slog.String("path", filePath), sl.Err(err))
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}
