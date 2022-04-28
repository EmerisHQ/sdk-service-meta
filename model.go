package sdkservicemeta

import (
	. "goa.design/goa/v3/dsl"
)

var Coin = Type("Coin", func() {
	Description("SDK service representation of a Cosmos SDK types.Coin")

	Attribute("denom")
	Field(1, "denom", String)
	Field(2, "amount", String)

	Required("denom", "amount")
})

var Pagination = Type("Pagination", func() {
	Description("Pagination used in the Cosmos SDK")

	Field(1, "next_key", String)
	Field(2, "total", String)
})

var Supply = Type("Supply", func() {
	Description("Supply of a given Cosmos SDK chain")

	Field(1, "coins", ArrayOf(Coin))
	Field(2, "pagination", Pagination)

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

var MintEpochProvisions = Type("MintEpochProvisions", func() {
	Description("Mint epoch provisions response")

	Field(1, "mintEpochProvisions", Bytes)

	Required("mintEpochProvisions")
})

var AccountNumbers = Type("AccountNumbers", func() {
	Description("Account and sequence numbers for a given account")

	Field(1, "accountNumber", Int64)
	Field(2, "sequenceNumber", Int64)
	Field(3, "bech32Address", String)

	Required("accountNumber", "sequenceNumber", "bech32Address")
})

var DelegatorRewards = Type("DelegatorRewards", func() {
	Description("Delegator reward for a given address")

	Field(1, "rewards", ArrayOf(DelegationDelegatorReward))
	Field(2, "total", ArrayOf(Coin))
})

var DelegationDelegatorReward = Type("DelegationDelegatorReward", func() {
	Description("Delegation reward as defined in the Cosmos SDK")

	Field(1, "validatorAddress", String)
	Field(2, "rewards", ArrayOf(Coin))

	Required("validatorAddress", "rewards")
})

var Simulation = Type("Simulation", func() {
	Description("Results from the Cosmos SDK transaction simulation endpoint")

	Field(1, "gasWanted", UInt64)
	Field(2, "gasUsed", UInt64)
	Field(3, "fees", ArrayOf(Coin))

	Required("gasWanted", "gasUsed")
})

var StakingParams = Type("StakingParams", func() {
	Description("Staking params response")
	Field(1, "stakingParams", Bytes)
	Required("stakingParams")
})

var StakingPool = Type("StakingPool", func() {
	Description("List of staking pool")
	Field(1, "stakingPool", Bytes)
	Required("stakingPool")
})

var EmoneyInflation = Type("EmoneyInflation", func() {
	Description("e-money specific inflation parameters")
	Field(1, "state", EmoneyState)

	Required("state")
})

var EmoneyState = Type("EmoneyState", func() {
	Description("e-money specific inflation parameters (state)")
	Field(1, "last_applied", String)
	Field(2, "last_applied_height", String)
	Field(3, "assets", ArrayOf(EmoneyAsset))

	Required("last_applied", "last_applied_height", "assets")
})

var EmoneyAsset = Type("EmoneyAsset", func() {
	Description("e-money specific inflation parameters (asset)")
	Field(1, "denom", String)
	Field(2, "inflation", String)
	Field(3, "accum", String)

	Required("denom", "inflation", "accum")
})

var BudgetParams = Type("BudgetParams", func() {
	Description("Budget params response")
	Field(1, "budgetParams", Bytes)
	Required("budgetParams")
})

var DistributionParams = Type("DistributionParams", func() {
	Description("Distribution params response")
	Field(1, "communityTax", String)
	Field(2, "baseProposerReward", String)
	Field(3, "bonusProposerReward", String)
	Field(4, "withdrawAddrEnabled", Boolean)

	Required("communityTax", "baseProposerReward", "bonusProposerReward", "withdrawAddrEnabled")
})
