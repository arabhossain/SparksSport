package models

type Admin struct {
	ID          int64  `gorm:"primaryKey;autoIncrement" json:"id"`
	FirstName   string `gorm:"size:100;not null" json:"first_name"`
	LastName    string `gorm:"size:100;not null" json:"last_name"`
	AccountType string `gorm:"size:50;not null" json:"account_type"`
	Email       string `gorm:"size:100;unique;not null" json:"email"`
	Password    string `gorm:"size:255;not null" json:"password"`
	LastLogin   int64  `gorm:"default:0" json:"last_login"`
	IsActive    bool   `gorm:"default:true" json:"is_active"`
}
