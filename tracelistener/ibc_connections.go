package tracelistener

import (
	. "goa.design/goa/v3/dsl"
)

func init() {
	registerServiceDefinition(ibcConnectionsDesign)
}

var IBCConnectionDesignType = Type("IBCConnection", func() {
	Description("IBC connection as unmarshaled from trace bytes")

	Field(1, "connectionID", String, func() {
		Default("")
	})

	Field(2, "clientID", String, func() {
		Default("")
	})

	Field(3, "state", String, func() {
		Default("")
	})

	Field(4, "counterConnectionID", String, func() {
		Default("")
	})

	Field(5, "counterClientID", String, func() {
		Default("")
	})
})

func ibcConnectionsDesign() {
	Method("ibc_connection", func() {
		Payload(func() {
			Field(1, "payload", ArrayOf(TracePayload))
		})

		Result(ArrayOf(IBCConnectionDesignType))

		defineProcessingError()

		GRPC(func() {})
	})
}
