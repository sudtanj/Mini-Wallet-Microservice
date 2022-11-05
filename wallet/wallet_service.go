package wallet

import (
	"errors"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type InitWalletReturn struct {
	Token string `json:"token"`
}

type WalletOutput struct {
	Id        uuid.UUID `json:"id"`
	OwnedBy   string    `json:"owned_by"`
	Status    string    `json:"status"`
	EnabledAt time.Time `json:"enabled_at"`
	Balance   uint      `json:"balance"`
}

func InitWallet(db *gorm.DB, customerIdx string) *InitWalletReturn {
	var customerData UserWalletModel
	result := db.First(&customerData, "owned_by = ?", customerIdx)
	if result.Error != nil {
		// transaction is use to prevent racing condition while find duplicate token in database
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

func EnabledWallet(db *gorm.DB, token string) (*WalletOutput, error) {
	var customerData UserWalletModel
	result := db.First(&customerData, "token = ?", token)
	if result.Error != nil {
		return nil, result.Error
	}

	if customerData.Status == "enabled" {
		return nil, errors.New("Already enabled")
	}

	customerData.Status = "enabled"
	customerData.EnabledAt = time.Now()

	db.Save(&customerData)

	return &WalletOutput{
		Id:        customerData.Id,
		OwnedBy:   customerData.OwnedBy,
		Status:    customerData.Status,
		EnabledAt: customerData.EnabledAt,
		Balance:   customerData.Balance,
	}, nil
}
