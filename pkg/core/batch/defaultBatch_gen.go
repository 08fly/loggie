package batch

// Code generated by github.com/tinylib/msgp DO NOT EDIT.

import (
	"github.com/loggie-io/loggie/pkg/core/event"
	"github.com/tinylib/msgp/msgp"
)

// DecodeMsg implements msgp.Decodable
func (z *DefaultBatchS) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, err = dc.ReadMapHeader()
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			err = msgp.WrapError(err)
			return
		}
		switch msgp.UnsafeString(field) {
		case "content":
			var zb0002 uint32
			zb0002, err = dc.ReadArrayHeader()
			if err != nil {
				err = msgp.WrapError(err, "Content")
				return
			}
			if cap(z.Content) >= int(zb0002) {
				z.Content = (z.Content)[:zb0002]
			} else {
				z.Content = make([]*event.DefaultEventS, zb0002)
			}
			for za0001 := range z.Content {
				if dc.IsNil() {
					err = dc.ReadNil()
					if err != nil {
						err = msgp.WrapError(err, "Content", za0001)
						return
					}
					z.Content[za0001] = nil
				} else {
					if z.Content[za0001] == nil {
						z.Content[za0001] = new(event.DefaultEventS)
					}
					err = z.Content[za0001].DecodeMsg(dc)
					if err != nil {
						err = msgp.WrapError(err, "Content", za0001)
						return
					}
				}
			}
		case "metadata":
			var zb0003 uint32
			zb0003, err = dc.ReadMapHeader()
			if err != nil {
				err = msgp.WrapError(err, "Metadata")
				return
			}
			if z.Metadata == nil {
				z.Metadata = make(map[string]interface{}, zb0003)
			} else if len(z.Metadata) > 0 {
				for key := range z.Metadata {
					delete(z.Metadata, key)
				}
			}
			for zb0003 > 0 {
				zb0003--
				var za0002 string
				var za0003 interface{}
				za0002, err = dc.ReadString()
				if err != nil {
					err = msgp.WrapError(err, "Metadata")
					return
				}
				za0003, err = dc.ReadIntf()
				if err != nil {
					err = msgp.WrapError(err, "Metadata", za0002)
					return
				}
				z.Metadata[za0002] = za0003
			}
		default:
			err = dc.Skip()
			if err != nil {
				err = msgp.WrapError(err)
				return
			}
		}
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z *DefaultBatchS) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 2
	// write "content"
	err = en.Append(0x82, 0xa7, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74)
	if err != nil {
		return
	}
	err = en.WriteArrayHeader(uint32(len(z.Content)))
	if err != nil {
		err = msgp.WrapError(err, "Content")
		return
	}
	for za0001 := range z.Content {
		if z.Content[za0001] == nil {
			err = en.WriteNil()
			if err != nil {
				return
			}
		} else {
			err = z.Content[za0001].EncodeMsg(en)
			if err != nil {
				err = msgp.WrapError(err, "Content", za0001)
				return
			}
		}
	}
	// write "metadata"
	err = en.Append(0xa8, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61)
	if err != nil {
		return
	}
	err = en.WriteMapHeader(uint32(len(z.Metadata)))
	if err != nil {
		err = msgp.WrapError(err, "Metadata")
		return
	}
	for za0002, za0003 := range z.Metadata {
		err = en.WriteString(za0002)
		if err != nil {
			err = msgp.WrapError(err, "Metadata")
			return
		}
		err = en.WriteIntf(za0003)
		if err != nil {
			err = msgp.WrapError(err, "Metadata", za0002)
			return
		}
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *DefaultBatchS) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 2
	// string "content"
	o = append(o, 0x82, 0xa7, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74)
	o = msgp.AppendArrayHeader(o, uint32(len(z.Content)))
	for za0001 := range z.Content {
		if z.Content[za0001] == nil {
			o = msgp.AppendNil(o)
		} else {
			o, err = z.Content[za0001].MarshalMsg(o)
			if err != nil {
				err = msgp.WrapError(err, "Content", za0001)
				return
			}
		}
	}
	// string "metadata"
	o = append(o, 0xa8, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61)
	o = msgp.AppendMapHeader(o, uint32(len(z.Metadata)))
	for za0002, za0003 := range z.Metadata {
		o = msgp.AppendString(o, za0002)
		o, err = msgp.AppendIntf(o, za0003)
		if err != nil {
			err = msgp.WrapError(err, "Metadata", za0002)
			return
		}
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *DefaultBatchS) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			err = msgp.WrapError(err)
			return
		}
		switch msgp.UnsafeString(field) {
		case "content":
			var zb0002 uint32
			zb0002, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "Content")
				return
			}
			if cap(z.Content) >= int(zb0002) {
				z.Content = (z.Content)[:zb0002]
			} else {
				z.Content = make([]*event.DefaultEventS, zb0002)
			}
			for za0001 := range z.Content {
				if msgp.IsNil(bts) {
					bts, err = msgp.ReadNilBytes(bts)
					if err != nil {
						return
					}
					z.Content[za0001] = nil
				} else {
					if z.Content[za0001] == nil {
						z.Content[za0001] = new(event.DefaultEventS)
					}
					bts, err = z.Content[za0001].UnmarshalMsg(bts)
					if err != nil {
						err = msgp.WrapError(err, "Content", za0001)
						return
					}
				}
			}
		case "metadata":
			var zb0003 uint32
			zb0003, bts, err = msgp.ReadMapHeaderBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "Metadata")
				return
			}
			if z.Metadata == nil {
				z.Metadata = make(map[string]interface{}, zb0003)
			} else if len(z.Metadata) > 0 {
				for key := range z.Metadata {
					delete(z.Metadata, key)
				}
			}
			for zb0003 > 0 {
				var za0002 string
				var za0003 interface{}
				zb0003--
				za0002, bts, err = msgp.ReadStringBytes(bts)
				if err != nil {
					err = msgp.WrapError(err, "Metadata")
					return
				}
				za0003, bts, err = msgp.ReadIntfBytes(bts)
				if err != nil {
					err = msgp.WrapError(err, "Metadata", za0002)
					return
				}
				z.Metadata[za0002] = za0003
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				err = msgp.WrapError(err)
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *DefaultBatchS) Msgsize() (s int) {
	s = 1 + 8 + msgp.ArrayHeaderSize
	for za0001 := range z.Content {
		if z.Content[za0001] == nil {
			s += msgp.NilSize
		} else {
			s += z.Content[za0001].Msgsize()
		}
	}
	s += 9 + msgp.MapHeaderSize
	if z.Metadata != nil {
		for za0002, za0003 := range z.Metadata {
			_ = za0003
			s += msgp.StringPrefixSize + len(za0002) + msgp.GuessSize(za0003)
		}
	}
	return
}
