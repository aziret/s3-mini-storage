package service

import (
	"context"
	"github.com/aziret/s3-mini-storage/internal/model"
)

type FileService interface {
	GetServerID(ctx context.Context) (string, error)
	SaveFile(ctx context.Context, file *model.File) error
	GetFileData(ctx context.Context, UUID string) ([]byte, error)
}
