package models

import (
	"time"

	"gorm.io/gorm"
)

type Company struct {
	gorm.Model
	ID      string `gorm:"primary_key"`
	Name    string
	Address string
	Email   string
}

type Period struct {
	gorm.Model
	ID         string `gorm:"primary_key"`
	PeriodName string
	Start      time.Time
	End        time.Time
	Open       bool
}

type COATemplate struct {
	gorm.Model
	ID     string `gorm:"primary_key"`
	Code   string
	Name   string
	Parent string
}

type COA struct {
	gorm.Model
	ID        string `gorm:"primary_key"`
	Period    string
	PeriodID  Period `gorm:"foreignKey:Period; constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Company   string
	CompanyID Company `gorm:"foreignKey:Company; constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Code      string
	Name      string
	Parent    string
}

type GeneralLedger struct {
	gorm.Model
	ID          string `gorm:"primary_key"`
	Period      string
	PeriodID    Period `gorm:"foreignKey:Period; constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Company     string
	CompanyID   Company `gorm:"foreignKey:Company; constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	COA         string
	COAID       COA `gorm:"foreignKey:COA; constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Description string
	Debet       float64
	Credit      float64
}

type SalesJournal struct {
}

type PurchaseJournal struct {
}

type CashReceiptJournal struct {
}

type CashPaymentJournal struct {
}

type Debtor struct {
	gorm.Model
	ID        string `gorm:"primary_key"`
	Period    string
	PeriodID  Period `gorm:"foreignKey:Period; constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Company   string
	CompanyID Company `gorm:"foreignKey:Company; constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Name      string
	Email     string
	Address   string
}

type Creditor struct {
	gorm.Model
	ID        string `gorm:"primary_key"`
	Period    string
	PeriodID  Period `gorm:"foreignKey:Period; constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Company   string
	CompanyID Company `gorm:"foreignKey:Company; constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Name      string
	Email     string
	Address   string
}
