// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

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
	// Supply implements supply.
	Supply(ctx context.Context, in *SupplyRequest, opts ...grpc.CallOption) (*SupplyResponse, error)
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
	// Auth implements auth.
	AuthEndpoint(ctx context.Context, in *AuthRequest, opts ...grpc.CallOption) (*AuthResponse, error)
	// Bank implements bank.
	Bank(ctx context.Context, in *BankRequest, opts ...grpc.CallOption) (*BankResponse, error)
	// Delegation implements delegation.
	DelegationEndpoint(ctx context.Context, in *DelegationRequest, opts ...grpc.CallOption) (*DelegationResponse, error)
	// IbcChannel implements ibc_channel.
	IbcChannel(ctx context.Context, in *IbcChannelRequest, opts ...grpc.CallOption) (*IbcChannelResponse, error)
	// IbcClientState implements ibc_client_state.
	IbcClientState(ctx context.Context, in *IbcClientStateRequest, opts ...grpc.CallOption) (*IbcClientStateResponse, error)
	// IbcConnection implements ibc_connection.
	IbcConnection(ctx context.Context, in *IbcConnectionRequest, opts ...grpc.CallOption) (*IbcConnectionResponse, error)
	// IbcDenomTrace implements ibc_denom_trace.
	IbcDenomTrace(ctx context.Context, in *IbcDenomTraceRequest, opts ...grpc.CallOption) (*IbcDenomTraceResponse, error)
	// UnbondingDelegation implements unbondingDelegation.
	UnbondingDelegationEndpoint(ctx context.Context, in *UnbondingDelegationRequest, opts ...grpc.CallOption) (*UnbondingDelegationResponse, error)
	// Validator implements validator.
	ValidatorEndpoint(ctx context.Context, in *ValidatorRequest, opts ...grpc.CallOption) (*ValidatorResponse, error)
}

type sdkUtilitiesClient struct {
	cc grpc.ClientConnInterface
}

func NewSdkUtilitiesClient(cc grpc.ClientConnInterface) SdkUtilitiesClient {
	return &sdkUtilitiesClient{cc}
}

