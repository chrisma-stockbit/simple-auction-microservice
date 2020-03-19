package entity

import (
	"database/sql"
	"time"

	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

type Bidder struct {
	ID            uint
	CreatedAt     time.Time
	Guid          string
	FullName      string
	AccountNumber string
	BankName      string
	PhoneNumber   string
	Balance       sql.NullInt64
	IsActive      bool
}

func NewActiveBidder() *Bidder {
	return &Bidder{IsActive: true}
}

func (b *Bidder) BeforeCreate(scope *gorm.Scope) error {
	uid := uuid.NewV4().String()
	return scope.SetColumn("Guid", uid)
}

func (b Bidder) TableName() string {
	return "bidder"
}
