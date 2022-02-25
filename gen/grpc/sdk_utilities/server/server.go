// Code generated by goa v3.5.5, DO NOT EDIT.
//
// sdk-utilities gRPC server
//
// Command:
// $ goa gen github.com/allinbits/sdk-service-meta

package server

import (
	"context"

	sdk_utilitiespb "github.com/allinbits/sdk-service-meta/gen/grpc/sdk_utilities/pb"
	sdkutilities "github.com/allinbits/sdk-service-meta/gen/sdk_utilities"
	goagrpc "goa.design/goa/v3/grpc"
	goa "goa.design/goa/v3/pkg"
)

// Server implements the sdk_utilitiespb.SdkUtilitiesServer interface.
type Server struct {
	AccountNumbersH      goagrpc.UnaryHandler
	SupplyH              goagrpc.UnaryHandler
	QueryTxH             goagrpc.UnaryHandler
	BroadcastTxH         goagrpc.UnaryHandler
	TxMetadataH          goagrpc.UnaryHandler
	BlockH               goagrpc.UnaryHandler
	LiquidityParamsH     goagrpc.UnaryHandler
	LiquidityPoolsH      goagrpc.UnaryHandler
	MintInflationH       goagrpc.UnaryHandler
	MintParamsH          goagrpc.UnaryHandler
	MintAnnualProvisionH goagrpc.UnaryHandler
	DelegatorRewardsH    goagrpc.UnaryHandler
	EstimateFeesH        goagrpc.UnaryHandler
	StakingParamsH       goagrpc.UnaryHandler
	sdk_utilitiespb.UnimplementedSdkUtilitiesServer
}

// ErrorNamer is an interface implemented by generated error structs that
// exposes the name of the error as defined in the expr.
type ErrorNamer interface {
	ErrorName() string
}

// New instantiates the server struct with the sdk-utilities service endpoints.
func New(e *sdkutilities.Endpoints, uh goagrpc.UnaryHandler) *Server {
	return &Server{
		AccountNumbersH:      NewAccountNumbersHandler(e.AccountNumbers, uh),
		SupplyH:              NewSupplyHandler(e.Supply, uh),
		QueryTxH:             NewQueryTxHandler(e.QueryTx, uh),
		BroadcastTxH:         NewBroadcastTxHandler(e.BroadcastTx, uh),
		TxMetadataH:          NewTxMetadataHandler(e.TxMetadata, uh),
		BlockH:               NewBlockHandler(e.Block, uh),
		LiquidityParamsH:     NewLiquidityParamsHandler(e.LiquidityParams, uh),
		LiquidityPoolsH:      NewLiquidityPoolsHandler(e.LiquidityPools, uh),
		MintInflationH:       NewMintInflationHandler(e.MintInflation, uh),
		MintParamsH:          NewMintParamsHandler(e.MintParams, uh),
		MintAnnualProvisionH: NewMintAnnualProvisionHandler(e.MintAnnualProvision, uh),
		DelegatorRewardsH:    NewDelegatorRewardsHandler(e.DelegatorRewards, uh),
		EstimateFeesH:        NewEstimateFeesHandler(e.EstimateFees, uh),
		StakingParamsH:       NewStakingParamsHandler(e.StakingParams, uh),
	}
}

// NewAccountNumbersHandler creates a gRPC handler which serves the
// "sdk-utilities" service "accountNumbers" endpoint.
func NewAccountNumbersHandler(endpoint goa.Endpoint, h goagrpc.UnaryHandler) goagrpc.UnaryHandler {
	if h == nil {
		h = goagrpc.NewUnaryHandler(endpoint, DecodeAccountNumbersRequest, EncodeAccountNumbersResponse)
	}
	return h
}

