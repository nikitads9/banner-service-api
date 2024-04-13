package convert

import (
	"database/sql"

	"github.com/nikitads9/banner-service-api/internal/app/model"
	desc "github.com/nikitads9/banner-service-api/pkg/banner-api"
)

func ToBanner(req *desc.CreateBannerRequest) *model.Banner {
	return &model.Banner{
		FeatureID: req.GetFeatureID(),
		TagIDs:    req.GetTagIds(),
		Content:   req.GetContent(),
		IsActive:  req.GetIsActive(),
	}
}

func ToGetBannersResponse(mod []*model.BannerInfo) []desc.GetBannersResponseItem {
	res := make([]desc.GetBannersResponseItem, 0, len(mod))
	for _, val := range mod {
		item := desc.GetBannersResponseItem{
			BannerID:  val.BannerID,
			TagIds:    val.TagIDs,
			FeatureID: val.FeatureID,
			Content:   val.Content,
			IsActive:  val.IsActive,
			CreatedAt: val.CreatedAt,
		}

		if val.UpdatedAt.Valid {
			item.UpdatedAt = desc.OptDateTime{
				Value: item.UpdatedAt.Value,
				Set:   true,
			}
		}

		res = append(res, item)
	}

	return res
}

func ToSetBannerInfo(bannerID int64, req *desc.SetBannerRequest) *model.SetBannerInfo {
	res := &model.SetBannerInfo{
		BannerID: bannerID,
	}

	featureID, ok := req.GetFeatureID().Get()
	res.FeatureID = sql.NullInt64{
		Int64: featureID,
		Valid: ok,
	}

	tagIDs, ok := req.GetTagIds().Get()
	if ok && len(tagIDs) > 0 {
		res.TagIDs = tagIDs
	}

	//TODO: check
	res.Content = model.NullJxRaw{JxRaw: req.GetContent(),
		Valid: len(req.GetContent().String()) != 0}

	isActive, ok := req.GetIsActive().Get()
	res.IsActive = sql.NullBool{
		Bool:  isActive,
		Valid: ok,
	}

	return res
}

func ToGetBannersParams(params desc.GetBannersParams, limit int64, offset int64) *model.GetBannersParams {
	return &model.GetBannersParams{
		FeatureID: sql.NullInt64{
			Int64: params.FeatureID.Value,
			Valid: params.FeatureID.IsSet(),
		},
		TagID: sql.NullInt64{
			Int64: params.TagID.Value,
			Valid: params.TagID.IsSet(),
		},
		Limit:  limit,
		Offset: offset,
	}
}
