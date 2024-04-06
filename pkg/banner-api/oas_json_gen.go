// Code generated by ogen, DO NOT EDIT.

package api

import (
	"time"

	"github.com/go-faster/errors"
	"github.com/go-faster/jx"

	"github.com/ogen-go/ogen/json"
)

// Encode implements json.Marshaler.
func (s *BannerGetInternalServerError) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s *BannerGetInternalServerError) encodeFields(e *jx.Encoder) {
	{
		if s.Error.Set {
			e.FieldStart("error")
			s.Error.Encode(e)
		}
	}
}

var jsonFieldsNameOfBannerGetInternalServerError = [1]string{
	0: "error",
}

// Decode decodes BannerGetInternalServerError from json.
func (s *BannerGetInternalServerError) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode BannerGetInternalServerError to nil")
	}

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "error":
			if err := func() error {
				s.Error.Reset()
				if err := s.Error.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"error\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode BannerGetInternalServerError")
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *BannerGetInternalServerError) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *BannerGetInternalServerError) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode encodes BannerGetOKApplicationJSON as json.
func (s BannerGetOKApplicationJSON) Encode(e *jx.Encoder) {
	unwrapped := []BannerGetOKItem(s)

	e.ArrStart()
	for _, elem := range unwrapped {
		elem.Encode(e)
	}
	e.ArrEnd()
}

// Decode decodes BannerGetOKApplicationJSON from json.
func (s *BannerGetOKApplicationJSON) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode BannerGetOKApplicationJSON to nil")
	}
	var unwrapped []BannerGetOKItem
	if err := func() error {
		unwrapped = make([]BannerGetOKItem, 0)
		if err := d.Arr(func(d *jx.Decoder) error {
			var elem BannerGetOKItem
			if err := elem.Decode(d); err != nil {
				return err
			}
			unwrapped = append(unwrapped, elem)
			return nil
		}); err != nil {
			return err
		}
		return nil
	}(); err != nil {
		return errors.Wrap(err, "alias")
	}
	*s = BannerGetOKApplicationJSON(unwrapped)
	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s BannerGetOKApplicationJSON) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *BannerGetOKApplicationJSON) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode implements json.Marshaler.
func (s *BannerGetOKItem) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s *BannerGetOKItem) encodeFields(e *jx.Encoder) {
	{
		if s.BannerID.Set {
			e.FieldStart("banner_id")
			s.BannerID.Encode(e)
		}
	}
	{
		if s.TagIds != nil {
			e.FieldStart("tag_ids")
			e.ArrStart()
			for _, elem := range s.TagIds {
				e.Int(elem)
			}
			e.ArrEnd()
		}
	}
	{
		if s.FeatureID.Set {
			e.FieldStart("feature_id")
			s.FeatureID.Encode(e)
		}
	}
	{
		if s.Content.Set {
			e.FieldStart("content")
			s.Content.Encode(e)
		}
	}
	{
		if s.IsActive.Set {
			e.FieldStart("is_active")
			s.IsActive.Encode(e)
		}
	}
	{
		if s.CreatedAt.Set {
			e.FieldStart("created_at")
			s.CreatedAt.Encode(e, json.EncodeDateTime)
		}
	}
	{
		if s.UpdatedAt.Set {
			e.FieldStart("updated_at")
			s.UpdatedAt.Encode(e, json.EncodeDateTime)
		}
	}
}

var jsonFieldsNameOfBannerGetOKItem = [7]string{
	0: "banner_id",
	1: "tag_ids",
	2: "feature_id",
	3: "content",
	4: "is_active",
	5: "created_at",
	6: "updated_at",
}

