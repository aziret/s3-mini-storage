package repository

import (
	"context"
	"errors"

	"github.com/aziret/s3-mini-storage/internal/model"
	"github.com/google/uuid"
)

var (
	ErrFileExists = errors.New("file already exists")
)

type FileRepository interface {
	GetServerID(ctx context.Context) (string, error)
	SaveFile(ctx context.Context, info *model.FileInfo) error
	GetFile(ctx context.Context, UUID uuid.UUID) (*model.FileInfo, error)
}
