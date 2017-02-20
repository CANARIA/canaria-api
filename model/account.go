package model

import (
	"time"

	"github.com/CANARIA/canaria-api/util"

	"fmt"

	"github.com/jinzhu/gorm"
)

type Account struct {
	UserId      int64     `db:"user_id"`
	UserName    string    `db:"user_name"`
	MailAddress string    `db:"mailaddress"`
	Password    string    `db:"password"`
	Roll        int8      `db:"roll"`
	CreatedAt   time.Time `db:"created_at"`
	UpdatedAt   time.Time `db:"updated_at"`
	IsDeleted   bool      `db:"is_deleted"`
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

	if res := tx.Create(account); res.Error != nil {
		return fmt.Errorf("failed Account create: %s", res.Error.Error())
	}

	// _, err := tx.InsertInto("accounts").
	// 	Columns("user_name", "mailaddress", "password").
	// 	Record(account).
	// 	Exec()

	return nil
}
