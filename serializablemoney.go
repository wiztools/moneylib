package moneylib

// SerializableMoney used for JSON serialization.
type SerializableMoney struct {
	Value    string `json:"value"`
	Currency string `json:"currency"`
}

// SerializableDispMoney with money value in human-displayable format.
type SerializableDispMoney struct {
	SerializableMoney
	Display string `json:"display"`
}

// Money comment.
func (m *SerializableMoney) Money() (*Money, error) {
	out, err := NewMoneyStr(m.Currency, m.Value)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SerializableMoney returns SerializableMoney instance of Money.
func (m *Money) SerializableMoney() SerializableMoney {
	var out SerializableMoney
	out.Value = m.String()
	curr := m.Currency()
	out.Currency = curr.Code()
	return out
}

// SerializableDispMoney returns SerializableDispMoney instance of Money.
func (m *Money) SerializableDispMoney() SerializableDispMoney {
	var out SerializableDispMoney
	out.Value = m.String()
	out.Display = m.StringHuman()
	curr := m.Currency()
	out.Currency = curr.Code()
	return out
}
