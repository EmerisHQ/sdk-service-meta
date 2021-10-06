package tracelistener

import (
	. "goa.design/goa/v3/dsl"
)

func init() {
	registerServiceDefinition(ibcDenomTracesDesign)
}

var IBCDenomTraceDesignType = Type("IBCDenomTrace", func() {
	Description("IBC denomination trace as unmarshaled from trace bytes")

	Field(1, "path", String, func() {
		Default("")
	})

	Field(2, "baseDenom", String, func() {
		Default("")
	})

	Field(3, "hash", String, func() {
		Default("")
	})
})

func ibcDenomTracesDesign() {
	Method("ibc_denom_trace", func() {
		Payload(func() {
			Field(1, "payload", ArrayOf(TracePayload))
		})

		Result(ArrayOf(IBCDenomTraceDesignType))

		defineProcessingError()

		GRPC(func() {
			Response("ProcessingError", CodeInvalidArgument)
		})
	})
}
