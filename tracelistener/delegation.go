package tracelistener

import (
	. "goa.design/goa/v3/dsl"
)

func init() {
	registerServiceDefinition(delegationsDesign)
}

const (
	TypeDeleteDelegation = "delete"
	TypeCreateDelegation = "create"
)

var DelegationDesignType = Type("Delegation", func() {
	Description("Staking delegation as unmarshaled from trace bytes")

	Field(1, "delegator", String, func() {
		Default("")
	})

	Field(2, "validator", String, func() {
		Default("")
	})

	Field(3, "amount", String, func() {
		Default("")
	})

	Field(4, "type", String, func() {
		Default(TypeCreateDelegation)
		Enum(TypeDeleteDelegation, TypeCreateDelegation)
	})
})

func delegationsDesign() {
	Method("delegation", func() {
		Payload(func() {
			Field(1, "payload", ArrayOf(TracePayload))
		})

		Result(ArrayOf(DelegationDesignType))

		defineProcessingError()

		GRPC(func() {})
	})
}
