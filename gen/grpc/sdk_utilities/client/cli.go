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
				return nil, fmt.Errorf("invalid JSON for message, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"chainName\": \"Sed enim atque dolores esse dicta.\",\n      \"port\": 7124529205354848620\n   }'")
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
				return nil, fmt.Errorf("invalid JSON for message, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"chainName\": \"Explicabo consequatur est ea est.\",\n      \"hash\": \"Laborum quam ducimus consequatur.\",\n      \"port\": 3065905437504226774\n   }'")
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
				return nil, fmt.Errorf("invalid JSON for message, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"chainName\": \"Officiis alias.\",\n      \"port\": 5526397674056279747,\n      \"txBytes\": \"UmVtIHZvbHVwdGFzIHNlZCBpZCBhbWV0IGRpY3RhLg==\"\n   }'")
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
				return nil, fmt.Errorf("invalid JSON for message, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"txBytes\": \"UXVpcyBjdWxwYSBldCBibGFuZGl0aWlzIHZlcml0YXRpcyBoYXJ1bSBwb3NzaW11cy4=\"\n   }'")
			}
		}
	}
	v := &sdkutilities.TxMetadataPayload{
		TxBytes: message.TxBytes,
	}

	return v, nil
}

// BuildAuthEndpointPayload builds the payload for the sdk-utilities auth
// endpoint from CLI flags.
func BuildAuthEndpointPayload(sdkUtilitiesAuthMessage string) (*sdkutilities.AuthPayload, error) {
	var err error
	var message sdk_utilitiespb.AuthRequest
	{
		if sdkUtilitiesAuthMessage != "" {
			err = json.Unmarshal([]byte(sdkUtilitiesAuthMessage), &message)
			if err != nil {
				return nil, fmt.Errorf("invalid JSON for message, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"payload\": [\n         {\n            \"key\": \"QXV0IHNpdCBub2JpcyBldCBzaXQu\",\n            \"operationType\": \"delete\",\n            \"value\": \"Q3VwaWRpdGF0ZSByZXByZWhlbmRlcml0IHF1aWEgc3VzY2lwaXQgcmVydW0gY29ycnVwdGku\"\n         },\n         {\n            \"key\": \"QXV0IHNpdCBub2JpcyBldCBzaXQu\",\n            \"operationType\": \"delete\",\n            \"value\": \"Q3VwaWRpdGF0ZSByZXByZWhlbmRlcml0IHF1aWEgc3VzY2lwaXQgcmVydW0gY29ycnVwdGku\"\n         }\n      ]\n   }'")
			}
		}
	}
	v := &sdkutilities.AuthPayload{}
	if message.Payload != nil {
		v.Payload = make([]*sdkutilities.TracePayload, len(message.Payload))
		for i, val := range message.Payload {
			v.Payload[i] = &sdkutilities.TracePayload{
				Key:   val.Key,
				Value: val.Value,
			}
			if val.OperationType != "" {
				v.Payload[i].OperationType = &val.OperationType
			}
		}
	}

	return v, nil
}

