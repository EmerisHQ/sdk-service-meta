// Code generated with goa v3.5.2, DO NOT EDIT.
//
// sdk-utilities protocol buffer definition
//
// Command:
// $ goa gen github.com/allinbits/sdk-service-meta

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.3
// source: sdk_utilities.proto

package sdk_utilitiespb

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

type SupplyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Chain to get data from
	ChainName string `protobuf:"bytes,1,opt,name=chain_name,json=chainName,proto3" json:"chain_name,omitempty"`
	// gRPC port for selected chain, defaults to 9090
	Port int32 `protobuf:"zigzag32,2,opt,name=port,proto3" json:"port,omitempty"`
}

func (x *SupplyRequest) Reset() {
	*x = SupplyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sdk_utilities_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SupplyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SupplyRequest) ProtoMessage() {}

func (x *SupplyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_sdk_utilities_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SupplyRequest.ProtoReflect.Descriptor instead.
func (*SupplyRequest) Descriptor() ([]byte, []int) {
	return file_sdk_utilities_proto_rawDescGZIP(), []int{0}
}

func (x *SupplyRequest) GetChainName() string {
	if x != nil {
		return x.ChainName
	}
	return ""
}

func (x *SupplyRequest) GetPort() int32 {
	if x != nil {
		return x.Port
	}
	return 0
}

type SupplyResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Coins []*Coin `protobuf:"bytes,1,rep,name=coins,proto3" json:"coins,omitempty"`
}

func (x *SupplyResponse) Reset() {
	*x = SupplyResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sdk_utilities_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SupplyResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SupplyResponse) ProtoMessage() {}

func (x *SupplyResponse) ProtoReflect() protoreflect.Message {
	mi := &file_sdk_utilities_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SupplyResponse.ProtoReflect.Descriptor instead.
func (*SupplyResponse) Descriptor() ([]byte, []int) {
	return file_sdk_utilities_proto_rawDescGZIP(), []int{1}
}

func (x *SupplyResponse) GetCoins() []*Coin {
	if x != nil {
		return x.Coins
	}
	return nil
}

// SDK service representation of a Cosmos SDK types.Coin
type Coin struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Denom  string `protobuf:"bytes,1,opt,name=denom,proto3" json:"denom,omitempty"`
	Amount string `protobuf:"bytes,2,opt,name=amount,proto3" json:"amount,omitempty"`
}

func (x *Coin) Reset() {
	*x = Coin{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sdk_utilities_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Coin) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Coin) ProtoMessage() {}

func (x *Coin) ProtoReflect() protoreflect.Message {
	mi := &file_sdk_utilities_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Coin.ProtoReflect.Descriptor instead.
func (*Coin) Descriptor() ([]byte, []int) {
	return file_sdk_utilities_proto_rawDescGZIP(), []int{2}
}

func (x *Coin) GetDenom() string {
	if x != nil {
		return x.Denom
	}
	return ""
}

func (x *Coin) GetAmount() string {
	if x != nil {
		return x.Amount
	}
	return ""
}

type QueryTxRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Chain to get data from
	ChainName string `protobuf:"bytes,1,opt,name=chain_name,json=chainName,proto3" json:"chain_name,omitempty"`
	// gRPC port for selected chain, defaults to 9090
	Port int32 `protobuf:"zigzag32,2,opt,name=port,proto3" json:"port,omitempty"`
	// Transaction hash to query
	Hash string `protobuf:"bytes,3,opt,name=hash,proto3" json:"hash,omitempty"`
}

func (x *QueryTxRequest) Reset() {
	*x = QueryTxRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sdk_utilities_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryTxRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryTxRequest) ProtoMessage() {}

func (x *QueryTxRequest) ProtoReflect() protoreflect.Message {
	mi := &file_sdk_utilities_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryTxRequest.ProtoReflect.Descriptor instead.
func (*QueryTxRequest) Descriptor() ([]byte, []int) {
	return file_sdk_utilities_proto_rawDescGZIP(), []int{3}
}

func (x *QueryTxRequest) GetChainName() string {
	if x != nil {
		return x.ChainName
	}
	return ""
}

func (x *QueryTxRequest) GetPort() int32 {
	if x != nil {
		return x.Port
	}
	return 0
}

func (x *QueryTxRequest) GetHash() string {
	if x != nil {
		return x.Hash
	}
	return ""
}

type QueryTxResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Field []byte `protobuf:"bytes,1,opt,name=field,proto3" json:"field,omitempty"`
}