// Decode decodes BannerGetOKItem from json.
func (s *BannerGetOKItem) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode BannerGetOKItem to nil")
	}

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "banner_id":
			if err := func() error {
				s.BannerID.Reset()
				if err := s.BannerID.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"banner_id\"")
			}
		case "tag_ids":
			if err := func() error {
				s.TagIds = make([]int, 0)
				if err := d.Arr(func(d *jx.Decoder) error {
					var elem int
					v, err := d.Int()
					elem = int(v)
					if err != nil {
						return err
					}
					s.TagIds = append(s.TagIds, elem)
					return nil
				}); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"tag_ids\"")
			}
		case "feature_id":
			if err := func() error {
				s.FeatureID.Reset()
				if err := s.FeatureID.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"feature_id\"")
			}
		case "content":
			if err := func() error {
				s.Content.Reset()
				if err := s.Content.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"content\"")
			}
		case "is_active":
			if err := func() error {
				s.IsActive.Reset()
				if err := s.IsActive.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"is_active\"")
			}
		case "created_at":
			if err := func() error {
				s.CreatedAt.Reset()
				if err := s.CreatedAt.Decode(d, json.DecodeDateTime); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"created_at\"")
			}
		case "updated_at":
			if err := func() error {
				s.UpdatedAt.Reset()
				if err := s.UpdatedAt.Decode(d, json.DecodeDateTime); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"updated_at\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode BannerGetOKItem")
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *BannerGetOKItem) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *BannerGetOKItem) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode implements json.Marshaler.
func (s BannerGetOKItemContent) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields implements json.Marshaler.
func (s BannerGetOKItemContent) encodeFields(e *jx.Encoder) {
	for k, elem := range s {
		e.FieldStart(k)

		if len(elem) != 0 {
			e.Raw(elem)
		}
	}
}

// Decode decodes BannerGetOKItemContent from json.
func (s *BannerGetOKItemContent) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode BannerGetOKItemContent to nil")
	}
	m := s.init()
	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		var elem jx.Raw
		if err := func() error {
			v, err := d.RawAppend(nil)
			elem = jx.Raw(v)
			if err != nil {
				return err
			}
			return nil
		}(); err != nil {
			return errors.Wrapf(err, "decode field %q", k)
		}
		m[string(k)] = elem
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode BannerGetOKItemContent")
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s BannerGetOKItemContent) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *BannerGetOKItemContent) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode implements json.Marshaler.
func (s *BannerIDDeleteBadRequest) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s *BannerIDDeleteBadRequest) encodeFields(e *jx.Encoder) {
	{
		if s.Error.Set {
			e.FieldStart("error")
			s.Error.Encode(e)
		}
	}
}

var jsonFieldsNameOfBannerIDDeleteBadRequest = [1]string{
	0: "error",
}

// Decode decodes BannerIDDeleteBadRequest from json.
func (s *BannerIDDeleteBadRequest) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode BannerIDDeleteBadRequest to nil")
	}

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "error":
			if err := func() error {
				s.Error.Reset()
				if err := s.Error.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"error\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode BannerIDDeleteBadRequest")
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *BannerIDDeleteBadRequest) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *BannerIDDeleteBadRequest) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode implements json.Marshaler.
func (s *BannerIDDeleteInternalServerError) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s *BannerIDDeleteInternalServerError) encodeFields(e *jx.Encoder) {
	{
		if s.Error.Set {
			e.FieldStart("error")
			s.Error.Encode(e)
		}
	}
}

var jsonFieldsNameOfBannerIDDeleteInternalServerError = [1]string{
	0: "error",
}

// Decode decodes BannerIDDeleteInternalServerError from json.
func (s *BannerIDDeleteInternalServerError) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode BannerIDDeleteInternalServerError to nil")
	}

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "error":
			if err := func() error {
				s.Error.Reset()
				if err := s.Error.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"error\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode BannerIDDeleteInternalServerError")
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *BannerIDDeleteInternalServerError) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *BannerIDDeleteInternalServerError) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode implements json.Marshaler.
func (s *BannerIDPatchBadRequest) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s *BannerIDPatchBadRequest) encodeFields(e *jx.Encoder) {
	{
		if s.Error.Set {
			e.FieldStart("error")
			s.Error.Encode(e)
		}
	}
}

var jsonFieldsNameOfBannerIDPatchBadRequest = [1]string{
	0: "error",
}

// Decode decodes BannerIDPatchBadRequest from json.
func (s *BannerIDPatchBadRequest) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode BannerIDPatchBadRequest to nil")
	}

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "error":
			if err := func() error {
				s.Error.Reset()
				if err := s.Error.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"error\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode BannerIDPatchBadRequest")
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *BannerIDPatchBadRequest) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *BannerIDPatchBadRequest) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode implements json.Marshaler.
func (s *BannerIDPatchInternalServerError) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s *BannerIDPatchInternalServerError) encodeFields(e *jx.Encoder) {
	{
		if s.Error.Set {
			e.FieldStart("error")
			s.Error.Encode(e)
		}
	}
}

var jsonFieldsNameOfBannerIDPatchInternalServerError = [1]string{
	0: "error",
}

