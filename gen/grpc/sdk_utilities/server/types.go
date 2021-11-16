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
		Height: message.Height,
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

// NewAuthPayload builds the payload of the "auth" endpoint of the
// "sdk-utilities" service from the gRPC request type.
func NewAuthPayload(message *sdk_utilitiespb.AuthRequest) *sdkutilities.AuthPayload {
	v := &sdkutilities.AuthPayload{}
	if message.Payload != nil {
		v.Payload = make([]*sdkutilities.TracePayload, len(message.Payload))
		for i, val := range message.Payload {
			v.Payload[i] = &sdkutilities.TracePayload{
				Key:   val.Key,
				Value: val.Value,
			}
			if val.OperationType != "" {
				v.Payload[i].OperationType = &val.OperationType
			}
		}
	}
	return v
}

// NewAuthResponse builds the gRPC response type from the result of the "auth"
// endpoint of the "sdk-utilities" service.
func NewAuthResponse(result []*sdkutilities.Auth) *sdk_utilitiespb.AuthResponse {
	message := &sdk_utilitiespb.AuthResponse{}
	message.Field = make([]*sdk_utilitiespb.Auth, len(result))
	for i, val := range result {
		message.Field[i] = &sdk_utilitiespb.Auth{
			Address:        val.Address,
			SequenceNumber: val.SequenceNumber,
			AccountNumber:  val.AccountNumber,
		}
	}
	return message
}

// NewAuthProcessingErrorError builds the gRPC error response type from the
// error of the "auth" endpoint of the "sdk-utilities" service.
func NewAuthProcessingErrorError(er *sdkutilities.ProcessingError) *sdk_utilitiespb.AuthProcessingErrorError {
	message := &sdk_utilitiespb.AuthProcessingErrorError{}
	if er.Name != nil {
		message.Name = *er.Name
	}
	if er.Code != nil {
		message.Code = int32(*er.Code)
	}
	if er.Errors != nil {
		message.Errors = make([]*sdk_utilitiespb.ErrorObject, len(er.Errors))
		for i, val := range er.Errors {
			message.Errors[i] = &sdk_utilitiespb.ErrorObject{
				Value:        val.Value,
				PayloadIndex: int32(val.PayloadIndex),
			}
		}
	}
	return message
}

// NewBankPayload builds the payload of the "bank" endpoint of the
// "sdk-utilities" service from the gRPC request type.
func NewBankPayload(message *sdk_utilitiespb.BankRequest) *sdkutilities.BankPayload {
	v := &sdkutilities.BankPayload{}
	if message.Payload != nil {
		v.Payload = make([]*sdkutilities.TracePayload, len(message.Payload))
		for i, val := range message.Payload {
			v.Payload[i] = &sdkutilities.TracePayload{
				Key:   val.Key,
				Value: val.Value,
			}
			if val.OperationType != "" {
				v.Payload[i].OperationType = &val.OperationType
			}
		}
	}
	return v
}

// NewBankResponse builds the gRPC response type from the result of the "bank"
// endpoint of the "sdk-utilities" service.
func NewBankResponse(result []*sdkutilities.Balance) *sdk_utilitiespb.BankResponse {
	message := &sdk_utilitiespb.BankResponse{}
	message.Field = make([]*sdk_utilitiespb.Balance, len(result))
	for i, val := range result {
		message.Field[i] = &sdk_utilitiespb.Balance{
			Address: val.Address,
			Amount:  val.Amount,
			Denom:   val.Denom,
		}
	}
	return message
}

// NewBankProcessingErrorError builds the gRPC error response type from the
// error of the "bank" endpoint of the "sdk-utilities" service.
func NewBankProcessingErrorError(er *sdkutilities.ProcessingError) *sdk_utilitiespb.BankProcessingErrorError {
	message := &sdk_utilitiespb.BankProcessingErrorError{}
	if er.Name != nil {
		message.Name = *er.Name
	}
	if er.Code != nil {
		message.Code = int32(*er.Code)
	}
	if er.Errors != nil {
		message.Errors = make([]*sdk_utilitiespb.ErrorObject, len(er.Errors))
		for i, val := range er.Errors {
			message.Errors[i] = &sdk_utilitiespb.ErrorObject{
				Value:        val.Value,
				PayloadIndex: int32(val.PayloadIndex),
			}
		}
	}
	return message
}