// AccountNumbers implements the "AccountNumbers" method in
// sdk_utilitiespb.SdkUtilitiesServer interface.
func (s *Server) AccountNumbers(ctx context.Context, message *sdk_utilitiespb.AccountNumbersRequest) (*sdk_utilitiespb.AccountNumbersResponse, error) {
	ctx = context.WithValue(ctx, goa.MethodKey, "accountNumbers")
	ctx = context.WithValue(ctx, goa.ServiceKey, "sdk-utilities")
	resp, err := s.AccountNumbersH.Handle(ctx, message)
	if err != nil {
		return nil, goagrpc.EncodeError(err)
	}
	return resp.(*sdk_utilitiespb.AccountNumbersResponse), nil
}

// NewSupplyHandler creates a gRPC handler which serves the "sdk-utilities"
// service "supply" endpoint.
func NewSupplyHandler(endpoint goa.Endpoint, h goagrpc.UnaryHandler) goagrpc.UnaryHandler {
	if h == nil {
		h = goagrpc.NewUnaryHandler(endpoint, DecodeSupplyRequest, EncodeSupplyResponse)
	}
	return h
}

// Supply implements the "Supply" method in sdk_utilitiespb.SdkUtilitiesServer
// interface.
func (s *Server) Supply(ctx context.Context, message *sdk_utilitiespb.SupplyRequest) (*sdk_utilitiespb.SupplyResponse, error) {
	ctx = context.WithValue(ctx, goa.MethodKey, "supply")
	ctx = context.WithValue(ctx, goa.ServiceKey, "sdk-utilities")
	resp, err := s.SupplyH.Handle(ctx, message)
	if err != nil {
		return nil, goagrpc.EncodeError(err)
	}
	return resp.(*sdk_utilitiespb.SupplyResponse), nil
}

// NewQueryTxHandler creates a gRPC handler which serves the "sdk-utilities"
// service "queryTx" endpoint.
func NewQueryTxHandler(endpoint goa.Endpoint, h goagrpc.UnaryHandler) goagrpc.UnaryHandler {
	if h == nil {
		h = goagrpc.NewUnaryHandler(endpoint, DecodeQueryTxRequest, EncodeQueryTxResponse)
	}
	return h
}

// QueryTx implements the "QueryTx" method in
// sdk_utilitiespb.SdkUtilitiesServer interface.
func (s *Server) QueryTx(ctx context.Context, message *sdk_utilitiespb.QueryTxRequest) (*sdk_utilitiespb.QueryTxResponse, error) {
	ctx = context.WithValue(ctx, goa.MethodKey, "queryTx")
	ctx = context.WithValue(ctx, goa.ServiceKey, "sdk-utilities")
	resp, err := s.QueryTxH.Handle(ctx, message)
	if err != nil {
		return nil, goagrpc.EncodeError(err)
	}
	return resp.(*sdk_utilitiespb.QueryTxResponse), nil
}

// NewBroadcastTxHandler creates a gRPC handler which serves the
// "sdk-utilities" service "broadcastTx" endpoint.
func NewBroadcastTxHandler(endpoint goa.Endpoint, h goagrpc.UnaryHandler) goagrpc.UnaryHandler {
	if h == nil {
		h = goagrpc.NewUnaryHandler(endpoint, DecodeBroadcastTxRequest, EncodeBroadcastTxResponse)
	}
	return h
}

// BroadcastTx implements the "BroadcastTx" method in
// sdk_utilitiespb.SdkUtilitiesServer interface.
func (s *Server) BroadcastTx(ctx context.Context, message *sdk_utilitiespb.BroadcastTxRequest) (*sdk_utilitiespb.BroadcastTxResponse, error) {
	ctx = context.WithValue(ctx, goa.MethodKey, "broadcastTx")
	ctx = context.WithValue(ctx, goa.ServiceKey, "sdk-utilities")
	resp, err := s.BroadcastTxH.Handle(ctx, message)
	if err != nil {
		return nil, goagrpc.EncodeError(err)
	}
	return resp.(*sdk_utilitiespb.BroadcastTxResponse), nil
}

