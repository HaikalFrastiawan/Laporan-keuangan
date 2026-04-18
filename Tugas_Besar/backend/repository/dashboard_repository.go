package repository

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/HaikalFrastiawan/Laporan-keuangan/backend/models"
)

type DashboardRepository struct {
	DB *sql.DB
}

func NewDashboardRepository(db *sql.DB) *DashboardRepository {
	return &DashboardRepository{DB: db}
}

func (r *DashboardRepository) GetDashboardData(ctx context.Context) (*models.DashboardResponse, error) {
	var response models.DashboardResponse

	// Ambil Total Pemasukan
	pemasukanQuery := "SELECT COALESCE(SUM(jumlah), 0) FROM transactions WHERE jenis = 'pemasukan'"
	err := r.DB.QueryRowContext(ctx, pemasukanQuery).Scan(&response.TotalPemasukan)
	if err != nil {
		return nil, fmt.Errorf("gagal menghitung pemasukan: %v", err)
	}

	// Ambil Total Pengeluaran
	pengeluaranQuery := "SELECT COALESCE(SUM(jumlah), 0) FROM transactions WHERE jenis = 'pengeluaran'"
	err = r.DB.QueryRowContext(ctx, pengeluaranQuery).Scan(&response.TotalPengeluaran)
	if err != nil {
		return nil, fmt.Errorf("gagal menghitung pengeluaran: %v", err)
	}

	response.SaldoAkhir = response.TotalPemasukan - response.TotalPengeluaran

	// Ambil 5 Transaksi Terbaru
	recentQuery := "SELECT transactions_id, deskripsi, jumlah, tanggal, jenis, catatan FROM transactions ORDER BY tanggal DESC LIMIT 5"
	rows, err := r.DB.QueryContext(ctx, recentQuery)
	if err != nil {
		return nil, fmt.Errorf("gagal mengambil data transaksi: %v", err)
	}
	defer rows.Close()

	response.RecentTransactions = make([]models.Transaction, 0)
	for rows.Next() {
		var t models.Transaction
		if err := rows.Scan(&t.ID, &t.Description, &t.Amount, &t.Date, &t.Jenis, &t.Catatan); err != nil {
			return nil, fmt.Errorf("gagal memetakan transaksi %v", err)
		}
		response.RecentTransactions = append(response.RecentTransactions, t)
	}

	return &response, nil
}
