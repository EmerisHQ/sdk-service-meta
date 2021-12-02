package sdkutilities

import "fmt"

// String implements fmt.Stringer interface on Coins.
func (c *Coin) String() string {
	return fmt.Sprintf("%s%s", c.Amount, c.Denom)
}

// Coins represents a slice of Coin.
type Coins []Coin

// String implements fmt.Stringer interface on Coins.
// Taken from Cosmos SDK implementation for compatibility purposes.
func (coins Coins) String() string {
	if len(coins) == 0 {
		return ""
	}

	out := ""
	for _, coin := range coins {
		out += fmt.Sprintf("%v,", coin.String())
	}

	return out[:len(out)-1]
}
