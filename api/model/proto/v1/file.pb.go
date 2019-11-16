// Code generated by protoc-gen-go. DO NOT EDIT.
// source: file.proto

package model

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type File_FileType int32

const (
	File_PDF  File_FileType = 0
	File_XLS  File_FileType = 1
	File_DOCX File_FileType = 2
	File_ZIP  File_FileType = 3
	File_JPEG File_FileType = 4
	File_PPT  File_FileType = 5
	File_MIME File_FileType = 6
)

var File_FileType_name = map[int32]string{
	0: "PDF",
	1: "XLS",
	2: "DOCX",
	3: "ZIP",
	4: "JPEG",
	5: "PPT",
	6: "MIME",
}

var File_FileType_value = map[string]int32{
	"PDF":  0,
	"XLS":  1,
	"DOCX": 2,
	"ZIP":  3,
	"JPEG": 4,
	"PPT":  5,
	"MIME": 6,
}

func (x File_FileType) String() string {
	return proto.EnumName(File_FileType_name, int32(x))
}

func (File_FileType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_9188e3b7e55e1162, []int{0, 0}
}

type File struct {
	File                 []byte   `protobuf:"bytes,1,opt,name=File,proto3" json:"File,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *File) Reset()         { *m = File{} }
func (m *File) String() string { return proto.CompactTextString(m) }
func (*File) ProtoMessage()    {}
func (*File) Descriptor() ([]byte, []int) {
	return fileDescriptor_9188e3b7e55e1162, []int{0}
}

func (m *File) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_File.Unmarshal(m, b)
}
func (m *File) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_File.Marshal(b, m, deterministic)
}
func (m *File) XXX_Merge(src proto.Message) {
	xxx_messageInfo_File.Merge(m, src)
}
func (m *File) XXX_Size() int {
	return xxx_messageInfo_File.Size(m)
}
func (m *File) XXX_DiscardUnknown() {
	xxx_messageInfo_File.DiscardUnknown(m)
}

var xxx_messageInfo_File proto.InternalMessageInfo

func (m *File) GetFile() []byte {
	if m != nil {
		return m.File
	}
	return nil
}

type CreateResponse struct {
	// API versioning: it is my best practice to specify version explicitly
	Api string `protobuf:"bytes,1,opt,name=api,proto3" json:"api,omitempty"`
	// ID of created file
	UUID                 int64    `protobuf:"varint,2,opt,name=UUID,proto3" json:"UUID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateResponse) Reset()         { *m = CreateResponse{} }
func (m *CreateResponse) String() string { return proto.CompactTextString(m) }
func (*CreateResponse) ProtoMessage()    {}
func (*CreateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_9188e3b7e55e1162, []int{1}
}

func (m *CreateResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateResponse.Unmarshal(m, b)
}
func (m *CreateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateResponse.Marshal(b, m, deterministic)
}
func (m *CreateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateResponse.Merge(m, src)
}
func (m *CreateResponse) XXX_Size() int {
	return xxx_messageInfo_CreateResponse.Size(m)
}
func (m *CreateResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateResponse proto.InternalMessageInfo

func (m *CreateResponse) GetApi() string {
	if m != nil {
		return m.Api
	}
	return ""
}

func (m *CreateResponse) GetUUID() int64 {
	if m != nil {
		return m.UUID
	}
	return 0
}

type DeleteResponse struct {
	// API versioning: it is my best practice to specify version explicitly
	Api string `protobuf:"bytes,1,opt,name=api,proto3" json:"api,omitempty"`
	// ID of created file
	UUID                 int64    `protobuf:"varint,2,opt,name=UUID,proto3" json:"UUID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteResponse) Reset()         { *m = DeleteResponse{} }
func (m *DeleteResponse) String() string { return proto.CompactTextString(m) }
func (*DeleteResponse) ProtoMessage()    {}
func (*DeleteResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_9188e3b7e55e1162, []int{2}
}

func (m *DeleteResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteResponse.Unmarshal(m, b)
}
func (m *DeleteResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteResponse.Marshal(b, m, deterministic)
}
func (m *DeleteResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteResponse.Merge(m, src)
}
func (m *DeleteResponse) XXX_Size() int {
	return xxx_messageInfo_DeleteResponse.Size(m)
}
func (m *DeleteResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteResponse proto.InternalMessageInfo

func (m *DeleteResponse) GetApi() string {
	if m != nil {
		return m.Api
	}
	return ""
}

func (m *DeleteResponse) GetUUID() int64 {
	if m != nil {
		return m.UUID
	}
	return 0
}

type ListResponse struct {
	// API versioning: it is my best practice to specify version explicitly
	Api string `protobuf:"bytes,1,opt,name=api,proto3" json:"api,omitempty"`
	// ID of created file
	UUID                 int64    `protobuf:"varint,2,opt,name=UUID,proto3" json:"UUID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListResponse) Reset()         { *m = ListResponse{} }
func (m *ListResponse) String() string { return proto.CompactTextString(m) }
func (*ListResponse) ProtoMessage()    {}
func (*ListResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_9188e3b7e55e1162, []int{3}
}

func (m *ListResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListResponse.Unmarshal(m, b)
}
func (m *ListResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListResponse.Marshal(b, m, deterministic)
}
func (m *ListResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListResponse.Merge(m, src)
}
func (m *ListResponse) XXX_Size() int {
	return xxx_messageInfo_ListResponse.Size(m)
}
func (m *ListResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListResponse proto.InternalMessageInfo

func (m *ListResponse) GetApi() string {
	if m != nil {
		return m.Api
	}
	return ""
}

func (m *ListResponse) GetUUID() int64 {
	if m != nil {
		return m.UUID
	}
	return 0
}

type UpdateResponse struct {
	// API versioning: it is my best practice to specify version explicitly
	Api string `protobuf:"bytes,1,opt,name=api,proto3" json:"api,omitempty"`
	// ID of created file
	UUID                 int64    `protobuf:"varint,2,opt,name=UUID,proto3" json:"UUID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateResponse) Reset()         { *m = UpdateResponse{} }
func (m *UpdateResponse) String() string { return proto.CompactTextString(m) }
func (*UpdateResponse) ProtoMessage()    {}
func (*UpdateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_9188e3b7e55e1162, []int{4}
}

func (m *UpdateResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateResponse.Unmarshal(m, b)
}
func (m *UpdateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateResponse.Marshal(b, m, deterministic)
}
func (m *UpdateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateResponse.Merge(m, src)
}
func (m *UpdateResponse) XXX_Size() int {
	return xxx_messageInfo_UpdateResponse.Size(m)
}
func (m *UpdateResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateResponse proto.InternalMessageInfo

func (m *UpdateResponse) GetApi() string {
	if m != nil {
		return m.Api
	}
	return ""
}

func (m *UpdateResponse) GetUUID() int64 {
	if m != nil {
		return m.UUID
	}
	return 0
}

func init() {
	proto.RegisterEnum("model.File_FileType", File_FileType_name, File_FileType_value)
	proto.RegisterType((*File)(nil), "model.File")
	proto.RegisterType((*CreateResponse)(nil), "model.CreateResponse")
	proto.RegisterType((*DeleteResponse)(nil), "model.DeleteResponse")
	proto.RegisterType((*ListResponse)(nil), "model.ListResponse")
	proto.RegisterType((*UpdateResponse)(nil), "model.UpdateResponse")
}

func init() { proto.RegisterFile("file.proto", fileDescriptor_9188e3b7e55e1162) }

var fileDescriptor_9188e3b7e55e1162 = []byte{
	// 267 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x92, 0xc1, 0x4a, 0xc3, 0x40,
	0x10, 0x86, 0xdd, 0x26, 0x5d, 0xdb, 0xb1, 0xc8, 0xb0, 0x22, 0x14, 0x4f, 0x25, 0xa7, 0x1c, 0x42,
	0x0e, 0x2a, 0xbe, 0x40, 0xd3, 0x4a, 0x24, 0xc5, 0x90, 0x36, 0x50, 0xbc, 0x55, 0x3a, 0x62, 0x20,
	0xba, 0x4b, 0x92, 0x8b, 0x4f, 0xe7, 0xab, 0x95, 0xd9, 0xf4, 0x92, 0x92, 0x4b, 0x4e, 0xf3, 0xcf,
	0xcf, 0xcc, 0xce, 0x7c, 0xc3, 0x02, 0x7c, 0x15, 0x25, 0x85, 0xa6, 0xd2, 0x8d, 0x56, 0xe3, 0x1f,
	0x7d, 0xa4, 0xd2, 0xfb, 0x06, 0x77, 0x5d, 0x94, 0xa4, 0x54, 0x1b, 0xe7, 0x62, 0x21, 0xfc, 0x59,
	0x66, 0xb5, 0x97, 0xc0, 0x84, 0xe3, 0xee, 0xcf, 0x90, 0xba, 0x06, 0x27, 0x8d, 0xd6, 0x78, 0xc5,
	0x62, 0x9f, 0x6c, 0x51, 0xa8, 0x09, 0xb8, 0xd1, 0xfb, 0x72, 0x8f, 0x23, 0xb6, 0x3e, 0xe2, 0x14,
	0x1d, 0xb6, 0xde, 0xd2, 0xd5, 0x2b, 0xba, 0xb6, 0x3c, 0xdd, 0xe1, 0x98, 0xad, 0x4d, 0xbc, 0x59,
	0xa1, 0xf4, 0x5e, 0xe0, 0x76, 0x59, 0xd1, 0xa1, 0xa1, 0x8c, 0x6a, 0xa3, 0x7f, 0x6b, 0x52, 0x08,
	0xce, 0xc1, 0x14, 0x76, 0xe4, 0x34, 0x63, 0xc9, 0x5b, 0xe4, 0x79, 0x1c, 0xcd, 0x47, 0x0b, 0xe1,
	0x3b, 0x99, 0xd5, 0xdc, 0x17, 0x51, 0x49, 0x83, 0xfb, 0x9e, 0x61, 0x96, 0x14, 0x75, 0x33, 0x7c,
	0x5a, 0x6e, 0x8e, 0x83, 0xb7, 0x7c, 0xfc, 0x17, 0x30, 0xdd, 0x36, 0xba, 0x22, 0x7b, 0xcd, 0x00,
	0x64, 0xcb, 0xaa, 0x6e, 0x42, 0x7b, 0xe7, 0x90, 0xed, 0x87, 0xfb, 0x73, 0x72, 0x71, 0x87, 0x00,
	0x64, 0x4b, 0xd8, 0x5f, 0x7d, 0x41, 0xef, 0x83, 0xcb, 0x5c, 0xdd, 0xda, 0xbb, 0x73, 0xd2, 0x21,
	0x0e, 0x40, 0xb6, 0x2c, 0xfd, 0xef, 0x76, 0x39, 0x3f, 0xa5, 0xfd, 0x17, 0x4f, 0xa7, 0x00, 0x00,
	0x00, 0xff, 0xff, 0x29, 0xfd, 0x60, 0x7b, 0x25, 0x02, 0x00, 0x00,
}