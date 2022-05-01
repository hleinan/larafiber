package models

type User struct {
	Base
	Name      string   `gorm:"type:string;" json:"name" form:"name"`
	Avatar    string   `gorm:"type:string;" json:"avatar" form:"avatar"`
	Admin     bool     `gorm:"type:bool; default:false" json:"admin"`
	Verified  bool     `gorm:"type:bool; default:false" json:"verified"`
	Marketing bool     `gorm:"type:bool; default:false" json:"marketing"`
	Account   *Account `gorm:"foreignkey:UserID" json:"account"`
	AddressID int64    `gorm:"type:int64" json:"-"`
}

var user User

func GetUser(id interface{}) User {
	db := GetDB()
	user := User{}
	iId, ok := id.(int64)

	if ok {
		db.
			Preload("Account").
			First(&user, iId)
		return user
	}
	return User{}
}
