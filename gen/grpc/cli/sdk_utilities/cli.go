// Code generated by goa v3.7.5, DO NOT EDIT.
//
// sdk-utilities gRPC client CLI support package
//
// Command:
// $ goa gen github.com/emerishq/sdk-service-meta

package cli

import (
	"flag"
	"fmt"
	"os"

	sdkutilitiesc "github.com/emerishq/sdk-service-meta/gen/grpc/sdk_utilities/client"
	goa "goa.design/goa/v3/pkg"
	grpc "google.golang.org/grpc"
)

// UsageCommands returns the set of commands and sub-commands using the format
//
//    command (subcommand1|subcommand2|...)
//
func UsageCommands() string {
	return `sdk-utilities (account-numbers|supply|supply-denom|query-tx|broadcast-tx|tx-metadata|block|liquidity-params|liquidity-pools|mint-inflation|mint-params|mint-annual-provision|mint-epoch-provisions|delegator-rewards|estimate-fees|staking-params|staking-pool|emoney-inflation|budget-params|distribution-params|osmo-pools|crescent-pools)
`
}

// UsageExamples produces an example of a valid invocation of the CLI tool.
func UsageExamples() string {
	return os.Args[0] + ` sdk-utilities account-numbers --message '{
      "addresHex": "Autem fugiat optio.",
      "bech32Prefix": "Et quo dolorum.",
      "chainName": "Ea eaque.",
      "port": 3852088934630280001
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

		sdkUtilitiesSupplyDenomFlags       = flag.NewFlagSet("supply-denom", flag.ExitOnError)
		sdkUtilitiesSupplyDenomMessageFlag = sdkUtilitiesSupplyDenomFlags.String("message", "", "")

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

		sdkUtilitiesMintEpochProvisionsFlags       = flag.NewFlagSet("mint-epoch-provisions", flag.ExitOnError)
		sdkUtilitiesMintEpochProvisionsMessageFlag = sdkUtilitiesMintEpochProvisionsFlags.String("message", "", "")

		sdkUtilitiesDelegatorRewardsFlags       = flag.NewFlagSet("delegator-rewards", flag.ExitOnError)
		sdkUtilitiesDelegatorRewardsMessageFlag = sdkUtilitiesDelegatorRewardsFlags.String("message", "", "")

		sdkUtilitiesEstimateFeesFlags       = flag.NewFlagSet("estimate-fees", flag.ExitOnError)
		sdkUtilitiesEstimateFeesMessageFlag = sdkUtilitiesEstimateFeesFlags.String("message", "", "")

		sdkUtilitiesStakingParamsFlags       = flag.NewFlagSet("staking-params", flag.ExitOnError)
		sdkUtilitiesStakingParamsMessageFlag = sdkUtilitiesStakingParamsFlags.String("message", "", "")

		sdkUtilitiesStakingPoolFlags       = flag.NewFlagSet("staking-pool", flag.ExitOnError)
		sdkUtilitiesStakingPoolMessageFlag = sdkUtilitiesStakingPoolFlags.String("message", "", "")

		sdkUtilitiesEmoneyInflationFlags       = flag.NewFlagSet("emoney-inflation", flag.ExitOnError)
		sdkUtilitiesEmoneyInflationMessageFlag = sdkUtilitiesEmoneyInflationFlags.String("message", "", "")

		sdkUtilitiesBudgetParamsFlags       = flag.NewFlagSet("budget-params", flag.ExitOnError)
		sdkUtilitiesBudgetParamsMessageFlag = sdkUtilitiesBudgetParamsFlags.String("message", "", "")

		sdkUtilitiesDistributionParamsFlags       = flag.NewFlagSet("distribution-params", flag.ExitOnError)
		sdkUtilitiesDistributionParamsMessageFlag = sdkUtilitiesDistributionParamsFlags.String("message", "", "")

		sdkUtilitiesOsmoPoolsFlags       = flag.NewFlagSet("osmo-pools", flag.ExitOnError)
		sdkUtilitiesOsmoPoolsMessageFlag = sdkUtilitiesOsmoPoolsFlags.String("message", "", "")

		sdkUtilitiesCrescentPoolsFlags       = flag.NewFlagSet("crescent-pools", flag.ExitOnError)
		sdkUtilitiesCrescentPoolsMessageFlag = sdkUtilitiesCrescentPoolsFlags.String("message", "", "")
	)
	sdkUtilitiesFlags.Usage = sdkUtilitiesUsage
	sdkUtilitiesAccountNumbersFlags.Usage = sdkUtilitiesAccountNumbersUsage
	sdkUtilitiesSupplyFlags.Usage = sdkUtilitiesSupplyUsage
	sdkUtilitiesSupplyDenomFlags.Usage = sdkUtilitiesSupplyDenomUsage
	sdkUtilitiesQueryTxFlags.Usage = sdkUtilitiesQueryTxUsage
	sdkUtilitiesBroadcastTxFlags.Usage = sdkUtilitiesBroadcastTxUsage
	sdkUtilitiesTxMetadataFlags.Usage = sdkUtilitiesTxMetadataUsage
	sdkUtilitiesBlockFlags.Usage = sdkUtilitiesBlockUsage
	sdkUtilitiesLiquidityParamsFlags.Usage = sdkUtilitiesLiquidityParamsUsage
	sdkUtilitiesLiquidityPoolsFlags.Usage = sdkUtilitiesLiquidityPoolsUsage
	sdkUtilitiesMintInflationFlags.Usage = sdkUtilitiesMintInflationUsage
	sdkUtilitiesMintParamsFlags.Usage = sdkUtilitiesMintParamsUsage
	sdkUtilitiesMintAnnualProvisionFlags.Usage = sdkUtilitiesMintAnnualProvisionUsage
	sdkUtilitiesMintEpochProvisionsFlags.Usage = sdkUtilitiesMintEpochProvisionsUsage
	sdkUtilitiesDelegatorRewardsFlags.Usage = sdkUtilitiesDelegatorRewardsUsage
	sdkUtilitiesEstimateFeesFlags.Usage = sdkUtilitiesEstimateFeesUsage
	sdkUtilitiesStakingParamsFlags.Usage = sdkUtilitiesStakingParamsUsage
	sdkUtilitiesStakingPoolFlags.Usage = sdkUtilitiesStakingPoolUsage
	sdkUtilitiesEmoneyInflationFlags.Usage = sdkUtilitiesEmoneyInflationUsage
	sdkUtilitiesBudgetParamsFlags.Usage = sdkUtilitiesBudgetParamsUsage
	sdkUtilitiesDistributionParamsFlags.Usage = sdkUtilitiesDistributionParamsUsage
	sdkUtilitiesOsmoPoolsFlags.Usage = sdkUtilitiesOsmoPoolsUsage
	sdkUtilitiesCrescentPoolsFlags.Usage = sdkUtilitiesCrescentPoolsUsage

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

			case "supply-denom":
				epf = sdkUtilitiesSupplyDenomFlags

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

			case "mint-epoch-provisions":
				epf = sdkUtilitiesMintEpochProvisionsFlags

			case "delegator-rewards":
				epf = sdkUtilitiesDelegatorRewardsFlags

			case "estimate-fees":
				epf = sdkUtilitiesEstimateFeesFlags

			case "staking-params":
				epf = sdkUtilitiesStakingParamsFlags

			case "staking-pool":
				epf = sdkUtilitiesStakingPoolFlags

			case "emoney-inflation":
				epf = sdkUtilitiesEmoneyInflationFlags

			case "budget-params":
				epf = sdkUtilitiesBudgetParamsFlags

			case "distribution-params":
				epf = sdkUtilitiesDistributionParamsFlags

			case "osmo-pools":
				epf = sdkUtilitiesOsmoPoolsFlags

			case "crescent-pools":
				epf = sdkUtilitiesCrescentPoolsFlags

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
			case "supply-denom":
				endpoint = c.SupplyDenom()
				data, err = sdkutilitiesc.BuildSupplyDenomPayload(*sdkUtilitiesSupplyDenomMessageFlag)
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
			case "mint-epoch-provisions":
				endpoint = c.MintEpochProvisions()
				data, err = sdkutilitiesc.BuildMintEpochProvisionsPayload(*sdkUtilitiesMintEpochProvisionsMessageFlag)
			case "delegator-rewards":
				endpoint = c.DelegatorRewards()
				data, err = sdkutilitiesc.BuildDelegatorRewardsPayload(*sdkUtilitiesDelegatorRewardsMessageFlag)
			case "estimate-fees":
				endpoint = c.EstimateFees()
				data, err = sdkutilitiesc.BuildEstimateFeesPayload(*sdkUtilitiesEstimateFeesMessageFlag)
			case "staking-params":
				endpoint = c.StakingParams()
				data, err = sdkutilitiesc.BuildStakingParamsPayload(*sdkUtilitiesStakingParamsMessageFlag)
			case "staking-pool":
				endpoint = c.StakingPool()
				data, err = sdkutilitiesc.BuildStakingPoolPayload(*sdkUtilitiesStakingPoolMessageFlag)
			case "emoney-inflation":
				endpoint = c.EmoneyInflation()
				data, err = sdkutilitiesc.BuildEmoneyInflationPayload(*sdkUtilitiesEmoneyInflationMessageFlag)
			case "budget-params":
				endpoint = c.BudgetParams()
				data, err = sdkutilitiesc.BuildBudgetParamsPayload(*sdkUtilitiesBudgetParamsMessageFlag)
			case "distribution-params":
				endpoint = c.DistributionParams()
				data, err = sdkutilitiesc.BuildDistributionParamsPayload(*sdkUtilitiesDistributionParamsMessageFlag)
			case "osmo-pools":
				endpoint = c.OsmoPools()
				data, err = sdkutilitiesc.BuildOsmoPoolsPayload(*sdkUtilitiesOsmoPoolsMessageFlag)
			case "crescent-pools":
				endpoint = c.CrescentPools()
				data, err = sdkutilitiesc.BuildCrescentPoolsPayload(*sdkUtilitiesCrescentPoolsMessageFlag)
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
    supply-denom: SupplyDenom implements supplyDenom.
    query-tx: QueryTx implements queryTx.
    broadcast-tx: BroadcastTx implements broadcastTx.
    tx-metadata: TxMetadata implements txMetadata.
    block: Block implements block.
    liquidity-params: LiquidityParams implements liquidityParams.
    liquidity-pools: LiquidityPools implements liquidityPools.
    mint-inflation: MintInflation implements mintInflation.
    mint-params: MintParams implements mintParams.
    mint-annual-provision: MintAnnualProvision implements mintAnnualProvision.
    mint-epoch-provisions: MintEpochProvisions implements mintEpochProvisions.
    delegator-rewards: DelegatorRewards implements delegatorRewards.
    estimate-fees: EstimateFees implements estimateFees.
    staking-params: StakingParams implements stakingParams.
    staking-pool: StakingPool implements stakingPool.
    emoney-inflation: EmoneyInflation implements emoneyInflation.
    budget-params: BudgetParams implements budgetParams.
    distribution-params: DistributionParams implements distributionParams.
    osmo-pools: OsmoPools implements osmoPools.
    crescent-pools: CrescentPools implements crescentPools.

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
      "addresHex": "Autem fugiat optio.",
      "bech32Prefix": "Et quo dolorum.",
      "chainName": "Ea eaque.",
      "port": 3852088934630280001
   }'
`, os.Args[0])
}

func sdkUtilitiesSupplyUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] sdk-utilities supply -message JSON

Supply implements supply.
    -message JSON: 

Example:
    %[1]s sdk-utilities supply --message '{
      "chainName": "Amet nostrum sed est nihil error odit.",
      "paginationKey": "Eos maxime esse nulla quis blanditiis aspernatur.",
      "port": 3640771888125048896
   }'
`, os.Args[0])
}

func sdkUtilitiesSupplyDenomUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] sdk-utilities supply-denom -message JSON

SupplyDenom implements supplyDenom.
    -message JSON: 

Example:
    %[1]s sdk-utilities supply-denom --message '{
      "chainName": "Hic eum nostrum quia ut reiciendis ratione.",
      "denom": "Tempore alias quo animi eos.",
      "port": 382605817140914569
   }'
`, os.Args[0])
}

func sdkUtilitiesQueryTxUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] sdk-utilities query-tx -message JSON

QueryTx implements queryTx.
    -message JSON: 

Example:
    %[1]s sdk-utilities query-tx --message '{
      "chainName": "Velit occaecati ut.",
      "hash": "Qui ut sequi animi vero eaque.",
      "port": 1410794041151114724
   }'
`, os.Args[0])
}

func sdkUtilitiesBroadcastTxUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] sdk-utilities broadcast-tx -message JSON

BroadcastTx implements broadcastTx.
    -message JSON: 

Example:
    %[1]s sdk-utilities broadcast-tx --message '{
      "chainName": "Ut voluptas aut excepturi.",
      "port": 5921602754859307309,
      "txBytes": "RnVnYSBibGFuZGl0aWlzIHBsYWNlYXQgaXVyZSBleGVyY2l0YXRpb25lbSBpdXN0by4="
   }'
`, os.Args[0])
}

func sdkUtilitiesTxMetadataUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] sdk-utilities tx-metadata -message JSON

TxMetadata implements txMetadata.
    -message JSON: 

Example:
    %[1]s sdk-utilities tx-metadata --message '{
      "txBytes": "UXVpZGVtIG1vbGVzdGlhZSBjdW1xdWUgc3VudCBkb2xvcmVtLg=="
   }'
`, os.Args[0])
}

func sdkUtilitiesBlockUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] sdk-utilities block -message JSON

Block implements block.
    -message JSON: 

Example:
    %[1]s sdk-utilities block --message '{
      "chainName": "Non sit sunt excepturi ullam delectus corrupti.",
      "height": 8894233288527025489,
      "port": 4031148209971103528
   }'
`, os.Args[0])
}

func sdkUtilitiesLiquidityParamsUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] sdk-utilities liquidity-params -message JSON

LiquidityParams implements liquidityParams.
    -message JSON: 

Example:
    %[1]s sdk-utilities liquidity-params --message '{
      "chainName": "Nobis magnam tenetur.",
      "port": 8010145524969550654
   }'
`, os.Args[0])
}

func sdkUtilitiesLiquidityPoolsUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] sdk-utilities liquidity-pools -message JSON

LiquidityPools implements liquidityPools.
    -message JSON: 

Example:
    %[1]s sdk-utilities liquidity-pools --message '{
      "chainName": "Dolorem voluptatem dignissimos nam suscipit autem.",
      "port": 2954009122666477961
   }'
`, os.Args[0])
}

func sdkUtilitiesMintInflationUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] sdk-utilities mint-inflation -message JSON

MintInflation implements mintInflation.
    -message JSON: 

Example:
    %[1]s sdk-utilities mint-inflation --message '{
      "chainName": "Consequatur laborum aperiam et.",
      "port": 4404881313969695921
   }'
`, os.Args[0])
}

func sdkUtilitiesMintParamsUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] sdk-utilities mint-params -message JSON

MintParams implements mintParams.
    -message JSON: 

Example:
    %[1]s sdk-utilities mint-params --message '{
      "chainName": "Nihil voluptatem.",
      "port": 5396156918032315178
   }'
`, os.Args[0])
}

func sdkUtilitiesMintAnnualProvisionUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] sdk-utilities mint-annual-provision -message JSON

MintAnnualProvision implements mintAnnualProvision.
    -message JSON: 

Example:
    %[1]s sdk-utilities mint-annual-provision --message '{
      "chainName": "Ipsam assumenda perferendis.",
      "port": 5815409125182454944
   }'
`, os.Args[0])
}

func sdkUtilitiesMintEpochProvisionsUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] sdk-utilities mint-epoch-provisions -message JSON

MintEpochProvisions implements mintEpochProvisions.
    -message JSON: 

Example:
    %[1]s sdk-utilities mint-epoch-provisions --message '{
      "chainName": "Commodi vel.",
      "port": 1645460152882063013
   }'
`, os.Args[0])
}

func sdkUtilitiesDelegatorRewardsUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] sdk-utilities delegator-rewards -message JSON

DelegatorRewards implements delegatorRewards.
    -message JSON: 

Example:
    %[1]s sdk-utilities delegator-rewards --message '{
      "addresHex": "Vero quis.",
      "bech32Prefix": "Minima eligendi itaque.",
      "chainName": "Et expedita quibusdam.",
      "port": 7735830947852742027
   }'
`, os.Args[0])
}

func sdkUtilitiesEstimateFeesUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] sdk-utilities estimate-fees -message JSON

EstimateFees implements estimateFees.
    -message JSON: 

Example:
    %[1]s sdk-utilities estimate-fees --message '{
      "chainName": "Rerum est vel.",
      "port": 5512773892388447471,
      "txBytes": "UXVvZCB2b2x1cHRhdHVtIHN1bnQgdmVuaWFtIG5vbi4="
   }'
`, os.Args[0])
}

func sdkUtilitiesStakingParamsUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] sdk-utilities staking-params -message JSON

StakingParams implements stakingParams.
    -message JSON: 

Example:
    %[1]s sdk-utilities staking-params --message '{
      "chainName": "Aut consequatur sequi eveniet nisi aut.",
      "port": 422519322067561036
   }'
`, os.Args[0])
}

func sdkUtilitiesStakingPoolUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] sdk-utilities staking-pool -message JSON

StakingPool implements stakingPool.
    -message JSON: 

Example:
    %[1]s sdk-utilities staking-pool --message '{
      "chainName": "Sint numquam dolore et nemo quia quia.",
      "port": 1379536919246739616
   }'
`, os.Args[0])
}

func sdkUtilitiesEmoneyInflationUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] sdk-utilities emoney-inflation -message JSON

EmoneyInflation implements emoneyInflation.
    -message JSON: 

Example:
    %[1]s sdk-utilities emoney-inflation --message '{
      "chainName": "Officiis quo aperiam sed aut.",
      "port": 3236519200905233883
   }'
`, os.Args[0])
}

func sdkUtilitiesBudgetParamsUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] sdk-utilities budget-params -message JSON

BudgetParams implements budgetParams.
    -message JSON: 

Example:
    %[1]s sdk-utilities budget-params --message '{
      "chainName": "Nesciunt laboriosam quia et ratione rerum nihil.",
      "port": 2751582522214065040
   }'
`, os.Args[0])
}

func sdkUtilitiesDistributionParamsUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] sdk-utilities distribution-params -message JSON

DistributionParams implements distributionParams.
    -message JSON: 

Example:
    %[1]s sdk-utilities distribution-params --message '{
      "chainName": "Autem impedit quae excepturi vel quod corporis.",
      "port": 7190931602230024663
   }'
`, os.Args[0])
}

func sdkUtilitiesOsmoPoolsUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] sdk-utilities osmo-pools -message JSON

OsmoPools implements osmoPools.
    -message JSON: 

Example:
    %[1]s sdk-utilities osmo-pools --message '{
      "chainName": "Suscipit sequi exercitationem.",
      "port": 4611927448828070193
   }'
`, os.Args[0])
}

func sdkUtilitiesCrescentPoolsUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] sdk-utilities crescent-pools -message JSON

CrescentPools implements crescentPools.
    -message JSON: 

Example:
    %[1]s sdk-utilities crescent-pools --message '{
      "chainName": "Quia atque id est expedita rerum.",
      "port": 588797028473812858
   }'
`, os.Args[0])
}
