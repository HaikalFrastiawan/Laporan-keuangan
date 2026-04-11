package database

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql" // Driver sebagai "Kendaraan" resmi
)

// GetConnection membangun pangkalan taksi (Connection Pool) 
func GetConnection() (*sql.DB, error) {
	// Komposisi Connection String (Paspor Sakti)
	// Format: user:password@tcp(host:port)/dbname
	dsn := "root:@tcp(127.0.0.1:3306)/laporan_keuangan?parseTime=true"

	// Membuka jembatan (Inisialisasi)
	// sql.Open tidak langsung mengoneksi ke DB, hanya memvalidasi argumen DSN
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, fmt.Errorf("gagal inisialisasi driver: %v", err)
	}

	// --- MANAJEMEN KONEKSI (Connection Pooling) ---

	// SetMaxIdleConns: Menentukan jumlah taksi yang stand-by di pangkalan
	db.SetMaxIdleConns(10)

	// SetMaxOpenConns: Membatasi jumlah total taksi agar gedung tidak roboh (beban berlebih)
	db.SetMaxOpenConns(100)

	// SetConnMaxIdleTime: Jika taksi menganggur terlalu lama, pulangkan ke garasi
	db.SetConnMaxIdleTime(5 * time.Minute)

	// SetConnMaxLifetime: Masa pensiun taksi (mencegah koneksi 'basi'/stale)
	db.SetConnMaxLifetime(60 * time.Minute)

	return db, nil
}