// NewDelegationPayload builds the payload of the "delegation" endpoint of the
// "sdk-utilities" service from the gRPC request type.
func NewDelegationPayload(message *sdk_utilitiespb.DelegationRequest) *sdkutilities.DelegationPayload {
	v := &sdkutilities.DelegationPayload{}
	if message.Payload != nil {
		v.Payload = make([]*sdkutilities.TracePayload, len(message.Payload))
		for i, val := range message.Payload {
			v.Payload[i] = &sdkutilities.TracePayload{
				Key:   val.Key,
				Value: val.Value,
			}
			if val.OperationType != "" {
				v.Payload[i].OperationType = &val.OperationType
			}
		}
	}
	return v
}

// NewDelegationResponse builds the gRPC response type from the result of the
// "delegation" endpoint of the "sdk-utilities" service.
func NewDelegationResponse(result []*sdkutilities.Delegation) *sdk_utilitiespb.DelegationResponse {
	message := &sdk_utilitiespb.DelegationResponse{}
	message.Field = make([]*sdk_utilitiespb.Delegation, len(result))
	for i, val := range result {
		message.Field[i] = &sdk_utilitiespb.Delegation{
			Delegator: val.Delegator,
			Validator: val.Validator,
			Amount:    val.Amount,
			Type:      val.Type,
		}
	}
	return message
}

// NewDelegationProcessingErrorError builds the gRPC error response type from
// the error of the "delegation" endpoint of the "sdk-utilities" service.
func NewDelegationProcessingErrorError(er *sdkutilities.ProcessingError) *sdk_utilitiespb.DelegationProcessingErrorError {
	message := &sdk_utilitiespb.DelegationProcessingErrorError{}
	if er.Name != nil {
		message.Name = *er.Name
	}
	if er.Code != nil {
		message.Code = int32(*er.Code)
	}
	if er.Errors != nil {
		message.Errors = make([]*sdk_utilitiespb.ErrorObject, len(er.Errors))
		for i, val := range er.Errors {
			message.Errors[i] = &sdk_utilitiespb.ErrorObject{
				Value:        val.Value,
				PayloadIndex: int32(val.PayloadIndex),
			}
		}
	}
	return message
}

// NewIbcChannelPayload builds the payload of the "ibc_channel" endpoint of the
// "sdk-utilities" service from the gRPC request type.
func NewIbcChannelPayload(message *sdk_utilitiespb.IbcChannelRequest) *sdkutilities.IbcChannelPayload {
	v := &sdkutilities.IbcChannelPayload{}
	if message.Payload != nil {
		v.Payload = make([]*sdkutilities.TracePayload, len(message.Payload))
		for i, val := range message.Payload {
			v.Payload[i] = &sdkutilities.TracePayload{
				Key:   val.Key,
				Value: val.Value,
			}
			if val.OperationType != "" {
				v.Payload[i].OperationType = &val.OperationType
			}
		}
	}
	return v
}

// NewIbcChannelResponse builds the gRPC response type from the result of the
// "ibc_channel" endpoint of the "sdk-utilities" service.
func NewIbcChannelResponse(result []*sdkutilities.IBCChannel) *sdk_utilitiespb.IbcChannelResponse {
	message := &sdk_utilitiespb.IbcChannelResponse{}
	message.Field = make([]*sdk_utilitiespb.IBCChannel, len(result))
	for i, val := range result {
		message.Field[i] = &sdk_utilitiespb.IBCChannel{
			ChannelId:        val.ChannelID,
			CounterChannelId: val.CounterChannelID,
			Port:             val.Port,
			State:            val.State,
		}
		if val.Hops != nil {
			message.Field[i].Hops = make([]string, len(val.Hops))
			for j, val := range val.Hops {
				message.Field[i].Hops[j] = val
			}
		}
	}
	return message
}

// NewIbcChannelProcessingErrorError builds the gRPC error response type from
// the error of the "ibc_channel" endpoint of the "sdk-utilities" service.
func NewIbcChannelProcessingErrorError(er *sdkutilities.ProcessingError) *sdk_utilitiespb.IbcChannelProcessingErrorError {
	message := &sdk_utilitiespb.IbcChannelProcessingErrorError{}
	if er.Name != nil {
		message.Name = *er.Name
	}
	if er.Code != nil {
		message.Code = int32(*er.Code)
	}
	if er.Errors != nil {
		message.Errors = make([]*sdk_utilitiespb.ErrorObject, len(er.Errors))
		for i, val := range er.Errors {
			message.Errors[i] = &sdk_utilitiespb.ErrorObject{
				Value:        val.Value,
				PayloadIndex: int32(val.PayloadIndex),
			}
		}
	}
	return message
}

