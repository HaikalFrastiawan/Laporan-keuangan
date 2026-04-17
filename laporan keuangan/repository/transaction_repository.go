package repository

import (
	"context"
	"database/sql"
	"fmt"
	"laporan_keuangan/models"
)

// TransactionRepository menangani operasi database tabel transactions.
type TransactionRepository struct {
	DB *sql.DB
}

func NewTransactionRepository(db *sql.DB) *TransactionRepository {
	return &TransactionRepository{DB: db}
}

// Upsert melakukan Insert jika deskripsi belum ada, atau Update akumulasi nilai jika sudah ada.
func (r *TransactionRepository) Upsert(ctx context.Context, t models.Transaction) error {
	checkQuery := "SELECT transactions_id, jumlah FROM transactions WHERE deskripsi = ?"
	
	stmtCheck, err := r.DB.PrepareContext(ctx, checkQuery)
	if err != nil {
		return fmt.Errorf("gagal menyiapkan statement select: %v", err)
	}
	defer stmtCheck.Close()

	var existingID int
	var existingAmount float64
	
	err = stmtCheck.QueryRowContext(ctx, t.Description).Scan(&existingID, &existingAmount)
	
	if err == sql.ErrNoRows {
		// Logika Insert untuk data baru
		insertQuery := "INSERT INTO transactions (deskripsi, jumlah, tanggal, jenis, catatan) VALUES (?, ?, ?, 'pengeluaran', '')"
		stmtInsert, err := r.DB.PrepareContext(ctx, insertQuery)
		if err != nil {
			return fmt.Errorf("gagal menyiapkan statement insert: %v", err)
		}
		defer stmtInsert.Close()
		
		_, err = stmtInsert.ExecContext(ctx, t.Description, t.Amount, t.Date)
		if err != nil {
			return fmt.Errorf("gagal mengeksekusi insert: %v", err)
		}
	} else if err != nil {
		return fmt.Errorf("gagal mengecek transaksi: %v", err)
	} else {
		// Logika Update untuk akumulasi amount
		newAmount := existingAmount + t.Amount
		updateQuery := "UPDATE transactions SET jumlah = ? WHERE transactions_id = ?"
		stmtUpdate, err := r.DB.PrepareContext(ctx, updateQuery)
		if err != nil {
			return fmt.Errorf("gagal menyiapkan statement update: %v", err)
		}
		defer stmtUpdate.Close()
		
		_, err = stmtUpdate.ExecContext(ctx, newAmount, existingID)
		if err != nil {
			return fmt.Errorf("gagal mengeksekusi update: %v", err)
		}
	}
	
	return nil
}

// GetAllTransactions mengambil seluruh data riwayat transaksi menggunakan prepared statement.
func (r *TransactionRepository) GetAllTransactions(ctx context.Context) ([]models.Transaction, error) {
	query := "SELECT transactions_id, deskripsi, jumlah, tanggal FROM transactions"
	
	stmt, err := r.DB.PrepareContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("gagal menyiapkan statement select: %v", err)
	}
	defer stmt.Close()

	rows, err := stmt.QueryContext(ctx)
	if err != nil {
		return nil, fmt.Errorf("gagal menjalankan query: %v", err)
	}
	defer rows.Close()

	transactions := make([]models.Transaction, 0)
	
	for rows.Next() {
		var t models.Transaction
		if err := rows.Scan(&t.ID, &t.Description, &t.Amount, &t.Date); err != nil {
			return nil, fmt.Errorf("gagal memetakan data: %v", err)
		}
		transactions = append(transactions, t)
	}
	
	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterasi data: %v", err)
	}
	
	return transactions, nil
}
