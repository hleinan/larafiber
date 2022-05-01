package models

type Task struct {
	Base
	UserID int64  `json:"-"`
	User   *User  `gorm:"foreignkey:UserID" json:"-"`
	Text   string `gorm:"type:string;" json:"text"`
	Done   bool   `gorm:"type:bool; default:false" json:"done"`
}
