package main

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/HaikalFrastiawan/Laporan-keuangan/backend/database"
	"github.com/HaikalFrastiawan/Laporan-keuangan/backend/models"
	"github.com/HaikalFrastiawan/Laporan-keuangan/backend/repository"
)

// CORSMiddleware untuk mengizinkan akses dari Frontend React (Port 5173)
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func main() {
	// Inisialisasi Database
	db, err := database.GetConnection()
	if err != nil {
		log.Fatalf("Gagal inisialisasi database: %v", err)
	}
	defer db.Close()

	// Inisialisasi Repositori
	repo := repository.NewDashboardRepository(db)
	transactionRepo := repository.NewTransactionRepository(db)

	// Inisialisasi Router Gin
	r := gin.Default()
	r.Use(CORSMiddleware())

	// Endpoint API Dasar
	r.GET("/api/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// Endpoint Dashboard
	r.GET("/api/dashboard", func(c *gin.Context) {
		data, err := repo.GetDashboardData(c.Request.Context())
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, data)
	})

	// Endpoint Transactions
	r.GET("/api/transactions", func(c *gin.Context) {
		transactions, err := transactionRepo.GetAllTransactions(c.Request.Context())
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, transactions)
	})

	r.POST("/api/transactions", func(c *gin.Context) {
		var req models.Transaction
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
			return
		}
		if err := transactionRepo.CreateTransaction(c.Request.Context(), req); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusCreated, gin.H{"message": "Transaction created successfully"})
	})

	r.PUT("/api/transactions/:id", func(c *gin.Context) {
		idParam := c.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
			return
		}
		var req models.Transaction
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
			return
		}
		if err := transactionRepo.UpdateTransaction(c.Request.Context(), id, req); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "Transaction updated successfully"})
	})

	r.DELETE("/api/transactions/:id", func(c *gin.Context) {
		idParam := c.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
			return
		}
		if err := transactionRepo.DeleteTransaction(c.Request.Context(), id); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "Transaction deleted successfully"})
	})

	log.Println("Backend menyala di Port 8080...")
	r.Run(":8080")
}
