package tracelistener

import (
	. "goa.design/goa/v3/dsl"
)

func init() {
	registerServiceDefinition(delegationsDesign)
}

var DelegationDesignType = Type("Delegation", func() {
	Description("Staking delegation as unmarshaled from trace bytes")

	Field(1, "delegator", String)
	Field(2, "validator", String)
	Field(3, "amount", String)
})

func delegationsDesign() {
	Method("delegation", func() {
		Payload(func() {
			Field(1, "payload", ArrayOf(TracePayload))
		})

		Result(ArrayOf(DelegationDesignType))

		GRPC(func() {})
	})
}
