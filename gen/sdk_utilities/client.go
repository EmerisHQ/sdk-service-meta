// Code generated by goa v3.6.1, DO NOT EDIT.
//
// sdk-utilities client
//
// Command:
// $ goa gen github.com/emerishq/sdk-service-meta

package sdkutilities

import (
	"context"

	goa "goa.design/goa/v3/pkg"
)

// Client is the "sdk-utilities" service client.
type Client struct {
	AccountNumbersEndpoint      goa.Endpoint
	SupplyEndpoint              goa.Endpoint
	QueryTxEndpoint             goa.Endpoint
	BroadcastTxEndpoint         goa.Endpoint
	TxMetadataEndpoint          goa.Endpoint
	BlockEndpoint               goa.Endpoint
	LiquidityParamsEndpoint     goa.Endpoint
	LiquidityPoolsEndpoint      goa.Endpoint
	MintInflationEndpoint       goa.Endpoint
	MintParamsEndpoint          goa.Endpoint
	MintAnnualProvisionEndpoint goa.Endpoint
	DelegatorRewardsEndpoint    goa.Endpoint
	EstimateFeesEndpoint        goa.Endpoint
	StakingParamsEndpoint       goa.Endpoint
}

// NewClient initializes a "sdk-utilities" service client given the endpoints.
func NewClient(accountNumbers, supply, queryTx, broadcastTx, txMetadata, block, liquidityParams, liquidityPools, mintInflation, mintParams, mintAnnualProvision, delegatorRewards, estimateFees, stakingParams goa.Endpoint) *Client {
	return &Client{
		AccountNumbersEndpoint:      accountNumbers,
		SupplyEndpoint:              supply,
		QueryTxEndpoint:             queryTx,
		BroadcastTxEndpoint:         broadcastTx,
		TxMetadataEndpoint:          txMetadata,
		BlockEndpoint:               block,
		LiquidityParamsEndpoint:     liquidityParams,
		LiquidityPoolsEndpoint:      liquidityPools,
		MintInflationEndpoint:       mintInflation,
		MintParamsEndpoint:          mintParams,
		MintAnnualProvisionEndpoint: mintAnnualProvision,
		DelegatorRewardsEndpoint:    delegatorRewards,
		EstimateFeesEndpoint:        estimateFees,
		StakingParamsEndpoint:       stakingParams,
	}
}

// AccountNumbers calls the "accountNumbers" endpoint of the "sdk-utilities"
// service.
func (c *Client) AccountNumbers(ctx context.Context, p *AccountNumbersPayload) (res *AccountNumbers2, err error) {
	var ires interface{}
	ires, err = c.AccountNumbersEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(*AccountNumbers2), nil
}

// Supply calls the "supply" endpoint of the "sdk-utilities" service.
func (c *Client) Supply(ctx context.Context, p *SupplyPayload) (res *Supply2, err error) {
	var ires interface{}
	ires, err = c.SupplyEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(*Supply2), nil
}

// QueryTx calls the "queryTx" endpoint of the "sdk-utilities" service.
func (c *Client) QueryTx(ctx context.Context, p *QueryTxPayload) (res []byte, err error) {
	var ires interface{}
	ires, err = c.QueryTxEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.([]byte), nil
}

// BroadcastTx calls the "broadcastTx" endpoint of the "sdk-utilities" service.
func (c *Client) BroadcastTx(ctx context.Context, p *BroadcastTxPayload) (res *TransactionResult, err error) {
	var ires interface{}
	ires, err = c.BroadcastTxEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(*TransactionResult), nil
}

// TxMetadata calls the "txMetadata" endpoint of the "sdk-utilities" service.
func (c *Client) TxMetadata(ctx context.Context, p *TxMetadataPayload) (res *TxMessagesMetadata, err error) {
	var ires interface{}
	ires, err = c.TxMetadataEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(*TxMessagesMetadata), nil
}

// Block calls the "block" endpoint of the "sdk-utilities" service.
func (c *Client) Block(ctx context.Context, p *BlockPayload) (res *BlockData, err error) {
	var ires interface{}
	ires, err = c.BlockEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(*BlockData), nil
}

// LiquidityParams calls the "liquidityParams" endpoint of the "sdk-utilities"
// service.
func (c *Client) LiquidityParams(ctx context.Context, p *LiquidityParamsPayload) (res *LiquidityParams2, err error) {
	var ires interface{}
	ires, err = c.LiquidityParamsEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(*LiquidityParams2), nil
}

// LiquidityPools calls the "liquidityPools" endpoint of the "sdk-utilities"
// service.
func (c *Client) LiquidityPools(ctx context.Context, p *LiquidityPoolsPayload) (res *LiquidityPools2, err error) {
	var ires interface{}
	ires, err = c.LiquidityPoolsEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(*LiquidityPools2), nil
}

// MintInflation calls the "mintInflation" endpoint of the "sdk-utilities"
// service.
func (c *Client) MintInflation(ctx context.Context, p *MintInflationPayload) (res *MintInflation2, err error) {
	var ires interface{}
	ires, err = c.MintInflationEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(*MintInflation2), nil
}

// MintParams calls the "mintParams" endpoint of the "sdk-utilities" service.
func (c *Client) MintParams(ctx context.Context, p *MintParamsPayload) (res *MintParams2, err error) {
	var ires interface{}
	ires, err = c.MintParamsEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(*MintParams2), nil
}

// MintAnnualProvision calls the "mintAnnualProvision" endpoint of the
// "sdk-utilities" service.
func (c *Client) MintAnnualProvision(ctx context.Context, p *MintAnnualProvisionPayload) (res *MintAnnualProvision2, err error) {
	var ires interface{}
	ires, err = c.MintAnnualProvisionEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(*MintAnnualProvision2), nil
}

// DelegatorRewards calls the "delegatorRewards" endpoint of the
// "sdk-utilities" service.
func (c *Client) DelegatorRewards(ctx context.Context, p *DelegatorRewardsPayload) (res *DelegatorRewards2, err error) {
	var ires interface{}
	ires, err = c.DelegatorRewardsEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(*DelegatorRewards2), nil
}

// EstimateFees calls the "estimateFees" endpoint of the "sdk-utilities"
// service.
func (c *Client) EstimateFees(ctx context.Context, p *EstimateFeesPayload) (res *Simulation, err error) {
	var ires interface{}
	ires, err = c.EstimateFeesEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(*Simulation), nil
}

// StakingParams calls the "stakingParams" endpoint of the "sdk-utilities"
// service.
func (c *Client) StakingParams(ctx context.Context, p *StakingParamsPayload) (res *StakingParams2, err error) {
	var ires interface{}
	ires, err = c.StakingParamsEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(*StakingParams2), nil
}
