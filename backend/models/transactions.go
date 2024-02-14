package models

type Transaction struct {
	Time   string  `json:"time"`
	From   string  `json:"from"`
	To     string  `json:"to"`
	Amount float64 `json:"amount"`
}
