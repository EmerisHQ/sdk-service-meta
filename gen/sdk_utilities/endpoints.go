// Code generated by goa v3.7.3, DO NOT EDIT.
//
// sdk-utilities endpoints
//
// Command:
// $ goa gen github.com/emerishq/sdk-service-meta

package sdkutilities

import (
	"context"

	goa "goa.design/goa/v3/pkg"
)

// Endpoints wraps the "sdk-utilities" service endpoints.
type Endpoints struct {
	AccountNumbers      goa.Endpoint
	Supply              goa.Endpoint
	SupplyDenom         goa.Endpoint
	QueryTx             goa.Endpoint
	BroadcastTx         goa.Endpoint
	TxMetadata          goa.Endpoint
	Block               goa.Endpoint
	LiquidityParams     goa.Endpoint
	LiquidityPools      goa.Endpoint
	MintInflation       goa.Endpoint
	MintParams          goa.Endpoint
	MintAnnualProvision goa.Endpoint
	MintEpochProvisions goa.Endpoint
	DelegatorRewards    goa.Endpoint
	EstimateFees        goa.Endpoint
	StakingParams       goa.Endpoint
	StakingPool         goa.Endpoint
	EmoneyInflation     goa.Endpoint
	BudgetParams        goa.Endpoint
	DistributionParams  goa.Endpoint
	CrescentPools       goa.Endpoint
}

// NewEndpoints wraps the methods of the "sdk-utilities" service with endpoints.
func NewEndpoints(s Service) *Endpoints {
	return &Endpoints{
		AccountNumbers:      NewAccountNumbersEndpoint(s),
		Supply:              NewSupplyEndpoint(s),
		SupplyDenom:         NewSupplyDenomEndpoint(s),
		QueryTx:             NewQueryTxEndpoint(s),
		BroadcastTx:         NewBroadcastTxEndpoint(s),
		TxMetadata:          NewTxMetadataEndpoint(s),
		Block:               NewBlockEndpoint(s),
		LiquidityParams:     NewLiquidityParamsEndpoint(s),
		LiquidityPools:      NewLiquidityPoolsEndpoint(s),
		MintInflation:       NewMintInflationEndpoint(s),
		MintParams:          NewMintParamsEndpoint(s),
		MintAnnualProvision: NewMintAnnualProvisionEndpoint(s),
		MintEpochProvisions: NewMintEpochProvisionsEndpoint(s),
		DelegatorRewards:    NewDelegatorRewardsEndpoint(s),
		EstimateFees:        NewEstimateFeesEndpoint(s),
		StakingParams:       NewStakingParamsEndpoint(s),
		StakingPool:         NewStakingPoolEndpoint(s),
		EmoneyInflation:     NewEmoneyInflationEndpoint(s),
		BudgetParams:        NewBudgetParamsEndpoint(s),
		DistributionParams:  NewDistributionParamsEndpoint(s),
		CrescentPools:       NewCrescentPoolsEndpoint(s),
	}
}

// Use applies the given middleware to all the "sdk-utilities" service
// endpoints.
func (e *Endpoints) Use(m func(goa.Endpoint) goa.Endpoint) {
	e.AccountNumbers = m(e.AccountNumbers)
	e.Supply = m(e.Supply)
	e.SupplyDenom = m(e.SupplyDenom)
	e.QueryTx = m(e.QueryTx)
	e.BroadcastTx = m(e.BroadcastTx)
	e.TxMetadata = m(e.TxMetadata)
	e.Block = m(e.Block)
	e.LiquidityParams = m(e.LiquidityParams)
	e.LiquidityPools = m(e.LiquidityPools)
	e.MintInflation = m(e.MintInflation)
	e.MintParams = m(e.MintParams)
	e.MintAnnualProvision = m(e.MintAnnualProvision)
	e.MintEpochProvisions = m(e.MintEpochProvisions)
	e.DelegatorRewards = m(e.DelegatorRewards)
	e.EstimateFees = m(e.EstimateFees)
	e.StakingParams = m(e.StakingParams)
	e.StakingPool = m(e.StakingPool)
	e.EmoneyInflation = m(e.EmoneyInflation)
	e.BudgetParams = m(e.BudgetParams)
	e.DistributionParams = m(e.DistributionParams)
	e.CrescentPools = m(e.CrescentPools)
}

// NewAccountNumbersEndpoint returns an endpoint function that calls the method
// "accountNumbers" of service "sdk-utilities".
func NewAccountNumbersEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*AccountNumbersPayload)
		return s.AccountNumbers(ctx, p)
	}
}

// NewSupplyEndpoint returns an endpoint function that calls the method
// "supply" of service "sdk-utilities".
func NewSupplyEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*SupplyPayload)
		return s.Supply(ctx, p)
	}
}

