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

func easyjsonFed59962DecodeGithubComAntihaxGoesiEsi(in *jlexer.Lexer, out *GetUniverseSystemsSystemIdPlanetList) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		in.Skip()
		*out = nil
	} else {
		in.Delim('[')
		if *out == nil {
			if !in.IsDelim(']') {
				*out = make(GetUniverseSystemsSystemIdPlanetList, 0, 1)
			} else {
				*out = GetUniverseSystemsSystemIdPlanetList{}
			}
		} else {
			*out = (*out)[:0]
		}
		for !in.IsDelim(']') {
			var v1 GetUniverseSystemsSystemIdPlanet
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
func easyjsonFed59962EncodeGithubComAntihaxGoesiEsi(out *jwriter.Writer, in GetUniverseSystemsSystemIdPlanetList) {
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
func (v GetUniverseSystemsSystemIdPlanetList) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonFed59962EncodeGithubComAntihaxGoesiEsi(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GetUniverseSystemsSystemIdPlanetList) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonFed59962EncodeGithubComAntihaxGoesiEsi(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GetUniverseSystemsSystemIdPlanetList) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonFed59962DecodeGithubComAntihaxGoesiEsi(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GetUniverseSystemsSystemIdPlanetList) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonFed59962DecodeGithubComAntihaxGoesiEsi(l, v)
}
func easyjsonFed59962DecodeGithubComAntihaxGoesiEsi1(in *jlexer.Lexer, out *GetUniverseSystemsSystemIdPlanet) {
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
		case "asteroid_belts":
			if in.IsNull() {
				in.Skip()
				out.AsteroidBelts = nil
			} else {
				in.Delim('[')
				if out.AsteroidBelts == nil {
					if !in.IsDelim(']') {
						out.AsteroidBelts = make([]int32, 0, 16)
					} else {
						out.AsteroidBelts = []int32{}
					}
				} else {
					out.AsteroidBelts = (out.AsteroidBelts)[:0]
				}
				for !in.IsDelim(']') {
					var v4 int32
					v4 = int32(in.Int32())
					out.AsteroidBelts = append(out.AsteroidBelts, v4)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "moons":
			if in.IsNull() {
				in.Skip()
				out.Moons = nil
			} else {
				in.Delim('[')
				if out.Moons == nil {
					if !in.IsDelim(']') {
						out.Moons = make([]int32, 0, 16)
					} else {
						out.Moons = []int32{}
					}
				} else {
					out.Moons = (out.Moons)[:0]
				}
				for !in.IsDelim(']') {
					var v5 int32
					v5 = int32(in.Int32())
					out.Moons = append(out.Moons, v5)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "planet_id":
			out.PlanetId = int32(in.Int32())
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
func easyjsonFed59962EncodeGithubComAntihaxGoesiEsi1(out *jwriter.Writer, in GetUniverseSystemsSystemIdPlanet) {
	out.RawByte('{')
	first := true
	_ = first
	if len(in.AsteroidBelts) != 0 {
		const prefix string = ",\"asteroid_belts\":"
		first = false
		out.RawString(prefix[1:])
		{
			out.RawByte('[')
			for v6, v7 := range in.AsteroidBelts {
				if v6 > 0 {
					out.RawByte(',')
				}
				out.Int32(int32(v7))
			}
			out.RawByte(']')
		}
	}
	if len(in.Moons) != 0 {
		const prefix string = ",\"moons\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		{
			out.RawByte('[')
			for v8, v9 := range in.Moons {
				if v8 > 0 {
					out.RawByte(',')
				}
				out.Int32(int32(v9))
			}
			out.RawByte(']')
		}
	}
	if in.PlanetId != 0 {
		const prefix string = ",\"planet_id\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int32(int32(in.PlanetId))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v GetUniverseSystemsSystemIdPlanet) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonFed59962EncodeGithubComAntihaxGoesiEsi1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GetUniverseSystemsSystemIdPlanet) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonFed59962EncodeGithubComAntihaxGoesiEsi1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GetUniverseSystemsSystemIdPlanet) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonFed59962DecodeGithubComAntihaxGoesiEsi1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GetUniverseSystemsSystemIdPlanet) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonFed59962DecodeGithubComAntihaxGoesiEsi1(l, v)
}
