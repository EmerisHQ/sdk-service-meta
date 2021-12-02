// Code generated by goa v3.5.2, DO NOT EDIT.
//
// sdk-utilities gRPC server types
//
// Command:
// $ goa gen github.com/allinbits/sdk-service-meta

package server

import (
	sdk_utilitiespb "github.com/allinbits/sdk-service-meta/gen/grpc/sdk_utilities/pb"
	sdkutilities "github.com/allinbits/sdk-service-meta/gen/sdk_utilities"
	goa "goa.design/goa/v3/pkg"
)

// NewAccountNumbersPayload builds the payload of the "accountNumbers" endpoint
// of the "sdk-utilities" service from the gRPC request type.
func NewAccountNumbersPayload(message *sdk_utilitiespb.AccountNumbersRequest) *sdkutilities.AccountNumbersPayload {
	v := &sdkutilities.AccountNumbersPayload{
		ChainName: message.ChainName,
	}
	if message.Port != 0 {
		portptr := int(message.Port)
		v.Port = &portptr
	}
	if message.Bech32Prefix != "" {
		v.Bech32Prefix = &message.Bech32Prefix
	}
	if message.AddresHex != "" {
		v.AddresHex = &message.AddresHex
	}
	return v
}

// NewAccountNumbersResponse builds the gRPC response type from the result of
// the "accountNumbers" endpoint of the "sdk-utilities" service.
func NewAccountNumbersResponse(result *sdkutilities.AccountNumbers2) *sdk_utilitiespb.AccountNumbersResponse {
	message := &sdk_utilitiespb.AccountNumbersResponse{
		AccountNumber:  result.AccountNumber,
		SequenceNumber: result.SequenceNumber,
		Bech32Address:  result.Bech32Address,
	}
	return message
}

// NewSupplyPayload builds the payload of the "supply" endpoint of the
// "sdk-utilities" service from the gRPC request type.
func NewSupplyPayload(message *sdk_utilitiespb.SupplyRequest) *sdkutilities.SupplyPayload {
	v := &sdkutilities.SupplyPayload{
		ChainName: message.ChainName,
	}
	if message.Port != 0 {
		portptr := int(message.Port)
		v.Port = &portptr
	}
	return v
}

// NewSupplyResponse builds the gRPC response type from the result of the
// "supply" endpoint of the "sdk-utilities" service.
func NewSupplyResponse(result *sdkutilities.Supply2) *sdk_utilitiespb.SupplyResponse {
	message := &sdk_utilitiespb.SupplyResponse{}
	if result.Coins != nil {
		message.Coins = make([]*sdk_utilitiespb.Coin, len(result.Coins))
		for i, val := range result.Coins {
			message.Coins[i] = &sdk_utilitiespb.Coin{
				Denom:  val.Denom,
				Amount: val.Amount,
			}
		}
	}
	return message
}

// NewQueryTxPayload builds the payload of the "queryTx" endpoint of the
// "sdk-utilities" service from the gRPC request type.
func NewQueryTxPayload(message *sdk_utilitiespb.QueryTxRequest) *sdkutilities.QueryTxPayload {
	v := &sdkutilities.QueryTxPayload{
		ChainName: message.ChainName,
		Hash:      message.Hash,
	}
	if message.Port != 0 {
		portptr := int(message.Port)
		v.Port = &portptr
	}
	return v
}

// NewQueryTxResponse builds the gRPC response type from the result of the
// "queryTx" endpoint of the "sdk-utilities" service.
func NewQueryTxResponse(result []byte) *sdk_utilitiespb.QueryTxResponse {
	message := &sdk_utilitiespb.QueryTxResponse{}
	message.Field = result
	return message
}

// NewBroadcastTxPayload builds the payload of the "broadcastTx" endpoint of
// the "sdk-utilities" service from the gRPC request type.
func NewBroadcastTxPayload(message *sdk_utilitiespb.BroadcastTxRequest) *sdkutilities.BroadcastTxPayload {
	v := &sdkutilities.BroadcastTxPayload{
		ChainName: message.ChainName,
		TxBytes:   message.TxBytes,
	}
	if message.Port != 0 {
		portptr := int(message.Port)
		v.Port = &portptr
	}
	return v
}

