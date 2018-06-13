// Copyright (c) 2016-2018 iQIYI.com.  All rights reserved.
// 
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// 
//     http://www.apache.org/licenses/LICENSE-2.0
// 
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
// 

// Code generated by protoc-gen-go. DO NOT EDIT.
// source: object.proto

/*
Package pack is a generated protocol buffer package.

It is generated from these files:
	object.proto
	rpc.proto

It has these top-level messages:
	ObjectMeta
	NeedleIndex
	DBIndex
	ObjectTimestamps
	CheckedObjects
	WantedParts
	WantedObjects
	Partition
	PartitionSuffixesReply
	SuffixHashesMsg
	SuffixHashesReply
	SyncMsg
	SyncReply
	PartitionDeletionReply
	PartitionAuditionReply
*/
package pack

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type ObjectMeta struct {
	Name       string            `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Timestamp  string            `protobuf:"bytes,2,opt,name=timestamp" json:"timestamp,omitempty"`
	DataSize   int64             `protobuf:"varint,3,opt,name=dataSize" json:"dataSize,omitempty"`
	SystemMeta map[string]string `protobuf:"bytes,4,rep,name=systemMeta" json:"systemMeta,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	UserMeta   map[string]string `protobuf:"bytes,5,rep,name=userMeta" json:"userMeta,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *ObjectMeta) Reset()                    { *m = ObjectMeta{} }
func (m *ObjectMeta) String() string            { return proto.CompactTextString(m) }
func (*ObjectMeta) ProtoMessage()               {}
func (*ObjectMeta) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *ObjectMeta) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ObjectMeta) GetTimestamp() string {
	if m != nil {
		return m.Timestamp
	}
	return ""
}

func (m *ObjectMeta) GetDataSize() int64 {
	if m != nil {
		return m.DataSize
	}
	return 0
}

func (m *ObjectMeta) GetSystemMeta() map[string]string {
	if m != nil {
		return m.SystemMeta
	}
	return nil
}

func (m *ObjectMeta) GetUserMeta() map[string]string {
	if m != nil {
		return m.UserMeta
	}
	return nil
}

type NeedleIndex struct {
	Offset     int64 `protobuf:"varint,1,opt,name=offset" json:"offset,omitempty"`
	Size       int64 `protobuf:"varint,2,opt,name=size" json:"size,omitempty"`
	DataOffset int64 `protobuf:"varint,3,opt,name=dataOffset" json:"dataOffset,omitempty"`
	DataSize   int64 `protobuf:"varint,4,opt,name=dataSize" json:"dataSize,omitempty"`
	MetaOffset int64 `protobuf:"varint,5,opt,name=metaOffset" json:"metaOffset,omitempty"`
	MetaSize   int32 `protobuf:"varint,6,opt,name=metaSize" json:"metaSize,omitempty"`
}

func (m *NeedleIndex) Reset()                    { *m = NeedleIndex{} }
func (m *NeedleIndex) String() string            { return proto.CompactTextString(m) }
func (*NeedleIndex) ProtoMessage()               {}
func (*NeedleIndex) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *NeedleIndex) GetOffset() int64 {
	if m != nil {
		return m.Offset
	}
	return 0
}

func (m *NeedleIndex) GetSize() int64 {
	if m != nil {
		return m.Size
	}
	return 0
}

func (m *NeedleIndex) GetDataOffset() int64 {
	if m != nil {
		return m.DataOffset
	}
	return 0
}

func (m *NeedleIndex) GetDataSize() int64 {
	if m != nil {
		return m.DataSize
	}
	return 0
}

func (m *NeedleIndex) GetMetaOffset() int64 {
	if m != nil {
		return m.MetaOffset
	}
	return 0
}

func (m *NeedleIndex) GetMetaSize() int32 {
	if m != nil {
		return m.MetaSize
	}
	return 0
}

type DBIndex struct {
	Index *NeedleIndex `protobuf:"bytes,1,opt,name=index" json:"index,omitempty"`
	Meta  *ObjectMeta  `protobuf:"bytes,2,opt,name=meta" json:"meta,omitempty"`
}

func (m *DBIndex) Reset()                    { *m = DBIndex{} }
func (m *DBIndex) String() string            { return proto.CompactTextString(m) }
func (*DBIndex) ProtoMessage()               {}
func (*DBIndex) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *DBIndex) GetIndex() *NeedleIndex {
	if m != nil {
		return m.Index
	}
	return nil
}

func (m *DBIndex) GetMeta() *ObjectMeta {
	if m != nil {
		return m.Meta
	}
	return nil
}

type ObjectTimestamps struct {
	DataTimestamp string `protobuf:"bytes,1,opt,name=dataTimestamp" json:"dataTimestamp,omitempty"`
	MetaTimestamp string `protobuf:"bytes,2,opt,name=metaTimestamp" json:"metaTimestamp,omitempty"`
}

