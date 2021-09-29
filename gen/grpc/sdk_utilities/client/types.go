// Code generated by goa v3.5.2, DO NOT EDIT.
//
// sdk-utilities gRPC client types
//
// Command:
// $ goa gen github.com/allinbits/sdk-service-meta

package client

import (
	sdk_utilitiespb "github.com/allinbits/sdk-service-meta/gen/grpc/sdk_utilities/pb"
	sdkutilities "github.com/allinbits/sdk-service-meta/gen/sdk_utilities"
	goa "goa.design/goa/v3/pkg"
)

// NewSupplyRequest builds the gRPC request type from the payload of the
// "supply" endpoint of the "sdk-utilities" service.
func NewSupplyRequest(payload *sdkutilities.SupplyPayload) *sdk_utilitiespb.SupplyRequest {
	message := &sdk_utilitiespb.SupplyRequest{
		ChainName: payload.ChainName,
	}
	if payload.Port != nil {
		message.Port = int32(*payload.Port)
	}
	return message
}

// NewSupplyResult builds the result type of the "supply" endpoint of the
// "sdk-utilities" service from the gRPC response type.
func NewSupplyResult(message *sdk_utilitiespb.SupplyResponse) *sdkutilities.Supply2 {
	result := &sdkutilities.Supply2{}
	if message.Coins != nil {
		result.Coins = make([]*sdkutilities.Coin, len(message.Coins))
		for i, val := range message.Coins {
			result.Coins[i] = &sdkutilities.Coin{
				Denom:  val.Denom,
				Amount: val.Amount,
			}
		}
	}
	return result
}

// NewQueryTxRequest builds the gRPC request type from the payload of the
// "queryTx" endpoint of the "sdk-utilities" service.
func NewQueryTxRequest(payload *sdkutilities.QueryTxPayload) *sdk_utilitiespb.QueryTxRequest {
	message := &sdk_utilitiespb.QueryTxRequest{
		ChainName: payload.ChainName,
		Hash:      payload.Hash,
	}
	if payload.Port != nil {
		message.Port = int32(*payload.Port)
	}
	return message
}

// NewQueryTxResult builds the result type of the "queryTx" endpoint of the
// "sdk-utilities" service from the gRPC response type.
func NewQueryTxResult(message *sdk_utilitiespb.QueryTxResponse) []byte {
	result := message.Field
	return result
}

// ValidateSupplyResponse runs the validations defined on SupplyResponse.
func ValidateSupplyResponse(message *sdk_utilitiespb.SupplyResponse) (err error) {
	if message.Coins == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("coins", "message"))
	}
	return
}

// ValidateCoin runs the validations defined on Coin.
func ValidateCoin(message *sdk_utilitiespb.Coin) (err error) {

	return
}
