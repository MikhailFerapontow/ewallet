package repository

import (
	"fmt"
	"server/models"
	"time"

	"github.com/jmoiron/sqlx"
)

type ApiPostgres struct {
	db *sqlx.DB
}

func newApiPostgres(db *sqlx.DB) *ApiPostgres {
	return &ApiPostgres{
		db: db,
	}
}

func (r *ApiPostgres) NewWallet() (models.Wallet, error) {
	var wallet models.Wallet

	query := fmt.Sprintf("INSERT INTO %s (balance) VALUES ('100') RETURNING id, balance", walletTable)
	row := r.db.QueryRow(query)
	if err := row.Scan(&wallet.Id, &wallet.Balance); err != nil {
		return models.Wallet{}, err
	}
	return wallet, nil
}

func (r *ApiPostgres) GetWallet(id string) (models.Wallet, error) {
	var wallet models.Wallet

	query := fmt.Sprintf("SELECT * FROM %s WHERE id = $1", walletTable)
	err := r.db.Get(&wallet, query, id)
	return wallet, err
}

func (r *ApiPostgres) SendMoney(fromId string, input models.SendMoneyInput) error {
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	currentRFC3339Time := time.Now().Format(time.RFC3339)
	createTransactionQuery := fmt.Sprintf("INSERT INTO %s (\"time\", \"from\", \"to\", \"amount\") VALUES ($1, $2, $3,  $4);", transactionTable)
	_, err = tx.Exec(createTransactionQuery, currentRFC3339Time, fromId, input.To, input.Amount)
	if err != nil {
		return err
	}

	removeFromBalanceQuery := fmt.Sprintf("UPDATE %s SET balance = balance - $1 WHERE id = $2", walletTable)
	_, err = tx.Exec(removeFromBalanceQuery, input.Amount, fromId)
	if err != nil {
		return err
	}

	addToBalanceQuery := fmt.Sprintf("UPDATE %s SET balance = balance + $1 WHERE id = $2", walletTable)
	_, err = tx.Exec(addToBalanceQuery, input.Amount, input.To)
	if err != nil {
		return err
	}

	return tx.Commit()
}

func (r *ApiPostgres) GetHistory(id string) ([]models.Transaction, error) {
	var transactions []models.Transaction

	query := fmt.Sprintf("SELECT * FROM %s WHERE \"from\" = $1 OR \"to\" = $1", transactionTable)
	err := r.db.Select(&transactions, query, id)

	return transactions, err
}
