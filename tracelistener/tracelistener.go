package tracelistener

import (
	. "goa.design/goa/v3/dsl"
)

var serviceDefinitions = []func(){}

// Operation is a kind of operations a TraceWatcher observes.
type Operation string

var (
	// WriteOp is a write trace operation
	WriteOp Operation = "write"

	// DeleteOp is a write trace operation
	DeleteOp Operation = "delete"

	// ReadOp is a write trace operation
	ReadOp Operation = "read"

	// IterRangeOp is a write trace operation
	IterRangeOp Operation = "iterRange"
)

var TracePayload = Type("TracePayload", func() {
	Description("Data read off Cosmos SDK tracing facility")

	Field(1, "key", Bytes)
	Field(2, "value", Bytes)
	Field(3, "operationType", func() {
		Enum(WriteOp, DeleteOp, ReadOp, IterRangeOp)
	})

	Required("key", "value")
})

var ErrorObjectDesignType = Type("ErrorObject", func() {
	Field(1, "value", String, func() {
		Default("")
	})

	Field(2, "payloadIndex", Int, func() {
		Default(0)
	})
})

func defineProcessingError() {
	Error("ProcessingError", func() {
		Description("ProcessingError is a set of indexed error strings, where the index matches a given payload index")

		Field(1, "Name", String, "Name of error")
		Field(2, "Code", Int, "Code of error")
		Field(3, "errors", ArrayOf(ErrorObjectDesignType))
	})
}

func registerServiceDefinition(sd func()) {
	serviceDefinitions = append(serviceDefinitions, sd)
}

// ServiceDefinitions calls each registered service definition to get assembled in the
// goa template.
func ServiceDefinitions() {
	for _, sd := range serviceDefinitions {
		sd()
	}
}
