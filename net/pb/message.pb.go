// Code generated by protoc-gen-go. DO NOT EDIT.
// source: message.proto

package pb

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	any "github.com/golang/protobuf/ptypes/any"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

type Message_MessageType int32

const (
	Message_ACK  Message_MessageType = 0
	Message_PING Message_MessageType = 1
	Message_CHAT Message_MessageType = 2
)

var Message_MessageType_name = map[int32]string{
	0: "ACK",
	1: "PING",
	2: "CHAT",
}

var Message_MessageType_value = map[string]int32{
	"ACK":  0,
	"PING": 1,
	"CHAT": 2,
}

func (x Message_MessageType) String() string {
	return proto.EnumName(Message_MessageType_name, int32(x))
}

func (Message_MessageType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_33c57e4bae7b9afd, []int{0, 0}
}

type ChatMessage_Flag int32

const (
	ChatMessage_MESSAGE ChatMessage_Flag = 0
	ChatMessage_TYPING  ChatMessage_Flag = 1
	ChatMessage_READ    ChatMessage_Flag = 2
)

var ChatMessage_Flag_name = map[int32]string{
	0: "MESSAGE",
	1: "TYPING",
	2: "READ",
}

var ChatMessage_Flag_value = map[string]int32{
	"MESSAGE": 0,
	"TYPING":  1,
	"READ":    2,
}

func (x ChatMessage_Flag) String() string {
	return proto.EnumName(ChatMessage_Flag_name, int32(x))
}

func (ChatMessage_Flag) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_33c57e4bae7b9afd, []int{1, 0}
}

type Message struct {
	MessageType          Message_MessageType `protobuf:"varint,1,opt,name=messageType,proto3,enum=Message_MessageType" json:"messageType,omitempty"`
	MessageID            string              `protobuf:"bytes,2,opt,name=messageID,proto3" json:"messageID,omitempty"`
	Payload              *any.Any            `protobuf:"bytes,3,opt,name=payload,proto3" json:"payload,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *Message) Reset()         { *m = Message{} }
func (m *Message) String() string { return proto.CompactTextString(m) }
func (*Message) ProtoMessage()    {}
func (*Message) Descriptor() ([]byte, []int) {
	return fileDescriptor_33c57e4bae7b9afd, []int{0}
}

func (m *Message) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Message.Unmarshal(m, b)
}
func (m *Message) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Message.Marshal(b, m, deterministic)
}
func (m *Message) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Message.Merge(m, src)
}
func (m *Message) XXX_Size() int {
	return xxx_messageInfo_Message.Size(m)
}
func (m *Message) XXX_DiscardUnknown() {
	xxx_messageInfo_Message.DiscardUnknown(m)
}

var xxx_messageInfo_Message proto.InternalMessageInfo

func (m *Message) GetMessageType() Message_MessageType {
	if m != nil {
		return m.MessageType
	}
	return Message_ACK
}

func (m *Message) GetMessageID() string {
	if m != nil {
		return m.MessageID
	}
	return ""
}

func (m *Message) GetPayload() *any.Any {
	if m != nil {
		return m.Payload
	}
	return nil
}

type ChatMessage struct {
	Subject   string               `protobuf:"bytes,1,opt,name=subject,proto3" json:"subject,omitempty"`
	Message   string               `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	Timestamp *timestamp.Timestamp `protobuf:"bytes,3,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Flag      ChatMessage_Flag     `protobuf:"varint,4,opt,name=flag,proto3,enum=ChatMessage_Flag" json:"flag,omitempty"`
	// Only used when Flag is READ.
	ReadID               string   `protobuf:"bytes,5,opt,name=readID,proto3" json:"readID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ChatMessage) Reset()         { *m = ChatMessage{} }
func (m *ChatMessage) String() string { return proto.CompactTextString(m) }
func (*ChatMessage) ProtoMessage()    {}
func (*ChatMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_33c57e4bae7b9afd, []int{1}
}

func (m *ChatMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ChatMessage.Unmarshal(m, b)
}
func (m *ChatMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ChatMessage.Marshal(b, m, deterministic)
}
func (m *ChatMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ChatMessage.Merge(m, src)
}
func (m *ChatMessage) XXX_Size() int {
	return xxx_messageInfo_ChatMessage.Size(m)
}
func (m *ChatMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_ChatMessage.DiscardUnknown(m)
}

var xxx_messageInfo_ChatMessage proto.InternalMessageInfo

func (m *ChatMessage) GetSubject() string {
	if m != nil {
		return m.Subject
	}
	return ""
}

func (m *ChatMessage) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *ChatMessage) GetTimestamp() *timestamp.Timestamp {
	if m != nil {
		return m.Timestamp
	}
	return nil
}

func (m *ChatMessage) GetFlag() ChatMessage_Flag {
	if m != nil {
		return m.Flag
	}
	return ChatMessage_MESSAGE
}

func (m *ChatMessage) GetReadID() string {
	if m != nil {
		return m.ReadID
	}
	return ""
}

