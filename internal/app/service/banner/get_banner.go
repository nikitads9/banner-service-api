package banner

import (
	"context"

	"github.com/go-faster/jx"
)

func (s *Service) GetBanner(ctx context.Context, featureID int64, tagID int64, useLastRevision bool) (map[string]jx.Raw, error) {
	return s.postgresRepository.GetBanner(ctx, featureID, tagID)
}
