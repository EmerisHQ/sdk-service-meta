// Code generated by goa v3.7.2, DO NOT EDIT.
//
// sdk-utilities gRPC server encoders and decoders
//
// Command:
// $ goa gen github.com/emerishq/sdk-service-meta

package server

import (
	"context"

	sdk_utilitiespb "github.com/emerishq/sdk-service-meta/gen/grpc/sdk_utilities/pb"
	sdkutilities "github.com/emerishq/sdk-service-meta/gen/sdk_utilities"
	goagrpc "goa.design/goa/v3/grpc"
	"google.golang.org/grpc/metadata"
)

// EncodeAccountNumbersResponse encodes responses from the "sdk-utilities"
// service "accountNumbers" endpoint.
func EncodeAccountNumbersResponse(ctx context.Context, v interface{}, hdr, trlr *metadata.MD) (interface{}, error) {
	result, ok := v.(*sdkutilities.AccountNumbers2)
	if !ok {
		return nil, goagrpc.ErrInvalidType("sdk-utilities", "accountNumbers", "*sdkutilities.AccountNumbers2", v)
	}
	resp := NewProtoAccountNumbersResponse(result)
	return resp, nil
}

// DecodeAccountNumbersRequest decodes requests sent to "sdk-utilities" service
// "accountNumbers" endpoint.
func DecodeAccountNumbersRequest(ctx context.Context, v interface{}, md metadata.MD) (interface{}, error) {
	var (
		message *sdk_utilitiespb.AccountNumbersRequest
		ok      bool
	)
	{
		if message, ok = v.(*sdk_utilitiespb.AccountNumbersRequest); !ok {
			return nil, goagrpc.ErrInvalidType("sdk-utilities", "accountNumbers", "*sdk_utilitiespb.AccountNumbersRequest", v)
		}
	}
	var payload *sdkutilities.AccountNumbersPayload
	{
		payload = NewAccountNumbersPayload(message)
	}
	return payload, nil
}

