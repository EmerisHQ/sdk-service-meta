package tracelistener

import (
	. "goa.design/goa/v3/dsl"
)

func init() {
	registerServiceDefinition(ibcClientStateDesign)
}

var IBCClientStateDesignType = Type("IBCClientState", func() {
	Description("IBC client state as unmarshaled from trace bytes")

	Field(1, "chainID", String, func() {
		Default("")
	})

	Field(2, "clientID", String, func() {
		Default("")
	})

	Field(3, "latestHeight", UInt64, func() {
		Default(0)
	})

	Field(4, "trustingPeriod", Int64, func() {
		Default(0)
	})
})

func ibcClientStateDesign() {
	Method("ibc_client_state", func() {
		Payload(func() {
			Field(1, "payload", ArrayOf(TracePayload))
		})

		Result(ArrayOf(IBCClientStateDesignType))

		defineProcessingError()

		GRPC(func() {
			Response("ProcessingError", CodeInvalidArgument)
		})
	})
}