// NewSupplyDenomEndpoint returns an endpoint function that calls the method
// "supplyDenom" of service "sdk-utilities".
func NewSupplyDenomEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*SupplyDenomPayload)
		return s.SupplyDenom(ctx, p)
	}
}

// NewQueryTxEndpoint returns an endpoint function that calls the method
// "queryTx" of service "sdk-utilities".
func NewQueryTxEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*QueryTxPayload)
		return s.QueryTx(ctx, p)
	}
}

// NewBroadcastTxEndpoint returns an endpoint function that calls the method
// "broadcastTx" of service "sdk-utilities".
func NewBroadcastTxEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*BroadcastTxPayload)
		return s.BroadcastTx(ctx, p)
	}
}

// NewTxMetadataEndpoint returns an endpoint function that calls the method
// "txMetadata" of service "sdk-utilities".
func NewTxMetadataEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*TxMetadataPayload)
		return s.TxMetadata(ctx, p)
	}
}

// NewBlockEndpoint returns an endpoint function that calls the method "block"
// of service "sdk-utilities".
func NewBlockEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*BlockPayload)
		return s.Block(ctx, p)
	}
}

// NewLiquidityParamsEndpoint returns an endpoint function that calls the
// method "liquidityParams" of service "sdk-utilities".
func NewLiquidityParamsEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*LiquidityParamsPayload)
		return s.LiquidityParams(ctx, p)
	}
}

// NewLiquidityPoolsEndpoint returns an endpoint function that calls the method
// "liquidityPools" of service "sdk-utilities".
func NewLiquidityPoolsEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*LiquidityPoolsPayload)
		return s.LiquidityPools(ctx, p)
	}
}

// NewMintInflationEndpoint returns an endpoint function that calls the method
// "mintInflation" of service "sdk-utilities".
func NewMintInflationEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*MintInflationPayload)
		return s.MintInflation(ctx, p)
	}
}

// NewMintParamsEndpoint returns an endpoint function that calls the method
// "mintParams" of service "sdk-utilities".
func NewMintParamsEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*MintParamsPayload)
		return s.MintParams(ctx, p)
	}
}

// NewMintAnnualProvisionEndpoint returns an endpoint function that calls the
// method "mintAnnualProvision" of service "sdk-utilities".
func NewMintAnnualProvisionEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*MintAnnualProvisionPayload)
		return s.MintAnnualProvision(ctx, p)
	}
}

// NewMintEpochProvisionsEndpoint returns an endpoint function that calls the
// method "mintEpochProvisions" of service "sdk-utilities".
func NewMintEpochProvisionsEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*MintEpochProvisionsPayload)
		return s.MintEpochProvisions(ctx, p)
	}
}

// NewDelegatorRewardsEndpoint returns an endpoint function that calls the
// method "delegatorRewards" of service "sdk-utilities".
func NewDelegatorRewardsEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*DelegatorRewardsPayload)
		return s.DelegatorRewards(ctx, p)
	}
}

// NewEstimateFeesEndpoint returns an endpoint function that calls the method
// "estimateFees" of service "sdk-utilities".
func NewEstimateFeesEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*EstimateFeesPayload)
		return s.EstimateFees(ctx, p)
	}
}

// NewStakingParamsEndpoint returns an endpoint function that calls the method
// "stakingParams" of service "sdk-utilities".
func NewStakingParamsEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*StakingParamsPayload)
		return s.StakingParams(ctx, p)
	}
}

// NewStakingPoolEndpoint returns an endpoint function that calls the method
// "stakingPool" of service "sdk-utilities".
func NewStakingPoolEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*StakingPoolPayload)
		return s.StakingPool(ctx, p)
	}
}

// NewEmoneyInflationEndpoint returns an endpoint function that calls the
// method "emoneyInflation" of service "sdk-utilities".
func NewEmoneyInflationEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*EmoneyInflationPayload)
		return s.EmoneyInflation(ctx, p)
	}
}

// NewBudgetParamsEndpoint returns an endpoint function that calls the method
// "budgetParams" of service "sdk-utilities".
func NewBudgetParamsEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*BudgetParamsPayload)
		return s.BudgetParams(ctx, p)
	}
}

// NewDistributionParamsEndpoint returns an endpoint function that calls the
// method "distributionParams" of service "sdk-utilities".
func NewDistributionParamsEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*DistributionParamsPayload)
		return s.DistributionParams(ctx, p)
	}
}

// NewCrescentPoolsEndpoint returns an endpoint function that calls the method
// "crescentPools" of service "sdk-utilities".
func NewCrescentPoolsEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*CrescentPoolsPayload)
		return s.CrescentPools(ctx, p)
	}
}
