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
	Message_ACK              Message_MessageType = 0
	Message_PING             Message_MessageType = 1
	Message_PONG             Message_MessageType = 2
	Message_CHAT             Message_MessageType = 3
	Message_FOLLOW           Message_MessageType = 4
	Message_UNFOLLOW         Message_MessageType = 5
	Message_STORE            Message_MessageType = 6
	Message_ORDER            Message_MessageType = 7
	Message_ADDRESS_REQUEST  Message_MessageType = 8
	Message_ADDRESS_RESPONSE Message_MessageType = 9
)

var Message_MessageType_name = map[int32]string{
	0: "ACK",
	1: "PING",
	2: "PONG",
	3: "CHAT",
	4: "FOLLOW",
	5: "UNFOLLOW",
	6: "STORE",
	7: "ORDER",
	8: "ADDRESS_REQUEST",
	9: "ADDRESS_RESPONSE",
}

var Message_MessageType_value = map[string]int32{
	"ACK":              0,
	"PING":             1,
	"PONG":             2,
	"CHAT":             3,
	"FOLLOW":           4,
	"UNFOLLOW":         5,
	"STORE":            6,
	"ORDER":            7,
	"ADDRESS_REQUEST":  8,
	"ADDRESS_RESPONSE": 9,
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

type OrderMessage_MessageType int32

const (
	OrderMessage_ORDER_OPEN         OrderMessage_MessageType = 0
	OrderMessage_ORDER_REJECT       OrderMessage_MessageType = 1
	OrderMessage_ORDER_CANCEL       OrderMessage_MessageType = 2
	OrderMessage_ORDER_CONFIRMATION OrderMessage_MessageType = 3
	OrderMessage_ORDER_FULFILLMENT  OrderMessage_MessageType = 4
	OrderMessage_ORDER_COMPLETE     OrderMessage_MessageType = 5
	OrderMessage_DISPUTE_OPEN       OrderMessage_MessageType = 6
	OrderMessage_DISPUTE_UPDATE     OrderMessage_MessageType = 7
	OrderMessage_DISPUTE_CLOSE      OrderMessage_MessageType = 8
	OrderMessage_REFUND             OrderMessage_MessageType = 9
	OrderMessage_PAYMENT_SENT       OrderMessage_MessageType = 10
	OrderMessage_PAYMENT_FINALIZED  OrderMessage_MessageType = 11
)

var OrderMessage_MessageType_name = map[int32]string{
	0:  "ORDER_OPEN",
	1:  "ORDER_REJECT",
	2:  "ORDER_CANCEL",
	3:  "ORDER_CONFIRMATION",
	4:  "ORDER_FULFILLMENT",
	5:  "ORDER_COMPLETE",
	6:  "DISPUTE_OPEN",
	7:  "DISPUTE_UPDATE",
	8:  "DISPUTE_CLOSE",
	9:  "REFUND",
	10: "PAYMENT_SENT",
	11: "PAYMENT_FINALIZED",
}

var OrderMessage_MessageType_value = map[string]int32{
	"ORDER_OPEN":         0,
	"ORDER_REJECT":       1,
	"ORDER_CANCEL":       2,
	"ORDER_CONFIRMATION": 3,
	"ORDER_FULFILLMENT":  4,
	"ORDER_COMPLETE":     5,
	"DISPUTE_OPEN":       6,
	"DISPUTE_UPDATE":     7,
	"DISPUTE_CLOSE":      8,
	"REFUND":             9,
	"PAYMENT_SENT":       10,
	"PAYMENT_FINALIZED":  11,
}

func (x OrderMessage_MessageType) String() string {
	return proto.EnumName(OrderMessage_MessageType_name, int32(x))
}

func (OrderMessage_MessageType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_33c57e4bae7b9afd, []int{4, 0}
}

type Message struct {
	MessageType          Message_MessageType `protobuf:"varint,1,opt,name=messageType,proto3,enum=Message_MessageType" json:"messageType,omitempty"`
	MessageID            string              `protobuf:"bytes,2,opt,name=messageID,proto3" json:"messageID,omitempty"`
	Sequence             uint32              `protobuf:"varint,3,opt,name=sequence,proto3" json:"sequence,omitempty"`
	Payload              *any.Any            `protobuf:"bytes,4,opt,name=payload,proto3" json:"payload,omitempty"`
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

func (m *Message) GetSequence() uint32 {
	if m != nil {
		return m.Sequence
	}
	return 0
}

func (m *Message) GetPayload() *any.Any {
	if m != nil {
		return m.Payload
	}
	return nil
}

type ChatMessage struct {
	OrderID   string               `protobuf:"bytes,1,opt,name=orderID,proto3" json:"orderID,omitempty"`
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

func (m *ChatMessage) GetOrderID() string {
	if m != nil {
		return m.OrderID
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

type StoreMessage struct {
	Cids                 [][]byte `protobuf:"bytes,1,rep,name=cids,proto3" json:"cids,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StoreMessage) Reset()         { *m = StoreMessage{} }
func (m *StoreMessage) String() string { return proto.CompactTextString(m) }
func (*StoreMessage) ProtoMessage()    {}
func (*StoreMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_33c57e4bae7b9afd, []int{2}
}

func (m *StoreMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StoreMessage.Unmarshal(m, b)
}
func (m *StoreMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StoreMessage.Marshal(b, m, deterministic)
}
func (m *StoreMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StoreMessage.Merge(m, src)
}
func (m *StoreMessage) XXX_Size() int {
	return xxx_messageInfo_StoreMessage.Size(m)
}
func (m *StoreMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_StoreMessage.DiscardUnknown(m)
}

var xxx_messageInfo_StoreMessage proto.InternalMessageInfo

func (m *StoreMessage) GetCids() [][]byte {
	if m != nil {
		return m.Cids
	}
	return nil
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
	return fileDescriptor_33c57e4bae7b9afd, []int{3}
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

type OrderMessage struct {
	OrderID              string                   `protobuf:"bytes,1,opt,name=orderID,proto3" json:"orderID,omitempty"`
	MessageType          OrderMessage_MessageType `protobuf:"varint,2,opt,name=messageType,proto3,enum=OrderMessage_MessageType" json:"messageType,omitempty"`
	Message              *any.Any                 `protobuf:"bytes,3,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *OrderMessage) Reset()         { *m = OrderMessage{} }
func (m *OrderMessage) String() string { return proto.CompactTextString(m) }
func (*OrderMessage) ProtoMessage()    {}
func (*OrderMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_33c57e4bae7b9afd, []int{4}
}

func (m *OrderMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OrderMessage.Unmarshal(m, b)
}
func (m *OrderMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OrderMessage.Marshal(b, m, deterministic)
}
func (m *OrderMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OrderMessage.Merge(m, src)
}
func (m *OrderMessage) XXX_Size() int {
	return xxx_messageInfo_OrderMessage.Size(m)
}
func (m *OrderMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_OrderMessage.DiscardUnknown(m)
}

var xxx_messageInfo_OrderMessage proto.InternalMessageInfo

func (m *OrderMessage) GetOrderID() string {
	if m != nil {
		return m.OrderID
	}
	return ""
}

func (m *OrderMessage) GetMessageType() OrderMessage_MessageType {
	if m != nil {
		return m.MessageType
	}
	return OrderMessage_ORDER_OPEN
}

func (m *OrderMessage) GetMessage() *any.Any {
	if m != nil {
		return m.Message
	}
	return nil
}

type OrderList struct {
	Messages             []*OrderMessage `protobuf:"bytes,1,rep,name=messages,proto3" json:"messages,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *OrderList) Reset()         { *m = OrderList{} }
func (m *OrderList) String() string { return proto.CompactTextString(m) }
func (*OrderList) ProtoMessage()    {}
func (*OrderList) Descriptor() ([]byte, []int) {
	return fileDescriptor_33c57e4bae7b9afd, []int{5}
}

func (m *OrderList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OrderList.Unmarshal(m, b)
}
func (m *OrderList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OrderList.Marshal(b, m, deterministic)
}
func (m *OrderList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OrderList.Merge(m, src)
}
func (m *OrderList) XXX_Size() int {
	return xxx_messageInfo_OrderList.Size(m)
}
func (m *OrderList) XXX_DiscardUnknown() {
	xxx_messageInfo_OrderList.DiscardUnknown(m)
}

var xxx_messageInfo_OrderList proto.InternalMessageInfo

func (m *OrderList) GetMessages() []*OrderMessage {
	if m != nil {
		return m.Messages
	}
	return nil
}

type AddressRequestMessage struct {
	Coin                 string   `protobuf:"bytes,1,opt,name=coin,proto3" json:"coin,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AddressRequestMessage) Reset()         { *m = AddressRequestMessage{} }
func (m *AddressRequestMessage) String() string { return proto.CompactTextString(m) }
func (*AddressRequestMessage) ProtoMessage()    {}
func (*AddressRequestMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_33c57e4bae7b9afd, []int{6}
}

func (m *AddressRequestMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddressRequestMessage.Unmarshal(m, b)
}
func (m *AddressRequestMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddressRequestMessage.Marshal(b, m, deterministic)
}
func (m *AddressRequestMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddressRequestMessage.Merge(m, src)
}
func (m *AddressRequestMessage) XXX_Size() int {
	return xxx_messageInfo_AddressRequestMessage.Size(m)
}
func (m *AddressRequestMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_AddressRequestMessage.DiscardUnknown(m)
}

var xxx_messageInfo_AddressRequestMessage proto.InternalMessageInfo

func (m *AddressRequestMessage) GetCoin() string {
	if m != nil {
		return m.Coin
	}
	return ""
}

type AddressResponseMessage struct {
	Address              string   `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	Coin                 string   `protobuf:"bytes,2,opt,name=coin,proto3" json:"coin,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AddressResponseMessage) Reset()         { *m = AddressResponseMessage{} }
func (m *AddressResponseMessage) String() string { return proto.CompactTextString(m) }
func (*AddressResponseMessage) ProtoMessage()    {}
func (*AddressResponseMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_33c57e4bae7b9afd, []int{7}
}

func (m *AddressResponseMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddressResponseMessage.Unmarshal(m, b)
}
func (m *AddressResponseMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddressResponseMessage.Marshal(b, m, deterministic)
}
func (m *AddressResponseMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddressResponseMessage.Merge(m, src)
}
func (m *AddressResponseMessage) XXX_Size() int {
	return xxx_messageInfo_AddressResponseMessage.Size(m)
}
func (m *AddressResponseMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_AddressResponseMessage.DiscardUnknown(m)
}

var xxx_messageInfo_AddressResponseMessage proto.InternalMessageInfo

func (m *AddressResponseMessage) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *AddressResponseMessage) GetCoin() string {
	if m != nil {
		return m.Coin
	}
	return ""
}

type Envelope struct {
	SenderPubkey         []byte   `protobuf:"bytes,1,opt,name=senderPubkey,proto3" json:"senderPubkey,omitempty"`
	Message              *Message `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	Signature            []byte   `protobuf:"bytes,3,opt,name=signature,proto3" json:"signature,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Envelope) Reset()         { *m = Envelope{} }
func (m *Envelope) String() string { return proto.CompactTextString(m) }
func (*Envelope) ProtoMessage()    {}
func (*Envelope) Descriptor() ([]byte, []int) {
	return fileDescriptor_33c57e4bae7b9afd, []int{8}
}

func (m *Envelope) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Envelope.Unmarshal(m, b)
}
func (m *Envelope) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Envelope.Marshal(b, m, deterministic)
}
func (m *Envelope) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Envelope.Merge(m, src)
}
func (m *Envelope) XXX_Size() int {
	return xxx_messageInfo_Envelope.Size(m)
}
func (m *Envelope) XXX_DiscardUnknown() {
	xxx_messageInfo_Envelope.DiscardUnknown(m)
}

var xxx_messageInfo_Envelope proto.InternalMessageInfo

func (m *Envelope) GetSenderPubkey() []byte {
	if m != nil {
		return m.SenderPubkey
	}
	return nil
}

func (m *Envelope) GetMessage() *Message {
	if m != nil {
		return m.Message
	}
	return nil
}

func (m *Envelope) GetSignature() []byte {
	if m != nil {
		return m.Signature
	}
	return nil
}

func init() {
	proto.RegisterEnum("Message_MessageType", Message_MessageType_name, Message_MessageType_value)
	proto.RegisterEnum("ChatMessage_Flag", ChatMessage_Flag_name, ChatMessage_Flag_value)
	proto.RegisterEnum("OrderMessage_MessageType", OrderMessage_MessageType_name, OrderMessage_MessageType_value)
	proto.RegisterType((*Message)(nil), "Message")
	proto.RegisterType((*ChatMessage)(nil), "ChatMessage")
	proto.RegisterType((*StoreMessage)(nil), "StoreMessage")
	proto.RegisterType((*AckMessage)(nil), "AckMessage")
	proto.RegisterType((*OrderMessage)(nil), "OrderMessage")
	proto.RegisterType((*OrderList)(nil), "OrderList")
	proto.RegisterType((*AddressRequestMessage)(nil), "AddressRequestMessage")
	proto.RegisterType((*AddressResponseMessage)(nil), "AddressResponseMessage")
	proto.RegisterType((*Envelope)(nil), "Envelope")
}

func init() { proto.RegisterFile("message.proto", fileDescriptor_33c57e4bae7b9afd) }

var fileDescriptor_33c57e4bae7b9afd = []byte{
	// 749 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x53, 0x5d, 0x8e, 0xa3, 0x46,
	0x10, 0x1e, 0x30, 0xb6, 0xa1, 0xc0, 0x4e, 0x4f, 0x67, 0x76, 0xe4, 0x1d, 0xad, 0x14, 0x0b, 0x29,
	0x91, 0x57, 0x91, 0x58, 0xc9, 0x89, 0x56, 0x91, 0xf2, 0x44, 0x4c, 0x33, 0x21, 0xc1, 0x40, 0x1a,
	0xac, 0x68, 0xf7, 0x65, 0x84, 0x4d, 0xaf, 0x63, 0x8d, 0x07, 0x08, 0xe0, 0x48, 0xbe, 0x42, 0x5e,
	0x72, 0xa2, 0x1c, 0x27, 0x17, 0xc8, 0x09, 0x22, 0xfe, 0x8c, 0x3d, 0x51, 0xb2, 0x4f, 0x54, 0x7d,
	0xf5, 0x51, 0x5d, 0xfd, 0xd5, 0xd7, 0x30, 0x7a, 0x62, 0x79, 0x1e, 0x6e, 0x99, 0x96, 0x66, 0x49,
	0x91, 0xdc, 0xbd, 0xdc, 0x26, 0xc9, 0x76, 0xcf, 0xde, 0x54, 0xd9, 0xfa, 0xf0, 0xe1, 0x4d, 0x18,
	0x1f, 0x9b, 0xd2, 0x67, 0xcf, 0x4b, 0xc5, 0xee, 0x89, 0xe5, 0x45, 0xf8, 0x94, 0xd6, 0x04, 0xf5,
	0x4f, 0x1e, 0x86, 0xcb, 0xba, 0x1b, 0x7e, 0x0b, 0x72, 0xd3, 0x38, 0x38, 0xa6, 0x6c, 0xc2, 0x4d,
	0xb9, 0xd9, 0x78, 0x7e, 0xa3, 0x35, 0xe5, 0xf6, 0x5b, 0xd6, 0xe8, 0x39, 0x11, 0xbf, 0x02, 0xa9,
	0x49, 0x2d, 0x63, 0xc2, 0x4f, 0xb9, 0x99, 0x44, 0x3b, 0x00, 0xdf, 0x81, 0x98, 0xb3, 0x5f, 0x0f,
	0x2c, 0xde, 0xb0, 0x49, 0x6f, 0xca, 0xcd, 0x46, 0xf4, 0x94, 0x63, 0x0d, 0x86, 0x69, 0x78, 0xdc,
	0x27, 0x61, 0x34, 0x11, 0xa6, 0xdc, 0x4c, 0x9e, 0xdf, 0x68, 0xf5, 0xc0, 0x5a, 0x3b, 0xb0, 0xa6,
	0xc7, 0x47, 0xda, 0x92, 0xd4, 0x3f, 0x38, 0x90, 0xcf, 0xc6, 0xc0, 0x43, 0xe8, 0xe9, 0x8b, 0x1f,
	0xd1, 0x15, 0x16, 0x41, 0xf0, 0x2c, 0xe7, 0x1e, 0x71, 0x55, 0xe4, 0x3a, 0xf7, 0x88, 0x2f, 0xa3,
	0xc5, 0xf7, 0x7a, 0x80, 0x7a, 0x18, 0x60, 0x60, 0xba, 0xb6, 0xed, 0xfe, 0x8c, 0x04, 0xac, 0x80,
	0xb8, 0x72, 0x9a, 0xac, 0x8f, 0x25, 0xe8, 0xfb, 0x81, 0x4b, 0x09, 0x1a, 0x94, 0xa1, 0x4b, 0x0d,
	0x42, 0xd1, 0x10, 0x7f, 0x0a, 0x9f, 0xe8, 0x86, 0x41, 0x89, 0xef, 0x3f, 0x50, 0xf2, 0xd3, 0x8a,
	0xf8, 0x01, 0x12, 0xf1, 0x0d, 0xa0, 0x0e, 0xf4, 0x3d, 0xd7, 0xf1, 0x09, 0x92, 0xd4, 0xbf, 0x38,
	0x90, 0x17, 0xbf, 0x84, 0x45, 0xab, 0xe1, 0x04, 0x86, 0x49, 0x16, 0xb1, 0xcc, 0x32, 0x2a, 0xfd,
	0x24, 0xda, 0xa6, 0x65, 0xa5, 0x11, 0xa5, 0xd1, 0xa8, 0x4d, 0xf1, 0x37, 0x20, 0x9d, 0xd6, 0x52,
	0x49, 0x24, 0xcf, 0xef, 0xfe, 0xa5, 0x43, 0xd0, 0x32, 0x68, 0x47, 0xc6, 0x9f, 0x83, 0xf0, 0x61,
	0x1f, 0x6e, 0x2b, 0xf1, 0xc6, 0xf3, 0x6b, 0xed, 0x6c, 0x12, 0xcd, 0xdc, 0x87, 0x5b, 0x5a, 0x95,
	0xf1, 0x2d, 0x0c, 0x32, 0x16, 0x46, 0x96, 0x31, 0xe9, 0x57, 0x27, 0x37, 0x99, 0xfa, 0x1a, 0x84,
	0x92, 0x85, 0x65, 0x18, 0x2e, 0x89, 0xef, 0xeb, 0xf7, 0x04, 0x5d, 0x95, 0x62, 0x05, 0xef, 0x3a,
	0x31, 0x29, 0xd1, 0x0d, 0xc4, 0xab, 0x2a, 0x28, 0x7e, 0x91, 0x64, 0xac, 0xbd, 0x27, 0x06, 0x61,
	0xb3, 0x8b, 0xf2, 0x09, 0x37, 0xed, 0xcd, 0x14, 0x5a, 0xc5, 0xea, 0xd7, 0x00, 0xfa, 0xe6, 0xb1,
	0x65, 0x7c, 0x01, 0xe3, 0x70, 0xf3, 0xc8, 0xa2, 0xe5, 0xc9, 0x1a, 0xb5, 0x20, 0xcf, 0x50, 0xf5,
	0xf7, 0x1e, 0x28, 0x6e, 0xa9, 0xd1, 0xc7, 0x25, 0xfc, 0xf6, 0xd2, 0xa0, 0x7c, 0x75, 0xeb, 0x97,
	0xda, 0xf9, 0xdf, 0xff, 0xed, 0x52, 0xad, 0xd3, 0xbf, 0xf7, 0x7f, 0x5e, 0x6b, 0x48, 0xea, 0xdf,
	0xcf, 0xbc, 0x36, 0x06, 0xa8, 0xfc, 0xf1, 0xe0, 0x7a, 0xc4, 0x41, 0x57, 0x18, 0x81, 0x52, 0xe7,
	0x94, 0xfc, 0x40, 0x16, 0x01, 0xe2, 0x3a, 0x64, 0xa1, 0x3b, 0x0b, 0x62, 0x23, 0x1e, 0xdf, 0x02,
	0x6e, 0x10, 0xd7, 0x31, 0x2d, 0xba, 0xd4, 0x03, 0xcb, 0x75, 0x50, 0x0f, 0xbf, 0x80, 0xeb, 0x1a,
	0x37, 0x57, 0xb6, 0x69, 0xd9, 0xf6, 0x92, 0x38, 0x01, 0x12, 0x30, 0x86, 0x71, 0x4b, 0x5f, 0x7a,
	0x36, 0x09, 0x08, 0xea, 0x97, 0x4d, 0x0d, 0xcb, 0xf7, 0x56, 0x01, 0xa9, 0x0f, 0x1e, 0x94, 0xac,
	0x16, 0x59, 0x79, 0x86, 0x1e, 0x10, 0x34, 0xc4, 0xd7, 0x30, 0x6a, 0xb1, 0x85, 0xed, 0xfa, 0x04,
	0x89, 0xe5, 0x1e, 0x29, 0x31, 0x57, 0x8e, 0x81, 0xa4, 0xb2, 0x89, 0xa7, 0xbf, 0x2b, 0x4f, 0x79,
	0xf0, 0xcb, 0xa3, 0xa0, 0x9c, 0xa0, 0x45, 0x4c, 0xcb, 0xd1, 0x6d, 0xeb, 0x3d, 0x31, 0x90, 0xac,
	0xbe, 0x05, 0xa9, 0x52, 0xd3, 0xde, 0xe5, 0x05, 0x7e, 0x0d, 0x62, 0x23, 0x46, 0xbd, 0x67, 0x79,
	0x3e, 0xba, 0xd0, 0x9a, 0x9e, 0xca, 0xea, 0x97, 0xf0, 0x42, 0x8f, 0xa2, 0x8c, 0xe5, 0x39, 0x2d,
	0xdf, 0x76, 0x5e, 0x9c, 0xfb, 0x24, 0xd9, 0xc5, 0xcd, 0x26, 0xab, 0x58, 0x35, 0xe1, 0xf6, 0x44,
	0xce, 0xd3, 0x24, 0xce, 0xd9, 0xd9, 0xea, 0xc3, 0xba, 0xd2, 0xae, 0xbe, 0x49, 0x4f, 0x7d, 0xf8,
	0xb3, 0x3e, 0x29, 0x88, 0x24, 0xfe, 0x8d, 0xed, 0x93, 0x94, 0x61, 0x15, 0x94, 0x9c, 0xc5, 0x11,
	0xcb, 0xbc, 0xc3, 0xfa, 0x91, 0x1d, 0xab, 0xdf, 0x15, 0x7a, 0x81, 0x61, 0xf5, 0xf2, 0x05, 0xca,
	0x73, 0xb1, 0x75, 0x4b, 0xf7, 0x16, 0x5f, 0x81, 0x94, 0xef, 0xb6, 0x71, 0x58, 0x1c, 0xb2, 0xda,
	0x27, 0x0a, 0xed, 0x80, 0xef, 0x84, 0xf7, 0x7c, 0xba, 0x5e, 0x0f, 0x2a, 0xc3, 0x7c, 0xf5, 0x4f,
	0x00, 0x00, 0x00, 0xff, 0xff, 0xd4, 0xa8, 0xd6, 0xf0, 0x87, 0x05, 0x00, 0x00,
}