// Decode decodes BannerIDPatchInternalServerError from json.
func (s *BannerIDPatchInternalServerError) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode BannerIDPatchInternalServerError to nil")
	}

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "error":
			if err := func() error {
				s.Error.Reset()
				if err := s.Error.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"error\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode BannerIDPatchInternalServerError")
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *BannerIDPatchInternalServerError) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *BannerIDPatchInternalServerError) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode implements json.Marshaler.
func (s *BannerIDPatchReq) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s *BannerIDPatchReq) encodeFields(e *jx.Encoder) {
	{
		if s.TagIds.Set {
			e.FieldStart("tag_ids")
			s.TagIds.Encode(e)
		}
	}
	{
		if s.FeatureID.Set {
			e.FieldStart("feature_id")
			s.FeatureID.Encode(e)
		}
	}
	{
		if s.Content.Set {
			e.FieldStart("content")
			s.Content.Encode(e)
		}
	}
	{
		if s.IsActive.Set {
			e.FieldStart("is_active")
			s.IsActive.Encode(e)
		}
	}
}

var jsonFieldsNameOfBannerIDPatchReq = [4]string{
	0: "tag_ids",
	1: "feature_id",
	2: "content",
	3: "is_active",
}

// Decode decodes BannerIDPatchReq from json.
func (s *BannerIDPatchReq) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode BannerIDPatchReq to nil")
	}

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "tag_ids":
			if err := func() error {
				s.TagIds.Reset()
				if err := s.TagIds.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"tag_ids\"")
			}
		case "feature_id":
			if err := func() error {
				s.FeatureID.Reset()
				if err := s.FeatureID.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"feature_id\"")
			}
		case "content":
			if err := func() error {
				s.Content.Reset()
				if err := s.Content.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"content\"")
			}
		case "is_active":
			if err := func() error {
				s.IsActive.Reset()
				if err := s.IsActive.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"is_active\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode BannerIDPatchReq")
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *BannerIDPatchReq) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *BannerIDPatchReq) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode implements json.Marshaler.
func (s BannerIDPatchReqContent) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields implements json.Marshaler.
func (s BannerIDPatchReqContent) encodeFields(e *jx.Encoder) {
	for k, elem := range s {
		e.FieldStart(k)

		if len(elem) != 0 {
			e.Raw(elem)
		}
	}
}

// Decode decodes BannerIDPatchReqContent from json.
func (s *BannerIDPatchReqContent) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode BannerIDPatchReqContent to nil")
	}
	m := s.init()
	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		var elem jx.Raw
		if err := func() error {
			v, err := d.RawAppend(nil)
			elem = jx.Raw(v)
			if err != nil {
				return err
			}
			return nil
		}(); err != nil {
			return errors.Wrapf(err, "decode field %q", k)
		}
		m[string(k)] = elem
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode BannerIDPatchReqContent")
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s BannerIDPatchReqContent) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *BannerIDPatchReqContent) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode implements json.Marshaler.
func (s *BannerPostBadRequest) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s *BannerPostBadRequest) encodeFields(e *jx.Encoder) {
	{
		if s.Error.Set {
			e.FieldStart("error")
			s.Error.Encode(e)
		}
	}
}

var jsonFieldsNameOfBannerPostBadRequest = [1]string{
	0: "error",
}

// Decode decodes BannerPostBadRequest from json.
func (s *BannerPostBadRequest) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode BannerPostBadRequest to nil")
	}

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "error":
			if err := func() error {
				s.Error.Reset()
				if err := s.Error.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"error\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode BannerPostBadRequest")
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *BannerPostBadRequest) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *BannerPostBadRequest) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode implements json.Marshaler.
func (s *BannerPostCreated) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s *BannerPostCreated) encodeFields(e *jx.Encoder) {
	{
		if s.BannerID.Set {
			e.FieldStart("banner_id")
			s.BannerID.Encode(e)
		}
	}
}

var jsonFieldsNameOfBannerPostCreated = [1]string{
	0: "banner_id",
}

// Decode decodes BannerPostCreated from json.
func (s *BannerPostCreated) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode BannerPostCreated to nil")
	}

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "banner_id":
			if err := func() error {
				s.BannerID.Reset()
				if err := s.BannerID.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"banner_id\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode BannerPostCreated")
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *BannerPostCreated) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *BannerPostCreated) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode implements json.Marshaler.
func (s *BannerPostInternalServerError) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s *BannerPostInternalServerError) encodeFields(e *jx.Encoder) {
	{
		if s.Error.Set {
			e.FieldStart("error")
			s.Error.Encode(e)
		}
	}
}

