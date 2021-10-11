// Code generated by goa v3.5.2, DO NOT EDIT.
//
// sdk-utilities gRPC client types
//
// Command:
// $ goa gen github.com/allinbits/sdk-service-meta

package client

import (
	sdk_utilitiespb "github.com/allinbits/sdk-service-meta/gen/grpc/sdk_utilities/pb"
	sdkutilities "github.com/allinbits/sdk-service-meta/gen/sdk_utilities"
	goa "goa.design/goa/v3/pkg"
)

// NewSupplyRequest builds the gRPC request type from the payload of the
// "supply" endpoint of the "sdk-utilities" service.
func NewSupplyRequest(payload *sdkutilities.SupplyPayload) *sdk_utilitiespb.SupplyRequest {
	message := &sdk_utilitiespb.SupplyRequest{
		ChainName: payload.ChainName,
	}
	if payload.Port != nil {
		message.Port = int32(*payload.Port)
	}
	return message
}

// NewSupplyResult builds the result type of the "supply" endpoint of the
// "sdk-utilities" service from the gRPC response type.
func NewSupplyResult(message *sdk_utilitiespb.SupplyResponse) *sdkutilities.Supply2 {
	result := &sdkutilities.Supply2{}
	if message.Coins != nil {
		result.Coins = make([]*sdkutilities.Coin, len(message.Coins))
		for i, val := range message.Coins {
			result.Coins[i] = &sdkutilities.Coin{
				Denom:  val.Denom,
				Amount: val.Amount,
			}
		}
	}
	return result
}

// NewQueryTxRequest builds the gRPC request type from the payload of the
// "queryTx" endpoint of the "sdk-utilities" service.
func NewQueryTxRequest(payload *sdkutilities.QueryTxPayload) *sdk_utilitiespb.QueryTxRequest {
	message := &sdk_utilitiespb.QueryTxRequest{
		ChainName: payload.ChainName,
		Hash:      payload.Hash,
	}
	if payload.Port != nil {
		message.Port = int32(*payload.Port)
	}
	return message
}

// NewQueryTxResult builds the result type of the "queryTx" endpoint of the
// "sdk-utilities" service from the gRPC response type.
func NewQueryTxResult(message *sdk_utilitiespb.QueryTxResponse) []byte {
	result := message.Field
	return result
}

// NewBroadcastTxRequest builds the gRPC request type from the payload of the
// "broadcastTx" endpoint of the "sdk-utilities" service.
func NewBroadcastTxRequest(payload *sdkutilities.BroadcastTxPayload) *sdk_utilitiespb.BroadcastTxRequest {
	message := &sdk_utilitiespb.BroadcastTxRequest{
		ChainName: payload.ChainName,
		TxBytes:   payload.TxBytes,
	}
	if payload.Port != nil {
		message.Port = int32(*payload.Port)
	}
	return message
}

// NewBroadcastTxResult builds the result type of the "broadcastTx" endpoint of
// the "sdk-utilities" service from the gRPC response type.
func NewBroadcastTxResult(message *sdk_utilitiespb.BroadcastTxResponse) *sdkutilities.TransactionResult {
	result := &sdkutilities.TransactionResult{
		Hash: message.Hash,
	}
	return result
}

// NewTxMetadataRequest builds the gRPC request type from the payload of the
// "txMetadata" endpoint of the "sdk-utilities" service.
func NewTxMetadataRequest(payload *sdkutilities.TxMetadataPayload) *sdk_utilitiespb.TxMetadataRequest {
	message := &sdk_utilitiespb.TxMetadataRequest{
		TxBytes: payload.TxBytes,
	}
	return message
}

// NewTxMetadataResult builds the result type of the "txMetadata" endpoint of
// the "sdk-utilities" service from the gRPC response type.
func NewTxMetadataResult(message *sdk_utilitiespb.TxMetadataResponse) *sdkutilities.TxMessagesMetadata {
	result := &sdkutilities.TxMessagesMetadata{}
	if message.MessagesMetadata != nil {
		result.MessagesMetadata = make([]*sdkutilities.MsgMetadata, len(message.MessagesMetadata))
		for i, val := range message.MessagesMetadata {
			result.MessagesMetadata[i] = &sdkutilities.MsgMetadata{
				MsgType: val.MsgType,
			}
			if val.IbcTransferMetadata != nil {
				result.MessagesMetadata[i].IbcTransferMetadata = protobufSdkUtilitiespbIBCTransferMetadataToSdkutilitiesIBCTransferMetadata(val.IbcTransferMetadata)
			}
		}
	}
	return result
}