// NewIbcClientStatePayload builds the payload of the "ibc_client_state"
// endpoint of the "sdk-utilities" service from the gRPC request type.
func NewIbcClientStatePayload(message *sdk_utilitiespb.IbcClientStateRequest) *sdkutilities.IbcClientStatePayload {
	v := &sdkutilities.IbcClientStatePayload{}
	if message.Payload != nil {
		v.Payload = make([]*sdkutilities.TracePayload, len(message.Payload))
		for i, val := range message.Payload {
			v.Payload[i] = &sdkutilities.TracePayload{
				Key:   val.Key,
				Value: val.Value,
			}
			if val.OperationType != "" {
				v.Payload[i].OperationType = &val.OperationType
			}
		}
	}
	return v
}

// NewIbcClientStateResponse builds the gRPC response type from the result of
// the "ibc_client_state" endpoint of the "sdk-utilities" service.
func NewIbcClientStateResponse(result []*sdkutilities.IBCClientState) *sdk_utilitiespb.IbcClientStateResponse {
	message := &sdk_utilitiespb.IbcClientStateResponse{}
	message.Field = make([]*sdk_utilitiespb.IBCClientState, len(result))
	for i, val := range result {
		message.Field[i] = &sdk_utilitiespb.IBCClientState{
			ChainId:        val.ChainID,
			ClientId:       val.ClientID,
			LatestHeight:   val.LatestHeight,
			TrustingPeriod: val.TrustingPeriod,
		}
	}
	return message
}

// NewIbcClientStateProcessingErrorError builds the gRPC error response type
// from the error of the "ibc_client_state" endpoint of the "sdk-utilities"
// service.
func NewIbcClientStateProcessingErrorError(er *sdkutilities.ProcessingError) *sdk_utilitiespb.IbcClientStateProcessingErrorError {
	message := &sdk_utilitiespb.IbcClientStateProcessingErrorError{}
	if er.Name != nil {
		message.Name = *er.Name
	}
	if er.Code != nil {
		message.Code = int32(*er.Code)
	}
	if er.Errors != nil {
		message.Errors = make([]*sdk_utilitiespb.ErrorObject, len(er.Errors))
		for i, val := range er.Errors {
			message.Errors[i] = &sdk_utilitiespb.ErrorObject{
				Value:        val.Value,
				PayloadIndex: int32(val.PayloadIndex),
			}
		}
	}
	return message
}

// NewIbcConnectionPayload builds the payload of the "ibc_connection" endpoint
// of the "sdk-utilities" service from the gRPC request type.
func NewIbcConnectionPayload(message *sdk_utilitiespb.IbcConnectionRequest) *sdkutilities.IbcConnectionPayload {
	v := &sdkutilities.IbcConnectionPayload{}
	if message.Payload != nil {
		v.Payload = make([]*sdkutilities.TracePayload, len(message.Payload))
		for i, val := range message.Payload {
			v.Payload[i] = &sdkutilities.TracePayload{
				Key:   val.Key,
				Value: val.Value,
			}
			if val.OperationType != "" {
				v.Payload[i].OperationType = &val.OperationType
			}
		}
	}
	return v
}

// NewIbcConnectionResponse builds the gRPC response type from the result of
// the "ibc_connection" endpoint of the "sdk-utilities" service.
func NewIbcConnectionResponse(result []*sdkutilities.IBCConnection) *sdk_utilitiespb.IbcConnectionResponse {
	message := &sdk_utilitiespb.IbcConnectionResponse{}
	message.Field = make([]*sdk_utilitiespb.IBCConnection, len(result))
	for i, val := range result {
		message.Field[i] = &sdk_utilitiespb.IBCConnection{
			ConnectionId:        val.ConnectionID,
			ClientId:            val.ClientID,
			State:               val.State,
			CounterConnectionId: val.CounterConnectionID,
			CounterClientId:     val.CounterClientID,
		}
	}
	return message
}

