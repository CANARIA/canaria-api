package model

import (
	"fmt"
	"time"

	"github.com/CANARIA/canaria-api/config"
	"github.com/jinzhu/gorm"
)

var preAccount PreAccount

var auth Auth

type (
	Auth struct {
		UrlToken    *string
		MailAddress *string
	}

	PreAccount struct {
		Id          int64      `gorm:"column:id"`
		UrlToken    string     `gorm:"column:url_token"`
		CreatedAt   *time.Time `gorm:"column:created_at"`
		MailAddress string     `gorm:"column:mailaddress"`
		IsRegisterd bool       `gorm:"column:is_registered"`
	}

	preAccountDao struct {
		*gorm.DB
	}

	PreAccountDao interface {
		Dao
		AcctivateAccount(*PreAccount) error
		ValidPreAccountToken(*Auth) (*PreAccount, error)
	}
)

func PreAccountDaoFactory(db *gorm.DB) PreAccountDao {
	return &preAccountDao{
		DB: db,
	}
}

//--------------------------------------------
// Implementations for Dao interface
//--------------------------------------------

func (dao *preAccountDao) table() *gorm.DB {
	return dao.Table("pre_accounts")
}

//--------------------------------------------
// Implementations for Model interface
//--------------------------------------------

func PreAccountImpl(preRegister *PreRegister, token string) *PreAccount {
	return &PreAccount{
		Id:          0,
		UrlToken:    token,
		MailAddress: preRegister.MailAddress,
	}
}

func (preAccount *PreAccount) PreAccountCreate(tx *gorm.DB) error {
	if res := tx.Create(preAccount); res.Error != nil {
		return fmt.Errorf("failed PreAccount create: %s", res.Error.Error())
	}

	return nil
}

func BuildRegisterUrl(token string) string {
	url := "http://" + config.GetHost() + "/" + "register?register_token=" + token
	return url
}

func BuildPreAccountEntity(auth *Auth, preAccount *PreAccount) *PreAccount {
	return &PreAccount{
		UrlToken:    *auth.UrlToken,
		MailAddress: preAccount.MailAddress,
	}
}

/*
仮登録情報を本登録に更新
*/
func (dao *preAccountDao) AcctivateAccount(preAccountRow *PreAccount) error {

	preAccount := &PreAccount{IsRegisterd: true}

	res := dao.table().
		Where("url_token = ? AND mailaddress = ?", preAccountRow.UrlToken, &preAccountRow.MailAddress).
		Update(preAccount)

	if res.Error != nil {
		return fmt.Errorf("failed PreAccount update: %s", res.Error.Error())
	}

	return nil
}

/*
仮登録情報の妥当性確認
*/
func (dao *preAccountDao) ValidPreAccountToken(auth *Auth) (*PreAccount, error) {

	q := dao.table().
		Where("url_token = ? AND is_registered = 0 AND created_at > now() - interval 24 hour", *auth.UrlToken)

	if res := q.Find(&preAccount); res.Error != nil {
		return nil, fmt.Errorf("invalid token: %s", res.Error)
	}

	return &preAccount, nil
}
