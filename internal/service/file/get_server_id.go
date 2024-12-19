package file

import "context"

func (s *Service) GetServerID(ctx context.Context) (string, error) {
	return s.fileRepo.GetServerID(ctx)
}
