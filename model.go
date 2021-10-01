package sdkservicemeta

import (
	. "goa.design/goa/v3/dsl"
)

var Coin = Type("Coin", func() {
	Description("SDK service representation of a Cosmos SDK types.Coin")

	Attribute("denom", String)
	Attribute("amount", String)

	Field(1, "denom", String)
	Field(2, "amount", String)

	Required("denom", "amount")
})

var Supply = Type("Supply", func() {
	Description("Supply of a given Cosmos SDK chain")

	Attribute("coins", ArrayOf(Coin))

	Field(1, "coins", ArrayOf(Coin))

	Required("coins")
})

var TransactionResult = Type("TransactionResult", func() {
	Description("Result of a transaction broadcast operation")

	Attribute("hash", String)
	Attribute("error", String)

	Field(1, "hash", String)
	Field(2, "error", String)

	Required("hash")
})

var IBCHeight = Type("IBCHeight", func() {
	Description("The plain type associated with ibc-go types.Height struct")

	Attribute("revisionNumber", UInt64)
	Field(1, "revisionNumber", UInt64)

	Attribute("revisionHeight", UInt64)
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

var TxMetadata = Type("TxMetadata", func() {
	Description("Metadata related to some transaction bytes")

	Attribute("txType", String)
	Field(1, "txType", String)

	Attribute("ibcTransferMetadata", IBCTransferMetadata)
	Field(2, "ibcTransferMetadata", IBCTransferMetadata)

	Required("txType")
})
