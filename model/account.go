package model

import (
	"time"

	"github.com/CANARIA/canaria-api/util"

	"github.com/gocraft/dbr"
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

func AccountImpl(authRegister *AuthRegister) *Account {
	return &Account{
		UserId:      0,
		UserName:    authRegister.UserName,
		MailAddress: authRegister.MailAddress,
		Password:    util.ToHash(authRegister.Password),
	}
}

func (account *Account) AccountCreate(tx *dbr.Tx) error {
	_, err := tx.InsertInto("accounts").
		Columns("user_name", "mailaddress", "password").
		Record(account).
		Exec()

	return err
}
