// Code generated by goa v3.5.2, DO NOT EDIT.
//
// sdk-utilities service
//
// Command:
// $ goa gen github.com/allinbits/sdk-service-meta

package sdkutilities

import (
	"context"
)

// sdk-utilities performs Cosmos SDK-related operations
type Service interface {
	// Supply implements supply.
	Supply(context.Context, *SupplyPayload) (res *Supply2, err error)
	// QueryTx implements queryTx.
	QueryTx(context.Context, *QueryTxPayload) (res []byte, err error)
	// BroadcastTx implements broadcastTx.
	BroadcastTx(context.Context, *BroadcastTxPayload) (res *TransactionResult, err error)
	// TxMetadata implements txMetadata.
	TxMetadata(context.Context, *TxMetadataPayload) (res *TxMetadata2, err error)
}

// ServiceName is the name of the service as defined in the design. This is the
// same value that is set in the endpoint request contexts under the ServiceKey
// key.
const ServiceName = "sdk-utilities"

// MethodNames lists the service method names as defined in the design. These
// are the same values that are set in the endpoint request contexts under the
// MethodKey key.
var MethodNames = [4]string{"supply", "queryTx", "broadcastTx", "txMetadata"}

// SupplyPayload is the payload type of the sdk-utilities service supply method.
type SupplyPayload struct {
	// Chain to get data from
	ChainName string
	// gRPC port for selected chain, defaults to 9090
	Port *int
}

// Supply2 is the result type of the sdk-utilities service supply method.
type Supply2 struct {
	Coins []*Coin
}

// QueryTxPayload is the payload type of the sdk-utilities service queryTx
// method.
type QueryTxPayload struct {
	// Chain to get data from
	ChainName string
	// gRPC port for selected chain, defaults to 9090
	Port *int
	// Transaction hash to query
	Hash string
}

// BroadcastTxPayload is the payload type of the sdk-utilities service
// broadcastTx method.
type BroadcastTxPayload struct {
	// Chain to get data from
	ChainName string
	// gRPC port for selected chain, defaults to 9090
	Port *int
	// Transaction bytes
	TxBytes []byte
}

// TransactionResult is the result type of the sdk-utilities service
// broadcastTx method.
type TransactionResult struct {
	Hash string
}

// TxMetadataPayload is the payload type of the sdk-utilities service
// txMetadata method.
type TxMetadataPayload struct {
	// Transaction bytes
	TxBytes []byte
}

// TxMetadata2 is the result type of the sdk-utilities service txMetadata
// method.
type TxMetadata2 struct {
	TxType              string
	IbcTransferMetadata *IBCTransferMetadata
}

// SDK service representation of a Cosmos SDK types.Coin
type Coin struct {
	Denom  string
	Amount string
}

// Metadata related to an IBC Transfer
type IBCTransferMetadata struct {
	SourcePort       *string
	SourceCannel     *string
	Token            *Coin
	Sender           *string
	Receiver         *string
	TimeoutHeight    *IBCHeight
	TiemoutTimestamp *uint64
}

// The plain type associated with ibc-go types.Height struct
type IBCHeight struct {
	RevisionNumber *uint64
	RevisionHeight *uint64
}
