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

func easyjsonF3c409e3DecodeGithubComAntihaxGoesiEsi(in *jlexer.Lexer, out *PostFleetsFleetIdWingsCreatedList) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		in.Skip()
		*out = nil
	} else {
		in.Delim('[')
		if *out == nil {
			if !in.IsDelim(']') {
				*out = make(PostFleetsFleetIdWingsCreatedList, 0, 8)
			} else {
				*out = PostFleetsFleetIdWingsCreatedList{}
			}
		} else {
			*out = (*out)[:0]
		}
		for !in.IsDelim(']') {
			var v1 PostFleetsFleetIdWingsCreated
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
func easyjsonF3c409e3EncodeGithubComAntihaxGoesiEsi(out *jwriter.Writer, in PostFleetsFleetIdWingsCreatedList) {
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
func (v PostFleetsFleetIdWingsCreatedList) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonF3c409e3EncodeGithubComAntihaxGoesiEsi(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v PostFleetsFleetIdWingsCreatedList) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonF3c409e3EncodeGithubComAntihaxGoesiEsi(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *PostFleetsFleetIdWingsCreatedList) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonF3c409e3DecodeGithubComAntihaxGoesiEsi(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *PostFleetsFleetIdWingsCreatedList) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonF3c409e3DecodeGithubComAntihaxGoesiEsi(l, v)
}
func easyjsonF3c409e3DecodeGithubComAntihaxGoesiEsi1(in *jlexer.Lexer, out *PostFleetsFleetIdWingsCreated) {
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
		case "wing_id":
			out.WingId = int64(in.Int64())
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
func easyjsonF3c409e3EncodeGithubComAntihaxGoesiEsi1(out *jwriter.Writer, in PostFleetsFleetIdWingsCreated) {
	out.RawByte('{')
	first := true
	_ = first
	if in.WingId != 0 {
		const prefix string = ",\"wing_id\":"
		first = false
		out.RawString(prefix[1:])
		out.Int64(int64(in.WingId))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v PostFleetsFleetIdWingsCreated) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonF3c409e3EncodeGithubComAntihaxGoesiEsi1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v PostFleetsFleetIdWingsCreated) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonF3c409e3EncodeGithubComAntihaxGoesiEsi1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *PostFleetsFleetIdWingsCreated) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonF3c409e3DecodeGithubComAntihaxGoesiEsi1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *PostFleetsFleetIdWingsCreated) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonF3c409e3DecodeGithubComAntihaxGoesiEsi1(l, v)
}
