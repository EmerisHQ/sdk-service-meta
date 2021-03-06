// Code generated by goa v3.7.5, DO NOT EDIT.
//
// sdk-utilities gRPC client types
//
// Command:
// $ goa gen github.com/emerishq/sdk-service-meta

package client

import (
	sdk_utilitiespb "github.com/emerishq/sdk-service-meta/gen/grpc/sdk_utilities/pb"
	sdkutilities "github.com/emerishq/sdk-service-meta/gen/sdk_utilities"
	goa "goa.design/goa/v3/pkg"
)

// NewProtoAccountNumbersRequest builds the gRPC request type from the payload
// of the "accountNumbers" endpoint of the "sdk-utilities" service.
func NewProtoAccountNumbersRequest(payload *sdkutilities.AccountNumbersPayload) *sdk_utilitiespb.AccountNumbersRequest {
	message := &sdk_utilitiespb.AccountNumbersRequest{
		ChainName: payload.ChainName,
	}
	if payload.Port != nil {
		message.Port = int32(*payload.Port)
	}
	if payload.Bech32Prefix != nil {
		message.Bech32Prefix = *payload.Bech32Prefix
	}
	if payload.AddresHex != nil {
		message.AddresHex = *payload.AddresHex
	}
	return message
}

// NewAccountNumbersResult builds the result type of the "accountNumbers"
// endpoint of the "sdk-utilities" service from the gRPC response type.
func NewAccountNumbersResult(message *sdk_utilitiespb.AccountNumbersResponse) *sdkutilities.AccountNumbers2 {
	result := &sdkutilities.AccountNumbers2{
		AccountNumber:  message.AccountNumber,
		SequenceNumber: message.SequenceNumber,
		Bech32Address:  message.Bech32Address,
	}
	return result
}