// NewAuthRequest builds the gRPC request type from the payload of the "auth"
// endpoint of the "sdk-utilities" service.
func NewAuthRequest(payload *sdkutilities.AuthPayload) *sdk_utilitiespb.AuthRequest {
	message := &sdk_utilitiespb.AuthRequest{}
	if payload.Payload != nil {
		message.Payload = make([]*sdk_utilitiespb.TracePayload, len(payload.Payload))
		for i, val := range payload.Payload {
			message.Payload[i] = &sdk_utilitiespb.TracePayload{
				Key:   val.Key,
				Value: val.Value,
			}
			if val.OperationType != nil {
				message.Payload[i].OperationType = *val.OperationType
			}
		}
	}
	return message
}

// NewAuthResult builds the result type of the "auth" endpoint of the
// "sdk-utilities" service from the gRPC response type.
func NewAuthResult(message *sdk_utilitiespb.AuthResponse) []*sdkutilities.Auth {
	result := make([]*sdkutilities.Auth, len(message.Field))
	for i, val := range message.Field {
		result[i] = &sdkutilities.Auth{
			Address:        val.Address,
			SequenceNumber: val.SequenceNumber,
			AccountNumber:  val.AccountNumber,
		}
		if val.Address == "" {
			result[i].Address = ""
		}
		if val.SequenceNumber == 0 {
			result[i].SequenceNumber = 0
		}
		if val.AccountNumber == 0 {
			result[i].AccountNumber = 0
		}
	}
	return result
}

// NewAuthProcessingErrorError builds the error type of the "auth" endpoint of
// the "sdk-utilities" service from the gRPC error response type.
func NewAuthProcessingErrorError(message *sdk_utilitiespb.AuthProcessingErrorError) *sdkutilities.ProcessingError {
	er := &sdkutilities.ProcessingError{}
	if message.Name != "" {
		er.Name = &message.Name
	}
	if message.Code != 0 {
		codeptr := int(message.Code)
		er.Code = &codeptr
	}
	if message.Errors != nil {
		er.Errors = make([]*sdkutilities.ErrorObject, len(message.Errors))
		for i, val := range message.Errors {
			er.Errors[i] = &sdkutilities.ErrorObject{
				Value:        val.Value,
				PayloadIndex: int(val.PayloadIndex),
			}
			if val.Value == "" {
				er.Errors[i].Value = ""
			}
			if val.PayloadIndex == 0 {
				er.Errors[i].PayloadIndex = 0
			}
		}
	}
	return er
}

// NewBankRequest builds the gRPC request type from the payload of the "bank"
// endpoint of the "sdk-utilities" service.
func NewBankRequest(payload *sdkutilities.BankPayload) *sdk_utilitiespb.BankRequest {
	message := &sdk_utilitiespb.BankRequest{}
	if payload.Payload != nil {
		message.Payload = make([]*sdk_utilitiespb.TracePayload, len(payload.Payload))
		for i, val := range payload.Payload {
			message.Payload[i] = &sdk_utilitiespb.TracePayload{
				Key:   val.Key,
				Value: val.Value,
			}
			if val.OperationType != nil {
				message.Payload[i].OperationType = *val.OperationType
			}
		}
	}
	return message
}

// NewBankResult builds the result type of the "bank" endpoint of the
// "sdk-utilities" service from the gRPC response type.
func NewBankResult(message *sdk_utilitiespb.BankResponse) []*sdkutilities.Balance {
	result := make([]*sdkutilities.Balance, len(message.Field))
	for i, val := range message.Field {
		result[i] = &sdkutilities.Balance{
			Address: val.Address,
			Amount:  val.Amount,
			Denom:   val.Denom,
		}
	}
	return result
}

