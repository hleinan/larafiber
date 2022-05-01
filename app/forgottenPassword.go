package models

type ForgottenPassword struct {
	Base
	AccountID int64   `gorm:"unique;not null" json:"-"`
	Account   Account `gorm:"foreignkey:AccountID" json:"account"`
}
