package file

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/aziret/s3-mini-storage/internal/lib/logger/sl"
	"github.com/pkg/errors"
	"log/slog"
)

func (repo *repository) GetServerID(ctx context.Context) (string, error) {
	const op = "repository.file.CreateFileChunksForFile"

	log := repo.log.With(
		slog.String("op", op),
	)
	var id string

	err := repo.db.QueryRow("SELECT id FROM server_info LIMIT 1").Scan(&id)
	if errors.Is(err, sql.ErrNoRows) {
		log.Info("server info does not exist, inserting one...")
		err = repo.db.QueryRow("INSERT INTO server_info DEFAULT VALUES RETURNING id").Scan(&id)
		if err != nil {
			log.Error("failed to insert UUID", sl.Err(err))

			return "", fmt.Errorf("%s: %w", op, err)
		}
	} else if err != nil {
		log.Error("failed to query UUID", sl.Err(err))
		return "", fmt.Errorf("%s: %w", op, err)
	}

	return id, nil
}