// NewBroadcastTxResponse builds the gRPC response type from the result of the
// "broadcastTx" endpoint of the "sdk-utilities" service.
func NewBroadcastTxResponse(result *sdkutilities.TransactionResult) *sdk_utilitiespb.BroadcastTxResponse {
	message := &sdk_utilitiespb.BroadcastTxResponse{
		Hash: result.Hash,
	}
	return message
}

// NewTxMetadataPayload builds the payload of the "txMetadata" endpoint of the
// "sdk-utilities" service from the gRPC request type.
func NewTxMetadataPayload(message *sdk_utilitiespb.TxMetadataRequest) *sdkutilities.TxMetadataPayload {
	v := &sdkutilities.TxMetadataPayload{
		TxBytes: message.TxBytes,
	}
	return v
}

// NewTxMetadataResponse builds the gRPC response type from the result of the
// "txMetadata" endpoint of the "sdk-utilities" service.
func NewTxMetadataResponse(result *sdkutilities.TxMessagesMetadata) *sdk_utilitiespb.TxMetadataResponse {
	message := &sdk_utilitiespb.TxMetadataResponse{}
	if result.MessagesMetadata != nil {
		message.MessagesMetadata = make([]*sdk_utilitiespb.MsgMetadata, len(result.MessagesMetadata))
		for i, val := range result.MessagesMetadata {
			message.MessagesMetadata[i] = &sdk_utilitiespb.MsgMetadata{
				MsgType: val.MsgType,
			}
			if val.IbcTransferMetadata != nil {
				message.MessagesMetadata[i].IbcTransferMetadata = svcSdkutilitiesIBCTransferMetadataToSdkUtilitiespbIBCTransferMetadata(val.IbcTransferMetadata)
			}
		}
	}
	return message
}

// NewBlockPayload builds the payload of the "block" endpoint of the
// "sdk-utilities" service from the gRPC request type.
func NewBlockPayload(message *sdk_utilitiespb.BlockRequest) *sdkutilities.BlockPayload {
	v := &sdkutilities.BlockPayload{
		ChainName: message.ChainName,
		Height:    message.Height,
	}
	if message.Port != 0 {
		portptr := int(message.Port)
		v.Port = &portptr
	}
	return v
}

// NewBlockResponse builds the gRPC response type from the result of the
// "block" endpoint of the "sdk-utilities" service.
func NewBlockResponse(result *sdkutilities.BlockData) *sdk_utilitiespb.BlockResponse {
	message := &sdk_utilitiespb.BlockResponse{
		Height: result.Height,
		Block:  result.Block,
	}
	return message
}

// NewLiquidityParamsPayload builds the payload of the "liquidityParams"
// endpoint of the "sdk-utilities" service from the gRPC request type.
func NewLiquidityParamsPayload(message *sdk_utilitiespb.LiquidityParamsRequest) *sdkutilities.LiquidityParamsPayload {
	v := &sdkutilities.LiquidityParamsPayload{
		ChainName: message.ChainName,
	}
	if message.Port != 0 {
		portptr := int(message.Port)
		v.Port = &portptr
	}
	return v
}

// NewLiquidityParamsResponse builds the gRPC response type from the result of
// the "liquidityParams" endpoint of the "sdk-utilities" service.
func NewLiquidityParamsResponse(result *sdkutilities.LiquidityParams2) *sdk_utilitiespb.LiquidityParamsResponse {
	message := &sdk_utilitiespb.LiquidityParamsResponse{
		LiquidityParams: result.LiquidityParams,
	}
	return message
}

// NewLiquidityPoolsPayload builds the payload of the "liquidityPools" endpoint
// of the "sdk-utilities" service from the gRPC request type.
func NewLiquidityPoolsPayload(message *sdk_utilitiespb.LiquidityPoolsRequest) *sdkutilities.LiquidityPoolsPayload {
	v := &sdkutilities.LiquidityPoolsPayload{
		ChainName: message.ChainName,
	}
	if message.Port != 0 {
		portptr := int(message.Port)
		v.Port = &portptr
	}
	return v
}

// NewLiquidityPoolsResponse builds the gRPC response type from the result of
// the "liquidityPools" endpoint of the "sdk-utilities" service.
func NewLiquidityPoolsResponse(result *sdkutilities.LiquidityPools2) *sdk_utilitiespb.LiquidityPoolsResponse {
	message := &sdk_utilitiespb.LiquidityPoolsResponse{
		LiquidityPools: result.LiquidityPools,
	}
	return message
}

