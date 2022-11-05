package main

import (
	"Mini-Wallet-Microservice/wallet"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
	"strings"
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
				"status": "fail",
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

	v1.POST("/wallet", func(context *gin.Context) {
		token := strings.Split(context.Request.Header["Authorization"][0], " ")[1]

		if len(token) == 0 {
			context.JSON(400, gin.H{
				"status": "fail",
				"data": gin.H{
					"error": "Invalid authorization token!",
				},
			})
			return
		}

		result, err := wallet.EnabledWallet(db, token)
		if err != nil {
			context.JSON(400, gin.H{
				"status": "fail",
				"data": gin.H{
					"error": err.Error(),
				},
			})
			return
		}

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
