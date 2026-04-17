package main

import (
	"context"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"time"

	"laporan_keuangan/database"
	"laporan_keuangan/models"
	"laporan_keuangan/repository"
)

func main() {
	db, err := database.GetConnection()
	if err != nil {
		log.Fatalf("Gagal inisialisasi database: %v", err)
	}
	defer db.Close()

	repo := repository.NewTransactionRepository(db)
	tmpl := template.Must(template.ParseFiles("views/index.html"))

	// Halaman Utama: Menampilkan Daftar Transaksi
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			http.NotFound(w, r)
			return
		}

		ctx := context.Background()
		transactions, err := repo.GetAllTransactions(ctx)
		if err != nil {
			http.Error(w, "Gagal mengambil data dari DB: "+err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "text/html")
		tmpl.Execute(w, transactions)
	})

	// Endpoint Form: Menangani aksi Create/Update Transaksi
	http.HandleFunc("/add", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Redirect(w, r, "/", http.StatusSeeOther)
			return
		}

		if err := r.ParseForm(); err != nil {
			http.Error(w, "Gagal memproses isian form", http.StatusBadRequest)
			return
		}

		description := r.FormValue("description")
		amountStr := r.FormValue("amount")
		
		amount, err := strconv.ParseFloat(amountStr, 64)
		if err != nil {
			http.Error(w, "Format jumlah angka tidak valid", http.StatusBadRequest)
			return
		}

		tx := models.Transaction{
			Description: description,
			Amount:      amount,
			Date:        time.Now(),
		}

		ctx := context.Background()
		if err = repo.Upsert(ctx, tx); err != nil {
			http.Error(w, "Gagal menyimpan transaksi: "+err.Error(), http.StatusInternalServerError)
			return
		}

		http.Redirect(w, r, "/", http.StatusSeeOther)
	})

	fmt.Println("🚀 Server menyala port 8080. Buka http://localhost:8080/ di browser Anda.")
	log.Fatal(http.ListenAndServe(":8080", nil))
}