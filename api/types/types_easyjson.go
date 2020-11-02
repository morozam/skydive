// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package types

import (
	json "encoding/json"
	easyjson "github.com/mailru/easyjson"
	jlexer "github.com/mailru/easyjson/jlexer"
	jwriter "github.com/mailru/easyjson/jwriter"
	graph "github.com/skydive-project/skydive/graffiti/graph"
)

// suppress unused package warning
var (
	_ *json.RawMessage
	_ *jlexer.Lexer
	_ *jwriter.Writer
	_ easyjson.Marshaler
)

func easyjson6601e8cdDecodeGithubComSkydiveProjectSkydiveApiTypes(in *jlexer.Lexer, out *WorkflowParam) {
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
		case "Description":
			out.Description = string(in.String())
		case "Type":
			out.Type = string(in.String())
		case "Default":
			if m, ok := out.Default.(easyjson.Unmarshaler); ok {
				m.UnmarshalEasyJSON(in)
			} else if m, ok := out.Default.(json.Unmarshaler); ok {
				_ = m.UnmarshalJSON(in.Raw())
			} else {
				out.Default = in.Interface()
			}
		case "Values":
			if in.IsNull() {
				in.Skip()
				out.Values = nil
			} else {
				in.Delim('[')
				if out.Values == nil {
					if !in.IsDelim(']') {
						out.Values = make([]WorkflowChoice, 0, 2)
					} else {
						out.Values = []WorkflowChoice{}
					}
				} else {
					out.Values = (out.Values)[:0]
				}
				for !in.IsDelim(']') {
					var v1 WorkflowChoice
					(v1).UnmarshalEasyJSON(in)
					out.Values = append(out.Values, v1)
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
func easyjson6601e8cdEncodeGithubComSkydiveProjectSkydiveApiTypes(out *jwriter.Writer, in WorkflowParam) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"Name\":"
		out.RawString(prefix[1:])
		out.String(string(in.Name))
	}
	{
		const prefix string = ",\"Description\":"
		out.RawString(prefix)
		out.String(string(in.Description))
	}
	{
		const prefix string = ",\"Type\":"
		out.RawString(prefix)
		out.String(string(in.Type))
	}
	{
		const prefix string = ",\"Default\":"
		out.RawString(prefix)
		if m, ok := in.Default.(easyjson.Marshaler); ok {
			m.MarshalEasyJSON(out)
		} else if m, ok := in.Default.(json.Marshaler); ok {
			out.Raw(m.MarshalJSON())
		} else {
			out.Raw(json.Marshal(in.Default))
		}
	}
	{
		const prefix string = ",\"Values\":"
		out.RawString(prefix)
		if in.Values == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v2, v3 := range in.Values {
				if v2 > 0 {
					out.RawByte(',')
				}
				(v3).MarshalEasyJSON(out)
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v WorkflowParam) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson6601e8cdEncodeGithubComSkydiveProjectSkydiveApiTypes(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v WorkflowParam) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson6601e8cdEncodeGithubComSkydiveProjectSkydiveApiTypes(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *WorkflowParam) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson6601e8cdDecodeGithubComSkydiveProjectSkydiveApiTypes(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *WorkflowParam) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson6601e8cdDecodeGithubComSkydiveProjectSkydiveApiTypes(l, v)
}
func easyjson6601e8cdDecodeGithubComSkydiveProjectSkydiveApiTypes1(in *jlexer.Lexer, out *WorkflowChoice) {
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
		case "Value":
			out.Value = string(in.String())
		case "Description":
			out.Description = string(in.String())
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
func easyjson6601e8cdEncodeGithubComSkydiveProjectSkydiveApiTypes1(out *jwriter.Writer, in WorkflowChoice) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"Value\":"
		out.RawString(prefix[1:])
		out.String(string(in.Value))
	}
	{
		const prefix string = ",\"Description\":"
		out.RawString(prefix)
		out.String(string(in.Description))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v WorkflowChoice) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson6601e8cdEncodeGithubComSkydiveProjectSkydiveApiTypes1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v WorkflowChoice) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson6601e8cdEncodeGithubComSkydiveProjectSkydiveApiTypes1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *WorkflowChoice) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson6601e8cdDecodeGithubComSkydiveProjectSkydiveApiTypes1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *WorkflowChoice) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson6601e8cdDecodeGithubComSkydiveProjectSkydiveApiTypes1(l, v)
}
func easyjson6601e8cdDecodeGithubComSkydiveProjectSkydiveApiTypes2(in *jlexer.Lexer, out *Workflow) {
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
		case "Title":
			out.Title = string(in.String())
		case "Abstract":
			out.Abstract = string(in.String())
		case "Description":
			out.Description = string(in.String())
		case "Parameters":
			if in.IsNull() {
				in.Skip()
				out.Parameters = nil
			} else {
				in.Delim('[')
				if out.Parameters == nil {
					if !in.IsDelim(']') {
						out.Parameters = make([]WorkflowParam, 0, 0)
					} else {
						out.Parameters = []WorkflowParam{}
					}
				} else {
					out.Parameters = (out.Parameters)[:0]
				}
				for !in.IsDelim(']') {
					var v4 WorkflowParam
					(v4).UnmarshalEasyJSON(in)
					out.Parameters = append(out.Parameters, v4)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "Source":
			out.Source = string(in.String())
		case "UUID":
			out.UUID = string(in.String())
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
func easyjson6601e8cdEncodeGithubComSkydiveProjectSkydiveApiTypes2(out *jwriter.Writer, in Workflow) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"Name\":"
		out.RawString(prefix[1:])
		out.String(string(in.Name))
	}
	{
		const prefix string = ",\"Title\":"
		out.RawString(prefix)
		out.String(string(in.Title))
	}
	{
		const prefix string = ",\"Abstract\":"
		out.RawString(prefix)
		out.String(string(in.Abstract))
	}
	{
		const prefix string = ",\"Description\":"
		out.RawString(prefix)
		out.String(string(in.Description))
	}
	{
		const prefix string = ",\"Parameters\":"
		out.RawString(prefix)
		if in.Parameters == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v5, v6 := range in.Parameters {
				if v5 > 0 {
					out.RawByte(',')
				}
				(v6).MarshalEasyJSON(out)
			}
			out.RawByte(']')
		}
	}
	{
		const prefix string = ",\"Source\":"
		out.RawString(prefix)
		out.String(string(in.Source))
	}
	{
		const prefix string = ",\"UUID\":"
		out.RawString(prefix)
		out.String(string(in.UUID))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Workflow) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson6601e8cdEncodeGithubComSkydiveProjectSkydiveApiTypes2(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Workflow) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson6601e8cdEncodeGithubComSkydiveProjectSkydiveApiTypes2(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Workflow) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson6601e8cdDecodeGithubComSkydiveProjectSkydiveApiTypes2(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Workflow) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson6601e8cdDecodeGithubComSkydiveProjectSkydiveApiTypes2(l, v)
}
func easyjson6601e8cdDecodeGithubComSkydiveProjectSkydiveApiTypes3(in *jlexer.Lexer, out *TopologyParams) {
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
		case "GremlinQuery":
			out.GremlinQuery = string(in.String())
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
func easyjson6601e8cdEncodeGithubComSkydiveProjectSkydiveApiTypes3(out *jwriter.Writer, in TopologyParams) {
	out.RawByte('{')
	first := true
	_ = first
	if in.GremlinQuery != "" {
		const prefix string = ",\"GremlinQuery\":"
		first = false
		out.RawString(prefix[1:])
		out.String(string(in.GremlinQuery))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v TopologyParams) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson6601e8cdEncodeGithubComSkydiveProjectSkydiveApiTypes3(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v TopologyParams) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson6601e8cdEncodeGithubComSkydiveProjectSkydiveApiTypes3(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *TopologyParams) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson6601e8cdDecodeGithubComSkydiveProjectSkydiveApiTypes3(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *TopologyParams) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson6601e8cdDecodeGithubComSkydiveProjectSkydiveApiTypes3(l, v)
}
func easyjson6601e8cdDecodeGithubComSkydiveProjectSkydiveApiTypes4(in *jlexer.Lexer, out *PacketInjection) {
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
		case "Src":
			out.Src = string(in.String())
		case "Dst":
			out.Dst = string(in.String())
		case "SrcIP":
			out.SrcIP = string(in.String())
		case "DstIP":
			out.DstIP = string(in.String())
		case "SrcMAC":
			out.SrcMAC = string(in.String())
		case "DstMAC":
			out.DstMAC = string(in.String())
		case "SrcPort":
			out.SrcPort = uint16(in.Uint16())
		case "DstPort":
			out.DstPort = uint16(in.Uint16())
		case "Type":
			out.Type = string(in.String())
		case "Payload":
			out.Payload = string(in.String())
		case "ICMPID":
			out.ICMPID = uint16(in.Uint16())
		case "Count":
			out.Count = uint64(in.Uint64())
		case "Interval":
			out.Interval = uint64(in.Uint64())
		case "Mode":
			out.Mode = string(in.String())
		case "IncrementPayload":
			out.IncrementPayload = int64(in.Int64())
		case "StartTime":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.StartTime).UnmarshalJSON(data))
			}
		case "Pcap":
			if in.IsNull() {
				in.Skip()
				out.Pcap = nil
			} else {
				out.Pcap = in.Bytes()
			}
		case "TTL":
			out.TTL = uint8(in.Uint8())
		case "UUID":
			out.UUID = string(in.String())
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
func easyjson6601e8cdEncodeGithubComSkydiveProjectSkydiveApiTypes4(out *jwriter.Writer, in PacketInjection) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"Src\":"
		out.RawString(prefix[1:])
		out.String(string(in.Src))
	}
	{
		const prefix string = ",\"Dst\":"
		out.RawString(prefix)
		out.String(string(in.Dst))
	}
	{
		const prefix string = ",\"SrcIP\":"
		out.RawString(prefix)
		out.String(string(in.SrcIP))
	}
	{
		const prefix string = ",\"DstIP\":"
		out.RawString(prefix)
		out.String(string(in.DstIP))
	}
	{
		const prefix string = ",\"SrcMAC\":"
		out.RawString(prefix)
		out.String(string(in.SrcMAC))
	}
	{
		const prefix string = ",\"DstMAC\":"
		out.RawString(prefix)
		out.String(string(in.DstMAC))
	}
	{
		const prefix string = ",\"SrcPort\":"
		out.RawString(prefix)
		out.Uint16(uint16(in.SrcPort))
	}
	{
		const prefix string = ",\"DstPort\":"
		out.RawString(prefix)
		out.Uint16(uint16(in.DstPort))
	}
	{
		const prefix string = ",\"Type\":"
		out.RawString(prefix)
		out.String(string(in.Type))
	}
	{
		const prefix string = ",\"Payload\":"
		out.RawString(prefix)
		out.String(string(in.Payload))
	}
	{
		const prefix string = ",\"ICMPID\":"
		out.RawString(prefix)
		out.Uint16(uint16(in.ICMPID))
	}
	{
		const prefix string = ",\"Count\":"
		out.RawString(prefix)
		out.Uint64(uint64(in.Count))
	}
	{
		const prefix string = ",\"Interval\":"
		out.RawString(prefix)
		out.Uint64(uint64(in.Interval))
	}
	{
		const prefix string = ",\"Mode\":"
		out.RawString(prefix)
		out.String(string(in.Mode))
	}
	{
		const prefix string = ",\"IncrementPayload\":"
		out.RawString(prefix)
		out.Int64(int64(in.IncrementPayload))
	}
	{
		const prefix string = ",\"StartTime\":"
		out.RawString(prefix)
		out.Raw((in.StartTime).MarshalJSON())
	}
	{
		const prefix string = ",\"Pcap\":"
		out.RawString(prefix)
		out.Base64Bytes(in.Pcap)
	}
	{
		const prefix string = ",\"TTL\":"
		out.RawString(prefix)
		out.Uint8(uint8(in.TTL))
	}
	{
		const prefix string = ",\"UUID\":"
		out.RawString(prefix)
		out.String(string(in.UUID))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v PacketInjection) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson6601e8cdEncodeGithubComSkydiveProjectSkydiveApiTypes4(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v PacketInjection) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson6601e8cdEncodeGithubComSkydiveProjectSkydiveApiTypes4(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *PacketInjection) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson6601e8cdDecodeGithubComSkydiveProjectSkydiveApiTypes4(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *PacketInjection) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson6601e8cdDecodeGithubComSkydiveProjectSkydiveApiTypes4(l, v)
}
func easyjson6601e8cdDecodeGithubComSkydiveProjectSkydiveApiTypes5(in *jlexer.Lexer, out *NodeRule) {
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
		case "Description":
			out.Description = string(in.String())
		case "Metadata":
			if in.IsNull() {
				in.Skip()
			} else {
				in.Delim('{')
				out.Metadata = make(graph.Metadata)
				for !in.IsDelim('}') {
					key := string(in.String())
					in.WantColon()
					var v10 interface{}
					if m, ok := v10.(easyjson.Unmarshaler); ok {
						m.UnmarshalEasyJSON(in)
					} else if m, ok := v10.(json.Unmarshaler); ok {
						_ = m.UnmarshalJSON(in.Raw())
					} else {
						v10 = in.Interface()
					}
					(out.Metadata)[key] = v10
					in.WantComma()
				}
				in.Delim('}')
			}
		case "Action":
			out.Action = string(in.String())
		case "Query":
			out.Query = string(in.String())
		case "UUID":
			out.UUID = string(in.String())
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
func easyjson6601e8cdEncodeGithubComSkydiveProjectSkydiveApiTypes5(out *jwriter.Writer, in NodeRule) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"Name\":"
		out.RawString(prefix[1:])
		out.String(string(in.Name))
	}
	{
		const prefix string = ",\"Description\":"
		out.RawString(prefix)
		out.String(string(in.Description))
	}
	{
		const prefix string = ",\"Metadata\":"
		out.RawString(prefix)
		if in.Metadata == nil && (out.Flags&jwriter.NilMapAsEmpty) == 0 {
			out.RawString(`null`)
		} else {
			out.RawByte('{')
			v11First := true
			for v11Name, v11Value := range in.Metadata {
				if v11First {
					v11First = false
				} else {
					out.RawByte(',')
				}
				out.String(string(v11Name))
				out.RawByte(':')
				if m, ok := v11Value.(easyjson.Marshaler); ok {
					m.MarshalEasyJSON(out)
				} else if m, ok := v11Value.(json.Marshaler); ok {
					out.Raw(m.MarshalJSON())
				} else {
					out.Raw(json.Marshal(v11Value))
				}
			}
			out.RawByte('}')
		}
	}
	{
		const prefix string = ",\"Action\":"
		out.RawString(prefix)
		out.String(string(in.Action))
	}
	{
		const prefix string = ",\"Query\":"
		out.RawString(prefix)
		out.String(string(in.Query))
	}
	{
		const prefix string = ",\"UUID\":"
		out.RawString(prefix)
		out.String(string(in.UUID))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v NodeRule) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson6601e8cdEncodeGithubComSkydiveProjectSkydiveApiTypes5(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v NodeRule) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson6601e8cdEncodeGithubComSkydiveProjectSkydiveApiTypes5(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *NodeRule) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson6601e8cdDecodeGithubComSkydiveProjectSkydiveApiTypes5(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *NodeRule) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson6601e8cdDecodeGithubComSkydiveProjectSkydiveApiTypes5(l, v)
}
func easyjson6601e8cdDecodeGithubComSkydiveProjectSkydiveApiTypes6(in *jlexer.Lexer, out *Node) {
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
		case "ID":
			out.ID = graph.Identifier(in.String())
		case "Metadata":
			if in.IsNull() {
				in.Skip()
			} else {
				in.Delim('{')
				out.Metadata = make(graph.Metadata)
				for !in.IsDelim('}') {
					key := string(in.String())
					in.WantColon()
					var v12 interface{}
					if m, ok := v12.(easyjson.Unmarshaler); ok {
						m.UnmarshalEasyJSON(in)
					} else if m, ok := v12.(json.Unmarshaler); ok {
						_ = m.UnmarshalJSON(in.Raw())
					} else {
						v12 = in.Interface()
					}
					(out.Metadata)[key] = v12
					in.WantComma()
				}
				in.Delim('}')
			}
		case "Host":
			out.Host = string(in.String())
		case "Origin":
			out.Origin = string(in.String())
		case "CreatedAt":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.CreatedAt).UnmarshalJSON(data))
			}
		case "UpdatedAt":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.UpdatedAt).UnmarshalJSON(data))
			}
		case "DeletedAt":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.DeletedAt).UnmarshalJSON(data))
			}
		case "Revision":
			out.Revision = int64(in.Int64())
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
func easyjson6601e8cdEncodeGithubComSkydiveProjectSkydiveApiTypes6(out *jwriter.Writer, in Node) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"ID\":"
		out.RawString(prefix[1:])
		out.String(string(in.ID))
	}
	{
		const prefix string = ",\"Metadata\":"
		out.RawString(prefix)
		if in.Metadata == nil && (out.Flags&jwriter.NilMapAsEmpty) == 0 {
			out.RawString(`null`)
		} else {
			out.RawByte('{')
			v13First := true
			for v13Name, v13Value := range in.Metadata {
				if v13First {
					v13First = false
				} else {
					out.RawByte(',')
				}
				out.String(string(v13Name))
				out.RawByte(':')
				if m, ok := v13Value.(easyjson.Marshaler); ok {
					m.MarshalEasyJSON(out)
				} else if m, ok := v13Value.(json.Marshaler); ok {
					out.Raw(m.MarshalJSON())
				} else {
					out.Raw(json.Marshal(v13Value))
				}
			}
			out.RawByte('}')
		}
	}
	{
		const prefix string = ",\"Host\":"
		out.RawString(prefix)
		out.String(string(in.Host))
	}
	{
		const prefix string = ",\"Origin\":"
		out.RawString(prefix)
		out.String(string(in.Origin))
	}
	{
		const prefix string = ",\"CreatedAt\":"
		out.RawString(prefix)
		out.Raw((in.CreatedAt).MarshalJSON())
	}
	{
		const prefix string = ",\"UpdatedAt\":"
		out.RawString(prefix)
		out.Raw((in.UpdatedAt).MarshalJSON())
	}
	if true {
		const prefix string = ",\"DeletedAt\":"
		out.RawString(prefix)
		out.Raw((in.DeletedAt).MarshalJSON())
	}
	{
		const prefix string = ",\"Revision\":"
		out.RawString(prefix)
		out.Int64(int64(in.Revision))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Node) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson6601e8cdEncodeGithubComSkydiveProjectSkydiveApiTypes6(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Node) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson6601e8cdEncodeGithubComSkydiveProjectSkydiveApiTypes6(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Node) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson6601e8cdDecodeGithubComSkydiveProjectSkydiveApiTypes6(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Node) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson6601e8cdDecodeGithubComSkydiveProjectSkydiveApiTypes6(l, v)
}
func easyjson6601e8cdDecodeGithubComSkydiveProjectSkydiveApiTypes7(in *jlexer.Lexer, out *EdgeRule) {
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
		case "Description":
			out.Description = string(in.String())
		case "Src":
			out.Src = string(in.String())
		case "Dst":
			out.Dst = string(in.String())
		case "Metadata":
			if in.IsNull() {
				in.Skip()
			} else {
				in.Delim('{')
				out.Metadata = make(graph.Metadata)
				for !in.IsDelim('}') {
					key := string(in.String())
					in.WantColon()
					var v14 interface{}
					if m, ok := v14.(easyjson.Unmarshaler); ok {
						m.UnmarshalEasyJSON(in)
					} else if m, ok := v14.(json.Unmarshaler); ok {
						_ = m.UnmarshalJSON(in.Raw())
					} else {
						v14 = in.Interface()
					}
					(out.Metadata)[key] = v14
					in.WantComma()
				}
				in.Delim('}')
			}
		case "UUID":
			out.UUID = string(in.String())
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
func easyjson6601e8cdEncodeGithubComSkydiveProjectSkydiveApiTypes7(out *jwriter.Writer, in EdgeRule) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"Name\":"
		out.RawString(prefix[1:])
		out.String(string(in.Name))
	}
	{
		const prefix string = ",\"Description\":"
		out.RawString(prefix)
		out.String(string(in.Description))
	}
	{
		const prefix string = ",\"Src\":"
		out.RawString(prefix)
		out.String(string(in.Src))
	}
	{
		const prefix string = ",\"Dst\":"
		out.RawString(prefix)
		out.String(string(in.Dst))
	}
	{
		const prefix string = ",\"Metadata\":"
		out.RawString(prefix)
		if in.Metadata == nil && (out.Flags&jwriter.NilMapAsEmpty) == 0 {
			out.RawString(`null`)
		} else {
			out.RawByte('{')
			v15First := true
			for v15Name, v15Value := range in.Metadata {
				if v15First {
					v15First = false
				} else {
					out.RawByte(',')
				}
				out.String(string(v15Name))
				out.RawByte(':')
				if m, ok := v15Value.(easyjson.Marshaler); ok {
					m.MarshalEasyJSON(out)
				} else if m, ok := v15Value.(json.Marshaler); ok {
					out.Raw(m.MarshalJSON())
				} else {
					out.Raw(json.Marshal(v15Value))
				}
			}
			out.RawByte('}')
		}
	}
	{
		const prefix string = ",\"UUID\":"
		out.RawString(prefix)
		out.String(string(in.UUID))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v EdgeRule) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson6601e8cdEncodeGithubComSkydiveProjectSkydiveApiTypes7(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v EdgeRule) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson6601e8cdEncodeGithubComSkydiveProjectSkydiveApiTypes7(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *EdgeRule) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson6601e8cdDecodeGithubComSkydiveProjectSkydiveApiTypes7(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *EdgeRule) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson6601e8cdDecodeGithubComSkydiveProjectSkydiveApiTypes7(l, v)
}
func easyjson6601e8cdDecodeGithubComSkydiveProjectSkydiveApiTypes8(in *jlexer.Lexer, out *Edge) {
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
		case "Parent":
			out.Parent = graph.Identifier(in.String())
		case "Child":
			out.Child = graph.Identifier(in.String())
		case "ID":
			out.ID = graph.Identifier(in.String())
		case "Metadata":
			if in.IsNull() {
				in.Skip()
			} else {
				in.Delim('{')
				out.Metadata = make(graph.Metadata)
				for !in.IsDelim('}') {
					key := string(in.String())
					in.WantColon()
					var v16 interface{}
					if m, ok := v16.(easyjson.Unmarshaler); ok {
						m.UnmarshalEasyJSON(in)
					} else if m, ok := v16.(json.Unmarshaler); ok {
						_ = m.UnmarshalJSON(in.Raw())
					} else {
						v16 = in.Interface()
					}
					(out.Metadata)[key] = v16
					in.WantComma()
				}
				in.Delim('}')
			}
		case "Host":
			out.Host = string(in.String())
		case "Origin":
			out.Origin = string(in.String())
		case "CreatedAt":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.CreatedAt).UnmarshalJSON(data))
			}
		case "UpdatedAt":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.UpdatedAt).UnmarshalJSON(data))
			}
		case "DeletedAt":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.DeletedAt).UnmarshalJSON(data))
			}
		case "Revision":
			out.Revision = int64(in.Int64())
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
func easyjson6601e8cdEncodeGithubComSkydiveProjectSkydiveApiTypes8(out *jwriter.Writer, in Edge) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"Parent\":"
		out.RawString(prefix[1:])
		out.String(string(in.Parent))
	}
	{
		const prefix string = ",\"Child\":"
		out.RawString(prefix)
		out.String(string(in.Child))
	}
	{
		const prefix string = ",\"ID\":"
		out.RawString(prefix)
		out.String(string(in.ID))
	}
	{
		const prefix string = ",\"Metadata\":"
		out.RawString(prefix)
		if in.Metadata == nil && (out.Flags&jwriter.NilMapAsEmpty) == 0 {
			out.RawString(`null`)
		} else {
			out.RawByte('{')
			v17First := true
			for v17Name, v17Value := range in.Metadata {
				if v17First {
					v17First = false
				} else {
					out.RawByte(',')
				}
				out.String(string(v17Name))
				out.RawByte(':')
				if m, ok := v17Value.(easyjson.Marshaler); ok {
					m.MarshalEasyJSON(out)
				} else if m, ok := v17Value.(json.Marshaler); ok {
					out.Raw(m.MarshalJSON())
				} else {
					out.Raw(json.Marshal(v17Value))
				}
			}
			out.RawByte('}')
		}
	}
	{
		const prefix string = ",\"Host\":"
		out.RawString(prefix)
		out.String(string(in.Host))
	}
	{
		const prefix string = ",\"Origin\":"
		out.RawString(prefix)
		out.String(string(in.Origin))
	}
	{
		const prefix string = ",\"CreatedAt\":"
		out.RawString(prefix)
		out.Raw((in.CreatedAt).MarshalJSON())
	}
	{
		const prefix string = ",\"UpdatedAt\":"
		out.RawString(prefix)
		out.Raw((in.UpdatedAt).MarshalJSON())
	}
	if true {
		const prefix string = ",\"DeletedAt\":"
		out.RawString(prefix)
		out.Raw((in.DeletedAt).MarshalJSON())
	}
	{
		const prefix string = ",\"Revision\":"
		out.RawString(prefix)
		out.Int64(int64(in.Revision))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Edge) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson6601e8cdEncodeGithubComSkydiveProjectSkydiveApiTypes8(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Edge) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson6601e8cdEncodeGithubComSkydiveProjectSkydiveApiTypes8(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Edge) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson6601e8cdDecodeGithubComSkydiveProjectSkydiveApiTypes8(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Edge) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson6601e8cdDecodeGithubComSkydiveProjectSkydiveApiTypes8(l, v)
}
func easyjson6601e8cdDecodeGithubComSkydiveProjectSkydiveApiTypes9(in *jlexer.Lexer, out *Capture) {
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
		case "GremlinQuery":
			out.GremlinQuery = string(in.String())
		case "BPFFilter":
			out.BPFFilter = string(in.String())
		case "Name":
			out.Name = string(in.String())
		case "Description":
			out.Description = string(in.String())
		case "Type":
			out.Type = string(in.String())
		case "Count":
			out.Count = int(in.Int())
		case "Port":
			out.Port = int(in.Int())
		case "SamplingRate":
			out.SamplingRate = uint32(in.Uint32())
		case "PollingInterval":
			out.PollingInterval = uint32(in.Uint32())
		case "RawPacketLimit":
			out.RawPacketLimit = int(in.Int())
		case "HeaderSize":
			out.HeaderSize = int(in.Int())
		case "ExtraTCPMetric":
			out.ExtraTCPMetric = bool(in.Bool())
		case "IPDefrag":
			out.IPDefrag = bool(in.Bool())
		case "ReassembleTCP":
			out.ReassembleTCP = bool(in.Bool())
		case "LayerKeyMode":
			out.LayerKeyMode = string(in.String())
		case "ExtraLayers":
			out.ExtraLayers = int(in.Int())
		case "Target":
			out.Target = string(in.String())
		case "TargetType":
			out.TargetType = string(in.String())
		case "UUID":
			out.UUID = string(in.String())
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
func easyjson6601e8cdEncodeGithubComSkydiveProjectSkydiveApiTypes9(out *jwriter.Writer, in Capture) {
	out.RawByte('{')
	first := true
	_ = first
	if in.GremlinQuery != "" {
		const prefix string = ",\"GremlinQuery\":"
		first = false
		out.RawString(prefix[1:])
		out.String(string(in.GremlinQuery))
	}
	if in.BPFFilter != "" {
		const prefix string = ",\"BPFFilter\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.BPFFilter))
	}
	if in.Name != "" {
		const prefix string = ",\"Name\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Name))
	}
	if in.Description != "" {
		const prefix string = ",\"Description\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Description))
	}
	if in.Type != "" {
		const prefix string = ",\"Type\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Type))
	}
	{
		const prefix string = ",\"Count\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int(int(in.Count))
	}
	if in.Port != 0 {
		const prefix string = ",\"Port\":"
		out.RawString(prefix)
		out.Int(int(in.Port))
	}
	{
		const prefix string = ",\"SamplingRate\":"
		out.RawString(prefix)
		out.Uint32(uint32(in.SamplingRate))
	}
	{
		const prefix string = ",\"PollingInterval\":"
		out.RawString(prefix)
		out.Uint32(uint32(in.PollingInterval))
	}
	if in.RawPacketLimit != 0 {
		const prefix string = ",\"RawPacketLimit\":"
		out.RawString(prefix)
		out.Int(int(in.RawPacketLimit))
	}
	if in.HeaderSize != 0 {
		const prefix string = ",\"HeaderSize\":"
		out.RawString(prefix)
		out.Int(int(in.HeaderSize))
	}
	{
		const prefix string = ",\"ExtraTCPMetric\":"
		out.RawString(prefix)
		out.Bool(bool(in.ExtraTCPMetric))
	}
	{
		const prefix string = ",\"IPDefrag\":"
		out.RawString(prefix)
		out.Bool(bool(in.IPDefrag))
	}
	{
		const prefix string = ",\"ReassembleTCP\":"
		out.RawString(prefix)
		out.Bool(bool(in.ReassembleTCP))
	}
	if in.LayerKeyMode != "" {
		const prefix string = ",\"LayerKeyMode\":"
		out.RawString(prefix)
		out.String(string(in.LayerKeyMode))
	}
	if in.ExtraLayers != 0 {
		const prefix string = ",\"ExtraLayers\":"
		out.RawString(prefix)
		out.Int(int(in.ExtraLayers))
	}
	if in.Target != "" {
		const prefix string = ",\"Target\":"
		out.RawString(prefix)
		out.String(string(in.Target))
	}
	if in.TargetType != "" {
		const prefix string = ",\"TargetType\":"
		out.RawString(prefix)
		out.String(string(in.TargetType))
	}
	{
		const prefix string = ",\"UUID\":"
		out.RawString(prefix)
		out.String(string(in.UUID))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Capture) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson6601e8cdEncodeGithubComSkydiveProjectSkydiveApiTypes9(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Capture) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson6601e8cdEncodeGithubComSkydiveProjectSkydiveApiTypes9(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Capture) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson6601e8cdDecodeGithubComSkydiveProjectSkydiveApiTypes9(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Capture) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson6601e8cdDecodeGithubComSkydiveProjectSkydiveApiTypes9(l, v)
}
func easyjson6601e8cdDecodeGithubComSkydiveProjectSkydiveApiTypes10(in *jlexer.Lexer, out *Alert) {
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
		case "Description":
			out.Description = string(in.String())
		case "Expression":
			out.Expression = string(in.String())
		case "Action":
			out.Action = string(in.String())
		case "Trigger":
			out.Trigger = string(in.String())
		case "CreateTime":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.CreateTime).UnmarshalJSON(data))
			}
		case "UUID":
			out.UUID = string(in.String())
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
func easyjson6601e8cdEncodeGithubComSkydiveProjectSkydiveApiTypes10(out *jwriter.Writer, in Alert) {
	out.RawByte('{')
	first := true
	_ = first
	if in.Name != "" {
		const prefix string = ",\"Name\":"
		first = false
		out.RawString(prefix[1:])
		out.String(string(in.Name))
	}
	if in.Description != "" {
		const prefix string = ",\"Description\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Description))
	}
	if in.Expression != "" {
		const prefix string = ",\"Expression\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Expression))
	}
	if in.Action != "" {
		const prefix string = ",\"Action\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Action))
	}
	if in.Trigger != "" {
		const prefix string = ",\"Trigger\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Trigger))
	}
	{
		const prefix string = ",\"CreateTime\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Raw((in.CreateTime).MarshalJSON())
	}
	{
		const prefix string = ",\"UUID\":"
		out.RawString(prefix)
		out.String(string(in.UUID))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Alert) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson6601e8cdEncodeGithubComSkydiveProjectSkydiveApiTypes10(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Alert) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson6601e8cdEncodeGithubComSkydiveProjectSkydiveApiTypes10(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Alert) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson6601e8cdDecodeGithubComSkydiveProjectSkydiveApiTypes10(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Alert) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson6601e8cdDecodeGithubComSkydiveProjectSkydiveApiTypes10(l, v)
}