type AckMessage struct {
	AckedMessageID       string   `protobuf:"bytes,1,opt,name=ackedMessageID,proto3" json:"ackedMessageID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AckMessage) Reset()         { *m = AckMessage{} }
func (m *AckMessage) String() string { return proto.CompactTextString(m) }
func (*AckMessage) ProtoMessage()    {}
func (*AckMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_33c57e4bae7b9afd, []int{2}
}

func (m *AckMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AckMessage.Unmarshal(m, b)
}
func (m *AckMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AckMessage.Marshal(b, m, deterministic)
}
func (m *AckMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AckMessage.Merge(m, src)
}
func (m *AckMessage) XXX_Size() int {
	return xxx_messageInfo_AckMessage.Size(m)
}
func (m *AckMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_AckMessage.DiscardUnknown(m)
}

var xxx_messageInfo_AckMessage proto.InternalMessageInfo

func (m *AckMessage) GetAckedMessageID() string {
	if m != nil {
		return m.AckedMessageID
	}
	return ""
}

func init() {
	proto.RegisterEnum("Message_MessageType", Message_MessageType_name, Message_MessageType_value)
	proto.RegisterEnum("ChatMessage_Flag", ChatMessage_Flag_name, ChatMessage_Flag_value)
	proto.RegisterType((*Message)(nil), "Message")
	proto.RegisterType((*ChatMessage)(nil), "ChatMessage")
	proto.RegisterType((*AckMessage)(nil), "AckMessage")
}

func init() { proto.RegisterFile("message.proto", fileDescriptor_33c57e4bae7b9afd) }

var fileDescriptor_33c57e4bae7b9afd = []byte{
	// 343 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x90, 0xc1, 0x4e, 0xf2, 0x40,
	0x14, 0x85, 0x99, 0xd2, 0x9f, 0xfe, 0xbd, 0x8d, 0xa4, 0x4e, 0x88, 0xa9, 0xc4, 0x44, 0xd2, 0x44,
	0x83, 0x2e, 0x86, 0x04, 0x8d, 0x71, 0x5b, 0x01, 0x91, 0x18, 0x8c, 0x19, 0xba, 0xd1, 0xdd, 0x14,
	0x86, 0xaa, 0x14, 0xda, 0xd0, 0xb2, 0xe8, 0xb3, 0xf9, 0x3c, 0xbe, 0x87, 0xe9, 0x74, 0x06, 0x08,
	0xae, 0xda, 0x33, 0xe7, 0xe4, 0xde, 0xef, 0x5c, 0x38, 0x5a, 0xf2, 0x34, 0x65, 0x21, 0x27, 0xc9,
	0x3a, 0xce, 0xe2, 0xe6, 0x69, 0x18, 0xc7, 0x61, 0xc4, 0x3b, 0x42, 0x05, 0x9b, 0x79, 0x87, 0xad,
	0x72, 0x69, 0x9d, 0x1f, 0x5a, 0xd9, 0xe7, 0x92, 0xa7, 0x19, 0x5b, 0x26, 0x65, 0xc0, 0xfd, 0x46,
	0x60, 0x8c, 0xcb, 0x69, 0xf8, 0x0e, 0x2c, 0x39, 0xd8, 0xcf, 0x13, 0xee, 0xa0, 0x16, 0x6a, 0xd7,
	0xbb, 0x0d, 0x22, 0x6d, 0xf5, 0x2d, 0x3c, 0xba, 0x1f, 0xc4, 0x67, 0x60, 0x4a, 0x39, 0xea, 0x3b,
	0x5a, 0x0b, 0xb5, 0x4d, 0xba, 0x7b, 0xc0, 0x04, 0x8c, 0x84, 0xe5, 0x51, 0xcc, 0x66, 0x4e, 0xb5,
	0x85, 0xda, 0x56, 0xb7, 0x41, 0x4a, 0x28, 0xa2, 0xa0, 0x88, 0xb7, 0xca, 0xa9, 0x0a, 0xb9, 0xd7,
	0x60, 0xed, 0x6d, 0xc2, 0x06, 0x54, 0xbd, 0xde, 0xb3, 0x5d, 0xc1, 0xff, 0x41, 0x7f, 0x1d, 0xbd,
	0x0c, 0x6d, 0x54, 0xfc, 0xf5, 0x9e, 0x3c, 0xdf, 0xd6, 0xdc, 0x1f, 0x04, 0x56, 0xef, 0x83, 0x65,
	0xaa, 0x81, 0x03, 0x46, 0xba, 0x09, 0xbe, 0xf8, 0x34, 0x13, 0xf4, 0x26, 0x55, 0xb2, 0x70, 0x24,
	0x92, 0x24, 0x54, 0x12, 0xdf, 0x83, 0xb9, 0x3d, 0x8a, 0x24, 0x6c, 0xfe, 0x21, 0xf4, 0x55, 0x82,
	0xee, 0xc2, 0xf8, 0x02, 0xf4, 0x79, 0xc4, 0x42, 0x47, 0x17, 0x87, 0x3a, 0x26, 0x7b, 0x24, 0xe4,
	0x31, 0x62, 0x21, 0x15, 0x36, 0x3e, 0x81, 0xda, 0x9a, 0xb3, 0xd9, 0xa8, 0xef, 0xfc, 0x13, 0x9b,
	0xa5, 0x72, 0xaf, 0x40, 0x2f, 0x52, 0xd8, 0x02, 0x63, 0x3c, 0x98, 0x4c, 0xbc, 0xe1, 0xc0, 0xae,
	0x60, 0x80, 0x9a, 0xff, 0xb6, 0xeb, 0x49, 0x07, 0x5e, 0xdf, 0xd6, 0xdc, 0x5b, 0x00, 0x6f, 0xba,
	0x50, 0x2d, 0x2f, 0xa1, 0xce, 0xa6, 0x0b, 0x3e, 0x1b, 0x6f, 0x8f, 0x5e, 0x96, 0x3d, 0x78, 0x7d,
	0xd0, 0xdf, 0xb5, 0x24, 0x08, 0x6a, 0xa2, 0xc4, 0xcd, 0x6f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xe2,
	0x42, 0x4f, 0xc6, 0x35, 0x02, 0x00, 0x00,
}