// BuildBankPayload builds the payload for the sdk-utilities bank endpoint from
// CLI flags.
func BuildBankPayload(sdkUtilitiesBankMessage string) (*sdkutilities.BankPayload, error) {
	var err error
	var message sdk_utilitiespb.BankRequest
	{
		if sdkUtilitiesBankMessage != "" {
			err = json.Unmarshal([]byte(sdkUtilitiesBankMessage), &message)
			if err != nil {
				return nil, fmt.Errorf("invalid JSON for message, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"payload\": [\n         {\n            \"key\": \"QXV0IHNpdCBub2JpcyBldCBzaXQu\",\n            \"operationType\": \"delete\",\n            \"value\": \"Q3VwaWRpdGF0ZSByZXByZWhlbmRlcml0IHF1aWEgc3VzY2lwaXQgcmVydW0gY29ycnVwdGku\"\n         },\n         {\n            \"key\": \"QXV0IHNpdCBub2JpcyBldCBzaXQu\",\n            \"operationType\": \"delete\",\n            \"value\": \"Q3VwaWRpdGF0ZSByZXByZWhlbmRlcml0IHF1aWEgc3VzY2lwaXQgcmVydW0gY29ycnVwdGku\"\n         },\n         {\n            \"key\": \"QXV0IHNpdCBub2JpcyBldCBzaXQu\",\n            \"operationType\": \"delete\",\n            \"value\": \"Q3VwaWRpdGF0ZSByZXByZWhlbmRlcml0IHF1aWEgc3VzY2lwaXQgcmVydW0gY29ycnVwdGku\"\n         },\n         {\n            \"key\": \"QXV0IHNpdCBub2JpcyBldCBzaXQu\",\n            \"operationType\": \"delete\",\n            \"value\": \"Q3VwaWRpdGF0ZSByZXByZWhlbmRlcml0IHF1aWEgc3VzY2lwaXQgcmVydW0gY29ycnVwdGku\"\n         }\n      ]\n   }'")
			}
		}
	}
	v := &sdkutilities.BankPayload{}
	if message.Payload != nil {
		v.Payload = make([]*sdkutilities.TracePayload, len(message.Payload))
		for i, val := range message.Payload {
			v.Payload[i] = &sdkutilities.TracePayload{
				Key:   val.Key,
				Value: val.Value,
			}
			if val.OperationType != "" {
				v.Payload[i].OperationType = &val.OperationType
			}
		}
	}

	return v, nil
}

// BuildDelegationEndpointPayload builds the payload for the sdk-utilities
// delegation endpoint from CLI flags.
func BuildDelegationEndpointPayload(sdkUtilitiesDelegationMessage string) (*sdkutilities.DelegationPayload, error) {
	var err error
	var message sdk_utilitiespb.DelegationRequest
	{
		if sdkUtilitiesDelegationMessage != "" {
			err = json.Unmarshal([]byte(sdkUtilitiesDelegationMessage), &message)
			if err != nil {
				return nil, fmt.Errorf("invalid JSON for message, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"payload\": [\n         {\n            \"key\": \"QXV0IHNpdCBub2JpcyBldCBzaXQu\",\n            \"operationType\": \"delete\",\n            \"value\": \"Q3VwaWRpdGF0ZSByZXByZWhlbmRlcml0IHF1aWEgc3VzY2lwaXQgcmVydW0gY29ycnVwdGku\"\n         },\n         {\n            \"key\": \"QXV0IHNpdCBub2JpcyBldCBzaXQu\",\n            \"operationType\": \"delete\",\n            \"value\": \"Q3VwaWRpdGF0ZSByZXByZWhlbmRlcml0IHF1aWEgc3VzY2lwaXQgcmVydW0gY29ycnVwdGku\"\n         }\n      ]\n   }'")
			}
		}
	}
	v := &sdkutilities.DelegationPayload{}
	if message.Payload != nil {
		v.Payload = make([]*sdkutilities.TracePayload, len(message.Payload))
		for i, val := range message.Payload {
			v.Payload[i] = &sdkutilities.TracePayload{
				Key:   val.Key,
				Value: val.Value,
			}
			if val.OperationType != "" {
				v.Payload[i].OperationType = &val.OperationType
			}
		}
	}

	return v, nil
}

