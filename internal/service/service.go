package service

import "context"

type FileService interface {
	GetServerID(ctx context.Context) (string, error)
}