func (x *QueryTxResponse) Reset() {
	*x = QueryTxResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sdk_utilities_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryTxResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryTxResponse) ProtoMessage() {}

func (x *QueryTxResponse) ProtoReflect() protoreflect.Message {
	mi := &file_sdk_utilities_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryTxResponse.ProtoReflect.Descriptor instead.
func (*QueryTxResponse) Descriptor() ([]byte, []int) {
	return file_sdk_utilities_proto_rawDescGZIP(), []int{4}
}

func (x *QueryTxResponse) GetField() []byte {
	if x != nil {
		return x.Field
	}
	return nil
}

type BroadcastTxRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Chain to get data from
	ChainName string `protobuf:"bytes,1,opt,name=chain_name,json=chainName,proto3" json:"chain_name,omitempty"`
	// gRPC port for selected chain, defaults to 9090
	Port int32 `protobuf:"zigzag32,2,opt,name=port,proto3" json:"port,omitempty"`
	// Transaction bytes
	TxBytes []byte `protobuf:"bytes,3,opt,name=tx_bytes,json=txBytes,proto3" json:"tx_bytes,omitempty"`
}

func (x *BroadcastTxRequest) Reset() {
	*x = BroadcastTxRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sdk_utilities_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BroadcastTxRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BroadcastTxRequest) ProtoMessage() {}

func (x *BroadcastTxRequest) ProtoReflect() protoreflect.Message {
	mi := &file_sdk_utilities_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BroadcastTxRequest.ProtoReflect.Descriptor instead.
func (*BroadcastTxRequest) Descriptor() ([]byte, []int) {
	return file_sdk_utilities_proto_rawDescGZIP(), []int{5}
}

func (x *BroadcastTxRequest) GetChainName() string {
	if x != nil {
		return x.ChainName
	}
	return ""
}

func (x *BroadcastTxRequest) GetPort() int32 {
	if x != nil {
		return x.Port
	}
	return 0
}

func (x *BroadcastTxRequest) GetTxBytes() []byte {
	if x != nil {
		return x.TxBytes
	}
	return nil
}

type BroadcastTxResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Hash string `protobuf:"bytes,1,opt,name=hash,proto3" json:"hash,omitempty"`
}

func (x *BroadcastTxResponse) Reset() {
	*x = BroadcastTxResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sdk_utilities_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BroadcastTxResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BroadcastTxResponse) ProtoMessage() {}

func (x *BroadcastTxResponse) ProtoReflect() protoreflect.Message {
	mi := &file_sdk_utilities_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BroadcastTxResponse.ProtoReflect.Descriptor instead.
func (*BroadcastTxResponse) Descriptor() ([]byte, []int) {
	return file_sdk_utilities_proto_rawDescGZIP(), []int{6}
}

func (x *BroadcastTxResponse) GetHash() string {
	if x != nil {
		return x.Hash
	}
	return ""
}

type TxMetadataRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Transaction bytes
	TxBytes []byte `protobuf:"bytes,1,opt,name=tx_bytes,json=txBytes,proto3" json:"tx_bytes,omitempty"`
}

func (x *TxMetadataRequest) Reset() {
	*x = TxMetadataRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sdk_utilities_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TxMetadataRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TxMetadataRequest) ProtoMessage() {}

func (x *TxMetadataRequest) ProtoReflect() protoreflect.Message {
	mi := &file_sdk_utilities_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TxMetadataRequest.ProtoReflect.Descriptor instead.
func (*TxMetadataRequest) Descriptor() ([]byte, []int) {
	return file_sdk_utilities_proto_rawDescGZIP(), []int{7}
}

func (x *TxMetadataRequest) GetTxBytes() []byte {
	if x != nil {
		return x.TxBytes
	}
	return nil
}

type TxMetadataResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MessagesMetadata []*TxMetadata `protobuf:"bytes,1,rep,name=messages_metadata,json=messagesMetadata,proto3" json:"messages_metadata,omitempty"`
}

