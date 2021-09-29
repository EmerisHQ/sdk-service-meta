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
