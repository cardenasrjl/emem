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
	Id                   int64                `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Title                string               `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Description          string               `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	CreatedAt            *timestamp.Timestamp `protobuf:"bytes,4,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt            *timestamp.Timestamp `protobuf:"bytes,5,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
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

func (m *Mem) GetCreatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

func (m *Mem) GetUpdatedAt() *timestamp.Timestamp {
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
	return fileDescriptor_d90a088e1d8e1ded, []int{1}
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
	return fileDescriptor_d90a088e1d8e1ded, []int{2}
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
	return fileDescriptor_d90a088e1d8e1ded, []int{3}
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
	Mems                 []*Mem   `protobuf:"bytes,1,rep,name=mems,proto3" json:"mems,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetMemsResponse) Reset()         { *m = GetMemsResponse{} }
func (m *GetMemsResponse) String() string { return proto.CompactTextString(m) }
func (*GetMemsResponse) ProtoMessage()    {}
func (*GetMemsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_d90a088e1d8e1ded, []int{4}
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

func (m *GetMemsResponse) GetMems() []*Mem {
	if m != nil {
		return m.Mems
	}
	return nil
}

//update entry
type UpdateMemRequest struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Title                string   `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Description          string   `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateMemRequest) Reset()         { *m = UpdateMemRequest{} }
func (m *UpdateMemRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateMemRequest) ProtoMessage()    {}
func (*UpdateMemRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_d90a088e1d8e1ded, []int{5}
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

func (m *UpdateMemRequest) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *UpdateMemRequest) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *UpdateMemRequest) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
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
	return fileDescriptor_d90a088e1d8e1ded, []int{6}
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
	return fileDescriptor_d90a088e1d8e1ded, []int{7}
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
	Mem                  *Mem     `protobuf:"bytes,1,opt,name=mem,proto3" json:"mem,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetMemResponse) Reset()         { *m = GetMemResponse{} }
func (m *GetMemResponse) String() string { return proto.CompactTextString(m) }
func (*GetMemResponse) ProtoMessage()    {}
func (*GetMemResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_d90a088e1d8e1ded, []int{8}
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

func (m *GetMemResponse) GetMem() *Mem {
	if m != nil {
		return m.Mem
	}
	return nil
}

//delete entry
type DeleteMemRequest struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteMemRequest) Reset()         { *m = DeleteMemRequest{} }
func (m *DeleteMemRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteMemRequest) ProtoMessage()    {}
func (*DeleteMemRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_d90a088e1d8e1ded, []int{9}
}

func (m *DeleteMemRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteMemRequest.Unmarshal(m, b)
}
func (m *DeleteMemRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteMemRequest.Marshal(b, m, deterministic)
}
func (m *DeleteMemRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteMemRequest.Merge(m, src)
}
func (m *DeleteMemRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteMemRequest.Size(m)
}
func (m *DeleteMemRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteMemRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteMemRequest proto.InternalMessageInfo

func (m *DeleteMemRequest) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type DeleteMemResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteMemResponse) Reset()         { *m = DeleteMemResponse{} }
func (m *DeleteMemResponse) String() string { return proto.CompactTextString(m) }
func (*DeleteMemResponse) ProtoMessage()    {}
func (*DeleteMemResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_d90a088e1d8e1ded, []int{10}
}

func (m *DeleteMemResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteMemResponse.Unmarshal(m, b)
}
func (m *DeleteMemResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteMemResponse.Marshal(b, m, deterministic)
}
func (m *DeleteMemResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteMemResponse.Merge(m, src)
}
func (m *DeleteMemResponse) XXX_Size() int {
	return xxx_messageInfo_DeleteMemResponse.Size(m)
}
func (m *DeleteMemResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteMemResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteMemResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*Mem)(nil), "v1.Mem")
	proto.RegisterType((*NewMemRequest)(nil), "v1.NewMemRequest")
	proto.RegisterType((*NewMemResponse)(nil), "v1.NewMemResponse")
	proto.RegisterType((*GetMemsRequest)(nil), "v1.GetMemsRequest")
	proto.RegisterType((*GetMemsResponse)(nil), "v1.GetMemsResponse")
	proto.RegisterType((*UpdateMemRequest)(nil), "v1.UpdateMemRequest")
	proto.RegisterType((*UpdateMemResponse)(nil), "v1.UpdateMemResponse")
	proto.RegisterType((*GetMemRequest)(nil), "v1.GetMemRequest")
	proto.RegisterType((*GetMemResponse)(nil), "v1.GetMemResponse")
	proto.RegisterType((*DeleteMemRequest)(nil), "v1.DeleteMemRequest")
	proto.RegisterType((*DeleteMemResponse)(nil), "v1.DeleteMemResponse")
}

