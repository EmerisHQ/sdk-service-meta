// Code generated by goa v3.5.2, DO NOT EDIT.
//
// sdk-utilities gRPC server encoders and decoders
//
// Command:
// $ goa gen github.com/allinbits/sdk-service-meta

package server

import (
	"context"

	sdk_utilitiespb "github.com/allinbits/sdk-service-meta/gen/grpc/sdk_utilities/pb"
	sdkutilities "github.com/allinbits/sdk-service-meta/gen/sdk_utilities"
	goagrpc "goa.design/goa/v3/grpc"
	"google.golang.org/grpc/metadata"
)

// EncodeSupplyResponse encodes responses from the "sdk-utilities" service
// "supply" endpoint.
func EncodeSupplyResponse(ctx context.Context, v interface{}, hdr, trlr *metadata.MD) (interface{}, error) {
	result, ok := v.(*sdkutilities.Supply2)
	if !ok {
		return nil, goagrpc.ErrInvalidType("sdk-utilities", "supply", "*sdkutilities.Supply2", v)
	}
	resp := NewSupplyResponse(result)
	return resp, nil
}

// DecodeSupplyRequest decodes requests sent to "sdk-utilities" service
// "supply" endpoint.
func DecodeSupplyRequest(ctx context.Context, v interface{}, md metadata.MD) (interface{}, error) {
	var (
		message *sdk_utilitiespb.SupplyRequest
		ok      bool
	)
	{
		if message, ok = v.(*sdk_utilitiespb.SupplyRequest); !ok {
			return nil, goagrpc.ErrInvalidType("sdk-utilities", "supply", "*sdk_utilitiespb.SupplyRequest", v)
		}
	}
	var payload *sdkutilities.SupplyPayload
	{
		payload = NewSupplyPayload(message)
	}
	return payload, nil
}

// EncodeQueryTxResponse encodes responses from the "sdk-utilities" service
// "queryTx" endpoint.
func EncodeQueryTxResponse(ctx context.Context, v interface{}, hdr, trlr *metadata.MD) (interface{}, error) {
	result, ok := v.([]byte)
	if !ok {
		return nil, goagrpc.ErrInvalidType("sdk-utilities", "queryTx", "[]byte", v)
	}
	resp := NewQueryTxResponse(result)
	return resp, nil
}

// DecodeQueryTxRequest decodes requests sent to "sdk-utilities" service
// "queryTx" endpoint.
func DecodeQueryTxRequest(ctx context.Context, v interface{}, md metadata.MD) (interface{}, error) {
	var (
		message *sdk_utilitiespb.QueryTxRequest
		ok      bool
	)
	{
		if message, ok = v.(*sdk_utilitiespb.QueryTxRequest); !ok {
			return nil, goagrpc.ErrInvalidType("sdk-utilities", "queryTx", "*sdk_utilitiespb.QueryTxRequest", v)
		}
	}
	var payload *sdkutilities.QueryTxPayload
	{
		payload = NewQueryTxPayload(message)
	}
	return payload, nil
}

// EncodeBroadcastTxResponse encodes responses from the "sdk-utilities" service
// "broadcastTx" endpoint.
func EncodeBroadcastTxResponse(ctx context.Context, v interface{}, hdr, trlr *metadata.MD) (interface{}, error) {
	result, ok := v.(*sdkutilities.TransactionResult)
	if !ok {
		return nil, goagrpc.ErrInvalidType("sdk-utilities", "broadcastTx", "*sdkutilities.TransactionResult", v)
	}
	resp := NewBroadcastTxResponse(result)
	return resp, nil
}

// DecodeBroadcastTxRequest decodes requests sent to "sdk-utilities" service
// "broadcastTx" endpoint.
func DecodeBroadcastTxRequest(ctx context.Context, v interface{}, md metadata.MD) (interface{}, error) {
	var (
		message *sdk_utilitiespb.BroadcastTxRequest
		ok      bool
	)
	{
		if message, ok = v.(*sdk_utilitiespb.BroadcastTxRequest); !ok {
			return nil, goagrpc.ErrInvalidType("sdk-utilities", "broadcastTx", "*sdk_utilitiespb.BroadcastTxRequest", v)
		}
		if err := ValidateBroadcastTxRequest(message); err != nil {
			return nil, err
		}
	}
	var payload *sdkutilities.BroadcastTxPayload
	{
		payload = NewBroadcastTxPayload(message)
	}
	return payload, nil
}

// EncodeTxMetadataResponse encodes responses from the "sdk-utilities" service
// "txMetadata" endpoint.
func EncodeTxMetadataResponse(ctx context.Context, v interface{}, hdr, trlr *metadata.MD) (interface{}, error) {
	result, ok := v.(*sdkutilities.TxMetadata2)
	if !ok {
		return nil, goagrpc.ErrInvalidType("sdk-utilities", "txMetadata", "*sdkutilities.TxMetadata2", v)
	}
	resp := NewTxMetadataResponse(result)
	return resp, nil
}

// DecodeTxMetadataRequest decodes requests sent to "sdk-utilities" service
// "txMetadata" endpoint.
func DecodeTxMetadataRequest(ctx context.Context, v interface{}, md metadata.MD) (interface{}, error) {
	var (
		message *sdk_utilitiespb.TxMetadataRequest
		ok      bool
	)
	{
		if message, ok = v.(*sdk_utilitiespb.TxMetadataRequest); !ok {
			return nil, goagrpc.ErrInvalidType("sdk-utilities", "txMetadata", "*sdk_utilitiespb.TxMetadataRequest", v)
		}
		if err := ValidateTxMetadataRequest(message); err != nil {
			return nil, err
		}
	}
	var payload *sdkutilities.TxMetadataPayload
	{
		payload = NewTxMetadataPayload(message)
	}
	return payload, nil
}
