package models

import "time"

// Transaction merepresentasikan struktur data untuk tabel transactions.
// Digunakan untuk data mapping dari row database ke objek Go.
type Transaction struct {
	ID          int       `json:"transactions_id"`
	Description string    `json:"deskripsi"`
	Amount      float64   `json:"jumlah"`
	Date        time.Time `json:"tanggal"`
}
