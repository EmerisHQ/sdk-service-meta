// Code generated by goa v3.5.2, DO NOT EDIT.
//
// sdk-utilities client
//
// Command:
// $ goa gen github.com/allinbits/sdk-service-meta

package sdkutilities

import (
	"context"

	goa "goa.design/goa/v3/pkg"
)

// Client is the "sdk-utilities" service client.
type Client struct {
	SupplyEndpoint      goa.Endpoint
	QueryTxEndpoint     goa.Endpoint
	BroadcastTxEndpoint goa.Endpoint
	TxMetadataEndpoint  goa.Endpoint
}

// NewClient initializes a "sdk-utilities" service client given the endpoints.
func NewClient(supply, queryTx, broadcastTx, txMetadata goa.Endpoint) *Client {
	return &Client{
		SupplyEndpoint:      supply,
		QueryTxEndpoint:     queryTx,
		BroadcastTxEndpoint: broadcastTx,
		TxMetadataEndpoint:  txMetadata,
	}
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