var jsonFieldsNameOfBannerPostInternalServerError = [1]string{
	0: "error",
}

// Decode decodes BannerPostInternalServerError from json.
func (s *BannerPostInternalServerError) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode BannerPostInternalServerError to nil")
	}

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "error":
			if err := func() error {
				s.Error.Reset()
				if err := s.Error.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"error\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode BannerPostInternalServerError")
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *BannerPostInternalServerError) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *BannerPostInternalServerError) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode implements json.Marshaler.
func (s *BannerPostReq) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s *BannerPostReq) encodeFields(e *jx.Encoder) {
	{
		if s.TagIds != nil {
			e.FieldStart("tag_ids")
			e.ArrStart()
			for _, elem := range s.TagIds {
				e.Int(elem)
			}
			e.ArrEnd()
		}
	}
	{
		if s.FeatureID.Set {
			e.FieldStart("feature_id")
			s.FeatureID.Encode(e)
		}
	}
	{
		if s.Content.Set {
			e.FieldStart("content")
			s.Content.Encode(e)
		}
	}
	{
		if s.IsActive.Set {
			e.FieldStart("is_active")
			s.IsActive.Encode(e)
		}
	}
}

var jsonFieldsNameOfBannerPostReq = [4]string{
	0: "tag_ids",
	1: "feature_id",
	2: "content",
	3: "is_active",
}

// Decode decodes BannerPostReq from json.
func (s *BannerPostReq) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode BannerPostReq to nil")
	}

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "tag_ids":
			if err := func() error {
				s.TagIds = make([]int, 0)
				if err := d.Arr(func(d *jx.Decoder) error {
					var elem int
					v, err := d.Int()
					elem = int(v)
					if err != nil {
						return err
					}
					s.TagIds = append(s.TagIds, elem)
					return nil
				}); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"tag_ids\"")
			}
		case "feature_id":
			if err := func() error {
				s.FeatureID.Reset()
				if err := s.FeatureID.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"feature_id\"")
			}
		case "content":
			if err := func() error {
				s.Content.Reset()
				if err := s.Content.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"content\"")
			}
		case "is_active":
			if err := func() error {
				s.IsActive.Reset()
				if err := s.IsActive.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"is_active\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode BannerPostReq")
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *BannerPostReq) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *BannerPostReq) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode implements json.Marshaler.
func (s BannerPostReqContent) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields implements json.Marshaler.
func (s BannerPostReqContent) encodeFields(e *jx.Encoder) {
	for k, elem := range s {
		e.FieldStart(k)

		if len(elem) != 0 {
			e.Raw(elem)
		}
	}
}

// Decode decodes BannerPostReqContent from json.
func (s *BannerPostReqContent) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode BannerPostReqContent to nil")
	}
	m := s.init()
	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		var elem jx.Raw
		if err := func() error {
			v, err := d.RawAppend(nil)
			elem = jx.Raw(v)
			if err != nil {
				return err
			}
			return nil
		}(); err != nil {
			return errors.Wrapf(err, "decode field %q", k)
		}
		m[string(k)] = elem
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode BannerPostReqContent")
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s BannerPostReqContent) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *BannerPostReqContent) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode encodes BannerGetOKItemContent as json.
func (o OptBannerGetOKItemContent) Encode(e *jx.Encoder) {
	if !o.Set {
		return
	}
	o.Value.Encode(e)
}

// Decode decodes BannerGetOKItemContent from json.
func (o *OptBannerGetOKItemContent) Decode(d *jx.Decoder) error {
	if o == nil {
		return errors.New("invalid: unable to decode OptBannerGetOKItemContent to nil")
	}
	o.Set = true
	o.Value = make(BannerGetOKItemContent)
	if err := o.Value.Decode(d); err != nil {
		return err
	}
	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s OptBannerGetOKItemContent) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *OptBannerGetOKItemContent) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode encodes BannerPostReqContent as json.
func (o OptBannerPostReqContent) Encode(e *jx.Encoder) {
	if !o.Set {
		return
	}
	o.Value.Encode(e)
}

