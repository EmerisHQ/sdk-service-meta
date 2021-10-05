package tracelistener

import (
	. "goa.design/goa/v3/dsl"
)

var serviceDefinitions = []func(){}

var TracePayload = Type("TracePayload", func() {
	Description("Data read off Cosmos SDK tracing facility")

	Field(1, "key", Bytes)
	Field(2, "value", Bytes)

	Required("key", "value")
})

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
