package model

import (
	"fmt"
	"time"

	"errors"

	"github.com/CANARIA/canaria-api/config"
	"github.com/jinzhu/gorm"
)

var preAccount PreAccount

type Auth struct {
	UrlToken    string
	MailAddress string
}

var auth Auth

type PreAccount struct {
	Id          int64     `db:"id"`
	UrlToken    string    `db:"url_token"`
	CreatedAt   time.Time `db:"created_at"`
	MailAddress string    `db:"mailaddress"`
	IsRegisterd bool      `db:"is_registered"`
}

func PreAccountImpl(preRegister *PreRegister, token string) *PreAccount {
	return &PreAccount{
		Id:          0,
		UrlToken:    token,
		CreatedAt:   time.Now(),
		MailAddress: preRegister.MailAddress,
	}
}

func (preAccount *PreAccount) PreAccountCreate(tx *gorm.DB) error {
	// _, err := tx.InsertInto("pre_accounts").
	// 	Columns("url_token", "created_at", "mailaddress").
	// 	Record(preAccount).
	// 	Exec()
	if res := tx.Create(preAccount); res.Error != nil {
		return fmt.Errorf("failed PreAccount create: %s", res.Error.Error())
	}

	return nil
}

func BuildRegisterUrl(token string) string {
	url := "http://" + config.GetHost() + "/" + "register?register_token=" + token
	return url
}

func AcctivateAccount(tx *gorm.DB, auth *Auth) error {
	// _, err := tx.Update("pre_accounts").
	// 	Set("is_registered", true).
	// 	Where("url_token = ? AND mailaddress = ?", auth.UrlToken, auth.MailAddress).
	// 	Exec()

	preAccount := PreAccount{IsRegisterd: true}

	res := tx.
		Model(&preAccount).
		Where("url_token = ? AND mailaddress = ?", auth.UrlToken, auth.MailAddress).
		Update("is_registered")

	if res.Error != nil {
		return fmt.Errorf("failed PreAccount update: %s", res.Error.Error())
	}

	return nil
}

func (auth *Auth) ValidPreAccountToken(tx *gorm.DB) (*PreAccount, error) {

	tx.Select("*").
		Table("pre_accounts").
		Where("url_token = ? AND is_registered = ? AND created_at > now() - interval 24 hour", auth.UrlToken, 0).
		Find(&preAccount)

	if (PreAccount{}) == preAccount {
		return &PreAccount{}, errors.New("invalid token: " + auth.UrlToken)
	}
	return &preAccount, nil
}
