// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.16.0
// source: goadesign_goagen_sdk_utilities.proto

package sdk_utilitiespb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// SdkUtilitiesClient is the client API for SdkUtilities service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SdkUtilitiesClient interface {
	// AccountNumbers implements accountNumbers.
	AccountNumbers(ctx context.Context, in *AccountNumbersRequest, opts ...grpc.CallOption) (*AccountNumbersResponse, error)
	// Supply implements supply.
	Supply(ctx context.Context, in *SupplyRequest, opts ...grpc.CallOption) (*SupplyResponse, error)
	// SupplyDenom implements supplyDenom.
	SupplyDenom(ctx context.Context, in *SupplyDenomRequest, opts ...grpc.CallOption) (*SupplyDenomResponse, error)
	// QueryTx implements queryTx.
	QueryTx(ctx context.Context, in *QueryTxRequest, opts ...grpc.CallOption) (*QueryTxResponse, error)
	// BroadcastTx implements broadcastTx.
	BroadcastTx(ctx context.Context, in *BroadcastTxRequest, opts ...grpc.CallOption) (*BroadcastTxResponse, error)
	// TxMetadata implements txMetadata.
	TxMetadata(ctx context.Context, in *TxMetadataRequest, opts ...grpc.CallOption) (*TxMetadataResponse, error)
	// Block implements block.
	Block(ctx context.Context, in *BlockRequest, opts ...grpc.CallOption) (*BlockResponse, error)
	// LiquidityParams implements liquidityParams.
	LiquidityParams(ctx context.Context, in *LiquidityParamsRequest, opts ...grpc.CallOption) (*LiquidityParamsResponse, error)
	// LiquidityPools implements liquidityPools.
	LiquidityPools(ctx context.Context, in *LiquidityPoolsRequest, opts ...grpc.CallOption) (*LiquidityPoolsResponse, error)
	// MintInflation implements mintInflation.
	MintInflation(ctx context.Context, in *MintInflationRequest, opts ...grpc.CallOption) (*MintInflationResponse, error)
	// MintParams implements mintParams.
	MintParams(ctx context.Context, in *MintParamsRequest, opts ...grpc.CallOption) (*MintParamsResponse, error)
	// MintAnnualProvision implements mintAnnualProvision.
	MintAnnualProvision(ctx context.Context, in *MintAnnualProvisionRequest, opts ...grpc.CallOption) (*MintAnnualProvisionResponse, error)
	// MintEpochProvisions implements mintEpochProvisions.
	MintEpochProvisions(ctx context.Context, in *MintEpochProvisionsRequest, opts ...grpc.CallOption) (*MintEpochProvisionsResponse, error)
	// DelegatorRewards implements delegatorRewards.
	DelegatorRewards(ctx context.Context, in *DelegatorRewardsRequest, opts ...grpc.CallOption) (*DelegatorRewardsResponse, error)
	// EstimateFees implements estimateFees.
	EstimateFees(ctx context.Context, in *EstimateFeesRequest, opts ...grpc.CallOption) (*EstimateFeesResponse, error)
	// StakingParams implements stakingParams.
	StakingParams(ctx context.Context, in *StakingParamsRequest, opts ...grpc.CallOption) (*StakingParamsResponse, error)
	// StakingPool implements stakingPool.
	StakingPool(ctx context.Context, in *StakingPoolRequest, opts ...grpc.CallOption) (*StakingPoolResponse, error)
	// EmoneyInflation implements emoneyInflation.
	EmoneyInflation(ctx context.Context, in *EmoneyInflationRequest, opts ...grpc.CallOption) (*EmoneyInflationResponse, error)
	// IbcChannelClientState implements ibcChannelClientState.
	IbcChannelClientState(ctx context.Context, in *IbcChannelClientStateRequest, opts ...grpc.CallOption) (*IbcChannelClientStateResponse, error)
}

type sdkUtilitiesClient struct {
	cc grpc.ClientConnInterface
}

func NewSdkUtilitiesClient(cc grpc.ClientConnInterface) SdkUtilitiesClient {
	return &sdkUtilitiesClient{cc}
}