// BuildIbcChannelPayload builds the payload for the sdk-utilities ibc_channel
// endpoint from CLI flags.
func BuildIbcChannelPayload(sdkUtilitiesIbcChannelMessage string) (*sdkutilities.IbcChannelPayload, error) {
	var err error
	var message sdk_utilitiespb.IbcChannelRequest
	{
		if sdkUtilitiesIbcChannelMessage != "" {
			err = json.Unmarshal([]byte(sdkUtilitiesIbcChannelMessage), &message)
			if err != nil {
				return nil, fmt.Errorf("invalid JSON for message, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"payload\": [\n         {\n            \"key\": \"QXV0IHNpdCBub2JpcyBldCBzaXQu\",\n            \"operationType\": \"delete\",\n            \"value\": \"Q3VwaWRpdGF0ZSByZXByZWhlbmRlcml0IHF1aWEgc3VzY2lwaXQgcmVydW0gY29ycnVwdGku\"\n         },\n         {\n            \"key\": \"QXV0IHNpdCBub2JpcyBldCBzaXQu\",\n            \"operationType\": \"delete\",\n            \"value\": \"Q3VwaWRpdGF0ZSByZXByZWhlbmRlcml0IHF1aWEgc3VzY2lwaXQgcmVydW0gY29ycnVwdGku\"\n         },\n         {\n            \"key\": \"QXV0IHNpdCBub2JpcyBldCBzaXQu\",\n            \"operationType\": \"delete\",\n            \"value\": \"Q3VwaWRpdGF0ZSByZXByZWhlbmRlcml0IHF1aWEgc3VzY2lwaXQgcmVydW0gY29ycnVwdGku\"\n         },\n         {\n            \"key\": \"QXV0IHNpdCBub2JpcyBldCBzaXQu\",\n            \"operationType\": \"delete\",\n            \"value\": \"Q3VwaWRpdGF0ZSByZXByZWhlbmRlcml0IHF1aWEgc3VzY2lwaXQgcmVydW0gY29ycnVwdGku\"\n         }\n      ]\n   }'")
			}
		}
	}
	v := &sdkutilities.IbcChannelPayload{}
	if message.Payload != nil {
		v.Payload = make([]*sdkutilities.TracePayload, len(message.Payload))
		for i, val := range message.Payload {
			v.Payload[i] = &sdkutilities.TracePayload{
				Key:   val.Key,
				Value: val.Value,
			}
			if val.OperationType != "" {
				v.Payload[i].OperationType = &val.OperationType
			}
		}
	}

	return v, nil
}

// BuildIbcClientStatePayload builds the payload for the sdk-utilities
// ibc_client_state endpoint from CLI flags.
func BuildIbcClientStatePayload(sdkUtilitiesIbcClientStateMessage string) (*sdkutilities.IbcClientStatePayload, error) {
	var err error
	var message sdk_utilitiespb.IbcClientStateRequest
	{
		if sdkUtilitiesIbcClientStateMessage != "" {
			err = json.Unmarshal([]byte(sdkUtilitiesIbcClientStateMessage), &message)
			if err != nil {
				return nil, fmt.Errorf("invalid JSON for message, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"payload\": [\n         {\n            \"key\": \"QXV0IHNpdCBub2JpcyBldCBzaXQu\",\n            \"operationType\": \"delete\",\n            \"value\": \"Q3VwaWRpdGF0ZSByZXByZWhlbmRlcml0IHF1aWEgc3VzY2lwaXQgcmVydW0gY29ycnVwdGku\"\n         },\n         {\n            \"key\": \"QXV0IHNpdCBub2JpcyBldCBzaXQu\",\n            \"operationType\": \"delete\",\n            \"value\": \"Q3VwaWRpdGF0ZSByZXByZWhlbmRlcml0IHF1aWEgc3VzY2lwaXQgcmVydW0gY29ycnVwdGku\"\n         }\n      ]\n   }'")
			}
		}
	}
	v := &sdkutilities.IbcClientStatePayload{}
	if message.Payload != nil {
		v.Payload = make([]*sdkutilities.TracePayload, len(message.Payload))
		for i, val := range message.Payload {
			v.Payload[i] = &sdkutilities.TracePayload{
				Key:   val.Key,
				Value: val.Value,
			}
			if val.OperationType != "" {
				v.Payload[i].OperationType = &val.OperationType
			}
		}
	}

	return v, nil
}

