// Code generated by protoc-gen-go. DO NOT EDIT.
// source: mem_service.proto

package v1

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	_ "github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger/options"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

//entry
type Mem struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Title                string   `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Description          string   `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Mem) Reset()         { *m = Mem{} }
func (m *Mem) String() string { return proto.CompactTextString(m) }
func (*Mem) ProtoMessage()    {}
func (*Mem) Descriptor() ([]byte, []int) {
	return fileDescriptor_d90a088e1d8e1ded, []int{0}
}

func (m *Mem) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Mem.Unmarshal(m, b)
}
func (m *Mem) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Mem.Marshal(b, m, deterministic)
}
func (m *Mem) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Mem.Merge(m, src)
}
func (m *Mem) XXX_Size() int {
	return xxx_messageInfo_Mem.Size(m)
}
func (m *Mem) XXX_DiscardUnknown() {
	xxx_messageInfo_Mem.DiscardUnknown(m)
}

var xxx_messageInfo_Mem proto.InternalMessageInfo

func (m *Mem) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Mem) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *Mem) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

type MemInfo struct {
	Mem                  *Mem                 `protobuf:"bytes,1,opt,name=mem,proto3" json:"mem,omitempty"`
	CreatedAt            *timestamp.Timestamp `protobuf:"bytes,2,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt            *timestamp.Timestamp `protobuf:"bytes,3,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *MemInfo) Reset()         { *m = MemInfo{} }
func (m *MemInfo) String() string { return proto.CompactTextString(m) }
func (*MemInfo) ProtoMessage()    {}
func (*MemInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_d90a088e1d8e1ded, []int{1}
}

func (m *MemInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MemInfo.Unmarshal(m, b)
}
func (m *MemInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MemInfo.Marshal(b, m, deterministic)
}
func (m *MemInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MemInfo.Merge(m, src)
}
func (m *MemInfo) XXX_Size() int {
	return xxx_messageInfo_MemInfo.Size(m)
}
func (m *MemInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_MemInfo.DiscardUnknown(m)
}

var xxx_messageInfo_MemInfo proto.InternalMessageInfo

func (m *MemInfo) GetMem() *Mem {
	if m != nil {
		return m.Mem
	}
	return nil
}

func (m *MemInfo) GetCreatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

func (m *MemInfo) GetUpdatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.UpdatedAt
	}
	return nil
}

//new entry
type NewMemRequest struct {
	Title                string   `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Description          string   `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NewMemRequest) Reset()         { *m = NewMemRequest{} }
func (m *NewMemRequest) String() string { return proto.CompactTextString(m) }
func (*NewMemRequest) ProtoMessage()    {}
func (*NewMemRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_d90a088e1d8e1ded, []int{2}
}

func (m *NewMemRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NewMemRequest.Unmarshal(m, b)
}
func (m *NewMemRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NewMemRequest.Marshal(b, m, deterministic)
}
func (m *NewMemRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NewMemRequest.Merge(m, src)
}
func (m *NewMemRequest) XXX_Size() int {
	return xxx_messageInfo_NewMemRequest.Size(m)
}
func (m *NewMemRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_NewMemRequest.DiscardUnknown(m)
}

var xxx_messageInfo_NewMemRequest proto.InternalMessageInfo

func (m *NewMemRequest) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *NewMemRequest) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

type NewMemResponse struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NewMemResponse) Reset()         { *m = NewMemResponse{} }
func (m *NewMemResponse) String() string { return proto.CompactTextString(m) }
func (*NewMemResponse) ProtoMessage()    {}
func (*NewMemResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_d90a088e1d8e1ded, []int{3}
}

func (m *NewMemResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NewMemResponse.Unmarshal(m, b)
}
func (m *NewMemResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NewMemResponse.Marshal(b, m, deterministic)
}
func (m *NewMemResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NewMemResponse.Merge(m, src)
}
func (m *NewMemResponse) XXX_Size() int {
	return xxx_messageInfo_NewMemResponse.Size(m)
}
func (m *NewMemResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_NewMemResponse.DiscardUnknown(m)
}

var xxx_messageInfo_NewMemResponse proto.InternalMessageInfo

func (m *NewMemResponse) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

//get entries
type GetMemsRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetMemsRequest) Reset()         { *m = GetMemsRequest{} }
func (m *GetMemsRequest) String() string { return proto.CompactTextString(m) }
func (*GetMemsRequest) ProtoMessage()    {}
func (*GetMemsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_d90a088e1d8e1ded, []int{4}
}

