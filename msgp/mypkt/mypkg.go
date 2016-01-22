//go:generate msgp
// Copyright (c) 2014 Luis Pabon
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package mypkt

import (
	"github.com/tinylib/msgp/msgp"
)

type MyMsg struct {
	Values []interface{}
}

type MyPkt struct {
	Id      uint32
	Block   uint64
	NBlocks uint64
	Path    string
}

func (z *MyPkt) ArrayMsgsize() (s int) {
	s = 1 + 1 + msgp.Uint32Size + 1 + msgp.Uint64Size + 1 + msgp.Uint64Size + 1 + msgp.StringPrefixSize + len(z.Path)
	return
}

func (z *MyPkt) MarshalArrayMsg() (o []byte, err error) {
	o = msgp.Require(nil, z.ArrayMsgsize())

	// array header, size 4
	o = msgp.AppendArrayHeader(o, 4)

	// "Id"
	o = msgp.AppendUint32(o, z.Id)
	// "Block"
	o = msgp.AppendUint64(o, z.Block)
	// "NBlocks"
	o = msgp.AppendUint64(o, z.NBlocks)
	// "Path"
	o = msgp.AppendString(o, z.Path)

	return
}

func (z *MyPkt) UnmarshalArrayMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var i, isz uint32
	isz, bts, err = msgp.ReadArrayHeaderBytes(bts)
	if err != nil {
		return
	}
	for i = 0; i < isz; i++ {
		switch i {
		case 0:
			z.Id, bts, err = msgp.ReadUint32Bytes(bts)
			if err != nil {
				return
			}
		case 1:
			z.Block, bts, err = msgp.ReadUint64Bytes(bts)
			if err != nil {
				return
			}
		case 2:
			z.NBlocks, bts, err = msgp.ReadUint64Bytes(bts)
			if err != nil {
				return
			}
		case 3:
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
