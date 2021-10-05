package tracelistener

import (
	. "goa.design/goa/v3/dsl"
)

func init() {
	registerServiceDefinition(ibcDenomTracesDesign)
}

var IBCDenomTraceDesignType = Type("IBCDenomTrace", func() {
	Description("IBC denomination trace as unmarshaled from trace bytes")

	Field(1, "path", String)
	Field(2, "baseDenom", String)
	Field(3, "hash", String)
})

func ibcDenomTracesDesign() {
	Method("ibc_denom_trace", func() {
		Payload(func() {
			Field(1, "payload", TracePayload)
		})

		Result(IBCDenomTraceDesignType)

		GRPC(func() {})
	})
}
