// Code generated by goa v3.5.3, DO NOT EDIT.
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

// AccountNumbers calls the "AccountNumbers" function in
// sdk_utilitiespb.SdkUtilitiesClient interface.
func (c *Client) AccountNumbers() goa.Endpoint {
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		inv := goagrpc.NewInvoker(
			BuildAccountNumbersFunc(c.grpccli, c.opts...),
			EncodeAccountNumbersRequest,
			DecodeAccountNumbersResponse)
		res, err := inv.Invoke(ctx, v)
		if err != nil {
			return nil, goa.Fault(err.Error())
		}
		return res, nil
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

// Block calls the "Block" function in sdk_utilitiespb.SdkUtilitiesClient
// interface.
func (c *Client) Block() goa.Endpoint {
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		inv := goagrpc.NewInvoker(
			BuildBlockFunc(c.grpccli, c.opts...),
			EncodeBlockRequest,
			DecodeBlockResponse)
		res, err := inv.Invoke(ctx, v)
		if err != nil {
			return nil, goa.Fault(err.Error())
		}
		return res, nil
	}
}

// LiquidityParams calls the "LiquidityParams" function in
// sdk_utilitiespb.SdkUtilitiesClient interface.
func (c *Client) LiquidityParams() goa.Endpoint {
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		inv := goagrpc.NewInvoker(
			BuildLiquidityParamsFunc(c.grpccli, c.opts...),
			EncodeLiquidityParamsRequest,
			DecodeLiquidityParamsResponse)
		res, err := inv.Invoke(ctx, v)
		if err != nil {
			return nil, goa.Fault(err.Error())
		}
		return res, nil
	}
}

// LiquidityPools calls the "LiquidityPools" function in
// sdk_utilitiespb.SdkUtilitiesClient interface.
func (c *Client) LiquidityPools() goa.Endpoint {
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		inv := goagrpc.NewInvoker(
			BuildLiquidityPoolsFunc(c.grpccli, c.opts...),
			EncodeLiquidityPoolsRequest,
			DecodeLiquidityPoolsResponse)
		res, err := inv.Invoke(ctx, v)
		if err != nil {
			return nil, goa.Fault(err.Error())
		}
		return res, nil
	}
}

// MintInflation calls the "MintInflation" function in
// sdk_utilitiespb.SdkUtilitiesClient interface.
func (c *Client) MintInflation() goa.Endpoint {
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		inv := goagrpc.NewInvoker(
			BuildMintInflationFunc(c.grpccli, c.opts...),
			EncodeMintInflationRequest,
			DecodeMintInflationResponse)
		res, err := inv.Invoke(ctx, v)
		if err != nil {
			return nil, goa.Fault(err.Error())
		}
		return res, nil
	}
}

// MintParams calls the "MintParams" function in
// sdk_utilitiespb.SdkUtilitiesClient interface.
func (c *Client) MintParams() goa.Endpoint {
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		inv := goagrpc.NewInvoker(
			BuildMintParamsFunc(c.grpccli, c.opts...),
			EncodeMintParamsRequest,
			DecodeMintParamsResponse)
		res, err := inv.Invoke(ctx, v)
		if err != nil {
			return nil, goa.Fault(err.Error())
		}
		return res, nil
	}
}

// MintAnnualProvision calls the "MintAnnualProvision" function in
// sdk_utilitiespb.SdkUtilitiesClient interface.
func (c *Client) MintAnnualProvision() goa.Endpoint {
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		inv := goagrpc.NewInvoker(
			BuildMintAnnualProvisionFunc(c.grpccli, c.opts...),
			EncodeMintAnnualProvisionRequest,
			DecodeMintAnnualProvisionResponse)
		res, err := inv.Invoke(ctx, v)
		if err != nil {
			return nil, goa.Fault(err.Error())
		}
		return res, nil
	}
}

// DelegatorRewards calls the "DelegatorRewards" function in
// sdk_utilitiespb.SdkUtilitiesClient interface.
func (c *Client) DelegatorRewards() goa.Endpoint {
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		inv := goagrpc.NewInvoker(
			BuildDelegatorRewardsFunc(c.grpccli, c.opts...),
			EncodeDelegatorRewardsRequest,
			DecodeDelegatorRewardsResponse)
		res, err := inv.Invoke(ctx, v)
		if err != nil {
			return nil, goa.Fault(err.Error())
		}
		return res, nil
	}
}

// TxFeeEstimate calls the "TxFeeEstimate" function in
// sdk_utilitiespb.SdkUtilitiesClient interface.
func (c *Client) TxFeeEstimate() goa.Endpoint {
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		inv := goagrpc.NewInvoker(
			BuildTxFeeEstimateFunc(c.grpccli, c.opts...),
			EncodeTxFeeEstimateRequest,
			DecodeTxFeeEstimateResponse)
		res, err := inv.Invoke(ctx, v)
		if err != nil {
			return nil, goa.Fault(err.Error())
		}
		return res, nil
	}
}
