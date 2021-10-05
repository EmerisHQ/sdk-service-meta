// Code generated by goa v3.5.2, DO NOT EDIT.
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
	SupplyH         goagrpc.UnaryHandler
	QueryTxH        goagrpc.UnaryHandler
	BroadcastTxH    goagrpc.UnaryHandler
	TxMetadataH     goagrpc.UnaryHandler
	AuthH           goagrpc.UnaryHandler
	BankH           goagrpc.UnaryHandler
	DelegationH     goagrpc.UnaryHandler
	IbcChannelH     goagrpc.UnaryHandler
	IbcClientStateH goagrpc.UnaryHandler
	IbcConnectionH  goagrpc.UnaryHandler
	IbcDenomTraceH  goagrpc.UnaryHandler
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
		SupplyH:         NewSupplyHandler(e.Supply, uh),
		QueryTxH:        NewQueryTxHandler(e.QueryTx, uh),
		BroadcastTxH:    NewBroadcastTxHandler(e.BroadcastTx, uh),
		TxMetadataH:     NewTxMetadataHandler(e.TxMetadata, uh),
		AuthH:           NewAuthHandler(e.Auth, uh),
		BankH:           NewBankHandler(e.Bank, uh),
		DelegationH:     NewDelegationHandler(e.Delegation, uh),
		IbcChannelH:     NewIbcChannelHandler(e.IbcChannel, uh),
		IbcClientStateH: NewIbcClientStateHandler(e.IbcClientState, uh),
		IbcConnectionH:  NewIbcConnectionHandler(e.IbcConnection, uh),
		IbcDenomTraceH:  NewIbcDenomTraceHandler(e.IbcDenomTrace, uh),
	}
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

// NewAuthHandler creates a gRPC handler which serves the "sdk-utilities"
// service "auth" endpoint.
func NewAuthHandler(endpoint goa.Endpoint, h goagrpc.UnaryHandler) goagrpc.UnaryHandler {
	if h == nil {
		h = goagrpc.NewUnaryHandler(endpoint, DecodeAuthRequest, EncodeAuthResponse)
	}
	return h
}

// Auth implements the "Auth" method in sdk_utilitiespb.SdkUtilitiesServer
// interface.
func (s *Server) Auth(ctx context.Context, message *sdk_utilitiespb.AuthRequest) (*sdk_utilitiespb.AuthResponse, error) {
	ctx = context.WithValue(ctx, goa.MethodKey, "auth")
	ctx = context.WithValue(ctx, goa.ServiceKey, "sdk-utilities")
	resp, err := s.AuthH.Handle(ctx, message)
	if err != nil {
		return nil, goagrpc.EncodeError(err)
	}
	return resp.(*sdk_utilitiespb.AuthResponse), nil
}

// NewBankHandler creates a gRPC handler which serves the "sdk-utilities"
// service "bank" endpoint.
func NewBankHandler(endpoint goa.Endpoint, h goagrpc.UnaryHandler) goagrpc.UnaryHandler {
	if h == nil {
		h = goagrpc.NewUnaryHandler(endpoint, DecodeBankRequest, EncodeBankResponse)
	}
	return h
}

// Bank implements the "Bank" method in sdk_utilitiespb.SdkUtilitiesServer
// interface.
func (s *Server) Bank(ctx context.Context, message *sdk_utilitiespb.BankRequest) (*sdk_utilitiespb.BankResponse, error) {
	ctx = context.WithValue(ctx, goa.MethodKey, "bank")
	ctx = context.WithValue(ctx, goa.ServiceKey, "sdk-utilities")
	resp, err := s.BankH.Handle(ctx, message)
	if err != nil {
		return nil, goagrpc.EncodeError(err)
	}
	return resp.(*sdk_utilitiespb.BankResponse), nil
}

// NewDelegationHandler creates a gRPC handler which serves the "sdk-utilities"
// service "delegation" endpoint.
func NewDelegationHandler(endpoint goa.Endpoint, h goagrpc.UnaryHandler) goagrpc.UnaryHandler {
	if h == nil {
		h = goagrpc.NewUnaryHandler(endpoint, DecodeDelegationRequest, EncodeDelegationResponse)
	}
	return h
}

// Delegation implements the "Delegation" method in
// sdk_utilitiespb.SdkUtilitiesServer interface.
func (s *Server) Delegation(ctx context.Context, message *sdk_utilitiespb.DelegationRequest) (*sdk_utilitiespb.DelegationResponse, error) {
	ctx = context.WithValue(ctx, goa.MethodKey, "delegation")
	ctx = context.WithValue(ctx, goa.ServiceKey, "sdk-utilities")
	resp, err := s.DelegationH.Handle(ctx, message)
	if err != nil {
		return nil, goagrpc.EncodeError(err)
	}
	return resp.(*sdk_utilitiespb.DelegationResponse), nil
}

