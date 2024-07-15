package moneylib

import (
	"fmt"
	"math/big"
	"testing"
)

func TestString(t *testing.T) {
	o, err := NewMoneyStr("SGD", "123.02")
	if err != nil {
		t.Error(err)
		return
	}
	if o.String() != "123.02" {
		t.Fail()
	}
}

func TestStringHuman(t *testing.T) {
	o, err := NewMoneyStr("INR", "122123.02")
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Printf("Human: %v.\n", o.StringHuman())
	if o.StringHuman() != "₹122,123.02" {
		t.Fail()
	}
}

func TestString0Dec(t *testing.T) {
	o, err := NewMoneyStr("JPY", "123")
	if err != nil {
		t.Error(err)
		return
	}
	if o.String() != "123" {
		t.Fail()
	}
}

func TestBigFloat(t *testing.T) {
	o, err := NewMoneyStr("SGD", "123.02")
	if err != nil {
		t.Error(err)
		return
	}
	exp := big.NewFloat(123.02)
	fmt.Printf("Exp: %v; Actual: %v.\n", exp, o.Float())
	if exp.Cmp(o.Float()) != 0 {
		t.Fail()
	}
}

func TestBigFloat0Dec(t *testing.T) {
	o, err := NewMoneyStr("JPY", "123")
	if err != nil {
		t.Error(err)
		return
	}
	exp := big.NewFloat(123)
	fmt.Printf("Exp: %v; Actual: %v.\n", exp, o.Float())
	if exp.Cmp(o.Float()) != 0 {
		t.Fail()
	}
}

func TestNewMoneyStr(t *testing.T) {
	{
		_, err := NewMoneyStr("INR", "1") // INR requires 2 decimal precision
		if err == nil {
			t.Fail()
		}
	}
	{
		m, err := NewMoneyStr("INR", "0.01")
		if err != nil {
			t.Error(err)
			return
		}
		if m.Whole() != 1 {
			t.Fail()
		}
	}
	{
		m, err := NewMoneyStr("JPY", "1234")
		if err != nil {
			t.Error(err)
			return
		}
		if m.Whole() != 1234 {
			t.Fail()
		}
	}
	{
		m, err := NewMoneyStr("OMR", "123.456")
		if err != nil {
			t.Error(err)
			return
		}
		if m.Whole() != 123456 {
			t.Fail()
		}
	}
	{
		m, err := NewMoneyStr("OMR", "123.001")
		if err != nil {
			t.Error(err)
			return
		}
		if m.Whole() != 123001 {
			t.Fail()
		}
	}
}

func Test0Int(t *testing.T) {
	{
		m, err := NewMoneyStr("INR", "0.01")
		fmt.Println(m)
		if err != nil {
			t.Fail()
		}
	}
	{
		_, err := NewMoneyStr("INR", ".01")
		if err == nil {
			t.Fail()
		}
	}
}

func TestPart(t *testing.T) {
	{
		_, err := NewMoneyPart("INR", "12", "1") // Decimal part MUST equal NOD len.
		if err == nil {
			t.Fail()
		}
	}
}

func TestPartUint(t *testing.T) {
	{
		m, err := NewMoneyPartUint("INR", 123, 45)
		if err != nil {
			t.Error(err)
			return
		}
		if m.String() != "123.45" {
			t.Fail()
		}
	}
	{
		_, err := NewMoneyPartUint("INR", 123, 456)
		if err == nil {
			t.Fail()
		}
	}
	{
		m, err := NewMoneyPartUint("INR", 123, 99)
		if err != nil {
			t.Error(err)
			return
		}
		if m.String() != "123.99" {
			t.Fail()
		}
	}
	{
		m, err := NewMoneyPartUint("INR", 123, 0)
		if err != nil {
			t.Error(err)
			return
		}
		if m.String() != "123.00" {
			t.Fail()
		}
	}
	{
		m, err := NewMoneyPartUint("INR", 123, 1)
		if err != nil {
			t.Error(err)
			return
		}
		if m.String() != "123.01" {
			t.Fail()
		}
	}
}

func TestNewMoneyWhole(t *testing.T) {
	{
		m, err := NewMoneyWhole("INR", 123450)
		if err != nil {
			t.Error(err)
			return
		}
		if m.String() != "1234.50" {
			t.Fail()
		}
	}
	{
		m, err := NewMoneyWhole("INR", 1)
		if err != nil {
			t.Error(err)
			return
		}
		if m.String() != "0.01" {
			t.Fail()
		}
	}
}

func TestWrongFormat(t *testing.T) {
	{
		_, err := NewMoneyStr("JPY", "123.12")
		if err == nil {
			t.Fail()
		}
	}

	{
		_, err := NewMoneyStr("SGD", "1000.02x")
		if err == nil {
			t.Fail()
		}
	}

	{ // 0-padding is needed!
		_, err := NewMoneyStr("SGD", "1000.2")
		if err == nil {
			t.Fail()
		}
	}

	{
		_, err := NewMoneyPart("SGD", "1000", "021")
		if err == nil {
			t.Fail()
		}
	}
}
