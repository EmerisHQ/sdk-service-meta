package tracelistener

import (
	. "goa.design/goa/v3/dsl"
)

func init() {
	registerServiceDefinition(bankDesign)
}

var BalanceDesignType = Type("Balance", func() {
	Description("Balance of a given address as unmarshaled from trace bytes")

	Field(1, "address", String, func() {
		Default("")
	})

	Field(2, "amount", String, func() {
		Default("")
	})

	Field(3, "denom", String, func() {
		Default("")
	})

	Required("address", "amount", "denom")
})

func bankDesign() {
	Method("bank", func() {
		Payload(func() {
			Field(1, "payload", ArrayOf(TracePayload))
		})

		Result(ArrayOf(BalanceDesignType))

		defineProcessingError()

		GRPC(func() {
			Response("ProcessingError", CodeInvalidArgument)
		})
	})
}
