// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package esi

import (
	json "encoding/json"

	easyjson "github.com/mailru/easyjson"
	jlexer "github.com/mailru/easyjson/jlexer"
	jwriter "github.com/mailru/easyjson/jwriter"
)

// suppress unused package warning
var (
	_ *json.RawMessage
	_ *jlexer.Lexer
	_ *jwriter.Writer
	_ easyjson.Marshaler
)

func easyjsonCdc2fd8fDecodeGithubComAntihaxGoesiEsi(in *jlexer.Lexer, out *GetMarketsRegionIdOrders200OkList) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		in.Skip()
		*out = nil
	} else {
		in.Delim('[')
		if *out == nil {
			if !in.IsDelim(']') {
				*out = make(GetMarketsRegionIdOrders200OkList, 0, 1)
			} else {
				*out = GetMarketsRegionIdOrders200OkList{}
			}
		} else {
			*out = (*out)[:0]
		}
		for !in.IsDelim(']') {
			var v1 GetMarketsRegionIdOrders200Ok
			(v1).UnmarshalEasyJSON(in)
			*out = append(*out, v1)
			in.WantComma()
		}
		in.Delim(']')
	}
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonCdc2fd8fEncodeGithubComAntihaxGoesiEsi(out *jwriter.Writer, in GetMarketsRegionIdOrders200OkList) {
	if in == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
		out.RawString("null")
	} else {
		out.RawByte('[')
		for v2, v3 := range in {
			if v2 > 0 {
				out.RawByte(',')
			}
			(v3).MarshalEasyJSON(out)
		}
		out.RawByte(']')
	}
}

// MarshalJSON supports json.Marshaler interface
func (v GetMarketsRegionIdOrders200OkList) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonCdc2fd8fEncodeGithubComAntihaxGoesiEsi(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GetMarketsRegionIdOrders200OkList) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonCdc2fd8fEncodeGithubComAntihaxGoesiEsi(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GetMarketsRegionIdOrders200OkList) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonCdc2fd8fDecodeGithubComAntihaxGoesiEsi(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GetMarketsRegionIdOrders200OkList) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonCdc2fd8fDecodeGithubComAntihaxGoesiEsi(l, v)
}
func easyjsonCdc2fd8fDecodeGithubComAntihaxGoesiEsi1(in *jlexer.Lexer, out *GetMarketsRegionIdOrders200Ok) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "duration":
			out.Duration = int32(in.Int32())
		case "is_buy_order":
			out.IsBuyOrder = bool(in.Bool())
		case "issued":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.Issued).UnmarshalJSON(data))
			}
		case "location_id":
			out.LocationId = int64(in.Int64())
		case "min_volume":
			out.MinVolume = int32(in.Int32())
		case "order_id":
			out.OrderId = int64(in.Int64())
		case "price":
			out.Price = float64(in.Float64())
		case "range":
			out.Range_ = string(in.String())
		case "system_id":
			out.SystemId = int32(in.Int32())
		case "type_id":
			out.TypeId = int32(in.Int32())
		case "volume_remain":
			out.VolumeRemain = int32(in.Int32())
		case "volume_total":
			out.VolumeTotal = int32(in.Int32())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonCdc2fd8fEncodeGithubComAntihaxGoesiEsi1(out *jwriter.Writer, in GetMarketsRegionIdOrders200Ok) {
	out.RawByte('{')
	first := true
	_ = first
	if in.Duration != 0 {
		const prefix string = ",\"duration\":"
		first = false
		out.RawString(prefix[1:])
		out.Int32(int32(in.Duration))
	}
	if in.IsBuyOrder {
		const prefix string = ",\"is_buy_order\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Bool(bool(in.IsBuyOrder))
	}
	if true {
		const prefix string = ",\"issued\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Raw((in.Issued).MarshalJSON())
	}
	if in.LocationId != 0 {
		const prefix string = ",\"location_id\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int64(int64(in.LocationId))
	}
	if in.MinVolume != 0 {
		const prefix string = ",\"min_volume\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int32(int32(in.MinVolume))
	}
	if in.OrderId != 0 {
		const prefix string = ",\"order_id\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int64(int64(in.OrderId))
	}
	if in.Price != 0 {
		const prefix string = ",\"price\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Float64(float64(in.Price))
	}
	if in.Range_ != "" {
		const prefix string = ",\"range\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Range_))
	}
	if in.SystemId != 0 {
		const prefix string = ",\"system_id\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int32(int32(in.SystemId))
	}
	if in.TypeId != 0 {
		const prefix string = ",\"type_id\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int32(int32(in.TypeId))
	}
	if in.VolumeRemain != 0 {
		const prefix string = ",\"volume_remain\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int32(int32(in.VolumeRemain))
	}
	if in.VolumeTotal != 0 {
		const prefix string = ",\"volume_total\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int32(int32(in.VolumeTotal))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v GetMarketsRegionIdOrders200Ok) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonCdc2fd8fEncodeGithubComAntihaxGoesiEsi1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GetMarketsRegionIdOrders200Ok) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonCdc2fd8fEncodeGithubComAntihaxGoesiEsi1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GetMarketsRegionIdOrders200Ok) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonCdc2fd8fDecodeGithubComAntihaxGoesiEsi1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GetMarketsRegionIdOrders200Ok) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonCdc2fd8fDecodeGithubComAntihaxGoesiEsi1(l, v)
}