// NewIbcConnectionProcessingErrorError builds the gRPC error response type
// from the error of the "ibc_connection" endpoint of the "sdk-utilities"
// service.
func NewIbcConnectionProcessingErrorError(er *sdkutilities.ProcessingError) *sdk_utilitiespb.IbcConnectionProcessingErrorError {
	message := &sdk_utilitiespb.IbcConnectionProcessingErrorError{}
	if er.Name != nil {
		message.Name = *er.Name
	}
	if er.Code != nil {
		message.Code = int32(*er.Code)
	}
	if er.Errors != nil {
		message.Errors = make([]*sdk_utilitiespb.ErrorObject, len(er.Errors))
		for i, val := range er.Errors {
			message.Errors[i] = &sdk_utilitiespb.ErrorObject{
				Value:        val.Value,
				PayloadIndex: int32(val.PayloadIndex),
			}
		}
	}
	return message
}

// NewIbcDenomTracePayload builds the payload of the "ibc_denom_trace" endpoint
// of the "sdk-utilities" service from the gRPC request type.
func NewIbcDenomTracePayload(message *sdk_utilitiespb.IbcDenomTraceRequest) *sdkutilities.IbcDenomTracePayload {
	v := &sdkutilities.IbcDenomTracePayload{}
	if message.Payload != nil {
		v.Payload = make([]*sdkutilities.TracePayload, len(message.Payload))
		for i, val := range message.Payload {
			v.Payload[i] = &sdkutilities.TracePayload{
				Key:   val.Key,
				Value: val.Value,
			}
			if val.OperationType != "" {
				v.Payload[i].OperationType = &val.OperationType
			}
		}
	}
	return v
}

// NewIbcDenomTraceResponse builds the gRPC response type from the result of
// the "ibc_denom_trace" endpoint of the "sdk-utilities" service.
func NewIbcDenomTraceResponse(result []*sdkutilities.IBCDenomTrace) *sdk_utilitiespb.IbcDenomTraceResponse {
	message := &sdk_utilitiespb.IbcDenomTraceResponse{}
	message.Field = make([]*sdk_utilitiespb.IBCDenomTrace, len(result))
	for i, val := range result {
		message.Field[i] = &sdk_utilitiespb.IBCDenomTrace{
			Path:      val.Path,
			BaseDenom: val.BaseDenom,
			Hash:      val.Hash,
		}
	}
	return message
}

// NewIbcDenomTraceProcessingErrorError builds the gRPC error response type
// from the error of the "ibc_denom_trace" endpoint of the "sdk-utilities"
// service.
func NewIbcDenomTraceProcessingErrorError(er *sdkutilities.ProcessingError) *sdk_utilitiespb.IbcDenomTraceProcessingErrorError {
	message := &sdk_utilitiespb.IbcDenomTraceProcessingErrorError{}
	if er.Name != nil {
		message.Name = *er.Name
	}
	if er.Code != nil {
		message.Code = int32(*er.Code)
	}
	if er.Errors != nil {
		message.Errors = make([]*sdk_utilitiespb.ErrorObject, len(er.Errors))
		for i, val := range er.Errors {
			message.Errors[i] = &sdk_utilitiespb.ErrorObject{
				Value:        val.Value,
				PayloadIndex: int32(val.PayloadIndex),
			}
		}
	}
	return message
}

// NewUnbondingDelegationPayload builds the payload of the
// "unbondingDelegation" endpoint of the "sdk-utilities" service from the gRPC
// request type.
func NewUnbondingDelegationPayload(message *sdk_utilitiespb.UnbondingDelegationRequest) *sdkutilities.UnbondingDelegationPayload {
	v := &sdkutilities.UnbondingDelegationPayload{}
	if message.Payload != nil {
		v.Payload = make([]*sdkutilities.TracePayload, len(message.Payload))
		for i, val := range message.Payload {
			v.Payload[i] = &sdkutilities.TracePayload{
				Key:   val.Key,
				Value: val.Value,
			}
			if val.OperationType != "" {
				v.Payload[i].OperationType = &val.OperationType
			}
		}
	}
	return v
}

// NewUnbondingDelegationResponse builds the gRPC response type from the result
// of the "unbondingDelegation" endpoint of the "sdk-utilities" service.
func NewUnbondingDelegationResponse(result []*sdkutilities.UnbondingDelegation) *sdk_utilitiespb.UnbondingDelegationResponse {
	message := &sdk_utilitiespb.UnbondingDelegationResponse{}
	message.Field = make([]*sdk_utilitiespb.UnbondingDelegation, len(result))
	for i, val := range result {
		message.Field[i] = &sdk_utilitiespb.UnbondingDelegation{
			Delegator: val.Delegator,
			Validator: val.Validator,
			Type:      val.Type,
		}
		if val.Entries != nil {
			message.Field[i].Entries = make([]*sdk_utilitiespb.UnbondingDelegationEntry, len(val.Entries))
			for j, val := range val.Entries {
				message.Field[i].Entries[j] = &sdk_utilitiespb.UnbondingDelegationEntry{
					Balance:        val.Balance,
					InitialBalance: val.InitialBalance,
					CreationHeight: val.CreationHeight,
					CompletionTime: val.CompletionTime,
				}
			}
		}
	}
	return message
}