// NewBankProcessingErrorError builds the error type of the "bank" endpoint of
// the "sdk-utilities" service from the gRPC error response type.
func NewBankProcessingErrorError(message *sdk_utilitiespb.BankProcessingErrorError) *sdkutilities.ProcessingError {
	er := &sdkutilities.ProcessingError{}
	if message.Name != "" {
		er.Name = &message.Name
	}
	if message.Code != 0 {
		codeptr := int(message.Code)
		er.Code = &codeptr
	}
	if message.Errors != nil {
		er.Errors = make([]*sdkutilities.ErrorObject, len(message.Errors))
		for i, val := range message.Errors {
			er.Errors[i] = &sdkutilities.ErrorObject{
				Value:        val.Value,
				PayloadIndex: int(val.PayloadIndex),
			}
			if val.Value == "" {
				er.Errors[i].Value = ""
			}
			if val.PayloadIndex == 0 {
				er.Errors[i].PayloadIndex = 0
			}
		}
	}
	return er
}

// NewDelegationRequest builds the gRPC request type from the payload of the
// "delegation" endpoint of the "sdk-utilities" service.
func NewDelegationRequest(payload *sdkutilities.DelegationPayload) *sdk_utilitiespb.DelegationRequest {
	message := &sdk_utilitiespb.DelegationRequest{}
	if payload.Payload != nil {
		message.Payload = make([]*sdk_utilitiespb.TracePayload, len(payload.Payload))
		for i, val := range payload.Payload {
			message.Payload[i] = &sdk_utilitiespb.TracePayload{
				Key:   val.Key,
				Value: val.Value,
			}
			if val.OperationType != nil {
				message.Payload[i].OperationType = *val.OperationType
			}
		}
	}
	return message
}

// NewDelegationResult builds the result type of the "delegation" endpoint of
// the "sdk-utilities" service from the gRPC response type.
func NewDelegationResult(message *sdk_utilitiespb.DelegationResponse) []*sdkutilities.Delegation {
	result := make([]*sdkutilities.Delegation, len(message.Field))
	for i, val := range message.Field {
		result[i] = &sdkutilities.Delegation{
			Delegator: val.Delegator,
			Validator: val.Validator,
			Amount:    val.Amount,
			Type:      val.Type,
		}
		if val.Delegator == "" {
			result[i].Delegator = ""
		}
		if val.Validator == "" {
			result[i].Validator = ""
		}
		if val.Amount == "" {
			result[i].Amount = ""
		}
		if val.Type == "" {
			result[i].Type = "create"
		}
	}
	return result
}

// NewDelegationProcessingErrorError builds the error type of the "delegation"
// endpoint of the "sdk-utilities" service from the gRPC error response type.
func NewDelegationProcessingErrorError(message *sdk_utilitiespb.DelegationProcessingErrorError) *sdkutilities.ProcessingError {
	er := &sdkutilities.ProcessingError{}
	if message.Name != "" {
		er.Name = &message.Name
	}
	if message.Code != 0 {
		codeptr := int(message.Code)
		er.Code = &codeptr
	}
	if message.Errors != nil {
		er.Errors = make([]*sdkutilities.ErrorObject, len(message.Errors))
		for i, val := range message.Errors {
			er.Errors[i] = &sdkutilities.ErrorObject{
				Value:        val.Value,
				PayloadIndex: int(val.PayloadIndex),
			}
			if val.Value == "" {
				er.Errors[i].Value = ""
			}
			if val.PayloadIndex == 0 {
				er.Errors[i].PayloadIndex = 0
			}
		}
	}
	return er
}

// NewIbcChannelRequest builds the gRPC request type from the payload of the
// "ibc_channel" endpoint of the "sdk-utilities" service.
func NewIbcChannelRequest(payload *sdkutilities.IbcChannelPayload) *sdk_utilitiespb.IbcChannelRequest {
	message := &sdk_utilitiespb.IbcChannelRequest{}
	if payload.Payload != nil {
		message.Payload = make([]*sdk_utilitiespb.TracePayload, len(payload.Payload))
		for i, val := range payload.Payload {
			message.Payload[i] = &sdk_utilitiespb.TracePayload{
				Key:   val.Key,
				Value: val.Value,
			}
			if val.OperationType != nil {
				message.Payload[i].OperationType = *val.OperationType
			}
		}
	}
	return message
}

