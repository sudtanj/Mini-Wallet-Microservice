package wallet

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type UserWalletDepositedModel struct {
	gorm.Model
	DepositedBy string
	Id          uuid.UUID
	Status      string
	DepositedAt time.Time
	Amount      uint
	ReferenceId string
}
