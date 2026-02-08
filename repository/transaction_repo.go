package repository

import (
	"database/sql"
	"product-api/model"
)

type TransactionRepositoryInterface interface {
	Create(tx *sql.Tx, transaction *model.Transaction) error
	GetAll(fromDate string, toDate string) ([]model.Transaction, error)
}

type transactionRepository struct {
	db *sql.DB
}

func NewTransactionRepository(db *sql.DB) TransactionRepositoryInterface {
	return &transactionRepository{db: db}
}

func (repo *transactionRepository) Create(tx *sql.Tx, transaction *model.Transaction) (error) {
	query := "INSERT INTO transactions (total_amount) VALUES ($1) RETURNING id, created_at"
	err := tx.QueryRow(query, transaction.TotalAmount).Scan(&transaction.ID, &transaction.CreatedAt)
	if err != nil {
		return err
	}

	detailQuery := "INSERT INTO transaction_details (transaction_id, product_id, quantity, subtotal) VALUES ($1, $2, $3, $4) RETURNING id"
	for i := range transaction.Details {
		transaction.Details[i].TransactionID = transaction.ID
		err = tx.QueryRow(
			detailQuery,
			transaction.Details[i].TransactionID,
			transaction.Details[i].ProductID,
			transaction.Details[i].Quantity,
			transaction.Details[i].Subtotal,
		).Scan(&transaction.Details[i].ID)
		if err != nil {
			return err
		}
	}

	return nil
}

func (repo *transactionRepository) GetAll(fromDate string, toDate string) ([]model.Transaction, error) {
	query := "SELECT id, total_amount, created_at FROM transactions"
	var rows *sql.Rows
	var err error

	if fromDate != "" && toDate != "" {
		query += " WHERE created_at BETWEEN $1 AND $2 ORDER BY created_at DESC"
		rows, err = repo.db.Query(query, fromDate, toDate)
	} else {
		query += " ORDER BY created_at DESC"
		rows, err = repo.db.Query(query)
	}

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	transactions := make([]model.Transaction, 0)
	for rows.Next() {
		var transaction model.Transaction
		err := rows.Scan(&transaction.ID, &transaction.TotalAmount, &transaction.CreatedAt)
		if err != nil {
			return nil, err
		}

		detailQuery := "SELECT id, transaction_id, product_id, quantity, subtotal FROM transaction_details WHERE transaction_id = $1"
		detailRows, err := repo.db.Query(detailQuery, transaction.ID)
		if err != nil {
			return nil, err
		}

		details := make([]model.TransactionDetail, 0)
		for detailRows.Next() {
			var detail model.TransactionDetail
			err := detailRows.Scan(&detail.ID, &detail.TransactionID, &detail.ProductID, &detail.Quantity, &detail.Subtotal)
			if err != nil {
				detailRows.Close()
				return nil, err
			}
			details = append(details, detail)
		}
		detailRows.Close()
		transaction.Details = details
		transactions = append(transactions, transaction)
	}

	return transactions, nil
}
