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

	// INI ADALAH BAGIAN READ
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		transactions, _ := repo.GetAllTransactions(context.Background())
		tmpl.Execute(w, transactions)
	})

	// INI ADALAH BAGIAN CREATE
	http.HandleFunc("/add", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			amount, _ := strconv.ParseFloat(r.FormValue("amount"), 64)
			repo.Create(context.Background(), models.Transaction{
				Description: r.FormValue("description"),
				Amount:      amount,
				Date:        time.Now(),
			})
		}
		http.Redirect(w, r, "/", http.StatusSeeOther)
	})

	fmt.Println("🚀 Server menyala port 8080. Buka http://localhost:8080/ di browser Anda.")
	log.Fatal(http.ListenAndServe(":8080", nil))
}