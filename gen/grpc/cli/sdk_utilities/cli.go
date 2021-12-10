// Code generated by goa v3.5.3, DO NOT EDIT.
//
// sdk-utilities gRPC client CLI support package
//
// Command:
// $ goa gen github.com/allinbits/sdk-service-meta

package cli

import (
	"flag"
	"fmt"
	"os"

	sdkutilitiesc "github.com/allinbits/sdk-service-meta/gen/grpc/sdk_utilities/client"
	goa "goa.design/goa/v3/pkg"
	grpc "google.golang.org/grpc"
)

// UsageCommands returns the set of commands and sub-commands using the format
//
//    command (subcommand1|subcommand2|...)
//
func UsageCommands() string {
	return `sdk-utilities (account-numbers|supply|query-tx|broadcast-tx|tx-metadata|block|liquidity-params|liquidity-pools|mint-inflation|mint-params|mint-annual-provision|delegator-rewards|tx-fee-estimate)
`
}

// UsageExamples produces an example of a valid invocation of the CLI tool.
func UsageExamples() string {
	return os.Args[0] + ` sdk-utilities account-numbers --message '{
      "addresHex": "Veritatis quia.",
      "bech32Prefix": "Quisquam sint quia et magnam.",
      "chainName": "Quaerat eaque quas ea.",
      "port": 7031085827166329699
   }'` + "\n" +
		""
}