// NewUnbondingDelegationProcessingErrorError builds the gRPC error response
// type from the error of the "unbondingDelegation" endpoint of the
// "sdk-utilities" service.
func NewUnbondingDelegationProcessingErrorError(er *sdkutilities.ProcessingError) *sdk_utilitiespb.UnbondingDelegationProcessingErrorError {
	message := &sdk_utilitiespb.UnbondingDelegationProcessingErrorError{}
	if er.Name != nil {
		message.Name = *er.Name
	}
	if er.Code != nil {
		message.Code = int32(*er.Code)
	}
	if er.Errors != nil {
		message.Errors = make([]*sdk_utilitiespb.ErrorObject, len(er.Errors))
		for i, val := range er.Errors {
			message.Errors[i] = &sdk_utilitiespb.ErrorObject{
				Value:        val.Value,
				PayloadIndex: int32(val.PayloadIndex),
			}
		}
	}
	return message
}

// NewValidatorPayload builds the payload of the "validator" endpoint of the
// "sdk-utilities" service from the gRPC request type.
func NewValidatorPayload(message *sdk_utilitiespb.ValidatorRequest) *sdkutilities.ValidatorPayload {
	v := &sdkutilities.ValidatorPayload{}
	if message.Payload != nil {
		v.Payload = make([]*sdkutilities.TracePayload, len(message.Payload))
		for i, val := range message.Payload {
			v.Payload[i] = &sdkutilities.TracePayload{
				Key:   val.Key,
				Value: val.Value,
			}
			if val.OperationType != "" {
				v.Payload[i].OperationType = &val.OperationType
			}
		}
	}
	return v
}

// NewValidatorResponse builds the gRPC response type from the result of the
// "validator" endpoint of the "sdk-utilities" service.
func NewValidatorResponse(result []*sdkutilities.Validator) *sdk_utilitiespb.ValidatorResponse {
	message := &sdk_utilitiespb.ValidatorResponse{}
	message.Field = make([]*sdk_utilitiespb.Validator, len(result))
	for i, val := range result {
		message.Field[i] = &sdk_utilitiespb.Validator{
			OperatorAddress:      val.OperatorAddress,
			ConsensusPubKeyType:  val.ConsensusPubKeyType,
			ConsensusPubKeyValue: val.ConsensusPubKeyValue,
			Jailed:               val.Jailed,
			Status:               val.Status,
			Tokens:               val.Tokens,
			DelegatorShares:      val.DelegatorShares,
			Moniker:              val.Moniker,
			Identity:             val.Identity,
			Website:              val.Website,
			SecurityContact:      val.SecurityContact,
			Details:              val.Details,
			UnbondingHeight:      val.UnbondingHeight,
			UnbondingTime:        val.UnbondingTime,
			CommissionRate:       val.CommissionRate,
			MaxRate:              val.MaxRate,
			MaxChangeRate:        val.MaxChangeRate,
			UpdateTime:           val.UpdateTime,
			MinSelfDelegation:    val.MinSelfDelegation,
			Type:                 val.Type,
		}
	}
	return message
}

// NewValidatorProcessingErrorError builds the gRPC error response type from
// the error of the "validator" endpoint of the "sdk-utilities" service.
func NewValidatorProcessingErrorError(er *sdkutilities.ProcessingError) *sdk_utilitiespb.ValidatorProcessingErrorError {
	message := &sdk_utilitiespb.ValidatorProcessingErrorError{}
	if er.Name != nil {
		message.Name = *er.Name
	}
	if er.Code != nil {
		message.Code = int32(*er.Code)
	}
	if er.Errors != nil {
		message.Errors = make([]*sdk_utilitiespb.ErrorObject, len(er.Errors))
		for i, val := range er.Errors {
			message.Errors[i] = &sdk_utilitiespb.ErrorObject{
				Value:        val.Value,
				PayloadIndex: int32(val.PayloadIndex),
			}
		}
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
