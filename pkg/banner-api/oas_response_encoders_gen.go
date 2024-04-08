// Code generated by ogen, DO NOT EDIT.

package api

import (
	"net/http"

	"github.com/go-faster/errors"
	"github.com/go-faster/jx"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
)

func encodeCreateBannerResponse(response CreateBannerRes, w http.ResponseWriter, span trace.Span) error {
	switch response := response.(type) {
	case *CreateBannerResponse:
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(201)
		span.SetStatus(codes.Ok, http.StatusText(201))

		e := new(jx.Encoder)
		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil

	case *CreateBannerBadRequest:
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(400)
		span.SetStatus(codes.Error, http.StatusText(400))

		e := new(jx.Encoder)
		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil

	case *CreateBannerUnauthorized:
		w.WriteHeader(401)
		span.SetStatus(codes.Error, http.StatusText(401))

		return nil

	case *CreateBannerForbidden:
		w.WriteHeader(403)
		span.SetStatus(codes.Error, http.StatusText(403))

		return nil

	case *CreateBannerInternalServerError:
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(500)
		span.SetStatus(codes.Error, http.StatusText(500))

		e := new(jx.Encoder)
		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil

	default:
		return errors.Errorf("unexpected response type: %T", response)
	}
}

func encodeDeleteBannerResponse(response DeleteBannerRes, w http.ResponseWriter, span trace.Span) error {
	switch response := response.(type) {
	case *DeleteBannerNoContent:
		w.WriteHeader(204)
		span.SetStatus(codes.Ok, http.StatusText(204))

		return nil

	case *DeleteBannerBadRequest:
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(400)
		span.SetStatus(codes.Error, http.StatusText(400))

		e := new(jx.Encoder)
		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil

	case *DeleteBannerUnauthorized:
		w.WriteHeader(401)
		span.SetStatus(codes.Error, http.StatusText(401))

		return nil

	case *DeleteBannerForbidden:
		w.WriteHeader(403)
		span.SetStatus(codes.Error, http.StatusText(403))

		return nil

	case *DeleteBannerNotFound:
		w.WriteHeader(404)
		span.SetStatus(codes.Error, http.StatusText(404))

		return nil

	case *DeleteBannerInternalServerError:
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(500)
		span.SetStatus(codes.Error, http.StatusText(500))

		e := new(jx.Encoder)
		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil

	default:
		return errors.Errorf("unexpected response type: %T", response)
	}
}

func encodeGetBannerResponse(response GetBannerRes, w http.ResponseWriter, span trace.Span) error {
	switch response := response.(type) {
	case *GetBannerResponse:
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(200)
		span.SetStatus(codes.Ok, http.StatusText(200))

		e := new(jx.Encoder)
		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil

	case *GetBannerBadRequest:
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(400)
		span.SetStatus(codes.Error, http.StatusText(400))

		e := new(jx.Encoder)
		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil

	case *GetBannerUnauthorized:
		w.WriteHeader(401)
		span.SetStatus(codes.Error, http.StatusText(401))

		return nil

	case *GetBannerForbidden:
		w.WriteHeader(403)
		span.SetStatus(codes.Error, http.StatusText(403))

		return nil

	case *GetBannerNotFound:
		w.WriteHeader(404)
		span.SetStatus(codes.Error, http.StatusText(404))

		return nil

	case *GetBannerInternalServerError:
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(500)
		span.SetStatus(codes.Error, http.StatusText(500))

		e := new(jx.Encoder)
		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil

	default:
		return errors.Errorf("unexpected response type: %T", response)
	}
}

func encodeGetBannersResponse(response GetBannersRes, w http.ResponseWriter, span trace.Span) error {
	switch response := response.(type) {
	case *GetBannersResponse:
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(200)
		span.SetStatus(codes.Ok, http.StatusText(200))

		e := new(jx.Encoder)
		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil

	case *GetBannersUnauthorized:
		w.WriteHeader(401)
		span.SetStatus(codes.Error, http.StatusText(401))

		return nil

	case *GetBannersForbidden:
		w.WriteHeader(403)
		span.SetStatus(codes.Error, http.StatusText(403))

		return nil

	case *Error:
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(500)
		span.SetStatus(codes.Error, http.StatusText(500))

		e := new(jx.Encoder)
		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil

	default:
		return errors.Errorf("unexpected response type: %T", response)
	}
}

func encodeSetBannerResponse(response SetBannerRes, w http.ResponseWriter, span trace.Span) error {
	switch response := response.(type) {
	case *SetBannerOK:
		w.WriteHeader(200)
		span.SetStatus(codes.Ok, http.StatusText(200))

		return nil

	case *SetBannerBadRequest:
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(400)
		span.SetStatus(codes.Error, http.StatusText(400))

		e := new(jx.Encoder)
		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil

	case *SetBannerUnauthorized:
		w.WriteHeader(401)
		span.SetStatus(codes.Error, http.StatusText(401))

		return nil

	case *SetBannerForbidden:
		w.WriteHeader(403)
		span.SetStatus(codes.Error, http.StatusText(403))

		return nil

	case *SetBannerNotFound:
		w.WriteHeader(404)
		span.SetStatus(codes.Error, http.StatusText(404))

		return nil

	case *SetBannerInternalServerError:
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(500)
		span.SetStatus(codes.Error, http.StatusText(500))

		e := new(jx.Encoder)
		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil

	default:
		return errors.Errorf("unexpected response type: %T", response)
	}
}