func (x *TxMetadataResponse) Reset() {
	*x = TxMetadataResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sdk_utilities_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TxMetadataResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TxMetadataResponse) ProtoMessage() {}

func (x *TxMetadataResponse) ProtoReflect() protoreflect.Message {
	mi := &file_sdk_utilities_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TxMetadataResponse.ProtoReflect.Descriptor instead.
func (*TxMetadataResponse) Descriptor() ([]byte, []int) {
	return file_sdk_utilities_proto_rawDescGZIP(), []int{8}
}

func (x *TxMetadataResponse) GetMessagesMetadata() []*TxMetadata {
	if x != nil {
		return x.MessagesMetadata
	}
	return nil
}

// Metadata related to some transaction bytes
type TxMetadata struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TxType              string               `protobuf:"bytes,1,opt,name=tx_type,json=txType,proto3" json:"tx_type,omitempty"`
	IbcTransferMetadata *IBCTransferMetadata `protobuf:"bytes,2,opt,name=ibc_transfer_metadata,json=ibcTransferMetadata,proto3" json:"ibc_transfer_metadata,omitempty"`
}

func (x *TxMetadata) Reset() {
	*x = TxMetadata{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sdk_utilities_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TxMetadata) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TxMetadata) ProtoMessage() {}

func (x *TxMetadata) ProtoReflect() protoreflect.Message {
	mi := &file_sdk_utilities_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TxMetadata.ProtoReflect.Descriptor instead.
func (*TxMetadata) Descriptor() ([]byte, []int) {
	return file_sdk_utilities_proto_rawDescGZIP(), []int{9}
}

func (x *TxMetadata) GetTxType() string {
	if x != nil {
		return x.TxType
	}
	return ""
}

func (x *TxMetadata) GetIbcTransferMetadata() *IBCTransferMetadata {
	if x != nil {
		return x.IbcTransferMetadata
	}
	return nil
}

// Metadata related to an IBC Transfer
type IBCTransferMetadata struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SourcePort       string     `protobuf:"bytes,1,opt,name=source_port,json=sourcePort,proto3" json:"source_port,omitempty"`
	SourceCannel     string     `protobuf:"bytes,2,opt,name=source_cannel,json=sourceCannel,proto3" json:"source_cannel,omitempty"`
	Token            *Coin      `protobuf:"bytes,3,opt,name=token,proto3" json:"token,omitempty"`
	Sender           string     `protobuf:"bytes,4,opt,name=sender,proto3" json:"sender,omitempty"`
	Receiver         string     `protobuf:"bytes,5,opt,name=receiver,proto3" json:"receiver,omitempty"`
	TimeoutHeight    *IBCHeight `protobuf:"bytes,6,opt,name=timeout_height,json=timeoutHeight,proto3" json:"timeout_height,omitempty"`
	TiemoutTimestamp uint64     `protobuf:"varint,7,opt,name=tiemout_timestamp,json=tiemoutTimestamp,proto3" json:"tiemout_timestamp,omitempty"`
}

func (x *IBCTransferMetadata) Reset() {
	*x = IBCTransferMetadata{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sdk_utilities_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IBCTransferMetadata) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IBCTransferMetadata) ProtoMessage() {}

func (x *IBCTransferMetadata) ProtoReflect() protoreflect.Message {
	mi := &file_sdk_utilities_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IBCTransferMetadata.ProtoReflect.Descriptor instead.
func (*IBCTransferMetadata) Descriptor() ([]byte, []int) {
	return file_sdk_utilities_proto_rawDescGZIP(), []int{10}
}

func (x *IBCTransferMetadata) GetSourcePort() string {
	if x != nil {
		return x.SourcePort
	}
	return ""
}

func (x *IBCTransferMetadata) GetSourceCannel() string {
	if x != nil {
		return x.SourceCannel
	}
	return ""
}

func (x *IBCTransferMetadata) GetToken() *Coin {
	if x != nil {
		return x.Token
	}
	return nil
}

func (x *IBCTransferMetadata) GetSender() string {
	if x != nil {
		return x.Sender
	}
	return ""
}

func (x *IBCTransferMetadata) GetReceiver() string {
	if x != nil {
		return x.Receiver
	}
	return ""
}

func (x *IBCTransferMetadata) GetTimeoutHeight() *IBCHeight {
	if x != nil {
		return x.TimeoutHeight
	}
	return nil
}

func (x *IBCTransferMetadata) GetTiemoutTimestamp() uint64 {
	if x != nil {
		return x.TiemoutTimestamp
	}
	return 0
}

// The plain type associated with ibc-go types.Height struct
type IBCHeight struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RevisionNumber uint64 `protobuf:"varint,1,opt,name=revision_number,json=revisionNumber,proto3" json:"revision_number,omitempty"`
	RevisionHeight uint64 `protobuf:"varint,2,opt,name=revision_height,json=revisionHeight,proto3" json:"revision_height,omitempty"`
}