// NewIbcChannelResult builds the result type of the "ibc_channel" endpoint of
// the "sdk-utilities" service from the gRPC response type.
func NewIbcChannelResult(message *sdk_utilitiespb.IbcChannelResponse) []*sdkutilities.IBCChannel {
	result := make([]*sdkutilities.IBCChannel, len(message.Field))
	for i, val := range message.Field {
		result[i] = &sdkutilities.IBCChannel{
			ChannelID:        val.ChannelId,
			CounterChannelID: val.CounterChannelId,
			Port:             val.Port,
			State:            val.State,
		}
		if val.ChannelId == "" {
			result[i].ChannelID = ""
		}
		if val.CounterChannelId == "" {
			result[i].CounterChannelID = ""
		}
		if val.Hops != nil {
			result[i].Hops = make([]string, len(val.Hops))
			for j, val := range val.Hops {
				result[i].Hops[j] = val
			}
		}
		if val.Port == "" {
			result[i].Port = ""
		}
		if val.State == 0 {
			result[i].State = 0
		}
	}
	return result
}

// NewIbcChannelProcessingErrorError builds the error type of the "ibc_channel"
// endpoint of the "sdk-utilities" service from the gRPC error response type.
func NewIbcChannelProcessingErrorError(message *sdk_utilitiespb.IbcChannelProcessingErrorError) *sdkutilities.ProcessingError {
	er := &sdkutilities.ProcessingError{}
	if message.Name != "" {
		er.Name = &message.Name
	}
	if message.Code != 0 {
		codeptr := int(message.Code)
		er.Code = &codeptr
	}
	if message.Errors != nil {
		er.Errors = make([]*sdkutilities.ErrorObject, len(message.Errors))
		for i, val := range message.Errors {
			er.Errors[i] = &sdkutilities.ErrorObject{
				Value:        val.Value,
				PayloadIndex: int(val.PayloadIndex),
			}
			if val.Value == "" {
				er.Errors[i].Value = ""
			}
			if val.PayloadIndex == 0 {
				er.Errors[i].PayloadIndex = 0
			}
		}
	}
	return er
}

// NewIbcClientStateRequest builds the gRPC request type from the payload of
// the "ibc_client_state" endpoint of the "sdk-utilities" service.
func NewIbcClientStateRequest(payload *sdkutilities.IbcClientStatePayload) *sdk_utilitiespb.IbcClientStateRequest {
	message := &sdk_utilitiespb.IbcClientStateRequest{}
	if payload.Payload != nil {
		message.Payload = make([]*sdk_utilitiespb.TracePayload, len(payload.Payload))
		for i, val := range payload.Payload {
			message.Payload[i] = &sdk_utilitiespb.TracePayload{
				Key:   val.Key,
				Value: val.Value,
			}
			if val.OperationType != nil {
				message.Payload[i].OperationType = *val.OperationType
			}
		}
	}
	return message
}

// NewIbcClientStateResult builds the result type of the "ibc_client_state"
// endpoint of the "sdk-utilities" service from the gRPC response type.
func NewIbcClientStateResult(message *sdk_utilitiespb.IbcClientStateResponse) []*sdkutilities.IBCClientState {
	result := make([]*sdkutilities.IBCClientState, len(message.Field))
	for i, val := range message.Field {
		result[i] = &sdkutilities.IBCClientState{
			ChainID:        val.ChainId,
			ClientID:       val.ClientId,
			LatestHeight:   val.LatestHeight,
			TrustingPeriod: val.TrustingPeriod,
		}
		if val.ChainId == "" {
			result[i].ChainID = ""
		}
		if val.ClientId == "" {
			result[i].ClientID = ""
		}
		if val.LatestHeight == 0 {
			result[i].LatestHeight = 0
		}
		if val.TrustingPeriod == 0 {
			result[i].TrustingPeriod = 0
		}
	}
	return result
}

