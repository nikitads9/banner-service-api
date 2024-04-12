package banner

import (
	"context"
)

func (s *Service) DeleteBanner(ctx context.Context, bannerID int64) error {
	return s.postgresRepository.DeleteBanner(ctx, bannerID)
}
