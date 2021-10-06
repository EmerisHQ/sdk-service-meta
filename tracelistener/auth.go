package tracelistener

import (
	. "goa.design/goa/v3/dsl"
)

func init() {
	registerServiceDefinition(authDesign)
}

var AuthDesignType = Type("Auth", func() {
	Description("Account as unmarshaled from trace bytes")

	Field(1, "address", String, func() {
		Default("")
	})

	Field(2, "sequenceNumber", UInt64, func() {
		Default(0)
	})

	Field(3, "accountNumber", UInt64, func() {
		Default(0)
	})
})

func authDesign() {
	Method("auth", func() {
		Payload(func() {
			Field(1, "payload", ArrayOf(TracePayload))
		})

		Result(ArrayOf(AuthDesignType))

		defineProcessingError()

		GRPC(func() {
			Response("ProcessingError", CodeInvalidArgument)
		})
	})
}
