package sdkutilities

import "fmt"

func (c *Coin) String() string {
	return fmt.Sprintf("%s%s", c.Amount, c.Denom)
}