// NewProtoSupplyRequest builds the gRPC request type from the payload of the
// "supply" endpoint of the "sdk-utilities" service.
func NewProtoSupplyRequest(payload *sdkutilities.SupplyPayload) *sdk_utilitiespb.SupplyRequest {
	message := &sdk_utilitiespb.SupplyRequest{
		ChainName: payload.ChainName,
	}
	if payload.Port != nil {
		message.Port = int32(*payload.Port)
	}
	if payload.PaginationKey != nil {
		message.PaginationKey = *payload.PaginationKey
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
	if message.Pagination != nil {
		result.Pagination = protobufSdkUtilitiespbPaginationToSdkutilitiesPagination(message.Pagination)
	}
	return result
}

// NewProtoSupplyDenomRequest builds the gRPC request type from the payload of
// the "supplyDenom" endpoint of the "sdk-utilities" service.
func NewProtoSupplyDenomRequest(payload *sdkutilities.SupplyDenomPayload) *sdk_utilitiespb.SupplyDenomRequest {
	message := &sdk_utilitiespb.SupplyDenomRequest{
		ChainName: payload.ChainName,
	}
	if payload.Port != nil {
		message.Port = int32(*payload.Port)
	}
	if payload.Denom != nil {
		message.Denom = *payload.Denom
	}
	return message
}

// NewSupplyDenomResult builds the result type of the "supplyDenom" endpoint of
// the "sdk-utilities" service from the gRPC response type.
func NewSupplyDenomResult(message *sdk_utilitiespb.SupplyDenomResponse) *sdkutilities.Supply2 {
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
	if message.Pagination != nil {
		result.Pagination = protobufSdkUtilitiespbPaginationToSdkutilitiesPagination(message.Pagination)
	}
	return result
}

// NewProtoQueryTxRequest builds the gRPC request type from the payload of the
// "queryTx" endpoint of the "sdk-utilities" service.
func NewProtoQueryTxRequest(payload *sdkutilities.QueryTxPayload) *sdk_utilitiespb.QueryTxRequest {
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

// NewProtoBroadcastTxRequest builds the gRPC request type from the payload of
// the "broadcastTx" endpoint of the "sdk-utilities" service.
func NewProtoBroadcastTxRequest(payload *sdkutilities.BroadcastTxPayload) *sdk_utilitiespb.BroadcastTxRequest {
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

// NewProtoTxMetadataRequest builds the gRPC request type from the payload of
// the "txMetadata" endpoint of the "sdk-utilities" service.
func NewProtoTxMetadataRequest(payload *sdkutilities.TxMetadataPayload) *sdk_utilitiespb.TxMetadataRequest {
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

// NewProtoBlockRequest builds the gRPC request type from the payload of the
// "block" endpoint of the "sdk-utilities" service.
func NewProtoBlockRequest(payload *sdkutilities.BlockPayload) *sdk_utilitiespb.BlockRequest {
	message := &sdk_utilitiespb.BlockRequest{
		ChainName: payload.ChainName,
		Height:    payload.Height,
	}
	if payload.Port != nil {
		message.Port = int32(*payload.Port)
	}
	return message
}

// NewBlockResult builds the result type of the "block" endpoint of the
// "sdk-utilities" service from the gRPC response type.
func NewBlockResult(message *sdk_utilitiespb.BlockResponse) *sdkutilities.BlockData {
	result := &sdkutilities.BlockData{
		Height: message.Height,
		Block:  message.Block,
	}
	return result
}

// NewProtoLiquidityParamsRequest builds the gRPC request type from the payload
// of the "liquidityParams" endpoint of the "sdk-utilities" service.
func NewProtoLiquidityParamsRequest(payload *sdkutilities.LiquidityParamsPayload) *sdk_utilitiespb.LiquidityParamsRequest {
	message := &sdk_utilitiespb.LiquidityParamsRequest{
		ChainName: payload.ChainName,
	}
	if payload.Port != nil {
		message.Port = int32(*payload.Port)
	}
	return message
}

// NewLiquidityParamsResult builds the result type of the "liquidityParams"
// endpoint of the "sdk-utilities" service from the gRPC response type.
func NewLiquidityParamsResult(message *sdk_utilitiespb.LiquidityParamsResponse) *sdkutilities.LiquidityParams2 {
	result := &sdkutilities.LiquidityParams2{
		LiquidityParams: message.LiquidityParams,
	}
	return result
}

// NewProtoLiquidityPoolsRequest builds the gRPC request type from the payload
// of the "liquidityPools" endpoint of the "sdk-utilities" service.
func NewProtoLiquidityPoolsRequest(payload *sdkutilities.LiquidityPoolsPayload) *sdk_utilitiespb.LiquidityPoolsRequest {
	message := &sdk_utilitiespb.LiquidityPoolsRequest{
		ChainName: payload.ChainName,
	}
	if payload.Port != nil {
		message.Port = int32(*payload.Port)
	}
	return message
}

// NewLiquidityPoolsResult builds the result type of the "liquidityPools"
// endpoint of the "sdk-utilities" service from the gRPC response type.
func NewLiquidityPoolsResult(message *sdk_utilitiespb.LiquidityPoolsResponse) *sdkutilities.LiquidityPools2 {
	result := &sdkutilities.LiquidityPools2{
		LiquidityPools: message.LiquidityPools,
	}
	return result
}

// NewProtoMintInflationRequest builds the gRPC request type from the payload
// of the "mintInflation" endpoint of the "sdk-utilities" service.
func NewProtoMintInflationRequest(payload *sdkutilities.MintInflationPayload) *sdk_utilitiespb.MintInflationRequest {
	message := &sdk_utilitiespb.MintInflationRequest{
		ChainName: payload.ChainName,
	}
	if payload.Port != nil {
		message.Port = int32(*payload.Port)
	}
	return message
}

// NewMintInflationResult builds the result type of the "mintInflation"
// endpoint of the "sdk-utilities" service from the gRPC response type.
func NewMintInflationResult(message *sdk_utilitiespb.MintInflationResponse) *sdkutilities.MintInflation2 {
	result := &sdkutilities.MintInflation2{
		MintInflation: message.MintInflation,
	}
	return result
}

// NewProtoMintParamsRequest builds the gRPC request type from the payload of
// the "mintParams" endpoint of the "sdk-utilities" service.
func NewProtoMintParamsRequest(payload *sdkutilities.MintParamsPayload) *sdk_utilitiespb.MintParamsRequest {
	message := &sdk_utilitiespb.MintParamsRequest{
		ChainName: payload.ChainName,
	}
	if payload.Port != nil {
		message.Port = int32(*payload.Port)
	}
	return message
}

// NewMintParamsResult builds the result type of the "mintParams" endpoint of
// the "sdk-utilities" service from the gRPC response type.
func NewMintParamsResult(message *sdk_utilitiespb.MintParamsResponse) *sdkutilities.MintParams2 {
	result := &sdkutilities.MintParams2{
		MintParams: message.MintParams,
	}
	return result
}

// NewProtoMintAnnualProvisionRequest builds the gRPC request type from the
// payload of the "mintAnnualProvision" endpoint of the "sdk-utilities" service.
func NewProtoMintAnnualProvisionRequest(payload *sdkutilities.MintAnnualProvisionPayload) *sdk_utilitiespb.MintAnnualProvisionRequest {
	message := &sdk_utilitiespb.MintAnnualProvisionRequest{
		ChainName: payload.ChainName,
	}
	if payload.Port != nil {
		message.Port = int32(*payload.Port)
	}
	return message
}

// NewMintAnnualProvisionResult builds the result type of the
// "mintAnnualProvision" endpoint of the "sdk-utilities" service from the gRPC
// response type.
func NewMintAnnualProvisionResult(message *sdk_utilitiespb.MintAnnualProvisionResponse) *sdkutilities.MintAnnualProvision2 {
	result := &sdkutilities.MintAnnualProvision2{
		MintAnnualProvision: message.MintAnnualProvision,
	}
	return result
}

// NewProtoMintEpochProvisionsRequest builds the gRPC request type from the
// payload of the "mintEpochProvisions" endpoint of the "sdk-utilities" service.
func NewProtoMintEpochProvisionsRequest(payload *sdkutilities.MintEpochProvisionsPayload) *sdk_utilitiespb.MintEpochProvisionsRequest {
	message := &sdk_utilitiespb.MintEpochProvisionsRequest{
		ChainName: payload.ChainName,
	}
	if payload.Port != nil {
		message.Port = int32(*payload.Port)
	}
	return message
}

// NewMintEpochProvisionsResult builds the result type of the
// "mintEpochProvisions" endpoint of the "sdk-utilities" service from the gRPC
// response type.
func NewMintEpochProvisionsResult(message *sdk_utilitiespb.MintEpochProvisionsResponse) *sdkutilities.MintEpochProvisions2 {
	result := &sdkutilities.MintEpochProvisions2{
		MintEpochProvisions: message.MintEpochProvisions,
	}
	return result
}

// NewProtoDelegatorRewardsRequest builds the gRPC request type from the
// payload of the "delegatorRewards" endpoint of the "sdk-utilities" service.
func NewProtoDelegatorRewardsRequest(payload *sdkutilities.DelegatorRewardsPayload) *sdk_utilitiespb.DelegatorRewardsRequest {
	message := &sdk_utilitiespb.DelegatorRewardsRequest{
		ChainName: payload.ChainName,
	}
	if payload.Port != nil {
		message.Port = int32(*payload.Port)
	}
	if payload.Bech32Prefix != nil {
		message.Bech32Prefix = *payload.Bech32Prefix
	}
	if payload.AddresHex != nil {
		message.AddresHex = *payload.AddresHex
	}
	return message
}

// NewDelegatorRewardsResult builds the result type of the "delegatorRewards"
// endpoint of the "sdk-utilities" service from the gRPC response type.
func NewDelegatorRewardsResult(message *sdk_utilitiespb.DelegatorRewardsResponse) *sdkutilities.DelegatorRewards2 {
	result := &sdkutilities.DelegatorRewards2{}
	if message.Rewards != nil {
		result.Rewards = make([]*sdkutilities.DelegationDelegatorReward, len(message.Rewards))
		for i, val := range message.Rewards {
			result.Rewards[i] = &sdkutilities.DelegationDelegatorReward{
				ValidatorAddress: val.ValidatorAddress,
			}
			if val.Rewards != nil {
				result.Rewards[i].Rewards = make([]*sdkutilities.Coin, len(val.Rewards))
				for j, val := range val.Rewards {
					result.Rewards[i].Rewards[j] = &sdkutilities.Coin{
						Denom:  val.Denom,
						Amount: val.Amount,
					}
				}
			}
		}
	}
	if message.Total != nil {
		result.Total = make([]*sdkutilities.Coin, len(message.Total))
		for i, val := range message.Total {
			result.Total[i] = &sdkutilities.Coin{
				Denom:  val.Denom,
				Amount: val.Amount,
			}
		}
	}
	return result
}

// NewProtoEstimateFeesRequest builds the gRPC request type from the payload of
// the "estimateFees" endpoint of the "sdk-utilities" service.
func NewProtoEstimateFeesRequest(payload *sdkutilities.EstimateFeesPayload) *sdk_utilitiespb.EstimateFeesRequest {
	message := &sdk_utilitiespb.EstimateFeesRequest{
		ChainName: payload.ChainName,
		TxBytes:   payload.TxBytes,
	}
	if payload.Port != nil {
		message.Port = int32(*payload.Port)
	}
	return message
}

// NewEstimateFeesResult builds the result type of the "estimateFees" endpoint
// of the "sdk-utilities" service from the gRPC response type.
func NewEstimateFeesResult(message *sdk_utilitiespb.EstimateFeesResponse) *sdkutilities.Simulation {
	result := &sdkutilities.Simulation{
		GasWanted: message.GasWanted,
		GasUsed:   message.GasUsed,
	}
	if message.Fees != nil {
		result.Fees = make([]*sdkutilities.Coin, len(message.Fees))
		for i, val := range message.Fees {
			result.Fees[i] = &sdkutilities.Coin{
				Denom:  val.Denom,
				Amount: val.Amount,
			}
		}
	}
	return result
}

// NewProtoStakingParamsRequest builds the gRPC request type from the payload
// of the "stakingParams" endpoint of the "sdk-utilities" service.
func NewProtoStakingParamsRequest(payload *sdkutilities.StakingParamsPayload) *sdk_utilitiespb.StakingParamsRequest {
	message := &sdk_utilitiespb.StakingParamsRequest{
		ChainName: payload.ChainName,
	}
	if payload.Port != nil {
		message.Port = int32(*payload.Port)
	}
	return message
}

// NewStakingParamsResult builds the result type of the "stakingParams"
// endpoint of the "sdk-utilities" service from the gRPC response type.
func NewStakingParamsResult(message *sdk_utilitiespb.StakingParamsResponse) *sdkutilities.StakingParams2 {
	result := &sdkutilities.StakingParams2{
		StakingParams: message.StakingParams,
	}
	return result
}

// NewProtoStakingPoolRequest builds the gRPC request type from the payload of
// the "stakingPool" endpoint of the "sdk-utilities" service.
func NewProtoStakingPoolRequest(payload *sdkutilities.StakingPoolPayload) *sdk_utilitiespb.StakingPoolRequest {
	message := &sdk_utilitiespb.StakingPoolRequest{
		ChainName: payload.ChainName,
	}
	if payload.Port != nil {
		message.Port = int32(*payload.Port)
	}
	return message
}

// NewStakingPoolResult builds the result type of the "stakingPool" endpoint of
// the "sdk-utilities" service from the gRPC response type.
func NewStakingPoolResult(message *sdk_utilitiespb.StakingPoolResponse) *sdkutilities.StakingPool2 {
	result := &sdkutilities.StakingPool2{
		StakingPool: message.StakingPool,
	}
	return result
}

// NewProtoEmoneyInflationRequest builds the gRPC request type from the payload
// of the "emoneyInflation" endpoint of the "sdk-utilities" service.
func NewProtoEmoneyInflationRequest(payload *sdkutilities.EmoneyInflationPayload) *sdk_utilitiespb.EmoneyInflationRequest {
	message := &sdk_utilitiespb.EmoneyInflationRequest{
		ChainName: payload.ChainName,
	}
	if payload.Port != nil {
		message.Port = int32(*payload.Port)
	}
	return message
}

// NewEmoneyInflationResult builds the result type of the "emoneyInflation"
// endpoint of the "sdk-utilities" service from the gRPC response type.
func NewEmoneyInflationResult(message *sdk_utilitiespb.EmoneyInflationResponse) *sdkutilities.EmoneyInflation2 {
	result := &sdkutilities.EmoneyInflation2{}
	if message.State != nil {
		result.State = protobufSdkUtilitiespbEmoneyStateToSdkutilitiesEmoneyState(message.State)
	}
	return result
}

// NewProtoBudgetParamsRequest builds the gRPC request type from the payload of
// the "budgetParams" endpoint of the "sdk-utilities" service.
func NewProtoBudgetParamsRequest(payload *sdkutilities.BudgetParamsPayload) *sdk_utilitiespb.BudgetParamsRequest {
	message := &sdk_utilitiespb.BudgetParamsRequest{
		ChainName: payload.ChainName,
	}
	if payload.Port != nil {
		message.Port = int32(*payload.Port)
	}
	return message
}

// NewBudgetParamsResult builds the result type of the "budgetParams" endpoint
// of the "sdk-utilities" service from the gRPC response type.
func NewBudgetParamsResult(message *sdk_utilitiespb.BudgetParamsResponse) *sdkutilities.BudgetParams2 {
	result := &sdkutilities.BudgetParams2{
		BudgetParams: message.BudgetParams,
	}
	return result
}

// NewProtoDistributionParamsRequest builds the gRPC request type from the
// payload of the "distributionParams" endpoint of the "sdk-utilities" service.
func NewProtoDistributionParamsRequest(payload *sdkutilities.DistributionParamsPayload) *sdk_utilitiespb.DistributionParamsRequest {
	message := &sdk_utilitiespb.DistributionParamsRequest{
		ChainName: payload.ChainName,
	}
	if payload.Port != nil {
		message.Port = int32(*payload.Port)
	}
	return message
}

// NewDistributionParamsResult builds the result type of the
// "distributionParams" endpoint of the "sdk-utilities" service from the gRPC
// response type.
func NewDistributionParamsResult(message *sdk_utilitiespb.DistributionParamsResponse) *sdkutilities.DistributionParams2 {
	result := &sdkutilities.DistributionParams2{
		DistributionParams: message.DistributionParams,
	}
	return result
}

// NewProtoOsmoPoolsRequest builds the gRPC request type from the payload of
// the "osmoPools" endpoint of the "sdk-utilities" service.
func NewProtoOsmoPoolsRequest(payload *sdkutilities.OsmoPoolsPayload) *sdk_utilitiespb.OsmoPoolsRequest {
	message := &sdk_utilitiespb.OsmoPoolsRequest{
		ChainName: payload.ChainName,
	}
	if payload.Port != nil {
		message.Port = int32(*payload.Port)
	}
	return message
}

// NewOsmoPoolsResult builds the result type of the "osmoPools" endpoint of the
// "sdk-utilities" service from the gRPC response type.
func NewOsmoPoolsResult(message *sdk_utilitiespb.OsmoPoolsResponse) *sdkutilities.OsmoPools2 {
	result := &sdkutilities.OsmoPools2{
		OsmoPools: message.OsmoPools,
	}
	return result
}

// NewProtoCrescentPoolsRequest builds the gRPC request type from the payload
// of the "crescentPools" endpoint of the "sdk-utilities" service.
func NewProtoCrescentPoolsRequest(payload *sdkutilities.CrescentPoolsPayload) *sdk_utilitiespb.CrescentPoolsRequest {
	message := &sdk_utilitiespb.CrescentPoolsRequest{
		ChainName: payload.ChainName,
	}
	if payload.Port != nil {
		message.Port = int32(*payload.Port)
	}
	return message
}

// NewCrescentPoolsResult builds the result type of the "crescentPools"
// endpoint of the "sdk-utilities" service from the gRPC response type.
func NewCrescentPoolsResult(message *sdk_utilitiespb.CrescentPoolsResponse) *sdkutilities.CrescentPools2 {
	result := &sdkutilities.CrescentPools2{
		CrescentPools: message.CrescentPools,
	}
	return result
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

// ValidatePagination runs the validations defined on Pagination.
func ValidatePagination(message *sdk_utilitiespb.Pagination) (err error) {

	return
}

// ValidateSupplyDenomResponse runs the validations defined on
// SupplyDenomResponse.
func ValidateSupplyDenomResponse(message *sdk_utilitiespb.SupplyDenomResponse) (err error) {
	if message.Coins == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("coins", "message"))
	}
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

// ValidateBlockResponse runs the validations defined on BlockResponse.
func ValidateBlockResponse(message *sdk_utilitiespb.BlockResponse) (err error) {
	if message.Block == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("block", "message"))
	}
	return
}

// ValidateLiquidityParamsResponse runs the validations defined on
// LiquidityParamsResponse.
func ValidateLiquidityParamsResponse(message *sdk_utilitiespb.LiquidityParamsResponse) (err error) {
	if message.LiquidityParams == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("liquidityParams", "message"))
	}
	return
}

