package main

import (
	"fmt"
	"laporan_keuangan/database"
	"net/http" 
	"log"
)

func main() {
	// Inisialisasi Database Pool (Pangkalan Taksi)
	db, err := database.GetConnection()
	if err != nil {
		log.Fatalf("Gagal inisialisasi database: %v", err)
	}

	// Untuk Menutup kembali setelah Database setelah tidak digunakan
	defer db.Close()

	// Membuat "Resepsionis" (Handler) untuk alamat localhost:8080/

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Setiap kali ada yang buka browser, kita "Ping" database-nya
		err := db.Ping()
		
		if err != nil {
			// Jika gagal, tampilkan pesan error di browser dengan warna merah (HTML)
			w.Header().Set("Content-Type", "text/html")
			fmt.Fprintf(w, "<h1 style='color:red;'>Koneksi Gagal!</h1>")
			fmt.Fprintf(w, "<p>Error: %v</p>", err)
			return
		}

		// Jika berhasil, tampilkan pesan sukses
		w.Header().Set("Content-Type", "text/html")
		fmt.Fprintf(w, "<h1 style='color:green;'>Koneksi Berhasil!</h1>")
		
		fmt.Fprintf(w, "<p>Selamat datang di database 'laporan_keuangan' aman terkendali.</p>")
		
		// Opsional: Tampilkan jumlah user
		var count int
		db.QueryRow("SELECT COUNT(*) FROM users").Scan(&count)
		fmt.Fprintf(w, "<p>Jumlah data di tabel users: <b>%d</b></p>", count)
	})

	// Menjalankan server di Port 8080
	fmt.Println("Koneksi Berhasil")
	fmt.Println("Server berjalan di http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}