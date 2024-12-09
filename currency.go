package moneylib

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
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

var currMap = make(map[string]Currency)

func init() {
	confFile := "conf/currency.json"
	b, err := os.ReadFile(confFile)
	if err != nil {
		log.Fatalf("Cannot load config: %v.", confFile)
	}
	currencies := make([]currTmp, 0)
	err = json.Unmarshal(b, &currencies)
	if err != nil {
		log.Fatalf("Error unmarshaling: %v.\n%v\n", confFile, err)
	}
	for _, v := range currencies {
		currMap[v.Code] = newCurrency(&v)
	}
}

// GetCurrency returns the currency object for the code.
func GetCurrency(code string) (Currency, error) {
	curr := currMap[code]
	if curr.code != code {
		return curr, errors.New("Currency code not available: " + code)
	}
	return currMap[code], nil
}

// ValidCurrency checks if the currency code is correct
func ValidCurrency(code string) bool {
	_, err := GetCurrency(code)
	return err == nil
}
