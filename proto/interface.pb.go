// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.7
// source: interface.proto

package token

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Status int32

const (
	Status_FAIL      Status = 0
	Status_SUCCESS   Status = 1
	Status_EXCEPTION Status = 2
)

// Enum value maps for Status.
var (
	Status_name = map[int32]string{
		0: "FAIL",
		1: "SUCCESS",
		2: "EXCEPTION",
	}
	Status_value = map[string]int32{
		"FAIL":      0,
		"SUCCESS":   1,
		"EXCEPTION": 2,
	}
)

func (x Status) Enum() *Status {
	p := new(Status)
	*p = x
	return p
}

func (x Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Status) Descriptor() protoreflect.EnumDescriptor {
	return file_interface_proto_enumTypes[0].Descriptor()
}

func (Status) Type() protoreflect.EnumType {
	return &file_interface_proto_enumTypes[0]
}

func (x Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Status.Descriptor instead.
func (Status) EnumDescriptor() ([]byte, []int) {
	return file_interface_proto_rawDescGZIP(), []int{0}
}

type Reply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Reply) Reset() {
	*x = Reply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_interface_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Reply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Reply) ProtoMessage() {}

func (x *Reply) ProtoReflect() protoreflect.Message {
	mi := &file_interface_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Reply.ProtoReflect.Descriptor instead.
func (*Reply) Descriptor() ([]byte, []int) {
	return file_interface_proto_rawDescGZIP(), []int{0}
}

type Void struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Void) Reset() {
	*x = Void{}
	if protoimpl.UnsafeEnabled {
		mi := &file_interface_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Void) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Void) ProtoMessage() {}

func (x *Void) ProtoReflect() protoreflect.Message {
	mi := &file_interface_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Void.ProtoReflect.Descriptor instead.
func (*Void) Descriptor() ([]byte, []int) {
	return file_interface_proto_rawDescGZIP(), []int{1}
}

type Beat struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Beat) Reset() {
	*x = Beat{}
	if protoimpl.UnsafeEnabled {
		mi := &file_interface_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Beat) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Beat) ProtoMessage() {}

func (x *Beat) ProtoReflect() protoreflect.Message {
	mi := &file_interface_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Beat.ProtoReflect.Descriptor instead.
func (*Beat) Descriptor() ([]byte, []int) {
	return file_interface_proto_rawDescGZIP(), []int{2}
}

type ElectionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ElectionRequest) Reset() {
	*x = ElectionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_interface_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ElectionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ElectionRequest) ProtoMessage() {}

func (x *ElectionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_interface_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ElectionRequest.ProtoReflect.Descriptor instead.
func (*ElectionRequest) Descriptor() ([]byte, []int) {
	return file_interface_proto_rawDescGZIP(), []int{3}
}

type Primary struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Primary) Reset() {
	*x = Primary{}
	if protoimpl.UnsafeEnabled {
		mi := &file_interface_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Primary) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Primary) ProtoMessage() {}

