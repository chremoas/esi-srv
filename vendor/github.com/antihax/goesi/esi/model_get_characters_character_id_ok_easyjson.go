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

func easyjsonD2c9dba0DecodeGithubComAntihaxGoesiEsi(in *jlexer.Lexer, out *GetCharactersCharacterIdOkList) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		in.Skip()
		*out = nil
	} else {
		in.Delim('[')
		if *out == nil {
			if !in.IsDelim(']') {
				*out = make(GetCharactersCharacterIdOkList, 0, 1)
			} else {
				*out = GetCharactersCharacterIdOkList{}
			}
		} else {
			*out = (*out)[:0]
		}
		for !in.IsDelim(']') {
			var v1 GetCharactersCharacterIdOk
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
func easyjsonD2c9dba0EncodeGithubComAntihaxGoesiEsi(out *jwriter.Writer, in GetCharactersCharacterIdOkList) {
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
func (v GetCharactersCharacterIdOkList) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonD2c9dba0EncodeGithubComAntihaxGoesiEsi(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GetCharactersCharacterIdOkList) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonD2c9dba0EncodeGithubComAntihaxGoesiEsi(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GetCharactersCharacterIdOkList) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonD2c9dba0DecodeGithubComAntihaxGoesiEsi(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GetCharactersCharacterIdOkList) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonD2c9dba0DecodeGithubComAntihaxGoesiEsi(l, v)
}
func easyjsonD2c9dba0DecodeGithubComAntihaxGoesiEsi1(in *jlexer.Lexer, out *GetCharactersCharacterIdOk) {
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
		case "alliance_id":
			out.AllianceId = int32(in.Int32())
		case "ancestry_id":
			out.AncestryId = int32(in.Int32())
		case "birthday":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.Birthday).UnmarshalJSON(data))
			}
		case "bloodline_id":
			out.BloodlineId = int32(in.Int32())
		case "corporation_id":
			out.CorporationId = int32(in.Int32())
		case "description":
			out.Description = string(in.String())
		case "faction_id":
			out.FactionId = int32(in.Int32())
		case "gender":
			out.Gender = string(in.String())
		case "name":
			out.Name = string(in.String())
		case "race_id":
			out.RaceId = int32(in.Int32())
		case "security_status":
			out.SecurityStatus = float32(in.Float32())
		case "title":
			out.Title = string(in.String())
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
func easyjsonD2c9dba0EncodeGithubComAntihaxGoesiEsi1(out *jwriter.Writer, in GetCharactersCharacterIdOk) {
	out.RawByte('{')
	first := true
	_ = first
	if in.AllianceId != 0 {
		const prefix string = ",\"alliance_id\":"
		first = false
		out.RawString(prefix[1:])
		out.Int32(int32(in.AllianceId))
	}
	if in.AncestryId != 0 {
		const prefix string = ",\"ancestry_id\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int32(int32(in.AncestryId))
	}
	if true {
		const prefix string = ",\"birthday\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Raw((in.Birthday).MarshalJSON())
	}
	if in.BloodlineId != 0 {
		const prefix string = ",\"bloodline_id\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int32(int32(in.BloodlineId))
	}
	if in.CorporationId != 0 {
		const prefix string = ",\"corporation_id\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int32(int32(in.CorporationId))
	}
	if in.Description != "" {
		const prefix string = ",\"description\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Description))
	}
	if in.FactionId != 0 {
		const prefix string = ",\"faction_id\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int32(int32(in.FactionId))
	}
	if in.Gender != "" {
		const prefix string = ",\"gender\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Gender))
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
	if in.RaceId != 0 {
		const prefix string = ",\"race_id\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int32(int32(in.RaceId))
	}
	if in.SecurityStatus != 0 {
		const prefix string = ",\"security_status\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Float32(float32(in.SecurityStatus))
	}
	if in.Title != "" {
		const prefix string = ",\"title\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Title))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v GetCharactersCharacterIdOk) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonD2c9dba0EncodeGithubComAntihaxGoesiEsi1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GetCharactersCharacterIdOk) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonD2c9dba0EncodeGithubComAntihaxGoesiEsi1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GetCharactersCharacterIdOk) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonD2c9dba0DecodeGithubComAntihaxGoesiEsi1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GetCharactersCharacterIdOk) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonD2c9dba0DecodeGithubComAntihaxGoesiEsi1(l, v)
}
