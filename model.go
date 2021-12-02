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
	Field(2, "sourceChannel", String)
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

var BlockData = Type("BlockData", func() {
	Description("Data for block at a given height")

	Field(1, "height", Int64)
	Field(2, "block", Bytes)

	Required("height")
	Required("block")
})

var LiquidityParams = Type("LiquidityParams", func() {
	Description("Liquidity params response")

	Field(1, "liquidityParams", Bytes)

	Required("liquidityParams")
})

var LiquidityPools = Type("LiquidityPools", func() {
	Description("Liquidity pools response")

	Field(1, "liquidityPools", Bytes)

	Required("liquidityPools")
})

var MintParams = Type("MintParams", func() {
	Description("Mint params response")

	Field(1, "mintParams", Bytes)

	Required("mintParams")
})

var MintInflation = Type("MintInflation", func() {
	Description("Mint inflation response")

	Field(1, "mintInflation", Bytes)

	Required("mintInflation")
})

var MintAnnualProvision = Type("MintAnnualProvision", func() {
	Description("Mint annual provision response")

	Field(1, "mintAnnualProvision", Bytes)

	Required("mintAnnualProvision")
})

var AccountNumbers = Type("AccountNumbers", func() {
	Description("Account and sequence numbers for a given account")

	Field(1, "accountNumber", Int64)
	Field(2, "sequenceNumber", Int64)
	Field(3, "bech32Address", String)

	Required("accountNumber", "sequenceNumber", "bech32Address")
})
