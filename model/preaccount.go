package model

import (
	"time"

	"errors"

	"github.com/CANARIA/canaria-api/config"
	"github.com/gocraft/dbr"
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

func (preAccount *PreAccount) PreAccountCreate(tx *dbr.Tx) error {
	_, err := tx.InsertInto("pre_accounts").
		Columns("url_token", "created_at", "mailaddress").
		Record(preAccount).
		Exec()

	return err
}

func BuildRegisterUrl(token string) string {
	url := "http://" + config.GetHost() + "/" + "auth/register?register_token=" + token
	return url
}

func (auth *Auth) ValidPreAccountToken(tx *dbr.Tx) (PreAccount, error) {

	tx.Select("*").
		From("pre_accounts").
		Where("url_token = ? AND is_registered = ? AND created_at > now() - interval 24 hour", auth.UrlToken, 0).
		Load(&preAccount)

	if (PreAccount{}) == preAccount {
		return PreAccount{}, errors.New("invalid token: " + auth.UrlToken)
	}
	return preAccount, nil
}
