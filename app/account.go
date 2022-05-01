package models

import "time"

type Account struct {
	Base
	Email                string     `gorm:"type:string;" json:"email" form:"email"`
	Password             []byte     `gorm:"type:[]byte;" json:"-" form:"password"`
	Token                *string    `gorm:"type:string;" json:"-"`
	PasswordReset        *string    `gorm:"type:string;" json:"-"`
	PasswordResetTimeout *time.Time `gorm:"type:datetime;" json:"-"`
	UserID               int64      `json:"-"`
	User                 User       `gorm:"foreignkey:UserID; PRELOAD:true" json:"user"`
}