// NewTxMetadataHandler creates a gRPC handler which serves the "sdk-utilities"
// service "txMetadata" endpoint.
func NewTxMetadataHandler(endpoint goa.Endpoint, h goagrpc.UnaryHandler) goagrpc.UnaryHandler {
	if h == nil {
		h = goagrpc.NewUnaryHandler(endpoint, DecodeTxMetadataRequest, EncodeTxMetadataResponse)
	}
	return h
}

// TxMetadata implements the "TxMetadata" method in
// sdk_utilitiespb.SdkUtilitiesServer interface.
func (s *Server) TxMetadata(ctx context.Context, message *sdk_utilitiespb.TxMetadataRequest) (*sdk_utilitiespb.TxMetadataResponse, error) {
	ctx = context.WithValue(ctx, goa.MethodKey, "txMetadata")
	ctx = context.WithValue(ctx, goa.ServiceKey, "sdk-utilities")
	resp, err := s.TxMetadataH.Handle(ctx, message)
	if err != nil {
		return nil, goagrpc.EncodeError(err)
	}
	return resp.(*sdk_utilitiespb.TxMetadataResponse), nil
}

// NewBlockHandler creates a gRPC handler which serves the "sdk-utilities"
// service "block" endpoint.
func NewBlockHandler(endpoint goa.Endpoint, h goagrpc.UnaryHandler) goagrpc.UnaryHandler {
	if h == nil {
		h = goagrpc.NewUnaryHandler(endpoint, DecodeBlockRequest, EncodeBlockResponse)
	}
	return h
}

// Block implements the "Block" method in sdk_utilitiespb.SdkUtilitiesServer
// interface.
func (s *Server) Block(ctx context.Context, message *sdk_utilitiespb.BlockRequest) (*sdk_utilitiespb.BlockResponse, error) {
	ctx = context.WithValue(ctx, goa.MethodKey, "block")
	ctx = context.WithValue(ctx, goa.ServiceKey, "sdk-utilities")
	resp, err := s.BlockH.Handle(ctx, message)
	if err != nil {
		return nil, goagrpc.EncodeError(err)
	}
	return resp.(*sdk_utilitiespb.BlockResponse), nil
}

// NewLiquidityParamsHandler creates a gRPC handler which serves the
// "sdk-utilities" service "liquidityParams" endpoint.
func NewLiquidityParamsHandler(endpoint goa.Endpoint, h goagrpc.UnaryHandler) goagrpc.UnaryHandler {
	if h == nil {
		h = goagrpc.NewUnaryHandler(endpoint, DecodeLiquidityParamsRequest, EncodeLiquidityParamsResponse)
	}
	return h
}

// LiquidityParams implements the "LiquidityParams" method in
// sdk_utilitiespb.SdkUtilitiesServer interface.
func (s *Server) LiquidityParams(ctx context.Context, message *sdk_utilitiespb.LiquidityParamsRequest) (*sdk_utilitiespb.LiquidityParamsResponse, error) {
	ctx = context.WithValue(ctx, goa.MethodKey, "liquidityParams")
	ctx = context.WithValue(ctx, goa.ServiceKey, "sdk-utilities")
	resp, err := s.LiquidityParamsH.Handle(ctx, message)
	if err != nil {
		return nil, goagrpc.EncodeError(err)
	}
	return resp.(*sdk_utilitiespb.LiquidityParamsResponse), nil
}

// NewLiquidityPoolsHandler creates a gRPC handler which serves the
// "sdk-utilities" service "liquidityPools" endpoint.
func NewLiquidityPoolsHandler(endpoint goa.Endpoint, h goagrpc.UnaryHandler) goagrpc.UnaryHandler {
	if h == nil {
		h = goagrpc.NewUnaryHandler(endpoint, DecodeLiquidityPoolsRequest, EncodeLiquidityPoolsResponse)
	}
	return h
}