func (c *sdkUtilitiesClient) Supply(ctx context.Context, in *SupplyRequest, opts ...grpc.CallOption) (*SupplyResponse, error) {
	out := new(SupplyResponse)
	err := c.cc.Invoke(ctx, "/sdk_utilities.SdkUtilities/Supply", in, out, opts...)
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

func (c *sdkUtilitiesClient) AuthEndpoint(ctx context.Context, in *AuthRequest, opts ...grpc.CallOption) (*AuthResponse, error) {
	out := new(AuthResponse)
	err := c.cc.Invoke(ctx, "/sdk_utilities.SdkUtilities/AuthEndpoint", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sdkUtilitiesClient) Bank(ctx context.Context, in *BankRequest, opts ...grpc.CallOption) (*BankResponse, error) {
	out := new(BankResponse)
	err := c.cc.Invoke(ctx, "/sdk_utilities.SdkUtilities/Bank", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sdkUtilitiesClient) DelegationEndpoint(ctx context.Context, in *DelegationRequest, opts ...grpc.CallOption) (*DelegationResponse, error) {
	out := new(DelegationResponse)
	err := c.cc.Invoke(ctx, "/sdk_utilities.SdkUtilities/DelegationEndpoint", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sdkUtilitiesClient) IbcChannel(ctx context.Context, in *IbcChannelRequest, opts ...grpc.CallOption) (*IbcChannelResponse, error) {
	out := new(IbcChannelResponse)
	err := c.cc.Invoke(ctx, "/sdk_utilities.SdkUtilities/IbcChannel", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sdkUtilitiesClient) IbcClientState(ctx context.Context, in *IbcClientStateRequest, opts ...grpc.CallOption) (*IbcClientStateResponse, error) {
	out := new(IbcClientStateResponse)
	err := c.cc.Invoke(ctx, "/sdk_utilities.SdkUtilities/IbcClientState", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sdkUtilitiesClient) IbcConnection(ctx context.Context, in *IbcConnectionRequest, opts ...grpc.CallOption) (*IbcConnectionResponse, error) {
	out := new(IbcConnectionResponse)
	err := c.cc.Invoke(ctx, "/sdk_utilities.SdkUtilities/IbcConnection", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sdkUtilitiesClient) IbcDenomTrace(ctx context.Context, in *IbcDenomTraceRequest, opts ...grpc.CallOption) (*IbcDenomTraceResponse, error) {
	out := new(IbcDenomTraceResponse)
	err := c.cc.Invoke(ctx, "/sdk_utilities.SdkUtilities/IbcDenomTrace", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sdkUtilitiesClient) UnbondingDelegationEndpoint(ctx context.Context, in *UnbondingDelegationRequest, opts ...grpc.CallOption) (*UnbondingDelegationResponse, error) {
	out := new(UnbondingDelegationResponse)
	err := c.cc.Invoke(ctx, "/sdk_utilities.SdkUtilities/UnbondingDelegationEndpoint", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sdkUtilitiesClient) ValidatorEndpoint(ctx context.Context, in *ValidatorRequest, opts ...grpc.CallOption) (*ValidatorResponse, error) {
	out := new(ValidatorResponse)
	err := c.cc.Invoke(ctx, "/sdk_utilities.SdkUtilities/ValidatorEndpoint", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SdkUtilitiesServer is the server API for SdkUtilities service.
// All implementations must embed UnimplementedSdkUtilitiesServer
// for forward compatibility
type SdkUtilitiesServer interface {
	// Supply implements supply.
	Supply(context.Context, *SupplyRequest) (*SupplyResponse, error)
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
	// Auth implements auth.
	AuthEndpoint(context.Context, *AuthRequest) (*AuthResponse, error)
	// Bank implements bank.
	Bank(context.Context, *BankRequest) (*BankResponse, error)
	// Delegation implements delegation.
	DelegationEndpoint(context.Context, *DelegationRequest) (*DelegationResponse, error)
	// IbcChannel implements ibc_channel.
	IbcChannel(context.Context, *IbcChannelRequest) (*IbcChannelResponse, error)
	// IbcClientState implements ibc_client_state.
	IbcClientState(context.Context, *IbcClientStateRequest) (*IbcClientStateResponse, error)
	// IbcConnection implements ibc_connection.
	IbcConnection(context.Context, *IbcConnectionRequest) (*IbcConnectionResponse, error)
	// IbcDenomTrace implements ibc_denom_trace.
	IbcDenomTrace(context.Context, *IbcDenomTraceRequest) (*IbcDenomTraceResponse, error)
	// UnbondingDelegation implements unbondingDelegation.
	UnbondingDelegationEndpoint(context.Context, *UnbondingDelegationRequest) (*UnbondingDelegationResponse, error)
	// Validator implements validator.
	ValidatorEndpoint(context.Context, *ValidatorRequest) (*ValidatorResponse, error)
	mustEmbedUnimplementedSdkUtilitiesServer()
}

// UnimplementedSdkUtilitiesServer must be embedded to have forward compatible implementations.
type UnimplementedSdkUtilitiesServer struct {
}

func (UnimplementedSdkUtilitiesServer) Supply(context.Context, *SupplyRequest) (*SupplyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Supply not implemented")
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
func (UnimplementedSdkUtilitiesServer) AuthEndpoint(context.Context, *AuthRequest) (*AuthResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AuthEndpoint not implemented")
}
func (UnimplementedSdkUtilitiesServer) Bank(context.Context, *BankRequest) (*BankResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Bank not implemented")
}
func (UnimplementedSdkUtilitiesServer) DelegationEndpoint(context.Context, *DelegationRequest) (*DelegationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DelegationEndpoint not implemented")
}
func (UnimplementedSdkUtilitiesServer) IbcChannel(context.Context, *IbcChannelRequest) (*IbcChannelResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IbcChannel not implemented")
}
func (UnimplementedSdkUtilitiesServer) IbcClientState(context.Context, *IbcClientStateRequest) (*IbcClientStateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IbcClientState not implemented")
}
func (UnimplementedSdkUtilitiesServer) IbcConnection(context.Context, *IbcConnectionRequest) (*IbcConnectionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IbcConnection not implemented")
}
func (UnimplementedSdkUtilitiesServer) IbcDenomTrace(context.Context, *IbcDenomTraceRequest) (*IbcDenomTraceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IbcDenomTrace not implemented")
}
func (UnimplementedSdkUtilitiesServer) UnbondingDelegationEndpoint(context.Context, *UnbondingDelegationRequest) (*UnbondingDelegationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UnbondingDelegationEndpoint not implemented")
}
func (UnimplementedSdkUtilitiesServer) ValidatorEndpoint(context.Context, *ValidatorRequest) (*ValidatorResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ValidatorEndpoint not implemented")
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

func _SdkUtilities_AuthEndpoint_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AuthRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SdkUtilitiesServer).AuthEndpoint(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sdk_utilities.SdkUtilities/AuthEndpoint",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SdkUtilitiesServer).AuthEndpoint(ctx, req.(*AuthRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SdkUtilities_Bank_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BankRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SdkUtilitiesServer).Bank(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sdk_utilities.SdkUtilities/Bank",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SdkUtilitiesServer).Bank(ctx, req.(*BankRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SdkUtilities_DelegationEndpoint_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DelegationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SdkUtilitiesServer).DelegationEndpoint(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sdk_utilities.SdkUtilities/DelegationEndpoint",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SdkUtilitiesServer).DelegationEndpoint(ctx, req.(*DelegationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SdkUtilities_IbcChannel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IbcChannelRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SdkUtilitiesServer).IbcChannel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sdk_utilities.SdkUtilities/IbcChannel",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SdkUtilitiesServer).IbcChannel(ctx, req.(*IbcChannelRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SdkUtilities_IbcClientState_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IbcClientStateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SdkUtilitiesServer).IbcClientState(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sdk_utilities.SdkUtilities/IbcClientState",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SdkUtilitiesServer).IbcClientState(ctx, req.(*IbcClientStateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SdkUtilities_IbcConnection_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IbcConnectionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SdkUtilitiesServer).IbcConnection(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sdk_utilities.SdkUtilities/IbcConnection",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SdkUtilitiesServer).IbcConnection(ctx, req.(*IbcConnectionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SdkUtilities_IbcDenomTrace_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IbcDenomTraceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SdkUtilitiesServer).IbcDenomTrace(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sdk_utilities.SdkUtilities/IbcDenomTrace",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SdkUtilitiesServer).IbcDenomTrace(ctx, req.(*IbcDenomTraceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SdkUtilities_UnbondingDelegationEndpoint_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UnbondingDelegationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SdkUtilitiesServer).UnbondingDelegationEndpoint(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sdk_utilities.SdkUtilities/UnbondingDelegationEndpoint",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SdkUtilitiesServer).UnbondingDelegationEndpoint(ctx, req.(*UnbondingDelegationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SdkUtilities_ValidatorEndpoint_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ValidatorRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SdkUtilitiesServer).ValidatorEndpoint(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sdk_utilities.SdkUtilities/ValidatorEndpoint",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SdkUtilitiesServer).ValidatorEndpoint(ctx, req.(*ValidatorRequest))
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
			MethodName: "Supply",
			Handler:    _SdkUtilities_Supply_Handler,
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
			MethodName: "AuthEndpoint",
			Handler:    _SdkUtilities_AuthEndpoint_Handler,
		},
		{
			MethodName: "Bank",
			Handler:    _SdkUtilities_Bank_Handler,
		},
		{
			MethodName: "DelegationEndpoint",
			Handler:    _SdkUtilities_DelegationEndpoint_Handler,
		},
		{
			MethodName: "IbcChannel",
			Handler:    _SdkUtilities_IbcChannel_Handler,
		},
		{
			MethodName: "IbcClientState",
			Handler:    _SdkUtilities_IbcClientState_Handler,
		},
		{
			MethodName: "IbcConnection",
			Handler:    _SdkUtilities_IbcConnection_Handler,
		},
		{
			MethodName: "IbcDenomTrace",
			Handler:    _SdkUtilities_IbcDenomTrace_Handler,
		},
		{
			MethodName: "UnbondingDelegationEndpoint",
			Handler:    _SdkUtilities_UnbondingDelegationEndpoint_Handler,
		},
		{
			MethodName: "ValidatorEndpoint",
			Handler:    _SdkUtilities_ValidatorEndpoint_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "sdk_utilities.proto",
}