func (x *IBCHeight) Reset() {
	*x = IBCHeight{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sdk_utilities_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IBCHeight) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IBCHeight) ProtoMessage() {}

func (x *IBCHeight) ProtoReflect() protoreflect.Message {
	mi := &file_sdk_utilities_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IBCHeight.ProtoReflect.Descriptor instead.
func (*IBCHeight) Descriptor() ([]byte, []int) {
	return file_sdk_utilities_proto_rawDescGZIP(), []int{11}
}

func (x *IBCHeight) GetRevisionNumber() uint64 {
	if x != nil {
		return x.RevisionNumber
	}
	return 0
}

func (x *IBCHeight) GetRevisionHeight() uint64 {
	if x != nil {
		return x.RevisionHeight
	}
	return 0
}

var File_sdk_utilities_proto protoreflect.FileDescriptor

var file_sdk_utilities_proto_rawDesc = []byte{
	0x0a, 0x13, 0x73, 0x64, 0x6b, 0x5f, 0x75, 0x74, 0x69, 0x6c, 0x69, 0x74, 0x69, 0x65, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x73, 0x64, 0x6b, 0x5f, 0x75, 0x74, 0x69, 0x6c, 0x69,
	0x74, 0x69, 0x65, 0x73, 0x22, 0x42, 0x0a, 0x0d, 0x53, 0x75, 0x70, 0x70, 0x6c, 0x79, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x5f, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x68, 0x61, 0x69, 0x6e,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x11, 0x52, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x22, 0x3b, 0x0a, 0x0e, 0x53, 0x75, 0x70, 0x70,
	0x6c, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x29, 0x0a, 0x05, 0x63, 0x6f,
	0x69, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x73, 0x64, 0x6b, 0x5f,
	0x75, 0x74, 0x69, 0x6c, 0x69, 0x74, 0x69, 0x65, 0x73, 0x2e, 0x43, 0x6f, 0x69, 0x6e, 0x52, 0x05,
	0x63, 0x6f, 0x69, 0x6e, 0x73, 0x22, 0x34, 0x0a, 0x04, 0x43, 0x6f, 0x69, 0x6e, 0x12, 0x14, 0x0a,
	0x05, 0x64, 0x65, 0x6e, 0x6f, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x64, 0x65,
	0x6e, 0x6f, 0x6d, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x57, 0x0a, 0x0e, 0x51,
	0x75, 0x65, 0x72, 0x79, 0x54, 0x78, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a,
	0x0a, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04,
	0x70, 0x6f, 0x72, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x11, 0x52, 0x04, 0x70, 0x6f, 0x72, 0x74,
	0x12, 0x12, 0x0a, 0x04, 0x68, 0x61, 0x73, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x68, 0x61, 0x73, 0x68, 0x22, 0x27, 0x0a, 0x0f, 0x51, 0x75, 0x65, 0x72, 0x79, 0x54, 0x78, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x66, 0x69, 0x65, 0x6c, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x22, 0x62, 0x0a,
	0x12, 0x42, 0x72, 0x6f, 0x61, 0x64, 0x63, 0x61, 0x73, 0x74, 0x54, 0x78, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x5f, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x11,
	0x52, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x74, 0x78, 0x5f, 0x62, 0x79, 0x74,
	0x65, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x74, 0x78, 0x42, 0x79, 0x74, 0x65,
	0x73, 0x22, 0x29, 0x0a, 0x13, 0x42, 0x72, 0x6f, 0x61, 0x64, 0x63, 0x61, 0x73, 0x74, 0x54, 0x78,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x68, 0x61, 0x73, 0x68,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x68, 0x61, 0x73, 0x68, 0x22, 0x2e, 0x0a, 0x11,
	0x54, 0x78, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x19, 0x0a, 0x08, 0x74, 0x78, 0x5f, 0x62, 0x79, 0x74, 0x65, 0x73, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0c, 0x52, 0x07, 0x74, 0x78, 0x42, 0x79, 0x74, 0x65, 0x73, 0x22, 0x5c, 0x0a, 0x12,
	0x54, 0x78, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x46, 0x0a, 0x11, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x5f, 0x6d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e,
	0x73, 0x64, 0x6b, 0x5f, 0x75, 0x74, 0x69, 0x6c, 0x69, 0x74, 0x69, 0x65, 0x73, 0x2e, 0x54, 0x78,
	0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x52, 0x10, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x73, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x22, 0x7d, 0x0a, 0x0a, 0x54, 0x78,
	0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x17, 0x0a, 0x07, 0x74, 0x78, 0x5f, 0x74,
	0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x74, 0x78, 0x54, 0x79, 0x70,
	0x65, 0x12, 0x56, 0x0a, 0x15, 0x69, 0x62, 0x63, 0x5f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65,
	0x72, 0x5f, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x22, 0x2e, 0x73, 0x64, 0x6b, 0x5f, 0x75, 0x74, 0x69, 0x6c, 0x69, 0x74, 0x69, 0x65, 0x73,
	0x2e, 0x49, 0x42, 0x43, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x4d, 0x65, 0x74, 0x61,
	0x64, 0x61, 0x74, 0x61, 0x52, 0x13, 0x69, 0x62, 0x63, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65,
	0x72, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x22, 0xa8, 0x02, 0x0a, 0x13, 0x49, 0x42,
	0x43, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x70, 0x6f, 0x72, 0x74,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x50, 0x6f,
	0x72, 0x74, 0x12, 0x23, 0x0a, 0x0d, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x63, 0x61, 0x6e,
	0x6e, 0x65, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x43, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x12, 0x29, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x73, 0x64, 0x6b, 0x5f, 0x75, 0x74, 0x69,
	0x6c, 0x69, 0x74, 0x69, 0x65, 0x73, 0x2e, 0x43, 0x6f, 0x69, 0x6e, 0x52, 0x05, 0x74, 0x6f, 0x6b,
	0x65, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65,
	0x63, 0x65, 0x69, 0x76, 0x65, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x65,
	0x63, 0x65, 0x69, 0x76, 0x65, 0x72, 0x12, 0x3f, 0x0a, 0x0e, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75,
	0x74, 0x5f, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18,
	0x2e, 0x73, 0x64, 0x6b, 0x5f, 0x75, 0x74, 0x69, 0x6c, 0x69, 0x74, 0x69, 0x65, 0x73, 0x2e, 0x49,
	0x42, 0x43, 0x48, 0x65, 0x69, 0x67, 0x68, 0x74, 0x52, 0x0d, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75,
	0x74, 0x48, 0x65, 0x69, 0x67, 0x68, 0x74, 0x12, 0x2b, 0x0a, 0x11, 0x74, 0x69, 0x65, 0x6d, 0x6f,
	0x75, 0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x10, 0x74, 0x69, 0x65, 0x6d, 0x6f, 0x75, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x22, 0x5d, 0x0a, 0x09, 0x49, 0x42, 0x43, 0x48, 0x65, 0x69, 0x67, 0x68,
	0x74, 0x12, 0x27, 0x0a, 0x0f, 0x72, 0x65, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x6e, 0x75,
	0x6d, 0x62, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0e, 0x72, 0x65, 0x76, 0x69,
	0x73, 0x69, 0x6f, 0x6e, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x27, 0x0a, 0x0f, 0x72, 0x65,
	0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x0e, 0x72, 0x65, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x48, 0x65, 0x69,
	0x67, 0x68, 0x74, 0x32, 0xd0, 0x02, 0x0a, 0x0c, 0x53, 0x64, 0x6b, 0x55, 0x74, 0x69, 0x6c, 0x69,
	0x74, 0x69, 0x65, 0x73, 0x12, 0x45, 0x0a, 0x06, 0x53, 0x75, 0x70, 0x70, 0x6c, 0x79, 0x12, 0x1c,
	0x2e, 0x73, 0x64, 0x6b, 0x5f, 0x75, 0x74, 0x69, 0x6c, 0x69, 0x74, 0x69, 0x65, 0x73, 0x2e, 0x53,
	0x75, 0x70, 0x70, 0x6c, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x73,
	0x64, 0x6b, 0x5f, 0x75, 0x74, 0x69, 0x6c, 0x69, 0x74, 0x69, 0x65, 0x73, 0x2e, 0x53, 0x75, 0x70,
	0x70, 0x6c, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x48, 0x0a, 0x07, 0x51,
	0x75, 0x65, 0x72, 0x79, 0x54, 0x78, 0x12, 0x1d, 0x2e, 0x73, 0x64, 0x6b, 0x5f, 0x75, 0x74, 0x69,
	0x6c, 0x69, 0x74, 0x69, 0x65, 0x73, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x54, 0x78, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x73, 0x64, 0x6b, 0x5f, 0x75, 0x74, 0x69, 0x6c,
	0x69, 0x74, 0x69, 0x65, 0x73, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x54, 0x78, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x54, 0x0a, 0x0b, 0x42, 0x72, 0x6f, 0x61, 0x64, 0x63, 0x61,
	0x73, 0x74, 0x54, 0x78, 0x12, 0x21, 0x2e, 0x73, 0x64, 0x6b, 0x5f, 0x75, 0x74, 0x69, 0x6c, 0x69,
	0x74, 0x69, 0x65, 0x73, 0x2e, 0x42, 0x72, 0x6f, 0x61, 0x64, 0x63, 0x61, 0x73, 0x74, 0x54, 0x78,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x73, 0x64, 0x6b, 0x5f, 0x75, 0x74,
	0x69, 0x6c, 0x69, 0x74, 0x69, 0x65, 0x73, 0x2e, 0x42, 0x72, 0x6f, 0x61, 0x64, 0x63, 0x61, 0x73,
	0x74, 0x54, 0x78, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x59, 0x0a, 0x12, 0x54,
	0x78, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e,
	0x74, 0x12, 0x20, 0x2e, 0x73, 0x64, 0x6b, 0x5f, 0x75, 0x74, 0x69, 0x6c, 0x69, 0x74, 0x69, 0x65,
	0x73, 0x2e, 0x54, 0x78, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x73, 0x64, 0x6b, 0x5f, 0x75, 0x74, 0x69, 0x6c, 0x69, 0x74,
	0x69, 0x65, 0x73, 0x2e, 0x54, 0x78, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x12, 0x5a, 0x10, 0x2f, 0x73, 0x64, 0x6b, 0x5f, 0x75,
	0x74, 0x69, 0x6c, 0x69, 0x74, 0x69, 0x65, 0x73, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_sdk_utilities_proto_rawDescOnce sync.Once
	file_sdk_utilities_proto_rawDescData = file_sdk_utilities_proto_rawDesc
)

func file_sdk_utilities_proto_rawDescGZIP() []byte {
	file_sdk_utilities_proto_rawDescOnce.Do(func() {
		file_sdk_utilities_proto_rawDescData = protoimpl.X.CompressGZIP(file_sdk_utilities_proto_rawDescData)
	})
	return file_sdk_utilities_proto_rawDescData
}

var file_sdk_utilities_proto_msgTypes = make([]protoimpl.MessageInfo, 12)
var file_sdk_utilities_proto_goTypes = []interface{}{
	(*SupplyRequest)(nil),       // 0: sdk_utilities.SupplyRequest
	(*SupplyResponse)(nil),      // 1: sdk_utilities.SupplyResponse
	(*Coin)(nil),                // 2: sdk_utilities.Coin
	(*QueryTxRequest)(nil),      // 3: sdk_utilities.QueryTxRequest
	(*QueryTxResponse)(nil),     // 4: sdk_utilities.QueryTxResponse
	(*BroadcastTxRequest)(nil),  // 5: sdk_utilities.BroadcastTxRequest
	(*BroadcastTxResponse)(nil), // 6: sdk_utilities.BroadcastTxResponse
	(*TxMetadataRequest)(nil),   // 7: sdk_utilities.TxMetadataRequest
	(*TxMetadataResponse)(nil),  // 8: sdk_utilities.TxMetadataResponse
	(*TxMetadata)(nil),          // 9: sdk_utilities.TxMetadata
	(*IBCTransferMetadata)(nil), // 10: sdk_utilities.IBCTransferMetadata
	(*IBCHeight)(nil),           // 11: sdk_utilities.IBCHeight
}
var file_sdk_utilities_proto_depIdxs = []int32{
	2,  // 0: sdk_utilities.SupplyResponse.coins:type_name -> sdk_utilities.Coin
	9,  // 1: sdk_utilities.TxMetadataResponse.messages_metadata:type_name -> sdk_utilities.TxMetadata
	10, // 2: sdk_utilities.TxMetadata.ibc_transfer_metadata:type_name -> sdk_utilities.IBCTransferMetadata
	2,  // 3: sdk_utilities.IBCTransferMetadata.token:type_name -> sdk_utilities.Coin
	11, // 4: sdk_utilities.IBCTransferMetadata.timeout_height:type_name -> sdk_utilities.IBCHeight
	0,  // 5: sdk_utilities.SdkUtilities.Supply:input_type -> sdk_utilities.SupplyRequest
	3,  // 6: sdk_utilities.SdkUtilities.QueryTx:input_type -> sdk_utilities.QueryTxRequest
	5,  // 7: sdk_utilities.SdkUtilities.BroadcastTx:input_type -> sdk_utilities.BroadcastTxRequest
	7,  // 8: sdk_utilities.SdkUtilities.TxMetadataEndpoint:input_type -> sdk_utilities.TxMetadataRequest
	1,  // 9: sdk_utilities.SdkUtilities.Supply:output_type -> sdk_utilities.SupplyResponse
	4,  // 10: sdk_utilities.SdkUtilities.QueryTx:output_type -> sdk_utilities.QueryTxResponse
	6,  // 11: sdk_utilities.SdkUtilities.BroadcastTx:output_type -> sdk_utilities.BroadcastTxResponse
	8,  // 12: sdk_utilities.SdkUtilities.TxMetadataEndpoint:output_type -> sdk_utilities.TxMetadataResponse
	9,  // [9:13] is the sub-list for method output_type
	5,  // [5:9] is the sub-list for method input_type
	5,  // [5:5] is the sub-list for extension type_name
	5,  // [5:5] is the sub-list for extension extendee
	0,  // [0:5] is the sub-list for field type_name
}

func init() { file_sdk_utilities_proto_init() }
func file_sdk_utilities_proto_init() {
	if File_sdk_utilities_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_sdk_utilities_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SupplyRequest); i {
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
		file_sdk_utilities_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SupplyResponse); i {
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
		file_sdk_utilities_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Coin); i {
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
		file_sdk_utilities_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryTxRequest); i {
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
		file_sdk_utilities_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryTxResponse); i {
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
		file_sdk_utilities_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BroadcastTxRequest); i {
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
		file_sdk_utilities_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BroadcastTxResponse); i {
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
		file_sdk_utilities_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TxMetadataRequest); i {
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
		file_sdk_utilities_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TxMetadataResponse); i {
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
		file_sdk_utilities_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TxMetadata); i {
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
		file_sdk_utilities_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IBCTransferMetadata); i {
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
		file_sdk_utilities_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IBCHeight); i {
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
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_sdk_utilities_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   12,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_sdk_utilities_proto_goTypes,
		DependencyIndexes: file_sdk_utilities_proto_depIdxs,
		MessageInfos:      file_sdk_utilities_proto_msgTypes,
	}.Build()
	File_sdk_utilities_proto = out.File
	file_sdk_utilities_proto_rawDesc = nil
	file_sdk_utilities_proto_goTypes = nil
	file_sdk_utilities_proto_depIdxs = nil
}
