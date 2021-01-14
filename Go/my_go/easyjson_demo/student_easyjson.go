// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package proto

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

func easyjsonB83d7b77DecodeEasyjson(in *jlexer.Lexer, out *Student) {
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
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "Id":
			out.Id = int(in.Int())
		case "Name":
			out.Name = string(in.String())
		case "School":
			(out.School).UnmarshalEasyJSON(in)
		case "Birthday":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.Birthday).UnmarshalJSON(data))
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
func easyjsonB83d7b77EncodeEasyjson(out *jwriter.Writer, in Student) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"Id\":"
		out.RawString(prefix[1:])
		out.Int(int(in.Id))
	}
	{
		const prefix string = ",\"Name\":"
		out.RawString(prefix)
		out.String(string(in.Name))
	}
	{
		const prefix string = ",\"School\":"
		out.RawString(prefix)
		(in.School).MarshalEasyJSON(out)
	}
	{
		const prefix string = ",\"Birthday\":"
		out.RawString(prefix)
		out.Raw((in.Birthday).MarshalJSON())
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Student) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonB83d7b77EncodeEasyjson(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Student) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonB83d7b77EncodeEasyjson(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Student) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonB83d7b77DecodeEasyjson(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Student) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonB83d7b77DecodeEasyjson(l, v)
}
func easyjsonB83d7b77DecodeEasyjson1(in *jlexer.Lexer, out *School) {
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
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "Name":
			out.Name = string(in.String())
		case "Addr":
			out.Addr = string(in.String())
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
func easyjsonB83d7b77EncodeEasyjson1(out *jwriter.Writer, in School) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"Name\":"
		out.RawString(prefix[1:])
		out.String(string(in.Name))
	}
	{
		const prefix string = ",\"Addr\":"
		out.RawString(prefix)
		out.String(string(in.Addr))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v School) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonB83d7b77EncodeEasyjson1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v School) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonB83d7b77EncodeEasyjson1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *School) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonB83d7b77DecodeEasyjson1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *School) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonB83d7b77DecodeEasyjson1(l, v)
}