// ParseEndpoint returns the endpoint and payload as specified on the command
// line.
func ParseEndpoint(cc *grpc.ClientConn, opts ...grpc.CallOption) (goa.Endpoint, interface{}, error) {
	var (
		sdkUtilitiesFlags = flag.NewFlagSet("sdk-utilities", flag.ContinueOnError)

		sdkUtilitiesAccountNumbersFlags       = flag.NewFlagSet("account-numbers", flag.ExitOnError)
		sdkUtilitiesAccountNumbersMessageFlag = sdkUtilitiesAccountNumbersFlags.String("message", "", "")

		sdkUtilitiesSupplyFlags       = flag.NewFlagSet("supply", flag.ExitOnError)
		sdkUtilitiesSupplyMessageFlag = sdkUtilitiesSupplyFlags.String("message", "", "")

		sdkUtilitiesQueryTxFlags       = flag.NewFlagSet("query-tx", flag.ExitOnError)
		sdkUtilitiesQueryTxMessageFlag = sdkUtilitiesQueryTxFlags.String("message", "", "")

		sdkUtilitiesBroadcastTxFlags       = flag.NewFlagSet("broadcast-tx", flag.ExitOnError)
		sdkUtilitiesBroadcastTxMessageFlag = sdkUtilitiesBroadcastTxFlags.String("message", "", "")

		sdkUtilitiesTxMetadataFlags       = flag.NewFlagSet("tx-metadata", flag.ExitOnError)
		sdkUtilitiesTxMetadataMessageFlag = sdkUtilitiesTxMetadataFlags.String("message", "", "")

		sdkUtilitiesBlockFlags       = flag.NewFlagSet("block", flag.ExitOnError)
		sdkUtilitiesBlockMessageFlag = sdkUtilitiesBlockFlags.String("message", "", "")

		sdkUtilitiesLiquidityParamsFlags       = flag.NewFlagSet("liquidity-params", flag.ExitOnError)
		sdkUtilitiesLiquidityParamsMessageFlag = sdkUtilitiesLiquidityParamsFlags.String("message", "", "")

		sdkUtilitiesLiquidityPoolsFlags       = flag.NewFlagSet("liquidity-pools", flag.ExitOnError)
		sdkUtilitiesLiquidityPoolsMessageFlag = sdkUtilitiesLiquidityPoolsFlags.String("message", "", "")

		sdkUtilitiesMintInflationFlags       = flag.NewFlagSet("mint-inflation", flag.ExitOnError)
		sdkUtilitiesMintInflationMessageFlag = sdkUtilitiesMintInflationFlags.String("message", "", "")

		sdkUtilitiesMintParamsFlags       = flag.NewFlagSet("mint-params", flag.ExitOnError)
		sdkUtilitiesMintParamsMessageFlag = sdkUtilitiesMintParamsFlags.String("message", "", "")

		sdkUtilitiesMintAnnualProvisionFlags       = flag.NewFlagSet("mint-annual-provision", flag.ExitOnError)
		sdkUtilitiesMintAnnualProvisionMessageFlag = sdkUtilitiesMintAnnualProvisionFlags.String("message", "", "")

		sdkUtilitiesDelegatorRewardsFlags       = flag.NewFlagSet("delegator-rewards", flag.ExitOnError)
		sdkUtilitiesDelegatorRewardsMessageFlag = sdkUtilitiesDelegatorRewardsFlags.String("message", "", "")

		sdkUtilitiesTxFeeEstimateFlags       = flag.NewFlagSet("tx-fee-estimate", flag.ExitOnError)
		sdkUtilitiesTxFeeEstimateMessageFlag = sdkUtilitiesTxFeeEstimateFlags.String("message", "", "")
	)
	sdkUtilitiesFlags.Usage = sdkUtilitiesUsage
	sdkUtilitiesAccountNumbersFlags.Usage = sdkUtilitiesAccountNumbersUsage
	sdkUtilitiesSupplyFlags.Usage = sdkUtilitiesSupplyUsage
	sdkUtilitiesQueryTxFlags.Usage = sdkUtilitiesQueryTxUsage
	sdkUtilitiesBroadcastTxFlags.Usage = sdkUtilitiesBroadcastTxUsage
	sdkUtilitiesTxMetadataFlags.Usage = sdkUtilitiesTxMetadataUsage
	sdkUtilitiesBlockFlags.Usage = sdkUtilitiesBlockUsage
	sdkUtilitiesLiquidityParamsFlags.Usage = sdkUtilitiesLiquidityParamsUsage
	sdkUtilitiesLiquidityPoolsFlags.Usage = sdkUtilitiesLiquidityPoolsUsage
	sdkUtilitiesMintInflationFlags.Usage = sdkUtilitiesMintInflationUsage
	sdkUtilitiesMintParamsFlags.Usage = sdkUtilitiesMintParamsUsage
	sdkUtilitiesMintAnnualProvisionFlags.Usage = sdkUtilitiesMintAnnualProvisionUsage
	sdkUtilitiesDelegatorRewardsFlags.Usage = sdkUtilitiesDelegatorRewardsUsage
	sdkUtilitiesTxFeeEstimateFlags.Usage = sdkUtilitiesTxFeeEstimateUsage

	if err := flag.CommandLine.Parse(os.Args[1:]); err != nil {
		return nil, nil, err
	}

	if flag.NArg() < 2 { // two non flag args are required: SERVICE and ENDPOINT (aka COMMAND)
		return nil, nil, fmt.Errorf("not enough arguments")
	}

	var (
		svcn string
		svcf *flag.FlagSet
	)
	{
		svcn = flag.Arg(0)
		switch svcn {
		case "sdk-utilities":
			svcf = sdkUtilitiesFlags
		default:
			return nil, nil, fmt.Errorf("unknown service %q", svcn)
		}
	}
	if err := svcf.Parse(flag.Args()[1:]); err != nil {
		return nil, nil, err
	}

	var (
		epn string
		epf *flag.FlagSet
	)
	{
		epn = svcf.Arg(0)
		switch svcn {
		case "sdk-utilities":
			switch epn {
			case "account-numbers":
				epf = sdkUtilitiesAccountNumbersFlags

			case "supply":
				epf = sdkUtilitiesSupplyFlags

			case "query-tx":
				epf = sdkUtilitiesQueryTxFlags

			case "broadcast-tx":
				epf = sdkUtilitiesBroadcastTxFlags

			case "tx-metadata":
				epf = sdkUtilitiesTxMetadataFlags

			case "block":
				epf = sdkUtilitiesBlockFlags

			case "liquidity-params":
				epf = sdkUtilitiesLiquidityParamsFlags

			case "liquidity-pools":
				epf = sdkUtilitiesLiquidityPoolsFlags

			case "mint-inflation":
				epf = sdkUtilitiesMintInflationFlags

			case "mint-params":
				epf = sdkUtilitiesMintParamsFlags

			case "mint-annual-provision":
				epf = sdkUtilitiesMintAnnualProvisionFlags

			case "delegator-rewards":
				epf = sdkUtilitiesDelegatorRewardsFlags

			case "tx-fee-estimate":
				epf = sdkUtilitiesTxFeeEstimateFlags

			}

		}
	}
	if epf == nil {
		return nil, nil, fmt.Errorf("unknown %q endpoint %q", svcn, epn)
	}

	// Parse endpoint flags if any
	if svcf.NArg() > 1 {
		if err := epf.Parse(svcf.Args()[1:]); err != nil {
			return nil, nil, err
		}
	}

	var (
		data     interface{}
		endpoint goa.Endpoint
		err      error
	)
	{
		switch svcn {
		case "sdk-utilities":
			c := sdkutilitiesc.NewClient(cc, opts...)
			switch epn {
			case "account-numbers":
				endpoint = c.AccountNumbers()
				data, err = sdkutilitiesc.BuildAccountNumbersPayload(*sdkUtilitiesAccountNumbersMessageFlag)
			case "supply":
				endpoint = c.Supply()
				data, err = sdkutilitiesc.BuildSupplyPayload(*sdkUtilitiesSupplyMessageFlag)
			case "query-tx":
				endpoint = c.QueryTx()
				data, err = sdkutilitiesc.BuildQueryTxPayload(*sdkUtilitiesQueryTxMessageFlag)
			case "broadcast-tx":
				endpoint = c.BroadcastTx()
				data, err = sdkutilitiesc.BuildBroadcastTxPayload(*sdkUtilitiesBroadcastTxMessageFlag)
			case "tx-metadata":
				endpoint = c.TxMetadata()
				data, err = sdkutilitiesc.BuildTxMetadataPayload(*sdkUtilitiesTxMetadataMessageFlag)
			case "block":
				endpoint = c.Block()
				data, err = sdkutilitiesc.BuildBlockPayload(*sdkUtilitiesBlockMessageFlag)
			case "liquidity-params":
				endpoint = c.LiquidityParams()
				data, err = sdkutilitiesc.BuildLiquidityParamsPayload(*sdkUtilitiesLiquidityParamsMessageFlag)
			case "liquidity-pools":
				endpoint = c.LiquidityPools()
				data, err = sdkutilitiesc.BuildLiquidityPoolsPayload(*sdkUtilitiesLiquidityPoolsMessageFlag)
			case "mint-inflation":
				endpoint = c.MintInflation()
				data, err = sdkutilitiesc.BuildMintInflationPayload(*sdkUtilitiesMintInflationMessageFlag)
			case "mint-params":
				endpoint = c.MintParams()
				data, err = sdkutilitiesc.BuildMintParamsPayload(*sdkUtilitiesMintParamsMessageFlag)
			case "mint-annual-provision":
				endpoint = c.MintAnnualProvision()
				data, err = sdkutilitiesc.BuildMintAnnualProvisionPayload(*sdkUtilitiesMintAnnualProvisionMessageFlag)
			case "delegator-rewards":
				endpoint = c.DelegatorRewards()
				data, err = sdkutilitiesc.BuildDelegatorRewardsPayload(*sdkUtilitiesDelegatorRewardsMessageFlag)
			case "tx-fee-estimate":
				endpoint = c.TxFeeEstimate()
				data, err = sdkutilitiesc.BuildTxFeeEstimatePayload(*sdkUtilitiesTxFeeEstimateMessageFlag)
			}
		}
	}
	if err != nil {
		return nil, nil, err
	}

	return endpoint, data, nil
}

