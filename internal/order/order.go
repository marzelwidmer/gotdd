package order

import "github.com/Rhymond/go-money"

type Order struct {
	ID                string
	Items             []Item
	CurrencyAlphacode string
}
type Item struct {
	ID        string
	Quantity  uint
	UnitPrice *money.Money
}

func (order Order) ComputeTotale(*money.Money, error) {
	amount := money.New(0, order.CurrencyAlphacode)
	for _, item := range order.Items {
		amount.Add(item.UnitPrice)
	}
}
