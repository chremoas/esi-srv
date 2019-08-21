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

func easyjsonFa8e3062DecodeGithubComAntihaxGoesiEsi(in *jlexer.Lexer, out *GetCorporationsCorporationIdDivisionsWalletWalletList) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		in.Skip()
		*out = nil
	} else {
		in.Delim('[')
		if *out == nil {
			if !in.IsDelim(']') {
				*out = make(GetCorporationsCorporationIdDivisionsWalletWalletList, 0, 2)
			} else {
				*out = GetCorporationsCorporationIdDivisionsWalletWalletList{}
			}
		} else {
			*out = (*out)[:0]
		}
		for !in.IsDelim(']') {
			var v1 GetCorporationsCorporationIdDivisionsWalletWallet
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
func easyjsonFa8e3062EncodeGithubComAntihaxGoesiEsi(out *jwriter.Writer, in GetCorporationsCorporationIdDivisionsWalletWalletList) {
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
func (v GetCorporationsCorporationIdDivisionsWalletWalletList) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonFa8e3062EncodeGithubComAntihaxGoesiEsi(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GetCorporationsCorporationIdDivisionsWalletWalletList) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonFa8e3062EncodeGithubComAntihaxGoesiEsi(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GetCorporationsCorporationIdDivisionsWalletWalletList) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonFa8e3062DecodeGithubComAntihaxGoesiEsi(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GetCorporationsCorporationIdDivisionsWalletWalletList) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonFa8e3062DecodeGithubComAntihaxGoesiEsi(l, v)
}
func easyjsonFa8e3062DecodeGithubComAntihaxGoesiEsi1(in *jlexer.Lexer, out *GetCorporationsCorporationIdDivisionsWalletWallet) {
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
		case "division":
			out.Division = int32(in.Int32())
		case "name":
			out.Name = string(in.String())
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
func easyjsonFa8e3062EncodeGithubComAntihaxGoesiEsi1(out *jwriter.Writer, in GetCorporationsCorporationIdDivisionsWalletWallet) {
	out.RawByte('{')
	first := true
	_ = first
	if in.Division != 0 {
		const prefix string = ",\"division\":"
		first = false
		out.RawString(prefix[1:])
		out.Int32(int32(in.Division))
	}
	if in.Name != "" {
		const prefix string = ",\"name\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Name))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v GetCorporationsCorporationIdDivisionsWalletWallet) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonFa8e3062EncodeGithubComAntihaxGoesiEsi1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GetCorporationsCorporationIdDivisionsWalletWallet) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonFa8e3062EncodeGithubComAntihaxGoesiEsi1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GetCorporationsCorporationIdDivisionsWalletWallet) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonFa8e3062DecodeGithubComAntihaxGoesiEsi1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GetCorporationsCorporationIdDivisionsWalletWallet) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonFa8e3062DecodeGithubComAntihaxGoesiEsi1(l, v)
}
