// Code generated by goa v3.7.5, DO NOT EDIT.
//
// sdk-utilities service
//
// Command:
// $ goa gen github.com/emerishq/sdk-service-meta

package sdkutilities

import (
	"context"
)

// sdk-utilities performs Cosmos SDK-related operations
type Service interface {
	// AccountNumbers implements accountNumbers.
	AccountNumbers(context.Context, *AccountNumbersPayload) (res *AccountNumbers2, err error)
	// Supply implements supply.
	Supply(context.Context, *SupplyPayload) (res *Supply2, err error)
	// SupplyDenom implements supplyDenom.
	SupplyDenom(context.Context, *SupplyDenomPayload) (res *Supply2, err error)
	// QueryTx implements queryTx.
	QueryTx(context.Context, *QueryTxPayload) (res []byte, err error)
	// BroadcastTx implements broadcastTx.
	BroadcastTx(context.Context, *BroadcastTxPayload) (res *TransactionResult, err error)
	// TxMetadata implements txMetadata.
	TxMetadata(context.Context, *TxMetadataPayload) (res *TxMessagesMetadata, err error)
	// Block implements block.
	Block(context.Context, *BlockPayload) (res *BlockData, err error)
	// LiquidityParams implements liquidityParams.
	LiquidityParams(context.Context, *LiquidityParamsPayload) (res *LiquidityParams2, err error)
	// LiquidityPools implements liquidityPools.
	LiquidityPools(context.Context, *LiquidityPoolsPayload) (res *LiquidityPools2, err error)
	// MintInflation implements mintInflation.
	MintInflation(context.Context, *MintInflationPayload) (res *MintInflation2, err error)
	// MintParams implements mintParams.
	MintParams(context.Context, *MintParamsPayload) (res *MintParams2, err error)
	// MintAnnualProvision implements mintAnnualProvision.
	MintAnnualProvision(context.Context, *MintAnnualProvisionPayload) (res *MintAnnualProvision2, err error)
	// MintEpochProvisions implements mintEpochProvisions.
	MintEpochProvisions(context.Context, *MintEpochProvisionsPayload) (res *MintEpochProvisions2, err error)
	// DelegatorRewards implements delegatorRewards.
	DelegatorRewards(context.Context, *DelegatorRewardsPayload) (res *DelegatorRewards2, err error)
	// EstimateFees implements estimateFees.
	EstimateFees(context.Context, *EstimateFeesPayload) (res *Simulation, err error)
	// StakingParams implements stakingParams.
	StakingParams(context.Context, *StakingParamsPayload) (res *StakingParams2, err error)
	// StakingPool implements stakingPool.
	StakingPool(context.Context, *StakingPoolPayload) (res *StakingPool2, err error)
	// EmoneyInflation implements emoneyInflation.
	EmoneyInflation(context.Context, *EmoneyInflationPayload) (res *EmoneyInflation2, err error)
	// BudgetParams implements budgetParams.
	BudgetParams(context.Context, *BudgetParamsPayload) (res *BudgetParams2, err error)
	// DistributionParams implements distributionParams.
	DistributionParams(context.Context, *DistributionParamsPayload) (res *DistributionParams2, err error)
	// OsmoPools implements osmoPools.
	OsmoPools(context.Context, *OsmoPoolsPayload) (res *OsmoPools2, err error)
	// CrescentPools implements crescentPools.
	CrescentPools(context.Context, *CrescentPoolsPayload) (res *CrescentPools2, err error)
}

// ServiceName is the name of the service as defined in the design. This is the
// same value that is set in the endpoint request contexts under the ServiceKey
// key.
const ServiceName = "sdk-utilities"

// MethodNames lists the service method names as defined in the design. These
// are the same values that are set in the endpoint request contexts under the
// MethodKey key.
var MethodNames = [22]string{"accountNumbers", "supply", "supplyDenom", "queryTx", "broadcastTx", "txMetadata", "block", "liquidityParams", "liquidityPools", "mintInflation", "mintParams", "mintAnnualProvision", "mintEpochProvisions", "delegatorRewards", "estimateFees", "stakingParams", "stakingPool", "emoneyInflation", "budgetParams", "distributionParams", "osmoPools", "crescentPools"}

// AccountNumbers2 is the result type of the sdk-utilities service
// accountNumbers method.
type AccountNumbers2 struct {
	AccountNumber  int64
	SequenceNumber int64
	Bech32Address  string
}

// AccountNumbersPayload is the payload type of the sdk-utilities service
// accountNumbers method.
type AccountNumbersPayload struct {
	// Chain to get data from
	ChainName string
	// gRPC port for selected chain, defaults to 9090
	Port *int
	// bech32-encoded prefix of the account
	Bech32Prefix *string
	// Hex-encoded address, without bech32 hrp
	AddresHex *string
}

