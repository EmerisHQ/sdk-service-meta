// Code generated by goa v3.5.2, DO NOT EDIT.
//
// sdk-utilities gRPC server types
//
// Command:
// $ goa gen github.com/allinbits/sdk-service-meta

package server

import (
	sdk_utilitiespb "github.com/allinbits/sdk-service-meta/gen/grpc/sdk_utilities/pb"
	sdkutilities "github.com/allinbits/sdk-service-meta/gen/sdk_utilities"
)

// NewSupplyPayload builds the payload of the "supply" endpoint of the
// "sdk-utilities" service from the gRPC request type.
func NewSupplyPayload(message *sdk_utilitiespb.SupplyRequest) *sdkutilities.SupplyPayload {
	v := &sdkutilities.SupplyPayload{
		ChainName: message.ChainName,
	}
	if message.Port != 0 {
		portptr := int(message.Port)
		v.Port = &portptr
	}
	return v
}

// NewSupplyResponse builds the gRPC response type from the result of the
// "supply" endpoint of the "sdk-utilities" service.
func NewSupplyResponse(result *sdkutilities.Supply2) *sdk_utilitiespb.SupplyResponse {
	message := &sdk_utilitiespb.SupplyResponse{}
	if result.Coins != nil {
		message.Coins = make([]*sdk_utilitiespb.Coin, len(result.Coins))
		for i, val := range result.Coins {
			message.Coins[i] = &sdk_utilitiespb.Coin{
				Denom:  val.Denom,
				Amount: val.Amount,
			}
		}
	}
	return message
}
