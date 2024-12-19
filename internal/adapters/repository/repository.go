package repository

import "context"

type FileRepository interface {
	GetServerID(ctx context.Context) (string, error)
}
