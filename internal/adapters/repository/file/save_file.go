package file

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/aziret/s3-mini-storage/internal/lib/logger/sl"
	"github.com/lib/pq"

	repoPackage "github.com/aziret/s3-mini-storage/internal/adapters/repository"
	"github.com/aziret/s3-mini-storage/internal/model"
)

func (repo *repository) SaveFile(ctx context.Context, info *model.FileInfo) error {
	const op = "repository.file.SaveFile"

	log := repo.log.With(
		slog.String("op", op),
	)

	stmt, err := repo.db.Prepare("INSERT INTO files (id, file_path) VALUES ($1, $2) ON CONFLICT DO NOTHING")

	if err != nil {
		log.Error("failed to prepare statement", sl.Err(err))
		return fmt.Errorf("%s: %w", op, err)
	}

	_, err = stmt.Exec(info.UUID, info.FilePath)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok && pqErr.Code.Name() == "unique_violation" {
			return fmt.Errorf("%s: %w", op, repoPackage.ErrFileExists)
		}

		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}
