package model

import (
	"time"

	"github.com/CANARIA/canaria-api/util"

	"fmt"

	"github.com/jinzhu/gorm"
)

type (
	Account struct {
		UserId      int64     `gorm:"column:user_id"`
		UserName    string    `gorm:"column:user_name"`
		MailAddress string    `gorm:"column:mailaddress"`
		Password    string    `gorm:"column:password"`
		Roll        int8      `gorm:"column:roll"`
		CreatedAt   time.Time `gorm:"column:created_at"`
		UpdatedAt   time.Time `gorm:"column:updated_at"`
		IsDeleted   bool      `gorm:"column:is_deleted"`
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
		UserId:      0,
		UserName:    authRegister.UserName,
		MailAddress: preAccount.MailAddress,
		Password:    util.ToCrypt(authRegister.Password),
	}
}

func (dao *accountDao) Create(account *Account) error {

	if res := dao.table().Create(account); res.Error != nil {
		return fmt.Errorf("failed account create{%v}", account)
	}

	return nil
}
