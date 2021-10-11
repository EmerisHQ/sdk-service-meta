package tracelistener

import (
	. "goa.design/goa/v3/dsl"
)

func init() {
	registerServiceDefinition(unbondingDelegationsDesign)
}

const (
	TypeDeleteUnbondingDelegation = "delete"
	TypeCreateUnbondingDelegation = "create"
)

var UnbondingDelegationDesignType = Type("UnbondingDelegation", func() {
	Description("Staking unbondingDelegation as unmarshaled from trace bytes")

	Field(1, "delegator", String, func() {
		Default("")
	})

	Field(2, "validator", String, func() {
		Default("")
	})

	Field(3, "type", String, func() {
		Default(TypeCreateUnbondingDelegation)
		Enum(TypeDeleteUnbondingDelegation, TypeCreateUnbondingDelegation)
	})

	Field(4, "entries", ArrayOf(UnbondingDelegationEntryDesignType))
})

var UnbondingDelegationEntryDesignType = Type("UnbondingDelegationEntry", func() {
	Description("A bonded user unbonding delegation entry")

	Field(1, "balance", String, func() {
		Default("")
	})

	Field(2, "initialBalance", String, func() {
		Default("")
	})

	Field(3, "creationHeight", Int64, func() {
		Default(0)
	})

	Field(4, "completionTime", Int64, func() {
		Default(0)
	})
})

func unbondingDelegationsDesign() {
	Method("unbondingDelegation", func() {
		Payload(func() {
			Field(1, "payload", ArrayOf(TracePayload))
		})

		Result(ArrayOf(UnbondingDelegationDesignType))

		defineProcessingError()

		GRPC(func() {
			Response("ProcessingError", CodeInvalidArgument)
		})
	})
}
