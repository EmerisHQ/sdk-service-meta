package sdkservicemeta

import (
	"fmt"

	. "goa.design/goa/v3/dsl"
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

	Method("supply", func() {
		Payload(func() {
			_ = standardArguments()
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

		Result(TxMetadata)

		GRPC(func() {
			Response(CodeOK)
		})
	})

	Files("/openapi.json", "./gen/http/openapi.json")
})
