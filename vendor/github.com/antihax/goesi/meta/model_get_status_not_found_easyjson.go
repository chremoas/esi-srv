// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package meta

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

func easyjsonFbb8b2acDecodeGithubComAntihaxGoesiMeta(in *jlexer.Lexer, out *GetStatusNotFoundList) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		in.Skip()
		*out = nil
	} else {
		in.Delim('[')
		if *out == nil {
			if !in.IsDelim(']') {
				*out = make(GetStatusNotFoundList, 0, 4)
			} else {
				*out = GetStatusNotFoundList{}
			}
		} else {
			*out = (*out)[:0]
		}
		for !in.IsDelim(']') {
			var v1 GetStatusNotFound
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
func easyjsonFbb8b2acEncodeGithubComAntihaxGoesiMeta(out *jwriter.Writer, in GetStatusNotFoundList) {
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
func (v GetStatusNotFoundList) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonFbb8b2acEncodeGithubComAntihaxGoesiMeta(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GetStatusNotFoundList) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonFbb8b2acEncodeGithubComAntihaxGoesiMeta(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GetStatusNotFoundList) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonFbb8b2acDecodeGithubComAntihaxGoesiMeta(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GetStatusNotFoundList) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonFbb8b2acDecodeGithubComAntihaxGoesiMeta(l, v)
}
func easyjsonFbb8b2acDecodeGithubComAntihaxGoesiMeta1(in *jlexer.Lexer, out *GetStatusNotFound) {
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
		case "error":
			out.Error_ = string(in.String())
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
func easyjsonFbb8b2acEncodeGithubComAntihaxGoesiMeta1(out *jwriter.Writer, in GetStatusNotFound) {
	out.RawByte('{')
	first := true
	_ = first
	if in.Error_ != "" {
		const prefix string = ",\"error\":"
		first = false
		out.RawString(prefix[1:])
		out.String(string(in.Error_))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v GetStatusNotFound) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonFbb8b2acEncodeGithubComAntihaxGoesiMeta1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GetStatusNotFound) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonFbb8b2acEncodeGithubComAntihaxGoesiMeta1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GetStatusNotFound) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonFbb8b2acDecodeGithubComAntihaxGoesiMeta1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GetStatusNotFound) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonFbb8b2acDecodeGithubComAntihaxGoesiMeta1(l, v)
}