// LiquidityPools implements the "LiquidityPools" method in
// sdk_utilitiespb.SdkUtilitiesServer interface.
func (s *Server) LiquidityPools(ctx context.Context, message *sdk_utilitiespb.LiquidityPoolsRequest) (*sdk_utilitiespb.LiquidityPoolsResponse, error) {
	ctx = context.WithValue(ctx, goa.MethodKey, "liquidityPools")
	ctx = context.WithValue(ctx, goa.ServiceKey, "sdk-utilities")
	resp, err := s.LiquidityPoolsH.Handle(ctx, message)
	if err != nil {
		return nil, goagrpc.EncodeError(err)
	}
	return resp.(*sdk_utilitiespb.LiquidityPoolsResponse), nil
}

// NewMintInflationHandler creates a gRPC handler which serves the
// "sdk-utilities" service "mintInflation" endpoint.
func NewMintInflationHandler(endpoint goa.Endpoint, h goagrpc.UnaryHandler) goagrpc.UnaryHandler {
	if h == nil {
		h = goagrpc.NewUnaryHandler(endpoint, DecodeMintInflationRequest, EncodeMintInflationResponse)
	}
	return h
}

// MintInflation implements the "MintInflation" method in
// sdk_utilitiespb.SdkUtilitiesServer interface.
func (s *Server) MintInflation(ctx context.Context, message *sdk_utilitiespb.MintInflationRequest) (*sdk_utilitiespb.MintInflationResponse, error) {
	ctx = context.WithValue(ctx, goa.MethodKey, "mintInflation")
	ctx = context.WithValue(ctx, goa.ServiceKey, "sdk-utilities")
	resp, err := s.MintInflationH.Handle(ctx, message)
	if err != nil {
		return nil, goagrpc.EncodeError(err)
	}
	return resp.(*sdk_utilitiespb.MintInflationResponse), nil
}

// NewMintParamsHandler creates a gRPC handler which serves the "sdk-utilities"
// service "mintParams" endpoint.
func NewMintParamsHandler(endpoint goa.Endpoint, h goagrpc.UnaryHandler) goagrpc.UnaryHandler {
	if h == nil {
		h = goagrpc.NewUnaryHandler(endpoint, DecodeMintParamsRequest, EncodeMintParamsResponse)
	}
	return h
}

// MintParams implements the "MintParams" method in
// sdk_utilitiespb.SdkUtilitiesServer interface.
func (s *Server) MintParams(ctx context.Context, message *sdk_utilitiespb.MintParamsRequest) (*sdk_utilitiespb.MintParamsResponse, error) {
	ctx = context.WithValue(ctx, goa.MethodKey, "mintParams")
	ctx = context.WithValue(ctx, goa.ServiceKey, "sdk-utilities")
	resp, err := s.MintParamsH.Handle(ctx, message)
	if err != nil {
		return nil, goagrpc.EncodeError(err)
	}
	return resp.(*sdk_utilitiespb.MintParamsResponse), nil
}

// NewMintAnnualProvisionHandler creates a gRPC handler which serves the
// "sdk-utilities" service "mintAnnualProvision" endpoint.
func NewMintAnnualProvisionHandler(endpoint goa.Endpoint, h goagrpc.UnaryHandler) goagrpc.UnaryHandler {
	if h == nil {
		h = goagrpc.NewUnaryHandler(endpoint, DecodeMintAnnualProvisionRequest, EncodeMintAnnualProvisionResponse)
	}
	return h
}

// MintAnnualProvision implements the "MintAnnualProvision" method in
// sdk_utilitiespb.SdkUtilitiesServer interface.
func (s *Server) MintAnnualProvision(ctx context.Context, message *sdk_utilitiespb.MintAnnualProvisionRequest) (*sdk_utilitiespb.MintAnnualProvisionResponse, error) {
	ctx = context.WithValue(ctx, goa.MethodKey, "mintAnnualProvision")
	ctx = context.WithValue(ctx, goa.ServiceKey, "sdk-utilities")
	resp, err := s.MintAnnualProvisionH.Handle(ctx, message)
	if err != nil {
		return nil, goagrpc.EncodeError(err)
	}
	return resp.(*sdk_utilitiespb.MintAnnualProvisionResponse), nil
}

