// Code generated by ogen, DO NOT EDIT.

package api

import (
	"net/http"
	"net/url"

	"github.com/go-faster/errors"

	"github.com/ogen-go/ogen/conv"
	"github.com/ogen-go/ogen/middleware"
	"github.com/ogen-go/ogen/ogenerrors"
	"github.com/ogen-go/ogen/uri"
	"github.com/ogen-go/ogen/validate"
)

// DeleteBannerParams is parameters of deleteBanner operation.
type DeleteBannerParams struct {
	ID int64
}

func unpackDeleteBannerParams(packed middleware.Parameters) (params DeleteBannerParams) {
	{
		key := middleware.ParameterKey{
			Name: "id",
			In:   "path",
		}
		params.ID = packed[key].(int64)
	}
	return params
}

func decodeDeleteBannerParams(args [1]string, argsEscaped bool, r *http.Request) (params DeleteBannerParams, _ error) {
	// Decode path: id.
	if err := func() error {
		param := args[0]
		if argsEscaped {
			unescaped, err := url.PathUnescape(args[0])
			if err != nil {
				return errors.Wrap(err, "unescape path")
			}
			param = unescaped
		}
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "id",
				Value:   param,
				Style:   uri.PathStyleSimple,
				Explode: false,
			})

			if err := func() error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToInt64(val)
				if err != nil {
					return err
				}

				params.ID = c
				return nil
			}(); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "id",
			In:   "path",
			Err:  err,
		}
	}
	return params, nil
}

// GetBannerParams is parameters of getBanner operation.
type GetBannerParams struct {
	TagID           int64
	FeatureID       int64
	UseLastRevision OptBool
}

func unpackGetBannerParams(packed middleware.Parameters) (params GetBannerParams) {
	{
		key := middleware.ParameterKey{
			Name: "tag_id",
			In:   "query",
		}
		params.TagID = packed[key].(int64)
	}
	{
		key := middleware.ParameterKey{
			Name: "feature_id",
			In:   "query",
		}
		params.FeatureID = packed[key].(int64)
	}
	{
		key := middleware.ParameterKey{
			Name: "use_last_revision",
			In:   "query",
		}
		if v, ok := packed[key]; ok {
			params.UseLastRevision = v.(OptBool)
		}
	}
	return params
}

func decodeGetBannerParams(args [0]string, argsEscaped bool, r *http.Request) (params GetBannerParams, _ error) {
	q := uri.NewQueryDecoder(r.URL.Query())
	// Decode query: tag_id.
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "tag_id",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToInt64(val)
				if err != nil {
					return err
				}

				params.TagID = c
				return nil
			}); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "tag_id",
			In:   "query",
			Err:  err,
		}
	}
	// Decode query: feature_id.
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "feature_id",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToInt64(val)
				if err != nil {
					return err
				}

				params.FeatureID = c
				return nil
			}); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "feature_id",
			In:   "query",
			Err:  err,
		}
	}
	// Set default value for query: use_last_revision.
	{
		val := bool(false)
		params.UseLastRevision.SetTo(val)
	}
	// Decode query: use_last_revision.
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "use_last_revision",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				var paramsDotUseLastRevisionVal bool
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToBool(val)
					if err != nil {
						return err
					}

					paramsDotUseLastRevisionVal = c
					return nil
				}(); err != nil {
					return err
				}
				params.UseLastRevision.SetTo(paramsDotUseLastRevisionVal)
				return nil
			}); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "use_last_revision",
			In:   "query",
			Err:  err,
		}
	}
	return params, nil
}

// GetBannersParams is parameters of getBanners operation.
type GetBannersParams struct {
	FeatureID OptInt64
	TagID     OptInt64
	Limit     OptInt64
	Offset    OptInt64
}

func unpackGetBannersParams(packed middleware.Parameters) (params GetBannersParams) {
	{
		key := middleware.ParameterKey{
			Name: "feature_id",
			In:   "query",
		}
		if v, ok := packed[key]; ok {
			params.FeatureID = v.(OptInt64)
		}
	}
	{
		key := middleware.ParameterKey{
			Name: "tag_id",
			In:   "query",
		}
		if v, ok := packed[key]; ok {
			params.TagID = v.(OptInt64)
		}
	}
	{
		key := middleware.ParameterKey{
			Name: "limit",
			In:   "query",
		}
		if v, ok := packed[key]; ok {
			params.Limit = v.(OptInt64)
		}
	}
	{
		key := middleware.ParameterKey{
			Name: "offset",
			In:   "query",
		}
		if v, ok := packed[key]; ok {
			params.Offset = v.(OptInt64)
		}
	}
	return params
}

func decodeGetBannersParams(args [0]string, argsEscaped bool, r *http.Request) (params GetBannersParams, _ error) {
	q := uri.NewQueryDecoder(r.URL.Query())
	// Decode query: feature_id.
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "feature_id",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				var paramsDotFeatureIDVal int64
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToInt64(val)
					if err != nil {
						return err
					}

					paramsDotFeatureIDVal = c
					return nil
				}(); err != nil {
					return err
				}
				params.FeatureID.SetTo(paramsDotFeatureIDVal)
				return nil
			}); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "feature_id",
			In:   "query",
			Err:  err,
		}
	}
	// Decode query: tag_id.
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "tag_id",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				var paramsDotTagIDVal int64
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToInt64(val)
					if err != nil {
						return err
					}

					paramsDotTagIDVal = c
					return nil
				}(); err != nil {
					return err
				}
				params.TagID.SetTo(paramsDotTagIDVal)
				return nil
			}); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "tag_id",
			In:   "query",
			Err:  err,
		}
	}
	// Decode query: limit.
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "limit",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				var paramsDotLimitVal int64
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToInt64(val)
					if err != nil {
						return err
					}

					paramsDotLimitVal = c
					return nil
				}(); err != nil {
					return err
				}
				params.Limit.SetTo(paramsDotLimitVal)
				return nil
			}); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "limit",
			In:   "query",
			Err:  err,
		}
	}
	// Decode query: offset.
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "offset",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				var paramsDotOffsetVal int64
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToInt64(val)
					if err != nil {
						return err
					}

					paramsDotOffsetVal = c
					return nil
				}(); err != nil {
					return err
				}
				params.Offset.SetTo(paramsDotOffsetVal)
				return nil
			}); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "offset",
			In:   "query",
			Err:  err,
		}
	}
	return params, nil
}

// SetBannerParams is parameters of setBanner operation.
type SetBannerParams struct {
	ID int64
}

func unpackSetBannerParams(packed middleware.Parameters) (params SetBannerParams) {
	{
		key := middleware.ParameterKey{
			Name: "id",
			In:   "path",
		}
		params.ID = packed[key].(int64)
	}
	return params
}

func decodeSetBannerParams(args [1]string, argsEscaped bool, r *http.Request) (params SetBannerParams, _ error) {
	// Decode path: id.
	if err := func() error {
		param := args[0]
		if argsEscaped {
			unescaped, err := url.PathUnescape(args[0])
			if err != nil {
				return errors.Wrap(err, "unescape path")
			}
			param = unescaped
		}
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "id",
				Value:   param,
				Style:   uri.PathStyleSimple,
				Explode: false,
			})

			if err := func() error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToInt64(val)
				if err != nil {
					return err
				}

				params.ID = c
				return nil
			}(); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "id",
			In:   "path",
			Err:  err,
		}
	}
	return params, nil
}
