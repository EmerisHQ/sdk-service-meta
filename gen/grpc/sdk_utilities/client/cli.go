// Code generated by goa v3.5.3, DO NOT EDIT.
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

// BuildAccountNumbersPayload builds the payload for the sdk-utilities
// accountNumbers endpoint from CLI flags.
func BuildAccountNumbersPayload(sdkUtilitiesAccountNumbersMessage string) (*sdkutilities.AccountNumbersPayload, error) {
	var err error
	var message sdk_utilitiespb.AccountNumbersRequest
	{
		if sdkUtilitiesAccountNumbersMessage != "" {
			err = json.Unmarshal([]byte(sdkUtilitiesAccountNumbersMessage), &message)
			if err != nil {
				return nil, fmt.Errorf("invalid JSON for message, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"addresHex\": \"Veritatis quia.\",\n      \"bech32Prefix\": \"Quisquam sint quia et magnam.\",\n      \"chainName\": \"Quaerat eaque quas ea.\",\n      \"port\": 7031085827166329699\n   }'")
			}
		}
	}
	v := &sdkutilities.AccountNumbersPayload{
		ChainName: message.ChainName,
	}
	if message.Port != 0 {
		portptr := int(message.Port)
		v.Port = &portptr
	}
	if message.Bech32Prefix != "" {
		v.Bech32Prefix = &message.Bech32Prefix
	}
	if message.AddresHex != "" {
		v.AddresHex = &message.AddresHex
	}

	return v, nil
}

// BuildSupplyPayload builds the payload for the sdk-utilities supply endpoint
// from CLI flags.
func BuildSupplyPayload(sdkUtilitiesSupplyMessage string) (*sdkutilities.SupplyPayload, error) {
	var err error
	var message sdk_utilitiespb.SupplyRequest
	{
		if sdkUtilitiesSupplyMessage != "" {
			err = json.Unmarshal([]byte(sdkUtilitiesSupplyMessage), &message)
			if err != nil {
				return nil, fmt.Errorf("invalid JSON for message, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"chainName\": \"Temporibus libero enim rerum dolor qui sed.\",\n      \"port\": 3525784811766824029\n   }'")
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
				return nil, fmt.Errorf("invalid JSON for message, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"chainName\": \"Esse dicta velit velit dicta.\",\n      \"hash\": \"Est ea est.\",\n      \"port\": 8651610214817669303\n   }'")
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
				return nil, fmt.Errorf("invalid JSON for message, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"chainName\": \"Eaque ut.\",\n      \"port\": 1219077044981623912,\n      \"txBytes\": \"UXVpc3F1YW0gZXN0IGV0IGFzcGVybmF0dXIgc3VudCBvbW5pcyBldC4=\"\n   }'")
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
				return nil, fmt.Errorf("invalid JSON for message, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"txBytes\": \"U3VudCByZW0gdm9sdXB0YXMgc2VkIGlkLg==\"\n   }'")
			}
		}
	}
	v := &sdkutilities.TxMetadataPayload{
		TxBytes: message.TxBytes,
	}

	return v, nil
}

// BuildBlockPayload builds the payload for the sdk-utilities block endpoint
// from CLI flags.
func BuildBlockPayload(sdkUtilitiesBlockMessage string) (*sdkutilities.BlockPayload, error) {
	var err error
	var message sdk_utilitiespb.BlockRequest
	{
		if sdkUtilitiesBlockMessage != "" {
			err = json.Unmarshal([]byte(sdkUtilitiesBlockMessage), &message)
			if err != nil {
				return nil, fmt.Errorf("invalid JSON for message, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"chainName\": \"Quia velit.\",\n      \"height\": 2966708842863342049,\n      \"port\": 3229902183543090429\n   }'")
			}
		}
	}
	v := &sdkutilities.BlockPayload{
		ChainName: message.ChainName,
		Height:    message.Height,
	}
	if message.Port != 0 {
		portptr := int(message.Port)
		v.Port = &portptr
	}

	return v, nil
}

// BuildLiquidityParamsPayload builds the payload for the sdk-utilities
// liquidityParams endpoint from CLI flags.
func BuildLiquidityParamsPayload(sdkUtilitiesLiquidityParamsMessage string) (*sdkutilities.LiquidityParamsPayload, error) {
	var err error
	var message sdk_utilitiespb.LiquidityParamsRequest
	{
		if sdkUtilitiesLiquidityParamsMessage != "" {
			err = json.Unmarshal([]byte(sdkUtilitiesLiquidityParamsMessage), &message)
			if err != nil {
				return nil, fmt.Errorf("invalid JSON for message, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"chainName\": \"Cumque itaque tenetur explicabo.\",\n      \"port\": 2970337066418489394\n   }'")
			}
		}
	}
	v := &sdkutilities.LiquidityParamsPayload{
		ChainName: message.ChainName,
	}
	if message.Port != 0 {
		portptr := int(message.Port)
		v.Port = &portptr
	}

	return v, nil
}

// BuildLiquidityPoolsPayload builds the payload for the sdk-utilities
// liquidityPools endpoint from CLI flags.
func BuildLiquidityPoolsPayload(sdkUtilitiesLiquidityPoolsMessage string) (*sdkutilities.LiquidityPoolsPayload, error) {
	var err error
	var message sdk_utilitiespb.LiquidityPoolsRequest
	{
		if sdkUtilitiesLiquidityPoolsMessage != "" {
			err = json.Unmarshal([]byte(sdkUtilitiesLiquidityPoolsMessage), &message)
			if err != nil {
				return nil, fmt.Errorf("invalid JSON for message, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"chainName\": \"Ab ut sit debitis dicta ullam quod.\",\n      \"port\": 4459115708693118891\n   }'")
			}
		}
	}
	v := &sdkutilities.LiquidityPoolsPayload{
		ChainName: message.ChainName,
	}
	if message.Port != 0 {
		portptr := int(message.Port)
		v.Port = &portptr
	}

	return v, nil
}

