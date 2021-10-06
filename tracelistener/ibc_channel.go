package tracelistener

import (
	. "goa.design/goa/v3/dsl"
)

func init() {
	registerServiceDefinition(ibcChannelsDesign)
}

var IBCChannelDesignType = Type("IBCChannel", func() {
	Description("IBC channel as unmarshaled from trace bytes")

	Field(1, "channelID", String)
	Field(2, "counterChannelID", String)
	Field(3, "hops", ArrayOf(String))
	Field(4, "port", String)
	Field(5, "state", Int32)
})

func ibcChannelsDesign() {
	Method("ibc_channel", func() {
		Payload(func() {
			Field(1, "payload", ArrayOf(TracePayload))
		})

		Result(ArrayOf(IBCChannelDesignType))

		defineProcessingError()

		GRPC(func() {})
	})
}