func (x *Primary) ProtoReflect() protoreflect.Message {
	mi := &file_interface_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Primary.ProtoReflect.Descriptor instead.
func (*Primary) Descriptor() ([]byte, []int) {
	return file_interface_proto_rawDescGZIP(), []int{4}
}

type Ack struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status  Status `protobuf:"varint,1,opt,name=status,proto3,enum=token.Status" json:"status,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *Ack) Reset() {
	*x = Ack{}
	if protoimpl.UnsafeEnabled {
		mi := &file_interface_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Ack) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Ack) ProtoMessage() {}

func (x *Ack) ProtoReflect() protoreflect.Message {
	mi := &file_interface_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Ack.ProtoReflect.Descriptor instead.
func (*Ack) Descriptor() ([]byte, []int) {
	return file_interface_proto_rawDescGZIP(), []int{5}
}

func (x *Ack) GetStatus() Status {
	if x != nil {
		return x.Status
	}
	return Status_FAIL
}

func (x *Ack) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type Data struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CurrentBid    *CurrentBid `protobuf:"bytes,1,opt,name=currentBid,proto3" json:"currentBid,omitempty"`
	TimeRemaining int32       `protobuf:"varint,2,opt,name=timeRemaining,proto3" json:"timeRemaining,omitempty"`
	Clients       []string    `protobuf:"bytes,3,rep,name=Clients,proto3" json:"Clients,omitempty"`
}

func (x *Data) Reset() {
	*x = Data{}
	if protoimpl.UnsafeEnabled {
		mi := &file_interface_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Data) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Data) ProtoMessage() {}

func (x *Data) ProtoReflect() protoreflect.Message {
	mi := &file_interface_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Data.ProtoReflect.Descriptor instead.
func (*Data) Descriptor() ([]byte, []int) {
	return file_interface_proto_rawDescGZIP(), []int{6}
}

func (x *Data) GetCurrentBid() *CurrentBid {
	if x != nil {
		return x.CurrentBid
	}
	return nil
}

func (x *Data) GetTimeRemaining() int32 {
	if x != nil {
		return x.TimeRemaining
	}
	return 0
}

func (x *Data) GetClients() []string {
	if x != nil {
		return x.Clients
	}
	return nil
}

type CurrentBid struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id     string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Amount int32  `protobuf:"varint,2,opt,name=amount,proto3" json:"amount,omitempty"`
}

func (x *CurrentBid) Reset() {
	*x = CurrentBid{}
	if protoimpl.UnsafeEnabled {
		mi := &file_interface_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CurrentBid) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CurrentBid) ProtoMessage() {}

func (x *CurrentBid) ProtoReflect() protoreflect.Message {
	mi := &file_interface_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CurrentBid.ProtoReflect.Descriptor instead.
func (*CurrentBid) Descriptor() ([]byte, []int) {
	return file_interface_proto_rawDescGZIP(), []int{7}
}

func (x *CurrentBid) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *CurrentBid) GetAmount() int32 {
	if x != nil {
		return x.Amount
	}
	return 0
}

type Coord struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *Coord) Reset() {
	*x = Coord{}
	if protoimpl.UnsafeEnabled {
		mi := &file_interface_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Coord) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Coord) ProtoMessage() {}

func (x *Coord) ProtoReflect() protoreflect.Message {
	mi := &file_interface_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Coord.ProtoReflect.Descriptor instead.
func (*Coord) Descriptor() ([]byte, []int) {
	return file_interface_proto_rawDescGZIP(), []int{8}
}

func (x *Coord) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type Bid struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Userid string `protobuf:"bytes,1,opt,name=Userid,proto3" json:"Userid,omitempty"`
	ItemId int32  `protobuf:"varint,2,opt,name=ItemId,proto3" json:"ItemId,omitempty"`
	Amount int32  `protobuf:"varint,3,opt,name=Amount,proto3" json:"Amount,omitempty"`
}

func (x *Bid) Reset() {
	*x = Bid{}
	if protoimpl.UnsafeEnabled {
		mi := &file_interface_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Bid) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Bid) ProtoMessage() {}

func (x *Bid) ProtoReflect() protoreflect.Message {
	mi := &file_interface_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Bid.ProtoReflect.Descriptor instead.
func (*Bid) Descriptor() ([]byte, []int) {
	return file_interface_proto_rawDescGZIP(), []int{9}
}

func (x *Bid) GetUserid() string {
	if x != nil {
		return x.Userid
	}
	return ""
}

func (x *Bid) GetItemId() int32 {
	if x != nil {
		return x.ItemId
	}
	return 0
}

func (x *Bid) GetAmount() int32 {
	if x != nil {
		return x.Amount
	}
	return 0
}

type Amount struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value int32 `protobuf:"varint,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *Amount) Reset() {
	*x = Amount{}
	if protoimpl.UnsafeEnabled {
		mi := &file_interface_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Amount) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Amount) ProtoMessage() {}

func (x *Amount) ProtoReflect() protoreflect.Message {
	mi := &file_interface_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Amount.ProtoReflect.Descriptor instead.
func (*Amount) Descriptor() ([]byte, []int) {
	return file_interface_proto_rawDescGZIP(), []int{10}
}

func (x *Amount) GetValue() int32 {
	if x != nil {
		return x.Value
	}
	return 0
}

type Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ItemId int32 `protobuf:"varint,1,opt,name=ItemId,proto3" json:"ItemId,omitempty"`
}

func (x *Request) Reset() {
	*x = Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_interface_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Request) ProtoMessage() {}

func (x *Request) ProtoReflect() protoreflect.Message {
	mi := &file_interface_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Request.ProtoReflect.Descriptor instead.
func (*Request) Descriptor() ([]byte, []int) {
	return file_interface_proto_rawDescGZIP(), []int{11}
}

func (x *Request) GetItemId() int32 {
	if x != nil {
		return x.ItemId
	}
	return 0
}

type AuctionResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id     string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Amount int32  `protobuf:"varint,2,opt,name=amount,proto3" json:"amount,omitempty"`
}

func (x *AuctionResult) Reset() {
	*x = AuctionResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_interface_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuctionResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuctionResult) ProtoMessage() {}