// NewMintInflationPayload builds the payload of the "mintInflation" endpoint
// of the "sdk-utilities" service from the gRPC request type.
func NewMintInflationPayload(message *sdk_utilitiespb.MintInflationRequest) *sdkutilities.MintInflationPayload {
	v := &sdkutilities.MintInflationPayload{
		ChainName: message.ChainName,
	}
	if message.Port != 0 {
		portptr := int(message.Port)
		v.Port = &portptr
	}
	return v
}

// NewMintInflationResponse builds the gRPC response type from the result of
// the "mintInflation" endpoint of the "sdk-utilities" service.
func NewMintInflationResponse(result *sdkutilities.MintInflation2) *sdk_utilitiespb.MintInflationResponse {
	message := &sdk_utilitiespb.MintInflationResponse{
		MintInflation: result.MintInflation,
	}
	return message
}

// NewMintParamsPayload builds the payload of the "mintParams" endpoint of the
// "sdk-utilities" service from the gRPC request type.
func NewMintParamsPayload(message *sdk_utilitiespb.MintParamsRequest) *sdkutilities.MintParamsPayload {
	v := &sdkutilities.MintParamsPayload{
		ChainName: message.ChainName,
	}
	if message.Port != 0 {
		portptr := int(message.Port)
		v.Port = &portptr
	}
	return v
}

// NewMintParamsResponse builds the gRPC response type from the result of the
// "mintParams" endpoint of the "sdk-utilities" service.
func NewMintParamsResponse(result *sdkutilities.MintParams2) *sdk_utilitiespb.MintParamsResponse {
	message := &sdk_utilitiespb.MintParamsResponse{
		MintParams: result.MintParams,
	}
	return message
}

// NewMintAnnualProvisionPayload builds the payload of the
// "mintAnnualProvision" endpoint of the "sdk-utilities" service from the gRPC
// request type.
func NewMintAnnualProvisionPayload(message *sdk_utilitiespb.MintAnnualProvisionRequest) *sdkutilities.MintAnnualProvisionPayload {
	v := &sdkutilities.MintAnnualProvisionPayload{
		ChainName: message.ChainName,
	}
	if message.Port != 0 {
		portptr := int(message.Port)
		v.Port = &portptr
	}
	return v
}

// NewMintAnnualProvisionResponse builds the gRPC response type from the result
// of the "mintAnnualProvision" endpoint of the "sdk-utilities" service.
func NewMintAnnualProvisionResponse(result *sdkutilities.MintAnnualProvision2) *sdk_utilitiespb.MintAnnualProvisionResponse {
	message := &sdk_utilitiespb.MintAnnualProvisionResponse{
		MintAnnualProvision: result.MintAnnualProvision,
	}
	return message
}

// ValidateBroadcastTxRequest runs the validations defined on
// BroadcastTxRequest.
func ValidateBroadcastTxRequest(message *sdk_utilitiespb.BroadcastTxRequest) (err error) {
	if message.TxBytes == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("txBytes", "message"))
	}
	return
}

// ValidateTxMetadataRequest runs the validations defined on TxMetadataRequest.
func ValidateTxMetadataRequest(message *sdk_utilitiespb.TxMetadataRequest) (err error) {
	if message.TxBytes == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("txBytes", "message"))
	}
	return
}

// svcSdkutilitiesIBCTransferMetadataToSdkUtilitiespbIBCTransferMetadata builds
// a value of type *sdk_utilitiespb.IBCTransferMetadata from a value of type
// *sdkutilities.IBCTransferMetadata.
func svcSdkutilitiesIBCTransferMetadataToSdkUtilitiespbIBCTransferMetadata(v *sdkutilities.IBCTransferMetadata) *sdk_utilitiespb.IBCTransferMetadata {
	if v == nil {
		return nil
	}
	res := &sdk_utilitiespb.IBCTransferMetadata{}
	if v.SourcePort != nil {
		res.SourcePort = *v.SourcePort
	}
	if v.SourceChannel != nil {
		res.SourceChannel = *v.SourceChannel
	}
	if v.Sender != nil {
		res.Sender = *v.Sender
	}
	if v.Receiver != nil {
		res.Receiver = *v.Receiver
	}
	if v.TiemoutTimestamp != nil {
		res.TiemoutTimestamp = *v.TiemoutTimestamp
	}
	if v.Token != nil {
		res.Token = svcSdkutilitiesCoinToSdkUtilitiespbCoin(v.Token)
	}
	if v.TimeoutHeight != nil {
		res.TimeoutHeight = svcSdkutilitiesIBCHeightToSdkUtilitiespbIBCHeight(v.TimeoutHeight)
	}

	return res
}

