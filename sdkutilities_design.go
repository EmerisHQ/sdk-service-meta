package sdkservicemeta

import (
	"fmt"

	//nolint this is a DSL definition, nothing to worry about
	. "goa.design/goa/v3/dsl"

	// we need the init() side effect of this import
	_ "goa.design/plugins/v3/zaplogger"
)

const (
	sdkUtilities = "sdk-utilities"
)

func standardArguments() int {
	Field(1, "chainName", String, "Chain to get data from")
	Field(2, "port", Int, "gRPC port for selected chain, defaults to 9090")
	Required("chainName")

	return 3
}

var _ = API(sdkUtilities, func() {
	Title("SDK utilities")
	Description(
		"Service for execution of various Cosmos SDK-related functionalities, with a focus on version-specific types and unmarshaling",
	)
	Server(sdkUtilities, func() {
		Host("localhost", func() {
			URI("http://localhost:8080")
			URI("grpc://localhost:9090")
		})
	})
})

var _ = Service(sdkUtilities, func() {
	Description(fmt.Sprintf("%s performs Cosmos SDK-related operations", sdkUtilities))

	Method("accountNumbers", func() {
		Payload(func() {
			nextFieldIdx := standardArguments()
			Field(nextFieldIdx, "bech32Prefix", String, "bech32-encoded prefix of the account")

			nextFieldIdx++

			Field(nextFieldIdx, "addresHex", String, "Hex-encoded address, without bech32 hrp")
		})

		Result(AccountNumbers)

		GRPC(func() {
			Response(CodeOK)
		})
	})

	Method("supply", func() {
		Payload(func() {
			nextFieldIdx := standardArguments()

			Field(nextFieldIdx, "paginationKey", String, "pagination key")
		})

		Result(Supply)

		GRPC(func() {
			Response(CodeOK)
		})
	})

	Method("supplyDenom", func() {
		Payload(func() {
			nextFieldIdx := standardArguments()

			Field(nextFieldIdx, "denom", String, "Denom name")
		})

		Result(Supply)

		GRPC(func() {
			Response(CodeOK)
		})
	})

	Method("queryTx", func() {
		Payload(func() {
			nextFieldIdx := standardArguments()
			Field(nextFieldIdx, "hash", String, "Transaction hash to query")

			Required("hash")
		})

		Result(Bytes)

		GRPC(func() {
			Response(CodeOK)
		})
	})

	Method("broadcastTx", func() {
		Payload(func() {
			nextFieldIdx := standardArguments()
			Field(nextFieldIdx, "txBytes", Bytes, "Transaction bytes")

			Required("txBytes")
		})

		Result(TransactionResult)

		GRPC(func() {
			Response(CodeOK)
		})
	})

	Method("txMetadata", func() {
		Payload(func() {
			Field(1, "txBytes", Bytes, "Transaction bytes")

			Required("txBytes")
		})

		Result(TxMessagesMetadata)

		GRPC(func() {
			Response(CodeOK)
		})
	})

	Method("block", func() {
		Payload(func() {
			nextFieldIdx := standardArguments()
			Field(nextFieldIdx, "height", Int64, "Height of the block to query")

			Required("height")
		})

		Result(BlockData)

		GRPC(func() {
			Response(CodeOK)
		})
	})

	Method("liquidityParams", func() {
		Payload(func() {
			_ = standardArguments()
		})

		Result(LiquidityParams)

		GRPC(func() {
			Response(CodeOK)
		})
	})

	Method("liquidityPools", func() {
		Payload(func() {
			_ = standardArguments()
		})

		Result(LiquidityPools)

		GRPC(func() {
			Response(CodeOK)
		})
	})

	Method("mintInflation", func() {
		Payload(func() {
			_ = standardArguments()
		})

		Result(MintInflation)

		GRPC(func() {
			Response(CodeOK)
		})
	})

	Method("mintParams", func() {
		Payload(func() {
			_ = standardArguments()
		})

		Result(MintParams)

		GRPC(func() {
			Response(CodeOK)
		})
	})

	Method("mintAnnualProvision", func() {
		Payload(func() {
			_ = standardArguments()
		})

		Result(MintAnnualProvision)

		GRPC(func() {
			Response(CodeOK)
		})
	})

	Method("mintEpochProvisions", func() {
		Payload(func() {
			_ = standardArguments()
		})

		Result(MintEpochProvisions)

		GRPC(func() {
			Response(CodeOK)
		})
	})

	Method("delegatorRewards", func() {
		Payload(func() {
			nextFieldIdx := standardArguments()

			Field(nextFieldIdx, "bech32Prefix", String, "bech32-encoded prefix of the account")

			nextFieldIdx++

			Field(nextFieldIdx, "addresHex", String, "Hex-encoded address, without bech32 hrp")

		})

		Result(DelegatorRewards)

		GRPC(func() {
			Response(CodeOK)
		})
	})

	Method("estimateFees", func() {
		Payload(func() {
			nextFieldIdx := standardArguments()
			Field(nextFieldIdx, "txBytes", Bytes, "Transaction bytes")

			Required("txBytes")
		})

		Result(Simulation)

		GRPC(func() {
			Response(CodeOK)
		})
	})

	Method("stakingParams", func() {
		Payload(func() {
			_ = standardArguments()
		})

		Result(StakingParams)

		GRPC(func() {
			Response(CodeOK)
		})
	})

	Method("stakingPool", func() {
		Payload(func() {
			_ = standardArguments()
		})

		Result(StakingPool)

		GRPC(func() {
			Response(CodeOK)
		})
	})

	Method("emoneyInflation", func() {
		Payload(func() {
			_ = standardArguments()
		})

		Result(EmoneyInflation)

		GRPC(func() {
			Response(CodeOK)
		})
	})

	Method("ibcChannelClientState", func() {
		Payload(func() {
			_ = standardArguments()
		})

		Result(IBCChannelClientState)

		GRPC(func() {
			Response(CodeOK)
		})
	})

	Files("/openapi.json", "./gen/http/openapi.json")
})
