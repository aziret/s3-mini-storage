package file

import (
	"database/sql"
	"fmt"
	"log/slog"
	"os"

	"github.com/aziret/s3-mini-storage/internal/lib/logger/sl"
)

type repository struct {
	db  *sql.DB
	log *slog.Logger
}

func NewRepository(log *slog.Logger) (*repository, error) {
	const op = "repository.file.NewRepository"

	logger := log.With(
		slog.String("op", op),
	)

	db, err := sql.Open("postgres", connectionString())
	if err != nil {
		logger.Error("failed to open database connection", sl.Err(err))

		return nil, fmt.Errorf("%s: %w", op, err)
	}

	err = db.Ping()
	if err != nil {
		logger.Error("Error pinging database: ", sl.Err(err))

		return nil, fmt.Errorf("%s: %w", op, err)
	}

	logger.Info("Successfully connected to database")

	return &repository{
		db:  db,
		log: log,
	}, nil
}

func connectionString() string {
	var (
		pgUser    = os.Getenv("PG_USER")
		pgPass    = os.Getenv("PG_PASS")
		pgHost    = os.Getenv("PG_HOST")
		pgPort    = os.Getenv("PG_PORT")
		pgDB      = os.Getenv("PG_DB")
		pgSSLMode = os.Getenv("PG_SSL_MODE")
	)
	return fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s sslmode=%s", pgUser, pgPass, pgDB, pgHost, pgPort, pgSSLMode)
}
