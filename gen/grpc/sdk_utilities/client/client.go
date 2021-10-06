// Code generated by goa v3.5.2, DO NOT EDIT.
//
// sdk-utilities gRPC client
//
// Command:
// $ goa gen github.com/allinbits/sdk-service-meta

package client

import (
	"context"

	sdk_utilitiespb "github.com/allinbits/sdk-service-meta/gen/grpc/sdk_utilities/pb"
	goagrpc "goa.design/goa/v3/grpc"
	goapb "goa.design/goa/v3/grpc/pb"
	goa "goa.design/goa/v3/pkg"
	"google.golang.org/grpc"
)

// Client lists the service endpoint gRPC clients.
type Client struct {
	grpccli sdk_utilitiespb.SdkUtilitiesClient
	opts    []grpc.CallOption
}

// NewClient instantiates gRPC client for all the sdk-utilities service servers.
func NewClient(cc *grpc.ClientConn, opts ...grpc.CallOption) *Client {
	return &Client{
		grpccli: sdk_utilitiespb.NewSdkUtilitiesClient(cc),
		opts:    opts,
	}
}

// Supply calls the "Supply" function in sdk_utilitiespb.SdkUtilitiesClient
// interface.
func (c *Client) Supply() goa.Endpoint {
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		inv := goagrpc.NewInvoker(
			BuildSupplyFunc(c.grpccli, c.opts...),
			EncodeSupplyRequest,
			DecodeSupplyResponse)
		res, err := inv.Invoke(ctx, v)
		if err != nil {
			return nil, goa.Fault(err.Error())
		}
		return res, nil
	}
}

// QueryTx calls the "QueryTx" function in sdk_utilitiespb.SdkUtilitiesClient
// interface.
func (c *Client) QueryTx() goa.Endpoint {
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		inv := goagrpc.NewInvoker(
			BuildQueryTxFunc(c.grpccli, c.opts...),
			EncodeQueryTxRequest,
			DecodeQueryTxResponse)
		res, err := inv.Invoke(ctx, v)
		if err != nil {
			return nil, goa.Fault(err.Error())
		}
		return res, nil
	}
}

// BroadcastTx calls the "BroadcastTx" function in
// sdk_utilitiespb.SdkUtilitiesClient interface.
func (c *Client) BroadcastTx() goa.Endpoint {
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		inv := goagrpc.NewInvoker(
			BuildBroadcastTxFunc(c.grpccli, c.opts...),
			EncodeBroadcastTxRequest,
			DecodeBroadcastTxResponse)
		res, err := inv.Invoke(ctx, v)
		if err != nil {
			return nil, goa.Fault(err.Error())
		}
		return res, nil
	}
}

// TxMetadata calls the "TxMetadata" function in
// sdk_utilitiespb.SdkUtilitiesClient interface.
func (c *Client) TxMetadata() goa.Endpoint {
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		inv := goagrpc.NewInvoker(
			BuildTxMetadataFunc(c.grpccli, c.opts...),
			EncodeTxMetadataRequest,
			DecodeTxMetadataResponse)
		res, err := inv.Invoke(ctx, v)
		if err != nil {
			return nil, goa.Fault(err.Error())
		}
		return res, nil
	}
}

// AuthEndpoint calls the "AuthEndpoint" function in
// sdk_utilitiespb.SdkUtilitiesClient interface.
func (c *Client) AuthEndpoint() goa.Endpoint {
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		inv := goagrpc.NewInvoker(
			BuildAuthEndpointFunc(c.grpccli, c.opts...),
			EncodeAuthEndpointRequest,
			DecodeAuthEndpointResponse)
		res, err := inv.Invoke(ctx, v)
		if err != nil {
			resp := goagrpc.DecodeError(err)
			switch message := resp.(type) {
			case *sdk_utilitiespb.AuthProcessingErrorError:
				return nil, NewAuthProcessingErrorError(message)
			case *goapb.ErrorResponse:
				return nil, goagrpc.NewServiceError(message)
			default:
				return nil, goa.Fault(err.Error())
			}
		}
		return res, nil
	}
}

// Bank calls the "Bank" function in sdk_utilitiespb.SdkUtilitiesClient
// interface.
func (c *Client) Bank() goa.Endpoint {
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		inv := goagrpc.NewInvoker(
			BuildBankFunc(c.grpccli, c.opts...),
			EncodeBankRequest,
			DecodeBankResponse)
		res, err := inv.Invoke(ctx, v)
		if err != nil {
			resp := goagrpc.DecodeError(err)
			switch message := resp.(type) {
			case *sdk_utilitiespb.BankProcessingErrorError:
				return nil, NewBankProcessingErrorError(message)
			case *goapb.ErrorResponse:
				return nil, goagrpc.NewServiceError(message)
			default:
				return nil, goa.Fault(err.Error())
			}
		}
		return res, nil
	}
}