// NewDelegatorRewardsHandler creates a gRPC handler which serves the
// "sdk-utilities" service "delegatorRewards" endpoint.
func NewDelegatorRewardsHandler(endpoint goa.Endpoint, h goagrpc.UnaryHandler) goagrpc.UnaryHandler {
	if h == nil {
		h = goagrpc.NewUnaryHandler(endpoint, DecodeDelegatorRewardsRequest, EncodeDelegatorRewardsResponse)
	}
	return h
}

// DelegatorRewards implements the "DelegatorRewards" method in
// sdk_utilitiespb.SdkUtilitiesServer interface.
func (s *Server) DelegatorRewards(ctx context.Context, message *sdk_utilitiespb.DelegatorRewardsRequest) (*sdk_utilitiespb.DelegatorRewardsResponse, error) {
	ctx = context.WithValue(ctx, goa.MethodKey, "delegatorRewards")
	ctx = context.WithValue(ctx, goa.ServiceKey, "sdk-utilities")
	resp, err := s.DelegatorRewardsH.Handle(ctx, message)
	if err != nil {
		return nil, goagrpc.EncodeError(err)
	}
	return resp.(*sdk_utilitiespb.DelegatorRewardsResponse), nil
}

// NewEstimateFeesHandler creates a gRPC handler which serves the
// "sdk-utilities" service "estimateFees" endpoint.
func NewEstimateFeesHandler(endpoint goa.Endpoint, h goagrpc.UnaryHandler) goagrpc.UnaryHandler {
	if h == nil {
		h = goagrpc.NewUnaryHandler(endpoint, DecodeEstimateFeesRequest, EncodeEstimateFeesResponse)
	}
	return h
}

// EstimateFees implements the "EstimateFees" method in
// sdk_utilitiespb.SdkUtilitiesServer interface.
func (s *Server) EstimateFees(ctx context.Context, message *sdk_utilitiespb.EstimateFeesRequest) (*sdk_utilitiespb.EstimateFeesResponse, error) {
	ctx = context.WithValue(ctx, goa.MethodKey, "estimateFees")
	ctx = context.WithValue(ctx, goa.ServiceKey, "sdk-utilities")
	resp, err := s.EstimateFeesH.Handle(ctx, message)
	if err != nil {
		return nil, goagrpc.EncodeError(err)
	}
	return resp.(*sdk_utilitiespb.EstimateFeesResponse), nil
}

// NewStakingParamsHandler creates a gRPC handler which serves the
// "sdk-utilities" service "stakingParams" endpoint.
func NewStakingParamsHandler(endpoint goa.Endpoint, h goagrpc.UnaryHandler) goagrpc.UnaryHandler {
	if h == nil {
		h = goagrpc.NewUnaryHandler(endpoint, DecodeStakingParamsRequest, EncodeStakingParamsResponse)
	}
	return h
}

// StakingParams implements the "StakingParams" method in
// sdk_utilitiespb.SdkUtilitiesServer interface.
func (s *Server) StakingParams(ctx context.Context, message *sdk_utilitiespb.StakingParamsRequest) (*sdk_utilitiespb.StakingParamsResponse, error) {
	ctx = context.WithValue(ctx, goa.MethodKey, "stakingParams")
	ctx = context.WithValue(ctx, goa.ServiceKey, "sdk-utilities")
	resp, err := s.StakingParamsH.Handle(ctx, message)
	if err != nil {
		return nil, goagrpc.EncodeError(err)
	}
	return resp.(*sdk_utilitiespb.StakingParamsResponse), nil
}
