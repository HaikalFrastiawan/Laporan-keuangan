package models

import "time"

// Transaction merepresentasikan struktur data dari tabel 'transactions'
type Transaction struct {
	ID          int       `json:"transactions_id"`
	Description string    `json:"deskripsi"`
	Amount      float64   `json:"jumlah"`
	Date        time.Time `json:"tanggal"`
	Jenis       string    `json:"jenis"` // 'pemasukan' atau 'pengeluaran'
	Catatan     string    `json:"catatan"`
}

// DashboardResponse adalah format JSON balasan dari API Dashboard
type DashboardResponse struct {
	TotalPemasukan     float64       `json:"total_pemasukan"`
	TotalPengeluaran   float64       `json:"total_pengeluaran"`
	SaldoAkhir         float64       `json:"saldo_akhir"`
	RecentTransactions []Transaction `json:"recent_transactions"`
}
