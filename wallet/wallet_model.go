package wallet

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type UserWalletModel struct {
	gorm.Model
	OwnedBy   string
	Token     string
	Id        uuid.UUID
	Status    string
	EnabledAt time.Time
	Balance   uint
}
