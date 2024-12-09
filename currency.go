package moneylib

import (
	"fmt"
)

type currTmp struct {
	Code   string `json:"code"`
	Number uint   `json:"number"`
	Symbol string `json:"sym"`
	NOD    uint   `json:"nod"` // Number of Decimals
}

// Currency type.
type Currency struct {
	code   string
	number uint
	symbol string
	nod    uint // Number of Decimals
}

// Code returns currency code.
func (c *Currency) Code() string {
	return c.code
}

// Number returns currency number.
func (c *Currency) Number() uint {
	return c.number
}

// Symbol returns currency symbol.
func (c *Currency) Symbol() string {
	return c.symbol
}

// NOD returns currency's Number of Decimal.
func (c *Currency) NOD() uint {
	return c.nod
}

func (c *Currency) String() string {
	return fmt.Sprintf("Currency{%v; %v; %v; %v}",
		c.code, c.number, c.symbol, c.nod)
}

func newCurrency(c *currTmp) Currency {
	o := Currency{code: c.Code, number: c.Number, symbol: c.Symbol, nod: c.NOD}
	return o
}