// BuildIbcConnectionPayload builds the payload for the sdk-utilities
// ibc_connection endpoint from CLI flags.
func BuildIbcConnectionPayload(sdkUtilitiesIbcConnectionMessage string) (*sdkutilities.IbcConnectionPayload, error) {
	var err error
	var message sdk_utilitiespb.IbcConnectionRequest
	{
		if sdkUtilitiesIbcConnectionMessage != "" {
			err = json.Unmarshal([]byte(sdkUtilitiesIbcConnectionMessage), &message)
			if err != nil {
				return nil, fmt.Errorf("invalid JSON for message, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"payload\": [\n         {\n            \"key\": \"QXV0IHNpdCBub2JpcyBldCBzaXQu\",\n            \"operationType\": \"delete\",\n            \"value\": \"Q3VwaWRpdGF0ZSByZXByZWhlbmRlcml0IHF1aWEgc3VzY2lwaXQgcmVydW0gY29ycnVwdGku\"\n         },\n         {\n            \"key\": \"QXV0IHNpdCBub2JpcyBldCBzaXQu\",\n            \"operationType\": \"delete\",\n            \"value\": \"Q3VwaWRpdGF0ZSByZXByZWhlbmRlcml0IHF1aWEgc3VzY2lwaXQgcmVydW0gY29ycnVwdGku\"\n         },\n         {\n            \"key\": \"QXV0IHNpdCBub2JpcyBldCBzaXQu\",\n            \"operationType\": \"delete\",\n            \"value\": \"Q3VwaWRpdGF0ZSByZXByZWhlbmRlcml0IHF1aWEgc3VzY2lwaXQgcmVydW0gY29ycnVwdGku\"\n         }\n      ]\n   }'")
			}
		}
	}
	v := &sdkutilities.IbcConnectionPayload{}
	if message.Payload != nil {
		v.Payload = make([]*sdkutilities.TracePayload, len(message.Payload))
		for i, val := range message.Payload {
			v.Payload[i] = &sdkutilities.TracePayload{
				Key:   val.Key,
				Value: val.Value,
			}
			if val.OperationType != "" {
				v.Payload[i].OperationType = &val.OperationType
			}
		}
	}

	return v, nil
}

// BuildIbcDenomTracePayload builds the payload for the sdk-utilities
// ibc_denom_trace endpoint from CLI flags.
func BuildIbcDenomTracePayload(sdkUtilitiesIbcDenomTraceMessage string) (*sdkutilities.IbcDenomTracePayload, error) {
	var err error
	var message sdk_utilitiespb.IbcDenomTraceRequest
	{
		if sdkUtilitiesIbcDenomTraceMessage != "" {
			err = json.Unmarshal([]byte(sdkUtilitiesIbcDenomTraceMessage), &message)
			if err != nil {
				return nil, fmt.Errorf("invalid JSON for message, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"payload\": [\n         {\n            \"key\": \"QXV0IHNpdCBub2JpcyBldCBzaXQu\",\n            \"operationType\": \"delete\",\n            \"value\": \"Q3VwaWRpdGF0ZSByZXByZWhlbmRlcml0IHF1aWEgc3VzY2lwaXQgcmVydW0gY29ycnVwdGku\"\n         },\n         {\n            \"key\": \"QXV0IHNpdCBub2JpcyBldCBzaXQu\",\n            \"operationType\": \"delete\",\n            \"value\": \"Q3VwaWRpdGF0ZSByZXByZWhlbmRlcml0IHF1aWEgc3VzY2lwaXQgcmVydW0gY29ycnVwdGku\"\n         },\n         {\n            \"key\": \"QXV0IHNpdCBub2JpcyBldCBzaXQu\",\n            \"operationType\": \"delete\",\n            \"value\": \"Q3VwaWRpdGF0ZSByZXByZWhlbmRlcml0IHF1aWEgc3VzY2lwaXQgcmVydW0gY29ycnVwdGku\"\n         }\n      ]\n   }'")
			}
		}
	}
	v := &sdkutilities.IbcDenomTracePayload{}
	if message.Payload != nil {
		v.Payload = make([]*sdkutilities.TracePayload, len(message.Payload))
		for i, val := range message.Payload {
			v.Payload[i] = &sdkutilities.TracePayload{
				Key:   val.Key,
				Value: val.Value,
			}
			if val.OperationType != "" {
				v.Payload[i].OperationType = &val.OperationType
			}
		}
	}

	return v, nil
}
