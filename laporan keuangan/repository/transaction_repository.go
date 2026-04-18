package repository

import (
	"context"
	"database/sql"
	"laporan_keuangan/models"
)

// TransactionRepository menangani operasi database tabel transactions.
type TransactionRepository struct {
	DB *sql.DB
}

func NewTransactionRepository(db *sql.DB) *TransactionRepository {
	return &TransactionRepository{DB: db}
}

// CREATE fungsi untuk menambahkan transaksi baru ke database.
func (r *TransactionRepository) Create(ctx context.Context, t models.Transaction) error {
	query := "INSERT INTO transactions (deskripsi, jumlah, tanggal, jenis, catatan) VALUES (?, ?, ?, 'pengeluaran', '')"
	_, err := r.DB.ExecContext(ctx, query, t.Description, t.Amount, t.Date)
	return err
}

// READ fungsi untuk mengambil semua transaksi dari database.
func (r *TransactionRepository) GetAllTransactions(ctx context.Context) ([]models.Transaction, error) {
	rows, err := r.DB.QueryContext(ctx, "SELECT transactions_id, deskripsi, jumlah, tanggal FROM transactions")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var transactions []models.Transaction
	for rows.Next() {
		var t models.Transaction
		if err := rows.Scan(&t.ID, &t.Description, &t.Amount, &t.Date); err == nil {
			transactions = append(transactions, t)
		}
	}
	return transactions, nil
}