// ValidateLiquidityPoolsResponse runs the validations defined on
// LiquidityPoolsResponse.
func ValidateLiquidityPoolsResponse(message *sdk_utilitiespb.LiquidityPoolsResponse) (err error) {
	if message.LiquidityPools == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("liquidityPools", "message"))
	}
	return
}

// ValidateMintInflationResponse runs the validations defined on
// MintInflationResponse.
func ValidateMintInflationResponse(message *sdk_utilitiespb.MintInflationResponse) (err error) {
	if message.MintInflation == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("mintInflation", "message"))
	}
	return
}

// ValidateMintParamsResponse runs the validations defined on
// MintParamsResponse.
func ValidateMintParamsResponse(message *sdk_utilitiespb.MintParamsResponse) (err error) {
	if message.MintParams == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("mintParams", "message"))
	}
	return
}

// ValidateMintAnnualProvisionResponse runs the validations defined on
// MintAnnualProvisionResponse.
func ValidateMintAnnualProvisionResponse(message *sdk_utilitiespb.MintAnnualProvisionResponse) (err error) {
	if message.MintAnnualProvision == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("mintAnnualProvision", "message"))
	}
	return
}

// ValidateMintEpochProvisionsResponse runs the validations defined on
// MintEpochProvisionsResponse.
func ValidateMintEpochProvisionsResponse(message *sdk_utilitiespb.MintEpochProvisionsResponse) (err error) {
	if message.MintEpochProvisions == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("mintEpochProvisions", "message"))
	}
	return
}

