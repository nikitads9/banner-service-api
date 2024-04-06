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

// BannerGetParams is parameters of GET /banner operation.
type BannerGetParams struct {
	FeatureID OptInt
	TagID     OptInt
	Limit     OptInt
	Offset    OptInt
}

func unpackBannerGetParams(packed middleware.Parameters) (params BannerGetParams) {
	{
		key := middleware.ParameterKey{
			Name: "feature_id",
			In:   "query",
		}
		if v, ok := packed[key]; ok {
			params.FeatureID = v.(OptInt)
		}
	}
	{
		key := middleware.ParameterKey{
			Name: "tag_id",
			In:   "query",
		}
		if v, ok := packed[key]; ok {
			params.TagID = v.(OptInt)
		}
	}
	{
		key := middleware.ParameterKey{
			Name: "limit",
			In:   "query",
		}
		if v, ok := packed[key]; ok {
			params.Limit = v.(OptInt)
		}
	}
	{
		key := middleware.ParameterKey{
			Name: "offset",
			In:   "query",
		}
		if v, ok := packed[key]; ok {
			params.Offset = v.(OptInt)
		}
	}
	return params
}

func decodeBannerGetParams(args [0]string, argsEscaped bool, r *http.Request) (params BannerGetParams, _ error) {
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
				var paramsDotFeatureIDVal int
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToInt(val)
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
				var paramsDotTagIDVal int
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToInt(val)
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
				var paramsDotLimitVal int
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToInt(val)
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
				var paramsDotOffsetVal int
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToInt(val)
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

// BannerIDDeleteParams is parameters of DELETE /banner/{id} operation.
type BannerIDDeleteParams struct {
	ID int
}

func unpackBannerIDDeleteParams(packed middleware.Parameters) (params BannerIDDeleteParams) {
	{
		key := middleware.ParameterKey{
			Name: "id",
			In:   "path",
		}
		params.ID = packed[key].(int)
	}
	return params
}

func decodeBannerIDDeleteParams(args [1]string, argsEscaped bool, r *http.Request) (params BannerIDDeleteParams, _ error) {
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

				c, err := conv.ToInt(val)
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

// BannerIDPatchParams is parameters of PATCH /banner/{id} operation.
type BannerIDPatchParams struct {
	ID int
}

func unpackBannerIDPatchParams(packed middleware.Parameters) (params BannerIDPatchParams) {
	{
		key := middleware.ParameterKey{
			Name: "id",
			In:   "path",
		}
		params.ID = packed[key].(int)
	}
	return params
}

func decodeBannerIDPatchParams(args [1]string, argsEscaped bool, r *http.Request) (params BannerIDPatchParams, _ error) {
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

				c, err := conv.ToInt(val)
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

// UserBannerGetParams is parameters of GET /user_banner operation.
type UserBannerGetParams struct {
	TagID           int
	FeatureID       int
	UseLastRevision OptBool
}

func unpackUserBannerGetParams(packed middleware.Parameters) (params UserBannerGetParams) {
	{
		key := middleware.ParameterKey{
			Name: "tag_id",
			In:   "query",
		}
		params.TagID = packed[key].(int)
	}
	{
		key := middleware.ParameterKey{
			Name: "feature_id",
			In:   "query",
		}
		params.FeatureID = packed[key].(int)
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

func decodeUserBannerGetParams(args [0]string, argsEscaped bool, r *http.Request) (params UserBannerGetParams, _ error) {
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

				c, err := conv.ToInt(val)
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

				c, err := conv.ToInt(val)
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
