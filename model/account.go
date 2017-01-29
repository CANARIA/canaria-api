package model

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Account struct {
	UserId      int64     `db:"user_id";gorm:"AUTO_INCREMENT;primary_key"`
	UserName    string    `db:"user_name"`
	MailAddress string    `db:"mailaddress"`
	Password    string    `db:"password"`
	Roll        int8      `db:"roll"`
	CreatedAt   time.Time `db:"created_at"`
	UpdatedAt   time.Time `db:"updated_at"`
	IsDeleted   bool      `db:"is_deleted"`
}

func AccountCreate(authRegister *AuthRegister, db *gorm.DB) {
	db.Create(authRegister)
}
