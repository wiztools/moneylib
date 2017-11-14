package moneylib

import "testing"
import "fmt"

func TestCurrency(t *testing.T) {
	inr, _ := GetCurrency("INR")
	fmt.Println(inr.String())
	if inr.Number() != 356 {
		t.Fail()
	}
}

func TestCurrencyInvalid(t *testing.T) {
	_, err := GetCurrency("inr")
	if err == nil {
		t.Fail()
	}
}

func TestValidCurrency(t *testing.T) {
	if ValidCurrency("♥") {
		t.Fail()
	}
	if !ValidCurrency("PAB") {
		t.Fail()
	}
}