func (x *AuctionResult) ProtoReflect() protoreflect.Message {
	mi := &file_interface_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuctionResult.ProtoReflect.Descriptor instead.
func (*AuctionResult) Descriptor() ([]byte, []int) {
	return file_interface_proto_rawDescGZIP(), []int{12}
}

func (x *AuctionResult) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *AuctionResult) GetAmount() int32 {
	if x != nil {
		return x.Amount
	}
	return 0
}

type Outcome struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TimeRemaining int32 `protobuf:"varint,1,opt,name=TimeRemaining,proto3" json:"TimeRemaining,omitempty"`
	// Types that are assignable to Outcome:
	//
	//	*Outcome_AuctionResult
	//	*Outcome_CurrentBid
	Outcome isOutcome_Outcome `protobuf_oneof:"outcome"`
}

func (x *Outcome) Reset() {
	*x = Outcome{}
	if protoimpl.UnsafeEnabled {
		mi := &file_interface_proto_msgTypes[13]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Outcome) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Outcome) ProtoMessage() {}

func (x *Outcome) ProtoReflect() protoreflect.Message {
	mi := &file_interface_proto_msgTypes[13]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Outcome.ProtoReflect.Descriptor instead.
func (*Outcome) Descriptor() ([]byte, []int) {
	return file_interface_proto_rawDescGZIP(), []int{13}
}

func (x *Outcome) GetTimeRemaining() int32 {
	if x != nil {
		return x.TimeRemaining
	}
	return 0
}

func (m *Outcome) GetOutcome() isOutcome_Outcome {
	if m != nil {
		return m.Outcome
	}
	return nil
}

func (x *Outcome) GetAuctionResult() *AuctionResult {
	if x, ok := x.GetOutcome().(*Outcome_AuctionResult); ok {
		return x.AuctionResult
	}
	return nil
}

func (x *Outcome) GetCurrentBid() *CurrentBid {
	if x, ok := x.GetOutcome().(*Outcome_CurrentBid); ok {
		return x.CurrentBid
	}
	return nil
}

type isOutcome_Outcome interface {
	isOutcome_Outcome()
}

type Outcome_AuctionResult struct {
	AuctionResult *AuctionResult `protobuf:"bytes,2,opt,name=auctionResult,proto3,oneof"`
}

type Outcome_CurrentBid struct {
	CurrentBid *CurrentBid `protobuf:"bytes,3,opt,name=currentBid,proto3,oneof"`
}

func (*Outcome_AuctionResult) isOutcome_Outcome() {}

func (*Outcome_CurrentBid) isOutcome_Outcome() {}

var File_interface_proto protoreflect.FileDescriptor

var file_interface_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x07, 0x0a, 0x05, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x22, 0x06, 0x0a, 0x04, 0x56, 0x6f, 0x69, 0x64, 0x22, 0x06, 0x0a, 0x04, 0x42, 0x65, 0x61,
	0x74, 0x22, 0x11, 0x0a, 0x0f, 0x45, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x22, 0x09, 0x0a, 0x07, 0x50, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x22,
	0x46, 0x0a, 0x03, 0x41, 0x63, 0x6b, 0x12, 0x25, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0d, 0x2e, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x2e, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x18, 0x0a,
	0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x79, 0x0a, 0x04, 0x44, 0x61, 0x74, 0x61, 0x12,
	0x31, 0x0a, 0x0a, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x42, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x2e, 0x43, 0x75, 0x72, 0x72,
	0x65, 0x6e, 0x74, 0x42, 0x69, 0x64, 0x52, 0x0a, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x42,
	0x69, 0x64, 0x12, 0x24, 0x0a, 0x0d, 0x74, 0x69, 0x6d, 0x65, 0x52, 0x65, 0x6d, 0x61, 0x69, 0x6e,
	0x69, 0x6e, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0d, 0x74, 0x69, 0x6d, 0x65, 0x52,
	0x65, 0x6d, 0x61, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x12, 0x18, 0x0a, 0x07, 0x43, 0x6c, 0x69, 0x65,
	0x6e, 0x74, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x43, 0x6c, 0x69, 0x65, 0x6e,
	0x74, 0x73, 0x22, 0x34, 0x0a, 0x0a, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x42, 0x69, 0x64,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x17, 0x0a, 0x05, 0x43, 0x6f, 0x6f, 0x72,
	0x64, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69,
	0x64, 0x22, 0x4d, 0x0a, 0x03, 0x42, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x55, 0x73, 0x65, 0x72,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x55, 0x73, 0x65, 0x72, 0x69, 0x64,
	0x12, 0x16, 0x0a, 0x06, 0x49, 0x74, 0x65, 0x6d, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x06, 0x49, 0x74, 0x65, 0x6d, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x41, 0x6d, 0x6f, 0x75,
	0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74,
	0x22, 0x1e, 0x0a, 0x06, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x22, 0x21, 0x0a, 0x07, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x49,
	0x74, 0x65, 0x6d, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x49, 0x74, 0x65,
	0x6d, 0x49, 0x64, 0x22, 0x37, 0x0a, 0x0d, 0x41, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0xad, 0x01, 0x0a,
	0x07, 0x4f, 0x75, 0x74, 0x63, 0x6f, 0x6d, 0x65, 0x12, 0x24, 0x0a, 0x0d, 0x54, 0x69, 0x6d, 0x65,
	0x52, 0x65, 0x6d, 0x61, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x0d, 0x54, 0x69, 0x6d, 0x65, 0x52, 0x65, 0x6d, 0x61, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x12, 0x3c,
	0x0a, 0x0d, 0x61, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x2e, 0x41, 0x75,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x48, 0x00, 0x52, 0x0d, 0x61,
	0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x33, 0x0a, 0x0a,
	0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x42, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x11, 0x2e, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x2e, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74,
	0x42, 0x69, 0x64, 0x48, 0x00, 0x52, 0x0a, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x42, 0x69,
	0x64, 0x42, 0x09, 0x0a, 0x07, 0x6f, 0x75, 0x74, 0x63, 0x6f, 0x6d, 0x65, 0x2a, 0x2e, 0x0a, 0x06,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x08, 0x0a, 0x04, 0x46, 0x41, 0x49, 0x4c, 0x10, 0x00,
	0x12, 0x0b, 0x0a, 0x07, 0x53, 0x55, 0x43, 0x43, 0x45, 0x53, 0x53, 0x10, 0x01, 0x12, 0x0d, 0x0a,
	0x09, 0x45, 0x58, 0x43, 0x45, 0x50, 0x54, 0x49, 0x4f, 0x4e, 0x10, 0x02, 0x32, 0x81, 0x02, 0x0a,
	0x07, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x12, 0x30, 0x0a, 0x08, 0x45, 0x6c, 0x65, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x16, 0x2e, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x2e, 0x45, 0x6c, 0x65,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0a, 0x2e, 0x74,
	0x6f, 0x6b, 0x65, 0x6e, 0x2e, 0x41, 0x63, 0x6b, 0x22, 0x00, 0x12, 0x2a, 0x0a, 0x0c, 0x43, 0x6f,
	0x6f, 0x72, 0x64, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x0c, 0x2e, 0x74, 0x6f, 0x6b,
	0x65, 0x6e, 0x2e, 0x43, 0x6f, 0x6f, 0x72, 0x64, 0x1a, 0x0a, 0x2e, 0x74, 0x6f, 0x6b, 0x65, 0x6e,
	0x2e, 0x41, 0x63, 0x6b, 0x22, 0x00, 0x12, 0x22, 0x0a, 0x03, 0x42, 0x69, 0x64, 0x12, 0x0d, 0x2e,
	0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x2e, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x1a, 0x0a, 0x2e, 0x74,
	0x6f, 0x6b, 0x65, 0x6e, 0x2e, 0x41, 0x63, 0x6b, 0x22, 0x00, 0x12, 0x27, 0x0a, 0x06, 0x52, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x12, 0x0b, 0x2e, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x2e, 0x56, 0x6f, 0x69,
	0x64, 0x1a, 0x0e, 0x2e, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x2e, 0x4f, 0x75, 0x74, 0x63, 0x6f, 0x6d,
	0x65, 0x22, 0x00, 0x12, 0x23, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x0b, 0x2e,
	0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x1a, 0x0a, 0x2e, 0x74, 0x6f, 0x6b,
	0x65, 0x6e, 0x2e, 0x41, 0x63, 0x6b, 0x22, 0x00, 0x12, 0x26, 0x0a, 0x09, 0x48, 0x65, 0x61, 0x72,
	0x74, 0x62, 0x65, 0x61, 0x74, 0x12, 0x0b, 0x2e, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x2e, 0x42, 0x65,
	0x61, 0x74, 0x1a, 0x0a, 0x2e, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x2e, 0x41, 0x63, 0x6b, 0x22, 0x00,
	0x32, 0x35, 0x0a, 0x06, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x12, 0x2b, 0x0a, 0x09, 0x48, 0x65,
	0x61, 0x72, 0x74, 0x62, 0x65, 0x61, 0x74, 0x12, 0x0e, 0x2e, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x2e,
	0x50, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x1a, 0x0c, 0x2e, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x2e,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x42, 0x2f, 0x5a, 0x2d, 0x68, 0x74, 0x74, 0x70, 0x73,
	0x3a, 0x2f, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6c, 0x75,
	0x63, 0x7a, 0x69, 0x74, 0x6f, 0x2f, 0x61, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x68, 0x6f, 0x75,
	0x73, 0x65, 0x3b, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_interface_proto_rawDescOnce sync.Once
	file_interface_proto_rawDescData = file_interface_proto_rawDesc
)