// Decode decodes BannerPostReqContent from json.
func (o *OptBannerPostReqContent) Decode(d *jx.Decoder) error {
	if o == nil {
		return errors.New("invalid: unable to decode OptBannerPostReqContent to nil")
	}
	o.Set = true
	o.Value = make(BannerPostReqContent)
	if err := o.Value.Decode(d); err != nil {
		return err
	}
	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s OptBannerPostReqContent) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *OptBannerPostReqContent) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode encodes bool as json.
func (o OptBool) Encode(e *jx.Encoder) {
	if !o.Set {
		return
	}
	e.Bool(bool(o.Value))
}

// Decode decodes bool from json.
func (o *OptBool) Decode(d *jx.Decoder) error {
	if o == nil {
		return errors.New("invalid: unable to decode OptBool to nil")
	}
	o.Set = true
	v, err := d.Bool()
	if err != nil {
		return err
	}
	o.Value = bool(v)
	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s OptBool) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *OptBool) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode encodes time.Time as json.
func (o OptDateTime) Encode(e *jx.Encoder, format func(*jx.Encoder, time.Time)) {
	if !o.Set {
		return
	}
	format(e, o.Value)
}

// Decode decodes time.Time from json.
func (o *OptDateTime) Decode(d *jx.Decoder, format func(*jx.Decoder) (time.Time, error)) error {
	if o == nil {
		return errors.New("invalid: unable to decode OptDateTime to nil")
	}
	o.Set = true
	v, err := format(d)
	if err != nil {
		return err
	}
	o.Value = v
	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s OptDateTime) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e, json.EncodeDateTime)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *OptDateTime) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d, json.DecodeDateTime)
}

// Encode encodes int as json.
func (o OptInt) Encode(e *jx.Encoder) {
	if !o.Set {
		return
	}
	e.Int(int(o.Value))
}

// Decode decodes int from json.
func (o *OptInt) Decode(d *jx.Decoder) error {
	if o == nil {
		return errors.New("invalid: unable to decode OptInt to nil")
	}
	o.Set = true
	v, err := d.Int()
	if err != nil {
		return err
	}
	o.Value = int(v)
	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s OptInt) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *OptInt) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode encodes BannerIDPatchReqContent as json.
func (o OptNilBannerIDPatchReqContent) Encode(e *jx.Encoder) {
	if !o.Set {
		return
	}
	if o.Null {
		e.Null()
		return
	}
	o.Value.Encode(e)
}

// Decode decodes BannerIDPatchReqContent from json.
func (o *OptNilBannerIDPatchReqContent) Decode(d *jx.Decoder) error {
	if o == nil {
		return errors.New("invalid: unable to decode OptNilBannerIDPatchReqContent to nil")
	}
	if d.Next() == jx.Null {
		if err := d.Null(); err != nil {
			return err
		}

		var v BannerIDPatchReqContent
		o.Value = v
		o.Set = true
		o.Null = true
		return nil
	}
	o.Set = true
	o.Null = false
	o.Value = make(BannerIDPatchReqContent)
	if err := o.Value.Decode(d); err != nil {
		return err
	}
	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s OptNilBannerIDPatchReqContent) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *OptNilBannerIDPatchReqContent) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode encodes bool as json.
func (o OptNilBool) Encode(e *jx.Encoder) {
	if !o.Set {
		return
	}
	if o.Null {
		e.Null()
		return
	}
	e.Bool(bool(o.Value))
}

// Decode decodes bool from json.
func (o *OptNilBool) Decode(d *jx.Decoder) error {
	if o == nil {
		return errors.New("invalid: unable to decode OptNilBool to nil")
	}
	if d.Next() == jx.Null {
		if err := d.Null(); err != nil {
			return err
		}

		var v bool
		o.Value = v
		o.Set = true
		o.Null = true
		return nil
	}
	o.Set = true
	o.Null = false
	v, err := d.Bool()
	if err != nil {
		return err
	}
	o.Value = bool(v)
	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s OptNilBool) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *OptNilBool) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode encodes int as json.
func (o OptNilInt) Encode(e *jx.Encoder) {
	if !o.Set {
		return
	}
	if o.Null {
		e.Null()
		return
	}
	e.Int(int(o.Value))
}

