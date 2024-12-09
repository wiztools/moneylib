package moneylib

import (
	"errors"
	"fmt"
	"math/big"
	"regexp"
	"strconv"
	"strings"

	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

// Money type. Use NewMoney*() functions to instantiate.
type Money struct {
	intPart string
	decPart string
	curr    Currency
}

// Currency returns the currency of the money.
func (m Money) Currency() Currency {
	return m.curr
}

// Float converts to Money to big.Float
func (m *Money) Float() *big.Float {
	// On magic number 53: https://golang.org/pkg/math/big/#NewFloat
	out, _, _ := big.ParseFloat(m.String(), 10, 53, big.ToNearestAway)
	return out
}

// Whole returns the whole number representation of the currency.
func (m *Money) Whole() uint64 {
	wholeStr := m.intPart + m.decPart
	o, _ := strconv.ParseUint(wholeStr, 10, 64)
	return o
}

// String implements fmt.Stringer
func (m *Money) String() string {
	curr := m.Currency()
	if curr.NOD() == 0 {
		return fmt.Sprintf("%v", m.intPart)
	}
	return fmt.Sprintf("%v.%v", m.intPart, m.decPart)
}

func (m *Money) HumanEN() string {
	return m.Human(language.English)
}

func (m *Money) Human(lan language.Tag) string {
	p := message.NewPrinter(lan)
	if m.Currency().nod == 0 {
		return p.Sprintf("%s%d", m.curr.symbol, m.Whole())
	}
	// Format the money value to the exact decimal value defined
	// for the currency:
	format := fmt.Sprintf("%%.%df", m.Currency().nod)
	value, _ := m.Float().Float64()
	return p.Sprintf("%s"+format, m.curr.symbol, value)
}

// StringHuman returns money representation for humans.
// func (m *Money) StringHuman() string {
// 	curr := m.Currency()
// 	f, _ := m.Float().Float64()
// 	o := money.Format(f, money.Options{"currency": curr.Code()})
// 	return o
// }

func pow10(x uint) uint {
	if x == 0 {
		return 0
	}
	var res uint = 1
	for i := 0; i < int(x); i++ {
		res = res * 10
	}
	return res
}

// Group 1: Integer part
// Group 3: Decimal part
const moneyRE = `^([0-9]+)(\.([0-9]{1,5}))?$`

func getMoney(curr Currency, value string) (*Money, error) {
	re := regexp.MustCompile(moneyRE)
	if !re.MatchString(value) {
		return nil, errors.New("Wrong format for money: " + value)
	}
	outArr := re.FindStringSubmatch(value)

	// First, copy the integer part:
	intPart := outArr[1]

	// Next, the decimal part after validation:
	// Validate decimal precision is correct:
	if uint(len([]rune(outArr[3]))) != curr.NOD() {
		msg := fmt.Sprintf("Decimal precision incorrect: %v (%v - %v)", value, curr.Code(), curr.NOD())
		return nil, errors.New(msg)
	}
	decPart := outArr[3]

	m := &Money{curr: curr, intPart: intPart, decPart: decPart}

	return m, nil
}

// -----------------------------
// Constructor functions follow:
// -----------------------------

// NewMoneyPartUint returns the money instance.
func NewMoneyPartUint(currStr string, intPart uint64, decPart uint) (*Money, error) {
	curr, err := GetCurrency(currStr)
	if err != nil {
		return nil, err
	}
	nod := curr.NOD()
	if nod == 0 {
		return getMoney(curr, strconv.FormatUint(intPart, 10))
	}
	if decPart >= pow10(nod) {
		return nil, errors.New("currency decimal part does not fit currency NOD")
	}
	format := "%v.%0#" + fmt.Sprintf("%v", nod) + "d"
	mStr := fmt.Sprintf(format, intPart, decPart)
	return getMoney(curr, mStr)
}

// NewMoneyStr returns the money instance.
func NewMoneyStr(currStr string, value string) (*Money, error) {
	curr, err := GetCurrency(currStr)
	if err != nil {
		return nil, err
	}
	return getMoney(curr, value)
}

// NewMoneyPart returns the money instance.
func NewMoneyPart(currStr string, intPart string, decPart string) (*Money, error) {
	curr, err := GetCurrency(currStr)
	if err != nil {
		return nil, err
	}
	nod := curr.NOD()
	if decPart == "" || (len(decPart) != int(nod)) {
		return nil, errors.New("currency decimal part can be empty OR length of NOD")
	}
	var combined string
	if decPart == "" {
		combined = intPart
	} else {
		combined = intPart + "." + decPart
	}
	return getMoney(curr, combined)
}

// NewMoneyWhole returns the money instance.
func NewMoneyWhole(currStr string, whole uint64) (*Money, error) {
	curr, err := GetCurrency(currStr)
	if err != nil {
		return nil, err
	}
	format := "%0#" + fmt.Sprintf("%v", curr.NOD()) + "d"
	str := fmt.Sprintf(format, whole)
	if curr.NOD() == 0 {
		return getMoney(curr, str)
	}
	lenOfStr := len([]rune(str))
	locOfDec := lenOfStr - int(curr.NOD())
	if locOfDec >= 0 {
		outStr := str[:locOfDec] + "." + str[locOfDec:]
		if strings.HasPrefix(outStr, ".") {
			outStr = "0" + outStr
		}
		return getMoney(curr, outStr)
	}
	return nil, errors.New("invalid currency amount '" + str + "'")
}
