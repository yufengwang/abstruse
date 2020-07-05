package model

// User defines `users` database table.
type User struct {
	ID       uint   `gorm:"primary_key;auto_increment;not null" json:"id"`
	Email    string `gorm:"not null;varchar(255);unique_index" json:"email"`
	Password string `gorm:"not null;varchar(255);column:password" json:"-"`
	Name     string `gorm:"not null;varchar(255)" json:"name"`
	Avatar   string `gorm:"not null;varchar(255);default:'/assets/images/avatars/avatar_1.svg'" json:"avatar"`
	Admin    bool   `gorm:"not null;default:false" json:"admin"`
	TimestampModel
}
