package file

import (
	"context"
	"fmt"
	"github.com/aziret/s3-mini-storage/internal/lib/logger/sl"
	"github.com/aziret/s3-mini-storage/internal/model"
	"log/slog"
	"os"
	"path/filepath"
)

func (s *Service) SaveFile(_ context.Context, file *model.File) error {
	const op = "service.file.SaveFile"

	log := s.logger.With(
		slog.String("op", op),
	)

	dir := os.Getenv("FILE_PATH")

	if _, err := os.Stat(dir); os.IsNotExist(err) {
		// Create the uploads directory if it doesn't exist
		err := os.MkdirAll(dir, 0755)
		if err != nil {
			log.Error("could not create directory", slog.String("path", dir), sl.Err(err))

			return fmt.Errorf("%s: %w", op, err)
		}
	}

	filePath := filepath.Join(dir, file.UUID)

	err := os.WriteFile(filePath, file.Data, 0644)
	if err != nil {
		log.Error("could not write file", slog.String("path", filePath), sl.Err(err))
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}