// NewIbcClientStateProcessingErrorError builds the error type of the
// "ibc_client_state" endpoint of the "sdk-utilities" service from the gRPC
// error response type.
func NewIbcClientStateProcessingErrorError(message *sdk_utilitiespb.IbcClientStateProcessingErrorError) *sdkutilities.ProcessingError {
	er := &sdkutilities.ProcessingError{}
	if message.Name != "" {
		er.Name = &message.Name
	}
	if message.Code != 0 {
		codeptr := int(message.Code)
		er.Code = &codeptr
	}
	if message.Errors != nil {
		er.Errors = make([]*sdkutilities.ErrorObject, len(message.Errors))
		for i, val := range message.Errors {
			er.Errors[i] = &sdkutilities.ErrorObject{
				Value:        val.Value,
				PayloadIndex: int(val.PayloadIndex),
			}
			if val.Value == "" {
				er.Errors[i].Value = ""
			}
			if val.PayloadIndex == 0 {
				er.Errors[i].PayloadIndex = 0
			}
		}
	}
	return er
}

// NewIbcConnectionRequest builds the gRPC request type from the payload of the
// "ibc_connection" endpoint of the "sdk-utilities" service.
func NewIbcConnectionRequest(payload *sdkutilities.IbcConnectionPayload) *sdk_utilitiespb.IbcConnectionRequest {
	message := &sdk_utilitiespb.IbcConnectionRequest{}
	if payload.Payload != nil {
		message.Payload = make([]*sdk_utilitiespb.TracePayload, len(payload.Payload))
		for i, val := range payload.Payload {
			message.Payload[i] = &sdk_utilitiespb.TracePayload{
				Key:   val.Key,
				Value: val.Value,
			}
			if val.OperationType != nil {
				message.Payload[i].OperationType = *val.OperationType
			}
		}
	}
	return message
}

// NewIbcConnectionResult builds the result type of the "ibc_connection"
// endpoint of the "sdk-utilities" service from the gRPC response type.
func NewIbcConnectionResult(message *sdk_utilitiespb.IbcConnectionResponse) []*sdkutilities.IBCConnection {
	result := make([]*sdkutilities.IBCConnection, len(message.Field))
	for i, val := range message.Field {
		result[i] = &sdkutilities.IBCConnection{
			ConnectionID:        val.ConnectionId,
			ClientID:            val.ClientId,
			State:               val.State,
			CounterConnectionID: val.CounterConnectionId,
			CounterClientID:     val.CounterClientId,
		}
		if val.ConnectionId == "" {
			result[i].ConnectionID = ""
		}
		if val.ClientId == "" {
			result[i].ClientID = ""
		}
		if val.State == "" {
			result[i].State = ""
		}
		if val.CounterConnectionId == "" {
			result[i].CounterConnectionID = ""
		}
		if val.CounterClientId == "" {
			result[i].CounterClientID = ""
		}
	}
	return result
}

// NewIbcConnectionProcessingErrorError builds the error type of the
// "ibc_connection" endpoint of the "sdk-utilities" service from the gRPC error
// response type.
func NewIbcConnectionProcessingErrorError(message *sdk_utilitiespb.IbcConnectionProcessingErrorError) *sdkutilities.ProcessingError {
	er := &sdkutilities.ProcessingError{}
	if message.Name != "" {
		er.Name = &message.Name
	}
	if message.Code != 0 {
		codeptr := int(message.Code)
		er.Code = &codeptr
	}
	if message.Errors != nil {
		er.Errors = make([]*sdkutilities.ErrorObject, len(message.Errors))
		for i, val := range message.Errors {
			er.Errors[i] = &sdkutilities.ErrorObject{
				Value:        val.Value,
				PayloadIndex: int(val.PayloadIndex),
			}
			if val.Value == "" {
				er.Errors[i].Value = ""
			}
			if val.PayloadIndex == 0 {
				er.Errors[i].PayloadIndex = 0
			}
		}
	}
	return er
}

// NewIbcDenomTraceRequest builds the gRPC request type from the payload of the
// "ibc_denom_trace" endpoint of the "sdk-utilities" service.
func NewIbcDenomTraceRequest(payload *sdkutilities.IbcDenomTracePayload) *sdk_utilitiespb.IbcDenomTraceRequest {
	message := &sdk_utilitiespb.IbcDenomTraceRequest{}
	if payload.Payload != nil {
		message.Payload = make([]*sdk_utilitiespb.TracePayload, len(payload.Payload))
		for i, val := range payload.Payload {
			message.Payload[i] = &sdk_utilitiespb.TracePayload{
				Key:   val.Key,
				Value: val.Value,
			}
			if val.OperationType != nil {
				message.Payload[i].OperationType = *val.OperationType
			}
		}
	}
	return message
}