// BlockData is the result type of the sdk-utilities service block method.
type BlockData struct {
	Height int64
	Block  []byte
}

// BlockPayload is the payload type of the sdk-utilities service block method.
type BlockPayload struct {
	// Chain to get data from
	ChainName string
	// gRPC port for selected chain, defaults to 9090
	Port *int
	// Height of the block to query
	Height int64
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

// BudgetParams2 is the result type of the sdk-utilities service budgetParams
// method.
type BudgetParams2 struct {
	BudgetParams []byte
}

// BudgetParamsPayload is the payload type of the sdk-utilities service
// budgetParams method.
type BudgetParamsPayload struct {
	// Chain to get data from
	ChainName string
	// gRPC port for selected chain, defaults to 9090
	Port *int
}

// SDK service representation of a Cosmos SDK types.Coin
type Coin struct {
	Denom  string
	Amount string
}

// CrescentPools2 is the result type of the sdk-utilities service crescentPools
// method.
type CrescentPools2 struct {
	CrescentPools []byte
}

// CrescentPoolsPayload is the payload type of the sdk-utilities service
// crescentPools method.
type CrescentPoolsPayload struct {
	// Chain to get data from
	ChainName string
	// gRPC port for selected chain, defaults to 9090
	Port *int
}

// Delegation reward as defined in the Cosmos SDK
type DelegationDelegatorReward struct {
	ValidatorAddress string
	Rewards          []*Coin
}

// DelegatorRewards2 is the result type of the sdk-utilities service
// delegatorRewards method.
type DelegatorRewards2 struct {
	Rewards []*DelegationDelegatorReward
	Total   []*Coin
}

// DelegatorRewardsPayload is the payload type of the sdk-utilities service
// delegatorRewards method.
type DelegatorRewardsPayload struct {
	// Chain to get data from
	ChainName string
	// gRPC port for selected chain, defaults to 9090
	Port *int
	// bech32-encoded prefix of the account
	Bech32Prefix *string
	// Hex-encoded address, without bech32 hrp
	AddresHex *string
}

// DistributionParams2 is the result type of the sdk-utilities service
// distributionParams method.
type DistributionParams2 struct {
	DistributionParams []byte
}

// DistributionParamsPayload is the payload type of the sdk-utilities service
// distributionParams method.
type DistributionParamsPayload struct {
	// Chain to get data from
	ChainName string
	// gRPC port for selected chain, defaults to 9090
	Port *int
}

// e-money specific inflation parameters (asset)
type EmoneyAsset struct {
	Denom     string
	Inflation string
	Accum     string
}

// EmoneyInflation2 is the result type of the sdk-utilities service
// emoneyInflation method.
type EmoneyInflation2 struct {
	State *EmoneyState
}

// EmoneyInflationPayload is the payload type of the sdk-utilities service
// emoneyInflation method.
type EmoneyInflationPayload struct {
	// Chain to get data from
	ChainName string
	// gRPC port for selected chain, defaults to 9090
	Port *int
}

// e-money specific inflation parameters (state)
type EmoneyState struct {
	LastApplied       string
	LastAppliedHeight string
	Assets            []*EmoneyAsset
}

// EstimateFeesPayload is the payload type of the sdk-utilities service
// estimateFees method.
type EstimateFeesPayload struct {
	// Chain to get data from
	ChainName string
	// gRPC port for selected chain, defaults to 9090
	Port *int
	// Transaction bytes
	TxBytes []byte
}

// The plain type associated with ibc-go types.Height struct
type IBCHeight struct {
	RevisionNumber *uint64
	RevisionHeight *uint64
}

// Metadata related to an IBC Transfer
type IBCTransferMetadata struct {
	SourcePort       *string
	SourceChannel    *string
	Token            *Coin
	Sender           *string
	Receiver         *string
	TimeoutHeight    *IBCHeight
	TiemoutTimestamp *uint64
}

// LiquidityParams2 is the result type of the sdk-utilities service
// liquidityParams method.
type LiquidityParams2 struct {
	LiquidityParams []byte
}

// LiquidityParamsPayload is the payload type of the sdk-utilities service
// liquidityParams method.
type LiquidityParamsPayload struct {
	// Chain to get data from
	ChainName string
	// gRPC port for selected chain, defaults to 9090
	Port *int
}

// LiquidityPools2 is the result type of the sdk-utilities service
// liquidityPools method.
type LiquidityPools2 struct {
	LiquidityPools []byte
}

// LiquidityPoolsPayload is the payload type of the sdk-utilities service
// liquidityPools method.
type LiquidityPoolsPayload struct {
	// Chain to get data from
	ChainName string
	// gRPC port for selected chain, defaults to 9090
	Port *int
}

// MintAnnualProvision2 is the result type of the sdk-utilities service
// mintAnnualProvision method.
type MintAnnualProvision2 struct {
	MintAnnualProvision []byte
}

// MintAnnualProvisionPayload is the payload type of the sdk-utilities service
// mintAnnualProvision method.
type MintAnnualProvisionPayload struct {
	// Chain to get data from
	ChainName string
	// gRPC port for selected chain, defaults to 9090
	Port *int
}

// MintEpochProvisions2 is the result type of the sdk-utilities service
// mintEpochProvisions method.
type MintEpochProvisions2 struct {
	MintEpochProvisions []byte
}

// MintEpochProvisionsPayload is the payload type of the sdk-utilities service
// mintEpochProvisions method.
type MintEpochProvisionsPayload struct {
	// Chain to get data from
	ChainName string
	// gRPC port for selected chain, defaults to 9090
	Port *int
}

// MintInflation2 is the result type of the sdk-utilities service mintInflation
// method.
type MintInflation2 struct {
	MintInflation []byte
}

// MintInflationPayload is the payload type of the sdk-utilities service
// mintInflation method.
type MintInflationPayload struct {
	// Chain to get data from
	ChainName string
	// gRPC port for selected chain, defaults to 9090
	Port *int
}

// MintParams2 is the result type of the sdk-utilities service mintParams
// method.
type MintParams2 struct {
	MintParams []byte
}

// MintParamsPayload is the payload type of the sdk-utilities service
// mintParams method.
type MintParamsPayload struct {
	// Chain to get data from
	ChainName string
	// gRPC port for selected chain, defaults to 9090
	Port *int
}

// Metadata related to some message contained in a transaction
type MsgMetadata struct {
	MsgType             string
	IbcTransferMetadata *IBCTransferMetadata
}

// OsmoPools2 is the result type of the sdk-utilities service osmoPools method.
type OsmoPools2 struct {
	OsmoPools []byte
}

// OsmoPoolsPayload is the payload type of the sdk-utilities service osmoPools
// method.
type OsmoPoolsPayload struct {
	// Chain to get data from
	ChainName string
	// gRPC port for selected chain, defaults to 9090
	Port *int
}

// Pagination used in the Cosmos SDK
type Pagination struct {
	NextKey *string
	Total   *string
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

// Simulation is the result type of the sdk-utilities service estimateFees
// method.
type Simulation struct {
	GasWanted uint64
	GasUsed   uint64
	Fees      []*Coin
}

// StakingParams2 is the result type of the sdk-utilities service stakingParams
// method.
type StakingParams2 struct {
	StakingParams []byte
}

// StakingParamsPayload is the payload type of the sdk-utilities service
// stakingParams method.
type StakingParamsPayload struct {
	// Chain to get data from
	ChainName string
	// gRPC port for selected chain, defaults to 9090
	Port *int
}

// StakingPool2 is the result type of the sdk-utilities service stakingPool
// method.
type StakingPool2 struct {
	StakingPool []byte
}

// StakingPoolPayload is the payload type of the sdk-utilities service
// stakingPool method.
type StakingPoolPayload struct {
	// Chain to get data from
	ChainName string
	// gRPC port for selected chain, defaults to 9090
	Port *int
}

// Supply2 is the result type of the sdk-utilities service supply method.
type Supply2 struct {
	Coins      []*Coin
	Pagination *Pagination
}

// SupplyDenomPayload is the payload type of the sdk-utilities service
// supplyDenom method.
type SupplyDenomPayload struct {
	// Chain to get data from
	ChainName string
	// gRPC port for selected chain, defaults to 9090
	Port *int
	// Denom name
	Denom *string
}

// SupplyPayload is the payload type of the sdk-utilities service supply method.
type SupplyPayload struct {
	// Chain to get data from
	ChainName string
	// gRPC port for selected chain, defaults to 9090
	Port *int
	// pagination key
	PaginationKey *string
}

// TransactionResult is the result type of the sdk-utilities service
// broadcastTx method.
type TransactionResult struct {
	Hash string
}

// TxMessagesMetadata is the result type of the sdk-utilities service
// txMetadata method.
type TxMessagesMetadata struct {
	MessagesMetadata []*MsgMetadata
}

// TxMetadataPayload is the payload type of the sdk-utilities service
// txMetadata method.
type TxMetadataPayload struct {
	// Transaction bytes
	TxBytes []byte
}
