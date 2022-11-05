package wallet

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type UserWalletWithdrawModel struct {
	gorm.Model
	WithdrawnBy string
	Id          uuid.UUID
	Status      string
	WithdrawnAt time.Time
	Amount      uint
	ReferenceId string
}