func (m *ObjectTimestamps) Reset()                    { *m = ObjectTimestamps{} }
func (m *ObjectTimestamps) String() string            { return proto.CompactTextString(m) }
func (*ObjectTimestamps) ProtoMessage()               {}
func (*ObjectTimestamps) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *ObjectTimestamps) GetDataTimestamp() string {
	if m != nil {
		return m.DataTimestamp
	}
	return ""
}

func (m *ObjectTimestamps) GetMetaTimestamp() string {
	if m != nil {
		return m.MetaTimestamp
	}
	return ""
}

type CheckedObjects struct {
	Objects map[string]*ObjectTimestamps `protobuf:"bytes,1,rep,name=objects" json:"objects,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *CheckedObjects) Reset()                    { *m = CheckedObjects{} }
func (m *CheckedObjects) String() string            { return proto.CompactTextString(m) }
func (*CheckedObjects) ProtoMessage()               {}
func (*CheckedObjects) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *CheckedObjects) GetObjects() map[string]*ObjectTimestamps {
	if m != nil {
		return m.Objects
	}
	return nil
}

type WantedParts struct {
	Data bool `protobuf:"varint,1,opt,name=data" json:"data,omitempty"`
	Meta bool `protobuf:"varint,2,opt,name=meta" json:"meta,omitempty"`
}

func (m *WantedParts) Reset()                    { *m = WantedParts{} }
func (m *WantedParts) String() string            { return proto.CompactTextString(m) }
func (*WantedParts) ProtoMessage()               {}
func (*WantedParts) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *WantedParts) GetData() bool {
	if m != nil {
		return m.Data
	}
	return false
}

func (m *WantedParts) GetMeta() bool {
	if m != nil {
		return m.Meta
	}
	return false
}

type WantedObjects struct {
	Objects map[string]*WantedParts `protobuf:"bytes,1,rep,name=objects" json:"objects,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *WantedObjects) Reset()                    { *m = WantedObjects{} }
func (m *WantedObjects) String() string            { return proto.CompactTextString(m) }
func (*WantedObjects) ProtoMessage()               {}
func (*WantedObjects) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *WantedObjects) GetObjects() map[string]*WantedParts {
	if m != nil {
		return m.Objects
	}
	return nil
}

func init() {
	proto.RegisterType((*ObjectMeta)(nil), "pack.ObjectMeta")
	proto.RegisterType((*NeedleIndex)(nil), "pack.NeedleIndex")
	proto.RegisterType((*DBIndex)(nil), "pack.DBIndex")
	proto.RegisterType((*ObjectTimestamps)(nil), "pack.ObjectTimestamps")
	proto.RegisterType((*CheckedObjects)(nil), "pack.CheckedObjects")
	proto.RegisterType((*WantedParts)(nil), "pack.WantedParts")
	proto.RegisterType((*WantedObjects)(nil), "pack.WantedObjects")
}

