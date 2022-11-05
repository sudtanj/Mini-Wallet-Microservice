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

type DepositOutput struct {
	DepositedBy string    `json:"deposited_by"`
	Id          uuid.UUID `json:"id"`
	Status      string    `json:"status"`
	DepositedAt time.Time `json:"deposited_at"`
	Amount      uint      `json:"amount"`
	ReferenceId string    `json:"reference_id"`
}

type WithdrawOutput struct {
	WithdrawnBy string    `json:"withdrawn_by"`
	Id          uuid.UUID `json:"id"`
	Status      string    `json:"status"`
	WithdrawnAt time.Time `json:"withdrawn_at"`
	Amount      uint      `json:"amount"`
	ReferenceId string    `json:"reference_id"`
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

func ViewWallet(db *gorm.DB, token string) (*WalletOutput, error) {
	var customerData UserWalletModel
	result := db.First(&customerData, "token = ?", token)
	if result.Error != nil {
		return nil, result.Error
	}

	if customerData.Status == "disabled" {
		return nil, errors.New("Disabled")
	}

	return &WalletOutput{
		Id:        customerData.Id,
		OwnedBy:   customerData.OwnedBy,
		Status:    customerData.Status,
		EnabledAt: customerData.EnabledAt,
		Balance:   customerData.Balance,
	}, nil
}

func AddMoney(db *gorm.DB, token string, amount uint, refId string) (*DepositOutput, error) {
	tx := db.Begin()

	var customerData UserWalletModel
	result := tx.First(&customerData, "token = ?", token)
	if result.Error != nil {
		return nil, result.Error
	}

	if customerData.Status == "disabled" {
		return nil, errors.New("Disabled")
	}

	var depositedWallet UserWalletDepositedModel
	res := tx.First(&depositedWallet, "reference_id = ?", refId)
	if res.Error == nil {
		return nil, errors.New("Duplicate reference id!")
	}

	customerData.Balance += amount
	tx.Save(&customerData)

	newDeposit := UserWalletDepositedModel{
		DepositedBy: customerData.OwnedBy,
		Id:          uuid.New(),
		Status:      "success",
		DepositedAt: time.Now(),
		Amount:      amount,
		ReferenceId: refId,
	}

	tx.Create(&newDeposit)

	tx.Commit()

	return &DepositOutput{
		DepositedBy: newDeposit.DepositedBy,
		Id:          newDeposit.Id,
		Status:      newDeposit.Status,
		DepositedAt: newDeposit.DepositedAt,
		Amount:      newDeposit.Amount,
		ReferenceId: newDeposit.ReferenceId,
	}, nil
}

func WithdrawMoney(db *gorm.DB, token string, amount uint, refId string) (*WithdrawOutput, error) {
	tx := db.Begin()

	var customerData UserWalletModel
	result := tx.First(&customerData, "token = ?", token)
	if result.Error != nil {
		return nil, result.Error
	}

	if customerData.Status == "disabled" {
		return nil, errors.New("Disabled")
	}

	var withdrawWallet UserWalletWithdrawModel
	res := tx.First(&withdrawWallet, "reference_id = ?", refId)
	if res.Error == nil {
		return nil, errors.New("Duplicate reference id!")
	}

	customerData.Balance -= amount
	if customerData.Balance < 0 {
		return nil, errors.New("Cannot withdraw due to insufficient balance!")
	}

	tx.Save(&customerData)

	newWithdraw := UserWalletWithdrawModel{
		WithdrawnBy: customerData.OwnedBy,
		Id:          uuid.New(),
		Status:      "success",
		WithdrawnAt: time.Now(),
		Amount:      amount,
		ReferenceId: refId,
	}

	tx.Create(&newWithdraw)

	tx.Commit()

	return &WithdrawOutput{
		WithdrawnBy: newWithdraw.WithdrawnBy,
		Id:          newWithdraw.Id,
		Status:      newWithdraw.Status,
		WithdrawnAt: newWithdraw.WithdrawnAt,
		Amount:      newWithdraw.Amount,
		ReferenceId: newWithdraw.ReferenceId,
	}, nil
}
