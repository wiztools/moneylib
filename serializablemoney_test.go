package moneylib

import "testing"

func TestSerializable(t *testing.T) {
	exp, _ := NewMoneyStr("SGD", "11.24")
	s := exp.SerializableMoney()
	res, _ := s.Money()
	if exp.Whole() != res.Whole() {
		t.Fail()
	}
}