func (m *GetMemsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetMemsRequest.Unmarshal(m, b)
}
func (m *GetMemsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetMemsRequest.Marshal(b, m, deterministic)
}
func (m *GetMemsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetMemsRequest.Merge(m, src)
}
func (m *GetMemsRequest) XXX_Size() int {
	return xxx_messageInfo_GetMemsRequest.Size(m)
}
func (m *GetMemsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetMemsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetMemsRequest proto.InternalMessageInfo

type GetMemsResponse struct {
	Mems                 []*MemInfo `protobuf:"bytes,1,rep,name=mems,proto3" json:"mems,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *GetMemsResponse) Reset()         { *m = GetMemsResponse{} }
func (m *GetMemsResponse) String() string { return proto.CompactTextString(m) }
func (*GetMemsResponse) ProtoMessage()    {}
func (*GetMemsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_d90a088e1d8e1ded, []int{5}
}

func (m *GetMemsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetMemsResponse.Unmarshal(m, b)
}
func (m *GetMemsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetMemsResponse.Marshal(b, m, deterministic)
}
func (m *GetMemsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetMemsResponse.Merge(m, src)
}
func (m *GetMemsResponse) XXX_Size() int {
	return xxx_messageInfo_GetMemsResponse.Size(m)
}
func (m *GetMemsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetMemsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetMemsResponse proto.InternalMessageInfo

func (m *GetMemsResponse) GetMems() []*MemInfo {
	if m != nil {
		return m.Mems
	}
	return nil
}

//update entry
type UpdateMemRequest struct {
	Mem                  *Mem     `protobuf:"bytes,1,opt,name=mem,proto3" json:"mem,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateMemRequest) Reset()         { *m = UpdateMemRequest{} }
func (m *UpdateMemRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateMemRequest) ProtoMessage()    {}
func (*UpdateMemRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_d90a088e1d8e1ded, []int{6}
}

func (m *UpdateMemRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateMemRequest.Unmarshal(m, b)
}
func (m *UpdateMemRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateMemRequest.Marshal(b, m, deterministic)
}
func (m *UpdateMemRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateMemRequest.Merge(m, src)
}
func (m *UpdateMemRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateMemRequest.Size(m)
}
func (m *UpdateMemRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateMemRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateMemRequest proto.InternalMessageInfo

func (m *UpdateMemRequest) GetMem() *Mem {
	if m != nil {
		return m.Mem
	}
	return nil
}

type UpdateMemResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateMemResponse) Reset()         { *m = UpdateMemResponse{} }
func (m *UpdateMemResponse) String() string { return proto.CompactTextString(m) }
func (*UpdateMemResponse) ProtoMessage()    {}
func (*UpdateMemResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_d90a088e1d8e1ded, []int{7}
}

func (m *UpdateMemResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateMemResponse.Unmarshal(m, b)
}
func (m *UpdateMemResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateMemResponse.Marshal(b, m, deterministic)
}
func (m *UpdateMemResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateMemResponse.Merge(m, src)
}
func (m *UpdateMemResponse) XXX_Size() int {
	return xxx_messageInfo_UpdateMemResponse.Size(m)
}
func (m *UpdateMemResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateMemResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateMemResponse proto.InternalMessageInfo

//get entry
type GetMemRequest struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetMemRequest) Reset()         { *m = GetMemRequest{} }
func (m *GetMemRequest) String() string { return proto.CompactTextString(m) }
func (*GetMemRequest) ProtoMessage()    {}
func (*GetMemRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_d90a088e1d8e1ded, []int{8}
}

func (m *GetMemRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetMemRequest.Unmarshal(m, b)
}
func (m *GetMemRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetMemRequest.Marshal(b, m, deterministic)
}
func (m *GetMemRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetMemRequest.Merge(m, src)
}
func (m *GetMemRequest) XXX_Size() int {
	return xxx_messageInfo_GetMemRequest.Size(m)
}
func (m *GetMemRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetMemRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetMemRequest proto.InternalMessageInfo

func (m *GetMemRequest) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type GetMemResponse struct {
	Mem                  *MemInfo `protobuf:"bytes,1,opt,name=mem,proto3" json:"mem,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetMemResponse) Reset()         { *m = GetMemResponse{} }
func (m *GetMemResponse) String() string { return proto.CompactTextString(m) }
func (*GetMemResponse) ProtoMessage()    {}
func (*GetMemResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_d90a088e1d8e1ded, []int{9}
}

func (m *GetMemResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetMemResponse.Unmarshal(m, b)
}
func (m *GetMemResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetMemResponse.Marshal(b, m, deterministic)
}
func (m *GetMemResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetMemResponse.Merge(m, src)
}
func (m *GetMemResponse) XXX_Size() int {
	return xxx_messageInfo_GetMemResponse.Size(m)
}
func (m *GetMemResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetMemResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetMemResponse proto.InternalMessageInfo

func (m *GetMemResponse) GetMem() *MemInfo {
	if m != nil {
		return m.Mem
	}
	return nil
}

func init() {
	proto.RegisterType((*Mem)(nil), "v1.Mem")
	proto.RegisterType((*MemInfo)(nil), "v1.MemInfo")
	proto.RegisterType((*NewMemRequest)(nil), "v1.NewMemRequest")
	proto.RegisterType((*NewMemResponse)(nil), "v1.NewMemResponse")
	proto.RegisterType((*GetMemsRequest)(nil), "v1.GetMemsRequest")
	proto.RegisterType((*GetMemsResponse)(nil), "v1.GetMemsResponse")
	proto.RegisterType((*UpdateMemRequest)(nil), "v1.UpdateMemRequest")
	proto.RegisterType((*UpdateMemResponse)(nil), "v1.UpdateMemResponse")
	proto.RegisterType((*GetMemRequest)(nil), "v1.GetMemRequest")
	proto.RegisterType((*GetMemResponse)(nil), "v1.GetMemResponse")
}

func init() { proto.RegisterFile("mem_service.proto", fileDescriptor_d90a088e1d8e1ded) }

var fileDescriptor_d90a088e1d8e1ded = []byte{
	// 493 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x52, 0xcb, 0x6e, 0xd3, 0x40,
	0x14, 0x95, 0x6d, 0x48, 0xc8, 0x8d, 0x92, 0x26, 0x93, 0x16, 0x05, 0xab, 0x28, 0x96, 0x57, 0x11,
	0x22, 0xb6, 0x62, 0x56, 0xb0, 0xab, 0x54, 0x51, 0xb1, 0x30, 0x0b, 0xf3, 0xd8, 0x56, 0x6e, 0x7c,
	0x6b, 0x8d, 0xd4, 0xb1, 0x8d, 0x67, 0x92, 0x2e, 0x50, 0x37, 0xfc, 0x02, 0x3b, 0xbe, 0x84, 0xff,
	0xe0, 0x17, 0xf8, 0x10, 0x34, 0x0f, 0xa7, 0x76, 0x2a, 0x60, 0x65, 0xf9, 0xdc, 0x73, 0xcf, 0x3d,
	0xe7, 0x68, 0x60, 0xca, 0x90, 0x5d, 0x72, 0xac, 0x77, 0x74, 0x83, 0x41, 0x55, 0x97, 0xa2, 0x24,
	0xf6, 0x6e, 0xed, 0x2e, 0xf2, 0xb2, 0xcc, 0x6f, 0x30, 0x54, 0xc8, 0xd5, 0xf6, 0x3a, 0x14, 0x94,
	0x21, 0x17, 0x29, 0xab, 0x34, 0xc9, 0x3d, 0x35, 0x84, 0xb4, 0xa2, 0x61, 0x5a, 0x14, 0xa5, 0x48,
	0x05, 0x2d, 0x0b, 0x6e, 0xa6, 0x2f, 0xd5, 0x67, 0xb3, 0xca, 0xb1, 0x58, 0xf1, 0xdb, 0x34, 0xcf,
	0xb1, 0x0e, 0xcb, 0x4a, 0x31, 0x1e, 0xb2, 0xfd, 0x18, 0x9c, 0x18, 0x19, 0x19, 0x83, 0x4d, 0xb3,
	0xb9, 0xe5, 0x59, 0x4b, 0x27, 0xb1, 0x69, 0x46, 0x8e, 0xe1, 0xb1, 0xa0, 0xe2, 0x06, 0xe7, 0xb6,
	0x67, 0x2d, 0x07, 0x89, 0xfe, 0x21, 0x1e, 0x0c, 0x33, 0xe4, 0x9b, 0x9a, 0x2a, 0xb9, 0xb9, 0xa3,
	0x66, 0x6d, 0xc8, 0xff, 0x61, 0x41, 0x3f, 0x46, 0xf6, 0xae, 0xb8, 0x2e, 0xc9, 0x33, 0x70, 0x18,
	0x32, 0x25, 0x3a, 0x8c, 0xfa, 0xc1, 0x6e, 0x1d, 0xc4, 0xc8, 0x12, 0x89, 0x91, 0xd7, 0x00, 0x9b,
	0x1a, 0x53, 0x81, 0xd9, 0x65, 0x2a, 0xd4, 0x8d, 0x61, 0xe4, 0x06, 0x3a, 0x56, 0xd0, 0xe4, 0x0e,
	0x3e, 0x36, 0xb9, 0x93, 0x81, 0x61, 0x9f, 0x09, 0xb9, 0xba, 0xad, 0xb2, 0x66, 0xd5, 0xf9, 0xff,
	0xaa, 0x61, 0x9f, 0x09, 0xff, 0x02, 0x46, 0xef, 0xf1, 0x56, 0x9a, 0xc0, 0x2f, 0x5b, 0xe4, 0xe2,
	0x3e, 0xa5, 0xf5, 0x8f, 0x94, 0xf6, 0xc3, 0x94, 0x1e, 0x8c, 0x1b, 0x21, 0x5e, 0x95, 0x05, 0xc7,
	0xc3, 0xfe, 0xfc, 0x09, 0x8c, 0x2f, 0x50, 0xc4, 0xc8, 0xb8, 0xb9, 0xe5, 0x47, 0x70, 0xb4, 0x47,
	0xcc, 0xd2, 0x02, 0x1e, 0x31, 0x64, 0x7c, 0x6e, 0x79, 0xce, 0x72, 0x18, 0x0d, 0x4d, 0x43, 0xb2,
	0xbb, 0x44, 0x0d, 0xfc, 0x15, 0x4c, 0x3e, 0x29, 0xf7, 0x2d, 0xcf, 0x7f, 0x6f, 0xd5, 0x9f, 0xc1,
	0xb4, 0x45, 0xd7, 0x47, 0xfc, 0x05, 0x8c, 0xf4, 0xdd, 0x46, 0xe0, 0xd0, 0x6a, 0xd8, 0x58, 0xdd,
	0xfb, 0x7a, 0xde, 0x3e, 0xd1, 0xb1, 0x25, 0xf1, 0xe8, 0xa7, 0x0d, 0x10, 0x23, 0xfb, 0xa0, 0x1f,
	0x2e, 0x39, 0x87, 0x9e, 0x2e, 0x83, 0x4c, 0x25, 0xb5, 0xd3, 0xb0, 0x4b, 0xda, 0x90, 0x71, 0x34,
	0xfb, 0xf6, 0xeb, 0xf7, 0x77, 0x7b, 0xe4, 0x3f, 0x09, 0x77, 0xeb, 0x50, 0xe6, 0x7c, 0x63, 0xbd,
	0x20, 0xe7, 0xd0, 0x37, 0xf5, 0x10, 0xb5, 0xd3, 0x6d, 0xcf, 0x9d, 0x75, 0x30, 0x23, 0x34, 0x51,
	0x42, 0x40, 0xf6, 0x42, 0xe4, 0x33, 0x0c, 0xf6, 0x0d, 0x90, 0x63, 0xb9, 0x73, 0xd8, 0x9f, 0x7b,
	0x72, 0x80, 0x1a, 0xad, 0x53, 0xa5, 0xf5, 0xd4, 0x9d, 0x36, 0x5a, 0xe1, 0x57, 0x86, 0x2c, 0xa0,
	0xd9, 0x9d, 0x74, 0xf7, 0x16, 0x7a, 0xfa, 0xb8, 0xce, 0xd8, 0x29, 0xd4, 0x25, 0x6d, 0xc8, 0xc8,
	0x9d, 0x28, 0xb9, 0x23, 0x32, 0xba, 0x97, 0xa3, 0xd9, 0xdd, 0x55, 0x4f, 0x3d, 0xd0, 0x57, 0x7f,
	0x02, 0x00, 0x00, 0xff, 0xff, 0xdd, 0x06, 0x83, 0xd5, 0xfa, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// MemServiceClient is the client API for MemService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MemServiceClient interface {
	//create a new mem
	NewMem(ctx context.Context, in *NewMemRequest, opts ...grpc.CallOption) (*NewMemResponse, error)
	//Get all mems
	GetMems(ctx context.Context, in *GetMemsRequest, opts ...grpc.CallOption) (*GetMemsResponse, error)
	//update mem
	UpdateMem(ctx context.Context, in *UpdateMemRequest, opts ...grpc.CallOption) (*UpdateMemResponse, error)
	GetMem(ctx context.Context, in *GetMemRequest, opts ...grpc.CallOption) (*GetMemResponse, error)
}

type memServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewMemServiceClient(cc grpc.ClientConnInterface) MemServiceClient {
	return &memServiceClient{cc}
}

func (c *memServiceClient) NewMem(ctx context.Context, in *NewMemRequest, opts ...grpc.CallOption) (*NewMemResponse, error) {
	out := new(NewMemResponse)
	err := c.cc.Invoke(ctx, "/v1.MemService/NewMem", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *memServiceClient) GetMems(ctx context.Context, in *GetMemsRequest, opts ...grpc.CallOption) (*GetMemsResponse, error) {
	out := new(GetMemsResponse)
	err := c.cc.Invoke(ctx, "/v1.MemService/GetMems", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *memServiceClient) UpdateMem(ctx context.Context, in *UpdateMemRequest, opts ...grpc.CallOption) (*UpdateMemResponse, error) {
	out := new(UpdateMemResponse)
	err := c.cc.Invoke(ctx, "/v1.MemService/UpdateMem", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *memServiceClient) GetMem(ctx context.Context, in *GetMemRequest, opts ...grpc.CallOption) (*GetMemResponse, error) {
	out := new(GetMemResponse)
	err := c.cc.Invoke(ctx, "/v1.MemService/GetMem", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MemServiceServer is the server API for MemService service.
type MemServiceServer interface {
	//create a new mem
	NewMem(context.Context, *NewMemRequest) (*NewMemResponse, error)
	//Get all mems
	GetMems(context.Context, *GetMemsRequest) (*GetMemsResponse, error)
	//update mem
	UpdateMem(context.Context, *UpdateMemRequest) (*UpdateMemResponse, error)
	GetMem(context.Context, *GetMemRequest) (*GetMemResponse, error)
}

// UnimplementedMemServiceServer can be embedded to have forward compatible implementations.
type UnimplementedMemServiceServer struct {
}

func (*UnimplementedMemServiceServer) NewMem(ctx context.Context, req *NewMemRequest) (*NewMemResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NewMem not implemented")
}
func (*UnimplementedMemServiceServer) GetMems(ctx context.Context, req *GetMemsRequest) (*GetMemsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMems not implemented")
}
func (*UnimplementedMemServiceServer) UpdateMem(ctx context.Context, req *UpdateMemRequest) (*UpdateMemResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateMem not implemented")
}
func (*UnimplementedMemServiceServer) GetMem(ctx context.Context, req *GetMemRequest) (*GetMemResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMem not implemented")
}

func RegisterMemServiceServer(s *grpc.Server, srv MemServiceServer) {
	s.RegisterService(&_MemService_serviceDesc, srv)
}

func _MemService_NewMem_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NewMemRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MemServiceServer).NewMem(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.MemService/NewMem",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MemServiceServer).NewMem(ctx, req.(*NewMemRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MemService_GetMems_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMemsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MemServiceServer).GetMems(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.MemService/GetMems",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MemServiceServer).GetMems(ctx, req.(*GetMemsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MemService_UpdateMem_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateMemRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MemServiceServer).UpdateMem(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.MemService/UpdateMem",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MemServiceServer).UpdateMem(ctx, req.(*UpdateMemRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MemService_GetMem_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMemRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MemServiceServer).GetMem(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.MemService/GetMem",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MemServiceServer).GetMem(ctx, req.(*GetMemRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _MemService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "v1.MemService",
	HandlerType: (*MemServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "NewMem",
			Handler:    _MemService_NewMem_Handler,
		},
		{
			MethodName: "GetMems",
			Handler:    _MemService_GetMems_Handler,
		},
		{
			MethodName: "UpdateMem",
			Handler:    _MemService_UpdateMem_Handler,
		},
		{
			MethodName: "GetMem",
			Handler:    _MemService_GetMem_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "mem_service.proto",
}