// NewIbcDenomTraceResult builds the result type of the "ibc_denom_trace"
// endpoint of the "sdk-utilities" service from the gRPC response type.
func NewIbcDenomTraceResult(message *sdk_utilitiespb.IbcDenomTraceResponse) []*sdkutilities.IBCDenomTrace {
	result := make([]*sdkutilities.IBCDenomTrace, len(message.Field))
	for i, val := range message.Field {
		result[i] = &sdkutilities.IBCDenomTrace{
			Path:      val.Path,
			BaseDenom: val.BaseDenom,
			Hash:      val.Hash,
		}
		if val.Path == "" {
			result[i].Path = ""
		}
		if val.BaseDenom == "" {
			result[i].BaseDenom = ""
		}
		if val.Hash == "" {
			result[i].Hash = ""
		}
	}
	return result
}

// NewIbcDenomTraceProcessingErrorError builds the error type of the
// "ibc_denom_trace" endpoint of the "sdk-utilities" service from the gRPC
// error response type.
func NewIbcDenomTraceProcessingErrorError(message *sdk_utilitiespb.IbcDenomTraceProcessingErrorError) *sdkutilities.ProcessingError {
	er := &sdkutilities.ProcessingError{}
	if message.Name != "" {
		er.Name = &message.Name
	}
	if message.Code != 0 {
		codeptr := int(message.Code)
		er.Code = &codeptr
	}
	if message.Errors != nil {
		er.Errors = make([]*sdkutilities.ErrorObject, len(message.Errors))
		for i, val := range message.Errors {
			er.Errors[i] = &sdkutilities.ErrorObject{
				Value:        val.Value,
				PayloadIndex: int(val.PayloadIndex),
			}
			if val.Value == "" {
				er.Errors[i].Value = ""
			}
			if val.PayloadIndex == 0 {
				er.Errors[i].PayloadIndex = 0
			}
		}
	}
	return er
}

// NewUnbondingDelegationRequest builds the gRPC request type from the payload
// of the "unbondingDelegation" endpoint of the "sdk-utilities" service.
func NewUnbondingDelegationRequest(payload *sdkutilities.UnbondingDelegationPayload) *sdk_utilitiespb.UnbondingDelegationRequest {
	message := &sdk_utilitiespb.UnbondingDelegationRequest{}
	if payload.Payload != nil {
		message.Payload = make([]*sdk_utilitiespb.TracePayload, len(payload.Payload))
		for i, val := range payload.Payload {
			message.Payload[i] = &sdk_utilitiespb.TracePayload{
				Key:   val.Key,
				Value: val.Value,
			}
			if val.OperationType != nil {
				message.Payload[i].OperationType = *val.OperationType
			}
		}
	}
	return message
}

// NewUnbondingDelegationResult builds the result type of the
// "unbondingDelegation" endpoint of the "sdk-utilities" service from the gRPC
// response type.
func NewUnbondingDelegationResult(message *sdk_utilitiespb.UnbondingDelegationResponse) []*sdkutilities.UnbondingDelegation {
	result := make([]*sdkutilities.UnbondingDelegation, len(message.Field))
	for i, val := range message.Field {
		result[i] = &sdkutilities.UnbondingDelegation{
			Delegator: val.Delegator,
			Validator: val.Validator,
			Type:      val.Type,
		}
		if val.Delegator == "" {
			result[i].Delegator = ""
		}
		if val.Validator == "" {
			result[i].Validator = ""
		}
		if val.Type == "" {
			result[i].Type = "create"
		}
		if val.Entries != nil {
			result[i].Entries = make([]*sdkutilities.UnbondingDelegationEntry, len(val.Entries))
			for j, val := range val.Entries {
				result[i].Entries[j] = &sdkutilities.UnbondingDelegationEntry{
					Balance:        val.Balance,
					InitialBalance: val.InitialBalance,
					CreationHeight: val.CreationHeight,
					CompletionTime: val.CompletionTime,
				}
				if val.Balance == "" {
					result[i].Entries[j].Balance = ""
				}
				if val.InitialBalance == "" {
					result[i].Entries[j].InitialBalance = ""
				}
				if val.CreationHeight == 0 {
					result[i].Entries[j].CreationHeight = 0
				}
				if val.CompletionTime == 0 {
					result[i].Entries[j].CompletionTime = 0
				}
			}
		}
	}
	return result
}

