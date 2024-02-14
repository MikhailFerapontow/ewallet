package models

type Wallet struct {
	Id      string  `json:"id" binding:"required" db:"id"`
	Balance float64 `json:"balance"`
}