// EncodeSupplyResponse encodes responses from the "sdk-utilities" service
// "supply" endpoint.
func EncodeSupplyResponse(ctx context.Context, v interface{}, hdr, trlr *metadata.MD) (interface{}, error) {
	result, ok := v.(*sdkutilities.Supply2)
	if !ok {
		return nil, goagrpc.ErrInvalidType("sdk-utilities", "supply", "*sdkutilities.Supply2", v)
	}
	resp := NewProtoSupplyResponse(result)
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

// EncodeSupplyDenomResponse encodes responses from the "sdk-utilities" service
// "supplyDenom" endpoint.
func EncodeSupplyDenomResponse(ctx context.Context, v interface{}, hdr, trlr *metadata.MD) (interface{}, error) {
	result, ok := v.(*sdkutilities.Supply2)
	if !ok {
		return nil, goagrpc.ErrInvalidType("sdk-utilities", "supplyDenom", "*sdkutilities.Supply2", v)
	}
	resp := NewProtoSupplyDenomResponse(result)
	return resp, nil
}

// DecodeSupplyDenomRequest decodes requests sent to "sdk-utilities" service
// "supplyDenom" endpoint.
func DecodeSupplyDenomRequest(ctx context.Context, v interface{}, md metadata.MD) (interface{}, error) {
	var (
		message *sdk_utilitiespb.SupplyDenomRequest
		ok      bool
	)
	{
		if message, ok = v.(*sdk_utilitiespb.SupplyDenomRequest); !ok {
			return nil, goagrpc.ErrInvalidType("sdk-utilities", "supplyDenom", "*sdk_utilitiespb.SupplyDenomRequest", v)
		}
	}
	var payload *sdkutilities.SupplyDenomPayload
	{
		payload = NewSupplyDenomPayload(message)
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
	resp := NewProtoQueryTxResponse(result)
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
	resp := NewProtoBroadcastTxResponse(result)
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
	result, ok := v.(*sdkutilities.TxMessagesMetadata)
	if !ok {
		return nil, goagrpc.ErrInvalidType("sdk-utilities", "txMetadata", "*sdkutilities.TxMessagesMetadata", v)
	}
	resp := NewProtoTxMetadataResponse(result)
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

// EncodeBlockResponse encodes responses from the "sdk-utilities" service
// "block" endpoint.
func EncodeBlockResponse(ctx context.Context, v interface{}, hdr, trlr *metadata.MD) (interface{}, error) {
	result, ok := v.(*sdkutilities.BlockData)
	if !ok {
		return nil, goagrpc.ErrInvalidType("sdk-utilities", "block", "*sdkutilities.BlockData", v)
	}
	resp := NewProtoBlockResponse(result)
	return resp, nil
}

// DecodeBlockRequest decodes requests sent to "sdk-utilities" service "block"
// endpoint.
func DecodeBlockRequest(ctx context.Context, v interface{}, md metadata.MD) (interface{}, error) {
	var (
		message *sdk_utilitiespb.BlockRequest
		ok      bool
	)
	{
		if message, ok = v.(*sdk_utilitiespb.BlockRequest); !ok {
			return nil, goagrpc.ErrInvalidType("sdk-utilities", "block", "*sdk_utilitiespb.BlockRequest", v)
		}
	}
	var payload *sdkutilities.BlockPayload
	{
		payload = NewBlockPayload(message)
	}
	return payload, nil
}

// EncodeLiquidityParamsResponse encodes responses from the "sdk-utilities"
// service "liquidityParams" endpoint.
func EncodeLiquidityParamsResponse(ctx context.Context, v interface{}, hdr, trlr *metadata.MD) (interface{}, error) {
	result, ok := v.(*sdkutilities.LiquidityParams2)
	if !ok {
		return nil, goagrpc.ErrInvalidType("sdk-utilities", "liquidityParams", "*sdkutilities.LiquidityParams2", v)
	}
	resp := NewProtoLiquidityParamsResponse(result)
	return resp, nil
}

// DecodeLiquidityParamsRequest decodes requests sent to "sdk-utilities"
// service "liquidityParams" endpoint.
func DecodeLiquidityParamsRequest(ctx context.Context, v interface{}, md metadata.MD) (interface{}, error) {
	var (
		message *sdk_utilitiespb.LiquidityParamsRequest
		ok      bool
	)
	{
		if message, ok = v.(*sdk_utilitiespb.LiquidityParamsRequest); !ok {
			return nil, goagrpc.ErrInvalidType("sdk-utilities", "liquidityParams", "*sdk_utilitiespb.LiquidityParamsRequest", v)
		}
	}
	var payload *sdkutilities.LiquidityParamsPayload
	{
		payload = NewLiquidityParamsPayload(message)
	}
	return payload, nil
}

// EncodeLiquidityPoolsResponse encodes responses from the "sdk-utilities"
// service "liquidityPools" endpoint.
func EncodeLiquidityPoolsResponse(ctx context.Context, v interface{}, hdr, trlr *metadata.MD) (interface{}, error) {
	result, ok := v.(*sdkutilities.LiquidityPools2)
	if !ok {
		return nil, goagrpc.ErrInvalidType("sdk-utilities", "liquidityPools", "*sdkutilities.LiquidityPools2", v)
	}
	resp := NewProtoLiquidityPoolsResponse(result)
	return resp, nil
}

// DecodeLiquidityPoolsRequest decodes requests sent to "sdk-utilities" service
// "liquidityPools" endpoint.
func DecodeLiquidityPoolsRequest(ctx context.Context, v interface{}, md metadata.MD) (interface{}, error) {
	var (
		message *sdk_utilitiespb.LiquidityPoolsRequest
		ok      bool
	)
	{
		if message, ok = v.(*sdk_utilitiespb.LiquidityPoolsRequest); !ok {
			return nil, goagrpc.ErrInvalidType("sdk-utilities", "liquidityPools", "*sdk_utilitiespb.LiquidityPoolsRequest", v)
		}
	}
	var payload *sdkutilities.LiquidityPoolsPayload
	{
		payload = NewLiquidityPoolsPayload(message)
	}
	return payload, nil
}

// EncodeMintInflationResponse encodes responses from the "sdk-utilities"
// service "mintInflation" endpoint.
func EncodeMintInflationResponse(ctx context.Context, v interface{}, hdr, trlr *metadata.MD) (interface{}, error) {
	result, ok := v.(*sdkutilities.MintInflation2)
	if !ok {
		return nil, goagrpc.ErrInvalidType("sdk-utilities", "mintInflation", "*sdkutilities.MintInflation2", v)
	}
	resp := NewProtoMintInflationResponse(result)
	return resp, nil
}

// DecodeMintInflationRequest decodes requests sent to "sdk-utilities" service
// "mintInflation" endpoint.
func DecodeMintInflationRequest(ctx context.Context, v interface{}, md metadata.MD) (interface{}, error) {
	var (
		message *sdk_utilitiespb.MintInflationRequest
		ok      bool
	)
	{
		if message, ok = v.(*sdk_utilitiespb.MintInflationRequest); !ok {
			return nil, goagrpc.ErrInvalidType("sdk-utilities", "mintInflation", "*sdk_utilitiespb.MintInflationRequest", v)
		}
	}
	var payload *sdkutilities.MintInflationPayload
	{
		payload = NewMintInflationPayload(message)
	}
	return payload, nil
}

// EncodeMintParamsResponse encodes responses from the "sdk-utilities" service
// "mintParams" endpoint.
func EncodeMintParamsResponse(ctx context.Context, v interface{}, hdr, trlr *metadata.MD) (interface{}, error) {
	result, ok := v.(*sdkutilities.MintParams2)
	if !ok {
		return nil, goagrpc.ErrInvalidType("sdk-utilities", "mintParams", "*sdkutilities.MintParams2", v)
	}
	resp := NewProtoMintParamsResponse(result)
	return resp, nil
}

// DecodeMintParamsRequest decodes requests sent to "sdk-utilities" service
// "mintParams" endpoint.
func DecodeMintParamsRequest(ctx context.Context, v interface{}, md metadata.MD) (interface{}, error) {
	var (
		message *sdk_utilitiespb.MintParamsRequest
		ok      bool
	)
	{
		if message, ok = v.(*sdk_utilitiespb.MintParamsRequest); !ok {
			return nil, goagrpc.ErrInvalidType("sdk-utilities", "mintParams", "*sdk_utilitiespb.MintParamsRequest", v)
		}
	}
	var payload *sdkutilities.MintParamsPayload
	{
		payload = NewMintParamsPayload(message)
	}
	return payload, nil
}

// EncodeMintAnnualProvisionResponse encodes responses from the "sdk-utilities"
// service "mintAnnualProvision" endpoint.
func EncodeMintAnnualProvisionResponse(ctx context.Context, v interface{}, hdr, trlr *metadata.MD) (interface{}, error) {
	result, ok := v.(*sdkutilities.MintAnnualProvision2)
	if !ok {
		return nil, goagrpc.ErrInvalidType("sdk-utilities", "mintAnnualProvision", "*sdkutilities.MintAnnualProvision2", v)
	}
	resp := NewProtoMintAnnualProvisionResponse(result)
	return resp, nil
}

// DecodeMintAnnualProvisionRequest decodes requests sent to "sdk-utilities"
// service "mintAnnualProvision" endpoint.
func DecodeMintAnnualProvisionRequest(ctx context.Context, v interface{}, md metadata.MD) (interface{}, error) {
	var (
		message *sdk_utilitiespb.MintAnnualProvisionRequest
		ok      bool
	)
	{
		if message, ok = v.(*sdk_utilitiespb.MintAnnualProvisionRequest); !ok {
			return nil, goagrpc.ErrInvalidType("sdk-utilities", "mintAnnualProvision", "*sdk_utilitiespb.MintAnnualProvisionRequest", v)
		}
	}
	var payload *sdkutilities.MintAnnualProvisionPayload
	{
		payload = NewMintAnnualProvisionPayload(message)
	}
	return payload, nil
}

// EncodeMintEpochProvisionsResponse encodes responses from the "sdk-utilities"
// service "mintEpochProvisions" endpoint.
func EncodeMintEpochProvisionsResponse(ctx context.Context, v interface{}, hdr, trlr *metadata.MD) (interface{}, error) {
	result, ok := v.(*sdkutilities.MintEpochProvisions2)
	if !ok {
		return nil, goagrpc.ErrInvalidType("sdk-utilities", "mintEpochProvisions", "*sdkutilities.MintEpochProvisions2", v)
	}
	resp := NewProtoMintEpochProvisionsResponse(result)
	return resp, nil
}

// DecodeMintEpochProvisionsRequest decodes requests sent to "sdk-utilities"
// service "mintEpochProvisions" endpoint.
func DecodeMintEpochProvisionsRequest(ctx context.Context, v interface{}, md metadata.MD) (interface{}, error) {
	var (
		message *sdk_utilitiespb.MintEpochProvisionsRequest
		ok      bool
	)
	{
		if message, ok = v.(*sdk_utilitiespb.MintEpochProvisionsRequest); !ok {
			return nil, goagrpc.ErrInvalidType("sdk-utilities", "mintEpochProvisions", "*sdk_utilitiespb.MintEpochProvisionsRequest", v)
		}
	}
	var payload *sdkutilities.MintEpochProvisionsPayload
	{
		payload = NewMintEpochProvisionsPayload(message)
	}
	return payload, nil
}

// EncodeDelegatorRewardsResponse encodes responses from the "sdk-utilities"
// service "delegatorRewards" endpoint.
func EncodeDelegatorRewardsResponse(ctx context.Context, v interface{}, hdr, trlr *metadata.MD) (interface{}, error) {
	result, ok := v.(*sdkutilities.DelegatorRewards2)
	if !ok {
		return nil, goagrpc.ErrInvalidType("sdk-utilities", "delegatorRewards", "*sdkutilities.DelegatorRewards2", v)
	}
	resp := NewProtoDelegatorRewardsResponse(result)
	return resp, nil
}

// DecodeDelegatorRewardsRequest decodes requests sent to "sdk-utilities"
// service "delegatorRewards" endpoint.
func DecodeDelegatorRewardsRequest(ctx context.Context, v interface{}, md metadata.MD) (interface{}, error) {
	var (
		message *sdk_utilitiespb.DelegatorRewardsRequest
		ok      bool
	)
	{
		if message, ok = v.(*sdk_utilitiespb.DelegatorRewardsRequest); !ok {
			return nil, goagrpc.ErrInvalidType("sdk-utilities", "delegatorRewards", "*sdk_utilitiespb.DelegatorRewardsRequest", v)
		}
	}
	var payload *sdkutilities.DelegatorRewardsPayload
	{
		payload = NewDelegatorRewardsPayload(message)
	}
	return payload, nil
}

// EncodeEstimateFeesResponse encodes responses from the "sdk-utilities"
// service "estimateFees" endpoint.
func EncodeEstimateFeesResponse(ctx context.Context, v interface{}, hdr, trlr *metadata.MD) (interface{}, error) {
	result, ok := v.(*sdkutilities.Simulation)
	if !ok {
		return nil, goagrpc.ErrInvalidType("sdk-utilities", "estimateFees", "*sdkutilities.Simulation", v)
	}
	resp := NewProtoEstimateFeesResponse(result)
	return resp, nil
}

// DecodeEstimateFeesRequest decodes requests sent to "sdk-utilities" service
// "estimateFees" endpoint.
func DecodeEstimateFeesRequest(ctx context.Context, v interface{}, md metadata.MD) (interface{}, error) {
	var (
		message *sdk_utilitiespb.EstimateFeesRequest
		ok      bool
	)
	{
		if message, ok = v.(*sdk_utilitiespb.EstimateFeesRequest); !ok {
			return nil, goagrpc.ErrInvalidType("sdk-utilities", "estimateFees", "*sdk_utilitiespb.EstimateFeesRequest", v)
		}
		if err := ValidateEstimateFeesRequest(message); err != nil {
			return nil, err
		}
	}
	var payload *sdkutilities.EstimateFeesPayload
	{
		payload = NewEstimateFeesPayload(message)
	}
	return payload, nil
}

// EncodeStakingParamsResponse encodes responses from the "sdk-utilities"
// service "stakingParams" endpoint.
func EncodeStakingParamsResponse(ctx context.Context, v interface{}, hdr, trlr *metadata.MD) (interface{}, error) {
	result, ok := v.(*sdkutilities.StakingParams2)
	if !ok {
		return nil, goagrpc.ErrInvalidType("sdk-utilities", "stakingParams", "*sdkutilities.StakingParams2", v)
	}
	resp := NewProtoStakingParamsResponse(result)
	return resp, nil
}

// DecodeStakingParamsRequest decodes requests sent to "sdk-utilities" service
// "stakingParams" endpoint.
func DecodeStakingParamsRequest(ctx context.Context, v interface{}, md metadata.MD) (interface{}, error) {
	var (
		message *sdk_utilitiespb.StakingParamsRequest
		ok      bool
	)
	{
		if message, ok = v.(*sdk_utilitiespb.StakingParamsRequest); !ok {
			return nil, goagrpc.ErrInvalidType("sdk-utilities", "stakingParams", "*sdk_utilitiespb.StakingParamsRequest", v)
		}
	}
	var payload *sdkutilities.StakingParamsPayload
	{
		payload = NewStakingParamsPayload(message)
	}
	return payload, nil
}

// EncodeStakingPoolResponse encodes responses from the "sdk-utilities" service
// "stakingPool" endpoint.
func EncodeStakingPoolResponse(ctx context.Context, v interface{}, hdr, trlr *metadata.MD) (interface{}, error) {
	result, ok := v.(*sdkutilities.StakingPool2)
	if !ok {
		return nil, goagrpc.ErrInvalidType("sdk-utilities", "stakingPool", "*sdkutilities.StakingPool2", v)
	}
	resp := NewProtoStakingPoolResponse(result)
	return resp, nil
}

// DecodeStakingPoolRequest decodes requests sent to "sdk-utilities" service
// "stakingPool" endpoint.
func DecodeStakingPoolRequest(ctx context.Context, v interface{}, md metadata.MD) (interface{}, error) {
	var (
		message *sdk_utilitiespb.StakingPoolRequest
		ok      bool
	)
	{
		if message, ok = v.(*sdk_utilitiespb.StakingPoolRequest); !ok {
			return nil, goagrpc.ErrInvalidType("sdk-utilities", "stakingPool", "*sdk_utilitiespb.StakingPoolRequest", v)
		}
	}
	var payload *sdkutilities.StakingPoolPayload
	{
		payload = NewStakingPoolPayload(message)
	}
	return payload, nil
}

// EncodeEmoneyInflationResponse encodes responses from the "sdk-utilities"
// service "emoneyInflation" endpoint.
func EncodeEmoneyInflationResponse(ctx context.Context, v interface{}, hdr, trlr *metadata.MD) (interface{}, error) {
	result, ok := v.(*sdkutilities.EmoneyInflation2)
	if !ok {
		return nil, goagrpc.ErrInvalidType("sdk-utilities", "emoneyInflation", "*sdkutilities.EmoneyInflation2", v)
	}
	resp := NewProtoEmoneyInflationResponse(result)
	return resp, nil
}

// DecodeEmoneyInflationRequest decodes requests sent to "sdk-utilities"
// service "emoneyInflation" endpoint.
func DecodeEmoneyInflationRequest(ctx context.Context, v interface{}, md metadata.MD) (interface{}, error) {
	var (
		message *sdk_utilitiespb.EmoneyInflationRequest
		ok      bool
	)
	{
		if message, ok = v.(*sdk_utilitiespb.EmoneyInflationRequest); !ok {
			return nil, goagrpc.ErrInvalidType("sdk-utilities", "emoneyInflation", "*sdk_utilitiespb.EmoneyInflationRequest", v)
		}
	}
	var payload *sdkutilities.EmoneyInflationPayload
	{
		payload = NewEmoneyInflationPayload(message)
	}
	return payload, nil
}

// EncodeIbcChannelClientStateResponse encodes responses from the
// "sdk-utilities" service "ibcChannelClientState" endpoint.
func EncodeIbcChannelClientStateResponse(ctx context.Context, v interface{}, hdr, trlr *metadata.MD) (interface{}, error) {
	result, ok := v.(*sdkutilities.IbcChannelClientState2)
	if !ok {
		return nil, goagrpc.ErrInvalidType("sdk-utilities", "ibcChannelClientState", "*sdkutilities.IbcChannelClientState2", v)
	}
	resp := NewProtoIbcChannelClientStateResponse(result)
	return resp, nil
}

// DecodeIbcChannelClientStateRequest decodes requests sent to "sdk-utilities"
// service "ibcChannelClientState" endpoint.
func DecodeIbcChannelClientStateRequest(ctx context.Context, v interface{}, md metadata.MD) (interface{}, error) {
	var (
		message *sdk_utilitiespb.IbcChannelClientStateRequest
		ok      bool
	)
	{
		if message, ok = v.(*sdk_utilitiespb.IbcChannelClientStateRequest); !ok {
			return nil, goagrpc.ErrInvalidType("sdk-utilities", "ibcChannelClientState", "*sdk_utilitiespb.IbcChannelClientStateRequest", v)
		}
	}
	var payload *sdkutilities.IbcChannelClientStatePayload
	{
		payload = NewIbcChannelClientStatePayload(message)
	}
	return payload, nil
}
