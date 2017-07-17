package model

import (
	"time"

	"github.com/CANARIA/canaria-api/core/util"

	"fmt"

	"github.com/jinzhu/gorm"
)

type (
	Account struct {
		MailAddress string    `datastore:"-" goon:"id"`
		//UserId      int64     `datastore:"-" goon:"parent"`

		UserName    string    `datastore:"user_name"`
		Password    string    `datastore:"password"`
		Roll        int8      `datastore:"roll"`
		CreatedAt   time.Time `datastore:"created_at"`
		UpdatedAt   time.Time `datastore:"updated_at"`
		IsDeleted   bool      `datastore:"is_deleted"`
	}
	accountDao struct {
		*gorm.DB
	}

	AccountDao interface {
		Dao
		Create(*Account) error
	}
)

func AccountDaoFactory(db *gorm.DB) AccountDao {
	return &accountDao{
		DB: db,
	}
}

//--------------------------------------------
// Implementations for Dao interface
//--------------------------------------------

func (dao *accountDao) table() *gorm.DB {
	return dao.Table("accounts")
}

//--------------------------------------------
// Implementations for Model interface
//--------------------------------------------

func AccountImpl(authRegister *AuthRegister, preAccount *PreAccount) *Account {
	return &Account{
		//UserId:      0,
		UserName:    authRegister.UserName,
		MailAddress: preAccount.MailAddress,
		Password:    util.ToCrypt(authRegister.Password),
	}
}

func (dao *accountDao) Create(account *Account) error {

	if res := dao.table().Save(&account); res.Error != nil {
		return fmt.Errorf("failed account create{%v}", account)
	}

	return nil
}