// NewIbcChannelHandler creates a gRPC handler which serves the "sdk-utilities"
// service "ibc_channel" endpoint.
func NewIbcChannelHandler(endpoint goa.Endpoint, h goagrpc.UnaryHandler) goagrpc.UnaryHandler {
	if h == nil {
		h = goagrpc.NewUnaryHandler(endpoint, DecodeIbcChannelRequest, EncodeIbcChannelResponse)
	}
	return h
}

// IbcChannel implements the "IbcChannel" method in
// sdk_utilitiespb.SdkUtilitiesServer interface.
func (s *Server) IbcChannel(ctx context.Context, message *sdk_utilitiespb.IbcChannelRequest) (*sdk_utilitiespb.IbcChannelResponse, error) {
	ctx = context.WithValue(ctx, goa.MethodKey, "ibc_channel")
	ctx = context.WithValue(ctx, goa.ServiceKey, "sdk-utilities")
	resp, err := s.IbcChannelH.Handle(ctx, message)
	if err != nil {
		return nil, goagrpc.EncodeError(err)
	}
	return resp.(*sdk_utilitiespb.IbcChannelResponse), nil
}

// NewIbcClientStateHandler creates a gRPC handler which serves the
// "sdk-utilities" service "ibc_client_state" endpoint.
func NewIbcClientStateHandler(endpoint goa.Endpoint, h goagrpc.UnaryHandler) goagrpc.UnaryHandler {
	if h == nil {
		h = goagrpc.NewUnaryHandler(endpoint, DecodeIbcClientStateRequest, EncodeIbcClientStateResponse)
	}
	return h
}

// IbcClientState implements the "IbcClientState" method in
// sdk_utilitiespb.SdkUtilitiesServer interface.
func (s *Server) IbcClientState(ctx context.Context, message *sdk_utilitiespb.IbcClientStateRequest) (*sdk_utilitiespb.IbcClientStateResponse, error) {
	ctx = context.WithValue(ctx, goa.MethodKey, "ibc_client_state")
	ctx = context.WithValue(ctx, goa.ServiceKey, "sdk-utilities")
	resp, err := s.IbcClientStateH.Handle(ctx, message)
	if err != nil {
		return nil, goagrpc.EncodeError(err)
	}
	return resp.(*sdk_utilitiespb.IbcClientStateResponse), nil
}

// NewIbcConnectionHandler creates a gRPC handler which serves the
// "sdk-utilities" service "ibc_connection" endpoint.
func NewIbcConnectionHandler(endpoint goa.Endpoint, h goagrpc.UnaryHandler) goagrpc.UnaryHandler {
	if h == nil {
		h = goagrpc.NewUnaryHandler(endpoint, DecodeIbcConnectionRequest, EncodeIbcConnectionResponse)
	}
	return h
}

// IbcConnection implements the "IbcConnection" method in
// sdk_utilitiespb.SdkUtilitiesServer interface.
func (s *Server) IbcConnection(ctx context.Context, message *sdk_utilitiespb.IbcConnectionRequest) (*sdk_utilitiespb.IbcConnectionResponse, error) {
	ctx = context.WithValue(ctx, goa.MethodKey, "ibc_connection")
	ctx = context.WithValue(ctx, goa.ServiceKey, "sdk-utilities")
	resp, err := s.IbcConnectionH.Handle(ctx, message)
	if err != nil {
		return nil, goagrpc.EncodeError(err)
	}
	return resp.(*sdk_utilitiespb.IbcConnectionResponse), nil
}

// NewIbcDenomTraceHandler creates a gRPC handler which serves the
// "sdk-utilities" service "ibc_denom_trace" endpoint.
func NewIbcDenomTraceHandler(endpoint goa.Endpoint, h goagrpc.UnaryHandler) goagrpc.UnaryHandler {
	if h == nil {
		h = goagrpc.NewUnaryHandler(endpoint, DecodeIbcDenomTraceRequest, EncodeIbcDenomTraceResponse)
	}
	return h
}

// IbcDenomTrace implements the "IbcDenomTrace" method in
// sdk_utilitiespb.SdkUtilitiesServer interface.
func (s *Server) IbcDenomTrace(ctx context.Context, message *sdk_utilitiespb.IbcDenomTraceRequest) (*sdk_utilitiespb.IbcDenomTraceResponse, error) {
	ctx = context.WithValue(ctx, goa.MethodKey, "ibc_denom_trace")
	ctx = context.WithValue(ctx, goa.ServiceKey, "sdk-utilities")
	resp, err := s.IbcDenomTraceH.Handle(ctx, message)
	if err != nil {
		return nil, goagrpc.EncodeError(err)
	}
	return resp.(*sdk_utilitiespb.IbcDenomTraceResponse), nil
}