// Decode decodes int from json.
func (o *OptNilInt) Decode(d *jx.Decoder) error {
	if o == nil {
		return errors.New("invalid: unable to decode OptNilInt to nil")
	}
	if d.Next() == jx.Null {
		if err := d.Null(); err != nil {
			return err
		}

		var v int
		o.Value = v
		o.Set = true
		o.Null = true
		return nil
	}
	o.Set = true
	o.Null = false
	v, err := d.Int()
	if err != nil {
		return err
	}
	o.Value = int(v)
	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s OptNilInt) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *OptNilInt) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode encodes []int as json.
func (o OptNilIntArray) Encode(e *jx.Encoder) {
	if !o.Set {
		return
	}
	if o.Null {
		e.Null()
		return
	}
	e.ArrStart()
	for _, elem := range o.Value {
		e.Int(elem)
	}
	e.ArrEnd()
}

// Decode decodes []int from json.
func (o *OptNilIntArray) Decode(d *jx.Decoder) error {
	if o == nil {
		return errors.New("invalid: unable to decode OptNilIntArray to nil")
	}
	if d.Next() == jx.Null {
		if err := d.Null(); err != nil {
			return err
		}

		var v []int
		o.Value = v
		o.Set = true
		o.Null = true
		return nil
	}
	o.Set = true
	o.Null = false
	o.Value = make([]int, 0)
	if err := d.Arr(func(d *jx.Decoder) error {
		var elem int
		v, err := d.Int()
		elem = int(v)
		if err != nil {
			return err
		}
		o.Value = append(o.Value, elem)
		return nil
	}); err != nil {
		return err
	}
	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s OptNilIntArray) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *OptNilIntArray) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode encodes string as json.
func (o OptString) Encode(e *jx.Encoder) {
	if !o.Set {
		return
	}
	e.Str(string(o.Value))
}

// Decode decodes string from json.
func (o *OptString) Decode(d *jx.Decoder) error {
	if o == nil {
		return errors.New("invalid: unable to decode OptString to nil")
	}
	o.Set = true
	v, err := d.Str()
	if err != nil {
		return err
	}
	o.Value = string(v)
	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s OptString) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *OptString) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode implements json.Marshaler.
func (s *UserBannerGetBadRequest) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s *UserBannerGetBadRequest) encodeFields(e *jx.Encoder) {
	{
		if s.Error.Set {
			e.FieldStart("error")
			s.Error.Encode(e)
		}
	}
}

var jsonFieldsNameOfUserBannerGetBadRequest = [1]string{
	0: "error",
}

// Decode decodes UserBannerGetBadRequest from json.
func (s *UserBannerGetBadRequest) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode UserBannerGetBadRequest to nil")
	}

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "error":
			if err := func() error {
				s.Error.Reset()
				if err := s.Error.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"error\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode UserBannerGetBadRequest")
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *UserBannerGetBadRequest) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *UserBannerGetBadRequest) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode implements json.Marshaler.
func (s *UserBannerGetInternalServerError) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s *UserBannerGetInternalServerError) encodeFields(e *jx.Encoder) {
	{
		if s.Error.Set {
			e.FieldStart("error")
			s.Error.Encode(e)
		}
	}
}

var jsonFieldsNameOfUserBannerGetInternalServerError = [1]string{
	0: "error",
}

// Decode decodes UserBannerGetInternalServerError from json.
func (s *UserBannerGetInternalServerError) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode UserBannerGetInternalServerError to nil")
	}

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "error":
			if err := func() error {
				s.Error.Reset()
				if err := s.Error.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"error\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode UserBannerGetInternalServerError")
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *UserBannerGetInternalServerError) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *UserBannerGetInternalServerError) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode implements json.Marshaler.
func (s UserBannerGetOK) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields implements json.Marshaler.
func (s UserBannerGetOK) encodeFields(e *jx.Encoder) {
	for k, elem := range s {
		e.FieldStart(k)

		if len(elem) != 0 {
			e.Raw(elem)
		}
	}
}

// Decode decodes UserBannerGetOK from json.
func (s *UserBannerGetOK) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode UserBannerGetOK to nil")
	}
	m := s.init()
	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		var elem jx.Raw
		if err := func() error {
			v, err := d.RawAppend(nil)
			elem = jx.Raw(v)
			if err != nil {
				return err
			}
			return nil
		}(); err != nil {
			return errors.Wrapf(err, "decode field %q", k)
		}
		m[string(k)] = elem
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode UserBannerGetOK")
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s UserBannerGetOK) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *UserBannerGetOK) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}
