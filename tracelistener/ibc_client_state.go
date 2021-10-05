package tracelistener

import (
	. "goa.design/goa/v3/dsl"
)

func init() {
	registerServiceDefinition(ibcClientStateDesign)
}

var IBCClientStateDesignType = Type("IBCClientState", func() {
	Description("IBC client state as unmarshaled from trace bytes")

	Field(1, "chainID", String)
	Field(2, "clientID", String)
	Field(3, "latestHeight", UInt64)
	Field(4, "trustingPeriod", Int64)
})

func ibcClientStateDesign() {
	Method("ibc_client_state", func() {
		Payload(func() {
			Field(1, "payload", TracePayload)
		})

		Result(IBCClientStateDesignType)

		GRPC(func() {})
	})
}
