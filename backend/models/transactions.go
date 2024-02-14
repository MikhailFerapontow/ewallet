package models

type Transaction struct {
	Time   string  `json:"time" db:"time"`
	From   string  `json:"from" db:"from"`
	To     string  `json:"to" db:"to"`
	Amount float64 `json:"amount" db:"amount"`
}

type SendMoneyInput struct {
	To     string  `json:"to"`
	Amount float64 `json:"amount"`
}
