package wallet

import (
	"fmt"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type InitWalletReturn struct {
	Token string `json:"token"`
}

func InitWallet(db *gorm.DB, customerIdx string) *InitWalletReturn {
	var customerData UserWalletModel
	result := db.First(&customerData, "owned_by = ?", customerIdx)
	fmt.Println(result.Error)
	if result.Error != nil {
		tx := db.Begin()
		// use to generate random token for auth
		// token := GenerateSecureToken(21)
		// hard coded for this example purposes
		token := "6b3f7dc70abe8aed3e56658b86fa508b472bf238"
		tx.Create(&UserWalletModel{
			OwnedBy:   customerIdx,
			Token:     token,
			Id:        uuid.New(),
			Status:    "enabled",
			EnabledAt: time.Now(),
			Balance:   0,
		})

		tx.Commit()

		return &InitWalletReturn{
			Token: token,
		}
	}

	return &InitWalletReturn{
		Token: customerData.Token,
	}
}
