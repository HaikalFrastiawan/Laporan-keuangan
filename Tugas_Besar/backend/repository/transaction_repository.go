package repository

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/HaikalFrastiawan/Laporan-keuangan/backend/models"
)

type TransactionRepository struct {
	DB *sql.DB
}

func NewTransactionRepository(db *sql.DB) *TransactionRepository {
	return &TransactionRepository{DB: db}
}

func (r *TransactionRepository) GetAllTransactions(ctx context.Context) ([]models.Transaction, error) {
	query := "SELECT transactions_id, deskripsi, jumlah, tanggal, jenis, catatan FROM transactions ORDER BY tanggal DESC"
	rows, err := r.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("gagal mengambil data transaksi: %v", err)
	}
	defer rows.Close()

	var transactions []models.Transaction
	for rows.Next() {
		var t models.Transaction
		if err := rows.Scan(&t.ID, &t.Description, &t.Amount, &t.Date, &t.Jenis, &t.Catatan); err != nil {
			return nil, fmt.Errorf("gagal memetakan transaksi %v", err)
		}
		transactions = append(transactions, t)
	}
	
	if transactions == nil {
		transactions = []models.Transaction{}
	}

	return transactions, nil
}

func (r *TransactionRepository) CreateTransaction(ctx context.Context, t models.Transaction) error {
	if t.Date.IsZero() {
		t.Date = time.Now()
	}

	query := "INSERT INTO transactions (deskripsi, jumlah, tanggal, jenis, catatan) VALUES (?, ?, ?, ?, ?)"
	_, err := r.DB.ExecContext(ctx, query, t.Description, t.Amount, t.Date, t.Jenis, t.Catatan)
	if err != nil {
		return fmt.Errorf("gagal menambahkan transaksi: %v", err)
	}
	return nil
}

func (r *TransactionRepository) UpdateTransaction(ctx context.Context, id int, t models.Transaction) error {
	query := "UPDATE transactions SET deskripsi = ?, jumlah = ?, tanggal = ?, jenis = ?, catatan = ? WHERE transactions_id = ?"
	result, err := r.DB.ExecContext(ctx, query, t.Description, t.Amount, t.Date, t.Jenis, t.Catatan, id)
	if err != nil {
		return fmt.Errorf("gagal mengubah transaksi: %v", err)
	}

	affected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if affected == 0 {
		return fmt.Errorf("transaksi dengan id %d tidak ditemukan", id)
	}

	return nil
}

func (r *TransactionRepository) DeleteTransaction(ctx context.Context, id int) error {
	query := "DELETE FROM transactions WHERE transactions_id = ?"
	result, err := r.DB.ExecContext(ctx, query, id)
	if err != nil {
		return fmt.Errorf("gagal menghapus transaksi: %v", err)
	}
	
	affected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if affected == 0 {
		return fmt.Errorf("transaksi dengan id %d tidak ditemukan", id)
	}

	return nil
}
