package file

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log/slog"

	"github.com/aziret/s3-mini-storage/internal/lib/logger/sl"
	"github.com/aziret/s3-mini-storage/internal/model"
)

func (repo *repository) GetFile(ctx context.Context, UUID string) (*model.FileInfo, error) {
	const op = "repository.file.GetFile"
	log := repo.log.With(
		slog.String("op", op),
		slog.String("UUID", UUID),
	)

	stmt, err := repo.db.Prepare("SELECT id, file_path FROM files WHERE uuid = $1 ")
	if err != nil {
		log.Error("failed to prepare statement", sl.Err(err))
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	defer func(stmt *sql.Stmt) {
		err := stmt.Close()
		if err != nil {
			log.Error("failed to close statement", sl.Err(err))
		}
	}(stmt)

	var fileInfo model.FileInfo

	err = stmt.QueryRow(UUID).Scan(&fileInfo.UUID, &fileInfo.FilePath)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			log.Error("no file by provided UUID", sl.Err(err))
			return nil, fmt.Errorf("%s: %w", op, err)
		} else {
			log.Error("Error querying file info:", sl.Err(err))
			return nil, fmt.Errorf("%s: %w", op, err)
		}
	}

	return &fileInfo, nil
}
