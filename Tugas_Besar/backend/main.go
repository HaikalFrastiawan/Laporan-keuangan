package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/HaikalFrastiawan/Laporan-keuangan/backend/database"
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

	log.Println("Backend menyala di Port 8080...")
	r.Run(":8080")
}
