package model

import (
	"time"

	"github.com/CANARIA/canaria-api/util"

	"fmt"

	"github.com/jinzhu/gorm"
)

type Account struct {
	UserId      int64     `gorm:"column:user_id"`
	UserName    string    `gorm:"column:user_name"`
	MailAddress string    `gorm:"column:mailaddress"`
	Password    string    `gorm:"column:password"`
	Roll        int8      `gorm:"column:roll"`
	CreatedAt   time.Time `gorm:"column:created_at"`
	UpdatedAt   time.Time `gorm:"column:updated_at"`
	IsDeleted   bool      `gorm:"column:is_deleted"`
}

func AccountImpl(authRegister *AuthRegister) *Account {
	return &Account{
		UserId:      0,
		UserName:    authRegister.UserName,
		MailAddress: authRegister.MailAddress,
		Password:    util.ToHash(authRegister.Password),
	}
}

func (account *Account) AccountCreate(tx *gorm.DB) error {

	if res := tx.Table("accounts").Create(account); res.Error != nil {
		return fmt.Errorf("failed account create{%v}", *account)
	}

	// _, err := tx.InsertInto("accounts").
	// 	Columns("user_name", "mailaddress", "password").
	// 	Record(account).
	// 	Exec()

	return nil
}
