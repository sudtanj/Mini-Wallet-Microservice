package main

import (
	"Mini-Wallet-Microservice/wallet"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
)

func main() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	db.AutoMigrate(&wallet.UserWalletModel{})

	r := gin.Default()

	v1 := r.Group("/api/v1")
	v1.POST("/init", func(context *gin.Context) {
		customerIdx := context.PostFormArray("customer_xid")
		if len(customerIdx) == 0 || customerIdx[0] == "" {
			context.JSON(400, gin.H{
				"status": "failed",
				"error": gin.H{
					"customer_xid": "Missing data for required field.",
				},
			})
			return
		}
		result := wallet.InitWallet(db, customerIdx[0])

		context.JSON(201, gin.H{
			"status": "success",
			"data":   result,
		})
		return
	})

	runErr := r.Run(":80")
	if runErr != nil {
		log.Fatalf("Failed to run server: %v", runErr)
	}
}
