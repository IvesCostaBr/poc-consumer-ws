// Code generated by protoc-gen-go. DO NOT EDIT.
// source: protos/exchange.proto

package pb

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

type SubscribeEventRequest struct {
	Channel              string   `protobuf:"bytes,1,opt,name=channel,proto3" json:"channel,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SubscribeEventRequest) Reset()         { *m = SubscribeEventRequest{} }
func (m *SubscribeEventRequest) String() string { return proto.CompactTextString(m) }
func (*SubscribeEventRequest) ProtoMessage()    {}
func (*SubscribeEventRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_8974edd908601500, []int{0}
}

func (m *SubscribeEventRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SubscribeEventRequest.Unmarshal(m, b)
}
func (m *SubscribeEventRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SubscribeEventRequest.Marshal(b, m, deterministic)
}
func (m *SubscribeEventRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SubscribeEventRequest.Merge(m, src)
}
func (m *SubscribeEventRequest) XXX_Size() int {
	return xxx_messageInfo_SubscribeEventRequest.Size(m)
}
func (m *SubscribeEventRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SubscribeEventRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SubscribeEventRequest proto.InternalMessageInfo

func (m *SubscribeEventRequest) GetChannel() string {
	if m != nil {
		return m.Channel
	}
	return ""
}

type DataReceiver struct {
	Data                 string   `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	Channel              string   `protobuf:"bytes,2,opt,name=channel,proto3" json:"channel,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DataReceiver) Reset()         { *m = DataReceiver{} }
func (m *DataReceiver) String() string { return proto.CompactTextString(m) }
func (*DataReceiver) ProtoMessage()    {}
func (*DataReceiver) Descriptor() ([]byte, []int) {
	return fileDescriptor_8974edd908601500, []int{1}
}

func (m *DataReceiver) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DataReceiver.Unmarshal(m, b)
}
func (m *DataReceiver) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DataReceiver.Marshal(b, m, deterministic)
}
func (m *DataReceiver) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DataReceiver.Merge(m, src)
}
func (m *DataReceiver) XXX_Size() int {
	return xxx_messageInfo_DataReceiver.Size(m)
}
func (m *DataReceiver) XXX_DiscardUnknown() {
	xxx_messageInfo_DataReceiver.DiscardUnknown(m)
}

var xxx_messageInfo_DataReceiver proto.InternalMessageInfo

func (m *DataReceiver) GetData() string {
	if m != nil {
		return m.Data
	}
	return ""
}

func (m *DataReceiver) GetChannel() string {
	if m != nil {
		return m.Channel
	}
	return ""
}

type ActionRequest struct {
	Action               string   `protobuf:"bytes,1,opt,name=action,proto3" json:"action,omitempty"`
	Market               string   `protobuf:"bytes,2,opt,name=market,proto3" json:"market,omitempty"`
	MessageId            int32    `protobuf:"varint,3,opt,name=messageId,proto3" json:"messageId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ActionRequest) Reset()         { *m = ActionRequest{} }
func (m *ActionRequest) String() string { return proto.CompactTextString(m) }
func (*ActionRequest) ProtoMessage()    {}
func (*ActionRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_8974edd908601500, []int{2}
}

func (m *ActionRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ActionRequest.Unmarshal(m, b)
}
func (m *ActionRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ActionRequest.Marshal(b, m, deterministic)
}
func (m *ActionRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ActionRequest.Merge(m, src)
}
func (m *ActionRequest) XXX_Size() int {
	return xxx_messageInfo_ActionRequest.Size(m)
}
func (m *ActionRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ActionRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ActionRequest proto.InternalMessageInfo

func (m *ActionRequest) GetAction() string {
	if m != nil {
		return m.Action
	}
	return ""
}

func (m *ActionRequest) GetMarket() string {
	if m != nil {
		return m.Market
	}
	return ""
}

func (m *ActionRequest) GetMessageId() int32 {
	if m != nil {
		return m.MessageId
	}
	return 0
}

func init() {
	proto.RegisterType((*SubscribeEventRequest)(nil), "exchange.SubscribeEventRequest")
	proto.RegisterType((*DataReceiver)(nil), "exchange.DataReceiver")
	proto.RegisterType((*ActionRequest)(nil), "exchange.ActionRequest")
}

func init() { proto.RegisterFile("protos/exchange.proto", fileDescriptor_8974edd908601500) }

var fileDescriptor_8974edd908601500 = []byte{
	// 248 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2d, 0x28, 0xca, 0x2f,
	0xc9, 0x2f, 0xd6, 0x4f, 0xad, 0x48, 0xce, 0x48, 0xcc, 0x4b, 0x4f, 0xd5, 0x03, 0xf3, 0x85, 0x38,
	0x60, 0x7c, 0x25, 0x43, 0x2e, 0xd1, 0xe0, 0xd2, 0xa4, 0xe2, 0xe4, 0xa2, 0xcc, 0xa4, 0x54, 0xd7,
	0xb2, 0xd4, 0xbc, 0x92, 0xa0, 0xd4, 0xc2, 0xd2, 0xd4, 0xe2, 0x12, 0x21, 0x09, 0x2e, 0x76, 0x90,
	0x92, 0xbc, 0xd4, 0x1c, 0x09, 0x46, 0x05, 0x46, 0x0d, 0xce, 0x20, 0x18, 0x57, 0xc9, 0x86, 0x8b,
	0xc7, 0x25, 0xb1, 0x24, 0x31, 0x28, 0x35, 0x39, 0x35, 0xb3, 0x2c, 0xb5, 0x48, 0x48, 0x88, 0x8b,
	0x25, 0x25, 0xb1, 0x24, 0x11, 0xaa, 0x0c, 0xcc, 0x46, 0xd6, 0xcd, 0x84, 0xaa, 0x3b, 0x96, 0x8b,
	0xd7, 0x31, 0xb9, 0x24, 0x33, 0x3f, 0x0f, 0x66, 0x91, 0x18, 0x17, 0x5b, 0x22, 0x58, 0x00, 0x6a,
	0x00, 0x94, 0x07, 0x12, 0xcf, 0x4d, 0x2c, 0xca, 0x4e, 0x2d, 0x81, 0x9a, 0x00, 0xe5, 0x09, 0xc9,
	0x70, 0x71, 0xe6, 0xa6, 0x16, 0x17, 0x27, 0xa6, 0xa7, 0x7a, 0xa6, 0x48, 0x30, 0x2b, 0x30, 0x6a,
	0xb0, 0x06, 0x21, 0x04, 0x8c, 0xe6, 0x31, 0x72, 0xf1, 0x06, 0x97, 0x14, 0xa5, 0x26, 0xe6, 0xba,
	0x56, 0x24, 0xe6, 0x16, 0xe4, 0xa4, 0x0a, 0x39, 0x70, 0xf1, 0x42, 0xad, 0x82, 0xd8, 0x2b, 0x24,
	0xae, 0x07, 0x0f, 0x0d, 0x14, 0x97, 0x48, 0x89, 0x21, 0x24, 0x50, 0x3c, 0xe8, 0xcd, 0xc5, 0x87,
	0x1a, 0x46, 0x42, 0xf2, 0x08, 0x95, 0x58, 0x43, 0x0f, 0x97, 0x51, 0x06, 0x8c, 0x4e, 0x6c, 0x51,
	0x2c, 0x7a, 0xfa, 0x05, 0x49, 0x49, 0x6c, 0xe0, 0x98, 0x30, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff,
	0xb0, 0x9d, 0x88, 0x54, 0xa2, 0x01, 0x00, 0x00,
}