// BuildMintInflationPayload builds the payload for the sdk-utilities
// mintInflation endpoint from CLI flags.
func BuildMintInflationPayload(sdkUtilitiesMintInflationMessage string) (*sdkutilities.MintInflationPayload, error) {
	var err error
	var message sdk_utilitiespb.MintInflationRequest
	{
		if sdkUtilitiesMintInflationMessage != "" {
			err = json.Unmarshal([]byte(sdkUtilitiesMintInflationMessage), &message)
			if err != nil {
				return nil, fmt.Errorf("invalid JSON for message, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"chainName\": \"Natus quasi.\",\n      \"port\": 4941727330865473130\n   }'")
			}
		}
	}
	v := &sdkutilities.MintInflationPayload{
		ChainName: message.ChainName,
	}
	if message.Port != 0 {
		portptr := int(message.Port)
		v.Port = &portptr
	}

	return v, nil
}

// BuildMintParamsPayload builds the payload for the sdk-utilities mintParams
// endpoint from CLI flags.
func BuildMintParamsPayload(sdkUtilitiesMintParamsMessage string) (*sdkutilities.MintParamsPayload, error) {
	var err error
	var message sdk_utilitiespb.MintParamsRequest
	{
		if sdkUtilitiesMintParamsMessage != "" {
			err = json.Unmarshal([]byte(sdkUtilitiesMintParamsMessage), &message)
			if err != nil {
				return nil, fmt.Errorf("invalid JSON for message, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"chainName\": \"Velit tenetur aut ducimus voluptatibus perferendis.\",\n      \"port\": 4590398434326812352\n   }'")
			}
		}
	}
	v := &sdkutilities.MintParamsPayload{
		ChainName: message.ChainName,
	}
	if message.Port != 0 {
		portptr := int(message.Port)
		v.Port = &portptr
	}

	return v, nil
}

// BuildMintAnnualProvisionPayload builds the payload for the sdk-utilities
// mintAnnualProvision endpoint from CLI flags.
func BuildMintAnnualProvisionPayload(sdkUtilitiesMintAnnualProvisionMessage string) (*sdkutilities.MintAnnualProvisionPayload, error) {
	var err error
	var message sdk_utilitiespb.MintAnnualProvisionRequest
	{
		if sdkUtilitiesMintAnnualProvisionMessage != "" {
			err = json.Unmarshal([]byte(sdkUtilitiesMintAnnualProvisionMessage), &message)
			if err != nil {
				return nil, fmt.Errorf("invalid JSON for message, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"chainName\": \"Occaecati dolorem tenetur cum dolores veniam.\",\n      \"port\": 435799008750423843\n   }'")
			}
		}
	}
	v := &sdkutilities.MintAnnualProvisionPayload{
		ChainName: message.ChainName,
	}
	if message.Port != 0 {
		portptr := int(message.Port)
		v.Port = &portptr
	}

	return v, nil
}

// BuildDelegatorRewardsPayload builds the payload for the sdk-utilities
// delegatorRewards endpoint from CLI flags.
func BuildDelegatorRewardsPayload(sdkUtilitiesDelegatorRewardsMessage string) (*sdkutilities.DelegatorRewardsPayload, error) {
	var err error
	var message sdk_utilitiespb.DelegatorRewardsRequest
	{
		if sdkUtilitiesDelegatorRewardsMessage != "" {
			err = json.Unmarshal([]byte(sdkUtilitiesDelegatorRewardsMessage), &message)
			if err != nil {
				return nil, fmt.Errorf("invalid JSON for message, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"addresHex\": \"Autem fugiat optio.\",\n      \"bech32Prefix\": \"Quo dolorum.\",\n      \"chainName\": \"Est doloribus similique similique ea eaque dolorem.\",\n      \"port\": 3447600267286363243\n   }'")
			}
		}
	}
	v := &sdkutilities.DelegatorRewardsPayload{
		ChainName: message.ChainName,
	}
	if message.Port != 0 {
		portptr := int(message.Port)
		v.Port = &portptr
	}
	if message.Bech32Prefix != "" {
		v.Bech32Prefix = &message.Bech32Prefix
	}
	if message.AddresHex != "" {
		v.AddresHex = &message.AddresHex
	}

	return v, nil
}

// BuildTxFeeEstimatePayload builds the payload for the sdk-utilities
// txFeeEstimate endpoint from CLI flags.
func BuildTxFeeEstimatePayload(sdkUtilitiesTxFeeEstimateMessage string) (*sdkutilities.TxFeeEstimatePayload, error) {
	var err error
	var message sdk_utilitiespb.TxFeeEstimateRequest
	{
		if sdkUtilitiesTxFeeEstimateMessage != "" {
			err = json.Unmarshal([]byte(sdkUtilitiesTxFeeEstimateMessage), &message)
			if err != nil {
				return nil, fmt.Errorf("invalid JSON for message, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"txBytes\": \"Tm9uIGFwZXJpYW0gc2VxdWkgcHJhZXNlbnRpdW0gdmVuaWFtIHV0Lg==\"\n   }'")
			}
		}
	}
	v := &sdkutilities.TxFeeEstimatePayload{
		TxBytes: message.TxBytes,
	}

	return v, nil
}
