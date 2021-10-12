package tracelistener

import (
	. "goa.design/goa/v3/dsl"
)

func init() {
	registerServiceDefinition(validatorDesign)
}

const (
	TypeDeleteValidator = "delete"
	TypeCreateValidator = "create"
)

var ValidatorDesignType = Type("Validator", func() {
	Description("Validator struct as unmarshaled from trace bytes")

	Field(1, "operatorAddress", String, func() {
		Default("")
	})

	Field(2, "consensusPubKeyType", String, func() {
		Default("")
	})

	Field(3, "consensusPubKeyValue", Bytes, func() {
		Default([]byte{})
	})

	Field(4, "jailed", Boolean, func() {
		Default(false)
	})

	Field(5, "status", Int32, func() {
		Default(0)
	})

	Field(6, "tokens", String, func() {
		Default("")
	})

	Field(7, "delegatorShares", String, func() {
		Default("")
	})

	Field(8, "moniker", String, func() {
		Default("")
	})

	Field(9, "identity", String, func() {
		Default("")
	})

	Field(10, "website", String, func() {
		Default("")
	})

	Field(11, "securityContact", String, func() {
		Default("")
	})

	Field(12, "details", String, func() {
		Default("")
	})

	Field(13, "unbondingHeight", Int64, func() {
		Default(0)
	})

	Field(14, "unbondingTime", Int64, func() {
		Default(0)
	})

	Field(15, "commissionRate", String, func() {
		Default("")
	})

	Field(16, "maxRate", String, func() {
		Default("")
	})

	Field(17, "maxChangeRate", String, func() {
		Default("")
	})

	Field(18, "updateTime", String, func() {
		Default("")
	})

	Field(19, "minSelfDelegation", String, func() {
		Default("")
	})

	Field(20, "type", String, func() {
		Default(TypeCreateValidator)
		Enum(TypeDeleteValidator, TypeCreateValidator)
	})
})

func validatorDesign() {
	Method("validator", func() {
		Payload(func() {
			Field(1, "payload", ArrayOf(TracePayload))
		})

		Result(ArrayOf(ValidatorDesignType))

		defineProcessingError()

		GRPC(func() {
			Response("ProcessingError", CodeInvalidArgument)
		})
	})
}
