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

func easyjson7a98a6fdDecodeGithubComAntihaxGoesiEsi(in *jlexer.Lexer, out *GetFwLeaderboardsCharactersOkList) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		in.Skip()
		*out = nil
	} else {
		in.Delim('[')
		if *out == nil {
			if !in.IsDelim(']') {
				*out = make(GetFwLeaderboardsCharactersOkList, 0, 1)
			} else {
				*out = GetFwLeaderboardsCharactersOkList{}
			}
		} else {
			*out = (*out)[:0]
		}
		for !in.IsDelim(']') {
			var v1 GetFwLeaderboardsCharactersOk
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
func easyjson7a98a6fdEncodeGithubComAntihaxGoesiEsi(out *jwriter.Writer, in GetFwLeaderboardsCharactersOkList) {
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
func (v GetFwLeaderboardsCharactersOkList) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson7a98a6fdEncodeGithubComAntihaxGoesiEsi(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GetFwLeaderboardsCharactersOkList) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson7a98a6fdEncodeGithubComAntihaxGoesiEsi(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GetFwLeaderboardsCharactersOkList) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson7a98a6fdDecodeGithubComAntihaxGoesiEsi(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GetFwLeaderboardsCharactersOkList) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson7a98a6fdDecodeGithubComAntihaxGoesiEsi(l, v)
}
func easyjson7a98a6fdDecodeGithubComAntihaxGoesiEsi1(in *jlexer.Lexer, out *GetFwLeaderboardsCharactersOk) {
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
		case "kills":
			easyjson7a98a6fdDecodeGithubComAntihaxGoesiEsi2(in, &out.Kills)
		case "victory_points":
			(out.VictoryPoints).UnmarshalEasyJSON(in)
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
func easyjson7a98a6fdEncodeGithubComAntihaxGoesiEsi1(out *jwriter.Writer, in GetFwLeaderboardsCharactersOk) {
	out.RawByte('{')
	first := true
	_ = first
	if true {
		const prefix string = ",\"kills\":"
		first = false
		out.RawString(prefix[1:])
		easyjson7a98a6fdEncodeGithubComAntihaxGoesiEsi2(out, in.Kills)
	}
	if true {
		const prefix string = ",\"victory_points\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		(in.VictoryPoints).MarshalEasyJSON(out)
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v GetFwLeaderboardsCharactersOk) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson7a98a6fdEncodeGithubComAntihaxGoesiEsi1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GetFwLeaderboardsCharactersOk) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson7a98a6fdEncodeGithubComAntihaxGoesiEsi1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GetFwLeaderboardsCharactersOk) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson7a98a6fdDecodeGithubComAntihaxGoesiEsi1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GetFwLeaderboardsCharactersOk) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson7a98a6fdDecodeGithubComAntihaxGoesiEsi1(l, v)
}
func easyjson7a98a6fdDecodeGithubComAntihaxGoesiEsi2(in *jlexer.Lexer, out *GetFwLeaderboardsCharactersKills) {
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
		case "active_total":
			if in.IsNull() {
				in.Skip()
				out.ActiveTotal = nil
			} else {
				in.Delim('[')
				if out.ActiveTotal == nil {
					if !in.IsDelim(']') {
						out.ActiveTotal = make([]GetFwLeaderboardsCharactersActiveTotalActiveTotal, 0, 8)
					} else {
						out.ActiveTotal = []GetFwLeaderboardsCharactersActiveTotalActiveTotal{}
					}
				} else {
					out.ActiveTotal = (out.ActiveTotal)[:0]
				}
				for !in.IsDelim(']') {
					var v4 GetFwLeaderboardsCharactersActiveTotalActiveTotal
					easyjson7a98a6fdDecodeGithubComAntihaxGoesiEsi3(in, &v4)
					out.ActiveTotal = append(out.ActiveTotal, v4)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "last_week":
			if in.IsNull() {
				in.Skip()
				out.LastWeek = nil
			} else {
				in.Delim('[')
				if out.LastWeek == nil {
					if !in.IsDelim(']') {
						out.LastWeek = make([]GetFwLeaderboardsCharactersLastWeekLastWeek, 0, 8)
					} else {
						out.LastWeek = []GetFwLeaderboardsCharactersLastWeekLastWeek{}
					}
				} else {
					out.LastWeek = (out.LastWeek)[:0]
				}
				for !in.IsDelim(']') {
					var v5 GetFwLeaderboardsCharactersLastWeekLastWeek
					easyjson7a98a6fdDecodeGithubComAntihaxGoesiEsi4(in, &v5)
					out.LastWeek = append(out.LastWeek, v5)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "yesterday":
			if in.IsNull() {
				in.Skip()
				out.Yesterday = nil
			} else {
				in.Delim('[')
				if out.Yesterday == nil {
					if !in.IsDelim(']') {
						out.Yesterday = make([]GetFwLeaderboardsCharactersYesterdayYesterday, 0, 8)
					} else {
						out.Yesterday = []GetFwLeaderboardsCharactersYesterdayYesterday{}
					}
				} else {
					out.Yesterday = (out.Yesterday)[:0]
				}
				for !in.IsDelim(']') {
					var v6 GetFwLeaderboardsCharactersYesterdayYesterday
					easyjson7a98a6fdDecodeGithubComAntihaxGoesiEsi5(in, &v6)
					out.Yesterday = append(out.Yesterday, v6)
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
func easyjson7a98a6fdEncodeGithubComAntihaxGoesiEsi2(out *jwriter.Writer, in GetFwLeaderboardsCharactersKills) {
	out.RawByte('{')
	first := true
	_ = first
	if len(in.ActiveTotal) != 0 {
		const prefix string = ",\"active_total\":"
		first = false
		out.RawString(prefix[1:])
		{
			out.RawByte('[')
			for v7, v8 := range in.ActiveTotal {
				if v7 > 0 {
					out.RawByte(',')
				}
				easyjson7a98a6fdEncodeGithubComAntihaxGoesiEsi3(out, v8)
			}
			out.RawByte(']')
		}
	}
	if len(in.LastWeek) != 0 {
		const prefix string = ",\"last_week\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		{
			out.RawByte('[')
			for v9, v10 := range in.LastWeek {
				if v9 > 0 {
					out.RawByte(',')
				}
				easyjson7a98a6fdEncodeGithubComAntihaxGoesiEsi4(out, v10)
			}
			out.RawByte(']')
		}
	}
	if len(in.Yesterday) != 0 {
		const prefix string = ",\"yesterday\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		{
			out.RawByte('[')
			for v11, v12 := range in.Yesterday {
				if v11 > 0 {
					out.RawByte(',')
				}
				easyjson7a98a6fdEncodeGithubComAntihaxGoesiEsi5(out, v12)
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}
func easyjson7a98a6fdDecodeGithubComAntihaxGoesiEsi5(in *jlexer.Lexer, out *GetFwLeaderboardsCharactersYesterdayYesterday) {
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
		case "amount":
			out.Amount = int32(in.Int32())
		case "character_id":
			out.CharacterId = int32(in.Int32())
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
func easyjson7a98a6fdEncodeGithubComAntihaxGoesiEsi5(out *jwriter.Writer, in GetFwLeaderboardsCharactersYesterdayYesterday) {
	out.RawByte('{')
	first := true
	_ = first
	if in.Amount != 0 {
		const prefix string = ",\"amount\":"
		first = false
		out.RawString(prefix[1:])
		out.Int32(int32(in.Amount))
	}
	if in.CharacterId != 0 {
		const prefix string = ",\"character_id\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int32(int32(in.CharacterId))
	}
	out.RawByte('}')
}
func easyjson7a98a6fdDecodeGithubComAntihaxGoesiEsi4(in *jlexer.Lexer, out *GetFwLeaderboardsCharactersLastWeekLastWeek) {
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
		case "amount":
			out.Amount = int32(in.Int32())
		case "character_id":
			out.CharacterId = int32(in.Int32())
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
func easyjson7a98a6fdEncodeGithubComAntihaxGoesiEsi4(out *jwriter.Writer, in GetFwLeaderboardsCharactersLastWeekLastWeek) {
	out.RawByte('{')
	first := true
	_ = first
	if in.Amount != 0 {
		const prefix string = ",\"amount\":"
		first = false
		out.RawString(prefix[1:])
		out.Int32(int32(in.Amount))
	}
	if in.CharacterId != 0 {
		const prefix string = ",\"character_id\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int32(int32(in.CharacterId))
	}
	out.RawByte('}')
}
func easyjson7a98a6fdDecodeGithubComAntihaxGoesiEsi3(in *jlexer.Lexer, out *GetFwLeaderboardsCharactersActiveTotalActiveTotal) {
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
		case "amount":
			out.Amount = int32(in.Int32())
		case "character_id":
			out.CharacterId = int32(in.Int32())
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
func easyjson7a98a6fdEncodeGithubComAntihaxGoesiEsi3(out *jwriter.Writer, in GetFwLeaderboardsCharactersActiveTotalActiveTotal) {
	out.RawByte('{')
	first := true
	_ = first
	if in.Amount != 0 {
		const prefix string = ",\"amount\":"
		first = false
		out.RawString(prefix[1:])
		out.Int32(int32(in.Amount))
	}
	if in.CharacterId != 0 {
		const prefix string = ",\"character_id\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int32(int32(in.CharacterId))
	}
	out.RawByte('}')
}