func (c *sdkUtilitiesClient) AccountNumbers(ctx context.Context, in *AccountNumbersRequest, opts ...grpc.CallOption) (*AccountNumbersResponse, error) {
	out := new(AccountNumbersResponse)
	err := c.cc.Invoke(ctx, "/sdk_utilities.SdkUtilities/AccountNumbers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sdkUtilitiesClient) Supply(ctx context.Context, in *SupplyRequest, opts ...grpc.CallOption) (*SupplyResponse, error) {
	out := new(SupplyResponse)
	err := c.cc.Invoke(ctx, "/sdk_utilities.SdkUtilities/Supply", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sdkUtilitiesClient) SupplyDenom(ctx context.Context, in *SupplyDenomRequest, opts ...grpc.CallOption) (*SupplyDenomResponse, error) {
	out := new(SupplyDenomResponse)
	err := c.cc.Invoke(ctx, "/sdk_utilities.SdkUtilities/SupplyDenom", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sdkUtilitiesClient) QueryTx(ctx context.Context, in *QueryTxRequest, opts ...grpc.CallOption) (*QueryTxResponse, error) {
	out := new(QueryTxResponse)
	err := c.cc.Invoke(ctx, "/sdk_utilities.SdkUtilities/QueryTx", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sdkUtilitiesClient) BroadcastTx(ctx context.Context, in *BroadcastTxRequest, opts ...grpc.CallOption) (*BroadcastTxResponse, error) {
	out := new(BroadcastTxResponse)
	err := c.cc.Invoke(ctx, "/sdk_utilities.SdkUtilities/BroadcastTx", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sdkUtilitiesClient) TxMetadata(ctx context.Context, in *TxMetadataRequest, opts ...grpc.CallOption) (*TxMetadataResponse, error) {
	out := new(TxMetadataResponse)
	err := c.cc.Invoke(ctx, "/sdk_utilities.SdkUtilities/TxMetadata", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sdkUtilitiesClient) Block(ctx context.Context, in *BlockRequest, opts ...grpc.CallOption) (*BlockResponse, error) {
	out := new(BlockResponse)
	err := c.cc.Invoke(ctx, "/sdk_utilities.SdkUtilities/Block", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sdkUtilitiesClient) LiquidityParams(ctx context.Context, in *LiquidityParamsRequest, opts ...grpc.CallOption) (*LiquidityParamsResponse, error) {
	out := new(LiquidityParamsResponse)
	err := c.cc.Invoke(ctx, "/sdk_utilities.SdkUtilities/LiquidityParams", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sdkUtilitiesClient) LiquidityPools(ctx context.Context, in *LiquidityPoolsRequest, opts ...grpc.CallOption) (*LiquidityPoolsResponse, error) {
	out := new(LiquidityPoolsResponse)
	err := c.cc.Invoke(ctx, "/sdk_utilities.SdkUtilities/LiquidityPools", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sdkUtilitiesClient) MintInflation(ctx context.Context, in *MintInflationRequest, opts ...grpc.CallOption) (*MintInflationResponse, error) {
	out := new(MintInflationResponse)
	err := c.cc.Invoke(ctx, "/sdk_utilities.SdkUtilities/MintInflation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sdkUtilitiesClient) MintParams(ctx context.Context, in *MintParamsRequest, opts ...grpc.CallOption) (*MintParamsResponse, error) {
	out := new(MintParamsResponse)
	err := c.cc.Invoke(ctx, "/sdk_utilities.SdkUtilities/MintParams", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sdkUtilitiesClient) MintAnnualProvision(ctx context.Context, in *MintAnnualProvisionRequest, opts ...grpc.CallOption) (*MintAnnualProvisionResponse, error) {
	out := new(MintAnnualProvisionResponse)
	err := c.cc.Invoke(ctx, "/sdk_utilities.SdkUtilities/MintAnnualProvision", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sdkUtilitiesClient) MintEpochProvisions(ctx context.Context, in *MintEpochProvisionsRequest, opts ...grpc.CallOption) (*MintEpochProvisionsResponse, error) {
	out := new(MintEpochProvisionsResponse)
	err := c.cc.Invoke(ctx, "/sdk_utilities.SdkUtilities/MintEpochProvisions", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sdkUtilitiesClient) DelegatorRewards(ctx context.Context, in *DelegatorRewardsRequest, opts ...grpc.CallOption) (*DelegatorRewardsResponse, error) {
	out := new(DelegatorRewardsResponse)
	err := c.cc.Invoke(ctx, "/sdk_utilities.SdkUtilities/DelegatorRewards", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sdkUtilitiesClient) EstimateFees(ctx context.Context, in *EstimateFeesRequest, opts ...grpc.CallOption) (*EstimateFeesResponse, error) {
	out := new(EstimateFeesResponse)
	err := c.cc.Invoke(ctx, "/sdk_utilities.SdkUtilities/EstimateFees", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sdkUtilitiesClient) StakingParams(ctx context.Context, in *StakingParamsRequest, opts ...grpc.CallOption) (*StakingParamsResponse, error) {
	out := new(StakingParamsResponse)
	err := c.cc.Invoke(ctx, "/sdk_utilities.SdkUtilities/StakingParams", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sdkUtilitiesClient) StakingPool(ctx context.Context, in *StakingPoolRequest, opts ...grpc.CallOption) (*StakingPoolResponse, error) {
	out := new(StakingPoolResponse)
	err := c.cc.Invoke(ctx, "/sdk_utilities.SdkUtilities/StakingPool", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sdkUtilitiesClient) EmoneyInflation(ctx context.Context, in *EmoneyInflationRequest, opts ...grpc.CallOption) (*EmoneyInflationResponse, error) {
	out := new(EmoneyInflationResponse)
	err := c.cc.Invoke(ctx, "/sdk_utilities.SdkUtilities/EmoneyInflation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sdkUtilitiesClient) IbcChannelClientState(ctx context.Context, in *IbcChannelClientStateRequest, opts ...grpc.CallOption) (*IbcChannelClientStateResponse, error) {
	out := new(IbcChannelClientStateResponse)
	err := c.cc.Invoke(ctx, "/sdk_utilities.SdkUtilities/IbcChannelClientState", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SdkUtilitiesServer is the server API for SdkUtilities service.
// All implementations must embed UnimplementedSdkUtilitiesServer
// for forward compatibility
type SdkUtilitiesServer interface {
	// AccountNumbers implements accountNumbers.
	AccountNumbers(context.Context, *AccountNumbersRequest) (*AccountNumbersResponse, error)
	// Supply implements supply.
	Supply(context.Context, *SupplyRequest) (*SupplyResponse, error)
	// SupplyDenom implements supplyDenom.
	SupplyDenom(context.Context, *SupplyDenomRequest) (*SupplyDenomResponse, error)
	// QueryTx implements queryTx.
	QueryTx(context.Context, *QueryTxRequest) (*QueryTxResponse, error)
	// BroadcastTx implements broadcastTx.
	BroadcastTx(context.Context, *BroadcastTxRequest) (*BroadcastTxResponse, error)
	// TxMetadata implements txMetadata.
	TxMetadata(context.Context, *TxMetadataRequest) (*TxMetadataResponse, error)
	// Block implements block.
	Block(context.Context, *BlockRequest) (*BlockResponse, error)
	// LiquidityParams implements liquidityParams.
	LiquidityParams(context.Context, *LiquidityParamsRequest) (*LiquidityParamsResponse, error)
	// LiquidityPools implements liquidityPools.
	LiquidityPools(context.Context, *LiquidityPoolsRequest) (*LiquidityPoolsResponse, error)
	// MintInflation implements mintInflation.
	MintInflation(context.Context, *MintInflationRequest) (*MintInflationResponse, error)
	// MintParams implements mintParams.
	MintParams(context.Context, *MintParamsRequest) (*MintParamsResponse, error)
	// MintAnnualProvision implements mintAnnualProvision.
	MintAnnualProvision(context.Context, *MintAnnualProvisionRequest) (*MintAnnualProvisionResponse, error)
	// MintEpochProvisions implements mintEpochProvisions.
	MintEpochProvisions(context.Context, *MintEpochProvisionsRequest) (*MintEpochProvisionsResponse, error)
	// DelegatorRewards implements delegatorRewards.
	DelegatorRewards(context.Context, *DelegatorRewardsRequest) (*DelegatorRewardsResponse, error)
	// EstimateFees implements estimateFees.
	EstimateFees(context.Context, *EstimateFeesRequest) (*EstimateFeesResponse, error)
	// StakingParams implements stakingParams.
	StakingParams(context.Context, *StakingParamsRequest) (*StakingParamsResponse, error)
	// StakingPool implements stakingPool.
	StakingPool(context.Context, *StakingPoolRequest) (*StakingPoolResponse, error)
	// EmoneyInflation implements emoneyInflation.
	EmoneyInflation(context.Context, *EmoneyInflationRequest) (*EmoneyInflationResponse, error)
	// IbcChannelClientState implements ibcChannelClientState.
	IbcChannelClientState(context.Context, *IbcChannelClientStateRequest) (*IbcChannelClientStateResponse, error)
	mustEmbedUnimplementedSdkUtilitiesServer()
}

// UnimplementedSdkUtilitiesServer must be embedded to have forward compatible implementations.
type UnimplementedSdkUtilitiesServer struct {
}

func (UnimplementedSdkUtilitiesServer) AccountNumbers(context.Context, *AccountNumbersRequest) (*AccountNumbersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AccountNumbers not implemented")
}
func (UnimplementedSdkUtilitiesServer) Supply(context.Context, *SupplyRequest) (*SupplyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Supply not implemented")
}
func (UnimplementedSdkUtilitiesServer) SupplyDenom(context.Context, *SupplyDenomRequest) (*SupplyDenomResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SupplyDenom not implemented")
}
func (UnimplementedSdkUtilitiesServer) QueryTx(context.Context, *QueryTxRequest) (*QueryTxResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryTx not implemented")
}
func (UnimplementedSdkUtilitiesServer) BroadcastTx(context.Context, *BroadcastTxRequest) (*BroadcastTxResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BroadcastTx not implemented")
}
func (UnimplementedSdkUtilitiesServer) TxMetadata(context.Context, *TxMetadataRequest) (*TxMetadataResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TxMetadata not implemented")
}
func (UnimplementedSdkUtilitiesServer) Block(context.Context, *BlockRequest) (*BlockResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Block not implemented")
}
func (UnimplementedSdkUtilitiesServer) LiquidityParams(context.Context, *LiquidityParamsRequest) (*LiquidityParamsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LiquidityParams not implemented")
}
func (UnimplementedSdkUtilitiesServer) LiquidityPools(context.Context, *LiquidityPoolsRequest) (*LiquidityPoolsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LiquidityPools not implemented")
}
func (UnimplementedSdkUtilitiesServer) MintInflation(context.Context, *MintInflationRequest) (*MintInflationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MintInflation not implemented")
}
func (UnimplementedSdkUtilitiesServer) MintParams(context.Context, *MintParamsRequest) (*MintParamsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MintParams not implemented")
}
func (UnimplementedSdkUtilitiesServer) MintAnnualProvision(context.Context, *MintAnnualProvisionRequest) (*MintAnnualProvisionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MintAnnualProvision not implemented")
}
func (UnimplementedSdkUtilitiesServer) MintEpochProvisions(context.Context, *MintEpochProvisionsRequest) (*MintEpochProvisionsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MintEpochProvisions not implemented")
}
func (UnimplementedSdkUtilitiesServer) DelegatorRewards(context.Context, *DelegatorRewardsRequest) (*DelegatorRewardsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DelegatorRewards not implemented")
}
func (UnimplementedSdkUtilitiesServer) EstimateFees(context.Context, *EstimateFeesRequest) (*EstimateFeesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EstimateFees not implemented")
}
func (UnimplementedSdkUtilitiesServer) StakingParams(context.Context, *StakingParamsRequest) (*StakingParamsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StakingParams not implemented")
}
func (UnimplementedSdkUtilitiesServer) StakingPool(context.Context, *StakingPoolRequest) (*StakingPoolResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StakingPool not implemented")
}
func (UnimplementedSdkUtilitiesServer) EmoneyInflation(context.Context, *EmoneyInflationRequest) (*EmoneyInflationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EmoneyInflation not implemented")
}
func (UnimplementedSdkUtilitiesServer) IbcChannelClientState(context.Context, *IbcChannelClientStateRequest) (*IbcChannelClientStateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IbcChannelClientState not implemented")
}
func (UnimplementedSdkUtilitiesServer) mustEmbedUnimplementedSdkUtilitiesServer() {}

// UnsafeSdkUtilitiesServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SdkUtilitiesServer will
// result in compilation errors.
type UnsafeSdkUtilitiesServer interface {
	mustEmbedUnimplementedSdkUtilitiesServer()
}

func RegisterSdkUtilitiesServer(s grpc.ServiceRegistrar, srv SdkUtilitiesServer) {
	s.RegisterService(&SdkUtilities_ServiceDesc, srv)
}

func _SdkUtilities_AccountNumbers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AccountNumbersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SdkUtilitiesServer).AccountNumbers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sdk_utilities.SdkUtilities/AccountNumbers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SdkUtilitiesServer).AccountNumbers(ctx, req.(*AccountNumbersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SdkUtilities_Supply_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SupplyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SdkUtilitiesServer).Supply(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sdk_utilities.SdkUtilities/Supply",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SdkUtilitiesServer).Supply(ctx, req.(*SupplyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SdkUtilities_SupplyDenom_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SupplyDenomRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SdkUtilitiesServer).SupplyDenom(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sdk_utilities.SdkUtilities/SupplyDenom",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SdkUtilitiesServer).SupplyDenom(ctx, req.(*SupplyDenomRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SdkUtilities_QueryTx_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryTxRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SdkUtilitiesServer).QueryTx(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sdk_utilities.SdkUtilities/QueryTx",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SdkUtilitiesServer).QueryTx(ctx, req.(*QueryTxRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SdkUtilities_BroadcastTx_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BroadcastTxRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SdkUtilitiesServer).BroadcastTx(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sdk_utilities.SdkUtilities/BroadcastTx",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SdkUtilitiesServer).BroadcastTx(ctx, req.(*BroadcastTxRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SdkUtilities_TxMetadata_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TxMetadataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SdkUtilitiesServer).TxMetadata(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sdk_utilities.SdkUtilities/TxMetadata",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SdkUtilitiesServer).TxMetadata(ctx, req.(*TxMetadataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SdkUtilities_Block_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BlockRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SdkUtilitiesServer).Block(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sdk_utilities.SdkUtilities/Block",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SdkUtilitiesServer).Block(ctx, req.(*BlockRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SdkUtilities_LiquidityParams_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LiquidityParamsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SdkUtilitiesServer).LiquidityParams(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sdk_utilities.SdkUtilities/LiquidityParams",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SdkUtilitiesServer).LiquidityParams(ctx, req.(*LiquidityParamsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SdkUtilities_LiquidityPools_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LiquidityPoolsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SdkUtilitiesServer).LiquidityPools(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sdk_utilities.SdkUtilities/LiquidityPools",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SdkUtilitiesServer).LiquidityPools(ctx, req.(*LiquidityPoolsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SdkUtilities_MintInflation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MintInflationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SdkUtilitiesServer).MintInflation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sdk_utilities.SdkUtilities/MintInflation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SdkUtilitiesServer).MintInflation(ctx, req.(*MintInflationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SdkUtilities_MintParams_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MintParamsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SdkUtilitiesServer).MintParams(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sdk_utilities.SdkUtilities/MintParams",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SdkUtilitiesServer).MintParams(ctx, req.(*MintParamsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SdkUtilities_MintAnnualProvision_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MintAnnualProvisionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SdkUtilitiesServer).MintAnnualProvision(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sdk_utilities.SdkUtilities/MintAnnualProvision",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SdkUtilitiesServer).MintAnnualProvision(ctx, req.(*MintAnnualProvisionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SdkUtilities_MintEpochProvisions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MintEpochProvisionsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SdkUtilitiesServer).MintEpochProvisions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sdk_utilities.SdkUtilities/MintEpochProvisions",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SdkUtilitiesServer).MintEpochProvisions(ctx, req.(*MintEpochProvisionsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SdkUtilities_DelegatorRewards_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DelegatorRewardsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SdkUtilitiesServer).DelegatorRewards(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sdk_utilities.SdkUtilities/DelegatorRewards",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SdkUtilitiesServer).DelegatorRewards(ctx, req.(*DelegatorRewardsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SdkUtilities_EstimateFees_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EstimateFeesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SdkUtilitiesServer).EstimateFees(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sdk_utilities.SdkUtilities/EstimateFees",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SdkUtilitiesServer).EstimateFees(ctx, req.(*EstimateFeesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SdkUtilities_StakingParams_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StakingParamsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SdkUtilitiesServer).StakingParams(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sdk_utilities.SdkUtilities/StakingParams",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SdkUtilitiesServer).StakingParams(ctx, req.(*StakingParamsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SdkUtilities_StakingPool_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StakingPoolRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SdkUtilitiesServer).StakingPool(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sdk_utilities.SdkUtilities/StakingPool",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SdkUtilitiesServer).StakingPool(ctx, req.(*StakingPoolRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SdkUtilities_EmoneyInflation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EmoneyInflationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SdkUtilitiesServer).EmoneyInflation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sdk_utilities.SdkUtilities/EmoneyInflation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SdkUtilitiesServer).EmoneyInflation(ctx, req.(*EmoneyInflationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SdkUtilities_IbcChannelClientState_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IbcChannelClientStateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SdkUtilitiesServer).IbcChannelClientState(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sdk_utilities.SdkUtilities/IbcChannelClientState",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SdkUtilitiesServer).IbcChannelClientState(ctx, req.(*IbcChannelClientStateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// SdkUtilities_ServiceDesc is the grpc.ServiceDesc for SdkUtilities service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SdkUtilities_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "sdk_utilities.SdkUtilities",
	HandlerType: (*SdkUtilitiesServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AccountNumbers",
			Handler:    _SdkUtilities_AccountNumbers_Handler,
		},
		{
			MethodName: "Supply",
			Handler:    _SdkUtilities_Supply_Handler,
		},
		{
			MethodName: "SupplyDenom",
			Handler:    _SdkUtilities_SupplyDenom_Handler,
		},
		{
			MethodName: "QueryTx",
			Handler:    _SdkUtilities_QueryTx_Handler,
		},
		{
			MethodName: "BroadcastTx",
			Handler:    _SdkUtilities_BroadcastTx_Handler,
		},
		{
			MethodName: "TxMetadata",
			Handler:    _SdkUtilities_TxMetadata_Handler,
		},
		{
			MethodName: "Block",
			Handler:    _SdkUtilities_Block_Handler,
		},
		{
			MethodName: "LiquidityParams",
			Handler:    _SdkUtilities_LiquidityParams_Handler,
		},
		{
			MethodName: "LiquidityPools",
			Handler:    _SdkUtilities_LiquidityPools_Handler,
		},
		{
			MethodName: "MintInflation",
			Handler:    _SdkUtilities_MintInflation_Handler,
		},
		{
			MethodName: "MintParams",
			Handler:    _SdkUtilities_MintParams_Handler,
		},
		{
			MethodName: "MintAnnualProvision",
			Handler:    _SdkUtilities_MintAnnualProvision_Handler,
		},
		{
			MethodName: "MintEpochProvisions",
			Handler:    _SdkUtilities_MintEpochProvisions_Handler,
		},
		{
			MethodName: "DelegatorRewards",
			Handler:    _SdkUtilities_DelegatorRewards_Handler,
		},
		{
			MethodName: "EstimateFees",
			Handler:    _SdkUtilities_EstimateFees_Handler,
		},
		{
			MethodName: "StakingParams",
			Handler:    _SdkUtilities_StakingParams_Handler,
		},
		{
			MethodName: "StakingPool",
			Handler:    _SdkUtilities_StakingPool_Handler,
		},
		{
			MethodName: "EmoneyInflation",
			Handler:    _SdkUtilities_EmoneyInflation_Handler,
		},
		{
			MethodName: "IbcChannelClientState",
			Handler:    _SdkUtilities_IbcChannelClientState_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "goadesign_goagen_sdk_utilities.proto",
}
