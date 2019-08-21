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

func easyjsonDebc599dDecodeGithubComAntihaxGoesiEsi(in *jlexer.Lexer, out *GetCharactersCharacterIdSearchOkList) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		in.Skip()
		*out = nil
	} else {
		in.Delim('[')
		if *out == nil {
			if !in.IsDelim(']') {
				*out = make(GetCharactersCharacterIdSearchOkList, 0, 1)
			} else {
				*out = GetCharactersCharacterIdSearchOkList{}
			}
		} else {
			*out = (*out)[:0]
		}
		for !in.IsDelim(']') {
			var v1 GetCharactersCharacterIdSearchOk
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
func easyjsonDebc599dEncodeGithubComAntihaxGoesiEsi(out *jwriter.Writer, in GetCharactersCharacterIdSearchOkList) {
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
func (v GetCharactersCharacterIdSearchOkList) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonDebc599dEncodeGithubComAntihaxGoesiEsi(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GetCharactersCharacterIdSearchOkList) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonDebc599dEncodeGithubComAntihaxGoesiEsi(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GetCharactersCharacterIdSearchOkList) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonDebc599dDecodeGithubComAntihaxGoesiEsi(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GetCharactersCharacterIdSearchOkList) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonDebc599dDecodeGithubComAntihaxGoesiEsi(l, v)
}
func easyjsonDebc599dDecodeGithubComAntihaxGoesiEsi1(in *jlexer.Lexer, out *GetCharactersCharacterIdSearchOk) {
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
		case "agent":
			if in.IsNull() {
				in.Skip()
				out.Agent = nil
			} else {
				in.Delim('[')
				if out.Agent == nil {
					if !in.IsDelim(']') {
						out.Agent = make([]int32, 0, 16)
					} else {
						out.Agent = []int32{}
					}
				} else {
					out.Agent = (out.Agent)[:0]
				}
				for !in.IsDelim(']') {
					var v4 int32
					v4 = int32(in.Int32())
					out.Agent = append(out.Agent, v4)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "alliance":
			if in.IsNull() {
				in.Skip()
				out.Alliance = nil
			} else {
				in.Delim('[')
				if out.Alliance == nil {
					if !in.IsDelim(']') {
						out.Alliance = make([]int32, 0, 16)
					} else {
						out.Alliance = []int32{}
					}
				} else {
					out.Alliance = (out.Alliance)[:0]
				}
				for !in.IsDelim(']') {
					var v5 int32
					v5 = int32(in.Int32())
					out.Alliance = append(out.Alliance, v5)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "character":
			if in.IsNull() {
				in.Skip()
				out.Character = nil
			} else {
				in.Delim('[')
				if out.Character == nil {
					if !in.IsDelim(']') {
						out.Character = make([]int32, 0, 16)
					} else {
						out.Character = []int32{}
					}
				} else {
					out.Character = (out.Character)[:0]
				}
				for !in.IsDelim(']') {
					var v6 int32
					v6 = int32(in.Int32())
					out.Character = append(out.Character, v6)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "constellation":
			if in.IsNull() {
				in.Skip()
				out.Constellation = nil
			} else {
				in.Delim('[')
				if out.Constellation == nil {
					if !in.IsDelim(']') {
						out.Constellation = make([]int32, 0, 16)
					} else {
						out.Constellation = []int32{}
					}
				} else {
					out.Constellation = (out.Constellation)[:0]
				}
				for !in.IsDelim(']') {
					var v7 int32
					v7 = int32(in.Int32())
					out.Constellation = append(out.Constellation, v7)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "corporation":
			if in.IsNull() {
				in.Skip()
				out.Corporation = nil
			} else {
				in.Delim('[')
				if out.Corporation == nil {
					if !in.IsDelim(']') {
						out.Corporation = make([]int32, 0, 16)
					} else {
						out.Corporation = []int32{}
					}
				} else {
					out.Corporation = (out.Corporation)[:0]
				}
				for !in.IsDelim(']') {
					var v8 int32
					v8 = int32(in.Int32())
					out.Corporation = append(out.Corporation, v8)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "faction":
			if in.IsNull() {
				in.Skip()
				out.Faction = nil
			} else {
				in.Delim('[')
				if out.Faction == nil {
					if !in.IsDelim(']') {
						out.Faction = make([]int32, 0, 16)
					} else {
						out.Faction = []int32{}
					}
				} else {
					out.Faction = (out.Faction)[:0]
				}
				for !in.IsDelim(']') {
					var v9 int32
					v9 = int32(in.Int32())
					out.Faction = append(out.Faction, v9)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "inventory_type":
			if in.IsNull() {
				in.Skip()
				out.InventoryType = nil
			} else {
				in.Delim('[')
				if out.InventoryType == nil {
					if !in.IsDelim(']') {
						out.InventoryType = make([]int32, 0, 16)
					} else {
						out.InventoryType = []int32{}
					}
				} else {
					out.InventoryType = (out.InventoryType)[:0]
				}
				for !in.IsDelim(']') {
					var v10 int32
					v10 = int32(in.Int32())
					out.InventoryType = append(out.InventoryType, v10)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "region":
			if in.IsNull() {
				in.Skip()
				out.Region = nil
			} else {
				in.Delim('[')
				if out.Region == nil {
					if !in.IsDelim(']') {
						out.Region = make([]int32, 0, 16)
					} else {
						out.Region = []int32{}
					}
				} else {
					out.Region = (out.Region)[:0]
				}
				for !in.IsDelim(']') {
					var v11 int32
					v11 = int32(in.Int32())
					out.Region = append(out.Region, v11)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "solar_system":
			if in.IsNull() {
				in.Skip()
				out.SolarSystem = nil
			} else {
				in.Delim('[')
				if out.SolarSystem == nil {
					if !in.IsDelim(']') {
						out.SolarSystem = make([]int32, 0, 16)
					} else {
						out.SolarSystem = []int32{}
					}
				} else {
					out.SolarSystem = (out.SolarSystem)[:0]
				}
				for !in.IsDelim(']') {
					var v12 int32
					v12 = int32(in.Int32())
					out.SolarSystem = append(out.SolarSystem, v12)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "station":
			if in.IsNull() {
				in.Skip()
				out.Station = nil
			} else {
				in.Delim('[')
				if out.Station == nil {
					if !in.IsDelim(']') {
						out.Station = make([]int32, 0, 16)
					} else {
						out.Station = []int32{}
					}
				} else {
					out.Station = (out.Station)[:0]
				}
				for !in.IsDelim(']') {
					var v13 int32
					v13 = int32(in.Int32())
					out.Station = append(out.Station, v13)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "structure":
			if in.IsNull() {
				in.Skip()
				out.Structure = nil
			} else {
				in.Delim('[')
				if out.Structure == nil {
					if !in.IsDelim(']') {
						out.Structure = make([]int64, 0, 8)
					} else {
						out.Structure = []int64{}
					}
				} else {
					out.Structure = (out.Structure)[:0]
				}
				for !in.IsDelim(']') {
					var v14 int64
					v14 = int64(in.Int64())
					out.Structure = append(out.Structure, v14)
					in.WantComma()
				}
				in.Delim(']')
			}
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
func easyjsonDebc599dEncodeGithubComAntihaxGoesiEsi1(out *jwriter.Writer, in GetCharactersCharacterIdSearchOk) {
	out.RawByte('{')
	first := true
	_ = first
	if len(in.Agent) != 0 {
		const prefix string = ",\"agent\":"
		first = false
		out.RawString(prefix[1:])
		{
			out.RawByte('[')
			for v15, v16 := range in.Agent {
				if v15 > 0 {
					out.RawByte(',')
				}
				out.Int32(int32(v16))
			}
			out.RawByte(']')
		}
	}
	if len(in.Alliance) != 0 {
		const prefix string = ",\"alliance\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		{
			out.RawByte('[')
			for v17, v18 := range in.Alliance {
				if v17 > 0 {
					out.RawByte(',')
				}
				out.Int32(int32(v18))
			}
			out.RawByte(']')
		}
	}
	if len(in.Character) != 0 {
		const prefix string = ",\"character\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		{
			out.RawByte('[')
			for v19, v20 := range in.Character {
				if v19 > 0 {
					out.RawByte(',')
				}
				out.Int32(int32(v20))
			}
			out.RawByte(']')
		}
	}
	if len(in.Constellation) != 0 {
		const prefix string = ",\"constellation\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		{
			out.RawByte('[')
			for v21, v22 := range in.Constellation {
				if v21 > 0 {
					out.RawByte(',')
				}
				out.Int32(int32(v22))
			}
			out.RawByte(']')
		}
	}
	if len(in.Corporation) != 0 {
		const prefix string = ",\"corporation\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		{
			out.RawByte('[')
			for v23, v24 := range in.Corporation {
				if v23 > 0 {
					out.RawByte(',')
				}
				out.Int32(int32(v24))
			}
			out.RawByte(']')
		}
	}
	if len(in.Faction) != 0 {
		const prefix string = ",\"faction\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		{
			out.RawByte('[')
			for v25, v26 := range in.Faction {
				if v25 > 0 {
					out.RawByte(',')
				}
				out.Int32(int32(v26))
			}
			out.RawByte(']')
		}
	}
	if len(in.InventoryType) != 0 {
		const prefix string = ",\"inventory_type\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		{
			out.RawByte('[')
			for v27, v28 := range in.InventoryType {
				if v27 > 0 {
					out.RawByte(',')
				}
				out.Int32(int32(v28))
			}
			out.RawByte(']')
		}
	}
	if len(in.Region) != 0 {
		const prefix string = ",\"region\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		{
			out.RawByte('[')
			for v29, v30 := range in.Region {
				if v29 > 0 {
					out.RawByte(',')
				}
				out.Int32(int32(v30))
			}
			out.RawByte(']')
		}
	}
	if len(in.SolarSystem) != 0 {
		const prefix string = ",\"solar_system\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		{
			out.RawByte('[')
			for v31, v32 := range in.SolarSystem {
				if v31 > 0 {
					out.RawByte(',')
				}
				out.Int32(int32(v32))
			}
			out.RawByte(']')
		}
	}
	if len(in.Station) != 0 {
		const prefix string = ",\"station\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		{
			out.RawByte('[')
			for v33, v34 := range in.Station {
				if v33 > 0 {
					out.RawByte(',')
				}
				out.Int32(int32(v34))
			}
			out.RawByte(']')
		}
	}
	if len(in.Structure) != 0 {
		const prefix string = ",\"structure\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		{
			out.RawByte('[')
			for v35, v36 := range in.Structure {
				if v35 > 0 {
					out.RawByte(',')
				}
				out.Int64(int64(v36))
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v GetCharactersCharacterIdSearchOk) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonDebc599dEncodeGithubComAntihaxGoesiEsi1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GetCharactersCharacterIdSearchOk) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonDebc599dEncodeGithubComAntihaxGoesiEsi1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GetCharactersCharacterIdSearchOk) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonDebc599dDecodeGithubComAntihaxGoesiEsi1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GetCharactersCharacterIdSearchOk) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonDebc599dDecodeGithubComAntihaxGoesiEsi1(l, v)
}
