package mypkt

// NOTE: THIS FILE WAS PRODUCED BY THE
// MSGP CODE GENERATION TOOL (github.com/tinylib/msgp)
// DO NOT EDIT

import (
	"github.com/tinylib/msgp/msgp"
)

// DecodeMsg implements msgp.Decodable
func (z *MyMsg) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var isz uint32
	isz, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	for isz > 0 {
		isz--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "Values":
			var xsz uint32
			xsz, err = dc.ReadArrayHeader()
			if err != nil {
				return
			}
			if cap(z.Values) >= int(xsz) {
				z.Values = z.Values[:xsz]
			} else {
				z.Values = make([]interface{}, xsz)
			}
			for xvk := range z.Values {
				z.Values[xvk], err = dc.ReadIntf()
				if err != nil {
					return
				}
			}
		default:
			err = dc.Skip()
			if err != nil {
				return
			}
		}
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z *MyMsg) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 1
	// write "Values"
	err = en.Append(0x81, 0xa6, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x73)
	if err != nil {
		return err
	}
	err = en.WriteArrayHeader(uint32(len(z.Values)))
	if err != nil {
		return
	}
	for xvk := range z.Values {
		err = en.WriteIntf(z.Values[xvk])
		if err != nil {
			return
		}
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *MyMsg) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 1
	// string "Values"
	o = append(o, 0x81, 0xa6, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x73)
	o = msgp.AppendArrayHeader(o, uint32(len(z.Values)))
	for xvk := range z.Values {
		o, err = msgp.AppendIntf(o, z.Values[xvk])
		if err != nil {
			return
		}
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *MyMsg) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var isz uint32
	isz, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for isz > 0 {
		isz--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "Values":
			var xsz uint32
			xsz, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				return
			}
			if cap(z.Values) >= int(xsz) {
				z.Values = z.Values[:xsz]
			} else {
				z.Values = make([]interface{}, xsz)
			}
			for xvk := range z.Values {
				z.Values[xvk], bts, err = msgp.ReadIntfBytes(bts)
				if err != nil {
					return
				}
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				return
			}
		}
	}
	o = bts
	return
}

func (z *MyMsg) Msgsize() (s int) {
	s = 1 + 7 + msgp.ArrayHeaderSize
	for xvk := range z.Values {
		s += msgp.GuessSize(z.Values[xvk])
	}
	return
}

// DecodeMsg implements msgp.Decodable
func (z *MyPkt) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var isz uint32
	isz, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	for isz > 0 {
		isz--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "Id":
			z.Id, err = dc.ReadUint32()
			if err != nil {
				return
			}
		case "Block":
			z.Block, err = dc.ReadUint64()
			if err != nil {
				return
			}
		case "NBlocks":
			z.NBlocks, err = dc.ReadUint64()
			if err != nil {
				return
			}
		case "Path":
			z.Path, err = dc.ReadString()
			if err != nil {
				return
			}
		default:
			err = dc.Skip()
			if err != nil {
				return
			}
		}
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z *MyPkt) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 4
	// write "Id"
	err = en.Append(0x84, 0xa2, 0x49, 0x64)
	if err != nil {
		return err
	}
	err = en.WriteUint32(z.Id)
	if err != nil {
		return
	}
	// write "Block"
	err = en.Append(0xa5, 0x42, 0x6c, 0x6f, 0x63, 0x6b)
	if err != nil {
		return err
	}
	err = en.WriteUint64(z.Block)
	if err != nil {
		return
	}
	// write "NBlocks"
	err = en.Append(0xa7, 0x4e, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x73)
	if err != nil {
		return err
	}
	err = en.WriteUint64(z.NBlocks)
	if err != nil {
		return
	}
	// write "Path"
	err = en.Append(0xa4, 0x50, 0x61, 0x74, 0x68)
	if err != nil {
		return err
	}
	err = en.WriteString(z.Path)
	if err != nil {
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *MyPkt) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 4
	// string "Id"
	o = append(o, 0x84, 0xa2, 0x49, 0x64)
	o = msgp.AppendUint32(o, z.Id)
	// string "Block"
	o = append(o, 0xa5, 0x42, 0x6c, 0x6f, 0x63, 0x6b)
	o = msgp.AppendUint64(o, z.Block)
	// string "NBlocks"
	o = append(o, 0xa7, 0x4e, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x73)
	o = msgp.AppendUint64(o, z.NBlocks)
	// string "Path"
	o = append(o, 0xa4, 0x50, 0x61, 0x74, 0x68)
	o = msgp.AppendString(o, z.Path)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *MyPkt) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var isz uint32
	isz, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for isz > 0 {
		isz--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "Id":
			z.Id, bts, err = msgp.ReadUint32Bytes(bts)
			if err != nil {
				return
			}
		case "Block":
			z.Block, bts, err = msgp.ReadUint64Bytes(bts)
			if err != nil {
				return
			}
		case "NBlocks":
			z.NBlocks, bts, err = msgp.ReadUint64Bytes(bts)
			if err != nil {
				return
			}
		case "Path":
			z.Path, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				return
			}
		}
	}
	o = bts
	return
}

func (z *MyPkt) Msgsize() (s int) {
	s = 1 + 3 + msgp.Uint32Size + 6 + msgp.Uint64Size + 8 + msgp.Uint64Size + 5 + msgp.StringPrefixSize + len(z.Path)
	return
}
