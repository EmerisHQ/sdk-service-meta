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
	Field(2, "error", String)

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

var TxMetadata = Type("TxMetadata", func() {
	Description("Metadata related to some transaction bytes")

	Field(1, "txType", String)

	Field(2, "ibcTransferMetadata", IBCTransferMetadata)

	Required("txType")
})
