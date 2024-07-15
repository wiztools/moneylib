package moneylib

import "testing"

func TestSerializable(t *testing.T) {
	exp, err := NewMoneyStr("SGD", "11.24")
	if err != nil {
		t.Error(err)
		return
	}
	s := exp.SerializableMoney()
	res, _ := s.Money()
	if exp.Whole() != res.Whole() {
		t.Fail()
	}
}
