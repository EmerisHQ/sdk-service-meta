package tracelistener

import (
	. "goa.design/goa/v3/dsl"
)

func init() {
	registerServiceDefinition(ibcConnectionsDesign)
}

var IBCConnectionDesignType = Type("IBCConnection", func() {
	Description("IBC connection as unmarshaled from trace bytes")

	Field(1, "connectionID", String)
	Field(2, "clientID", String)
	Field(3, "state", String)
	Field(4, "counterConnectionID", String)
	Field(5, "counterClientID", String)
})

func ibcConnectionsDesign() {
	Method("ibc_connection", func() {
		Payload(func() {
			Field(1, "payload", TracePayload)
		})

		Result(IBCConnectionDesignType)

		GRPC(func() {})
	})
}