func init() { proto.RegisterFile("object.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 488 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x54, 0xcd, 0x6e, 0xd3, 0x40,
	0x10, 0xd6, 0xc6, 0x76, 0x9a, 0x8e, 0x1b, 0x48, 0x57, 0xa8, 0xb2, 0x22, 0x14, 0x19, 0xab, 0x52,
	0x73, 0x40, 0x39, 0x04, 0x21, 0xa1, 0x54, 0x48, 0x88, 0x9f, 0x03, 0x87, 0x52, 0xb4, 0x2d, 0x82,
	0x13, 0xd2, 0x36, 0x9e, 0x8a, 0x90, 0x3a, 0x89, 0xb2, 0x1b, 0x44, 0x78, 0x14, 0x6e, 0x3c, 0x02,
	0xef, 0xc2, 0x03, 0xa1, 0x9d, 0xdd, 0xd8, 0x6b, 0x03, 0x42, 0xbd, 0xcd, 0x7c, 0x3b, 0xdf, 0xec,
	0x7c, 0xdf, 0x8e, 0x0d, 0x07, 0xcb, 0xab, 0xcf, 0x38, 0xd5, 0xa3, 0xd5, 0x7a, 0xa9, 0x97, 0x3c,
	0x5c, 0xc9, 0xe9, 0x3c, 0xfb, 0xd5, 0x02, 0x38, 0x27, 0xf8, 0x0c, 0xb5, 0xe4, 0x1c, 0xc2, 0x85,
	0x2c, 0x30, 0x61, 0x29, 0x1b, 0xee, 0x0b, 0x8a, 0xf9, 0x7d, 0xd8, 0xd7, 0xb3, 0x02, 0x95, 0x96,
	0xc5, 0x2a, 0x69, 0xd1, 0x41, 0x05, 0xf0, 0x3e, 0x74, 0x72, 0xa9, 0xe5, 0xc5, 0xec, 0x1b, 0x26,
	0x41, 0xca, 0x86, 0x81, 0x28, 0x73, 0xfe, 0x0c, 0x40, 0x6d, 0x95, 0xc6, 0xc2, 0xf4, 0x4e, 0xc2,
	0x34, 0x18, 0xc6, 0xe3, 0x74, 0x64, 0xee, 0x1d, 0x55, 0x77, 0x8e, 0x2e, 0xca, 0x92, 0x57, 0x0b,
	0xbd, 0xde, 0x0a, 0x8f, 0xc3, 0x27, 0xd0, 0xd9, 0x28, 0x5c, 0x13, 0x3f, 0x22, 0xfe, 0xe0, 0x0f,
	0xfe, 0x3b, 0x57, 0x60, 0xd9, 0x65, 0x7d, 0xff, 0x29, 0xdc, 0x6d, 0xb4, 0xe6, 0x3d, 0x08, 0xe6,
	0xb8, 0x75, 0xea, 0x4c, 0xc8, 0xef, 0x41, 0xf4, 0x45, 0xde, 0x6c, 0xd0, 0x09, 0xb3, 0xc9, 0xa4,
	0xf5, 0x84, 0xf5, 0x4f, 0xa1, 0x5b, 0xeb, 0x7c, 0x1b, 0x72, 0xf6, 0x93, 0x41, 0xfc, 0x06, 0x31,
	0xbf, 0xc1, 0xd7, 0x8b, 0x1c, 0xbf, 0xf2, 0x23, 0x68, 0x2f, 0xaf, 0xaf, 0x15, 0x6a, 0xa2, 0x07,
	0xc2, 0x65, 0xc6, 0x6f, 0x65, 0x9c, 0x6b, 0x11, 0x4a, 0x31, 0x1f, 0x00, 0x18, 0x07, 0xcf, 0x6d,
	0xbd, 0xf5, 0xd4, 0x43, 0x6a, 0x8e, 0x87, 0x0d, 0xc7, 0x07, 0x00, 0x05, 0x96, 0xdc, 0xc8, 0x72,
	0x2b, 0xc4, 0x70, 0x4d, 0x46, 0xdc, 0x76, 0xca, 0x86, 0x91, 0x28, 0xf3, 0xec, 0x03, 0xec, 0xbd,
	0x7c, 0x6e, 0xc7, 0x3d, 0x81, 0x68, 0x66, 0x02, 0x9a, 0x36, 0x1e, 0x1f, 0x5a, 0xcf, 0x3d, 0x41,
	0xc2, 0x9e, 0xf3, 0x63, 0x08, 0x0d, 0x9f, 0xe6, 0x8f, 0xc7, 0xbd, 0xe6, 0xdb, 0x08, 0x3a, 0xcd,
	0x3e, 0x42, 0xcf, 0x62, 0x97, 0xbb, 0xb5, 0x51, 0xfc, 0x18, 0xba, 0x66, 0xea, 0x12, 0x71, 0xbe,
	0xd6, 0x41, 0x53, 0x65, 0x3a, 0x5c, 0x36, 0xf6, 0xaf, 0x0e, 0x66, 0x3f, 0x18, 0xdc, 0x79, 0xf1,
	0x09, 0xa7, 0x73, 0xcc, 0xed, 0x3d, 0x8a, 0x9f, 0xc2, 0x9e, 0xdd, 0x76, 0x95, 0x30, 0xda, 0x9b,
	0x07, 0x76, 0xb6, 0x7a, 0x99, 0x1b, 0x55, 0xd9, 0xd5, 0xd9, 0x31, 0xfa, 0x02, 0x0e, 0xfc, 0x83,
	0xbf, 0xbc, 0xfc, 0x43, 0xff, 0xe5, 0xe3, 0xf1, 0x91, 0x2f, 0xbc, 0x12, 0xe9, 0x6f, 0xc4, 0x63,
	0x88, 0xdf, 0xcb, 0x85, 0xc6, 0xfc, 0xad, 0x5c, 0x6b, 0x65, 0x1e, 0xde, 0x28, 0xa5, 0x9e, 0x1d,
	0x41, 0xb1, 0xc1, 0x4a, 0x33, 0x3b, 0xce, 0xba, 0xef, 0x0c, 0xba, 0x96, 0xb7, 0x53, 0x36, 0x69,
	0x2a, 0x73, 0x5f, 0x54, 0xad, 0xea, 0x1f, 0xc2, 0xce, 0xfe, 0x2b, 0xec, 0xa4, 0x2e, 0xec, 0xd0,
	0xef, 0x4d, 0x93, 0x7b, 0x9a, 0xae, 0xda, 0xf4, 0x27, 0x79, 0xf4, 0x3b, 0x00, 0x00, 0xff, 0xff,
	0x34, 0x01, 0xbd, 0x51, 0x59, 0x04, 0x00, 0x00,
}