// svcSdkutilitiesCoinToSdkUtilitiespbCoin builds a value of type
// *sdk_utilitiespb.Coin from a value of type *sdkutilities.Coin.
func svcSdkutilitiesCoinToSdkUtilitiespbCoin(v *sdkutilities.Coin) *sdk_utilitiespb.Coin {
	if v == nil {
		return nil
	}
	res := &sdk_utilitiespb.Coin{
		Denom:  v.Denom,
		Amount: v.Amount,
	}

	return res
}

// svcSdkutilitiesIBCHeightToSdkUtilitiespbIBCHeight builds a value of type
// *sdk_utilitiespb.IBCHeight from a value of type *sdkutilities.IBCHeight.
func svcSdkutilitiesIBCHeightToSdkUtilitiespbIBCHeight(v *sdkutilities.IBCHeight) *sdk_utilitiespb.IBCHeight {
	if v == nil {
		return nil
	}
	res := &sdk_utilitiespb.IBCHeight{}
	if v.RevisionNumber != nil {
		res.RevisionNumber = *v.RevisionNumber
	}
	if v.RevisionHeight != nil {
		res.RevisionHeight = *v.RevisionHeight
	}

	return res
}

// protobufSdkUtilitiespbIBCTransferMetadataToSdkutilitiesIBCTransferMetadata
// builds a value of type *sdkutilities.IBCTransferMetadata from a value of
// type *sdk_utilitiespb.IBCTransferMetadata.
func protobufSdkUtilitiespbIBCTransferMetadataToSdkutilitiesIBCTransferMetadata(v *sdk_utilitiespb.IBCTransferMetadata) *sdkutilities.IBCTransferMetadata {
	if v == nil {
		return nil
	}
	res := &sdkutilities.IBCTransferMetadata{}
	if v.SourcePort != "" {
		res.SourcePort = &v.SourcePort
	}
	if v.SourceChannel != "" {
		res.SourceChannel = &v.SourceChannel
	}
	if v.Sender != "" {
		res.Sender = &v.Sender
	}
	if v.Receiver != "" {
		res.Receiver = &v.Receiver
	}
	if v.TiemoutTimestamp != 0 {
		res.TiemoutTimestamp = &v.TiemoutTimestamp
	}
	if v.Token != nil {
		res.Token = protobufSdkUtilitiespbCoinToSdkutilitiesCoin(v.Token)
	}
	if v.TimeoutHeight != nil {
		res.TimeoutHeight = protobufSdkUtilitiespbIBCHeightToSdkutilitiesIBCHeight(v.TimeoutHeight)
	}

	return res
}

// protobufSdkUtilitiespbCoinToSdkutilitiesCoin builds a value of type
// *sdkutilities.Coin from a value of type *sdk_utilitiespb.Coin.
func protobufSdkUtilitiespbCoinToSdkutilitiesCoin(v *sdk_utilitiespb.Coin) *sdkutilities.Coin {
	if v == nil {
		return nil
	}
	res := &sdkutilities.Coin{
		Denom:  v.Denom,
		Amount: v.Amount,
	}

	return res
}

// protobufSdkUtilitiespbIBCHeightToSdkutilitiesIBCHeight builds a value of
// type *sdkutilities.IBCHeight from a value of type *sdk_utilitiespb.IBCHeight.
func protobufSdkUtilitiespbIBCHeightToSdkutilitiesIBCHeight(v *sdk_utilitiespb.IBCHeight) *sdkutilities.IBCHeight {
	if v == nil {
		return nil
	}
	res := &sdkutilities.IBCHeight{}
	if v.RevisionNumber != 0 {
		res.RevisionNumber = &v.RevisionNumber
	}
	if v.RevisionHeight != 0 {
		res.RevisionHeight = &v.RevisionHeight
	}

	return res
}
