// Code generated by goa v3.7.2, DO NOT EDIT.
//
// sdk-utilities gRPC server types
//
// Command:
// $ goa gen github.com/emerishq/sdk-service-meta

package server

import (
	sdk_utilitiespb "github.com/emerishq/sdk-service-meta/gen/grpc/sdk_utilities/pb"
	sdkutilities "github.com/emerishq/sdk-service-meta/gen/sdk_utilities"
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

// NewProtoAccountNumbersResponse builds the gRPC response type from the result
// of the "accountNumbers" endpoint of the "sdk-utilities" service.
func NewProtoAccountNumbersResponse(result *sdkutilities.AccountNumbers2) *sdk_utilitiespb.AccountNumbersResponse {
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
	if message.PaginationKey != "" {
		v.PaginationKey = &message.PaginationKey
	}
	return v
}

// NewProtoSupplyResponse builds the gRPC response type from the result of the
// "supply" endpoint of the "sdk-utilities" service.
func NewProtoSupplyResponse(result *sdkutilities.Supply2) *sdk_utilitiespb.SupplyResponse {
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
	if result.Pagination != nil {
		message.Pagination = svcSdkutilitiesPaginationToSdkUtilitiespbPagination(result.Pagination)
	}
	return message
}

// NewSupplyDenomPayload builds the payload of the "supplyDenom" endpoint of
// the "sdk-utilities" service from the gRPC request type.
func NewSupplyDenomPayload(message *sdk_utilitiespb.SupplyDenomRequest) *sdkutilities.SupplyDenomPayload {
	v := &sdkutilities.SupplyDenomPayload{
		ChainName: message.ChainName,
	}
	if message.Port != 0 {
		portptr := int(message.Port)
		v.Port = &portptr
	}
	if message.Denom != "" {
		v.Denom = &message.Denom
	}
	return v
}

// NewProtoSupplyDenomResponse builds the gRPC response type from the result of
// the "supplyDenom" endpoint of the "sdk-utilities" service.
func NewProtoSupplyDenomResponse(result *sdkutilities.Supply2) *sdk_utilitiespb.SupplyDenomResponse {
	message := &sdk_utilitiespb.SupplyDenomResponse{}
	if result.Coins != nil {
		message.Coins = make([]*sdk_utilitiespb.Coin, len(result.Coins))
		for i, val := range result.Coins {
			message.Coins[i] = &sdk_utilitiespb.Coin{
				Denom:  val.Denom,
				Amount: val.Amount,
			}
		}
	}
	if result.Pagination != nil {
		message.Pagination = svcSdkutilitiesPaginationToSdkUtilitiespbPagination(result.Pagination)
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

// NewProtoQueryTxResponse builds the gRPC response type from the result of the
// "queryTx" endpoint of the "sdk-utilities" service.
func NewProtoQueryTxResponse(result []byte) *sdk_utilitiespb.QueryTxResponse {
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

// NewProtoBroadcastTxResponse builds the gRPC response type from the result of
// the "broadcastTx" endpoint of the "sdk-utilities" service.
func NewProtoBroadcastTxResponse(result *sdkutilities.TransactionResult) *sdk_utilitiespb.BroadcastTxResponse {
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

// NewProtoTxMetadataResponse builds the gRPC response type from the result of
// the "txMetadata" endpoint of the "sdk-utilities" service.
func NewProtoTxMetadataResponse(result *sdkutilities.TxMessagesMetadata) *sdk_utilitiespb.TxMetadataResponse {
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

// NewProtoBlockResponse builds the gRPC response type from the result of the
// "block" endpoint of the "sdk-utilities" service.
func NewProtoBlockResponse(result *sdkutilities.BlockData) *sdk_utilitiespb.BlockResponse {
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

// NewProtoLiquidityParamsResponse builds the gRPC response type from the
// result of the "liquidityParams" endpoint of the "sdk-utilities" service.
func NewProtoLiquidityParamsResponse(result *sdkutilities.LiquidityParams2) *sdk_utilitiespb.LiquidityParamsResponse {
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

// NewProtoLiquidityPoolsResponse builds the gRPC response type from the result
// of the "liquidityPools" endpoint of the "sdk-utilities" service.
func NewProtoLiquidityPoolsResponse(result *sdkutilities.LiquidityPools2) *sdk_utilitiespb.LiquidityPoolsResponse {
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

// NewProtoMintInflationResponse builds the gRPC response type from the result
// of the "mintInflation" endpoint of the "sdk-utilities" service.
func NewProtoMintInflationResponse(result *sdkutilities.MintInflation2) *sdk_utilitiespb.MintInflationResponse {
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

// NewProtoMintParamsResponse builds the gRPC response type from the result of
// the "mintParams" endpoint of the "sdk-utilities" service.
func NewProtoMintParamsResponse(result *sdkutilities.MintParams2) *sdk_utilitiespb.MintParamsResponse {
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

// NewProtoMintAnnualProvisionResponse builds the gRPC response type from the
// result of the "mintAnnualProvision" endpoint of the "sdk-utilities" service.
func NewProtoMintAnnualProvisionResponse(result *sdkutilities.MintAnnualProvision2) *sdk_utilitiespb.MintAnnualProvisionResponse {
	message := &sdk_utilitiespb.MintAnnualProvisionResponse{
		MintAnnualProvision: result.MintAnnualProvision,
	}
	return message
}

// NewMintEpochProvisionsPayload builds the payload of the
// "mintEpochProvisions" endpoint of the "sdk-utilities" service from the gRPC
// request type.
func NewMintEpochProvisionsPayload(message *sdk_utilitiespb.MintEpochProvisionsRequest) *sdkutilities.MintEpochProvisionsPayload {
	v := &sdkutilities.MintEpochProvisionsPayload{
		ChainName: message.ChainName,
	}
	if message.Port != 0 {
		portptr := int(message.Port)
		v.Port = &portptr
	}
	return v
}

// NewProtoMintEpochProvisionsResponse builds the gRPC response type from the
// result of the "mintEpochProvisions" endpoint of the "sdk-utilities" service.
func NewProtoMintEpochProvisionsResponse(result *sdkutilities.MintEpochProvisions2) *sdk_utilitiespb.MintEpochProvisionsResponse {
	message := &sdk_utilitiespb.MintEpochProvisionsResponse{
		MintEpochProvisions: result.MintEpochProvisions,
	}
	return message
}

// NewDelegatorRewardsPayload builds the payload of the "delegatorRewards"
// endpoint of the "sdk-utilities" service from the gRPC request type.
func NewDelegatorRewardsPayload(message *sdk_utilitiespb.DelegatorRewardsRequest) *sdkutilities.DelegatorRewardsPayload {
	v := &sdkutilities.DelegatorRewardsPayload{
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

// NewProtoDelegatorRewardsResponse builds the gRPC response type from the
// result of the "delegatorRewards" endpoint of the "sdk-utilities" service.
func NewProtoDelegatorRewardsResponse(result *sdkutilities.DelegatorRewards2) *sdk_utilitiespb.DelegatorRewardsResponse {
	message := &sdk_utilitiespb.DelegatorRewardsResponse{}
	if result.Rewards != nil {
		message.Rewards = make([]*sdk_utilitiespb.DelegationDelegatorReward, len(result.Rewards))
		for i, val := range result.Rewards {
			message.Rewards[i] = &sdk_utilitiespb.DelegationDelegatorReward{
				ValidatorAddress: val.ValidatorAddress,
			}
			if val.Rewards != nil {
				message.Rewards[i].Rewards = make([]*sdk_utilitiespb.Coin, len(val.Rewards))
				for j, val := range val.Rewards {
					message.Rewards[i].Rewards[j] = &sdk_utilitiespb.Coin{
						Denom:  val.Denom,
						Amount: val.Amount,
					}
				}
			}
		}
	}
	if result.Total != nil {
		message.Total = make([]*sdk_utilitiespb.Coin, len(result.Total))
		for i, val := range result.Total {
			message.Total[i] = &sdk_utilitiespb.Coin{
				Denom:  val.Denom,
				Amount: val.Amount,
			}
		}
	}
	return message
}

// NewEstimateFeesPayload builds the payload of the "estimateFees" endpoint of
// the "sdk-utilities" service from the gRPC request type.
func NewEstimateFeesPayload(message *sdk_utilitiespb.EstimateFeesRequest) *sdkutilities.EstimateFeesPayload {
	v := &sdkutilities.EstimateFeesPayload{
		ChainName: message.ChainName,
		TxBytes:   message.TxBytes,
	}
	if message.Port != 0 {
		portptr := int(message.Port)
		v.Port = &portptr
	}
	return v
}

// NewProtoEstimateFeesResponse builds the gRPC response type from the result
// of the "estimateFees" endpoint of the "sdk-utilities" service.
func NewProtoEstimateFeesResponse(result *sdkutilities.Simulation) *sdk_utilitiespb.EstimateFeesResponse {
	message := &sdk_utilitiespb.EstimateFeesResponse{
		GasWanted: result.GasWanted,
		GasUsed:   result.GasUsed,
	}
	if result.Fees != nil {
		message.Fees = make([]*sdk_utilitiespb.Coin, len(result.Fees))
		for i, val := range result.Fees {
			message.Fees[i] = &sdk_utilitiespb.Coin{
				Denom:  val.Denom,
				Amount: val.Amount,
			}
		}
	}
	return message
}

// NewStakingParamsPayload builds the payload of the "stakingParams" endpoint
// of the "sdk-utilities" service from the gRPC request type.
func NewStakingParamsPayload(message *sdk_utilitiespb.StakingParamsRequest) *sdkutilities.StakingParamsPayload {
	v := &sdkutilities.StakingParamsPayload{
		ChainName: message.ChainName,
	}
	if message.Port != 0 {
		portptr := int(message.Port)
		v.Port = &portptr
	}
	return v
}

// NewProtoStakingParamsResponse builds the gRPC response type from the result
// of the "stakingParams" endpoint of the "sdk-utilities" service.
func NewProtoStakingParamsResponse(result *sdkutilities.StakingParams2) *sdk_utilitiespb.StakingParamsResponse {
	message := &sdk_utilitiespb.StakingParamsResponse{
		StakingParams: result.StakingParams,
	}
	return message
}

// NewStakingPoolPayload builds the payload of the "stakingPool" endpoint of
// the "sdk-utilities" service from the gRPC request type.
func NewStakingPoolPayload(message *sdk_utilitiespb.StakingPoolRequest) *sdkutilities.StakingPoolPayload {
	v := &sdkutilities.StakingPoolPayload{
		ChainName: message.ChainName,
	}
	if message.Port != 0 {
		portptr := int(message.Port)
		v.Port = &portptr
	}
	return v
}

// NewProtoStakingPoolResponse builds the gRPC response type from the result of
// the "stakingPool" endpoint of the "sdk-utilities" service.
func NewProtoStakingPoolResponse(result *sdkutilities.StakingPool2) *sdk_utilitiespb.StakingPoolResponse {
	message := &sdk_utilitiespb.StakingPoolResponse{
		StakingPool: result.StakingPool,
	}
	return message
}

// NewEmoneyInflationPayload builds the payload of the "emoneyInflation"
// endpoint of the "sdk-utilities" service from the gRPC request type.
func NewEmoneyInflationPayload(message *sdk_utilitiespb.EmoneyInflationRequest) *sdkutilities.EmoneyInflationPayload {
	v := &sdkutilities.EmoneyInflationPayload{
		ChainName: message.ChainName,
	}
	if message.Port != 0 {
		portptr := int(message.Port)
		v.Port = &portptr
	}
	return v
}

// NewProtoEmoneyInflationResponse builds the gRPC response type from the
// result of the "emoneyInflation" endpoint of the "sdk-utilities" service.
func NewProtoEmoneyInflationResponse(result *sdkutilities.EmoneyInflation2) *sdk_utilitiespb.EmoneyInflationResponse {
	message := &sdk_utilitiespb.EmoneyInflationResponse{}
	if result.State != nil {
		message.State = svcSdkutilitiesEmoneyStateToSdkUtilitiespbEmoneyState(result.State)
	}
	return message
}

// NewBudgetParamsPayload builds the payload of the "budgetParams" endpoint of
// the "sdk-utilities" service from the gRPC request type.
func NewBudgetParamsPayload(message *sdk_utilitiespb.BudgetParamsRequest) *sdkutilities.BudgetParamsPayload {
	v := &sdkutilities.BudgetParamsPayload{
		ChainName: message.ChainName,
	}
	if message.Port != 0 {
		portptr := int(message.Port)
		v.Port = &portptr
	}
	return v
}

// NewProtoBudgetParamsResponse builds the gRPC response type from the result
// of the "budgetParams" endpoint of the "sdk-utilities" service.
func NewProtoBudgetParamsResponse(result *sdkutilities.BudgetParams2) *sdk_utilitiespb.BudgetParamsResponse {
	message := &sdk_utilitiespb.BudgetParamsResponse{
		BudgetParams: result.BudgetParams,
	}
	return message
}

// NewDistributionParamsPayload builds the payload of the "distributionParams"
// endpoint of the "sdk-utilities" service from the gRPC request type.
func NewDistributionParamsPayload(message *sdk_utilitiespb.DistributionParamsRequest) *sdkutilities.DistributionParamsPayload {
	v := &sdkutilities.DistributionParamsPayload{
		ChainName: message.ChainName,
	}
	if message.Port != 0 {
		portptr := int(message.Port)
		v.Port = &portptr
	}
	return v
}

// NewProtoDistributionParamsResponse builds the gRPC response type from the
// result of the "distributionParams" endpoint of the "sdk-utilities" service.
func NewProtoDistributionParamsResponse(result *sdkutilities.DistributionParams2) *sdk_utilitiespb.DistributionParamsResponse {
	message := &sdk_utilitiespb.DistributionParamsResponse{
		CommunityTax:        result.CommunityTax,
		BaseProposerReward:  result.BaseProposerReward,
		BonusProposerReward: result.BonusProposerReward,
		WithdrawAddrEnabled: result.WithdrawAddrEnabled,
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

// ValidateEstimateFeesRequest runs the validations defined on
// EstimateFeesRequest.
func ValidateEstimateFeesRequest(message *sdk_utilitiespb.EstimateFeesRequest) (err error) {
	if message.TxBytes == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("txBytes", "message"))
	}
	return
}

// svcSdkutilitiesPaginationToSdkUtilitiespbPagination builds a value of type
// *sdk_utilitiespb.Pagination from a value of type *sdkutilities.Pagination.
func svcSdkutilitiesPaginationToSdkUtilitiespbPagination(v *sdkutilities.Pagination) *sdk_utilitiespb.Pagination {
	if v == nil {
		return nil
	}
	res := &sdk_utilitiespb.Pagination{}
	if v.NextKey != nil {
		res.NextKey = *v.NextKey
	}
	if v.Total != nil {
		res.Total = *v.Total
	}

	return res
}

// protobufSdkUtilitiespbPaginationToSdkutilitiesPagination builds a value of
// type *sdkutilities.Pagination from a value of type
// *sdk_utilitiespb.Pagination.
func protobufSdkUtilitiespbPaginationToSdkutilitiesPagination(v *sdk_utilitiespb.Pagination) *sdkutilities.Pagination {
	if v == nil {
		return nil
	}
	res := &sdkutilities.Pagination{}
	if v.NextKey != "" {
		res.NextKey = &v.NextKey
	}
	if v.Total != "" {
		res.Total = &v.Total
	}

	return res
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

// svcSdkutilitiesEmoneyStateToSdkUtilitiespbEmoneyState builds a value of type
// *sdk_utilitiespb.EmoneyState from a value of type *sdkutilities.EmoneyState.
func svcSdkutilitiesEmoneyStateToSdkUtilitiespbEmoneyState(v *sdkutilities.EmoneyState) *sdk_utilitiespb.EmoneyState {
	res := &sdk_utilitiespb.EmoneyState{
		LastApplied:       v.LastApplied,
		LastAppliedHeight: v.LastAppliedHeight,
	}
	if v.Assets != nil {
		res.Assets = make([]*sdk_utilitiespb.EmoneyAsset, len(v.Assets))
		for i, val := range v.Assets {
			res.Assets[i] = &sdk_utilitiespb.EmoneyAsset{
				Denom:     val.Denom,
				Inflation: val.Inflation,
				Accum:     val.Accum,
			}
		}
	}

	return res
}

// protobufSdkUtilitiespbEmoneyStateToSdkutilitiesEmoneyState builds a value of
// type *sdkutilities.EmoneyState from a value of type
// *sdk_utilitiespb.EmoneyState.
func protobufSdkUtilitiespbEmoneyStateToSdkutilitiesEmoneyState(v *sdk_utilitiespb.EmoneyState) *sdkutilities.EmoneyState {
	res := &sdkutilities.EmoneyState{
		LastApplied:       v.LastApplied,
		LastAppliedHeight: v.LastAppliedHeight,
	}
	if v.Assets != nil {
		res.Assets = make([]*sdkutilities.EmoneyAsset, len(v.Assets))
		for i, val := range v.Assets {
			res.Assets[i] = &sdkutilities.EmoneyAsset{
				Denom:     val.Denom,
				Inflation: val.Inflation,
				Accum:     val.Accum,
			}
		}
	}

	return res
}
