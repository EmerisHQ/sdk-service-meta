// Code generated by goa v3.5.2, DO NOT EDIT.
//
// sdk-utilities gRPC client CLI support package
//
// Command:
// $ goa gen github.com/allinbits/sdk-service-meta

package client

import (
	"encoding/json"
	"fmt"

	sdk_utilitiespb "github.com/allinbits/sdk-service-meta/gen/grpc/sdk_utilities/pb"
	sdkutilities "github.com/allinbits/sdk-service-meta/gen/sdk_utilities"
)

// BuildSupplyPayload builds the payload for the sdk-utilities supply endpoint
// from CLI flags.
func BuildSupplyPayload(sdkUtilitiesSupplyMessage string) (*sdkutilities.SupplyPayload, error) {
	var err error
	var message sdk_utilitiespb.SupplyRequest
	{
		if sdkUtilitiesSupplyMessage != "" {
			err = json.Unmarshal([]byte(sdkUtilitiesSupplyMessage), &message)
			if err != nil {
				return nil, fmt.Errorf("invalid JSON for message, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"chainName\": \"Molestias aut sit nobis et sit temporibus.\",\n      \"port\": 4276607318535473132\n   }'")
			}
		}
	}
	v := &sdkutilities.SupplyPayload{
		ChainName: message.ChainName,
	}
	if message.Port != 0 {
		portptr := int(message.Port)
		v.Port = &portptr
	}

	return v, nil
}

// BuildQueryTxPayload builds the payload for the sdk-utilities queryTx
// endpoint from CLI flags.
func BuildQueryTxPayload(sdkUtilitiesQueryTxMessage string) (*sdkutilities.QueryTxPayload, error) {
	var err error
	var message sdk_utilitiespb.QueryTxRequest
	{
		if sdkUtilitiesQueryTxMessage != "" {
			err = json.Unmarshal([]byte(sdkUtilitiesQueryTxMessage), &message)
			if err != nil {
				return nil, fmt.Errorf("invalid JSON for message, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"chainName\": \"Suscipit rerum corrupti.\",\n      \"hash\": \"Vitae porro architecto dolorem.\",\n      \"port\": 8026972786014809817\n   }'")
			}
		}
	}
	v := &sdkutilities.QueryTxPayload{
		ChainName: message.ChainName,
		Hash:      message.Hash,
	}
	if message.Port != 0 {
		portptr := int(message.Port)
		v.Port = &portptr
	}

	return v, nil
}

// BuildBroadcastTxPayload builds the payload for the sdk-utilities broadcastTx
// endpoint from CLI flags.
func BuildBroadcastTxPayload(sdkUtilitiesBroadcastTxMessage string) (*sdkutilities.BroadcastTxPayload, error) {
	var err error
	var message sdk_utilitiespb.BroadcastTxRequest
	{
		if sdkUtilitiesBroadcastTxMessage != "" {
			err = json.Unmarshal([]byte(sdkUtilitiesBroadcastTxMessage), &message)
			if err != nil {
				return nil, fmt.Errorf("invalid JSON for message, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"chainName\": \"Voluptas magnam voluptatem dolore quae quasi quia.\",\n      \"port\": 5026848979338219803,\n      \"txBytes\": \"QmVhdGFlIHNlZCBjb25zZXF1dW50dXIgYmVhdGFlLg==\"\n   }'")
			}
		}
	}
	v := &sdkutilities.BroadcastTxPayload{
		ChainName: message.ChainName,
		TxBytes:   message.TxBytes,
	}
	if message.Port != 0 {
		portptr := int(message.Port)
		v.Port = &portptr
	}

	return v, nil
}

// BuildTxMetadataPayload builds the payload for the sdk-utilities txMetadata
// endpoint from CLI flags.
func BuildTxMetadataPayload(sdkUtilitiesTxMetadataMessage string) (*sdkutilities.TxMetadataPayload, error) {
	var err error
	var message sdk_utilitiespb.TxMetadataRequest
	{
		if sdkUtilitiesTxMetadataMessage != "" {
			err = json.Unmarshal([]byte(sdkUtilitiesTxMetadataMessage), &message)
			if err != nil {
				return nil, fmt.Errorf("invalid JSON for message, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"txBytes\": \"SXBzdW0gZXZlbmlldCBlbmltIHF1aWEgZG9sb3J1bSBhbWV0IGVvcy4=\"\n   }'")
			}
		}
	}
	v := &sdkutilities.TxMetadataPayload{
		TxBytes: message.TxBytes,
	}

	return v, nil
}