// NewUnbondingDelegationProcessingErrorError builds the error type of the
// "unbondingDelegation" endpoint of the "sdk-utilities" service from the gRPC
// error response type.
func NewUnbondingDelegationProcessingErrorError(message *sdk_utilitiespb.UnbondingDelegationProcessingErrorError) *sdkutilities.ProcessingError {
	er := &sdkutilities.ProcessingError{}
	if message.Name != "" {
		er.Name = &message.Name
	}
	if message.Code != 0 {
		codeptr := int(message.Code)
		er.Code = &codeptr
	}
	if message.Errors != nil {
		er.Errors = make([]*sdkutilities.ErrorObject, len(message.Errors))
		for i, val := range message.Errors {
			er.Errors[i] = &sdkutilities.ErrorObject{
				Value:        val.Value,
				PayloadIndex: int(val.PayloadIndex),
			}
			if val.Value == "" {
				er.Errors[i].Value = ""
			}
			if val.PayloadIndex == 0 {
				er.Errors[i].PayloadIndex = 0
			}
		}
	}
	return er
}

// ValidateSupplyResponse runs the validations defined on SupplyResponse.
func ValidateSupplyResponse(message *sdk_utilitiespb.SupplyResponse) (err error) {
	if message.Coins == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("coins", "message"))
	}
	return
}

// ValidateCoin runs the validations defined on Coin.
func ValidateCoin(message *sdk_utilitiespb.Coin) (err error) {

	return
}

// ValidateTxMetadataResponse runs the validations defined on
// TxMetadataResponse.
func ValidateTxMetadataResponse(message *sdk_utilitiespb.TxMetadataResponse) (err error) {
	if message.MessagesMetadata == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("messagesMetadata", "message"))
	}
	return
}

// ValidateMsgMetadata runs the validations defined on MsgMetadata.
func ValidateMsgMetadata(message *sdk_utilitiespb.MsgMetadata) (err error) {

	return
}

// ValidateIBCTransferMetadata runs the validations defined on
// IBCTransferMetadata.
func ValidateIBCTransferMetadata(message *sdk_utilitiespb.IBCTransferMetadata) (err error) {

	return
}

// ValidateIBCHeight runs the validations defined on IBCHeight.
func ValidateIBCHeight(message *sdk_utilitiespb.IBCHeight) (err error) {

	return
}

// ValidateDelegationResponse runs the validations defined on
// DelegationResponse.
func ValidateDelegationResponse(message *sdk_utilitiespb.DelegationResponse) (err error) {
	for _, e := range message.Field {
		if e != nil {
			if err2 := ValidateDelegation(e); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	return
}

// ValidateDelegation runs the validations defined on Delegation.
func ValidateDelegation(message *sdk_utilitiespb.Delegation) (err error) {
	if !(message.Type == "delete" || message.Type == "create") {
		err = goa.MergeErrors(err, goa.InvalidEnumValueError("message.type", message.Type, []interface{}{"delete", "create"}))
	}
	return
}

// ValidateUnbondingDelegationResponse runs the validations defined on
// UnbondingDelegationResponse.
func ValidateUnbondingDelegationResponse(message *sdk_utilitiespb.UnbondingDelegationResponse) (err error) {
	for _, e := range message.Field {
		if e != nil {
			if err2 := ValidateUnbondingDelegation(e); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	return
}

// ValidateUnbondingDelegation runs the validations defined on
// UnbondingDelegation.
func ValidateUnbondingDelegation(message *sdk_utilitiespb.UnbondingDelegation) (err error) {
	if !(message.Type == "delete" || message.Type == "create") {
		err = goa.MergeErrors(err, goa.InvalidEnumValueError("message.type", message.Type, []interface{}{"delete", "create"}))
	}
	return
}

// ValidateUnbondingDelegationEntry runs the validations defined on
// UnbondingDelegationEntry.
func ValidateUnbondingDelegationEntry(message *sdk_utilitiespb.UnbondingDelegationEntry) (err error) {

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
