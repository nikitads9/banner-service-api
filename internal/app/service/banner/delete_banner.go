package banner

import (
	"context"
)

// DeleteBanner это удаление баннера по идентификатору.
func (s *Service) DeleteBanner(ctx context.Context, bannerID int64) error {
	return s.postgresRepository.DeleteBanner(ctx, bannerID)
}
