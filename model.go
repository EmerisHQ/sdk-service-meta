package sdkservicemeta

import (
	. "goa.design/goa/v3/dsl"
)

var Coin = Type("Coin", func() {
	Description("SDK service representation of a Cosmos SDK types.Coin")

	Field(1, "denom", String)
	Field(2, "amount", String)

	Required("denom", "amount")
})

var Supply = Type("Supply", func() {
	Description("Supply of a given Cosmos SDK chain")

	Field(1, "coins", ArrayOf(Coin))

	Required("coins")
})

var TransactionResult = Type("TransactionResult", func() {
	Description("Result of a transaction broadcast operation")

	Field(1, "hash", String)

	Required("hash")
})

var IBCHeight = Type("IBCHeight", func() {
	Description("The plain type associated with ibc-go types.Height struct")

	Field(1, "revisionNumber", UInt64)
	Field(2, "revisionHeight", UInt64)
})

var IBCTransferMetadata = Type("IBCTransferMetadata", func() {
	Description("Metadata related to an IBC Transfer")

	Field(1, "sourcePort", String)
	Field(2, "sourceCannel", String)
	Field(3, "token", Coin)
	Field(4, "sender", String)
	Field(5, "receiver", String)
	Field(6, "timeoutHeight", IBCHeight)
	Field(7, "tiemoutTimestamp", UInt64)
})

var MsgMetadata = Type("MsgMetadata", func() {
	Description("Metadata related to some message contained in a transaction")

	Field(1, "msgType", String)

	Field(2, "ibcTransferMetadata", IBCTransferMetadata)

	Required("msgType")
})

var TxMessagesMetadata = Type("TxMessagesMetadata", func() {
	Description("Metadata for all messages contained in a transaction")

	Field(1, "messagesMetadata", ArrayOf(MsgMetadata))

	Required("messagesMetadata")
})