// sdk-utilitiesUsage displays the usage of the sdk-utilities command and its
// subcommands.
func sdkUtilitiesUsage() {
	fmt.Fprintf(os.Stderr, `sdk-utilities performs Cosmos SDK-related operations
Usage:
    %[1]s [globalflags] sdk-utilities COMMAND [flags]

COMMAND:
    account-numbers: AccountNumbers implements accountNumbers.
    supply: Supply implements supply.
    query-tx: QueryTx implements queryTx.
    broadcast-tx: BroadcastTx implements broadcastTx.
    tx-metadata: TxMetadata implements txMetadata.
    block: Block implements block.
    liquidity-params: LiquidityParams implements liquidityParams.
    liquidity-pools: LiquidityPools implements liquidityPools.
    mint-inflation: MintInflation implements mintInflation.
    mint-params: MintParams implements mintParams.
    mint-annual-provision: MintAnnualProvision implements mintAnnualProvision.
    delegator-rewards: DelegatorRewards implements delegatorRewards.
    tx-fee-estimate: TxFeeEstimate implements txFeeEstimate.

Additional help:
    %[1]s sdk-utilities COMMAND --help
`, os.Args[0])
}
func sdkUtilitiesAccountNumbersUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] sdk-utilities account-numbers -message JSON

AccountNumbers implements accountNumbers.
    -message JSON: 

Example:
    %[1]s sdk-utilities account-numbers --message '{
      "addresHex": "Veritatis quia.",
      "bech32Prefix": "Quisquam sint quia et magnam.",
      "chainName": "Quaerat eaque quas ea.",
      "port": 7031085827166329699
   }'
`, os.Args[0])
}

func sdkUtilitiesSupplyUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] sdk-utilities supply -message JSON

Supply implements supply.
    -message JSON: 

Example:
    %[1]s sdk-utilities supply --message '{
      "chainName": "Temporibus libero enim rerum dolor qui sed.",
      "port": 3525784811766824029
   }'
`, os.Args[0])
}

func sdkUtilitiesQueryTxUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] sdk-utilities query-tx -message JSON

QueryTx implements queryTx.
    -message JSON: 

Example:
    %[1]s sdk-utilities query-tx --message '{
      "chainName": "Esse dicta velit velit dicta.",
      "hash": "Est ea est.",
      "port": 8651610214817669303
   }'
`, os.Args[0])
}

func sdkUtilitiesBroadcastTxUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] sdk-utilities broadcast-tx -message JSON

BroadcastTx implements broadcastTx.
    -message JSON: 

Example:
    %[1]s sdk-utilities broadcast-tx --message '{
      "chainName": "Eaque ut.",
      "port": 1219077044981623912,
      "txBytes": "UXVpc3F1YW0gZXN0IGV0IGFzcGVybmF0dXIgc3VudCBvbW5pcyBldC4="
   }'
`, os.Args[0])
}

func sdkUtilitiesTxMetadataUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] sdk-utilities tx-metadata -message JSON

TxMetadata implements txMetadata.
    -message JSON: 

Example:
    %[1]s sdk-utilities tx-metadata --message '{
      "txBytes": "U3VudCByZW0gdm9sdXB0YXMgc2VkIGlkLg=="
   }'
`, os.Args[0])
}

func sdkUtilitiesBlockUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] sdk-utilities block -message JSON

Block implements block.
    -message JSON: 

Example:
    %[1]s sdk-utilities block --message '{
      "chainName": "Quia velit.",
      "height": 2966708842863342049,
      "port": 3229902183543090429
   }'
`, os.Args[0])
}

func sdkUtilitiesLiquidityParamsUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] sdk-utilities liquidity-params -message JSON

LiquidityParams implements liquidityParams.
    -message JSON: 

Example:
    %[1]s sdk-utilities liquidity-params --message '{
      "chainName": "Cumque itaque tenetur explicabo.",
      "port": 2970337066418489394
   }'
`, os.Args[0])
}

func sdkUtilitiesLiquidityPoolsUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] sdk-utilities liquidity-pools -message JSON

LiquidityPools implements liquidityPools.
    -message JSON: 

Example:
    %[1]s sdk-utilities liquidity-pools --message '{
      "chainName": "Ab ut sit debitis dicta ullam quod.",
      "port": 4459115708693118891
   }'
`, os.Args[0])
}

func sdkUtilitiesMintInflationUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] sdk-utilities mint-inflation -message JSON

MintInflation implements mintInflation.
    -message JSON: 

Example:
    %[1]s sdk-utilities mint-inflation --message '{
      "chainName": "Natus quasi.",
      "port": 4941727330865473130
   }'
`, os.Args[0])
}

func sdkUtilitiesMintParamsUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] sdk-utilities mint-params -message JSON

MintParams implements mintParams.
    -message JSON: 

Example:
    %[1]s sdk-utilities mint-params --message '{
      "chainName": "Velit tenetur aut ducimus voluptatibus perferendis.",
      "port": 4590398434326812352
   }'
`, os.Args[0])
}

func sdkUtilitiesMintAnnualProvisionUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] sdk-utilities mint-annual-provision -message JSON

MintAnnualProvision implements mintAnnualProvision.
    -message JSON: 

Example:
    %[1]s sdk-utilities mint-annual-provision --message '{
      "chainName": "Occaecati dolorem tenetur cum dolores veniam.",
      "port": 435799008750423843
   }'
`, os.Args[0])
}

func sdkUtilitiesDelegatorRewardsUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] sdk-utilities delegator-rewards -message JSON

DelegatorRewards implements delegatorRewards.
    -message JSON: 

Example:
    %[1]s sdk-utilities delegator-rewards --message '{
      "addresHex": "Autem fugiat optio.",
      "bech32Prefix": "Quo dolorum.",
      "chainName": "Est doloribus similique similique ea eaque dolorem.",
      "port": 3447600267286363243
   }'
`, os.Args[0])
}

func sdkUtilitiesTxFeeEstimateUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] sdk-utilities tx-fee-estimate -message JSON

TxFeeEstimate implements txFeeEstimate.
    -message JSON: 

Example:
    %[1]s sdk-utilities tx-fee-estimate --message '{
      "txBytes": "Tm9uIGFwZXJpYW0gc2VxdWkgcHJhZXNlbnRpdW0gdmVuaWFtIHV0Lg=="
   }'
`, os.Args[0])
}