// ValidateStakingParamsResponse runs the validations defined on
// StakingParamsResponse.
func ValidateStakingParamsResponse(message *sdk_utilitiespb.StakingParamsResponse) (err error) {
	if message.StakingParams == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("stakingParams", "message"))
	}
	return
}

// ValidateStakingPoolResponse runs the validations defined on
// StakingPoolResponse.
func ValidateStakingPoolResponse(message *sdk_utilitiespb.StakingPoolResponse) (err error) {
	if message.StakingPool == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("stakingPool", "message"))
	}
	return
}

// ValidateEmoneyInflationResponse runs the validations defined on
// EmoneyInflationResponse.
func ValidateEmoneyInflationResponse(message *sdk_utilitiespb.EmoneyInflationResponse) (err error) {
	if message.State == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("state", "message"))
	}
	if message.State != nil {
		if err2 := ValidateEmoneyState(message.State); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	return
}

// ValidateEmoneyState runs the validations defined on EmoneyState.
func ValidateEmoneyState(message *sdk_utilitiespb.EmoneyState) (err error) {
	if message.Assets == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("assets", "message"))
	}
	return
}

// ValidateEmoneyAsset runs the validations defined on EmoneyAsset.
func ValidateEmoneyAsset(message *sdk_utilitiespb.EmoneyAsset) (err error) {

	return
}

// ValidateBudgetParamsResponse runs the validations defined on
// BudgetParamsResponse.
func ValidateBudgetParamsResponse(message *sdk_utilitiespb.BudgetParamsResponse) (err error) {
	if message.BudgetParams == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("budgetParams", "message"))
	}
	return
}

// ValidateDistributionParamsResponse runs the validations defined on
// DistributionParamsResponse.
func ValidateDistributionParamsResponse(message *sdk_utilitiespb.DistributionParamsResponse) (err error) {
	if message.DistributionParams == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("distributionParams", "message"))
	}
	return
}

// ValidateOsmoPoolsResponse runs the validations defined on OsmoPoolsResponse.
func ValidateOsmoPoolsResponse(message *sdk_utilitiespb.OsmoPoolsResponse) (err error) {
	if message.OsmoPools == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("osmoPools", "message"))
	}
	return
}

// ValidateCrescentPoolsResponse runs the validations defined on
// CrescentPoolsResponse.
func ValidateCrescentPoolsResponse(message *sdk_utilitiespb.CrescentPoolsResponse) (err error) {
	if message.CrescentPools == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("crescentPools", "message"))
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
