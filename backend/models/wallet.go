package models

type Wallet struct {
	Id      string  `json:"id"`
	Balance float64 `json:"balance" binding:"required"`
}