func init() { proto.RegisterFile("mem_service.proto", fileDescriptor_d90a088e1d8e1ded) }

var fileDescriptor_d90a088e1d8e1ded = []byte{
	// 505 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x93, 0xc1, 0x6e, 0xd3, 0x4c,
	0x14, 0x85, 0x65, 0xbb, 0x4d, 0xff, 0xdc, 0x28, 0x69, 0x32, 0x69, 0x25, 0xff, 0x06, 0xa9, 0xd6,
	0xac, 0xa2, 0x42, 0x6d, 0x25, 0xac, 0x60, 0x57, 0x29, 0xa2, 0xab, 0x80, 0x64, 0x60, 0xc3, 0xa6,
	0x72, 0xe3, 0x8b, 0x35, 0x52, 0xc6, 0x36, 0x9e, 0x49, 0xba, 0x40, 0x6c, 0x90, 0x78, 0x02, 0x1e,
	0x89, 0x47, 0xe0, 0x15, 0x78, 0x10, 0xe4, 0x99, 0xb1, 0xb1, 0x13, 0x51, 0x36, 0xac, 0xa2, 0x9c,
	0x39, 0xf7, 0x9b, 0x73, 0xcf, 0x24, 0x30, 0xe1, 0xc8, 0x6f, 0x05, 0x96, 0x3b, 0xb6, 0xc6, 0xa0,
	0x28, 0x73, 0x99, 0x13, 0x7b, 0x37, 0xf7, 0x2e, 0xd2, 0x3c, 0x4f, 0x37, 0x18, 0x2a, 0xe5, 0x6e,
	0xfb, 0x21, 0x94, 0x8c, 0xa3, 0x90, 0x31, 0x2f, 0xb4, 0xc9, 0x7b, 0x6c, 0x0c, 0x71, 0xc1, 0xc2,
	0x38, 0xcb, 0x72, 0x19, 0x4b, 0x96, 0x67, 0xc2, 0x9c, 0x3e, 0x55, 0x1f, 0xeb, 0xab, 0x14, 0xb3,
	0x2b, 0x71, 0x1f, 0xa7, 0x29, 0x96, 0x61, 0x5e, 0x28, 0xc7, 0xa1, 0x9b, 0x7e, 0xb7, 0xc0, 0x59,
	0x21, 0x27, 0x23, 0xb0, 0x59, 0xe2, 0x5a, 0xbe, 0x35, 0x73, 0x22, 0x9b, 0x25, 0xe4, 0x0c, 0x8e,
	0x25, 0x93, 0x1b, 0x74, 0x6d, 0xdf, 0x9a, 0xf5, 0x23, 0xfd, 0x85, 0xf8, 0x30, 0x48, 0x50, 0xac,
	0x4b, 0xa6, 0x78, 0xae, 0xa3, 0xce, 0xda, 0x12, 0x79, 0x0e, 0xb0, 0x2e, 0x31, 0x96, 0x98, 0xdc,
	0xc6, 0xd2, 0x3d, 0xf2, 0xad, 0xd9, 0x60, 0xe1, 0x05, 0x3a, 0x70, 0x50, 0x6f, 0x14, 0xbc, 0xad,
	0x37, 0x8a, 0xfa, 0xc6, 0x7d, 0x2d, 0xab, 0xd1, 0x6d, 0x91, 0xd4, 0xa3, 0xc7, 0x7f, 0x1f, 0x35,
	0xee, 0x6b, 0x49, 0x6f, 0x60, 0xf8, 0x0a, 0xef, 0x57, 0xc8, 0x23, 0xfc, 0xb8, 0x45, 0x21, 0x7f,
	0xc7, 0xb7, 0x1e, 0x88, 0x6f, 0x1f, 0xc4, 0xa7, 0x3e, 0x8c, 0x6a, 0x90, 0x28, 0xf2, 0x4c, 0xe0,
	0x7e, 0x31, 0x74, 0x0c, 0xa3, 0x1b, 0x94, 0x2b, 0xe4, 0xc2, 0xdc, 0x45, 0x03, 0x38, 0x6d, 0x14,
	0x33, 0xf4, 0x08, 0x8e, 0x38, 0x72, 0xe1, 0x5a, 0xbe, 0x33, 0x1b, 0x2c, 0x4e, 0x82, 0xdd, 0x3c,
	0xa8, 0x98, 0x4a, 0xa4, 0xef, 0x61, 0xfc, 0x4e, 0x25, 0x6f, 0xe5, 0xfd, 0x47, 0xf5, 0xd3, 0x29,
	0x4c, 0x5a, 0x6c, 0x9d, 0x86, 0x5e, 0xc0, 0x50, 0x07, 0xfc, 0xc3, 0x6d, 0xf4, 0x49, 0xbd, 0x53,
	0xb3, 0xc0, 0xff, 0xe0, 0x70, 0xe4, 0xca, 0xd2, 0xca, 0x5f, 0x69, 0x94, 0xc2, 0x78, 0x89, 0x1b,
	0x7c, 0x28, 0x7e, 0x15, 0xa3, 0xe5, 0xd1, 0xcc, 0xc5, 0x57, 0x07, 0x60, 0x85, 0xfc, 0x8d, 0xfe,
	0xc1, 0x93, 0x25, 0xf4, 0x74, 0xd5, 0x64, 0x52, 0xf1, 0x3b, 0xef, 0xe7, 0x91, 0xb6, 0x64, 0xd6,
	0x98, 0x7e, 0xf9, 0xf1, 0xf3, 0x9b, 0x3d, 0xa4, 0xff, 0x85, 0xbb, 0x79, 0x58, 0x35, 0xf9, 0xc2,
	0xba, 0x24, 0x4b, 0x38, 0x31, 0xe5, 0x13, 0x35, 0xd3, 0x7d, 0x1b, 0x6f, 0xda, 0xd1, 0x0c, 0x68,
	0xac, 0x40, 0x40, 0x1a, 0x10, 0x89, 0xa0, 0xdf, 0xd4, 0x46, 0xce, 0xaa, 0x99, 0xfd, 0x17, 0xf2,
	0xce, 0xf7, 0x54, 0xc3, 0x72, 0x15, 0x8b, 0x78, 0xc3, 0x9a, 0x15, 0x7e, 0x62, 0xc9, 0xe7, 0x2a,
	0xd9, 0x4b, 0xe8, 0xe9, 0x8b, 0xf5, 0x7e, 0x9d, 0x17, 0xf0, 0x48, 0x5b, 0x32, 0xa8, 0x73, 0x85,
	0x3a, 0x25, 0x5d, 0x14, 0x79, 0x0d, 0xfd, 0xa6, 0x4b, 0x9d, 0x6d, 0xbf, 0x7e, 0x9d, 0xed, 0xa0,
	0xf0, 0x1a, 0x78, 0xd9, 0x05, 0xde, 0xf5, 0xd4, 0x7f, 0xe9, 0xd9, 0xaf, 0x00, 0x00, 0x00, 0xff,
	0xff, 0xc5, 0x17, 0xb4, 0x79, 0x7f, 0x04, 0x00, 0x00,
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
	//creates a new mem
	NewMem(ctx context.Context, in *NewMemRequest, opts ...grpc.CallOption) (*NewMemResponse, error)
	//Get all mems
	GetMems(ctx context.Context, in *GetMemsRequest, opts ...grpc.CallOption) (*GetMemsResponse, error)
	//updates a particular mem
	UpdateMem(ctx context.Context, in *UpdateMemRequest, opts ...grpc.CallOption) (*UpdateMemResponse, error)
	//gets a particular mem
	GetMem(ctx context.Context, in *GetMemRequest, opts ...grpc.CallOption) (*GetMemResponse, error)
	//deletes a particular mem
	DeleteMem(ctx context.Context, in *DeleteMemRequest, opts ...grpc.CallOption) (*DeleteMemResponse, error)
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

func (c *memServiceClient) DeleteMem(ctx context.Context, in *DeleteMemRequest, opts ...grpc.CallOption) (*DeleteMemResponse, error) {
	out := new(DeleteMemResponse)
	err := c.cc.Invoke(ctx, "/v1.MemService/DeleteMem", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MemServiceServer is the server API for MemService service.
type MemServiceServer interface {
	//creates a new mem
	NewMem(context.Context, *NewMemRequest) (*NewMemResponse, error)
	//Get all mems
	GetMems(context.Context, *GetMemsRequest) (*GetMemsResponse, error)
	//updates a particular mem
	UpdateMem(context.Context, *UpdateMemRequest) (*UpdateMemResponse, error)
	//gets a particular mem
	GetMem(context.Context, *GetMemRequest) (*GetMemResponse, error)
	//deletes a particular mem
	DeleteMem(context.Context, *DeleteMemRequest) (*DeleteMemResponse, error)
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
func (*UnimplementedMemServiceServer) DeleteMem(ctx context.Context, req *DeleteMemRequest) (*DeleteMemResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteMem not implemented")
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

func _MemService_DeleteMem_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteMemRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MemServiceServer).DeleteMem(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.MemService/DeleteMem",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MemServiceServer).DeleteMem(ctx, req.(*DeleteMemRequest))
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
		{
			MethodName: "DeleteMem",
			Handler:    _MemService_DeleteMem_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "mem_service.proto",
}