// DelegationEndpoint calls the "DelegationEndpoint" function in
// sdk_utilitiespb.SdkUtilitiesClient interface.
func (c *Client) DelegationEndpoint() goa.Endpoint {
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		inv := goagrpc.NewInvoker(
			BuildDelegationEndpointFunc(c.grpccli, c.opts...),
			EncodeDelegationEndpointRequest,
			DecodeDelegationEndpointResponse)
		res, err := inv.Invoke(ctx, v)
		if err != nil {
			resp := goagrpc.DecodeError(err)
			switch message := resp.(type) {
			case *sdk_utilitiespb.DelegationProcessingErrorError:
				return nil, NewDelegationProcessingErrorError(message)
			case *goapb.ErrorResponse:
				return nil, goagrpc.NewServiceError(message)
			default:
				return nil, goa.Fault(err.Error())
			}
		}
		return res, nil
	}
}

// IbcChannel calls the "IbcChannel" function in
// sdk_utilitiespb.SdkUtilitiesClient interface.
func (c *Client) IbcChannel() goa.Endpoint {
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		inv := goagrpc.NewInvoker(
			BuildIbcChannelFunc(c.grpccli, c.opts...),
			EncodeIbcChannelRequest,
			DecodeIbcChannelResponse)
		res, err := inv.Invoke(ctx, v)
		if err != nil {
			resp := goagrpc.DecodeError(err)
			switch message := resp.(type) {
			case *sdk_utilitiespb.IbcChannelProcessingErrorError:
				return nil, NewIbcChannelProcessingErrorError(message)
			case *goapb.ErrorResponse:
				return nil, goagrpc.NewServiceError(message)
			default:
				return nil, goa.Fault(err.Error())
			}
		}
		return res, nil
	}
}

// IbcClientState calls the "IbcClientState" function in
// sdk_utilitiespb.SdkUtilitiesClient interface.
func (c *Client) IbcClientState() goa.Endpoint {
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		inv := goagrpc.NewInvoker(
			BuildIbcClientStateFunc(c.grpccli, c.opts...),
			EncodeIbcClientStateRequest,
			DecodeIbcClientStateResponse)
		res, err := inv.Invoke(ctx, v)
		if err != nil {
			resp := goagrpc.DecodeError(err)
			switch message := resp.(type) {
			case *sdk_utilitiespb.IbcClientStateProcessingErrorError:
				return nil, NewIbcClientStateProcessingErrorError(message)
			case *goapb.ErrorResponse:
				return nil, goagrpc.NewServiceError(message)
			default:
				return nil, goa.Fault(err.Error())
			}
		}
		return res, nil
	}
}

// IbcConnection calls the "IbcConnection" function in
// sdk_utilitiespb.SdkUtilitiesClient interface.
func (c *Client) IbcConnection() goa.Endpoint {
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		inv := goagrpc.NewInvoker(
			BuildIbcConnectionFunc(c.grpccli, c.opts...),
			EncodeIbcConnectionRequest,
			DecodeIbcConnectionResponse)
		res, err := inv.Invoke(ctx, v)
		if err != nil {
			resp := goagrpc.DecodeError(err)
			switch message := resp.(type) {
			case *sdk_utilitiespb.IbcConnectionProcessingErrorError:
				return nil, NewIbcConnectionProcessingErrorError(message)
			case *goapb.ErrorResponse:
				return nil, goagrpc.NewServiceError(message)
			default:
				return nil, goa.Fault(err.Error())
			}
		}
		return res, nil
	}
}

// IbcDenomTrace calls the "IbcDenomTrace" function in
// sdk_utilitiespb.SdkUtilitiesClient interface.
func (c *Client) IbcDenomTrace() goa.Endpoint {
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		inv := goagrpc.NewInvoker(
			BuildIbcDenomTraceFunc(c.grpccli, c.opts...),
			EncodeIbcDenomTraceRequest,
			DecodeIbcDenomTraceResponse)
		res, err := inv.Invoke(ctx, v)
		if err != nil {
			resp := goagrpc.DecodeError(err)
			switch message := resp.(type) {
			case *sdk_utilitiespb.IbcDenomTraceProcessingErrorError:
				return nil, NewIbcDenomTraceProcessingErrorError(message)
			case *goapb.ErrorResponse:
				return nil, goagrpc.NewServiceError(message)
			default:
				return nil, goa.Fault(err.Error())
			}
		}
		return res, nil
	}
}