func file_interface_proto_rawDescGZIP() []byte {
	file_interface_proto_rawDescOnce.Do(func() {
		file_interface_proto_rawDescData = protoimpl.X.CompressGZIP(file_interface_proto_rawDescData)
	})
	return file_interface_proto_rawDescData
}

var file_interface_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_interface_proto_msgTypes = make([]protoimpl.MessageInfo, 14)
var file_interface_proto_goTypes = []interface{}{
	(Status)(0),             // 0: token.Status
	(*Reply)(nil),           // 1: token.Reply
	(*Void)(nil),            // 2: token.Void
	(*Beat)(nil),            // 3: token.Beat
	(*ElectionRequest)(nil), // 4: token.ElectionRequest
	(*Primary)(nil),         // 5: token.Primary
	(*Ack)(nil),             // 6: token.Ack
	(*Data)(nil),            // 7: token.Data
	(*CurrentBid)(nil),      // 8: token.CurrentBid
	(*Coord)(nil),           // 9: token.Coord
	(*Bid)(nil),             // 10: token.Bid
	(*Amount)(nil),          // 11: token.Amount
	(*Request)(nil),         // 12: token.Request
	(*AuctionResult)(nil),   // 13: token.AuctionResult
	(*Outcome)(nil),         // 14: token.Outcome
}
var file_interface_proto_depIdxs = []int32{
	0,  // 0: token.Ack.status:type_name -> token.Status
	8,  // 1: token.Data.currentBid:type_name -> token.CurrentBid
	13, // 2: token.Outcome.auctionResult:type_name -> token.AuctionResult
	8,  // 3: token.Outcome.currentBid:type_name -> token.CurrentBid
	4,  // 4: token.Manager.Election:input_type -> token.ElectionRequest
	9,  // 5: token.Manager.Coordination:input_type -> token.Coord
	11, // 6: token.Manager.Bid:input_type -> token.Amount
	2,  // 7: token.Manager.Result:input_type -> token.Void
	7,  // 8: token.Manager.Update:input_type -> token.Data
	3,  // 9: token.Manager.Heartbeat:input_type -> token.Beat
	5,  // 10: token.Client.Heartbeat:input_type -> token.Primary
	6,  // 11: token.Manager.Election:output_type -> token.Ack
	6,  // 12: token.Manager.Coordination:output_type -> token.Ack
	6,  // 13: token.Manager.Bid:output_type -> token.Ack
	14, // 14: token.Manager.Result:output_type -> token.Outcome
	6,  // 15: token.Manager.Update:output_type -> token.Ack
	6,  // 16: token.Manager.Heartbeat:output_type -> token.Ack
	1,  // 17: token.Client.Heartbeat:output_type -> token.Reply
	11, // [11:18] is the sub-list for method output_type
	4,  // [4:11] is the sub-list for method input_type
	4,  // [4:4] is the sub-list for extension type_name
	4,  // [4:4] is the sub-list for extension extendee
	0,  // [0:4] is the sub-list for field type_name
}

func init() { file_interface_proto_init() }
func file_interface_proto_init() {
	if File_interface_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_interface_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Reply); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_interface_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Void); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_interface_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Beat); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_interface_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ElectionRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_interface_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Primary); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_interface_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Ack); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_interface_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Data); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_interface_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CurrentBid); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_interface_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Coord); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_interface_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Bid); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_interface_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Amount); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_interface_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Request); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_interface_proto_msgTypes[12].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuctionResult); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_interface_proto_msgTypes[13].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Outcome); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	file_interface_proto_msgTypes[13].OneofWrappers = []interface{}{
		(*Outcome_AuctionResult)(nil),
		(*Outcome_CurrentBid)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_interface_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   14,
			NumExtensions: 0,
			NumServices:   2,
		},
		GoTypes:           file_interface_proto_goTypes,
		DependencyIndexes: file_interface_proto_depIdxs,
		EnumInfos:         file_interface_proto_enumTypes,
		MessageInfos:      file_interface_proto_msgTypes,
	}.Build()
	File_interface_proto = out.File
	file_interface_proto_rawDesc = nil
	file_interface_proto_goTypes = nil
	file_interface_proto_depIdxs = nil
}
