// Code generated by goa v3.5.2, DO NOT EDIT.
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
	return `sdk-utilities (account-numbers|supply|query-tx|broadcast-tx|tx-metadata|block|liquidity-params|liquidity-pools|mint-inflation|mint-params|mint-annual-provision)
`
}

// UsageExamples produces an example of a valid invocation of the CLI tool.
func UsageExamples() string {
	return os.Args[0] + ` sdk-utilities account-numbers --message '{
      "addresHex": "Dolorum ut eos qui earum.",
      "bech32Prefix": "Blanditiis ut consequatur.",
      "chainName": "Ipsam nesciunt voluptate nulla numquam dolorum.",
      "port": 1457976898164366168
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
      "addresHex": "Dolorum ut eos qui earum.",
      "bech32Prefix": "Blanditiis ut consequatur.",
      "chainName": "Ipsam nesciunt voluptate nulla numquam dolorum.",
      "port": 1457976898164366168
   }'
`, os.Args[0])
}

func sdkUtilitiesSupplyUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] sdk-utilities supply -message JSON

Supply implements supply.
    -message JSON: 

Example:
    %[1]s sdk-utilities supply --message '{
      "chainName": "Repudiandae hic unde distinctio consectetur.",
      "port": 6548433300935100839
   }'
`, os.Args[0])
}

func sdkUtilitiesQueryTxUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] sdk-utilities query-tx -message JSON

QueryTx implements queryTx.
    -message JSON: 

Example:
    %[1]s sdk-utilities query-tx --message '{
      "chainName": "Aut id.",
      "hash": "Veritatis blanditiis quaerat eaque.",
      "port": 4799841495418416071
   }'
`, os.Args[0])
}

func sdkUtilitiesBroadcastTxUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] sdk-utilities broadcast-tx -message JSON

BroadcastTx implements broadcastTx.
    -message JSON: 

Example:
    %[1]s sdk-utilities broadcast-tx --message '{
      "chainName": "Rerum fugiat ut est molestias ut non.",
      "port": 8213706707786077595,
      "txBytes": "TGliZXJvIGVuaW0gcmVydW0gZG9sb3IgcXVpIHNlZC4="
   }'
`, os.Args[0])
}

func sdkUtilitiesTxMetadataUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] sdk-utilities tx-metadata -message JSON

TxMetadata implements txMetadata.
    -message JSON: 

Example:
    %[1]s sdk-utilities tx-metadata --message '{
      "txBytes": "Q29uc2VxdWF0dXIgZXN0IGVhIGVzdC4="
   }'
`, os.Args[0])
}

func sdkUtilitiesBlockUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] sdk-utilities block -message JSON

Block implements block.
    -message JSON: 

Example:
    %[1]s sdk-utilities block --message '{
      "chainName": "Laborum quam ducimus consequatur.",
      "height": 4738969913563760839,
      "port": 984187946088149780
   }'
`, os.Args[0])
}

func sdkUtilitiesLiquidityParamsUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] sdk-utilities liquidity-params -message JSON

LiquidityParams implements liquidityParams.
    -message JSON: 

Example:
    %[1]s sdk-utilities liquidity-params --message '{
      "chainName": "Dolor quisquam est et aspernatur.",
      "port": 894623405773491359
   }'
`, os.Args[0])
}

func sdkUtilitiesLiquidityPoolsUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] sdk-utilities liquidity-pools -message JSON

LiquidityPools implements liquidityPools.
    -message JSON: 

Example:
    %[1]s sdk-utilities liquidity-pools --message '{
      "chainName": "Alias accusantium sunt rem voluptas sed.",
      "port": 4442055250494698477
   }'
`, os.Args[0])
}

func sdkUtilitiesMintInflationUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] sdk-utilities mint-inflation -message JSON

MintInflation implements mintInflation.
    -message JSON: 

Example:
    %[1]s sdk-utilities mint-inflation --message '{
      "chainName": "Sequi quis culpa et blanditiis veritatis.",
      "port": 2428174418171390862
   }'
`, os.Args[0])
}

func sdkUtilitiesMintParamsUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] sdk-utilities mint-params -message JSON

MintParams implements mintParams.
    -message JSON: 

Example:
    %[1]s sdk-utilities mint-params --message '{
      "chainName": "Est neque porro consectetur.",
      "port": 5472785720082760056
   }'
`, os.Args[0])
}

func sdkUtilitiesMintAnnualProvisionUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] sdk-utilities mint-annual-provision -message JSON

MintAnnualProvision implements mintAnnualProvision.
    -message JSON: 

Example:
    %[1]s sdk-utilities mint-annual-provision --message '{
      "chainName": "Ullam quod.",
      "port": 4459115708693118891
   }'
`, os.Args[0])
}
